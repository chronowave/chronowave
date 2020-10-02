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
	"bytes"

	"github.com/chronowave/chronowave/ssql"

	"github.com/circonus-labs/circonusllhist"
)

type percentile struct {
	histogram *circonusllhist.Histogram
}

func NewPercentile() Aggregator {
	return &percentile{circonusllhist.NewNoLocks()}
}

func DecodePercentile(data []byte) Aggregator {
	hist, err := circonusllhist.Deserialize(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return &percentile{histogram: hist}
}

func (p *percentile) StepInt64(int64) {
	//p.histogram.Merge()

}

func (p *percentile) StepFloat64(float64) {
}

func (p *percentile) DoneInt() int64 {
	return 0
}

func (p *percentile) DoneFloat(v float64) float64 {
	return p.histogram.ValueAtQuantile(v)
}

func (p *percentile) Bytes() []byte {
	o := bytes.NewBuffer(make([]byte, 0, 1024))
	p.histogram.Serialize(o)
	return append([]byte{byte(ssql.Function_PCTL)}, o.Bytes()...)
}

func (p *percentile) MergeBytes(data []byte) {
	if data[0] == byte(ssql.Function_PCTL) {
		hist, err := circonusllhist.Deserialize(bytes.NewReader(data[1:]))
		if err != nil {
			panic(err)
		}
		p.histogram.Merge(hist)
	}
}

func (p *percentile) Merge(m Aggregator) {
	if o, ok := m.(*percentile); ok {
		p.histogram.Merge(o.histogram)
	}
}
