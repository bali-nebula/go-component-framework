/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadBinary(t *testing.T) {
	var _, ok = strings.BinaryFromString("'.'")
	assert.False(t, ok)
}

func TestEmptyBinary(t *testing.T) {
	var v, ok = strings.BinaryFromString("''")
	assert.True(t, ok)
	assert.Equal(t, "''", v.AsString())
	assert.True(t, v.IsEmpty())
	assert.Equal(t, 0, v.GetSize())
	assert.Equal(t, 0, v.GetIndex(byte(0x69)))
}

func TestBinary(t *testing.T) {
	var v, ok = strings.BinaryFromString("'abcd\n1234'")
	assert.True(t, ok)
	assert.Equal(t, "'abcd1234'", v.AsString())
	assert.False(t, v.IsEmpty())
	assert.Equal(t, 6, v.GetSize())
	assert.Equal(t, byte(0x69), v.GetItem(1))
	assert.Equal(t, byte(0xf8), v.GetItem(-1))
	assert.Equal(t, v.String(), strings.BinaryFromBytes(v.AsArray()).AsString())
	assert.Equal(t, "'abcd'", strings.BinaryFromBytes(v.GetItems(1, 3)).AsString())
	assert.Equal(t, 1, v.GetIndex(byte(0x69)))
}

func TestBinariesLibrary(t *testing.T) {
	var v1, _ = strings.BinaryFromString("'abcd'")
	var v2, _ = strings.BinaryFromString("'12345678'")
	assert.Equal(t, "'abcd12345678'", strings.Binaries.Concatenate(v1, v2).AsString())

	v1 = strings.BinaryFromBytes([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = strings.BinaryFromBytes([]byte{0x03, 0x00, 0x01, 0x02})
	var not = strings.BinaryFromBytes([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = strings.BinaryFromBytes([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = strings.BinaryFromBytes([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = strings.BinaryFromBytes([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = strings.BinaryFromBytes([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = strings.BinaryFromBytes([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	assert.Equal(t, not, strings.Binaries.Not(v1))
	assert.Equal(t, and, strings.Binaries.And(v1, v2))
	assert.Equal(t, sans, strings.Binaries.Sans(v1, v2))
	assert.Equal(t, or, strings.Binaries.Or(v1, v2))
	assert.Equal(t, xor, strings.Binaries.Xor(v1, v2))
	assert.Equal(t, sans2, strings.Binaries.Sans(v2, v1))
}
