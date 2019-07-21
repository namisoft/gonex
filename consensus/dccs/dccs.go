// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package dccs implements the proof-of-foundation consensus engine.
package dccs

import (
	"bytes"
	"errors"
	"io"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/deployer"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/stable"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/volatile"
	"github.com/ethereum/go-ethereum/contracts/nexty/governance"
	"github.com/ethereum/go-ethereum/contracts/nexty/ntf"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

const (
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

// Dccs proof-of-foundation protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint and reset the pending votes

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures

	rewards = []*big.Int{
		big.NewInt(1e+4),
		big.NewInt(5e+3),
		big.NewInt(25e+2),
		big.NewInt(1250),
		big.NewInt(625),
		big.NewInt(500),
	} // rewards per year in percent of current total supply
	initialSupply = big.NewInt(18e+10)   // initial total supply in NTY
	blockPerYear  = big.NewInt(15778476) // Number of blocks per year with blocktime = 2s
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidPrice is returned if price block contains invalid price value in
	// their extra-data fields.
	errInvalidPrice = errors.New("price block contains invalid price value")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes, or not the correct
	// ones).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 or 2, or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errWaitTransactions is returned if an empty block is attempted to be sealed
	// on an instant chain (0 second period). It's important to refuse these as the
	// block reward is zero, so an empty block just bloats the chain... fast.
	errWaitTransactions = errors.New("waiting for transactions")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")
)

// SignerFn is a signer callback function to request a header to be signed by a
// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// Dccs is the proof-of-foundation consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Dccs struct {
	config *params.DccsConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	proposals map[common.Address]bool // Current list of proposals we are pushing

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields

	priceEngine     *PriceEngine
	priceEngineOnce sync.Once
}

// PriceEngine creates and returns the PriceEngine singleton instance
func (d *Dccs) PriceEngine() *PriceEngine {
	d.priceEngineOnce.Do(func() {
		d.priceEngine = newPriceEngine(d.config)
	})
	return d.priceEngine
}

// New creates a Dccs proof-of-foundation consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.DccsConfig, db ethdb.Database) *Dccs {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	if conf.ThangLongBlock != nil && conf.ThangLongBlock.Sign() > 0 && conf.ThangLongEpoch > 0 {
		// ensure that the hard-fork block must be divisible by both the old and new epoch value
		tlBlock := conf.ThangLongBlock.Uint64()
		if tlBlock%conf.Epoch != 0 {
			log.Crit("Unable to create DCCS consensus engine", "ThangLong block", tlBlock, "Epoch", conf.Epoch)
			return nil
		}
		if tlBlock%conf.ThangLongEpoch != 0 {
			log.Crit("Unable to create DCCS consensus engine", "ThangLong block", tlBlock, "ThangLongEpoch", conf.ThangLongEpoch)
			return nil
		}
	}

	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Dccs{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		proposals:  make(map[common.Address]bool),
	}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (d *Dccs) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, d.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (d *Dccs) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.verifyHeader2(chain, header, nil)
	}
	return d.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (d *Dccs) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			var err error
			if chain.Config().IsThangLong(header.Number) {
				err = d.verifyHeader2(chain, header, headers[:i])
			} else {
				err = d.verifyHeader(chain, header, headers[:i])
			}

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (d *Dccs) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}
	// Checkpoint blocks need to enforce zero beneficiary
	checkpoint := d.config.IsCheckpoint(number)
	if checkpoint && header.Coinbase != (common.Address{}) {
		return errInvalidCheckpointBeneficiary
	}
	// Nonces must be 0x00..0 or 0xff..f, zeroes enforced on checkpoints
	if !bytes.Equal(header.Nonce[:], nonceAuthVote) && !bytes.Equal(header.Nonce[:], nonceDropVote) {
		return errInvalidVote
	}
	if checkpoint && !bytes.Equal(header.Nonce[:], nonceDropVote) {
		return errInvalidCheckpointVote
	}
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !checkpoint && signersBytes != 0 {
		return errExtraSigners
	}
	if checkpoint && signersBytes%common.AddressLength != 0 {
		return errInvalidCheckpointSigners
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return d.verifyCascadingFields(chain, header, parents)
}

