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

import "github.com/rleiwang/hfmi"

const (
	beg = iota
	end
)

func findEndOfKeyBound(fmi hfmi.FMI, e uint, key []byte, c byte) []uint {
	if len(key) == 0 {
		return nil
	}

	// in SOH block
	bound := findEndRange(fmi, key[0], e)
	if bound == nil {
		return nil
	}
	if len(key) > 1 {
		bound = findKeyRange(fmi, bound, key[1:], c)
	} else {
		bound = findBound(fmi, c, bound)
	}

	return bound
}

func findEndRange(fmi hfmi.FMI, c byte, end uint) []uint {
	s, _, ok := fmi.GetBound(c)
	if !ok {
		return nil
	}
	rend, ok := fmi.Rank(c, end)
	if !ok || rend == 0 {
		// not found
		return nil
	}

	return []uint{s, s + rend}
}

// cc: control character, returns starting range of cc
func findKeyRange(fmi hfmi.FMI, bound []uint, key []byte, cc byte) []uint {
	for _, k := range key {
		bound = findBound(fmi, k, bound)
		if bound == nil {
			return nil
		}
	}

	return findBound(fmi, cc, bound)
}

// (s, e]
func findBound(fmi hfmi.FMI, c byte, bound []uint) []uint {
	// (s, e]
	s, _, ok := fmi.GetBound(c)
	if !ok {
		return nil
	}
	rbeg, ok := fmi.Rank(c, bound[beg])
	if !ok {
		// not found
		return nil
	}
	rend, ok := fmi.Rank(c, bound[end])
	if !ok || rbeg == rend {
		// not found
		return nil
	}

	return []uint{s + rbeg, s + rend}
}

func findKeyBoundFromSOH(fmi hfmi.FMI, e uint, key []byte) []uint {
	if len(key) == 0 {
		return nil
	}

	// in SOH block
	bound := findEndRange(fmi, key[0], e)
	if bound != nil {
		for _, k := range key[1:] {
			bound = findBound(fmi, k, bound)
			if bound == nil {
				return nil
			}
		}
	}

	return bound
}

func findKeyBound(fmi hfmi.FMI, bound []uint, key []byte) []uint {
	for _, k := range key {
		if bound = findBound(fmi, k, bound); bound == nil {
			break
		}
	}

	return bound
}
