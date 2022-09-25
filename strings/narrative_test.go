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

func TestBadNarrative(t *testing.T) {
	var _, ok = strings.NarrativeFromString(bad)
	assert.False(t, ok)
}

func TestEmptyNarrative(t *testing.T) {
	var v, ok = strings.NarrativeFromString(n0)
	assert.True(t, ok)
	assert.Equal(t, n0, v.AsString())
	assert.True(t, v.IsEmpty())
	assert.Equal(t, 0, v.GetSize())
}

func TestNarrative(t *testing.T) {
	var v, ok = strings.NarrativeFromString(n1)
	assert.True(t, ok)
	assert.Equal(t, n1, v.AsString())
	assert.False(t, v.IsEmpty())
	assert.Equal(t, 10, v.GetSize())
	assert.Equal(t, 'a', v.GetItem(2))
	assert.Equal(t, '4', v.GetItem(-1))
	assert.Equal(t, v.String(), strings.NarrativeFromRunes(v.AsArray()).AsString())
	assert.Equal(t, n2, strings.NarrativeFromRunes(v.GetItems(5, 7)).AsString())
	assert.Equal(t, 6, v.GetIndex('本'))
}

func TestNarrativesLibrary(t *testing.T) {
	var v1, _ = strings.NarrativeFromString(n3)
	var v2, _ = strings.NarrativeFromString(n4)
	assert.Equal(t, n1, strings.Narratives.Concatenate(v1, v2).AsString())
}
