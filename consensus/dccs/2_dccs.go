package dccs

import (
	"bytes"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/deployer"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/contracts/nexty/governance"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	nonceSealerNone  = hexutil.MustDecode("0x0000000000000000")
	nonceSealerJoin  = hexutil.MustDecode("0x1111111111111111")
	nonceSealerLeave = hexutil.MustDecode("0x2222222222222222")
	nonceSealerSlash = hexutil.MustDecode("0xffffffffffffffff")
)

// verifyHeader2 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (d *Dccs) verifyHeader2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	log.Debug("verifyHeader2", "number", header.Number)
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	checkpoint := d.config.CoLoaBlock.Uint64() == number
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on genesis block of CoLoa hardfork
	signersBytes := len(header.Extra) - extraVanity - extraSeal
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

// verifyCascadingFields2 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (d *Dccs) verifyCascadingFields2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	log.Debug("verifyCascadingFields2", "number", header.Number)
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 || d.config.CoLoaBlock.Cmp(header.Number) == 0 {
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
	// If the block is a checkpoint block, verify the signer list
	// But if the checkpoint is CoLoa hardfork block then just take it as genesis data without verification
	// And from CoLoa hardfork, we don't save the signer list to checkpoint header block anymore
	// if d.config.CoLoaBlock.Cmp(header.Number) == 0 {
	// 	// Retrieve the snapshot needed to verify this header and cache it
	// 	snap, err := d.snapshot3(chain, number-1, header.ParentHash, parents)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	extraData := make([]byte, len(snap.Sealers)*2*common.AddressLength)
	// 	for i, signer := range snap.signers() {
	// 		copy(extraData[i*2*common.AddressLength:], signer.Signer[:])
	// 		copy(extraData[((i*2)+1)*common.AddressLength:], signer.Beneficiary[:])
	// 	}
	// 	extraSuffix := len(header.Extra) - extraSeal
	// 	if !bytes.Equal(header.Extra[extraVanity:extraSuffix], extraData) {
	// 		return errInvalidCheckpointSigners
	// 	}
	// }
	// All basic checks passed, verify the seal and return
	return d.verifySeal2(chain, header, parents)
}

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	log.Debug("verifySeal2", "number", header.Number)
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

	beneficiary, ok := snap.Sealers[signer]
	if !ok {
		return errUnauthorizedSigner
	}

	if header.Coinbase != beneficiary {
		return errInvalidBeneficiary
	}

	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only fail if the current block doesn't shift it out
			if limit := uint64(len(snap.Sealers)/2 + 1); seen > number-limit {
				return errRecentlySigned
			}
		}
	}
	// Ensure that the difficulty corresponds to the turn-ness of the signer
	var parent *types.Header
	if len(parents) == 0 {
		parent = chain.GetHeaderByNumber(number - 1)
	} else {
		parent = parents[len(parents)-1]
	}
	signerDifficulty := snap.difficulty(signer, parent)
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal2(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()
	log.Debug("seal2", "number", header.Number)

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
	if _, authorized := snap.Sealers[signer]; !authorized {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only wait if the current block doesn't shift it out
			if limit := uint64(len(snap.Sealers)/2 + 1); number < limit || seen > number-limit {
				log.Info("Signed recently, must wait for others")
				return nil
			}
		}
	}
	parent := chain.GetHeaderByNumber(number - 1)
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if !snap.inturn(signer, parent) {
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

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare2(chain consensus.ChainReader, header *types.Header) error {
	log.Debug("prepare2", "number", header.Number)
	header.Nonce = types.BlockNonce{}
	// Get the beneficiary of signer from snapshot and set to header's coinbase to give sealing reward later
	number := header.Number.Uint64()
	snap, err := d.snapshot2(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	header.Coinbase = snap.Sealers[header.Coinbase]

	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Set the correct difficulty
	header.Difficulty = CalcDifficulty2(snap, d.signer, parent)
	log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// snapshot2 retrieves the authorization snapshot at a given point in time.
func (d *Dccs) snapshot2(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot2, error) {
	log.Debug("snapshot2", "number", number)
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot2
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := d.recents.Get(hash); ok {
			switch s.(type) {
			case Snapshot2:
				snap = s.(*Snapshot2)
				break
			default:
				log.Trace("Ignore ancient snapshot from pre-hardfork")
			}
		}
		// If an on-disk checkpoint snapshot can be found, use that
		if number%checkpointInterval == 0 {
			if s, err := loadSnapshot2(d.config, d.signatures, d.db, hash); err == nil {
				log.Trace("Loaded voting snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}
		// If we're at the genesis, snapshot the initial state. Alternatively if we're
		// at a checkpoint block without a parent (light client CHT), or we have piled
		// up more headers than allowed to be reorged (chain reinit from a freezer),
		// consider the checkpoint trusted and snapshot it.
		// TODO: Need to support freezer and light client later
		if number == d.config.CoLoaBlock.Uint64() {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				signers := make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraSeal)/common.AddressLength/2)
				beneficiaries := make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraSeal)/common.AddressLength/2)
				for i := 0; i < len(signers); i++ {
					copy(signers[i][:], checkpoint.Extra[extraVanity+(2*i)*common.AddressLength:])
					copy(beneficiaries[i][:], checkpoint.Extra[extraVanity+(2*i+1)*common.AddressLength:])
				}
				// CoLoa block is the checkpoint block then snapshot's sealer hash is block hash
				snap = newSnapshot2(d.config, d.signatures, number, hash, hash, signers, beneficiaries)
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

	// Calculate snapshot's sealer hash used to order sealer list
	cpNumber := d.config.Snapshot(snap.Number)
	cp := chain.GetHeaderByNumber(cpNumber)
	if cp != nil {
		snap.SealerHash = cp.Hash()
	} else {
		for i := len(headers) - 1; i >= 0; i-- {
			if headers[i].Number.Uint64() == cpNumber {
				snap.SealerHash = headers[i].Hash()
				break
			}
		}
	}

	d.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
		if err := snap.store(d.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}

	return snap, nil
}

// CalcDifficulty2 is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty2(snap *Snapshot2, signer common.Address, parent *types.Header) *big.Int {
	return new(big.Int).SetUint64(snap.difficulty(signer, parent))
}

func upgradeGovernanceContract(state *state.StateDB, chainConfig *params.ChainConfig) error {
	// Generate contract code and data using a simulated backend
	code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
		address, _, _, err := governance.DeployNextyGovernance(auth, sim)
		return address, err
	})
	if err != nil {
		return err
	}
	// Upgrade code only, keeping all ancient contract state
	deployContract(state, chainConfig.Dccs.Contract, code, storage, true, true)

	return nil
}
