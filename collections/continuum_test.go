/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	col "github.com/craterdog-bali/go-component-framework/collections"
	ele "github.com/craterdog-bali/go-component-framework/elements"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestContinuaWithAngles(t *tes.T) {
	var s = col.Continuum(0, abs.LEFT, ele.Pi)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, -1, s.GetSize())
	ass.Equal(t, []ele.Angle{}, s.AsArray())
}

func TestContinuaWithPercentages(t *tes.T) {
	var s = col.Continuum(0, abs.INCLUSIVE, ele.Percentage(100))
	ass.False(t, s.IsEmpty())
	ass.Equal(t, -1, s.GetSize())
	ass.Equal(t, []ele.Percentage{}, s.AsArray())
}

func TestContinuaWithProbabilities(t *tes.T) {
	var s = col.Continuum(0, abs.RIGHT, ele.Probability(1))
	ass.False(t, s.IsEmpty())
	ass.Equal(t, -1, s.GetSize())
	ass.Equal(t, []ele.Probability{}, s.AsArray())
}

func TestContinuaWithReals(t *tes.T) {
	var s = col.Continuum(col.Real(mat.Inf(-1)), abs.EXCLUSIVE, col.Real(mat.Inf(1)))
	ass.False(t, s.IsEmpty())
	ass.Equal(t, -1, s.GetSize())
	ass.Equal(t, []col.Real{}, s.AsArray())
}
