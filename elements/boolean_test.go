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
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFalseBooleans(t *testing.T) {
	var v = elements.Boolean(false)
	assert.False(t, bool(v))
	assert.False(t, v.AsBoolean())
	assert.Equal(t, "false", v.AsString())
	assert.Equal(t, 0, v.AsInteger())
}

func TestTrueBooleans(t *testing.T) {
	var v = elements.Boolean(true)
	assert.True(t, bool(v))
	assert.True(t, v.AsBoolean())
	assert.Equal(t, "true", v.AsString())
	assert.Equal(t, 1, v.AsInteger())
}

func TestStringBooleans(t *testing.T) {
	var v, ok = elements.BooleanFromString("false")
	assert.True(t, ok)
	assert.False(t, bool(v))
	assert.False(t, v.AsBoolean())
	assert.Equal(t, "false", v.AsString())
	assert.Equal(t, 0, v.AsInteger())

	v, ok = elements.BooleanFromString("true")
	assert.True(t, bool(v))
	assert.True(t, v.AsBoolean())
	assert.Equal(t, "true", v.AsString())
	assert.Equal(t, 1, v.AsInteger())

	v, ok = elements.BooleanFromString("bad")
	assert.False(t, ok)
	assert.False(t, bool(v))
}

func TestBooleansLibrary(t *testing.T) {
	var T = elements.Boolean(true)
	var F = elements.Boolean(false)

	var andNot = elements.Booleans.And(elements.Booleans.Not(T), elements.Booleans.Not(T))
	var notOr = elements.Booleans.Not(elements.Booleans.Or(T, T))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Booleans.And(elements.Booleans.Not(T), elements.Booleans.Not(F))
	notOr = elements.Booleans.Not(elements.Booleans.Or(T, F))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Booleans.And(elements.Booleans.Not(F), elements.Booleans.Not(T))
	notOr = elements.Booleans.Not(elements.Booleans.Or(F, T))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Booleans.And(elements.Booleans.Not(F), elements.Booleans.Not(F))
	notOr = elements.Booleans.Not(elements.Booleans.Or(F, F))
	assert.Equal(t, andNot, notOr)

	var sans = elements.Booleans.And(T, elements.Booleans.Not(T))
	assert.Equal(t, sans, elements.Booleans.Sans(T, T))

	sans = elements.Booleans.And(T, elements.Booleans.Not(F))
	assert.Equal(t, sans, elements.Booleans.Sans(T, F))

	sans = elements.Booleans.And(F, elements.Booleans.Not(T))
	assert.Equal(t, sans, elements.Booleans.Sans(F, T))

	sans = elements.Booleans.And(F, elements.Booleans.Not(F))
	assert.Equal(t, sans, elements.Booleans.Sans(F, F))

	var xor = elements.Booleans.Or(elements.Booleans.Sans(T, T), elements.Booleans.Sans(T, T))
	assert.Equal(t, xor, elements.Booleans.Xor(T, T))

	xor = elements.Booleans.Or(elements.Booleans.Sans(T, F), elements.Booleans.Sans(F, T))
	assert.Equal(t, xor, elements.Booleans.Xor(T, F))

	xor = elements.Booleans.Or(elements.Booleans.Sans(F, T), elements.Booleans.Sans(T, F))
	assert.Equal(t, xor, elements.Booleans.Xor(F, T))

	xor = elements.Booleans.Or(elements.Booleans.Sans(F, F), elements.Booleans.Sans(F, F))
	assert.Equal(t, xor, elements.Booleans.Xor(F, F))
}
