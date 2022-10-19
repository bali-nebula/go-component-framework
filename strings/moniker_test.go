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
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBadMoniker(t *tes.T) {
	var _, ok = str.MonikerFromString("bad")
	ass.False(t, ok)
}

func TestEmptyMoniker(t *tes.T) {
	var _, ok = str.MonikerFromString("/")
	ass.False(t, ok)
}

func TestMoniker(t *tes.T) {
	var v1, ok = str.MonikerFromString("/bali/abstractions/String/v1.2.3")
	ass.True(t, ok)
	ass.Equal(t, "/bali/abstractions/String/v1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, v1.GetSize())
	ass.Equal(t, "bali", v1.GetItem(1))
	ass.Equal(t, "v1.2.3", v1.GetItem(-1))
	var v2 = str.MonikerFromNames(v1.AsArray())
	ass.Equal(t, v1.String(), v2.String())
	var v3 = str.MonikerFromNames(v1.GetItems(1, 2))
	ass.Equal(t, "/bali/abstractions", string(v3))
	ass.Equal(t, 3, v1.GetIndex("String"))
}

func TestMonikersLibrary(t *tes.T) {
	var v1, _ = str.MonikerFromString("/bali/abstractions")
	var v2, _ = str.MonikerFromString("/String/v1.2.3")
	ass.Equal(t, "/bali/abstractions/String/v1.2.3", str.Monikers.Concatenate(v1, v2).AsString())
}
