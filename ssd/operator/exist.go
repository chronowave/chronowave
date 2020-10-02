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

func Exist(index *ssd.IndexedBlock, key []byte) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Exist panic: '%v', trace %v, block %v, path %v",
				err, string(debug.Stack()), index.ID, string(key))
		}
	}()

	// GetBound (s, e]
	_, sohEND, ok := index.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	// verify entity attribute exists, +1 count sentinel
	bound := findKeyBoundFromSOH(index.Entity, sohEND+1, key)
	if bound == nil {
		// key doesn't exist
		return []uint16{}, []uint16{}
	}

	for s, e := bound[beg]+1, bound[end]+1; s < e; s++ {
		if c, _, ok := index.Entity.Access(s); ok && ssd.TEXT <= c && c <= ssd.JSON {
			return existingEntity(index, bound, c)
		}
	}

	return []uint16{}, []uint16{}
}

func existingEntity(index *ssd.IndexedBlock, bound []uint, c byte) ([]uint16, []uint16) {
	bound = findBound(index.Entity, c, bound)
	var hlt *ssd.HeaderISA
	switch c {
	case ssd.TEXT:
		hlt = &index.HLT.Text
	case ssd.FLT64:
		hlt = &index.HLT.Float64
	case ssd.INT64:
		hlt = &index.HLT.Int64
	case ssd.BOOL:
		hlt = &index.HLT.Bool
	case ssd.NULL:
		hlt = &index.HLT.Null
	}

	if hlt != nil {
		s, _, _ := index.Entity.GetBound(c)
		entity, attribute := make([]uint16, bound[end]-bound[beg]), make([]uint16, bound[end]-bound[beg])
		copy(entity, hlt.Entity[bound[beg]-s:bound[end]-s])
		copy(attribute, hlt.Attribute[bound[beg]-s:bound[end]-s])
		return entity, attribute
	}

	return []uint16{}, []uint16{}
}
