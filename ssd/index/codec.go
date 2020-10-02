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
	"encoding/binary"
	"math"
	"unsafe"

	hfmi "github.com/rleiwang/hfmi/ctor"

	"github.com/chronowave/chronowave/ssd"
)

func EncodeIndexBlock(index *ssd.IndexedBlock) ([]byte, error) {
	meta, err := index.Meta.Bytes()
	if err != nil {
		return nil, err
	}

	// entity
	entity, entdict, entcnt := index.Entity.Bytes(), index.Entity.Dictionary(), index.Entity.Len()

	// content
	var (
		content  []byte
		contdict []byte
		contcnt  uint
	)
	if index.Content != nil {
		content, contdict, contcnt = index.Content.Bytes(), index.Content.Dictionary(), index.Content.Len()
	}

	sz := 1 + 8
	sz += len(index.EntityID)*4 + 4
	sz += len(meta) + 4
	sz += len(entity) + len(entdict) + 4*3
	csz, cols := sizingColumnar(index)
	sz += csz + 4*len(cols)
	sz += len(content) + len(contdict) + 4*3
	hsz, hlt := sizingHLT(index)
	sz += hsz + 4*len(hlt)
	sz += len(index.HeaderDA)*4 + 4
	sz += len(index.FragDA)*4 + 4

	buf := make([]byte, sz)
	// version
	buf[0] = 0
	working := buf[1:]
	binary.LittleEndian.PutUint64(working[:8], index.ID)
	working = working[8:]

	binary.LittleEndian.PutUint32(working[:4], uint32(len(index.EntityID)))
	working = working[4:]
	for _, v := range index.EntityID {
		binary.LittleEndian.PutUint32(working[:4], v)
		working = working[4:]
	}

	binary.LittleEndian.PutUint32(working[:4], uint32(len(meta)))
	working = working[4:]
	copy(working, meta)
	working = working[len(meta):]

	binary.LittleEndian.PutUint32(working[:4], uint32(entcnt))
	working = working[4:]
	binary.LittleEndian.PutUint32(working[:4], uint32(len(entity)))
	binary.LittleEndian.PutUint32(working[4:8], uint32(len(entdict)))
	working = working[8:]
	copy(working, entity)
	working = working[len(entity):]
	copy(working, entdict)
	working = working[len(entdict):]

	for _, c := range cols {
		binary.LittleEndian.PutUint32(working[:4], c)
		working = working[4:]
	}

	working = encodeColumnar(index, working)

	binary.LittleEndian.PutUint32(working[:4], uint32(contcnt))
	working = working[4:]
	binary.LittleEndian.PutUint32(working[:4], uint32(len(content)))
	binary.LittleEndian.PutUint32(working[4:8], uint32(len(contdict)))
	working = working[8:]
	copy(working, content)
	working = working[len(content):]
	copy(working, contdict)
	working = working[len(contdict):]

	for _, c := range hlt {
		binary.LittleEndian.PutUint32(working[:4], c)
		working = working[4:]
	}

	working = encodeHLT(index, working)

	binary.LittleEndian.PutUint32(working[:4], uint32(len(index.HeaderDA)))
	working = working[4:]
	for _, v := range index.HeaderDA {
		binary.LittleEndian.PutUint32(working[:4], v)
		working = working[4:]
	}

	binary.LittleEndian.PutUint32(working[:4], uint32(len(index.FragDA)))
	working = working[4:]
	for _, v := range index.FragDA {
		binary.LittleEndian.PutUint32(working[:4], v)
		working = working[4:]
	}

	return buf, nil
}

func DecodeIndexBlock(buf []byte) (*ssd.IndexedBlock, error) {
	// version buf[0]
	buf = buf[1:]

	index := &ssd.IndexedBlock{
		ID: binary.LittleEndian.Uint64(buf[:8]),
	}
	buf = buf[8:]

	sz := binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	index.EntityID = *(*[]uint32)(unsafe.Pointer(&buf))
	index.EntityID = index.EntityID[:sz]
	buf = buf[sz*4:]

	sz = binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]

	var err error
	if index.Meta, err = ssd.EntityMetaFromBytes(buf[:sz]); err != nil {
		return nil, err
	}
	buf = buf[sz:]

	cnt := binary.LittleEndian.Uint32(buf[:4])
	sz = binary.LittleEndian.Uint32(buf[4:8])
	dict := binary.LittleEndian.Uint32(buf[8:12])
	buf = buf[12:]
	index.Entity = hfmi.Build(uint(cnt), buf[sz:sz+dict], buf[:sz])
	buf = buf[sz+dict:]

	// columnar
	cols := *(*[]uint32)(unsafe.Pointer(&buf))
	cols = cols[:4]
	buf = buf[16:]

	buf = decodeColumnar(index, buf, cols)

	cnt = binary.LittleEndian.Uint32(buf[:4])
	sz = binary.LittleEndian.Uint32(buf[4:8])
	dict = binary.LittleEndian.Uint32(buf[8:12])
	buf = buf[12:]
	index.Content = hfmi.Build(uint(cnt), buf[sz:sz+dict], buf[:sz])
	buf = buf[sz+dict:]

	// HLT
	hlt := *(*[]uint32)(unsafe.Pointer(&buf))
	hlt = hlt[:5]
	buf = buf[20:]
	buf = decodeHLT(index, buf, hlt)

	sz = binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	index.HeaderDA = *(*[]uint32)(unsafe.Pointer(&buf))
	index.HeaderDA = index.HeaderDA[:sz]
	buf = buf[sz*4:]

	sz = binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	index.FragDA = *(*[]uint32)(unsafe.Pointer(&buf))
	index.FragDA = index.FragDA[:sz]

	return index, nil
}

