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

import "github.com/chronowave/chronowave/ssql"

type Aggregator interface {
	StepInt64(int64)
	StepFloat64(float64)
	DoneInt() int64
	DoneFloat(float64) float64
	Bytes() []byte
	MergeBytes([]byte)
	Merge(Aggregator)
}

func FromBytes(buf []byte) Aggregator {
	switch buf[0] {
	case byte(ssql.Function_AVG):
		return DecodeAverage(buf[1:])
	case byte(ssql.Function_SUM):
		return DecodeSum(buf[1:])
	case byte(ssql.Function_MIN):
		return DecodeMin(buf[1:])
	case byte(ssql.Function_MAX):
		return DecodeMax(buf[1:])
	case byte(ssql.Function_COUNT):
		return DecodeCount(buf[1:])
	case byte(ssql.Function_PCTL):
		return DecodePercentile(buf[1:])
	case byte(ssql.Function_PART):
	}

	return nil
}
