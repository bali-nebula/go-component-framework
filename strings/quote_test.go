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

func TestBadQuote(t *testing.T) {
	var _, ok = strings.QuoteFromString(`"`)
	assert.False(t, ok)
}

func TestEmptyQuote(t *testing.T) {
	var v, ok = strings.QuoteFromString(`""`)
	assert.True(t, ok)
	assert.Equal(t, `""`, v.AsString())
	assert.True(t, v.IsEmpty())
	assert.Equal(t, 0, v.GetSize())
}

func TestQuote(t *testing.T) {
	var v, ok = strings.QuoteFromString(`"abcd本1234"`)
	assert.True(t, ok)
	assert.Equal(t, `"abcd本1234"`, v.AsString())
	assert.False(t, v.IsEmpty())
	assert.Equal(t, 9, v.GetSize())
	assert.Equal(t, 'a', v.GetItem(1))
	assert.Equal(t, '4', v.GetItem(-1))
	assert.Equal(t, v.String(), strings.QuoteFromRunes(v.AsArray()).AsString())
	assert.Equal(t, `"d本1"`, strings.QuoteFromRunes(v.GetItems(4, 6)).AsString())
	assert.Equal(t, 5, v.GetIndex('本'))
}

func TestQuotesLibrary(t *testing.T) {
	var v1, _ = strings.QuoteFromString(`"abcd本"`)
	var v2, _ = strings.QuoteFromString(`"1234"`)
	assert.Equal(t, `"abcd本1234"`, strings.Quotes.Concatenate(v1, v2).AsString())
}
