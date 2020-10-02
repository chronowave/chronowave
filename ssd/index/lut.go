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
	"errors"

	"github.com/chronowave/chronowave/ssd"
)

var (
	ErrNoSOH            = errors.New("no SOH headers")
	ErrNotFound         = errors.New("not found char")
	ErrDocumentOverflow = errors.New("number of attributes in a document exceeds 65536")
)

type handler func(uint, uint16, uint16)

// fillHLT fill Header Lookup Table
func fillHLT(parsed *ssd.ParsedBlock, indexed *ssd.IndexedBlock) error {
	_, sohEnd, ok := indexed.Entity.GetBound(ssd.SOH)
	if !ok {
		sohEnd = 0
	}
	// SENTINEL
	sohEnd++

	rows := make([]uint, 32)
	handlers := make([]handler, 32)

	if beg, end, ok := indexed.Entity.GetBound(ssd.TEXT); ok {
		sz := end - beg
		buf := make([]uint16, 2*sz)
		indexed.HLT.Text = ssd.HeaderISA{Entity: buf[:sz], Attribute: buf[sz:]}
		handlers[ssd.TEXT] = func(rank uint, doc, offset uint16) {
			indexed.Columnar.Text[rank] = parsed.Columnar.Text[rows[ssd.TEXT]]
			indexed.HLT.Text.Entity[rank] = doc
			indexed.HLT.Text.Attribute[rank] = offset
		}
	}

	if beg, end, ok := indexed.Entity.GetBound(ssd.FLT64); ok {
		sz := end - beg
		buf := make([]uint16, 2*sz)
		indexed.HLT.Float64 = ssd.HeaderISA{Entity: buf[:sz], Attribute: buf[sz:]}

		handlers[ssd.FLT64] = func(rank uint, doc, offset uint16) {
			indexed.Columnar.Float64[rank] = parsed.Columnar.Float64[rows[ssd.FLT64]]
			indexed.HLT.Float64.Entity[rank] = doc
			indexed.HLT.Float64.Attribute[rank] = offset
		}
	}

	if beg, end, ok := indexed.Entity.GetBound(ssd.INT64); ok {
		sz := end - beg
		buf := make([]uint16, 2*sz)
		indexed.HLT.Int64 = ssd.HeaderISA{Entity: buf[:sz], Attribute: buf[sz:]}

		handlers[ssd.INT64] = func(rank uint, doc, offset uint16) {
			indexed.Columnar.Int64[rank] = parsed.Columnar.Int64[rows[ssd.INT64]]
			indexed.HLT.Int64.Entity[rank] = doc
			indexed.HLT.Int64.Attribute[rank] = offset
		}
	}

	if beg, end, ok := indexed.Entity.GetBound(ssd.BOOL); ok {
		sz := end - beg
		buf := make([]uint16, 2*sz)
		indexed.HLT.Bool = ssd.HeaderISA{Entity: buf[:sz], Attribute: buf[sz:]}

		handlers[ssd.BOOL] = func(rank uint, doc, offset uint16) {
			indexed.Columnar.Bool[rank] = parsed.Columnar.Bool[rows[ssd.BOOL]]
			indexed.HLT.Bool.Entity[rank] = doc
			indexed.HLT.Bool.Attribute[rank] = offset
		}
	}

	if beg, end, ok := indexed.Entity.GetBound(ssd.NULL); ok {
		sz := end - beg
		buf := make([]uint16, 2*sz)
		indexed.HLT.Null = ssd.HeaderISA{Entity: buf[:sz], Attribute: buf[sz:]}

		handlers[ssd.NULL] = func(rank uint, doc, offset uint16) {
			indexed.HLT.Null.Entity[rank] = doc
			indexed.HLT.Null.Attribute[rank] = offset
		}
	}

	var (
		ent    uint16
		nested []uint16
		attr   uint16
	)
	indexed.EntityID[ent] = 0
	// [0, end)
	for p := uint(0); p < sohEnd; p++ {
		a, r, ok := indexed.Entity.Access(p)
		if !ok {
			return ErrNotFound
		}

		if ssd.IsControlCharacter(a) {
			if a == ssd.EOO {
				// 0-th offset in SOH
				ent++
				indexed.EntityID[ent], attr, nested = uint32(p+1), 0, nil
			} else if a == ssd.AED {
				diff := p - uint(indexed.EntityID[ent])
				if diff >= 1<<16 {
					return ErrDocumentOverflow
				}
				attr = uint16(diff)
				nested[len(nested)-1] = attr
			}
			continue
		}

		for {
			s, _, ok := indexed.Entity.GetBound(a)
			if !ok {
				return ErrNotFound
			}

			a, r, ok = indexed.Entity.Access(s + r)
			if !ok {
				return ErrNotFound
			}

			if a <= ssd.SOH {
				break
			} else if a == ssd.SOA {
				diff := p - uint(indexed.EntityID[ent])
				if diff >= 1<<16 {
					return ErrDocumentOverflow
				}
				attr = uint16(diff) + 1
				nested = append(nested, attr)
			} else if a == ssd.EOA {
				sz := len(nested) - 1
				nested = nested[:sz]
				if sz > 0 {
					attr = nested[sz-1]
				} else {
					attr = 0
				}
			} else if ssd.TEXT <= a && a <= ssd.NULL {
				handlers[a](r-1, ent, attr)
				rows[a]++
				break
			}
		}
	}

	return nil
}

func fillDocumentArray(indexed *ssd.IndexedBlock) error {
	_, sohEND, ok := indexed.Content.GetBound(ssd.SOH)
	if !ok {
		sohEND = 0
	}

	// count sentinel
	indexed.HeaderDA = make([]uint32, sohEND)

	if fragBEG, fragEND, ok := indexed.Content.GetBound(ssd.FRAG); ok && fragEND > fragBEG {
		indexed.FragDA = make([]uint32, fragEND-fragBEG)
	}

	// [0, end)
	for p := uint(0); p <= sohEND; p++ {
		a, r, ok := indexed.Content.Access(p)
		if !ok {
			return ErrNotFound
		}

		for {
			if a == ssd.SENTINEL {
				// last one
				break
			} else if a == ssd.SOH {
				indexed.HeaderDA[r-1] = uint32(p)
				break
			} else if a == ssd.FRAG {
				indexed.FragDA[r-1] = uint32(p)
			}

			s, _, ok := indexed.Content.GetBound(a)
			if !ok {
				return ErrNotFound
			}

			a, r, ok = indexed.Content.Access(s + r)
			if !ok {
				return ErrNotFound
			}
		}
	}

	return nil
}
