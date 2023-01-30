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
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestEmptyBytecode(t *tes.T) {
	var v = str.BytecodeFromString("")
	ass.Equal(t, "", v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestBytecode(t *tes.T) {
	var v = str.BytecodeFromString("abcd1234")
	ass.Equal(t, "abcd1234", v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, v.GetSize())
	ass.Equal(t, v.AsArray(), str.BytecodeFromArray(v.AsArray()).AsArray())
}
