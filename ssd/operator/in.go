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
	"runtime/debug"
	"sort"

	ext "github.com/chronowave/ext/operator"

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
)

func InFloat(index *ssd.IndexedBlock, key []byte, value []float64) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("InFloat panic: '%v', trace %v, block %v, path %v",
				err, string(debug.Stack()), index.ID, string(key))
		}
	}()

	s, _, ok := index.Entity.GetBound(ssd.FLT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// GetBound (s, e]
	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	// verify entity attribute exists, +1 count sentinel
	floats := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.FLT64)
	if floats == nil {
		// key doesn't exist
		return []uint16{}, []uint16{}
	}

	sz := floats[end] - floats[beg]
	keyBeg := floats[beg] - s
	matched := make([]uint32, sz)
	cnt := ext.InFloat64(index.Columnar.Float64[keyBeg:keyBeg+sz], value, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Float64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Float64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func InInt(index *ssd.IndexedBlock, key []byte, value []int64) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("InInt panic: '%v', trace %v, block %v, path %v",
				err, string(debug.Stack()), index.ID, string(key))
		}
	}()

	s, _, ok := index.Entity.GetBound(ssd.INT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// GetBound (s, e]
	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	// verify entity attribute exists, +1 count sentinel
	ints := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.INT64)
	if ints == nil {
		// key doesn't exist
		return []uint16{}, []uint16{}
	}

	sz := ints[end] - ints[beg]
	keyBeg := ints[beg] - s
	matched := make([]uint32, sz)
	cnt := ext.InInt64(index.Columnar.Int64[keyBeg:keyBeg+sz], value, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Int64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Int64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func InText(index *ssd.IndexedBlock, key []byte, value []string) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("InText panic: '%v', trace %v, block %v, path %v",
				err, string(debug.Stack()), index.ID, string(key))
		}
	}()

	s, _, ok := index.Entity.GetBound(ssd.TEXT)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// GetBound (s, e]
	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	// verify entity attribute exists, +1 count sentinel
	text := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.TEXT)
	if text == nil {
		// key doesn't exist
		return []uint16{}, []uint16{}
	}
	sz := text[end] - text[beg]

	matched := searchText(index, value)
	if len(matched) == 0 {
		return []uint16{}, []uint16{}
	}

	keyBeg := text[beg] - s
	x := index.Columnar.Text[keyBeg : keyBeg+sz]
	offsets := intersectUint32(x, matched)
	if len(offsets) == 0 {
		return []uint16{}, []uint16{}
	}

	entity, attribute := make([]uint16, len(offsets)), make([]uint16, len(offsets))
	for i, v := range offsets {
		entity[i] = index.HLT.Text.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Text.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func searchText(index *ssd.IndexedBlock, text []string) []uint32 {
	m := map[uint32]bool{}

	for _, t := range text {
		sanitized, prefix, suffix := sanitizePattern([]byte(t))
		if len(sanitized) == 0 {
			continue
		}

		rs := searchContent(index, sanitized, prefix, suffix)
		for _, v := range rs {
			m[v] = true
		}
	}

	keys := make([]uint32, len(m))
	j := 0
	for k := range m {
		keys[j] = k
		j++
	}

	// asc
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
