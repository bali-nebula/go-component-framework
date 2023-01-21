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
	bal "github.com/bali-nebula/go-component-framework/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

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
		var component = bal.ParseComponent(s)
		var s = bal.FormatComponent(component)
		ass.Equal(t, rangeStrings[index], s)
	}
}