// verifyHeader2 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (d *Dccs) verifyHeader2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	checkpoint := d.config.IsCheckpoint(number)
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Only check header's signer list before Endurio hardfork
	if !d.config.IsEndurio(header.Number) {
		// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
		signersBytes := len(header.Extra) - extraVanity - extraSeal
		if !checkpoint && signersBytes != 0 {
			return errExtraSigners
		}
		if checkpoint && signersBytes%common.AddressLength != 0 {
			return errInvalidCheckpointSigners
		}
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return d.verifyCascadingFields2(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (d *Dccs) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+d.config.Period > header.Time {
		return ErrInvalidTimestamp
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// If the block is a checkpoint block, verify the signer list
	if d.config.IsCheckpoint(number) {
		signers := make([]byte, len(snap.Signers)*common.AddressLength)
		for i, signer := range snap.signers() {
			copy(signers[i*common.AddressLength:], signer[:])
		}
		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
			return errInvalidCheckpointSigners
		}
	}
	// All basic checks passed, verify the seal and return
	return d.verifySeal(chain, header, parents)
}

// verifyCascadingFields2 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (d *Dccs) verifyCascadingFields2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+d.config.Period > header.Time {
		return ErrInvalidTimestamp
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot2(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// Stop recording signers list in checkpoint after Endurio hardfork
	if !d.config.IsEndurio(header.Number) {
		// If the block is a checkpoint block, verify the signer list
		if d.config.IsCheckpoint(number) {
			signers := make([]byte, len(snap.Signers)*common.AddressLength)
			for i, signer := range snap.signers2() {
				copy(signers[i*common.AddressLength:], signer.Address[:])
			}
			// for checkpoint: extra = [vanity(32), signers(...), signature(65)]
			extraSuffix := len(header.Extra) - extraSeal
			if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
				return errInvalidCheckpointSigners
			}
		}
	} else if d.config.IsPriceBlock(number) {
		// for price block: extra = [vanity(32), price(...), signature(65)]
		price := PriceDecodeFromExtra(header.Extra)
		if price == nil {
			log.Warn("Missing price data in block", "number", number)
		} else if price.Rat().Cmp(common.Rat0) <= 0 {
			log.Error("Invalid price data in block", "number", number, "price", price.Rat().RatString())
			return errInvalidPrice
		} else {
			log.Info("Block price data found", "number", number, "price", price)
		}
	} else {
		// for regular block: extra = [vanity(32), signature(65)]
	}
	// All basic checks passed, verify the seal and return
	return d.verifySeal2(chain, header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (d *Dccs) snapshot(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := d.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}
		// If an on-disk checkpoint snapshot can be found, use that
		if number%checkpointInterval == 0 {
			if s, err := loadSnapshot(d.config, d.signatures, d.db, hash); err == nil {
				log.Trace("Loaded voting snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}
		// If we're at an checkpoint block, make a snapshot if it's known
		if number == 0 || (d.config.IsCheckpoint(number) && chain.GetHeaderByNumber(number-1) == nil) {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				signers := make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraSeal)/common.AddressLength)
				for i := 0; i < len(signers); i++ {
					copy(signers[i][:], checkpoint.Extra[extraVanity+i*common.AddressLength:])
				}
				snap = newSnapshot(d.config, d.signatures, number, hash, signers)
				if err := snap.store(d.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}
		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}
	snap, err := snap.apply(headers)
	if err != nil {
		return nil, err
	}
	d.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
		if err = snap.store(d.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// snapshot2 retrieves the authorization snapshot at a given point in time.
func (d *Dccs) snapshot2(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		snap *Snapshot
	)
	cp := d.config.Snapshot(number + 1)
	// looping until data/state are available on local
	for snap == nil {
		// Get signers from Nexty staking smart contract at the latest epoch checkpoint from block number
		checkpoint := chain.GetHeaderByNumber(cp)
		if checkpoint == nil {
			log.Trace("snapshot header not available", "number", cp)
			continue
		}
		hash := checkpoint.Hash()
		log.Trace("Reading signers from epoch checkpoint", "number", cp, "hash", hash)
		// If an in-memory snapshot was found, use that
		if s, ok := d.recents.Get(hash); ok && number+1 != d.config.ThangLongBlock.Uint64() {
			snap = s.(*Snapshot)
			log.Trace("Loading snapshot from mem-cache", "hash", snap.Hash, "length", len(snap.signers()))
			break
		}
		state, err := chain.StateAt(checkpoint.Root)
		if state == nil || err != nil {
			log.Trace("snapshot state not available", "number", cp, "err", err)
			continue
		}
		size := state.GetCodeSize(chain.Config().Dccs.Contract)
		if size <= 0 || state.Error() != nil {
			log.Trace("snapshot contract state not available", "number", cp, "err", state.Error())
			continue
		}
		index := common.BigToHash(common.Big0)
		result := state.GetState(chain.Config().Dccs.Contract, index)
		var length int64
		if (result == common.Hash{}) {
			length = 0
		} else {
			length = result.Big().Int64()
		}
		log.Trace("Total number of signer from staking smart contract", "length", length)
		signers := make([]common.Address, length)
		key := crypto.Keccak256Hash(hexutil.MustDecode(index.String()))
		for i := 0; i < len(signers); i++ {
			log.Trace("key hash", "key", key)
			singer := state.GetState(chain.Config().Dccs.Contract, key)
			signers[i] = common.HexToAddress(singer.Hex())
			key = key.Plus()
		}
		snap = newSnapshot(d.config, d.signatures, number, hash, signers)
		// Store found snapshot into mem-cache
		d.recents.Add(snap.Hash, snap)
		break
	}

	// Set current block number for snapshot to calculate the inturn & difficulty
	snap.Number = number
	return snap, nil
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (d *Dccs) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (d *Dccs) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.verifySeal2(chain, header, nil)
	}
	return d.verifySeal(chain, header, nil)
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}
	if _, ok := snap.Signers[signer]; !ok {
		return errUnauthorizedSigner
	}
	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only fail if the current block doesn't shift it out
			if limit := uint64(len(snap.Signers)/2 + 1); seen > number-limit {
				return errRecentlySigned
			}
		}
	}
	// Ensure that the difficulty corresponds to the turn-ness of the signer
	inturn := snap.inturn(header.Number.Uint64(), signer)
	if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
		return errInvalidDifficulty
	}
	if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
		return errInvalidDifficulty
	}
	return nil
}

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot2(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}
	if _, ok := snap.Signers[signer]; !ok {
		return errUnauthorizedSigner
	}

	headers, err := d.GetRecentHeaders(snap, chain, header, parents)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents, only fail if the current block doesn't shift it out
			return errRecentlySigned
		}
	}

	var parent *types.Header
	if len(headers) > 0 {
		parent = headers[0]
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	signerDifficulty := snap.difficulty(signer, parent)
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) Prepare(chain consensus.ChainReader, header *types.Header) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.prepare2(chain, header)
	}
	return d.prepare(chain, header)
}

// prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare(chain consensus.ChainReader, header *types.Header) error {
	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Coinbase = common.Address{}
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()
	// Assemble the voting snapshot to check which votes make sense
	snap, err := d.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	if !d.config.IsCheckpoint(number) {
		d.lock.RLock()

		// Gather all the proposals that make sense voting on
		addresses := make([]common.Address, 0, len(d.proposals))
		for address, authorize := range d.proposals {
			if snap.validVote(address, authorize) {
				addresses = append(addresses, address)
			}
		}
		// If there's pending proposals, cast a vote on them
		if len(addresses) > 0 {
			header.Coinbase = addresses[rand.Intn(len(addresses))]
			if d.proposals[header.Coinbase] {
				copy(header.Nonce[:], nonceAuthVote)
			} else {
				copy(header.Nonce[:], nonceDropVote)
			}
		}
		d.lock.RUnlock()
	}
	// Set the correct difficulty
	header.Difficulty = CalcDifficulty(snap, d.signer)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if d.config.IsCheckpoint(number) {
		for _, signer := range snap.signers() {
			header.Extra = append(header.Extra, signer[:]...)
		}
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare2(chain consensus.ChainReader, header *types.Header) error {
	header.Nonce = types.BlockNonce{}
	// Get the beneficiary of signer from smart contract and set to header's coinbase to give sealing reward later
	number := header.Number.Uint64()
	cp := d.config.Snapshot(number)
	checkpoint := chain.GetHeaderByNumber(cp)
	if checkpoint != nil {
		root, _ := chain.StateAt(checkpoint.Root)
		index := common.BigToHash(common.Big1).String()[2:]
		coinbase := "0x000000000000000000000000" + header.Coinbase.String()[2:]
		key := crypto.Keccak256Hash(hexutil.MustDecode(coinbase + index))
		result := root.GetState(chain.Config().Dccs.Contract, key)
		beneficiary := common.HexToAddress(result.Hex())
		header.Coinbase = beneficiary
	} else {
		log.Error("state is not available at the block number", "number", cp)
		return errors.New("state is not available at the block number")
	}

	// Set the correct difficulty
	snap, err := d.snapshot2(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = CalcDifficulty2(snap, d.signer, parent)
	log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	// Stop recording signers list in checkpoint after Endurio hardfork
	if !d.config.IsEndurio(header.Number) {
		if d.config.IsCheckpoint(number) {
			for _, signer := range snap.signers2() {
				header.Extra = append(header.Extra, signer.Address[:]...)
			}
		}
	} else if d.config.IsPriceBlock(number) {
		price := d.PriceEngine().CurrentPrice()
		if price == nil {
			log.Warn("No price to record in block", "number", number)
		} else if price.Rat().Cmp(common.Rat0) <= 0 {
			log.Error("Skip recording invalid price data", "price", price.Rat().RatString())
		} else {
			log.Info("Encode price to block extra", "price", price.Rat().RatString())
			header.Extra = append(header.Extra, PriceEncode(price)...)
		}
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// deployConsensusContracts deploys the consensus contract without any owner
func deployConsensusContracts(state *state.StateDB, chainConfig *params.ChainConfig, signers []common.Address) error {
	// Deploy NTF ERC20 Token Contract
	{
		owner := common.HexToAddress("0x000000270840d8ebdffc7d162193cc5ba1ad8707")
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := ntf.DeployNtfToken(auth, sim, owner)
			return address, err
		})
		if err != nil {
			return err
		}

		// replace the random generated sender address in MultiOwnable's manager field
		storage[common.BigToHash(common.Big0)] = owner.Hash()

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.TokenAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "NTFToken")
	}

	// Deploy Nexty Governance Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			stakeRequire := new(big.Int).Mul(new(big.Int).SetUint64(chainConfig.Dccs.StakeRequire), new(big.Int).SetUint64(1e+18))
			stakeLockHeight := new(big.Int).SetUint64(chainConfig.Dccs.StakeLockHeight)
			address, _, _, err := governance.DeployNextyGovernance(auth, sim, params.TokenAddress, stakeRequire, stakeLockHeight, signers)
			return address, err
		})
		if err != nil {
			return err
		}
		// Deploy or update
		deployer.CopyContractToAddress(state, chainConfig.Dccs.Contract, code, storage, true)
		log.Info("⚙ Contract deployed successful", "contract", "NextyGovernance")
	}

	return nil
}

