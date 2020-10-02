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

package index

import (
	"runtime/debug"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
)

func Build(parsed *ssd.ParsedBlock) (*ssd.IndexedBlock, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("error %v filling HLT, trace %v", err, string(debug.Stack()))
		}
	}()
	index := &ssd.IndexedBlock{
		EntityID: make([]uint32, parsed.Count),
		Meta:     parsed.Meta,
	}

	index.Columnar.Float64 = make([]float64, len(parsed.Columnar.Float64))
	index.Columnar.Int64 = make([]int64, len(parsed.Columnar.Int64))
	index.Columnar.Text = make([]uint32, len(parsed.Columnar.Text))
	index.Columnar.Bool = make([]bool, len(parsed.Columnar.Bool))

	buf := parsed.Content.Value()
	var err error
	if len(buf) > 0 {
		index.Content = hfmi.New(parsed.Content.Value())
		err = fillDocumentArray(index)
	}

	if err != nil {
		return nil, err
	}

	buf = parsed.Entity.Bytes()
	if len(buf) > 0 {
		index.Entity = hfmi.New(parsed.Entity.Bytes())
		err = fillHLT(parsed, index)
	}

	return index, err
}