func sizingHLT(index *ssd.IndexedBlock) (int, []uint32) {
	sz := make([]uint32, 5)
	sz[0] = uint32(len(index.HLT.Text.Entity))
	sz[1] = uint32(len(index.HLT.Float64.Entity))
	sz[2] = uint32(len(index.HLT.Int64.Entity))
	sz[3] = uint32(len(index.HLT.Bool.Entity))
	sz[4] = uint32(len(index.HLT.Null.Entity))

	bytes := 0
	for _, s := range sz {
		bytes += int(s) * 4
	}

	return bytes, sz
}

func encodeHLT(index *ssd.IndexedBlock, buf []byte) []byte {
	for _, v := range index.HLT.Text.Entity {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}
	for _, v := range index.HLT.Text.Attribute {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Float64.Entity {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}
	for _, v := range index.HLT.Float64.Attribute {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Int64.Entity {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Int64.Attribute {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Bool.Entity {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Bool.Attribute {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Null.Entity {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	for _, v := range index.HLT.Null.Attribute {
		binary.LittleEndian.PutUint16(buf[:2], v)
		buf = buf[2:]
	}

	return buf
}

func decodeHLT(index *ssd.IndexedBlock, buf []byte, sz []uint32) []byte {
	i := sz[0]
	index.HLT.Text.Entity = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Text.Entity = index.HLT.Text.Entity[:i]
	buf = buf[i*2:]

	index.HLT.Text.Attribute = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Text.Attribute = index.HLT.Text.Attribute[:i]
	buf = buf[i*2:]

	i = sz[1]
	index.HLT.Float64.Entity = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Float64.Entity = index.HLT.Float64.Entity[:i]
	buf = buf[i*2:]

	index.HLT.Float64.Attribute = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Float64.Attribute = index.HLT.Float64.Attribute[:i]
	buf = buf[i*2:]

	i = sz[2]
	index.HLT.Int64.Entity = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Int64.Entity = index.HLT.Int64.Entity[:i]
	buf = buf[i*2:]

	index.HLT.Int64.Attribute = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Int64.Attribute = index.HLT.Int64.Attribute[:i]
	buf = buf[i*2:]

	i = sz[3]
	index.HLT.Bool.Entity = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Bool.Entity = index.HLT.Bool.Entity[:i]
	buf = buf[i*2:]

	index.HLT.Bool.Attribute = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Bool.Attribute = index.HLT.Bool.Attribute[:i]
	buf = buf[i*2:]

	i = sz[4]
	index.HLT.Null.Entity = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Null.Entity = index.HLT.Null.Entity[:i]
	buf = buf[i*2:]

	index.HLT.Null.Attribute = *(*[]uint16)(unsafe.Pointer(&buf))
	index.HLT.Null.Attribute = index.HLT.Null.Attribute[:i]

	return buf[i*2:]
}

func sizingColumnar(index *ssd.IndexedBlock) (int, []uint32) {
	sz := make([]uint32, 4)
	sz[0] = uint32(len(index.Columnar.Float64)) * 8
	sz[1] = uint32(len(index.Columnar.Int64)) * 8
	sz[2] = uint32(len(index.Columnar.Bool))
	sz[3] = uint32(len(index.Columnar.Text)) * 4

	bytes := 0
	for _, s := range sz {
		bytes += int(s)
	}

	return bytes, sz
}

func encodeColumnar(index *ssd.IndexedBlock, buf []byte) []byte {
	for _, f := range index.Columnar.Float64 {
		binary.LittleEndian.PutUint64(buf[:8], math.Float64bits(f))
		buf = buf[8:]
	}

	for _, v := range index.Columnar.Int64 {
		binary.LittleEndian.PutUint64(buf[:8], uint64(v))
		buf = buf[8:]
	}

	for i, v := range index.Columnar.Bool {
		if v {
			buf[i] = 1
		}
	}
	buf = buf[len(index.Columnar.Bool):]

	for _, v := range index.Columnar.Text {
		binary.LittleEndian.PutUint32(buf[:4], v)
		buf = buf[4:]
	}

	return buf
}

func decodeColumnar(index *ssd.IndexedBlock, buf []byte, sz []uint32) []byte {
	index.Columnar.Float64 = *(*[]float64)(unsafe.Pointer(&buf))
	index.Columnar.Float64 = index.Columnar.Float64[:sz[0]/8]
	buf = buf[sz[0]:]

	index.Columnar.Int64 = *(*[]int64)(unsafe.Pointer(&buf))
	index.Columnar.Int64 = index.Columnar.Int64[:sz[1]/8]
	buf = buf[sz[1]:]

	index.Columnar.Bool = *(*[]bool)(unsafe.Pointer(&buf))
	index.Columnar.Bool = index.Columnar.Bool[:sz[2]]
	buf = buf[sz[2]:]

	index.Columnar.Text = *(*[]uint32)(unsafe.Pointer(&buf))
	index.Columnar.Text = index.Columnar.Text[:sz[3]/4]

	return buf[sz[3]:]
}
