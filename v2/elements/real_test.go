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

func TestUndefinedReals(t *tes.T) {
	var v = ele.RealFromString(`undefined`)
	ass.Equal(t, `undefined`, v.AsString())
}

func TestInfiniteReals(t *tes.T) {
	var v = ele.RealFromString(`+infinity`)
	ass.Equal(t, `+∞`, v.AsString())

	v = ele.RealFromString(`+∞`)
	ass.Equal(t, `+∞`, v.AsString())

	v = ele.RealFromString(`-infinity`)
	ass.Equal(t, `-∞`, v.AsString())

	v = ele.RealFromString(`-∞`)
	ass.Equal(t, `-∞`, v.AsString())
}

func TestZeroReals(t *tes.T) {
	var v = ele.RealFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveReals(t *tes.T) {
	var v = ele.RealFromString(`12.3`)
	ass.Equal(t, `12.3`, v.AsString())

	v = ele.RealFromString(`+π`)
	ass.Equal(t, `π`, v.AsString())
}

func TestNegativeReals(t *tes.T) {
	var v = ele.RealFromString(`-e`)
	ass.Equal(t, `-e`, v.AsString())
}
