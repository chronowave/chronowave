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
	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
)

func Nested(index *ssd.IndexedBlock, docs []internal.Document) ([]uint16, []uint16) {
	if len(docs) == 0 {
		return []uint16{}, []uint16{}
	}

	flat := make([]uint16, len(docs[0].Attribute)*2)
	cnt := len(flat) / 2
	entity, attribute := flat[:cnt], flat[cnt:]
	copy(entity, docs[0].Entity)
	copy(attribute, docs[0].Attribute)
	for _, d := range docs[1:] {
		cnt = 0
		for x, y := 0, 0; x < len(attribute) && y < len(d.Attribute); {
			if entity[x] == d.Entity[y] {
				if attribute[x] == d.Attribute[y] {
					entity[cnt], attribute[cnt] = entity[x], attribute[x]
					cnt++
					x++
					y++
				} else if attribute[x] > d.Attribute[y] {
					y++
				} else {
					x++
				}
			} else if entity[x] > d.Entity[y] {
				y++
			} else {
				x++
			}
		}
		entity, attribute = entity[:cnt], attribute[:cnt]

		if cnt == 0 {
			break
		}
	}

	return entity, attribute
}
