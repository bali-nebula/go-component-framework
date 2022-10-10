/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/craterdog-bali/go-bali-document-notation/language"
	"github.com/craterdog-bali/go-bali-document-notation/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

const n = `">
	This is a narrative... it
	contains a " character and
	spans multiple lines.
<"`

const c = `(
    $base: 64
    $en_coding: $utf8
)`

const d = "~-P12Y3M4DT5H6M7.890S"

const m = "<-1009-08-07T06:05:04.321>"

const p = `"ca+t"?`

const r = "<https://google.com>"

const s = "$getItem"

const v = "v1.2.3"

func TestParserWithElementTypes(t *testing.T) {
	var ok bool
	var component *language.Component

	// Angle
	component, ok = language.ParseSource("~pi($units: $radians)")
	assert.True(t, ok)
	var angle elements.Angle = component.Entity.(elements.Angle)
	assert.Equal(t, elements.Pi, angle)

	// Boolean
	component, ok = language.ParseSource("true")
	assert.True(t, ok)
	var boolean elements.Boolean = component.Entity.(elements.Boolean)
	assert.True(t, boolean.AsBoolean())

	// Duration
	component, ok = language.ParseSource(d)
	assert.True(t, ok)
	var duration elements.Duration = component.Entity.(elements.Duration)
	assert.Equal(t, d, duration.AsString())

	// Moment
	component, ok = language.ParseSource(m)
	assert.True(t, ok)
	var moment elements.Moment = component.Entity.(elements.Moment)
	assert.Equal(t, m, moment.AsString())

	// Number
	component, ok = language.ParseSource("(3, -i)")
	assert.True(t, ok)
	var number elements.Number = component.Entity.(elements.Number)
	assert.Equal(t, "(3, -i)", number.AsString())

	// Pattern
	component, ok = language.ParseSource(p)
	assert.True(t, ok)
	var pattern elements.Pattern = component.Entity.(elements.Pattern)
	assert.Equal(t, p, pattern.AsString())

	// Percentage
	component, ok = language.ParseSource("50%")
	assert.True(t, ok)
	var percentage elements.Percentage = component.Entity.(elements.Percentage)
	assert.Equal(t, 0.5, percentage.AsReal())

	// Probability
	component, ok = language.ParseSource(".75")
	assert.True(t, ok)
	var probability elements.Probability = component.Entity.(elements.Probability)
	assert.Equal(t, 0.75, probability.AsReal())

	// Resource
	component, ok = language.ParseSource(r)
	assert.True(t, ok)
	var resource elements.Resource = component.Entity.(elements.Resource)
	assert.Equal(t, r, resource.AsString())

	// Tag
	component, ok = language.ParseSource("#ABC")
	assert.True(t, ok)
	var tag elements.Tag = component.Entity.(elements.Tag)
	assert.Equal(t, "#ABC", tag.AsString())
}

func TestParserWithStringTypes(t *testing.T) {
	var ok bool
	var component *language.Component

	// Binary
	component, ok = language.ParseSource("'AAAAAAAA' ($base: 64, $encoding: $utf8)")
	assert.True(t, ok)
	var binary strings.Binary = component.Entity.(strings.Binary)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, binary.AsArray())

	// Moniker
	component, ok = language.ParseSource("/bali/types/Number/v1.2.3")
	assert.True(t, ok)
	var moniker strings.Moniker = component.Entity.(strings.Moniker)
	assert.Equal(t, []string{"bali", "types", "Number", "v1.2.3"}, moniker.AsArray())

	// Narrative
	component, ok = language.ParseSource(n + c)
	assert.True(t, ok)
	var narrative strings.Narrative = component.Entity.(strings.Narrative)
	assert.Equal(t, n, narrative.AsString())

	// Quote
	component, ok = language.ParseSource(`"Hello World!"`)
	assert.True(t, ok)
	var quote strings.Quote = component.Entity.(strings.Quote)
	assert.Equal(t, `"Hello World!"`, quote.AsString())

	// Symbol
	component, ok = language.ParseSource(s)
	assert.True(t, ok)
	var symbol elements.Symbol = component.Entity.(elements.Symbol)
	assert.Equal(t, s, symbol.AsString())

	// Version
	component, ok = language.ParseSource(v)
	assert.True(t, ok)
	var version strings.Version = component.Entity.(strings.Version)
	assert.Equal(t, v, version.AsString())
}

const l = `[
	$foo
	$bar
	$baz
]`

const k = `[
	0: false
	1: true
]`

func TestParserWithSequenceTypes(t *testing.T) {
	var ok bool
	var component *language.Component

	// List
	component, ok = language.ParseSource("[ ]")
	assert.True(t, ok)
	list := component.Entity.(abstractions.ListLike[any])
	assert.Equal(t, 0, list.GetSize())

	component, ok = language.ParseSource("[$foo]")
	assert.True(t, ok)
	list = component.Entity.(abstractions.ListLike[any])
	assert.Equal(t, 1, list.GetSize())

	component, ok = language.ParseSource("[$foo, $bar, $baz]")
	assert.True(t, ok)
	list = component.Entity.(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	component, ok = language.ParseSource(l)
	assert.True(t, ok)
	list = component.Entity.(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	// Catalog
	component, ok = language.ParseSource("[:]")
	assert.True(t, ok)
	catalog := component.Entity.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 0, catalog.GetSize())

	component, ok = language.ParseSource("[0: false]")
	assert.True(t, ok)
	catalog = component.Entity.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 1, catalog.GetSize())

	component, ok = language.ParseSource("[0: false, 1: true]")
	assert.True(t, ok)
	catalog = component.Entity.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	component, ok = language.ParseSource(k)
	assert.True(t, ok)
	catalog = component.Entity.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	// Range
	component, ok = language.ParseSource("[1..1]")
	assert.True(t, ok)
	rng := component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 1, rng.GetSize())

	component, ok = language.ParseSource("[1<..1]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[1<..<1]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[1..<1]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[-1..5]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 7, rng.GetSize())

	component, ok = language.ParseSource("[-1<..5]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 6, rng.GetSize())

	component, ok = language.ParseSource("[-1<..<5]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 5, rng.GetSize())

	component, ok = language.ParseSource("[-1..<5]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 6, rng.GetSize())

	component, ok = language.ParseSource("[..]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[<..]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[<..<]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component, ok = language.ParseSource("[..<]")
	assert.True(t, ok)
	rng = component.Entity.(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())
}
