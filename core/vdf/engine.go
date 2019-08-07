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

package vdf

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	"github.com/harmony-one/vdf/src/vdf_go"
)

var (
	engine     *Engine
	engineOnce sync.Once
)

// Engine defines the structure of the engine
type Engine struct {
	cli string
}

// Instance returns the singleton instance of the VDF Engine
func Instance() *Engine {
	engineOnce.Do(func() {
		engine = newEngine()
	})
	return engine
}

func newEngine() *Engine {
	cli, err := exec.LookPath("vdf-cli")
	if err != nil {
		log.Warn("vdf.newEngine", "vdf-cli", "not found")
	}
	return &Engine{cli}
}

// Verify verifies the generated output against the seed
func (e *Engine) Verify(seed, output []byte, iteration uint64, bitSize uint64) (valid bool) {
	defer func() {
		if x := recover(); x != nil {
			log.Error("vdf.Verify: verification process panic", "reason", x)
			valid = false
		}
	}()
	return vdf_go.VerifyVDF(seed, output, int(iteration), int(bitSize))
}

// Generate generates the vdf output = (y, proof)
func (e *Engine) Generate(seed []byte, iteration uint64, bitSize uint64, stop <-chan struct{}) (output []byte, err error) {
	if len(e.cli) == 0 {
		defer func() {
			if x := recover(); x != nil {
				log.Error("vdf.Generate: generation process panic", "reason", x)
				err = fmt.Errorf("%v", x)
			}
		}()
		y, proof := vdf_go.GenerateVDFWithStopChan(seed, int(iteration), int(bitSize), stop)
		if y == nil || proof == nil {
			return nil, nil
		}
		return append(y, proof...), nil
	}

	cmd := exec.Command(e.cli,
		"-l"+strconv.Itoa(int(bitSize)),
		common.Bytes2Hex(seed),
		strconv.Itoa(int(iteration)))

	log.Trace(e.cli + " -l" + strconv.Itoa(int(bitSize)) + " " + common.Bytes2Hex(seed) + " " + strconv.Itoa(int(iteration)))

	var done chan struct{}
	if stop != nil {
		done = make(chan struct{})
		go func() {
			select {
			case <-stop:
				if cmd == nil || cmd.Process == nil {
					return
				}
				log.Trace("vdf.Generate: vdf-cli interrupted")
				if err := cmd.Process.Kill(); err != nil {
					log.Error("vdf.Generate: failed to kill vdf-cli process", "err", err)
				}
				return
			case <-done:
				log.Trace("vdf.Generate: vdf-cli done")
				return
			}
		}()
	}

	log.Trace("vdf.Generate", "seed", common.Bytes2Hex(seed), "iteration", iteration, "output", common.Bytes2Hex(output))
	output, err = cmd.Output()

	if stop != nil && done != nil {
		// release the stopping goroutine above
		select {
		case done <- struct{}{}:
		default:
		}
	}

	if err, ok := err.(*exec.ExitError); ok {
		// verification failed
		log.Trace("vdf.Generate", "error code", err.Error())
		return nil, err
	}
	if err != nil {
		// sum ting wong
		log.Error("vdf.Generate", "error", err)
		return nil, err
	}

	strOutput := strings.TrimSpace(string(output))

	log.Trace("vdf.Generate", "output", strOutput)
	return common.Hex2Bytes(strOutput), nil
}
