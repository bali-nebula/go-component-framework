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

func TestBadQuote(t *tes.T) {
	var _, ok = str.QuoteFromString(`"`)
	ass.False(t, ok)
}

func TestEmptyQuote(t *tes.T) {
	var v, ok = str.QuoteFromString(`""`)
	ass.True(t, ok)
	ass.Equal(t, `""`, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestQuote(t *tes.T) {
	var v, ok = str.QuoteFromString(`"abcd本1234"`)
	ass.True(t, ok)
	ass.Equal(t, `"abcd本1234"`, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, v.GetSize())
	ass.Equal(t, 'a', v.GetItem(1))
	ass.Equal(t, '4', v.GetItem(-1))
	ass.Equal(t, v.String(), str.QuoteFromRunes(v.AsArray()).AsString())
	ass.Equal(t, `"d本1"`, str.QuoteFromRunes(v.GetItems(4, 6)).AsString())
	ass.Equal(t, 5, v.GetIndex('本'))
}

func TestQuotesLibrary(t *tes.T) {
	var v1, _ = str.QuoteFromString(`"abcd本"`)
	var v2, _ = str.QuoteFromString(`"1234"`)
	ass.Equal(t, `"abcd本1234"`, str.Quotes.Concatenate(v1, v2).AsString())
}
