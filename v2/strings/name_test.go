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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestName(t *tes.T) {
	var v1 = str.NameFromString("/bali/types/abstractions/String")
	ass.Equal(t, "/bali/types/abstractions/String", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, v1.GetSize())
	ass.Equal(t, abs.Identifier("bali"), v1.GetValue(1))
	ass.Equal(t, abs.Identifier("String"), v1.GetValue(-1))
	var v2 = str.NameFromArray(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = str.NameFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "/bali/types", v3.AsString())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = str.NameFromString("/bali/types/abstractions")
	var v2 = str.NameFromString("/String")
	ass.Equal(t, "/bali/types/abstractions/String", str.Names.Concatenate(v1, v2).AsString())
}
