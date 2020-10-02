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

type average struct {
	sum   float64
	count uint64
}

func NewAverage() Aggregator {
	return &average{}
}

func DecodeAverage(data []byte) Aggregator {
	return &average{
		sum:   math.Float64frombits(binary.LittleEndian.Uint64(data[1:9])),
		count: binary.LittleEndian.Uint64(data[9:]),
	}
}

func (a *average) StepInt64(v int64) {
	a.sum += float64(v)
	a.count++
}

func (a *average) StepFloat64(v float64) {
	a.sum += v
	a.count++
}

func (a *average) DoneInt() int64 {
	return 0
}

func (a *average) DoneFloat(v float64) float64 {
	if a.count == 0 {
		return 0
	}

	return a.sum / float64(a.count)
}

func (a *average) Bytes() []byte {
	o := make([]byte, 17)
	o[0] = byte(ssql.Function_AVG)
	binary.LittleEndian.PutUint64(o[1:9], math.Float64bits(a.sum))
	binary.LittleEndian.PutUint64(o[9:], a.count)

	return o
}

func (a *average) MergeBytes(data []byte) {
	if data[0] == byte(ssql.Function_AVG) && len(data) == 17 {
		a.sum += math.Float64frombits(binary.LittleEndian.Uint64(data[1:9]))
		a.count += binary.LittleEndian.Uint64(data[9:])
	}
}

func (a *average) Merge(m Aggregator) {
	if o, ok := m.(*average); ok {
		a.sum += o.sum
		a.count += o.count
	}
}
