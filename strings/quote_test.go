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
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestEmptyQuote(t *tes.T) {
	var v = str.QuoteFromString("")
	ass.Equal(t, "", v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestQuote(t *tes.T) {
	var v = str.QuoteFromString("abcd本1234")
	ass.Equal(t, "abcd本1234", v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, v.GetSize())
	ass.Equal(t, rune('a'), v.GetValue(1))
	ass.Equal(t, rune('4'), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), str.QuoteFromArray(v.AsArray()).AsArray())
	ass.Equal(t, "d本1", str.QuoteFromSequence(v.GetValues(4, 6)).AsString())
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = str.QuoteFromString("abcd本")
	var v2 = str.QuoteFromString("1234")
	ass.Equal(t, "abcd本1234", str.Quotes.Concatenate(v1, v2).AsString())
}
