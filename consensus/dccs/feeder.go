// Copyright 2015 The go-ethereum Authors
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

package dccs

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
)

// Data represents the external data feeded from outside
type Data struct {
	Value          interface{}
	Timestamp      int64
	Source         string
	reentranceFlag int64 // prevent request routine to run twice
}

// PriceData represents the external price feeded from outside
type PriceData struct {
	Value     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	Exchange  string `json:"exchange"`
}

// feeder is the main object which takes care of feeding data from outside to consensus
// engine and gathering the sealing result.
type feeder struct {
	mutex sync.RWMutex
	data  map[string]*Data
}

func newFeeder() *feeder {
	feeder := &feeder{}
	return feeder
}

func (f *feeder) getCurrent(url string) *Data {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	var data *Data
	var ok bool
	if data, ok = f.data[url]; !ok {
		return nil
	}

	if data.Value == nil {
		// data is being fetched the first time
		return nil
	}

	return data
}

// Yielding non-reentrant async request.
func (f *feeder) requestUpdate(url string) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	var data *Data
	var ok bool
	if data, ok = f.data[url]; !ok {
		f.mutex.Lock()
		defer f.mutex.Unlock()
		data = &Data{}
		f.data[url] = data
	}

	if !atomic.CompareAndSwapInt64(&data.reentranceFlag, 0, 1) {
		// one routine for one url only
		return
	}

	go func() {
		defer atomic.StoreInt64(&data.reentranceFlag, 0)
		// fetch the data here
		response, err := http.Get(url)
		if err != nil {
			log.Error("Failed to request for data", "url", url, "error", err)
			return
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error("Failed to read response body", "url", url, "error", err, "reponse", response)
			return
		}

		var priceData PriceData
		err = json.Unmarshal(body, &priceData)
		if err != nil {
			log.Error("Failed to unmarshal price json", "url", url, "error", err, "body", body)
			return
		}

		log.Trace("PriceData", "priceData", priceData)

		price := PriceFromString(priceData.Value)
		if price == nil {
			log.Error("Failed to parse price value", "url", url, "error", err, "priceData.Value", priceData.Value)
			return
		}

		f.mutex.Lock()
		defer f.mutex.Unlock()
		data.Value = &price
		data.Timestamp = priceData.Timestamp
		data.Source = priceData.Exchange
	}()
}

func (f *feeder) Price() *Price {
	return (*Price)(big.NewRat(1, 1))
}

// Price encoded in Rat
type Price big.Rat

// PriceDecode returns the price derivation encoded in Header's extra
func PriceDecode(bytes []byte) *Price {
	if len(bytes) == 0 {
		return nil
	}
	var rat big.Rat
	err := rat.GobDecode(bytes)
	if err != nil {
		log.Info("Input bytes array is not price derivation", "bytes", bytes, "error", err)
		return nil
	}
	return (*Price)(&rat)
}

// PriceEncode encodes the price derivation in Header's extra
func PriceEncode(price *Price) []byte {
	if price == nil || (*big.Rat)(price).Sign() == 0 {
		return nil
	}
	bytes, err := (*big.Rat)(price).GobEncode()
	if err != nil {
		log.Info("Failed to encode price derivation", "price", price, "error", err)
		return nil
	}
	return bytes
}

// PriceFromString decodes the price string fed from feeder service
func PriceFromString(s string) *Price {
	price, ok := new(big.Rat).SetString(s)
	if !ok {
		return nil
	}
	return (*Price)(price)
}
