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

package internal

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/chronowave/chronowave/ssd"
	"github.com/chronowave/chronowave/ssd/aggregator"
)

func (i64 *ColumnInt64) Bytes() []byte {
	// number of rows (4) + len(rows) * (index of rows (4) + int64 value (8))
	buf := make([]byte, uint(4+len(i64.Rows)*12))
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(i64.Rows)))

	j, n := 4, 0
	for _, r := range i64.Rows {
		n = j + 4
		binary.LittleEndian.PutUint32(buf[j:n], r)
		j = n
	}

	for _, c := range i64.Cols {
		n = j + 8
		binary.LittleEndian.PutUint64(buf[j:n], uint64(i64.Columnar[c]))
		j = n
	}

	return buf
}

func (i64 *ColumnInt64) FromBytes(buf []byte) {
	// number of rows (4) + len(rows) * (index of rows (4) + int64 value (8))
	sz := binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	i64.Rows = *(*[]uint32)(unsafe.Pointer(&buf))
	i64.Rows = i64.Rows[:sz]
	buf = buf[sz*4:]
	i64.Columnar = *(*[]int64)(unsafe.Pointer(&buf))
	i64.Columnar = i64.Columnar[:sz]
}

func (txt *ColumnText) Bytes() []byte {
	// number of rows (4) + number of texts (4) + len(rows) * (index of rows (4) + uint32 value (4))
	sz := uint(8 + len(txt.Rows)*8)
	for _, v := range txt.Columnar {
		// key (4) + buf size (4) + len(buf)
		sz += uint(len(v)) + 8
	}
	// number of text entries
	sz += 4

	buf := make([]byte, sz)
	ptr := buf
	// # of rows (4)
	binary.LittleEndian.PutUint32(ptr[:4], uint32(len(txt.Rows)))
	ptr = ptr[4:]

	for _, r := range txt.Rows {
		binary.LittleEndian.PutUint32(ptr[:4], r)
		ptr = ptr[4:]
	}

	for _, c := range txt.Cols {
		binary.LittleEndian.PutUint32(ptr[:4], c)
		ptr = ptr[4:]
	}

	binary.LittleEndian.PutUint32(ptr[:4], uint32(len(txt.Columnar)))
	ptr = ptr[4:]
	text := ptr[len(txt.Columnar)*8:]
	for k, v := range txt.Columnar {
		binary.LittleEndian.PutUint32(ptr[:4], k)
		binary.LittleEndian.PutUint32(ptr[4:8], uint32(len(v)))
		copy(text, v)
		ptr, text = ptr[8:], text[len(v):]
	}

	return buf
}

func (txt *ColumnText) FromBytes(buf []byte) {
	// # of rows (4)
	sz := binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	txt.Rows = *(*[]uint32)(unsafe.Pointer(&buf))
	txt.Rows = txt.Rows[:sz]
	buf = buf[sz*4:]
	txt.Cols = *(*[]uint32)(unsafe.Pointer(&buf))
	txt.Cols = txt.Cols[:sz]
	buf = buf[sz*4:]

	// # of text entries
	sz, buf = binary.LittleEndian.Uint32(buf[:4]), buf[4:]
	text := buf[sz*8:]

	txt.Columnar = make(map[uint32][]byte, sz)
	for i := uint32(0); i < sz; i++ {
		key, bsz := binary.LittleEndian.Uint32(buf[:4]), binary.LittleEndian.Uint32(buf[4:8])
		txt.Columnar[key] = text[:bsz]
		buf, text = buf[8:], text[bsz:]
	}
}

func SerializeResultSet(rs *ssd.ResultSet) []byte {
	w := bytes.NewBuffer(make([]byte, 0, 2048))

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(rs.RowId)))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(rs.ColumnType)))
	w.Write(buf)

	binary.LittleEndian.PutUint32(buf[:4], uint32(len(rs.Text)))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(rs.Aggregate)))
	w.Write(buf)

	for _, rid := range rs.RowId {
		binary.LittleEndian.PutUint64(buf, rid)
		w.Write(buf)
	}

	w.Write(rs.ColumnType)

	for _, c := range rs.Column {
		w.Write(c.RowIdx)
		for _, v := range c.Value {
			binary.LittleEndian.PutUint64(buf, v)
			w.Write(buf)
		}
	}

	if len(rs.Text) > 0 {
		var data []byte
		for _, v := range rs.Text {
			binary.LittleEndian.PutUint32(buf[:4], uint32(len(v)))
			w.Write(buf[:4])
			data = append(data, v...)
		}
		w.Write(data)
	}

	if len(rs.Aggregate) > 0 {
		var data []byte
		for _, a := range rs.Aggregate {
			tmp := a.Bytes()
			data = append(data, tmp...)
			binary.LittleEndian.PutUint32(buf[:4], uint32(len(tmp)))
			w.Write(buf[:4])
		}
		w.Write(data)
	}

	return w.Bytes()
}

func DeserializeResultSet(data []byte) *ssd.ResultSet {
	nor := binary.LittleEndian.Uint32(data[:4])
	noc := binary.LittleEndian.Uint32(data[4:8])
	lot := binary.LittleEndian.Uint32(data[8:12])
	loa := binary.LittleEndian.Uint32(data[12:16])
	data = data[16:]

	rs := &ssd.ResultSet{}

	if nor > 0 {
		rs.RowId = *(*[]uint64)(unsafe.Pointer(&data))
		rs.RowId, data = rs.RowId[:nor], data[nor*8:]
	}

	if noc > 0 {
		rs.ColumnType, data = data[:noc], data[noc:]
		rs.Column = make([]ssd.Column, noc)
		for i := range rs.Column {
			rs.Column[i].RowIdx, data = data[:nor], data[nor:]
			rs.Column[i].Value = *(*[]uint64)(unsafe.Pointer(&data))
			rs.Column[i].Value, data = rs.Column[i].Value[:nor], data[nor*8:]
		}
	}

	if lot > 0 {
		sz := *(*[]uint32)(unsafe.Pointer(&data))
		sz, data = sz[:lot], data[lot*4:]
		rs.Text = make([][]byte, lot)
		for i := range rs.Text {
			rs.Text[i], data = data[:sz[i]], data[sz[i]:]
		}
	}

	if loa > 0 {
		sz := *(*[]uint32)(unsafe.Pointer(&data))
		sz, data = sz[:loa], data[loa*4:]
		rs.Aggregate = make([]aggregator.Aggregator, loa)
		for i := range rs.Aggregate {
			rs.Aggregate[i], data = aggregator.FromBytes(data[:sz[i]]), data[sz[i]:]
		}
	}

	return rs
}
