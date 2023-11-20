/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package ranges_test

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	ran "github.com/bali-nebula/go-component-framework/v2/ranges"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSpectrWithNames(t *tes.T) {
	var first = str.NameFromString("/bali/types/abstractions")
	var last = str.NameFromString("/bali/types/strings")
	var s = ran.Spectrum(first, abs.RIGHT, last)
	ass.False(t, s.ContainsValue(str.NameFromString("/bali")))
	ass.False(t, s.ContainsValue(first))
	ass.True(t, s.ContainsValue(str.NameFromString("/bali/types/abstractions/Sequential/v1")))
	ass.True(t, s.ContainsValue(last))
	ass.False(t, s.ContainsValue(str.NameFromString("/bali/types/strings/Version/v1")))
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}

func TestSpectrWithQuotes(t *tes.T) {
	var first = str.QuoteFromString("A")
	var last = str.QuoteFromString("Fe")
	var s = ran.Spectrum(first, abs.EXCLUSIVE, last)
	ass.False(t, s.ContainsValue(str.QuoteFromString("1")))
	ass.False(t, s.ContainsValue(first))
	ass.True(t, s.ContainsValue(str.QuoteFromString("Boo")))
	ass.False(t, s.ContainsValue(last))
	ass.False(t, s.ContainsValue(str.QuoteFromString("Google")))
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}

func TestSpectrWithVersions(t *tes.T) {
	var first = str.VersionFromString("1.2")
	var last = str.VersionFromString("1.6.1")
	var s = ran.Spectrum(first, abs.LEFT, last)
	ass.False(t, s.ContainsValue(str.VersionFromString("1")))
	ass.True(t, s.ContainsValue(first))
	ass.True(t, s.ContainsValue(str.VersionFromString("1.5.3")))
	ass.False(t, s.ContainsValue(last))
	ass.False(t, s.ContainsValue(str.VersionFromString("1.6.2")))
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}
