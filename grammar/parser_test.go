/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/craterdog-bali/go-bali-document-notation/grammar"
	"github.com/craterdog-bali/go-bali-document-notation/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

const n = `"
This is a narrative... it
contains a " character and
spans multiple lines.
"`

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
	var ok bool
	var component *grammar.Component

	// Angle
	component, ok = grammar.ParseSource("~pi ($units: $radians) ! note")
	assert.True(t, ok)
	var angle elements.Angle = component.Value.(elements.Angle)
	assert.Equal(t, elements.Pi, angle)

	// Boolean
	component, ok = grammar.ParseSource("true ! false")
	assert.True(t, ok)
	var boolean elements.Boolean = component.Value.(elements.Boolean)
	assert.True(t, boolean.AsBoolean())

	// Duration
	component, ok = grammar.ParseSource(d)
	assert.True(t, ok)
	var duration elements.Duration = component.Value.(elements.Duration)
	assert.Equal(t, d, duration.AsString())

	// Moment
	component, ok = grammar.ParseSource(m)
	assert.True(t, ok)
	var moment elements.Moment = component.Value.(elements.Moment)
	assert.Equal(t, m, moment.AsString())

	// Number
	component, ok = grammar.ParseSource("(3, 4i)")
	assert.True(t, ok)
	var number elements.Number = component.Value.(elements.Number)
	assert.Equal(t, "(3, 4i)", number.AsString())

	// Pattern
	component, ok = grammar.ParseSource(p)
	assert.True(t, ok)
	var pattern elements.Pattern = component.Value.(elements.Pattern)
	assert.Equal(t, p, pattern.AsString())

	// Percentage
	component, ok = grammar.ParseSource("50%")
	assert.True(t, ok)
	var percentage elements.Percentage = component.Value.(elements.Percentage)
	assert.Equal(t, 0.5, percentage.AsReal())

	// Probability
	component, ok = grammar.ParseSource(".75")
	assert.True(t, ok)
	var probability elements.Probability = component.Value.(elements.Probability)
	assert.Equal(t, 0.75, probability.AsReal())

	// Resource
	component, ok = grammar.ParseSource(r)
	assert.True(t, ok)
	var resource elements.Resource = component.Value.(elements.Resource)
	assert.Equal(t, r, resource.AsString())

	// Tag
	component, ok = grammar.ParseSource("#ABC")
	assert.True(t, ok)
	var tag elements.Tag = component.Value.(elements.Tag)
	assert.Equal(t, "#ABC", tag.AsString())
}

func TestParserWithStringTypes(t *testing.T) {
	var ok bool
	var component *grammar.Component

	// Binary
	component, ok = grammar.ParseSource("'AAAAAAAA' ($base: 64, $encoding: $utf8)")
	assert.True(t, ok)
	var binary strings.Binary = component.Value.(strings.Binary)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, binary.AsArray())

	// Moniker
	component, ok = grammar.ParseSource("/bali/types/Number/v1.2.3")
	assert.True(t, ok)
	var moniker strings.Moniker = component.Value.(strings.Moniker)
	assert.Equal(t, []string{"bali", "types", "Number", "v1.2.3"}, moniker.AsArray())

	// Narrative
	component, ok = grammar.ParseSource(n + c)
	assert.True(t, ok)
	var narrative strings.Narrative = component.Value.(strings.Narrative)
	assert.Equal(t, n, narrative.AsString())

	// Quote
	component, ok = grammar.ParseSource(`"Hello World!"`)
	assert.True(t, ok)
	var quote strings.Quote = component.Value.(strings.Quote)
	assert.Equal(t, `"Hello World!"`, quote.AsString())

	// Symbol
	component, ok = grammar.ParseSource(s)
	assert.True(t, ok)
	var symbol elements.Symbol = component.Value.(elements.Symbol)
	assert.Equal(t, s, symbol.AsString())

	// Version
	component, ok = grammar.ParseSource(v)
	assert.True(t, ok)
	var version strings.Version = component.Value.(strings.Version)
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
	var component *grammar.Component

	// List
	component, ok = grammar.ParseSource("[ ]")
	assert.True(t, ok)
	list := component.Value.(abstractions.ListLike[any])
	assert.Equal(t, 0, list.GetSize())

	component, ok = grammar.ParseSource("[$foo]")
	assert.True(t, ok)
	list = component.Value.(abstractions.ListLike[any])
	assert.Equal(t, 1, list.GetSize())

	component, ok = grammar.ParseSource("[$foo, $bar, $baz]")
	assert.True(t, ok)
	list = component.Value.(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	component, ok = grammar.ParseSource(l)
	assert.True(t, ok)
	list = component.Value.(abstractions.ListLike[any])
	assert.Equal(t, 3, list.GetSize())

	// Catalog
	component, ok = grammar.ParseSource("[:]")
	assert.True(t, ok)
	catalog := component.Value.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 0, catalog.GetSize())

	component, ok = grammar.ParseSource("[0: false]")
	assert.True(t, ok)
	catalog = component.Value.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 1, catalog.GetSize())

	component, ok = grammar.ParseSource("[0: false, 1: true]")
	assert.True(t, ok)
	catalog = component.Value.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	component, ok = grammar.ParseSource(k)
	assert.True(t, ok)
	catalog = component.Value.(abstractions.CatalogLike[any, any])
	assert.Equal(t, 2, catalog.GetSize())

	// Slice
	component, ok = grammar.ParseSource("[1..1]")
	assert.True(t, ok)
	slice := component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 1, slice.GetSize())

	component, ok = grammar.ParseSource("[1<..1]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[1<..<1]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[1..<1]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[-1..5]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 7, slice.GetSize())

	component, ok = grammar.ParseSource("[-1<..5]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 6, slice.GetSize())

	component, ok = grammar.ParseSource("[-1<..<5]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 5, slice.GetSize())

	component, ok = grammar.ParseSource("[-1..<5]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 6, slice.GetSize())

	component, ok = grammar.ParseSource("[..]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[<..]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[<..<]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())

	component, ok = grammar.ParseSource("[..<]")
	assert.True(t, ok)
	slice = component.Value.(abstractions.SliceLike[any])
	assert.Equal(t, 0, slice.GetSize())
}
