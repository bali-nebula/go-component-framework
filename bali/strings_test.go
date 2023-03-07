/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	bal "github.com/bali-nebula/go-component-framework/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBinary(t *tes.T) {
	var v = bal.Binary(64)
	ass.Equal(t, 64, v.GetSize())
	v = bal.Binary(`'>
    abcd1234
<'`)
	ass.Equal(t, "abcd1234", v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, v.GetSize())
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), bal.Binary(v.AsArray()).AsArray())
	ass.Equal(t, "abcd", bal.Binary(v.GetValues(1, 3)).AsString())
}

func TestMoniker(t *tes.T) {
	var v1 = bal.Moniker("/bali/types/abstractions/String/v1.2.3")
	ass.Equal(t, "/bali/types/abstractions/String/v1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 5, v1.GetSize())
	ass.Equal(t, abs.Name("bali"), v1.GetValue(1))
	ass.Equal(t, abs.Name("v1.2.3"), v1.GetValue(-1))
	var v2 = bal.Moniker(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = bal.Moniker(v1.GetValues(1, 2))
	ass.Equal(t, "/bali/types", v3.AsString())
}

func TestEmptyNarrative(t *tes.T) {
	var v = bal.Narrative(`">

<"`)
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 1, v.GetSize())
	ass.Equal(t, 1, len(v.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v = bal.Narrative(`">
    This is a narrative
    containing " marks.
<"`)
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, v.GetSize())
	ass.Equal(t, "This is a narrative", string(v.GetValue(1)))
	ass.Equal(t, 2, len(v.AsArray()))
}

func TestEmptyQuote(t *tes.T) {
	var v = bal.Quote(`""`)
	ass.Equal(t, "", v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, v.GetSize())
}

func TestQuote(t *tes.T) {
	var v = bal.Quote(`"abcd本1234"`)
	ass.Equal(t, "abcd本1234", v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, v.GetSize())
	ass.Equal(t, 'a', v.GetValue(1))
	ass.Equal(t, '4', v.GetValue(-1))
	ass.Equal(t, v.AsArray(), bal.Quote(v.AsArray()).AsArray())
	ass.Equal(t, "d本1", bal.Quote(v.GetValues(4, 6)).AsString())
}

func TestVersion(t *tes.T) {
	var v1 = bal.Version("v1.2.3")
	ass.Equal(t, "1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, v1.GetSize())
	ass.Equal(t, abs.Ordinal(1), v1.GetValue(1))
	ass.Equal(t, abs.Ordinal(3), v1.GetValue(-1))
	var v2 = bal.Version(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = bal.Version(v1.GetValues(1, 2))
	ass.Equal(t, "1.2", v3.AsString())
}
