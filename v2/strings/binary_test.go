/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestEmptyBinary(t *tes.T) {
	var binary = `'>
<'`
	var v = str.BinaryFromString(binary)
	ass.Equal(t, binary, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestBinary(t *tes.T) {
	var b1 = `'>
    abcd1234
<'`
	var b2 = `'>
    abcd
<'`
	var v = str.BinaryFromString(b1)
	ass.Equal(t, b1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, v.GetSize())
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), str.BinaryFromArray(v.AsArray()).AsArray())
	ass.Equal(t, b2, str.BinaryFromSequence(v.GetValues(1, 3)).AsString())
}

func TestBinaryLibrary(t *tes.T) {
	var b1 = `'>
    abcd
<'`
	var b2 = `'>
    12345678
<'`
	var b3 = `'>
    abcd12345678
<'`
	var v1 = str.BinaryFromString(b1)
	var v2 = str.BinaryFromString(b2)
	ass.Equal(t, b3, str.Binary.Concatenate(v1, v2).AsString())

	v1 = str.BinaryFromArray([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = str.BinaryFromArray([]byte{0x03, 0x00, 0x01, 0x02})
	var not = str.BinaryFromArray([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = str.BinaryFromArray([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = str.BinaryFromArray([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = str.BinaryFromArray([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = str.BinaryFromArray([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = str.BinaryFromArray([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	ass.Equal(t, not, str.Binary.Not(v1))
	ass.Equal(t, and, str.Binary.And(v1, v2))
	ass.Equal(t, sans, str.Binary.Sans(v1, v2))
	ass.Equal(t, or, str.Binary.Or(v1, v2))
	ass.Equal(t, xor, str.Binary.Xor(v1, v2))
	ass.Equal(t, sans2, str.Binary.Sans(v2, v1))
}
