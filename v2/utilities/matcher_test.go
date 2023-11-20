/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities_test

import (
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestMatchers(t *tes.T) {
	var matches []string

	matches = uti.AngleMatcher.FindStringSubmatch(`~pi`)
	ass.True(t, len(matches) == 2)

	matches = uti.BinaryMatcher.FindStringSubmatch(`'>
    0123456789abcdef
<'`)
	ass.True(t, len(matches) == 2)

	matches = uti.BooleanMatcher.FindStringSubmatch(`true`)
	ass.True(t, len(matches) == 1)

	matches = uti.BytecodeMatcher.FindStringSubmatch(`'1234 abcd'`)
	ass.True(t, len(matches) == 2)

	matches = uti.CharacterMatcher.FindStringSubmatch(`"界"`)
	ass.True(t, len(matches) == 2)

	matches = uti.CommentMatcher.FindStringSubmatch(`!>
    This is a multiline
	comment.
<!`)
	ass.True(t, len(matches) == 2)

	matches = uti.CitationMatcher.FindStringSubmatch(`/bali/types/Set/v1.2.3`)
	ass.True(t, len(matches) == 4)

	matches = uti.DelimiterMatcher.FindStringSubmatch(`:`)
	ass.True(t, len(matches) == 1)

	matches = uti.DurationMatcher.FindStringSubmatch(`~-P3DT4H5M`)
	ass.True(t, len(matches) == 10)

	matches = uti.EolMatcher.FindStringSubmatch(`
`)
	ass.True(t, len(matches) == 1)

	matches = uti.IdentifierMatcher.FindStringSubmatch(`A3D`)
	ass.True(t, len(matches) == 1)

	matches = uti.IntegerMatcher.FindStringSubmatch(`5`)
	ass.True(t, len(matches) == 1)

	matches = uti.IntrinsicMatcher.FindStringSubmatch(`ANY`)
	ass.True(t, len(matches) == 1)

	matches = uti.KeywordMatcher.FindStringSubmatch(`publish`)
	ass.True(t, len(matches) == 1)

	matches = uti.MomentMatcher.FindStringSubmatch(`<-10000-10-15T03:04:05.678>`)
	ass.True(t, len(matches) == 8)

	matches = uti.NameMatcher.FindStringSubmatch(`/bali/types/Set`)
	ass.True(t, len(matches) == 1)

	matches = uti.NarrativeMatcher.FindStringSubmatch(`">
    This is a multiline
	narrative.
<"`)
	ass.True(t, len(matches) == 2)

	matches = uti.NoteMatcher.FindStringSubmatch(`! This is a note.`)
	ass.True(t, len(matches) == 1)

	matches = uti.NumberMatcher.FindStringSubmatch(`(1.0e^~pii)`)
	ass.True(t, len(matches) == 8)

	matches = uti.PatternMatcher.FindStringSubmatch(`"ca+t"?`)
	ass.True(t, len(matches) == 2)

	matches = uti.PercentageMatcher.FindStringSubmatch(`100%`)
	ass.True(t, len(matches) == 2)

	matches = uti.ProbabilityMatcher.FindStringSubmatch(`1.`)
	ass.True(t, len(matches) == 1)

	matches = uti.QuoteMatcher.FindStringSubmatch(`"Hello World!"`)
	ass.True(t, len(matches) == 2)

	matches = uti.RealMatcher.FindStringSubmatch(`1.2E+10`)
	ass.True(t, len(matches) == 1)

	matches = uti.ResourceMatcher.FindStringSubmatch(`<https://google.com>`)
	ass.True(t, len(matches) == 7)

	matches = uti.SymbolMatcher.FindStringSubmatch(`$symbol`)
	ass.True(t, len(matches) == 2)

	matches = uti.TagMatcher.FindStringSubmatch(`#ABCD1234`)
	ass.True(t, len(matches) == 2)

	matches = uti.VersionMatcher.FindStringSubmatch(`v1.2.3`)
	ass.True(t, len(matches) == 2)

	matches = uti.WhitespaceMatcher.FindStringSubmatch(` `)
	ass.True(t, len(matches) == 1)
}
