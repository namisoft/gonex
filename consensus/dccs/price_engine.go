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
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

const priceServiceURL = "http://localhost:3000/price/NUSD_USD"

// PriceData represents the external price feeded from outside
type PriceData struct {
	Value     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	Exchange  string `json:"exchange"`
}

type PriceEngine struct {
	feeder *Feeder
	ticker *time.Ticker
	prices *lru.Cache
}

func newPriceEngine(conf *params.DccsConfig) *PriceEngine {
	maxPriceCount := int(conf.PriceDuration / conf.PriceInterval)
	priceLRU, err := lru.New(maxPriceCount)
	if err != nil {
		log.Crit("Unable to create price LRU", "Endurio block", conf.EndurioBlock, "pricesCount", (conf.PriceDuration / conf.PriceInterval), "error", err)
		return nil
	}

	priceInterval := time.Duration(conf.PriceInterval*conf.Period) * time.Second

	e := &PriceEngine{
		feeder: &Feeder{},
		ticker: time.NewTicker(priceInterval / 3),
		prices: priceLRU,
	}

	go e.fetchingLoop()
	return e
}

func (e *PriceEngine) fetchingLoop() {
	for range e.ticker.C {
		e.feeder.data.Range(func(key interface{}, _ interface{}) bool {
			e.feeder.requestUpdate(key.(string), parsePriceFn)
			return true
		})
	}
}

func (e *PriceEngine) getBlockPrice(number uint64) (*Price, bool) {
	price, ok := e.prices.Get(number)
	if !ok {
		// TODO: read and cache the price from headers
		return nil, ok
	}
	return price.(*Price), ok
}

func (e *PriceEngine) CurrentPrice() *Price {
	data := e.feeder.getCurrent(priceServiceURL)
	if data == nil {
		// first request
		e.feeder.requestUpdate(priceServiceURL, parsePriceFn)
		return nil
	}
	return data.Value.(*Price)
}

func parsePriceFn(body []byte) (*Data, error) {
	var priceData PriceData

	if err := json.Unmarshal(body, &priceData); err != nil {
		log.Error("Failed to unmarshal price json", "error", err, "body", body)
		return nil, err
	}

	log.Trace("PriceData", "priceData", priceData)

	price := PriceFromString(priceData.Value)
	if price == nil {
		log.Error("Failed to parse price value", "priceData.Value", priceData.Value)
		return nil, errors.New("Not a price value")
	}

	return &Data{
		Value:             &price,
		DataTimestamp:     time.Unix(priceData.Timestamp, 0),
		ResponseTimestamp: time.Now(),
		Source:            priceData.Exchange,
	}, nil
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
