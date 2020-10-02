/*
 *  Copyright 2020 ChronoWave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package operator

import (
	jsonenc "encoding/json"
	"math/rand"
	"os"
	"testing"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd/index"
	"github.com/chronowave/chronowave/ssd/internal"
)

var (
	benchSize  = 1 << 16
	benchIndex *ssd.IndexedBlock
	testData   []JsonData
)

type JsonData struct {
	OneInt64 int64 `json:"a"`
}

func TestMain(m *testing.M) {
	hfmi.SetSegmentCache(uint(benchSize * 20))
	internal.FragmentSize = 2

	testData = make([]JsonData, benchSize)
	for i := range testData {
		testData[i].OneInt64 = rand.Int63()
	}

	data, err := jsonenc.Marshal(testData)
	if err != nil {
		panic(err)
	}

	parsed, err := codec.ParseJson(data)
	if err != nil {
		panic(err)
	}

	index, err := index.Build(parsed)
	if err != nil {
		panic(err)
	}

	benchIndex = index

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
