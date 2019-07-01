package deployer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// DeployContract deploy a smart contract to simulated chain to get out the contract's code and state
func DeployContract(deployCallback func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error)) (code []byte, storage map[common.Hash]common.Hash, err error) {
	// Generate a new random account and a funded simulator
	prvKey, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(prvKey)
	auth.GasLimit = 12344321
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: new(big.Int).Lsh(big.NewInt(1), 256-7)}}, auth.GasLimit)
	address, err := deployCallback(sim, auth)
	if err != nil {
		log.Error("Unable to deploy consensus contract", "error", err)
		return nil, nil, err
	}
	sim.Commit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code, _ = sim.CodeAt(ctx, address, nil)
	storage = make(map[common.Hash]common.Hash)
	sim.ForEachStorageAt(ctx, address, nil, func(key, val common.Hash) bool {
		log.Trace("DecodeBytes", "key", key, "val", val.Hex(), "storage value", storage[key].Hex())
		storage[key] = val
		return true
	})
	return code, storage, err
}
