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
	"math/bits"
	"strconv"

	"github.com/buger/jsonparser"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/internal"
)

func ParseJson(json []byte) (*ssd.ParsedBlock, error) {
	f := &ssd.ParsedBlock{
		Meta:    ssd.NewEntityMeta(),
		Entity:  bytes.NewBuffer(make([]byte, 0, len(json))),
		Content: ssd.NewTextColumnar(),
	}

	for {
		// offset to json buffer, number of object,
		var exec func(*ssd.ParsedBlock, [][]byte, []byte) (int, error)
	open:
		for _, b := range json {
			switch b {
			case '{':
				exec = parseObject
				break open
			case '[':
				exec = parseArray
				break open
			}
		}

		if exec == nil {
			break
		}

		// end of entity, } or ], note: eoe would be at before closing } or ], including white space
		eoe, err := exec(f, nil, json)
		if err != nil {
			return f, err
		}

		json = json[eoe+1:]
	}

	return f, nil
}

// offset in json doc, number of objects, error if any
func parseObject(f *ssd.ParsedBlock, path [][]byte, data []byte) (int, error) {
	if path == nil {
		// root
		if f.Count > 0 {
			// end of previous object
			f.Entity.WriteByte(ssd.SOH)
			f.Entity.WriteByte(ssd.EOO)
		}

		f.Count++
	}

	// eoo -> end of object, offset in json doc
	eoo := 0
	handle := func(key []byte, value []byte, valueType jsonparser.ValueType, offset int) error {
		eoo = offset
		return parseObjectValue(f, valueType, append(path, replaceSlashWithUnderScore(key)), value)
	}

	err := jsonparser.ObjectEach(data, handle)
	if err != nil {
		return eoo, err
	}

	return eoo, nil
}

func parseArray(f *ssd.ParsedBlock, path [][]byte, data []byte) (int, error) {
	var objerr error
	eoa, err := jsonparser.ArrayEach(data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if dataType == jsonparser.Object {
				if _, err := parseObject(f, path, value); err != nil {
					objerr = err
				}
			}
		})

	if err == nil {
		err = objerr
	}
	return eoa, err
}

func parseArrayValue(f *ssd.ParsedBlock, path [][]byte, data []byte) (int, error) {
	// start of array element
	pathCode := f.Meta.GenerateCode(path)
	if f.Entity.Len() > 0 {
		f.Entity.WriteByte(ssd.SOH)
	}
	f.Entity.Write(pathCode)
	f.Entity.WriteByte(ssd.SOA)

	sep := false
	handle := func(value []byte, valueType jsonparser.ValueType, offset int, err error) {
		if sep {
			f.Entity.WriteByte(ssd.SOH)
			f.Entity.WriteByte(ssd.AED)
		}
		parseObjectValue(f, valueType, path, value)
		sep = true
	}

	eoa, err := jsonparser.ArrayEach(data, handle)

	// end of array
	f.Entity.WriteByte(ssd.SOH)
	f.Entity.Write(pathCode)
	f.Entity.WriteByte(ssd.EOA)
	return eoa, err
}

// https://tools.ietf.org/html/rfc7159#section-3
func parseObjectValue(f *ssd.ParsedBlock, valueType jsonparser.ValueType, path [][]byte, value []byte) error {
	pathCode := f.Meta.GenerateCode(path)
	writePath := func() {
		if f.Entity.Len() > 0 {
			f.Entity.WriteByte(ssd.SOH)
		}
		f.Entity.Write(pathCode)
	}

	var err error
	switch valueType {
	case jsonparser.String:
		// start of string value
		writePath()
		f.Entity.WriteByte(ssd.TEXT)
		if len(value) == 0 {
			// empty string
			f.Columnar.Text = append(f.Columnar.Text, 0)
		} else {
			f.Columnar.Text = append(f.Columnar.Text, f.Content.Add(string(fragmentLongText(value, internal.FragmentSize))))
		}
	case jsonparser.Number:
		// start of number value
		writePath()
		dot, rank := byte('.'), [8]uint{}
		for _, b := range value {
			rank[bits.OnesCount8(dot^b)]++
		}
		if rank[0] > 0 {
			if v, e := strconv.ParseFloat(string(value), 64); e == nil {
				// encode as float64
				f.Entity.WriteByte(ssd.FLT64)
				f.Columnar.Float64 = append(f.Columnar.Float64, v)
			} else {
				err = e
			}
		} else {
			if v, e := strconv.ParseInt(string(value), 10, 64); e == nil {
				// encode as int64
				f.Entity.WriteByte(ssd.INT64)
				f.Columnar.Int64 = append(f.Columnar.Int64, v)
			} else {
				err = e
			}
		}
	case jsonparser.Object:
		_, err = parseObject(f, path, value)
	case jsonparser.Array:
		_, err = parseArrayValue(f, path, value)
	case jsonparser.Boolean:
		writePath()
		f.Entity.WriteByte(ssd.BOOL)
		if value[0] == 't' || value[0] == 'T' {
			f.Columnar.Bool = append(f.Columnar.Bool, true)
		} else {
			f.Columnar.Bool = append(f.Columnar.Bool, false)
		}
	case jsonparser.Null:
		writePath()
		f.Entity.WriteByte(ssd.NULL)
	case jsonparser.Unknown:
		err = jsonparser.UnknownValueTypeError
	}

	return err
}

func replaceSlashWithUnderScore(key []byte) []byte {
	for i, b := range key {
		if b == '/' || b <= ssd.EOA {
			key[i] = '_'
		}
	}
	return key
}

func fragmentLongText(value []byte, fsz int) []byte {
	i, fragments := len(value)-1, (len(value)-1)/fsz
	if fragments > 0 {
		fragmented := make([]byte, len(value)+fragments)
		copy(fragmented, value)
		value = fragmented
		for j := len(value) - 1; fragments > 0; j-- {
			b := value[i]
			if ssd.IsControlCharacter(b) {
				value[j] = ' '
			} else {
				value[j] = b
			}

			if i%fsz == 0 {
				j--
				value[j] = ssd.FRAG
				fragments--
			}
			i--
		}
	}

	for ; i >= 0; i-- {
		if ssd.IsControlCharacter(value[i]) {
			value[i] = ' '
		}
	}

	return value
}
