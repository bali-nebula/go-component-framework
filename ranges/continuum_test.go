/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package ranges_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ran "github.com/bali-nebula/go-component-framework/ranges"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestContinuaWithAngles(t *tes.T) {
	var s = ran.Continuum[ele.Angle](0, abs.LEFT, ele.Pi)
	ass.True(t, s.ContainsValue(ele.Angle(0)))
	ass.False(t, s.ContainsValue(ele.Pi))
	ass.True(t, s.ContainsValue(ele.Angle(2)))
	ass.False(t, s.ContainsValue(ele.Angle(4)))
	ass.Equal(t, ele.Angle(0), s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, ele.Pi, s.GetLast())
}

func TestContinuaWithPercentages(t *tes.T) {
	var s = ran.Continuum[ele.Percentage](-25, abs.INCLUSIVE, 100)
	ass.True(t, s.ContainsValue(-25))
	ass.True(t, s.ContainsValue(100))
	ass.False(t, s.ContainsValue(-30))
	ass.True(t, s.ContainsValue(50))
	ass.False(t, s.ContainsValue(150))
	ass.Equal(t, ele.Percentage(-25), s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, ele.Percentage(100), s.GetLast())
}

func TestContinuaWithProbabilities(t *tes.T) {
	var s = ran.Continuum[ele.Probability](0, abs.RIGHT, 1)
	ass.False(t, s.ContainsValue(0))
	ass.True(t, s.ContainsValue(1))
	ass.True(t, s.ContainsValue(0.5))
	ass.Equal(t, ele.Probability(0), s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, ele.Probability(1), s.GetLast())
}

func TestContinuaWithReals(t *tes.T) {
	var s = ran.Continuum[ran.Real](ran.Real(mat.Inf(-1)), abs.EXCLUSIVE, ran.Real(mat.Inf(1)))
	ass.True(t, s.ContainsValue(0))
	ass.Equal(t, ran.Real(mat.Inf(-1)), s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, ran.Real(mat.Inf(1)), s.GetLast())
}
