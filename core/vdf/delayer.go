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

// Package vdf implements the VDF engine.
package vdf

import (
	"bytes"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Delayer is a single process/routine for delay task of <seed, iteration>.
// The same task request will be ignored; a new task request will interrupt
// and replace the current one.
//
// bitSize value of 2^n-1 is recommened
// output size (in bytes) will be ((bitSize+16)>>4)*4
type Delayer struct {
	bitSize    uint64
	loopOnce   sync.Once
	requestCh  chan task
	responseCh chan []byte
}

// NewDelayer creates a new Delayer instance
func NewDelayer(bitSize uint64) *Delayer {
	return &Delayer{
		bitSize:    bitSize,
		requestCh:  make(chan task),
		responseCh: make(chan []byte),
	}
}

// OutputSize returns size in bytes of the VDF output (y+proof)
func (d *Delayer) OutputSize() uint64 {
	return ((d.bitSize + 16) >> 4) << 2
}

// Request request new delay task.
// The same task request will be ignored; a new task request will interrupt
// and replace the current one.
func (d *Delayer) Request(seed []byte, iteration uint64) <-chan []byte {
	d.loopOnce.Do(func() {
		go d.loop()
	})
	d.requestCh <- task{
		seed:      seed,
		iteration: iteration,
	}
	return d.responseCh
}

// Stop stops the current delay task.
func (d *Delayer) Stop() {
	d.requestCh <- task{}
}

func (d *Delayer) loop() {
	var current task
	stopCh := make(chan struct{})
	for {
		select {
		case t := <-d.requestCh:
			if t.Equal(current) {
				log.Trace("Delayer: discarding duplicated request", "seed", common.Bytes2Hex(t.seed), "iteration", t.iteration)
				continue
			}
			if current.iteration > 0 {
				// cancel the current task
				select {
				case stopCh <- struct{}{}:
				default:
				}
			}
			if t.iteration == 0 {
				// don't start new worker routine
				log.Error("Delayer: cancelation")
				continue
			}
			current = t
			// start new worker routine
			go func(t task) {
				output, err := Instance().Generate(t.seed, t.iteration, d.bitSize, stopCh)
				if err != nil {
					log.Error("Delayer: VDF worker loop failed", "err", err)
					return
				}
				select {
				case d.responseCh <- output:
				default:
					log.Warn("Delayer: VDF result is not read by miner", "output", common.Bytes2Hex(output))
				}
			}(t)
		}
	}
}

// send empty task{} to stop the calcuation
type task struct {
	seed      []byte
	iteration uint64
}

func (t task) Equal(u task) bool {
	return t.iteration == u.iteration && bytes.Equal(t.seed, u.seed)
}
