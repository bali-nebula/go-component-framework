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

func TestSymbol(t *tes.T) {
	var foobar = "foobar"
	var v = str.SymbolFromString(foobar)
	ass.Equal(t, foobar, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, v.GetSize())
	ass.Equal(t, []rune(foobar), v.AsArray())
}
