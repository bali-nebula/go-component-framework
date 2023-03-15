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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const n0 = ""

const n1 = "	abcd本\n1234"

const n2 = "	abcd本"

const n3 = "1234"

func TestEmptyNarrative(t *tes.T) {
	var v = str.NarrativeFromString(n0)
	ass.Equal(t, n0, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 1, v.GetSize())
	ass.Equal(t, 1, len(v.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v = str.NarrativeFromString(n1)
	ass.Equal(t, n1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, v.GetSize())
	ass.Equal(t, n2, string(v.GetValue(1)))
	ass.Equal(t, n3, string(v.GetValue(-1)))
	ass.Equal(t, n1, str.NarrativeFromArray([]abs.Line{n2, n3}).AsString())
	ass.Equal(t, 2, len(v.AsArray()))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = str.NarrativeFromString(n2)
	var v2 = str.NarrativeFromString(n3)
	ass.Equal(t, n1, str.Narratives.Concatenate(v1, v2).AsString())
}
