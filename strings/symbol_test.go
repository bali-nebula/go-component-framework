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

func TestBadSymbol(t *testing.T) {
	var _, ok = strings.SymbolFromString(`bad`)
	assert.False(t, ok)
}

func TestEmptySymbol(t *testing.T) {
	var _, ok = strings.SymbolFromString(`$`)
	assert.False(t, ok)
}

func TestSymbol(t *testing.T) {
	var v, ok = strings.SymbolFromString("$foobar")
	assert.True(t, ok)
	assert.Equal(t, "$foobar", v.AsString())
	assert.False(t, v.IsEmpty())
	assert.Equal(t, 6, v.GetSize())
	assert.Equal(t, 'f', v.GetItem(1))
	assert.Equal(t, 'r', v.GetItem(-1))
	assert.Equal(t, v.String(), strings.SymbolFromRunes(v.AsArray()).AsString())
	assert.Equal(t, "$foo", strings.SymbolFromRunes(v.GetItems(1, 3)).AsString())
	assert.Equal(t, 4, v.GetIndex('b'))
}

func TestSymbolsLibrary(t *testing.T) {
	var v1, _ = strings.SymbolFromString("$foo")
	var v2, _ = strings.SymbolFromString("$bar")
	assert.Equal(t, "$foobar", strings.Symbols.Concatenate(v1, v2).AsString())
}
