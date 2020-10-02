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

	"github.com/chronowave/chronowave/ssql"
)

type count struct {
	cnt int64
}

func NewCount() Aggregator {
	return &count{}
}

func DecodeCount(data []byte) Aggregator {
	var v int64
	binary.LittleEndian.Uint64(data[:8])

	return &count{v}
}

func (c *count) StepInt64(int64) {
	c.cnt++
}

func (c *count) StepFloat64(float64) {
	c.cnt++
}

func (c *count) DoneInt() int64 {
	return c.cnt
}

func (c *count) DoneFloat(v float64) float64 {
	return float64(c.cnt)
}

func (c *count) Bytes() []byte {
	buf := make([]byte, 9)
	buf[0] = byte(ssql.Function_COUNT)
	binary.LittleEndian.PutUint64(buf[1:], uint64(c.cnt))
	return buf
}

func (c *count) MergeBytes(data []byte) {
	if data[0] == byte(ssql.Function_COUNT) {
		c.cnt += int64(binary.LittleEndian.Uint64(data[1:]))
	}
}

func (c *count) Merge(f Aggregator) {
	if o, ok := f.(*count); ok {
		c.cnt += o.cnt
	}
}
