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
	"math"

	"github.com/rleiwang/hfmi"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/codec"
	"github.com/chronowave/chronowave/ssd/internal"
)

func extractContent(fmi hfmi.FMI, content map[uint32][]byte) {
	for k := range content {
		if k == 0 {
			content[k] = []byte{}
			continue
		}
		if data, _, ok := fmi.ForwardExtractToChar(uint(k-1), ssd.SOH); ok {
			content[k] = concat(data)
		}
	}
}

func concat(value []byte) []byte {
	j := len(value)
	for i := len(value) - 1; i >= 0; i-- {
		if value[i] == ssd.FRAG {
			continue
		}
		j--
		value[j] = value[i]
	}

	return value[j:]
}

// bound (s, e]
func extractJson(index *ssd.IndexedBlock, parent []byte, bound []uint, padding uint) []byte {
	var (
		c     byte
		r     uint
		ok    bool
		start uint
		i     int
		eav   = make([]internal.Attribute, bound[end]-bound[beg])
		text  = make(map[uint32][]byte, len(eav))
	)

	for s, e := bound[beg]+padding, bound[end]+padding; s < e; s++ {
		code := make([]byte, 0, 64)
		ok = true
		for p := s; ok; {
			if c, r, ok = index.Entity.Access(p); !ok {
				break
			}

			if ssd.IsControlCharacter(c) {
				ok = c != ssd.AED
				break
			}

			code = append(code, c)
			if start, _, ok = index.Entity.GetBound(c); !ok {
				break
			}

			p = start + r
		}

		if ok {
			eav[i].Code = code
			eav[i].ValueType = c

			switch c {
			case ssd.TEXT:
				eav[i].Offset = index.HLT.Text.Attribute[r-1]
				eav[i].Value = uint64(index.Columnar.Text[r-1])
				text[index.Columnar.Text[r-1]] = nil
			case ssd.FLT64:
				eav[i].Offset = index.HLT.Float64.Attribute[r-1]
				eav[i].Value = math.Float64bits(index.Columnar.Float64[r-1])
			case ssd.INT64:
				eav[i].Offset = index.HLT.Int64.Attribute[r-1]
				eav[i].Value = uint64(index.Columnar.Int64[r-1])
			}

			i++
		}
	}

	if len(text) > 0 {
		extractContent(index.Content, text)
	}

	return codec.MarshalToJson(index, parent, eav[:i], text)
}
