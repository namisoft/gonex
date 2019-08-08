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
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// Data represents the external data feeded from outside
type Data struct {
	Value             interface{}
	Source            string
	DataTimestamp     time.Time
	RequestTimestamp  time.Time
	ResponseTimestamp time.Time
	reentranceFlag    int64 // prevent request routine to run twice
}

// Feeder is the main object which takes care of feeding data from outside to consensus
// engine and gathering the sealing result.
type Feeder struct {
	data sync.Map
}

func (f *Feeder) getCurrent(url string) *Data {
	value, _ := f.data.Load(url)
	if value == nil {
		// data has never be fetched before
		return nil
	}

	data := value.(*Data)
	if data.Value == nil {
		// data is being fetched the first time
		return nil
	}

	return data
}

// Yielding non-reentrant async request.
func (f *Feeder) requestUpdate(url string, parsePriceFn func([]byte) (*Data, error)) {
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
		// make sure the Body will be closed, but only after the error check
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error("Failed to read response body", "url", url, "error", err, "reponse", response)
			return
		}

		parsed, err := parsePriceFn(body)
		if err != nil {
			log.Error("Failed to parse response data body", "url", url, "error", err)
			return
		}

		data.Value = parsed.Value
		data.DataTimestamp = parsed.DataTimestamp
		data.ResponseTimestamp = parsed.ResponseTimestamp
		data.Source = parsed.Source
	}()
}
