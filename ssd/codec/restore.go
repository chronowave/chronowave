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

package codec

import (
	"bytes"
	"math"
	"strconv"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
)

func MarshalToJson(index *ssd.IndexedBlock, parent []byte, eav []internal.Attribute, text map[uint32][]byte) []byte {
	var (
		depth      int
		aligned    int
		prev       = make([][][]byte, 1, len(eav))
		nested     = make([]int, 1, len(eav))
		diffs      = make([]uint16, 1, len(eav))
		w          = bytes.NewBuffer(make([]byte, 0, 2048))
		attributes = make([]map[string]bool, 1, len(eav))
	)

	attributes[0] = map[string]bool{}
	for _, attr := range eav {
		path := index.Meta.GetPath(append(parent, attr.Code...))
		path = path[len(parent):]
		if depth > 0 {
			if attr.ValueType == ssd.EOA {
				for i := 0; i < len(prev[depth]); i++ {
					w.WriteByte('}')
				}
				w.WriteByte(']')

				nested = nested[:depth]
				prev = prev[:depth]
				diffs = diffs[:depth]
				attributes = attributes[:depth]
				depth--
				continue
			}

			j := 0
			for i := 0; i <= depth; i++ {
				j += nested[i]
			}
			path = path[j:]
			if attr.Offset > diffs[depth] {
				if diffs[depth] > 0 {
					for i := 0; i < len(path); i++ {
						w.WriteByte('}')
					}
					w.WriteByte(',')
				}
				diffs[depth] = attr.Offset
				prev[depth] = nil
				for k := range attributes[depth] {
					delete(attributes[depth], k)
				}
			} else if attr.ValueType == ssd.SOA && len(prev[depth]) > 0 &&
				bytes.Equal(prev[depth][len(prev[depth])-1], path[len(path)-1]) {
				// nested
				for i := 0; i < len(prev[depth]); i++ {
					w.WriteByte('}')
				}
				w.WriteByte(',')
				prev[depth] = nil
				for k := range attributes[depth] {
					delete(attributes[depth], k)
				}
			}
		}

		aligned = 0
		if len(prev[depth]) > 0 {
			aligned = align(prev[depth], path)
			for i := len(prev[depth]) - 1; i > aligned; i-- {
				w.WriteByte('}')
			}
		}

		var prefix []byte
		for i := 0; i <= aligned && i < len(path); i++ {
			prefix = append(prefix, path[i]...)
		}
		key, sep := string(prefix), byte(',')
		if len(attributes[depth]) == 0 {
			sep = '{'
		}
		attributes[depth][key] = true
		for ; aligned < len(path); aligned++ {
			w.WriteByte(sep)
			w.WriteByte('"')
			w.Write(path[aligned])
			w.WriteByte('"')
			w.WriteByte(':')
			sep = '{'
		}
		prev[depth] = path

		switch attr.ValueType {
		case ssd.SOA:
			nested = append(nested, len(path))
			prev = append(prev, nil)
			diffs = append(diffs, 0)
			attributes = append(attributes, map[string]bool{})
			depth++
			w.WriteByte('[')
			continue
		case ssd.TEXT:
			w.WriteString(strconv.Quote(string(text[uint32(attr.Value)])))
		case ssd.FLT64:
			w.WriteString(strconv.FormatFloat(math.Float64frombits(attr.Value), 'f', -1, 64))
		case ssd.INT64:
			w.WriteString(strconv.FormatInt(int64(attr.Value), 10))
		}
	}

	for i := len(prev[depth]) - 1; i >= 0; i-- {
		w.WriteByte('}')
	}

	return w.Bytes()
}

func align(x, y [][]byte) int {
	i := 0
	for ; i < len(x) && i < len(y); i++ {
		if bytes.Compare(x[i], y[i]) != 0 {
			break
		}
	}

	return i
}