func deployEndurioContracts(chain consensus.ChainReader, state *state.StateDB) error {
	prefundAddress := common.HexToAddress("0x95e2fcBa1EB33dc4b8c6DCBfCC6352f0a253285d")

	// Deploy Seigniorage Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := endurio.DeploySeigniorage(auth, sim,
				new(big.Int).SetUint64(chain.Config().Dccs.AbsorptionDuration),
				new(big.Int).SetUint64(chain.Config().Dccs.AbsorptionExpiration),
				new(big.Int).SetUint64(chain.Config().Dccs.SlashingDuration),
				new(big.Int).SetUint64(chain.Config().Dccs.LockdownExpiration),
			)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.SeigniorageAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "Seigniorage")
	}

	// Deploy VolatileToken Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := volatile.DeployVolatileToken(auth, sim, params.SeigniorageAddress, prefundAddress, common.Big1000)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.VolatileTokenAddress, code, storage, false)
		// Pre-fund the contract with the token amount
		state.SetBalance(params.VolatileTokenAddress, new(big.Int).Mul(common.Big1000, common.Big1e24))
		log.Info("⚙ Contract deployed successful", "contract", "VolatileToken")
	}

	// Deploy StableToken Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := stable.DeployStableToken(auth, sim, params.SeigniorageAddress, prefundAddress, common.Big1000)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.StableTokenAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "StableToken")
	}

	// Link them together
	{
		//state.IntermediateRoot(false)
		backend := backends.NewRealBackend(state, chain, nil)
		seign, err := endurio.NewSeigniorage(params.SeigniorageAddress, backend)
		if err != nil {
			log.Error("Failed to create new Seigniorage contract executor", "err", err)
			return err
		}

		consensusTransactOpts := &bind.TransactOpts{
			GasLimit: math.MaxUint64, // it's over 9000
			Signer: func(_ types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return tx, nil
			},
		}

		_, err = seign.RegisterTokens(consensusTransactOpts, params.VolatileTokenAddress, params.StableTokenAddress)
		if err != nil {
			log.Error("Failed to execute Seigniorage.OnBlockInitialized", "err", err)
			return err
		}
		state.Commit(false)
	}

	return nil
}

