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

func TestBooleanProbabilities(t *testing.T) {
	var v1 = elements.ProbabilityFromBoolean(false)
	assert.False(t, v1.AsBoolean())
	assert.Equal(t, 0, v1.AsInteger())
	assert.Equal(t, 0.0, v1.AsReal())
	assert.Equal(t, 0.0, float64(v1))

	var v2 = elements.ProbabilityFromBoolean(true)
	assert.True(t, v2.AsBoolean())
	assert.Equal(t, 1, v2.AsInteger())
	assert.Equal(t, 1.0, v2.AsReal())
	assert.Equal(t, 1.0, float64(v2))
}

func TestZeroProbabilities(t *testing.T) {
	var v = elements.ProbabilityFromFloat(0.0)
	assert.False(t, v.AsBoolean())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, float64(v))
}

func TestOneProbabilities(t *testing.T) {
	var v = elements.ProbabilityFromFloat(1.0)
	assert.True(t, v.AsBoolean())
	assert.Equal(t, 1, v.AsInteger())
	assert.Equal(t, 1.0, v.AsReal())
	assert.Equal(t, 1.0, float64(v))
}

func TestRandomProbabilities(t *testing.T) {
	var v = elements.Probabilities.Random()
	assert.True(t, v.AsBoolean() || !v.AsBoolean())
}

func TestOtherProbabilities(t *testing.T) {
	var v1 = elements.ProbabilityFromFloat(0.25)
	assert.False(t, v1.AsBoolean())
	assert.Equal(t, 0, v1.AsInteger())
	assert.Equal(t, 0.25, v1.AsReal())
	assert.Equal(t, 0.25, float64(v1))

	var v2 = elements.ProbabilityFromFloat(0.5)
	assert.True(t, v2.AsBoolean())
	assert.Equal(t, 1, v2.AsInteger())
	assert.Equal(t, 0.5, v2.AsReal())
	assert.Equal(t, 0.5, float64(v2))

	var v3 = elements.ProbabilityFromFloat(0.75)
	assert.True(t, v3.AsBoolean())
	assert.Equal(t, 1, v3.AsInteger())
	assert.Equal(t, 0.75, v3.AsReal())
	assert.Equal(t, 0.75, float64(v3))
}

func TestStringProbabilities(t *testing.T) {
	var v, ok = elements.ProbabilityFromString(".0")
	assert.True(t, ok)
	assert.Equal(t, 0, int(v)) // Truncates the value.
	assert.False(t, v.AsBoolean())
	assert.Equal(t, 0, v.AsInteger()) // Rounds the value.
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, ".0", v.AsString())

	v, ok = elements.ProbabilityFromString(".5")
	assert.True(t, ok)
	assert.Equal(t, 0, int(v)) // Truncates the value.
	assert.True(t, v.AsBoolean())
	assert.Equal(t, 1, v.AsInteger()) // Rounds the value.
	assert.Equal(t, 0.5, v.AsReal())
	assert.Equal(t, ".5", v.AsString())

	v, ok = elements.ProbabilityFromString("1.")
	assert.True(t, ok)
	assert.Equal(t, 1, int(v))
	assert.True(t, v.AsBoolean())
	assert.Equal(t, 1, v.AsInteger())
	assert.Equal(t, 1.0, v.AsReal())
	assert.Equal(t, "1.", v.AsString())

	v, ok = elements.ProbabilityFromString("1.0")
	assert.False(t, ok)
	assert.Equal(t, 0, int(v))
}

func TestProbabilitiesLibrary(t *testing.T) {
	var T = elements.ProbabilityFromFloat(0.75)
	var F = elements.ProbabilityFromFloat(0.25)

	var andNot = elements.Probabilities.And(elements.Probabilities.Not(T), elements.Probabilities.Not(T))
	var notOr = elements.Probabilities.Not(elements.Probabilities.Or(T, T))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Probabilities.And(elements.Probabilities.Not(T), elements.Probabilities.Not(F))
	notOr = elements.Probabilities.Not(elements.Probabilities.Or(T, F))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Probabilities.And(elements.Probabilities.Not(F), elements.Probabilities.Not(T))
	notOr = elements.Probabilities.Not(elements.Probabilities.Or(F, T))
	assert.Equal(t, andNot, notOr)

	andNot = elements.Probabilities.And(elements.Probabilities.Not(F), elements.Probabilities.Not(F))
	notOr = elements.Probabilities.Not(elements.Probabilities.Or(F, F))
	assert.Equal(t, andNot, notOr)

	var sans = elements.Probabilities.And(T, elements.Probabilities.Not(T))
	assert.Equal(t, sans, elements.Probabilities.Sans(T, T))

	sans = elements.Probabilities.And(T, elements.Probabilities.Not(F))
	assert.Equal(t, sans, elements.Probabilities.Sans(T, F))

	sans = elements.Probabilities.And(F, elements.Probabilities.Not(T))
	assert.Equal(t, sans, elements.Probabilities.Sans(F, T))

	sans = elements.Probabilities.And(F, elements.Probabilities.Not(F))
	assert.Equal(t, sans, elements.Probabilities.Sans(F, F))

	var xor = elements.Probabilities.Or(elements.Probabilities.Sans(T, T), elements.Probabilities.Sans(T, T))
	assert.Equal(t, xor, elements.Probabilities.Xor(T, T))

	xor = elements.Probabilities.Or(elements.Probabilities.Sans(T, F), elements.Probabilities.Sans(F, T))
	assert.Equal(t, xor, elements.Probabilities.Xor(T, F))

	xor = elements.Probabilities.Or(elements.Probabilities.Sans(F, T), elements.Probabilities.Sans(T, F))
	assert.Equal(t, xor, elements.Probabilities.Xor(F, T))

	xor = elements.Probabilities.Or(elements.Probabilities.Sans(F, F), elements.Probabilities.Sans(F, F))
	assert.Equal(t, xor, elements.Probabilities.Xor(F, F))
}
