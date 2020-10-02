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

package aggregator

import (
	"encoding/binary"
	"math"

	"github.com/chronowave/chronowave/ssql"
)

type min struct {
	val float64
}

func NewMin() Aggregator {
	return &min{math.Inf(+1)}
}

func DecodeMin(data []byte) Aggregator {
	v := binary.LittleEndian.Uint64(data[:8])
	return &min{math.Float64frombits(v)}
}

func (m *min) StepInt64(v int64) {
	if m.val > float64(v) {
		m.val = float64(v)
	}
}

func (m *min) StepFloat64(v float64) {
	if m.val > v {
		m.val = v
	}
}

func (m *min) DoneInt() int64 {
	return int64(m.val)
}

func (m *min) DoneFloat(v float64) float64 {
	return m.val
}

func (m *min) Bytes() []byte {
	buf := make([]byte, 9)
	buf[0] = byte(ssql.Function_MIN)
	binary.LittleEndian.PutUint64(buf[1:], math.Float64bits(m.val))
	return buf
}

func (m *min) MergeBytes(data []byte) {
	if data[0] == byte(ssql.Function_MIN) {
		v := math.Float64frombits(binary.LittleEndian.Uint64(data[1:]))
		if v < m.val {
			m.val = v
		}
	}
}

func (m *min) Merge(f Aggregator) {
	if o, ok := f.(*min); ok {
		if m.val > o.val {
			m.val = o.val
		}
	}
}
