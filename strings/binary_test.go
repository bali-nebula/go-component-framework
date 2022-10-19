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
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBadBinary(t *tes.T) {
	var _, ok = str.BinaryFromString("'.'")
	ass.False(t, ok)
}

func TestEmptyBinary(t *tes.T) {
	var v, ok = str.BinaryFromString("''")
	ass.True(t, ok)
	ass.Equal(t, "''", v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
	ass.Equal(t, 0, v.GetIndex(byte(0x69)))
}

func TestBinary(t *tes.T) {
	var v, ok = str.BinaryFromString("'abcd\n1234'")
	ass.True(t, ok)
	ass.Equal(t, "'abcd1234'", v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, v.GetSize())
	ass.Equal(t, byte(0x69), v.GetItem(1))
	ass.Equal(t, byte(0xf8), v.GetItem(-1))
	ass.Equal(t, v.String(), str.BinaryFromBytes(v.AsArray()).AsString())
	ass.Equal(t, "'abcd'", str.BinaryFromBytes(v.GetItems(1, 3)).AsString())
	ass.Equal(t, 1, v.GetIndex(byte(0x69)))
}

func TestBinariesLibrary(t *tes.T) {
	var v1, _ = str.BinaryFromString("'abcd'")
	var v2, _ = str.BinaryFromString("'12345678'")
	ass.Equal(t, "'abcd12345678'", str.Binaries.Concatenate(v1, v2).AsString())

	v1 = str.BinaryFromBytes([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = str.BinaryFromBytes([]byte{0x03, 0x00, 0x01, 0x02})
	var not = str.BinaryFromBytes([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = str.BinaryFromBytes([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = str.BinaryFromBytes([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = str.BinaryFromBytes([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = str.BinaryFromBytes([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = str.BinaryFromBytes([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	ass.Equal(t, not, str.Binaries.Not(v1))
	ass.Equal(t, and, str.Binaries.And(v1, v2))
	ass.Equal(t, sans, str.Binaries.Sans(v1, v2))
	ass.Equal(t, or, str.Binaries.Or(v1, v2))
	ass.Equal(t, xor, str.Binaries.Xor(v1, v2))
	ass.Equal(t, sans2, str.Binaries.Sans(v2, v1))
}
