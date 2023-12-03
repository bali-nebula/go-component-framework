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

var Probability = ele.Probability()

func TestBooleanProbabilities(t *tes.T) {
	var v1 = Probability.FromBool(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = Probability.FromBool(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = Probability.FromFloat(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = Probability.FromFloat(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbabilities(t *tes.T) {
	Probability.Random()
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = Probability.FromFloat(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = Probability.FromFloat(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = Probability.FromFloat(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = Probability.FromFloat(0.75)
	var F = Probability.FromFloat(0.25)

	var andNot = Probability.And(Probability.Not(T), Probability.Not(T))
	var notOr = Probability.Not(Probability.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = Probability.And(Probability.Not(T), Probability.Not(F))
	notOr = Probability.Not(Probability.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = Probability.And(Probability.Not(F), Probability.Not(T))
	notOr = Probability.Not(Probability.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = Probability.And(Probability.Not(F), Probability.Not(F))
	notOr = Probability.Not(Probability.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = Probability.And(T, Probability.Not(T))
	ass.Equal(t, sans, Probability.Sans(T, T))

	sans = Probability.And(T, Probability.Not(F))
	ass.Equal(t, sans, Probability.Sans(T, F))

	sans = Probability.And(F, Probability.Not(T))
	ass.Equal(t, sans, Probability.Sans(F, T))

	sans = Probability.And(F, Probability.Not(F))
	ass.Equal(t, sans, Probability.Sans(F, F))

	var xor = Probability.FromFloat(Probability.Sans(T, T).AsFloat() + Probability.Sans(T, T).AsFloat())
	ass.Equal(t, xor, Probability.Xor(T, T))

	xor = Probability.FromFloat(Probability.Sans(T, F).AsFloat() + Probability.Sans(F, T).AsFloat())
	ass.Equal(t, xor, Probability.Xor(T, F))

	xor = Probability.FromFloat(Probability.Sans(F, T).AsFloat() + Probability.Sans(T, F).AsFloat())
	ass.Equal(t, xor, Probability.Xor(F, T))

	xor = Probability.FromFloat(Probability.Sans(F, F).AsFloat() + Probability.Sans(F, F).AsFloat())
	ass.Equal(t, xor, Probability.Xor(F, F))
}
