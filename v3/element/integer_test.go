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

func TestZeroIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`42`)
	ass.Equal(t, `42`, v.AsString())

	v = Integer.MakeFromString(`+5`)
	ass.Equal(t, `5`, v.AsString())
}

func TestNegativeIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`-1`)
	ass.Equal(t, `-1`, v.AsString())
}
