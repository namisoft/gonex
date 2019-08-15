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
	"math"
	"math/big"
	"time"

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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// errInvalidPrice is returned if price block contains invalid price value in
	// their extra-data fields.
	errInvalidPrice = errors.New("price block contains invalid price value")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	d.init1()
	return d
}

// PriceEngine creates and returns the PriceEngine singleton instance
func (d *Dccs) PriceEngine() *PriceEngine {
	d.priceEngineOnce.Do(func() {
		d.priceEngine = newPriceEngine(d.config)
	})
	return d.priceEngine
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
	// Only check header's signer list before CoLoa hardfork
	if !d.config.IsCoLoa(header.Number) {
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
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// Stop recording signers list in checkpoint after CoLoa hardfork
	if !d.config.IsCoLoa(header.Number) {
		// If the block is a checkpoint block, verify the signer list
		if d.config.IsCheckpoint(number) {
			signers := make([]byte, len(snap.Signers)*common.AddressLength)
			for i, signer := range snap.signers1() {
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

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal2(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	return d.verifySeal1(chain, header, parents)
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
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = CalcDifficulty1(snap, d.signer, parent)
	log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	// Stop recording signers list in checkpoint after CoLoa hardfork
	if !d.config.IsCoLoa(header.Number) {
		if d.config.IsCheckpoint(number) {
			for _, signer := range snap.signers1() {
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

// initialize implements the consensus.Engine
func (d *Dccs) initialize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB) (types.Transactions, types.Receipts, error) {
	if header.Number.Cmp(d.config.CoLoaBlock) == 0 {
		if err := deployCoLoaContracts(chain, header, state); err != nil {
			log.Error("Failed to deploy CoLoa stablecoin contracts", "err", err)
			return nil, nil, err
		}
		header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
		log.Info("⚙ Successfully deploy CoLoa stablecoin contracts")
		return nil, nil, nil
	}

	medianPrice, err := d.PriceEngine().CalcMedianPrice(chain, header.Number.Uint64()-params.CanonicalDepth)
	if err != nil {
		log.Trace("Failed to calculate canonical median price", "err", err, "number", header.Number)
	}

	txs, receipts, err := OnBlockInitialized(chain, header, state, medianPrice)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	return txs, receipts, err
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	d.finalize1(chain, header, state, txs, uncles)
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble2(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	return d.finalizeAndAssemble1(chain, header, state, txs, uncles, receipts)
}

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal2(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return d.seal1(chain, block, results, stop)
}

func deployCoLoaContracts(chain consensus.ChainReader, header *types.Header, state *state.StateDB) error {
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
			address, _, _, err := volatile.DeployVolatileToken(auth, sim, params.SeigniorageAddress, common.Address{}, common.Big0)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.VolatileTokenAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "VolatileToken")
	}

	// Deploy StableToken Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := stable.DeployStableToken(auth, sim, params.SeigniorageAddress, common.Address{}, common.Big0)
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
		backend := backends.NewRealBackend(chain, header, state, nil)
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
			log.Error("Failed to execute Seigniorage.RegisterTokens", "err", err)
			return err
		}
		state.Commit(false)
	}
	return nil
}
