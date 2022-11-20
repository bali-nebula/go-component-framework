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

func TestEmptyQuote(t *tes.T) {
	var v = str.Quote("")
	ass.Equal(t, "", string(v))
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestQuote(t *tes.T) {
	var v = str.Quote("abcd本1234")
	ass.Equal(t, "abcd本1234", string(v))
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, v.GetSize())
	ass.Equal(t, 'a', v.GetValue(1))
	ass.Equal(t, '4', v.GetValue(-1))
	ass.Equal(t, v.AsArray(), str.QuoteFromArray(v.AsArray()).AsArray())
	ass.Equal(t, "d本1", string(str.QuoteFromSequence(v.GetValues(4, 6))))
	ass.Equal(t, 5, v.GetIndex('本'))
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = str.Quote("abcd本")
	var v2 = str.Quote("1234")
	ass.Equal(t, "abcd本1234", string(str.Quotes.Concatenate(v1, v2)))
}
