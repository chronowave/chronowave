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

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
)

func UnaryInt64(index *ssd.IndexedBlock, key []byte, value int64, op func([]int64, int64, []uint32) uint) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Unary int64 panic: '%v', trace %v, block %v, path %v, value %v",
				err, string(debug.Stack()), index.ID, string(key), value)
		}
	}()

	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	s, _, ok := index.Entity.GetBound(ssd.INT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// verify entity attribute exists, +1 count sentinel
	ints := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.INT64)
	if len(ints) == 0 {
		// key doesn't exist
		log.Info().Msgf("path %v with TEXT value doesn't exist in block %v", string(key), index.ID)
		return []uint16{}, []uint16{}
	}
	sz := ints[end] - ints[beg]
	keyBeg := ints[beg] - s
	matched := make([]uint32, sz)
	cnt := op(index.Columnar.Int64[keyBeg:keyBeg+sz], value, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Int64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Int64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func UnaryFloat64(index *ssd.IndexedBlock, key []byte, value float64, op func([]float64, float64, []uint32) uint) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Unary float panic: '%v', trace %v, block %v, path %v, value %v",
				err, string(debug.Stack()), index.ID, string(key), value)
		}
	}()

	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	s, _, ok := index.Entity.GetBound(ssd.FLT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// verify entity attribute exists, +1 count sentinel
	floats := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.FLT64)
	if len(floats) == 0 {
		// key doesn't exist
		log.Info().Msgf("path %v with TEXT value doesn't exist in block %v", string(key), index.ID)
		return []uint16{}, []uint16{}
	}
	sz := floats[end] - floats[beg]
	keyBeg := floats[beg] - s
	matched := make([]uint32, sz)
	cnt := op(index.Columnar.Float64[keyBeg:keyBeg+sz], value, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Float64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Float64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func BinaryInt64(index *ssd.IndexedBlock, key []byte, x, y int64, op func([]int64, int64, int64, []uint32) uint) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("binary int64 panic: '%v', trace %v, block %v, path %v, [%v, %v]",
				err, string(debug.Stack()), index.ID, string(key), x, y)
		}
	}()

	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	s, _, ok := index.Entity.GetBound(ssd.INT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// verify entity attribute exists, +1 count sentinel
	ints := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.INT64)
	if len(ints) == 0 {
		return []uint16{}, []uint16{}
	}
	sz := ints[end] - ints[beg]
	keyBeg := ints[beg] - s
	matched := make([]uint32, sz)
	cnt := op(index.Columnar.Int64[keyBeg:keyBeg+sz], x, y, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Int64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Int64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}

func BinaryFloat64(index *ssd.IndexedBlock, key []byte, x, y float64, op func([]float64, float64, float64, []uint32) uint) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("binary int64 panic: '%v', trace %v, block %v, path %v, [%v, %v]",
				err, string(debug.Stack()), index.ID, string(key), x, y)
		}
	}()

	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	s, _, ok := index.Entity.GetBound(ssd.FLT64)
	if !ok {
		return []uint16{}, []uint16{}
	}

	// verify entity attribute exists, +1 count sentinel
	ints := findEndOfKeyBound(index.Entity, sohEND+1, key, ssd.FLT64)
	if len(ints) == 0 {
		// key doesn't exist
		log.Info().Msgf("path %v with TEXT value doesn't exist in block %v", string(key), index.ID)
		return []uint16{}, []uint16{}
	}
	sz := ints[end] - ints[beg]
	keyBeg := ints[beg] - s
	matched := make([]uint32, sz)
	cnt := op(index.Columnar.Float64[keyBeg:keyBeg+sz], x, y, matched)

	entity, attribute := make([]uint16, cnt), make([]uint16, cnt)
	for i, v := range matched[:cnt] {
		entity[i] = index.HLT.Float64.Entity[keyBeg+uint(v)]
		attribute[i] = index.HLT.Float64.Attribute[keyBeg+uint(v)]
	}

	return entity, attribute
}
