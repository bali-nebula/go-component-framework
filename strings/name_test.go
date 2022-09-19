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

func TestBadName(t *testing.T) {
	var _, ok = strings.NameFromString("bad")
	assert.False(t, ok)
}

func TestEmptyName(t *testing.T) {
	var _, ok = strings.NameFromString("/")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	var v1, ok = strings.NameFromString("/bali/abstractions/String/v1.2.3")
	assert.True(t, ok)
	assert.Equal(t, "/bali/abstractions/String/v1.2.3", v1.AsString())
	assert.False(t, v1.IsEmpty())
	assert.Equal(t, 4, v1.GetSize())
	assert.Equal(t, "bali", v1.GetItem(1))
	assert.Equal(t, "v1.2.3", v1.GetItem(-1))
	var v2 = strings.NameFromLabels(v1.AsArray())
	assert.Equal(t, v1.String(), v2.String())
	var v3 = strings.NameFromLabels(v1.GetItems(1, 2))
	assert.Equal(t, "/bali/abstractions", string(v3))
	assert.Equal(t, 3, v1.GetIndex("String"))
}

func TestNamesLibrary(t *testing.T) {
	var v1, _ = strings.NameFromString("/bali/abstractions")
	var v2, _ = strings.NameFromString("/String/v1.2.3")
	assert.Equal(t, "/bali/abstractions/String/v1.2.3", strings.Names.Concatenate(v1, v2).AsString())
}
