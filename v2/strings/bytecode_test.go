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

func TestEmptyBytecode(t *tes.T) {
	var bytecode = `''`
	var v = str.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestBytecode(t *tes.T) {
	var bytecode = `'abcd1234'`
	var v = str.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, v.GetSize())
	ass.Equal(t, v.AsArray(), str.BytecodeFromArray(v.AsArray()).AsArray())
}
