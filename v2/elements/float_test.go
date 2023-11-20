/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	//fmt "fmt"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestInfiniteFloats(t *tes.T) {
	var v = ele.FloatFromString(`+infinity`)
	ass.Equal(t, `∞`, v.AsString())

	v = ele.FloatFromString(`+∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = ele.FloatFromString(`∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = ele.FloatFromString(`-infinity`)
	ass.Equal(t, `-∞`, v.AsString())

	v = ele.FloatFromString(`-∞`)
	ass.Equal(t, `-∞`, v.AsString())
}

func TestZeroFloats(t *tes.T) {
	var v = ele.FloatFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveFloats(t *tes.T) {
	var v = ele.FloatFromString(`12.3`)
	ass.Equal(t, `12.3`, v.AsString())

	v = ele.FloatFromString(`+π`)
	ass.Equal(t, `π`, v.AsString())
}

func TestNegativeFloats(t *tes.T) {
	var v = ele.FloatFromString(`-e`)
	ass.Equal(t, `-e`, v.AsString())
}
