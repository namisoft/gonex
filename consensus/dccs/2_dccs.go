// Copyright 2019 The gonex Authors
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
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vdf"
	"github.com/ethereum/go-ethereum/log"
)

var (
	// errOutOfTurn is returned if a header is signed by an authorized entity
	// that already way passed its turn, thus is temporarily not allowed to.
	errOutOfTurn = errors.New("out of turn")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
	d.blockDelayer = vdf.NewDelayer(127)
	return d
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

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	// signersBytes := len(header.Extra) - extraVanity - extraSeal

	// checkpoint := d.config.IsCheckpoint(number)
	// if !checkpoint && signersBytes != 0 {
	// 	return errExtraSigners
	// }
	// if checkpoint && signersBytes%common.AddressLength != 0 {
	// 	return errInvalidCheckpointSigners
	// }

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

	// All basic checks passed, verify the seal and return
	return d.verifySeal2(chain, header, parents)
}

const epoch = 256 // = core.blockCacheLimit

func (d *Dccs) rank(targetSealer common.Address, parentHash common.Hash, parents []*types.Header, chain consensus.ChainReader) (int, int, int, error) {
	offset, last := -1, -1
	sealers := make(map[common.Address]struct{})
	for i, hash := 0, parentHash; i < epoch; i++ {
		var header *types.Header
		if i < len(parents) && parents[i].Hash() == hash {
			header = parents[i]
		} else {
			header = chain.GetHeaderByHash(hash)
		}
		if header == nil {
			log.Error("no header", "hash", hash, "i", i)
			return -1, -1, 0, fmt.Errorf("no header with hash %x", hash)
		}
		if (header.ParentHash == common.Hash{}) {
			// genesis block reached
			break
		}
		signer, err := ecrecover(header, d.signatures)
		if err != nil {
			log.Error("err", "err", err)
			return -1, -1, 0, err
		}
		if _, ok := sealers[signer]; !ok {
			// not in set
			sealers[signer] = struct{}{}
			last = i // first appearance of the least recent sealer
			if signer == targetSealer {
				offset = i // first appearance of target sealer
			}
		}
		hash = header.ParentHash
	}
	return offset, last, len(sealers), nil
}

// unused
func (d *Dccs) sealers2(chain consensus.ChainReader, parentHash common.Hash, parents []*types.Header) (map[common.Address]uint64, error) {
	sealers := make(map[common.Address]uint64)
	for i, hash := 0, parentHash; i < epoch; i++ {
		var header *types.Header
		if i < len(parents) && parents[i].Hash() == hash {
			header = parents[i]
		} else {
			header = chain.GetHeaderByHash(hash)
		}
		if header == nil {
			log.Error("no header", "hash", hash, "i", i)
			return nil, fmt.Errorf("no header with hash %x", hash)
		}
		if (header.ParentHash == common.Hash{}) {
			// genesis block reached
			break
		}
		signer, err := ecrecover(header, d.signatures)
		if err != nil {
			log.Error("err", "err", err)
			return nil, err
		}
		if _, ok := sealers[signer]; !ok {
			// not in set
			sealers[signer] = uint64(i)
		}
		hash = header.ParentHash
	}
	return sealers, nil
}

// unused
func (d *Dccs) sealers23(chain consensus.ChainReader, number uint64, parents []*types.Header) (map[common.Address]uint64, error) {
	sealers := make(map[common.Address]uint64)

	for i := 0; i < epoch; i++ {
		var header *types.Header
		n := number - uint64(i) - 1
		if i < len(parents) {
			header = parents[i]
		} else {
			if n <= 0 {
				break
			}
			header = chain.GetHeaderByNumber(n)
		}
		if header == nil {
			log.Error("no header", "n", n, "i", i)
			return nil, fmt.Errorf("no header at number %v", n)
		}
		signer, err := ecrecover(header, d.signatures)
		if err != nil {
			log.Error("err", "err", err)
			return nil, err
		}
		if _, ok := sealers[signer]; !ok {
			// not in set
			sealers[signer] = uint64(i)
		}
	}
	log.Trace("sealers", "sealers", sealers)
	return sealers, nil
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

	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}

	offset, _, _, err := d.rank(signer, header.ParentHash, parents, chain)
	if err != nil {
		return fmt.Errorf("dccs.verifySeal2: unable to get sealer rank: %v", err)
	}

	if offset < 0 {
		return errUnauthorizedSigner
	}

	if offset == 0 {
		return errRecentlySigned
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	// if header.Difficulty.Cmp(common.Big1) != 0 {
	// 	log.Error("verifySeal2: invalid diff", "actual", header.Difficulty, "expected", 1)
	// 	return errInvalidDifficulty
	// }
	return nil
}

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare2(chain consensus.ChainReader, header *types.Header) error {
	header.Nonce = types.BlockNonce{}
	number := header.Number.Uint64()
	//header.Coinbase = params.ZeroAddress // TODO

	// offset, last, count, err := d.rank(params.ZeroAddress, header.ParentHash, nil, chain)
	// if err != nil {
	// 	return err
	// }

	//header.Difficulty = common.Big1
	//log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	// if d.config.IsCheckpoint(number) {
	// 	for _, signer := range snap.signers2() {
	// 		header.Extra = append(header.Extra, signer.Address[:]...)
	// 	}
	// }
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	//d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Calculate any block reward for the sealer and commit the final state root
	//d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

//const baseIteration = 9876 / 2 // for vdf_go, 120 bit -> 2s
const baseIteration = 234567 / 2 // for vdf-cli, 120 bit -> 2s

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

	offset, last, count, err := d.rank(signer, header.ParentHash, nil, chain)
	if err != nil {
		return err
	}

	if err != nil {
		return fmt.Errorf("dccs.seal2: unable to get sealer rank: %v", err)
	}

	if offset < 0 {
		return errUnauthorizedSigner
	}

	if offset == 0 {
		return errRecentlySigned
	}

	iteration := baseIteration
	iteration += baseIteration * (count - offset) / 4
	if iteration <= 0 {
		iteration = last - offset + 1
		//return errOutOfTurn
	}

	signTheBlock := func() error {
		// Sign all the things!
		sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DccsRLP(header))
		if err != nil {
			return err
		}
		copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
		return nil
	}

	log.Info("Delayed signing requested", "offset", offset, "last", last, "count", count, "iteration", iteration)
	resCh := d.blockDelayer.Request(header.ParentHash[:], uint64(iteration))

	go func() {
		select {
		case <-stop:
			return
		case output := <-resCh:
			header.MixDigest = common.BytesToHash(output)
			header.Difficulty = big.NewInt(int64(diff(output[:])))
		}
		signTheBlock()
	}()

	return nil
}

func diff(output []byte) (diff int) {
	return 255 + int(output[0])
}
