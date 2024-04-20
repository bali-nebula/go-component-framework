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
	ele "github.com/bali-nebula/go-component-framework/v3/element"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestFalseBooleans(t *tes.T) {
	var Boolean = ele.Boolean()
	var v = Boolean.MakeFromBoolean(false)
	ass.False(t, v.AsBoolean())
}

func TestTrueBooleans(t *tes.T) {
	var Boolean = ele.Boolean()
	var v = Boolean.MakeFromBoolean(true)
	ass.True(t, v.AsBoolean())
}

func TestBooleansLibrary(t *tes.T) {
	var Boolean = ele.Boolean()
	var T = Boolean.MakeFromBoolean(true)
	var F = Boolean.MakeFromBoolean(false)

	var andNot = Boolean.And(Boolean.Not(T), Boolean.Not(T))
	var notOr = Boolean.Not(Boolean.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(T), Boolean.Not(F))
	notOr = Boolean.Not(Boolean.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(F), Boolean.Not(T))
	notOr = Boolean.Not(Boolean.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(F), Boolean.Not(F))
	notOr = Boolean.Not(Boolean.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = Boolean.And(T, Boolean.Not(T))
	ass.Equal(t, sans, Boolean.Sans(T, T))

	sans = Boolean.And(T, Boolean.Not(F))
	ass.Equal(t, sans, Boolean.Sans(T, F))

	sans = Boolean.And(F, Boolean.Not(T))
	ass.Equal(t, sans, Boolean.Sans(F, T))

	sans = Boolean.And(F, Boolean.Not(F))
	ass.Equal(t, sans, Boolean.Sans(F, F))

	var xor = Boolean.Or(Boolean.Sans(T, T), Boolean.Sans(T, T))
	ass.Equal(t, xor, Boolean.Xor(T, T))

	xor = Boolean.Or(Boolean.Sans(T, F), Boolean.Sans(F, T))
	ass.Equal(t, xor, Boolean.Xor(T, F))

	xor = Boolean.Or(Boolean.Sans(F, T), Boolean.Sans(T, F))
	ass.Equal(t, xor, Boolean.Xor(F, T))

	xor = Boolean.Or(Boolean.Sans(F, F), Boolean.Sans(F, F))
	ass.Equal(t, xor, Boolean.Xor(F, F))
}
