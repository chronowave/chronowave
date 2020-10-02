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

package ssd

import (
	"bytes"
	"encoding/gob"
	"unicode/utf8"
)

type Node struct {
	Name     []byte
	Code     []byte
	Children []*Node
}

type EntityMeta struct {
	root  Node
	codes map[string]int32
}

func NewEntityMeta() *EntityMeta {
	return &EntityMeta{root: Node{}, codes: map[string]int32{}}
}

func EntityMetaFromBytes(d []byte) (*EntityMeta, error) {
	dec := gob.NewDecoder(bytes.NewReader(d))
	var root Node
	err := dec.Decode(&root)
	return &EntityMeta{root: root}, err
}

func (s *EntityMeta) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(s.root)
	return buf.Bytes(), err
}

func (s *EntityMeta) GenerateCode(path [][]byte) []byte {
	var code []byte
	parent := &s.root
	for i := 0; i < len(path); i++ {
		current := matchPath(path[i], parent.Children)
		if current != nil {
			parent = current
			code = append(code, current.Code...)
			continue
		}

		// new node
		for j := i; j < len(path); j++ {
			key := string(path[j])
			c, ok := s.codes[key]
			if !ok {
				// code starts from printable char: space, ascii 32
				c = int32(len(s.codes)) + 32
				s.codes[key] = c
			}

			// convert to utf-8 char
			tmp := make([]byte, 4)
			cnt := utf8.EncodeRune(tmp, c)
			parent.Children = append(parent.Children, &Node{
				Name:     append(path[j][:0:0], path[j]...),
				Code:     tmp[:cnt],
				Children: nil,
			})
			parent = parent.Children[len(parent.Children)-1]
			code = append(code, parent.Code...)
		}

		break
	}

	return code
}

// GetCode given entity path, returns UTF-8 code represents the path, true if found, otherwise false
func (s *EntityMeta) GetCode(path [][]byte) ([]byte, bool) {
	var code []byte
	parent := &s.root
	for _, p := range path {
		if len(p) == 0 {
			// skip empty
			continue
		}
		current := matchPath(p, parent.Children)
		if current == nil {
			return nil, false
		}
		parent = current
		code = append(code, current.Code...)
	}

	return code, true
}

// GetPath given UTF-8 code, returns entity path
func (s *EntityMeta) GetPath(code []byte) [][]byte {
	var path [][]byte

	parent := &s.root
	for len(code) > 0 {
		_, sz := utf8.DecodeRune(code)
		current := matchCode(code[:sz], parent.Children)
		if current == nil {
			return nil
		}
		path = append(path, current.Name)
		code, parent = code[sz:], current
	}
	return path
}

func matchPath(path []byte, nodes []*Node) *Node {
	for _, n := range nodes {
		if bytes.Compare(n.Name, path) == 0 {
			return n
		}
	}

	return nil
}

func matchCode(code []byte, nodes []*Node) *Node {
	for _, n := range nodes {
		if bytes.Compare(n.Code, code) == 0 {
			return n
		}
	}

	return nil
}
