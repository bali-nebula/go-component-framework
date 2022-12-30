/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package temporary_test

import (
	tem "github.com/bali-nebula/go-component-framework/temporary"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestContinuaWithAngles(t *tes.T) {
	var s = tem.Continuum[ele.Angle](0, tem.LEFT, ele.Pi)
	ass.Equal(t, ele.Angle(0), s.GetFirst())
	ass.Equal(t, tem.LEFT, s.GetExtent())
	ass.Equal(t, ele.Pi, s.GetLast())
}

func TestContinuaWithPercentages(t *tes.T) {
	var s = tem.Continuum[ele.Percentage](-25, tem.INCLUSIVE, 100)
	ass.Equal(t, ele.Percentage(-25), s.GetFirst())
	ass.Equal(t, tem.INCLUSIVE, s.GetExtent())
	ass.Equal(t, ele.Percentage(100), s.GetLast())
}

func TestContinuaWithProbabilities(t *tes.T) {
	var s = tem.Continuum[ele.Probability](0, tem.RIGHT, 1)
	ass.Equal(t, ele.Probability(0), s.GetFirst())
	ass.Equal(t, tem.RIGHT, s.GetExtent())
	ass.Equal(t, ele.Probability(1), s.GetLast())
}

func TestContinuaWithReals(t *tes.T) {
	var s = tem.Continuum[tem.Real](tem.Real(mat.Inf(-1)), tem.EXCLUSIVE, tem.Real(mat.Inf(1)))
	ass.Equal(t, tem.Real(mat.Inf(-1)), s.GetFirst())
	ass.Equal(t, tem.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, tem.Real(mat.Inf(1)), s.GetLast())
}
