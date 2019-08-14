// Copyright 2015 The gonex Authors
// This file is part of the gonex library.
//
// The gonex library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gonex library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gonex library. If not, see <http://www.gnu.org/licenses/>.

package dccs

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/stable"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/volatile"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	emptyHash = common.Hash{}
)

// OnBlockInitialized handles supply absorption on each block initialization
func OnBlockInitialized(chain consensus.ChainReader, header *types.Header, state *state.StateDB, medianPrice *Price) (types.Transactions, types.Receipts, error) {
	backend := backends.NewRealBackend(chain, header, state, nil) // consensus only
	target := common.Big0

	if medianPrice != nil {
		stableToken, err := stable.NewStableTokenCaller(params.StableTokenAddress, backend)
		if err != nil {
			log.Error("Failed to create StableToken contract caller", "err", err)
			return nil, nil, err
		}
		target, err = stableToken.TotalSupply(nil)
		if err != nil {
			log.Error("Failed to get stable token supply", "err", err)
			return nil, nil, err
		}

		target.Mul(target, medianPrice.Rat().Num())
		target.Div(target, medianPrice.Rat().Denom())
	}

	volatileToken, err := volatile.NewVolatileTokenCaller(params.VolatileTokenAddress, backend)
	if err != nil {
		log.Error("Failed to create VolatileToken contract caller", "err", err)
		return nil, nil, err
	}
	supply, err := volatileToken.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to get volatile token supply", "err", err)
		return nil, nil, err
	}

	log.Trace("VolatileToken supply before", "supply", supply)

	seign, err := endurio.NewSeigniorage(params.SeigniorageAddress, backend)
	if err != nil {
		log.Error("Failed to create new Seigniorage contract executor", "err", err)
		return nil, nil, err
	}

	consensusTransactOpts := &bind.TransactOpts{
		GasLimit: math.MaxUint64, // it's over 9000
		Signer: func(_ types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	state.Prepare(emptyHash, emptyHash, 0)

	tx, err := seign.OnBlockInitialized(consensusTransactOpts, target)
	if err != nil {
		log.Error("Failed to execute Seigniorage.OnBlockInitialized", "err", err)
		return nil, nil, err
	}
	failed := err != nil

	newSupply, err := volatileToken.TotalSupply(nil)
	if newSupply.Cmp(supply) != 0 {
		log.Trace("VolatileToken supply after", "new supply", newSupply, "change", new(big.Int).Sub(newSupply, supply))
		stateObject := state.GetOrNewStateObject(params.VolatileTokenAddress)
		stateObject.SetBalance(newSupply)
		stateObject.CommitTrie(state.Database())
	}

	root := state.IntermediateRoot(chain.Config().IsEIP158(header.Number)).Bytes()
	receipt := types.NewReceipt(root, failed, 0)
	receipt.Logs = state.GetLogs(emptyHash)
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	log.Trace("OnBlockInitialized", "receipt.Logs len", len(receipt.Logs))
	for i, l := range receipt.Logs {
		log.Trace("OnBlockInitialized", "i", i,
			"address", l.Address,
			"data", common.Bytes2Hex(l.Data),
		)
		for _, topic := range l.Topics {
			log.Trace("+ Topic", "topic", topic)
		}
	}

	return types.Transactions{tx}, types.Receipts{receipt}, nil
}

// AbsorbedStat returns ethstats data for stablecoin supply absorbed by the block
func AbsorbedStat(chain consensus.ChainReader, number uint64) string {
	if number <= 0 {
		return ""
	}
	header := chain.GetHeaderByNumber(number - 1)
	state, err := chain.StateAt(header.Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + string(number-1)
	}
	oldSupply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if oldSupply == nil {
		return "No Old Supply"
	}

	header = chain.GetHeaderByNumber(number)
	state, err = chain.StateAt(header.Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + string(number)
	}
	supply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if supply == nil {
		return "No New Supply"
	}
	return supply.Sub(supply, oldSupply).String()
}

// RemainToAbsorbStat returns ethstats data for stablecoin supply remain to absorb
func RemainToAbsorbStat(chain consensus.ChainReader, number uint64) string {
	header := chain.GetHeaderByNumber(number)
	if header == nil {
		return "No Header"
	}
	state, err := chain.StateAt(chain.GetHeaderByNumber(number).Root)
	if err != nil {
		return "Failed to get chain state: " + err.Error()
	}
	if state == nil {
		return "No state at " + string(number)
	}
	// Random key to make sure no one has any special right
	backend := backends.NewRealBackend(chain, header, state, nil)

	caller, err := endurio.NewSeigniorageCaller(params.SeigniorageAddress, backend)
	if err != nil {
		return "Failed to create Seigniorage caller: " + err.Error()
	}
	hasAbsorption, remain, err := caller.GetRemainToAbsorb(nil)
	if err != nil {
		return "Failed to call Seigniorage.GetRemainToAbsorb: " + err.Error()
	}
	if !hasAbsorption || remain == nil {
		return "0"
	}
	return remain.String()
}

// StableSupplyStat returns ethstats data for stablecoin supply
func StableSupplyStat(chain consensus.ChainReader, number uint64) string {
	header := chain.GetHeaderByNumber(number)
	if header == nil {
		return "No Header"
	}
	state, err := chain.StateAt(chain.GetHeaderByNumber(number).Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + string(number)
	}
	supply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if supply == nil {
		return "No Supply"
	}
	return supply.String()
}

// GetStableTokenSupply returns the current supply of the stable token in the stateDB
func GetStableTokenSupply(chain consensus.ChainReader, header *types.Header, state *state.StateDB) (*big.Int, error) {
	backend := backends.NewRealBackend(chain, header, state, nil)
	caller, err := stable.NewStableTokenCaller(params.StableTokenAddress, backend)
	if err != nil {
		return nil, err
	}
	return caller.TotalSupply(nil)
}
