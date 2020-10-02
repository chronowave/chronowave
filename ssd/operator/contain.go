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

	"github.com/rleiwang/hfmi"

	"github.com/rs/zerolog/log"

	"github.com/chronowave/chronowave/ssd"
)

const (
	escapeChar = '\\'
)

func Contain(index *ssd.IndexedBlock, key, pattern []byte) ([]uint16, []uint16) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Contains panic: '%v', trace %v, block %v, path %v, pattern %v",
				err, string(debug.Stack()), index.ID, string(key), string(pattern))
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
	if len(text) == 0 {
		// key doesn't exist
		log.Info().Msgf("path %v with TEXT value doesn't exist in block %v", string(key), index.ID)
		return []uint16{}, []uint16{}
	}

	sanitized, prefix, suffix := sanitizePattern(pattern)
	if len(sanitized) == 0 {
		log.Info().Msgf("sanitized search pattern %v is empty", string(pattern))
		return []uint16{}, []uint16{}
	}

	matched := searchContent(index, sanitized, prefix, suffix)
	if len(matched) == 0 {
		return []uint16{}, []uint16{}
	}

	sz := text[end] - text[beg]
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

// returns pattern: text pattern to search, prefix: true if search prefix, suffix: true if search suffix
func sanitizePattern(pattern []byte) ([]byte, bool, bool) {
	prefix, suffix := pattern[0] == '^', pattern[len(pattern)-1] == '$'
	if suffix {
		if len(pattern) > 1 && pattern[len(pattern)-2] == escapeChar {
			// pattern is "...\$"
			suffix = false
		} else {
			// strip off $
			pattern = pattern[:len(pattern)-1]
		}
	}

	if prefix {
		// strip off ^
		pattern = pattern[1:]
	}

	// unescape
	idx, lastb := 0, byte(0)
	for _, b := range pattern {
		if b == escapeChar {
			if lastb == escapeChar {
				// two escape chars \\
				pattern[idx] = b
				idx++
				lastb = byte(0)
			}
			continue
		}
		lastb = b
		if b == '*' {
			// wild card
			b = ssd.SENTINEL
		}
		pattern[idx] = b
		idx++
	}

	// trim wild card from right
	i := len(pattern) - 1
	for i >= 0 && pattern[i] == ssd.SENTINEL {
		i--
	}
	if i < len(pattern)-1 {
		pattern = pattern[:i+1]
	}

	// trim wild card from left
	i = 0
	for i < len(pattern) && pattern[i] == ssd.SENTINEL {
		i++
	}
	if i > 0 {
		pattern = pattern[i:]
	}

	return pattern, prefix, suffix
}

func searchContent(index *ssd.IndexedBlock, sanitized []byte, prefix, suffix bool) []uint32 {
	_, sohEND, ok := index.Content.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	var textSearchBound []uint
	if prefix {
		// prefix, starts from soh (s, e], +1 to count sentinel
		textSearchBound = findEndRange(index.Content, sanitized[0], sohEND+1)
	} else {
		// find the first character in pattern
		if s, e, ok := index.Content.GetBound(sanitized[0]); ok {
			textSearchBound = []uint{s, e}
		}
	}

	if len(textSearchBound) == 0 {
		return []uint32{}
	}

	matched := make([]uint32, sohEND+1)
	onMatch := func(marker byte, bound []uint) {
		// bound: (s, e]
		switch marker {
		case ssd.SENTINEL:
			// this is the last one
			matched[sohEND] = 1
		case ssd.SOH:
			// note: soh bound (s, sohEND]
			for s, e := bound[beg], bound[end]; s < e; s++ {
				// get entity id from header document array
				matched[index.HeaderDA[s]] = 1
			}
		case ssd.FRAG:
			frag, _, _ := index.Content.GetBound(ssd.FRAG)
			// note: calculate the bound
			for s, e := bound[beg], bound[end]; s < e; s++ {
				// get entity id from fragment document array
				matched[index.FragDA[s-frag]] = 1
			}
		}
	}

	textSearch(index.Content, textSearchBound, sanitized[1:], suffix, onMatch)

	cnt := uint32(0)
	for i, v := range matched {
		// note: text 0 is empty
		matched[cnt] = uint32(i) + 1
		cnt += v
	}

	return matched[:cnt]
}

// full text pattern search will return key codes
func textSearch(fmi hfmi.FMI, bound []uint, pattern []byte, suffix bool, onMatch func(byte, []uint)) {
	for len(pattern) > 0 {
		if pattern[0] == ssd.SENTINEL {
			// contains wild card search
			// note: trim off wild from head and tail, there must have extra chars
			dfs(fmi, bound, pattern[1:], suffix, onMatch)
			return
		}

		if frag := findBound(fmi, ssd.FRAG, bound); frag != nil {
			textSearch(fmi, frag, pattern, suffix, onMatch)
		}

		bound = findBound(fmi, pattern[0], bound)
		if bound == nil {
			return
		}

		pattern = pattern[1:]
	}

	// pattern matched
	if suffix {
		for _, m := range []byte{ssd.SENTINEL, ssd.SOH} {
			terminate := findBound(fmi, m, bound)
			if len(terminate) > 0 {
				onMatch(m, terminate)
			}
		}
	} else {
		dfs(fmi, bound, nil, suffix, onMatch)
	}
}

// depth first search
func dfs(fmi hfmi.FMI, bound []uint, pattern []byte, suffix bool, onMatch func(byte, []uint)) {
	v, ni, bounds, terminated := 0, []int{0}, [][]uint{bound}, []uint{0}
	checked, est, chars := uint(0), bound[end]-bound[beg], [][]byte{fmi.CharsInBound(bound[beg], bound[end])}

	hasMore := func() bool {
		for v >= 0 {
			if ni[v] < len(chars[v]) {
				return true
			}
			v--
		}

		return false
	}

	// depth first
	for est > checked && hasMore() {
		if terminated[v] == 0 {
			for _, m := range []byte{ssd.SENTINEL, ssd.SOH, ssd.FRAG} {
				terminate := findBound(fmi, m, bounds[v])
				if len(terminate) > 0 {
					terminated[v] += terminate[end] - terminate[beg]
					if pattern == nil {
						onMatch(m, terminate)
					}
				}
			}

			checked += terminated[v]
			if terminated[v] == bounds[v][end]-bounds[v][beg] {
				v--
				continue
			}

			// assign non zero value
			terminated[v] = 1
		}

		k := chars[v][ni[v]]
		ni[v]++
		nb := findBound(fmi, k, bounds[v])
		if nb == nil {
			continue
		}

		if pattern != nil && k == pattern[0] {
			checked += nb[end] - nb[beg]
			textSearch(fmi, nb, pattern[1:], suffix, onMatch)
			v--
			continue
		}

		v++
		nc := fmi.CharsInBound(nb[beg], nb[end])
		if len(ni) > v {
			bounds[v] = nb
			chars[v] = nc
			ni[v] = 0
			terminated[v] = 0
		} else {
			bounds = append(bounds, nb)
			chars = append(chars, nc)
			ni = append(ni, 0)
			terminated = append(terminated, 0)
		}
	}
}
