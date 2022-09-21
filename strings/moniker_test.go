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

func TestBadMoniker(t *testing.T) {
	var _, ok = strings.MonikerFromString("bad")
	assert.False(t, ok)
}

func TestEmptyMoniker(t *testing.T) {
	var _, ok = strings.MonikerFromString("/")
	assert.False(t, ok)
}

func TestMoniker(t *testing.T) {
	var v1, ok = strings.MonikerFromString("/bali/abstractions/String/v1.2.3")
	assert.True(t, ok)
	assert.Equal(t, "/bali/abstractions/String/v1.2.3", v1.AsString())
	assert.False(t, v1.IsEmpty())
	assert.Equal(t, 4, v1.GetSize())
	assert.Equal(t, "bali", v1.GetItem(1))
	assert.Equal(t, "v1.2.3", v1.GetItem(-1))
	var v2 = strings.MonikerFromNames(v1.AsArray())
	assert.Equal(t, v1.String(), v2.String())
	var v3 = strings.MonikerFromNames(v1.GetItems(1, 2))
	assert.Equal(t, "/bali/abstractions", string(v3))
	assert.Equal(t, 3, v1.GetIndex("String"))
}

func TestMonikersLibrary(t *testing.T) {
	var v1, _ = strings.MonikerFromString("/bali/abstractions")
	var v2, _ = strings.MonikerFromString("/String/v1.2.3")
	assert.Equal(t, "/bali/abstractions/String/v1.2.3", strings.Monikers.Concatenate(v1, v2).AsString())
}
