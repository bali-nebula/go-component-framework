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
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestFalseBooleans(t *tes.T) {
	var v = ele.Boolean.FromBoolean(false)
	ass.False(t, v.AsBoolean())
}

func TestTrueBooleans(t *tes.T) {
	var v = ele.Boolean.FromBoolean(true)
	ass.True(t, v.AsBoolean())
}

func TestBooleansLibrary(t *tes.T) {
	var T = ele.Boolean.FromBoolean(true)
	var F = ele.Boolean.FromBoolean(false)

	var andNot = ele.Boolean.And(ele.Boolean.Not(T), ele.Boolean.Not(T))
	var notOr = ele.Boolean.Not(ele.Boolean.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Boolean.And(ele.Boolean.Not(T), ele.Boolean.Not(F))
	notOr = ele.Boolean.Not(ele.Boolean.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Boolean.And(ele.Boolean.Not(F), ele.Boolean.Not(T))
	notOr = ele.Boolean.Not(ele.Boolean.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Boolean.And(ele.Boolean.Not(F), ele.Boolean.Not(F))
	notOr = ele.Boolean.Not(ele.Boolean.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = ele.Boolean.And(T, ele.Boolean.Not(T))
	ass.Equal(t, sans, ele.Boolean.Sans(T, T))

	sans = ele.Boolean.And(T, ele.Boolean.Not(F))
	ass.Equal(t, sans, ele.Boolean.Sans(T, F))

	sans = ele.Boolean.And(F, ele.Boolean.Not(T))
	ass.Equal(t, sans, ele.Boolean.Sans(F, T))

	sans = ele.Boolean.And(F, ele.Boolean.Not(F))
	ass.Equal(t, sans, ele.Boolean.Sans(F, F))

	var xor = ele.Boolean.Or(ele.Boolean.Sans(T, T), ele.Boolean.Sans(T, T))
	ass.Equal(t, xor, ele.Boolean.Xor(T, T))

	xor = ele.Boolean.Or(ele.Boolean.Sans(T, F), ele.Boolean.Sans(F, T))
	ass.Equal(t, xor, ele.Boolean.Xor(T, F))

	xor = ele.Boolean.Or(ele.Boolean.Sans(F, T), ele.Boolean.Sans(T, F))
	ass.Equal(t, xor, ele.Boolean.Xor(F, T))

	xor = ele.Boolean.Or(ele.Boolean.Sans(F, F), ele.Boolean.Sans(F, F))
	ass.Equal(t, xor, ele.Boolean.Xor(F, F))
}
