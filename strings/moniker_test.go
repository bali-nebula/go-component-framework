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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestMoniker(t *tes.T) {
	var v1 = str.MonikerFromString("/bali/abstractions/String/v1.2.3")
	ass.Equal(t, "/bali/abstractions/String/v1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, v1.GetSize())
	ass.Equal(t, abs.Name("bali"), v1.GetValue(1))
	ass.Equal(t, abs.Name("v1.2.3"), v1.GetValue(-1))
	var v2 = str.MonikerFromArray(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = str.MonikerFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "/bali/abstractions", v3.AsString())
}

func TestMonikersLibrary(t *tes.T) {
	var v1 = str.MonikerFromString("/bali/abstractions")
	var v2 = str.MonikerFromString("/String/v1.2.3")
	ass.Equal(t, "/bali/abstractions/String/v1.2.3", str.Monikers.Concatenate(v1, v2).AsString())
}
