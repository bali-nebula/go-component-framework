/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestFalseBooleans(t *tes.T) {
	var v = ele.Boolean(false)
	ass.False(t, bool(v))
	ass.False(t, v.AsBoolean())
	ass.Equal(t, 0, v.AsInteger())
}

func TestTrueBooleans(t *tes.T) {
	var v = ele.Boolean(true)
	ass.True(t, bool(v))
	ass.True(t, v.AsBoolean())
	ass.Equal(t, 1, v.AsInteger())
}

func TestBooleansLibrary(t *tes.T) {
	var T = ele.Boolean(true)
	var F = ele.Boolean(false)

	var andNot = ele.Booleans.And(ele.Booleans.Not(T), ele.Booleans.Not(T))
	var notOr = ele.Booleans.Not(ele.Booleans.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Booleans.And(ele.Booleans.Not(T), ele.Booleans.Not(F))
	notOr = ele.Booleans.Not(ele.Booleans.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Booleans.And(ele.Booleans.Not(F), ele.Booleans.Not(T))
	notOr = ele.Booleans.Not(ele.Booleans.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Booleans.And(ele.Booleans.Not(F), ele.Booleans.Not(F))
	notOr = ele.Booleans.Not(ele.Booleans.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = ele.Booleans.And(T, ele.Booleans.Not(T))
	ass.Equal(t, sans, ele.Booleans.Sans(T, T))

	sans = ele.Booleans.And(T, ele.Booleans.Not(F))
	ass.Equal(t, sans, ele.Booleans.Sans(T, F))

	sans = ele.Booleans.And(F, ele.Booleans.Not(T))
	ass.Equal(t, sans, ele.Booleans.Sans(F, T))

	sans = ele.Booleans.And(F, ele.Booleans.Not(F))
	ass.Equal(t, sans, ele.Booleans.Sans(F, F))

	var xor = ele.Booleans.Or(ele.Booleans.Sans(T, T), ele.Booleans.Sans(T, T))
	ass.Equal(t, xor, ele.Booleans.Xor(T, T))

	xor = ele.Booleans.Or(ele.Booleans.Sans(T, F), ele.Booleans.Sans(F, T))
	ass.Equal(t, xor, ele.Booleans.Xor(T, F))

	xor = ele.Booleans.Or(ele.Booleans.Sans(F, T), ele.Booleans.Sans(T, F))
	ass.Equal(t, xor, ele.Booleans.Xor(F, T))

	xor = ele.Booleans.Or(ele.Booleans.Sans(F, F), ele.Booleans.Sans(F, F))
	ass.Equal(t, xor, ele.Booleans.Xor(F, F))
}
