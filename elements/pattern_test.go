/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/bali-nebula/go-component-framework/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestNonePattern(t *tes.T) {
	var v = ele.PatternFromString(`^none$`)
	ass.Equal(t, `^none$`, v.AsString())

	var text = ""
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *tes.T) {
	var v = ele.PatternFromString(`.*`)
	ass.Equal(t, `.*`, v.AsString())

	var text = ""
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *tes.T) {
	var v = ele.PatternFromString(`c(.+t)`)
	ass.Equal(t, `c(.+t)`, v.AsString())

	var text = "ct"
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}
