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

func TestZeroIntegers(t *tes.T) {
	var v = ele.Integer.FromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveIntegers(t *tes.T) {
	var v = ele.Integer.FromString(`42`)
	ass.Equal(t, `42`, v.AsString())

	v = ele.Integer.FromString(`+5`)
	ass.Equal(t, `5`, v.AsString())
}

func TestNegativeIntegers(t *tes.T) {
	var v = ele.Integer.FromString(`-1`)
	ass.Equal(t, `-1`, v.AsString())
}
