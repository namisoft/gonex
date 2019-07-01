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
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

const (
	priceServiceURL       = "http://localhost:3000/price/NUSD_USD"
	nonCanonicalCacheSize = params.CanonicalDepth
	medianPriceCacheSize  = 6
)

// PriceData represents the external price feeded from outside
type PriceData struct {
	Value     json.Number `json:"price"`
	Timestamp int64       `json:"timestamp"`
	Exchange  string      `json:"exchange"`
}

type PriceEngine struct {
	feeder       *Feeder
	ticker       *time.Ticker
	canonPrices  *lru.Cache // canonical prices: number -> Price
	nonacPrices  *lru.Cache // non-canonical prices: hash -> Price
	medianPrices *lru.Cache // calculated median price: hash -> Price
	ttl          time.Duration
	config       *params.DccsConfig
}

func newPriceEngine(conf *params.DccsConfig) *PriceEngine {
	priceSamplingInterval := time.Duration(conf.PriceSamplingInterval*conf.Period) * time.Second

	// the longest time for a price to stay valid = max(blocktime, priceSamplingInterval / 2)
	ttl := priceSamplingInterval / 2
	if ttl < time.Duration(conf.Period) {
		ttl = time.Duration(conf.Period)
	}

	e := &PriceEngine{
		feeder: &Feeder{},
		ticker: time.NewTicker(priceSamplingInterval / 3),
		ttl:    ttl,
		config: conf,
	}

	var err error

	maxPriceCount := int(conf.PriceSamplingDuration / conf.PriceSamplingInterval)
	e.canonPrices, err = lru.New(maxPriceCount) // add some extra buffer for sidechain values
	if err != nil {
		log.Crit("Unable to create the canonical price cache", "Endurio block", conf.EndurioBlock, "pricesCount", (conf.PriceSamplingDuration / conf.PriceSamplingInterval), "error", err)
		return nil
	}

	e.nonacPrices, err = lru.New(nonCanonicalCacheSize)
	if err != nil {
		log.Crit("Unable to create the non-canonical price cache", "Endurio block", conf.EndurioBlock, "nonCanonicalCacheSize", nonCanonicalCacheSize, "error", err)
		return nil
	}

	e.medianPrices, err = lru.New(medianPriceCacheSize)
	if err != nil {
		log.Crit("Unable to create the median price cache", "Endurio block", conf.EndurioBlock, "medianPriceCacheSize", medianPriceCacheSize, "error", err)
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

type ByPrice []*Price

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Less(i, j int) bool { return a[i].Rat().Cmp(a[j].Rat()) < 0 }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// CalcMedianPrice calculates the median price of a price block and cache it.
func (e *PriceEngine) CalcMedianPrice(chain consensus.ChainReader, number uint64) (*Price, error) {
	if !e.config.IsPriceBlock(number) {
		// not a price block
		return nil, errors.New("Not a price block")
	}
	header := chain.GetHeaderByNumber(number)
	if header == nil {
		return nil, errors.New("Block number too high")
	}
	if median, ok := e.medianPrices.Get(header.Hash()); ok {
		// cache found
		return median.(*Price), nil
	}
	cap := int(e.config.PriceSamplingDuration / e.config.PriceSamplingInterval)
	prices := make([]*Price, 0, cap)
	for n := number; n > number-e.config.PriceSamplingDuration; n -= e.config.PriceSamplingInterval {
		price := e.GetBlockPrice(chain, n)
		if price != nil {
			prices = append(prices, price)
		}
	}
	count := len(prices)
	if count*3 < cap*2 {
		// require atleast 2/3 of maximum price feed
		// TODO: make this configurable
		return nil, errors.New("Not enough block with price to come to a consensus")
	}
	sort.Sort(ByPrice(prices))
	if count&1 == 1 {
		return prices[count/2], nil
	}
	median := new(big.Rat).Add(prices[count/2-1].Rat(), prices[count/2].Rat())
	median.Mul(median, common.Rat1_2)
	e.medianPrices.Add(header.Hash(), median)
	return (*Price)(median), nil
}

func (e *PriceEngine) GetBlockPrice(chain consensus.ChainReader, number uint64) *Price {
	if !e.config.IsPriceBlock(number) {
		// not a price block
		return nil
	}
	currentNumber := chain.CurrentHeader().Number.Uint64()
	if number >= currentNumber-params.CanonicalDepth {
		// cache non-canonical price by block hash
		header := chain.GetHeaderByNumber(number)
		if header == nil {
			log.Error("PriceEngine.GetBlockPrice: failed to get header by number ", "number", number)
			return nil
		}
		hash := header.Hash()
		price, ok := e.nonacPrices.Get(hash)
		if ok {
			// non-canonical cache found
			return price.(*Price)
		}
		log.Info("Fetching non-canonical block price", "number", number)
		price = PriceDecodeFromExtra(header.Extra)
		if price == nil {
			log.Error("PriceEngine.GetBlockPrice: failed to decode price from non-canonical header extra", "number", number, "extra", header.Extra)
			return nil
		}
		e.nonacPrices.Add(hash, price)
		return price.(*Price)
	}

	// cache canonical price by block number
	price, ok := e.canonPrices.Get(number)
	if ok {
		// canonical cache found
		return price.(*Price)
	}
	header := chain.GetHeaderByNumber(number)
	log.Info("Fetching canonical block price", "number", number)
	price = PriceDecodeFromExtra(header.Extra)
	if price == nil {
		log.Error("PriceEngine.GetBlockPrice: failed to decode price from canonical header extra", "number", number, "extra", header.Extra)
		return nil
	}
	e.canonPrices.Add(number, price)
	return price.(*Price)
}

func (e *PriceEngine) CurrentPrice() *Price {
	data := e.feeder.getCurrent(priceServiceURL)
	if data == nil {
		e.feeder.requestUpdate(priceServiceURL, parsePriceFn)
		return nil
	}
	if time.Now().Sub(data.ResponseTimestamp) > e.ttl {
		// expired data
		return nil
	}
	return data.Value.(*Price)
}

func parsePriceFn(body []byte) (*Data, error) {
	var priceData PriceData

	if err := json.Unmarshal(body, &priceData); err != nil {
		log.Error("Failed to unmarshal price json", "error", err, "body", string(body))
		return nil, err
	}

	log.Trace("PriceData", "priceData", priceData)

	price := PriceFromString(priceData.Value.String())
	if price == nil {
		log.Error("Failed to parse price value", "priceData.Value", priceData.Value)
		return nil, errors.New("Not a price value")
	}

	return &Data{
		Value:             price,
		DataTimestamp:     time.Unix(priceData.Timestamp, 0),
		ResponseTimestamp: time.Now(),
		Source:            priceData.Exchange,
	}, nil
}

// Price encoded in Rat
type Price big.Rat

// PriceDerivation is (Price - 1/1)
type PriceDerivation Price

func (p *PriceDerivation) Rat() *big.Rat {
	return (*big.Rat)(p)
}

func (p *Price) Rat() *big.Rat {
	return (*big.Rat)(p)
}

func (p *Price) Derivation() *PriceDerivation {
	return (*PriceDerivation)(p.Rat().Sub(p.Rat(), common.Rat1))
}

// PriceDecodeFromExtra returns the price derivation encoded in Header's extra
// extra = [vanity(32), price(...), signature(65)]
func PriceDecodeFromExtra(extra []byte) *Price {
	extraSuffix := len(extra) - extraSeal
	extraBytes := extra[extraVanity:extraSuffix]
	return PriceDecode(extraBytes)
}

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

