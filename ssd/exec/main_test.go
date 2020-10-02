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

package exec

import (
	"os"
	"testing"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/index"
)

func TestMain(m *testing.M) {
	hfmi.SetSegmentCache(2378)
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func buildTestIndex(parsed *ssd.ParsedBlock) (*ssd.IndexedBlock, error) {
	indexed, err := index.Build(parsed)
	if err != nil {
		return nil, err
	}

	buf, err := index.EncodeIndexBlock(indexed)
	if err != nil {
		return nil, err
	}

	clone := make([]byte, len(buf))
	copy(clone, buf)

	return index.DecodeIndexBlock(clone)
}
