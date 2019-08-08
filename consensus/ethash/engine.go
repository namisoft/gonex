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

package ethash

import (
	"bytes"
	"math/big"
	"runtime"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var (
	engine     *Ethash
	engineOnce sync.Once
)

// Instance returns the singleton instance of the VDF Engine
func Instance() *Ethash {
	engineOnce.Do(func() {
		engine = NewEngine()
	})
	return engine
}

// NewEngine creates a light version of ethash PoW scheme useful only for header verification
func NewEngine() *Ethash {
	config := Config{
		CacheDir:     "gonex-ethash",
		CachesInMem:  1,
		CachesOnDisk: 2,
	}
	ethash := &Ethash{
		config: config,
		caches: newlru("cache", config.CachesInMem, newCache),
	}
	return ethash
}

// Verify checks whether a block satisfies the PoW difficulty requirements,
// either using the usual ethash cache for it, or alternatively using a full DAG
// to make remote mining fast.
func (ethash *Ethash) Verify(headerNumber, difficulty, nonce *big.Int, mixDigest, sealHash common.Hash) error {
	// Ensure that we have a valid difficulty for the block
	if difficulty.Sign() <= 0 {
		return errInvalidDifficulty
	}
	// Recompute the digest and PoW values
	number := headerNumber.Uint64()

	// Slow-but-light PoW verification only, use an ethash cache
	cache := ethash.cache(number)
	size := datasetSize(number)
	digest, result := hashimotoLight(size, cache.cache, sealHash[:], nonce.Uint64())

	// Caches are unmapped in a finalizer. Ensure that the cache stays alive
	// until after the call to hashimotoLight so it's not unmapped while being used.
	runtime.KeepAlive(cache)

	// Verify the calculated values against the ones provided in the header
	if !bytes.Equal(mixDigest[:], digest) {
		return errInvalidMixDigest
	}
	target := new(big.Int).Div(two256, difficulty)
	if new(big.Int).SetBytes(result).Cmp(target) > 0 {
		return errInvalidPoW
	}
	return nil
}
