/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBadSymbol(t *tes.T) {
	var _, ok = ele.SymbolFromString(`bad`)
	ass.False(t, ok)
}

func TestEmptySymbol(t *tes.T) {
	var _, ok = ele.SymbolFromString(`$`)
	ass.False(t, ok)
}

func TestSymbol(t *tes.T) {
	var v, ok = ele.SymbolFromString("$foobar")
	ass.True(t, ok)
	ass.Equal(t, "$foobar", lan.FormatValue(v))
}
