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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithCatalogs(t *tes.T) {
	var catalogStrings = []string{
		`[:]`,
		"[\n" +
		"    $key: \"value\"\n" +
		"]",
		"[\n" +
		"    $foo: 5\n" +
		"    $bar: ~π\n" +
		"]",
	}

	for index, s := range catalogStrings {
		var catalog = lan.ParseSource(s).GetEntity().(abs.CatalogLike[any, any])
		var s = lan.FormatValue(catalog)
		ass.Equal(t, catalogStrings[index], s)
	}
}

func TestRoundtripWithLists(t *tes.T) {
	var listStrings = []string{
		`[ ]`,
		"[\n" +
		"    42\n" +
		"]",
		"[\n" +
		"    $foo\n" +
		"    $bar\n" +
		"]",
	}

	for index, s := range listStrings {
		var list = lan.ParseSource(s).GetEntity().(abs.ListLike[any])
		var s = lan.FormatValue(list)
		ass.Equal(t, listStrings[index], s)
	}
}

func TestRoundtripWithRanges(t *tes.T) {
	var rangeStrings = []string{
		`[none..none]`,
		`[0..infinity]`,
	}

	for index, s := range rangeStrings {
		var r = lan.ParseSource(s).GetEntity().(abs.RangeLike[any])
		var s = lan.FormatValue(r)
		ass.Equal(t, rangeStrings[index], s)
	}
}
