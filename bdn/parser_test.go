/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bdn_test

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	bdn "github.com/craterdog-bali/go-bali-document-notation/bdn"
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
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

func TestParserWithElementTypes(t *tes.T) {
	var component abs.ComponentLike

	// Angle
	component = bdn.ParseSource("~pi($units: $radians)")
	var angle ele.Angle = component.GetEntity().(ele.Angle)
	ass.Equal(t, ele.Pi, angle)

	// Boolean
	component = bdn.ParseSource("true")
	var boolean ele.Boolean = component.GetEntity().(ele.Boolean)
	ass.True(t, boolean.AsBoolean())

	// Duration
	component = bdn.ParseSource(d)
	var duration ele.Duration = component.GetEntity().(ele.Duration)
	ass.Equal(t, d, bdn.FormatValue(duration))

	// Moment
	component = bdn.ParseSource(m)
	var moment ele.Moment = component.GetEntity().(ele.Moment)
	ass.Equal(t, m, bdn.FormatValue(moment))

	// Number
	component = bdn.ParseSource("(3, -i)")
	var number ele.Number = component.GetEntity().(ele.Number)
	ass.Equal(t, "(3, -i)", bdn.FormatValue(number))

	// Pattern
	component = bdn.ParseSource(p)
	var pattern ele.Pattern = component.GetEntity().(ele.Pattern)
	ass.Equal(t, p, bdn.FormatValue(pattern))

	// Percentage
	component = bdn.ParseSource("50%")
	var percentage ele.Percentage = component.GetEntity().(ele.Percentage)
	ass.Equal(t, 0.5, percentage.AsReal())

	// Probability
	component = bdn.ParseSource(".75")
	var probability ele.Probability = component.GetEntity().(ele.Probability)
	ass.Equal(t, 0.75, probability.AsReal())

	// Resource
	component = bdn.ParseSource(r)
	var resource ele.Resource = component.GetEntity().(ele.Resource)
	ass.Equal(t, r, bdn.FormatValue(resource))

	// Tag
	component = bdn.ParseSource("#ABC")
	var tag ele.Tag = component.GetEntity().(ele.Tag)
	ass.Equal(t, "#ABC", bdn.FormatValue(tag))
}

func TestParserWithStringTypes(t *tes.T) {
	var component abs.ComponentLike

	// Binary
	component = bdn.ParseSource("'AAAAAAAA' ($base: 64, $encoding: $utf8)")
	var binary str.Binary = component.GetEntity().(str.Binary)
	ass.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, binary.AsArray())

	// Moniker
	component = bdn.ParseSource("/bali/types/Number/v1.2.3")
	var moniker str.Moniker = component.GetEntity().(str.Moniker)
	ass.Equal(t, []string{"bali", "types", "Number", "v1.2.3"}, moniker.AsArray())

	// Narrative
	component = bdn.ParseSource(n + c)
	var narrative str.Narrative = component.GetEntity().(str.Narrative)
	ass.Equal(t, n, bdn.FormatValue(narrative))

	// Quote
	component = bdn.ParseSource(`"Hello World!"`)
	var quote str.Quote = component.GetEntity().(str.Quote)
	ass.Equal(t, `"Hello World!"`, bdn.FormatValue(quote))

	// Symbol
	component = bdn.ParseSource(s)
	var symbol ele.Symbol = component.GetEntity().(ele.Symbol)
	ass.Equal(t, s, bdn.FormatValue(symbol))

	// Version
	component = bdn.ParseSource(v)
	var version str.Version = component.GetEntity().(str.Version)
	ass.Equal(t, v, bdn.FormatValue(version))
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

func TestParserWithSequenceTypes(t *tes.T) {
	var component abs.ComponentLike

	// List
	component = bdn.ParseSource("[ ]")
	var list = component.GetEntity().(abs.ListLike[abs.ItemLike])
	ass.Equal(t, 0, list.GetSize())

	component = bdn.ParseSource("[$foo]")
	list = component.GetEntity().(abs.ListLike[abs.ItemLike])
	ass.Equal(t, 1, list.GetSize())

	component = bdn.ParseSource("[$foo, $bar, $baz]")
	list = component.GetEntity().(abs.ListLike[abs.ItemLike])
	ass.Equal(t, 3, list.GetSize())

	component = bdn.ParseSource(l)
	list = component.GetEntity().(abs.ListLike[abs.ItemLike])
	ass.Equal(t, 3, list.GetSize())

	// Catalog
	component = bdn.ParseSource("[:]")
	var catalog = component.GetEntity().(abs.CatalogLike[abs.PrimitiveLike, abs.ComponentLike])
	ass.Equal(t, 0, catalog.GetSize())

	component = bdn.ParseSource("[0: false]")
	catalog = component.GetEntity().(abs.CatalogLike[abs.PrimitiveLike, abs.ComponentLike])
	ass.Equal(t, 1, catalog.GetSize())

	component = bdn.ParseSource("[0: false, 1: true]")
	catalog = component.GetEntity().(abs.CatalogLike[abs.PrimitiveLike, abs.ComponentLike])
	ass.Equal(t, 2, catalog.GetSize())

	component = bdn.ParseSource(k)
	catalog = component.GetEntity().(abs.CatalogLike[abs.PrimitiveLike, abs.ComponentLike])
	ass.Equal(t, 2, catalog.GetSize())

	// Range
	component = bdn.ParseSource("[1..1]")
	var rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 1, rng.GetSize())

	component = bdn.ParseSource("(1..1]")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("(1..1)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("[1..1)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("[-1..5]")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 7, rng.GetSize())

	component = bdn.ParseSource("(-1..5]")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 6, rng.GetSize())

	component = bdn.ParseSource("(-1..5)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 5, rng.GetSize())

	component = bdn.ParseSource("[-1..5)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 6, rng.GetSize())

	component = bdn.ParseSource("[..]")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("(..]")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("(..)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())

	component = bdn.ParseSource("[..)")
	rng = component.GetEntity().(abs.RangeLike[abs.PrimitiveLike])
	ass.Equal(t, 0, rng.GetSize())
}
