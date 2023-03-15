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
	mat "math"
	tes "testing"
)

func TestZeroAngles(t *tes.T) {
	var v = ele.AngleFromFloat(0)
	ass.Equal(t, 0.0, v.AsReal())
}

func TestPositiveAngles(t *tes.T) {
	var v = ele.AngleFromFloat(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsReal())
}

func TestNegativeAngles(t *tes.T) {
	var v1 = ele.AngleFromFloat(-mat.Pi)
	ass.Equal(t, mat.Pi, v1.AsReal())

	var v2 = ele.AngleFromFloat(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v2.AsReal())
}

func TestAnglesLibrary(t *tes.T) {
	var v0 = ele.AngleFromFloat(0)
	var v1 = ele.AngleFromFloat(mat.Pi * 0.25)
	var v2 = ele.AngleFromFloat(mat.Pi * 0.5)
	var v3 = ele.AngleFromFloat(mat.Pi * 0.75)
	var v4 = ele.AngleFromFloat(mat.Pi)
	var v5 = ele.AngleFromFloat(mat.Pi * 1.25)
	var v6 = ele.AngleFromFloat(mat.Pi * 1.5)
	var v7 = ele.AngleFromFloat(mat.Pi * 1.75)
	var v8 = ele.Tau

	ass.Equal(t, v4, ele.Angles.Inverse(v0))
	ass.Equal(t, v5, ele.Angles.Inverse(v1))
	ass.Equal(t, v6, ele.Angles.Inverse(v2))
	ass.Equal(t, v7, ele.Angles.Inverse(v3))
	ass.Equal(t, v0, ele.Angles.Inverse(v4))
	ass.Equal(t, v4, ele.Angles.Inverse(v8))

	ass.Equal(t, v1, ele.Angles.Sum(v0, v1))
	ass.Equal(t, v0, ele.Angles.Difference(v1, v1))
	ass.Equal(t, v3, ele.Angles.Sum(v1, v2))
	ass.Equal(t, v1, ele.Angles.Difference(v3, v2))
	ass.Equal(t, v5, ele.Angles.Sum(v2, v3))
	ass.Equal(t, v2, ele.Angles.Difference(v5, v3))
	ass.Equal(t, v7, ele.Angles.Sum(v3, v4))
	ass.Equal(t, v3, ele.Angles.Difference(v7, v4))
	ass.Equal(t, v1, ele.Angles.Sum(v8, v1))
	ass.Equal(t, v0, ele.Angles.Difference(v8, v8))

	ass.Equal(t, v3, ele.Angles.Scaled(v1, 3.0))
	ass.Equal(t, v0, ele.Angles.Scaled(v4, 2.0))
	ass.Equal(t, v4, ele.Angles.Scaled(v4, -1.0))
	ass.Equal(t, v0, ele.Angles.Scaled(v8, 1.0))

	ass.Equal(t, v0, ele.Angles.ArcCosine(ele.Angles.Cosine(v0)))
	ass.Equal(t, v1, ele.Angles.ArcCosine(ele.Angles.Cosine(v1)))
	ass.Equal(t, v2, ele.Angles.ArcCosine(ele.Angles.Cosine(v2)))
	ass.Equal(t, v3, ele.Angles.ArcCosine(ele.Angles.Cosine(v3)))
	ass.Equal(t, v4, ele.Angles.ArcCosine(ele.Angles.Cosine(v4)))
	ass.Equal(t, v0, ele.Angles.ArcCosine(ele.Angles.Cosine(v8)))

	ass.Equal(t, v0, ele.Angles.ArcSine(ele.Angles.Sine(v0)))
	ass.Equal(t, v1, ele.Angles.ArcSine(ele.Angles.Sine(v1)))
	ass.Equal(t, v2, ele.Angles.ArcSine(ele.Angles.Sine(v2)))
	ass.Equal(t, v6, ele.Angles.ArcSine(ele.Angles.Sine(v6)))
	ass.Equal(t, v7, ele.Angles.ArcSine(ele.Angles.Sine(v7)))
	ass.Equal(t, v0, ele.Angles.ArcSine(ele.Angles.Sine(v8)))

	ass.Equal(t, v0, ele.Angles.ArcTangent(ele.Angles.Cosine(v0), ele.Angles.Sine(v0)))
	ass.Equal(t, v1, ele.Angles.ArcTangent(ele.Angles.Cosine(v1), ele.Angles.Sine(v1)))
	ass.Equal(t, v2, ele.Angles.ArcTangent(ele.Angles.Cosine(v2), ele.Angles.Sine(v2)))
	ass.Equal(t, v3, ele.Angles.ArcTangent(ele.Angles.Cosine(v3), ele.Angles.Sine(v3)))
	ass.Equal(t, v4, ele.Angles.ArcTangent(ele.Angles.Cosine(v4), ele.Angles.Sine(v4)))
	ass.Equal(t, v5, ele.Angles.ArcTangent(ele.Angles.Cosine(v5), ele.Angles.Sine(v5)))
	ass.Equal(t, v0, ele.Angles.ArcTangent(ele.Angles.Cosine(v8), ele.Angles.Sine(v8)))
}
