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
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

const feederServiceURL = "http://localhost:3000/price/NUSD_USD"

// Data represents the external data feeded from outside
type Data struct {
	Value             interface{}
	Source            string
	DataTimestamp     time.Time
	RequestTimestamp  time.Time
	ResponseTimestamp time.Time
	reentranceFlag    int64 // prevent request routine to run twice
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
	data   sync.Map
	ticker *time.Ticker
}

func newFeeder(interval time.Duration) *feeder {
	f := &feeder{
		ticker: time.NewTicker(interval),
	}
	go f.fetchingLoop()
	return f
}

func (f *feeder) fetchingLoop() {
	for range f.ticker.C {
		f.data.Range(func(key interface{}, _ interface{}) bool {
			f.requestUpdate(key.(string))
			return true
		})
	}
}

func (f *feeder) getCurrent(url string) *Data {
	value, _ := f.data.Load(url)
	data := value.(*Data)

	if data.Value == nil {
		// data is being fetched the first time
		return nil
	}

	return data
}

// Yielding non-reentrant async request.
func (f *feeder) requestUpdate(url string) {
	value, _ := f.data.LoadOrStore(url, &Data{RequestTimestamp: time.Now()})
	data := value.(*Data)

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

		data.Value = &price
		data.DataTimestamp = time.Unix(priceData.Timestamp, 0)
		data.ResponseTimestamp = time.Now()
		data.Source = priceData.Exchange
	}()
}

func (f *feeder) Price() *Price {
	data := f.getCurrent(feederServiceURL)
	if data == nil {
		// first request
		f.requestUpdate(feederServiceURL)
		return nil
	}
	return data.Value.(*Price)
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
