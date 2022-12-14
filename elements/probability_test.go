/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/bali-nebula/go-component-framework/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBooleanProbabilities(t *tes.T) {
	var v1 = ele.ProbabilityFromBoolean(false)
	ass.Equal(t, 0.0, v1.AsReal())
	ass.Equal(t, 0.0, float64(v1))

	var v2 = ele.ProbabilityFromBoolean(true)
	ass.Equal(t, 1.0, v2.AsReal())
	ass.Equal(t, 1.0, float64(v2))
}

func TestZeroProbabilities(t *tes.T) {
	var v = ele.ProbabilityFromFloat(0.0)
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, float64(v))
}

func TestOneProbabilities(t *tes.T) {
	var v = ele.ProbabilityFromFloat(1.0)
	ass.Equal(t, 1.0, v.AsReal())
	ass.Equal(t, 1.0, float64(v))
}

func TestRandomProbabilities(t *tes.T) {
	ele.Probabilities.Random()
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = ele.ProbabilityFromFloat(0.25)
	ass.Equal(t, 0.25, v1.AsReal())
	ass.Equal(t, 0.25, float64(v1))

	var v2 = ele.ProbabilityFromFloat(0.5)
	ass.Equal(t, 0.5, v2.AsReal())
	ass.Equal(t, 0.5, float64(v2))

	var v3 = ele.ProbabilityFromFloat(0.75)
	ass.Equal(t, 0.75, v3.AsReal())
	ass.Equal(t, 0.75, float64(v3))
}

func TestProbabilitiesLibrary(t *tes.T) {
	var T = ele.ProbabilityFromFloat(0.75)
	var F = ele.ProbabilityFromFloat(0.25)

	var andNot = ele.Probabilities.And(ele.Probabilities.Not(T), ele.Probabilities.Not(T))
	var notOr = ele.Probabilities.Not(ele.Probabilities.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probabilities.And(ele.Probabilities.Not(T), ele.Probabilities.Not(F))
	notOr = ele.Probabilities.Not(ele.Probabilities.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probabilities.And(ele.Probabilities.Not(F), ele.Probabilities.Not(T))
	notOr = ele.Probabilities.Not(ele.Probabilities.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = ele.Probabilities.And(ele.Probabilities.Not(F), ele.Probabilities.Not(F))
	notOr = ele.Probabilities.Not(ele.Probabilities.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = ele.Probabilities.And(T, ele.Probabilities.Not(T))
	ass.Equal(t, sans, ele.Probabilities.Sans(T, T))

	sans = ele.Probabilities.And(T, ele.Probabilities.Not(F))
	ass.Equal(t, sans, ele.Probabilities.Sans(T, F))

	sans = ele.Probabilities.And(F, ele.Probabilities.Not(T))
	ass.Equal(t, sans, ele.Probabilities.Sans(F, T))

	sans = ele.Probabilities.And(F, ele.Probabilities.Not(F))
	ass.Equal(t, sans, ele.Probabilities.Sans(F, F))

	var xor = ele.Probabilities.Or(ele.Probabilities.Sans(T, T), ele.Probabilities.Sans(T, T))
	ass.Equal(t, xor, ele.Probabilities.Xor(T, T))

	xor = ele.Probabilities.Or(ele.Probabilities.Sans(T, F), ele.Probabilities.Sans(F, T))
	ass.Equal(t, xor, ele.Probabilities.Xor(T, F))

	xor = ele.Probabilities.Or(ele.Probabilities.Sans(F, T), ele.Probabilities.Sans(T, F))
	ass.Equal(t, xor, ele.Probabilities.Xor(F, T))

	xor = ele.Probabilities.Or(ele.Probabilities.Sans(F, F), ele.Probabilities.Sans(F, F))
	ass.Equal(t, xor, ele.Probabilities.Xor(F, F))
}
