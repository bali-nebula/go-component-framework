/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	bal "github.com/craterdog-bali/go-component-framework/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithCatalogs(t *tes.T) {
	var catalogStrings = []string{
		`[:]`,
		"[$key: \"value\"]",
		"[\n" +
		"    $foo: 5\n" +
		"    $bar: ~π\n" +
		"]",
	}

	for index, s := range catalogStrings {
		var component = bal.ParseSource(s).(abs.ComponentLike)
		var s = bal.FormatComponent(component)
		ass.Equal(t, catalogStrings[index], s)
	}
}

func TestRoundtripWithLists(t *tes.T) {
	var listStrings = []string{
		`[ ]`,
		"[42]",
		"[\n" +
		"    $foo\n" +
		"    $bar\n" +
		"]",
	}

	for index, s := range listStrings {
		var component = bal.ParseSource(s).(abs.ComponentLike)
		var s = bal.FormatComponent(component)
		ass.Equal(t, listStrings[index], s)
	}
}

func TestRoundtripWithRanges(t *tes.T) {
	var rangeStrings = []string{
		`[1..+∞]`,
		`[-∞..0)`,
		`[~0..~π)`,
		`[~π..~τ)`,
		`[1..100]`,
		`(-1..1)`,
		`[.25...75)`,
	}

	for index, s := range rangeStrings {
		var component = bal.ParseSource(s).(abs.ComponentLike)
		var s = bal.FormatComponent(component)
		ass.Equal(t, rangeStrings[index], s)
	}
}
