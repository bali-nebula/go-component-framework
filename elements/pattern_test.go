/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoPattern(t *testing.T) {
	var v, ok = elements.PatternFromString(`""?`)
	assert.False(t, ok)
	assert.Equal(t, ``, string(v))
}

func TestNonePattern(t *testing.T) {
	var v, ok = elements.PatternFromString(`none`)
	assert.True(t, ok)
	assert.Equal(t, `none`, string(v))

	var text = ""
	assert.False(t, v.MatchesText(text))
	assert.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	assert.False(t, v.MatchesText(text))
	assert.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *testing.T) {
	var v, ok = elements.PatternFromString(`any`)
	assert.True(t, ok)
	assert.Equal(t, `any`, string(v))

	var text = ""
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *testing.T) {
	var v, ok = elements.PatternFromString(`"c(.+t)"?`)
	assert.True(t, ok)
	assert.Equal(t, `"c(.+t)"?`, string(v))

	var text = "ct"
	assert.False(t, v.MatchesText(text))
	assert.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	assert.True(t, v.MatchesText(text))
	assert.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}

func TestBadPattern(t *testing.T) {
	var v, ok = elements.PatternFromString(`?`)
	assert.False(t, ok)
	assert.Equal(t, ``, string(v))
}
