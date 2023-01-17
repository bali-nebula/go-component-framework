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
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestEmptyBinary(t *tes.T) {
	var v = str.Binary("")
	ass.Equal(t, "", string(v))
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestBinary(t *tes.T) {
	var v = str.Binary("abcd1234")
	ass.Equal(t, "abcd1234", string(v))
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, v.GetSize())
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), str.BinaryFromArray(v.AsArray()).AsArray())
	ass.Equal(t, "abcd", string(str.BinaryFromSequence(v.GetValues(1, 3))))
}

func TestBinariesLibrary(t *tes.T) {
	var v1 = str.Binary("abcd")
	var v2 = str.Binary("12345678")
	ass.Equal(t, "abcd12345678", string(str.Binaries.Concatenate(v1, v2)))

	v1 = str.BinaryFromArray([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = str.BinaryFromArray([]byte{0x03, 0x00, 0x01, 0x02})
	var not = str.BinaryFromArray([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = str.BinaryFromArray([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = str.BinaryFromArray([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = str.BinaryFromArray([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = str.BinaryFromArray([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = str.BinaryFromArray([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	ass.Equal(t, not, str.Binaries.Not(v1))
	ass.Equal(t, and, str.Binaries.And(v1, v2))
	ass.Equal(t, sans, str.Binaries.Sans(v1, v2))
	ass.Equal(t, or, str.Binaries.Or(v1, v2))
	ass.Equal(t, xor, str.Binaries.Xor(v1, v2))
	ass.Equal(t, sans2, str.Binaries.Sans(v2, v1))
}
