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
    $encoding: $utf8
)`

const d = "~-P12Y3M4DT5H6M7.890S"

const m = "<-1009-08-07T06:05:04.321>"

const p = `"ca+t"?`

const r = "<https://google.com>"

const s = "$getItem"

const v = "v1.2.3"

func TestParserWithElementTypes(t *testing.T) {
	var component abstractions.ComponentLike[any]

	// Angle
	component = language.ParseSource("~pi($units: $radians)")
	var angle elements.Angle = component.GetEntity().(elements.Angle)
	assert.Equal(t, elements.Pi, angle)

	// Boolean
	component = language.ParseSource("true")
	var boolean elements.Boolean = component.GetEntity().(elements.Boolean)
	assert.True(t, boolean.AsBoolean())

	// Duration
	component = language.ParseSource(d)
	var duration elements.Duration = component.GetEntity().(elements.Duration)
	assert.Equal(t, d, duration.AsString())

	// Moment
	component = language.ParseSource(m)
	var moment elements.Moment = component.GetEntity().(elements.Moment)
	assert.Equal(t, m, moment.AsString())

	// Number
	component = language.ParseSource("(3, -i)")
	var number elements.Number = component.GetEntity().(elements.Number)
	assert.Equal(t, "(3, -i)", number.AsString())

	// Pattern
	component = language.ParseSource(p)
	var pattern elements.Pattern = component.GetEntity().(elements.Pattern)
	assert.Equal(t, p, pattern.AsString())

	// Percentage
	component = language.ParseSource("50%")
	var percentage elements.Percentage = component.GetEntity().(elements.Percentage)
	assert.Equal(t, 0.5, percentage.AsReal())

	// Probability
	component = language.ParseSource(".75")
	var probability elements.Probability = component.GetEntity().(elements.Probability)
	assert.Equal(t, 0.75, probability.AsReal())

	// Resource
	component = language.ParseSource(r)
	var resource elements.Resource = component.GetEntity().(elements.Resource)
	assert.Equal(t, r, resource.AsString())

	// Tag
	component = language.ParseSource("#ABC")
	var tag elements.Tag = component.GetEntity().(elements.Tag)
	assert.Equal(t, "#ABC", tag.AsString())
}

func TestParserWithStringTypes(t *testing.T) {
	var component abstractions.ComponentLike[any]

	// Binary
	component = language.ParseSource("'AAAAAAAA' ($base: 64, $encoding: $utf8)")
	var binary strings.Binary = component.GetEntity().(strings.Binary)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, binary.AsArray())

	// Moniker
	component = language.ParseSource("/bali/types/Number/v1.2.3")
	var moniker strings.Moniker = component.GetEntity().(strings.Moniker)
	assert.Equal(t, []string{"bali", "types", "Number", "v1.2.3"}, moniker.AsArray())

	// Narrative
	component = language.ParseSource(n + c)
	var narrative strings.Narrative = component.GetEntity().(strings.Narrative)
	assert.Equal(t, n, narrative.AsString())

	// Quote
	component = language.ParseSource(`"Hello World!"`)
	var quote strings.Quote = component.GetEntity().(strings.Quote)
	assert.Equal(t, `"Hello World!"`, quote.AsString())

	// Symbol
	component = language.ParseSource(s)
	var symbol elements.Symbol = component.GetEntity().(elements.Symbol)
	assert.Equal(t, s, symbol.AsString())

	// Version
	component = language.ParseSource(v)
	var version strings.Version = component.GetEntity().(strings.Version)
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
	var component abstractions.ComponentLike[any]

	// List
	component = language.ParseSource("[ ]")
	var list = component.GetEntity().(abstractions.ListLike[any])
	assert.Equal(t, 0, list.GetSize())

	component = language.ParseSource("[$foo]")
	list = component.GetEntity().(abstractions.ListLike[any])
	assert.Equal(t, 1, list.GetSize())

	component = language.ParseSource("[$foo, $bar, $baz]")
	list = component.GetEntity().(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	component = language.ParseSource(l)
	list = component.GetEntity().(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	// Catalog
	component = language.ParseSource("[:]")
	var catalog = component.GetEntity().(abstractions.CatalogLike[any, any])
	assert.Equal(t, 0, catalog.GetSize())

	component = language.ParseSource("[0: false]")
	catalog = component.GetEntity().(abstractions.CatalogLike[any, any])
	assert.Equal(t, 1, catalog.GetSize())

	component = language.ParseSource("[0: false, 1: true]")
	catalog = component.GetEntity().(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	component = language.ParseSource(k)
	catalog = component.GetEntity().(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	// Range
	component = language.ParseSource("[1..1]")
	var rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 1, rng.GetSize())

	component = language.ParseSource("[1<..1]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[1<..<1]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[1..<1]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[-1..5]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 7, rng.GetSize())

	component = language.ParseSource("[-1<..5]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 6, rng.GetSize())

	component = language.ParseSource("[-1<..<5]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 5, rng.GetSize())

	component = language.ParseSource("[-1..<5]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 6, rng.GetSize())

	component = language.ParseSource("[..]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[<..]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[<..<]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())

	component = language.ParseSource("[..<]")
	rng = component.GetEntity().(abstractions.RangeLike[any])
	assert.Equal(t, 0, rng.GetSize())
}
