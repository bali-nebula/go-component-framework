/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	tes "testing"
)

func TestContinuaWithAngles(t *tes.T) {
	var s = ran.Continuum[abs.AngleLike](ele.AngleFromFloat(0), abs.LEFT, ele.Pi)
	ass.True(t, s.ContainsValue(ele.AngleFromFloat(0)))
	ass.False(t, s.ContainsValue(ele.Pi))
	ass.True(t, s.ContainsValue(ele.AngleFromFloat(2)))
	ass.False(t, s.ContainsValue(ele.AngleFromFloat(4)))
	ass.Equal(t, ele.AngleFromFloat(0), s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, ele.Pi, s.GetLast())
}

func TestContinuaWithPercentages(t *tes.T) {
	var pMinus30 = ele.PercentageFromInt(-30)
	var pMinus25 = ele.PercentageFromInt(-25)
	var p100 = ele.PercentageFromInt(100)
	var p50 = ele.PercentageFromInt(50)
	var p150 = ele.PercentageFromInt(150)
	var s = ran.Continuum[abs.PercentageLike](pMinus25, abs.INCLUSIVE, p100)
	ass.True(t, s.ContainsValue(pMinus25))
	ass.True(t, s.ContainsValue(p100))
	ass.False(t, s.ContainsValue(pMinus30))
	ass.True(t, s.ContainsValue(p50))
	ass.False(t, s.ContainsValue(p150))
	ass.Equal(t, pMinus25, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, p100, s.GetLast())
}

func TestContinuaWithProbabilities(t *tes.T) {
	var p0 = ele.ProbabilityFromFloat(0)
	var pHalf = ele.ProbabilityFromFloat(0.5)
	var p1 = ele.ProbabilityFromFloat(1)
	var s = ran.Continuum[abs.ProbabilityLike](p0, abs.RIGHT, p1)
	ass.False(t, s.ContainsValue(p0))
	ass.True(t, s.ContainsValue(p1))
	ass.True(t, s.ContainsValue(pHalf))
	ass.Equal(t, p0, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, p1, s.GetLast())
}

func TestContinuaWithReals(t *tes.T) {
	var minimum = ele.MinimumReal()
	var zero = ele.Real(0)
	var maximum = ele.MaximumReal()
	var s = ran.Continuum[abs.RealLike](minimum, abs.EXCLUSIVE, maximum)
	ass.True(t, s.ContainsValue(zero))
	ass.Equal(t, minimum, s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, maximum, s.GetLast())
}
