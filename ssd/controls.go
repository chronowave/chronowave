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

const (
	// SENTINEL smallest one and only in the alphabet
	SENTINEL byte = iota

	// SOH start of heading 1, field separator
	SOH

	// FRAG fragment long text, 2
	FRAG

	// SOA start of array, 3
	SOA

	// EOA end of array, 4
	EOA

	// AED array element divider, 5
	AED

	// EOO end of object, 6
	EOO

	// TEXT direct string start, 14
	TEXT = iota + 7

	// FLT32 float numeric start, 15
	FLT32

	// FLT64 double numeric start, 16
	FLT64

	// INT8 signed byte integer start, 17
	INT8

	// INT16 signed short integer start, 18
	INT16

	// INT32 unsigned int integer start, 19
	INT32

	// INT64 unsigned long integer start, 20
	INT64

	// BOOL boolean true start, 21
	BOOL

	// NULL NULL start, 22
	NULL

	// JSON JSON start, 23
	JSON
)

var (
	VALUES = []byte{TEXT, FLT32, FLT64, INT8, INT16, INT32, INT64, BOOL, NULL, JSON}
)

func IsControlCharacter(c byte) bool {
	return c <= EOO || (TEXT <= c && c <= JSON)
}
