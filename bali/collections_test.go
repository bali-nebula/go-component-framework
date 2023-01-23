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
	fmt "fmt"
	bal "github.com/bali-nebula/go-component-framework/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestDummy(t *tes.T) {
	var set = bal.Set(
		bal.Components{
			1, 3, 4, 2,
	}, bal.Parameters{
		{"$foobar", 2},
	})
	fmt.Println(bal.FormatComponent(set))
	var list = bal.List(
		bal.Components{
			1, 2, 3, 4,
	}, bal.Parameters{
	})
	fmt.Println(bal.FormatComponent(list))
	var catalog = bal.Catalog(
		bal.Associations{
		{1, 2},
		{3, 4},
		{5, 6},
	}, bal.Parameters{
		{"$foobar", 2},
	})
	fmt.Println(bal.FormatComponent(catalog))
}

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
		var component = bal.ParseComponent(s)
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
		var component = bal.ParseComponent(s)
		var s = bal.FormatComponent(component)
		ass.Equal(t, listStrings[index], s)
	}
}