// Initialize implements the consensus.Engine
func (d *Dccs) Initialize(chain consensus.ChainReader, header *types.Header, state *state.StateDB) (types.Transactions, types.Receipts, error) {
	if !chain.Config().IsEndurio(header.Number) {
		return nil, nil, nil
	}
	if header.Number.Cmp(d.config.EndurioBlock) == 0 {
		if err := deployEndurioContracts(chain, state); err != nil {
			log.Error("Failed to deploy Endurio stablecoin contracts", "err", err)
			return nil, nil, err
		}
		header.Root = state.IntermediateRoot(false)
		log.Info("⚙ Successfully deploy Endurio stablecoin contracts")
		return nil, nil, nil
	}

	medianPrice, err := d.PriceEngine().CalcMedianPrice(chain, header.Number.Uint64()-params.CanonicalDepth)
	if err != nil {
		log.Trace("Failed to calculate canonical median price", "err", err, "number", header.Number)
	}

	OnBlockInitialized(chain, state, medianPrice)
	header.Root = state.IntermediateRoot(false)
	return nil, nil, nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	if chain.Config().IsThangLong(header.Number) {
		d.finalize2(chain, header, state, txs, uncles)
	}
	d.finalize(chain, header, state, txs, uncles)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if chain.Config().IsThangLong(header.Number) {
		return d.finalizeAndAssemble2(chain, header, state, txs, uncles, receipts)
	}
	return d.finalizeAndAssemble(chain, header, state, txs, uncles, receipts)
}

// finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	if chain.Config().IsThangLongPreparationBlock(header.Number) {
		// Retrieve the pre-fork signers list
		s, err := d.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
		if err != nil {
			return
		}
		sigs := s.signers()
		// Deploy the contract and ininitalize it with pre-fork signers
		if err = deployConsensusContracts(state, chain.Config(), sigs); err != nil {
			log.Error("Unable to deploy Nexty governance smart contract", "err", err)
			return
		}
		log.Info("Successfully deploy Nexty governance contract", "Number of sealers", len(sigs))
	}

	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// finalizeAndAssemble implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if chain.Config().IsThangLongPreparationBlock(header.Number) {
		// Retrieve the pre-fork signers list
		s, err := d.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
		if err != nil {
			return nil, err
		}
		sigs := s.signers()
		// Deploy the contract and ininitalize it with pre-fork signers
		if err = deployConsensusContracts(state, chain.Config(), sigs); err != nil {
			log.Error("Unable to deploy Nexty governance smart contract", "err", err)
			return nil, err
		}
		log.Info("Successfully deploy Nexty governance contract", "Number of sealers", len(sigs))
	}

	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (d *Dccs) Authorize(signer common.Address, signFn SignerFn, state *state.StateDB, header *types.Header) {
	if d.config.IsThangLong(header.Number) {
		size := state.GetCodeSize(d.config.Contract)
		log.Info("smart contract size", "size", size)
		if size > 0 && state.Error() == nil {
			// Get token holder from coinbase
			index := common.BigToHash(common.Big1).String()[2:]
			coinbase := "0x000000000000000000000000" + signer.String()[2:]
			key := crypto.Keccak256Hash(hexutil.MustDecode(coinbase + index))
			result := state.GetState(d.config.Contract, key)

			if (result == common.Hash{}) {
				log.Warn("Validator is not in activation sealer set")
			}
		}
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	d.signer = signer
	d.signFn = signFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()
	if chain.Config().IsThangLong(header.Number) {
		return d.seal2(chain, block, results, stop)
	}
	return d.seal(chain, block, results, stop)
}

// seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if d.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	d.lock.RLock()
	signer, signFn := d.signer, d.signFn
	d.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := d.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only wait if the current block doesn't shift it out
			if limit := uint64(len(snap.Signers)/2 + 1); number < limit || seen > number-limit {
				log.Info("Signed recently, must wait for others")
				return nil
			}
		}
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := time.Duration(len(snap.Signers)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))

		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}
	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DccsRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal2(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if d.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	d.lock.RLock()
	signer, signFn := d.signer, d.signFn
	d.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := d.snapshot2(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	headers, err := d.GetRecentHeaders(snap, chain, header, nil)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents
			log.Info("Signed recently, must wait for others")
			return nil
		}
	}
	var parent *types.Header
	if len(headers) > 0 {
		parent = headers[0]
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if !snap.inturn2(signer, parent) {
		// It's not our turn explicitly to sign, delay it a bit
		offset, err := snap.offset(signer, parent)
		if err != nil {
			return err
		}
		wiggle := d.calcDelayTimeForOffset(offset)
		delay += wiggle
		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}
	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DccsRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// calcDelayTime calculate delay time for current sealing node
func (d *Dccs) calcDelayTime(snap *Snapshot, block *types.Block, signer common.Address) time.Duration {
	header := block.Header()
	number := header.Number.Uint64()
	sigs := snap.signers2()
	pos := 0
	for seen, sig := range sigs {
		if sig.Address == signer {
			pos = seen
		}
	}
	cp := d.config.Checkpoint(number)
	total := uint64(len(sigs))
	offset := (number - cp) - (number-cp)/total*total
	log.Trace("calcDelayTime", "number", number, "checkpoint", cp, "len", uint64(len(sigs)), "offset", offset)
	if pos >= int(offset) {
		pos -= int(offset)
	} else {
		pos += len(sigs) - int(offset)
	}
	return d.calcDelayTimeForOffset(pos)
}

// calcDelayTime calculate delay time for current sealing node
func (d *Dccs) calcDelayTimeForOffset(pos int) time.Duration {
	delay := time.Duration(0)
	wiggle := float64(0.0)
	for i := 1; i <= pos; i++ {
		wiggle += math.Floor(float64(1.387978)/(float64(0.002313279)*float64(i)+float64(0.00462659)) + float64(499.9994))
	}
	wiggle = wiggle * float64(time.Millisecond)
	delay += time.Duration(int64(wiggle))
	return delay
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (d *Dccs) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	if chain.Config().IsThangLong(parent.Number) {
		snap, err := d.snapshot2(chain, parent.Number.Uint64(), parent.Hash(), nil)
		if err != nil {
			return nil
		}
		return CalcDifficulty2(snap, d.signer, parent)
	}
	snap, err := d.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return CalcDifficulty(snap, d.signer)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty(snap *Snapshot, signer common.Address) *big.Int {
	if snap.inturn(snap.Number+1, signer) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

// CalcDifficulty2 is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty2(snap *Snapshot, signer common.Address, parent *types.Header) *big.Int {
	return new(big.Int).SetUint64(snap.difficulty(signer, parent))
}

// SealHash returns the hash of a block prior to it being sealed.
func (d *Dccs) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

// Close implements consensus.Engine. It's a noop for clique as there is are no background threads.
func (d *Dccs) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (d *Dccs) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "dccs",
		Version:   "1.1",
		Service:   &API{chain: chain, dccs: d},
		Public:    false,
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

// DccsRLP returns the rlp bytes which needs to be signed for the proof-of-foundation
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func DccsRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

// calculateRewards calculate reward for block sealer
func (d *Dccs) calculateRewards(chain consensus.ChainReader, state *state.StateDB, header *types.Header) {
	number := header.Number.Uint64()
	yo := (number - chain.Config().Dccs.ThangLongBlock.Uint64()) / blockPerYear.Uint64()
	per := yo
	if per > 5 {
		per = 5
	}
	totalSupply := new(big.Int).Mul(initialSupply, big.NewInt(1e+18)) // total supply in Wei
	for i := uint64(1); i <= yo; i++ {
		r := i
		if r > 5 {
			r = 5
		}
		totalReward := new(big.Int).Mul(totalSupply, rewards[r])
		totalReward = totalReward.Div(totalReward, big.NewInt(1e+5))
		totalSupply = totalSupply.Add(totalSupply, totalReward)
	}
	totalYearReward := new(big.Int).Mul(totalSupply, rewards[per])
	totalYearReward = totalYearReward.Div(totalYearReward, big.NewInt(1e+5))
	log.Trace("Total reward for current year", "reward", totalYearReward, "total sypply", totalSupply)
	blockReward := new(big.Int).Div(totalYearReward, blockPerYear)
	log.Trace("Give reward for sealer", "beneficiary", header.Coinbase, "reward", blockReward, "number", number, "hash", header.Hash)
	state.AddBalance(header.Coinbase, blockReward)
}

// GetRecentHeaders get some recent headers back from the current header.
// Return empty header list at checkpoint, a.k.a reset the recent headers at every checkpoint.
func (d *Dccs) GetRecentHeaders(snap *Snapshot, chain consensus.ChainReader, header *types.Header, parents []*types.Header) ([]*types.Header, error) {
	var headers []*types.Header
	number := header.Number.Uint64()
	// Reset the recent headers at checkpoint to ensure the in-turn signer
	if d.config.IsCheckpoint(number) {
		return headers, nil
	}
	limit := len(snap.Signers) / 2
	num, hash := number-1, header.ParentHash
	for i := 1; i <= limit; i++ {
		// shortcut for genesis block because it has no signature
		if num == 0 {
			break
		}
		var h *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			h = parents[len(parents)-1]
			if h.Hash() != hash || h.Number.Uint64() != num {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			h = chain.GetHeader(hash, num)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, h)
		if d.config.IsCheckpoint(num) {
			break
		}
		num, hash = num-1, h.ParentHash
	}
	return headers, nil
}
