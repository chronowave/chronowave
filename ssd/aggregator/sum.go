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

type sum struct {
	sum float64
}

func NewSum() Aggregator {
	return &sum{}
}

func DecodeSum(data []byte) Aggregator {
	v := binary.LittleEndian.Uint64(data[:8])
	return &sum{math.Float64frombits(v)}
}

func (s *sum) StepInt64(v int64) {
	s.sum += float64(v)
}

func (s *sum) StepFloat64(v float64) {
	s.sum += v
}

func (s *sum) DoneInt() int64 {
	return int64(s.sum)
}

func (s *sum) DoneFloat(v float64) float64 {
	return s.sum
}

func (s *sum) Bytes() []byte {
	buf := make([]byte, 9)
	buf[0] = byte(ssql.Function_SUM)
	binary.LittleEndian.PutUint64(buf[1:], math.Float64bits(s.sum))
	return buf
}

func (s *sum) MergeBytes(data []byte) {
	if data[0] == byte(ssql.Function_SUM) {
		v := binary.LittleEndian.Uint64(data[1:])
		s.sum += math.Float64frombits(v)
	}
}

func (s *sum) Merge(m Aggregator) {
	if o, ok := m.(*sum); ok {
		s.sum += o.sum
	}
}
