/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element_test

import (
	//fmt "fmt"
	ele "github.com/bali-nebula/go-component-framework/v3/element"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestInfiniteFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`+infinity`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`+∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`-infinity`)
	ass.Equal(t, `-∞`, v.AsString())

	v = Float.MakeFromString(`-∞`)
	ass.Equal(t, `-∞`, v.AsString())
}

func TestZeroFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`12.3`)
	ass.Equal(t, `12.3`, v.AsString())

	v = Float.MakeFromString(`+π`)
	ass.Equal(t, `π`, v.AsString())
}

func TestNegativeFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`-e`)
	ass.Equal(t, `-e`, v.AsString())
}
