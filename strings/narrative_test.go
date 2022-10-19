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

const bad = `">
<"`

const n0 = `">

<"`

const n1 = `">
	abcd本1234
<"`

const n2 = `">
d本1
<"`

const n3 = `">
	abcd本
<"`

const n4 = `">
1234
<"`

func TestBadNarrative(t *tes.T) {
	var _, ok = str.NarrativeFromString(bad)
	ass.False(t, ok)
}

func TestEmptyNarrative(t *tes.T) {
	var v, ok = str.NarrativeFromString(n0)
	ass.True(t, ok)
	ass.Equal(t, n0, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestNarrative(t *tes.T) {
	var v, ok = str.NarrativeFromString(n1)
	ass.True(t, ok)
	ass.Equal(t, n1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 10, v.GetSize())
	ass.Equal(t, 'a', v.GetItem(2))
	ass.Equal(t, '4', v.GetItem(-1))
	ass.Equal(t, v.String(), str.NarrativeFromRunes(v.AsArray()).AsString())
	ass.Equal(t, n2, str.NarrativeFromRunes(v.GetItems(5, 7)).AsString())
	ass.Equal(t, 6, v.GetIndex('本'))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1, _ = str.NarrativeFromString(n3)
	var v2, _ = str.NarrativeFromString(n4)
	ass.Equal(t, n1, str.Narratives.Concatenate(v1, v2).AsString())
}
