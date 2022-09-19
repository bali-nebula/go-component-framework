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
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadSymbol(t *testing.T) {
	var _, ok = elements.SymbolFromString(`bad`)
	assert.False(t, ok)
}

func TestEmptySymbol(t *testing.T) {
	var _, ok = elements.SymbolFromString(`$`)
	assert.False(t, ok)
}

func TestSymbol(t *testing.T) {
	var v, ok = elements.SymbolFromString("$foobar")
	assert.True(t, ok)
	assert.Equal(t, "$foobar", v.AsString())
}
