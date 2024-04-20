/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element_test

import (
	ele "github.com/bali-nebula/go-component-framework/v3/element"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestZeroAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(0)
	ass.Equal(t, 0.0, v.AsFloat())

	v = Angle.MakeFromString("~0")
	ass.Equal(t, "~0", v.AsString())
}

func TestPositiveAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = Angle.MakeFromString("~π")
	ass.Equal(t, "~π", v.AsString())
}

func TestNegativeAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = Angle.MakeFromFloat(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var Angle = ele.Angle()
	var v0 = Angle.Zero()
	var v1 = Angle.MakeFromFloat(mat.Pi * 0.25)
	var v2 = Angle.MakeFromFloat(mat.Pi * 0.5)
	var v3 = Angle.MakeFromFloat(mat.Pi * 0.75)
	var v4 = Angle.Pi()
	var v5 = Angle.MakeFromFloat(mat.Pi * 1.25)
	var v6 = Angle.MakeFromFloat(mat.Pi * 1.5)
	var v7 = Angle.MakeFromFloat(mat.Pi * 1.75)
	var v8 = Angle.Tau()

	ass.Equal(t, v4, Angle.Inverse(v0))
	ass.Equal(t, v5, Angle.Inverse(v1))
	ass.Equal(t, v6, Angle.Inverse(v2))
	ass.Equal(t, v7, Angle.Inverse(v3))
	ass.Equal(t, v0, Angle.Inverse(v4))
	ass.Equal(t, v4, Angle.Inverse(v8))

	ass.Equal(t, v1, Angle.Sum(v0, v1))
	ass.Equal(t, v0, Angle.Difference(v1, v1))
	ass.Equal(t, v3, Angle.Sum(v1, v2))
	ass.Equal(t, v1, Angle.Difference(v3, v2))
	ass.Equal(t, v5, Angle.Sum(v2, v3))
	ass.Equal(t, v2, Angle.Difference(v5, v3))
	ass.Equal(t, v7, Angle.Sum(v3, v4))
	ass.Equal(t, v3, Angle.Difference(v7, v4))
	ass.Equal(t, v1, Angle.Sum(v8, v1))
	ass.Equal(t, v0, Angle.Difference(v8, v8))

	ass.Equal(t, v3, Angle.Scaled(v1, 3.0))
	ass.Equal(t, v0, Angle.Scaled(v4, 2.0))
	ass.Equal(t, v4, Angle.Scaled(v4, -1.0))
	ass.Equal(t, v0, Angle.Scaled(v8, 1.0))

	ass.Equal(t, v0, Angle.ArcCosine(Angle.Cosine(v0)))
	ass.Equal(t, v1, Angle.ArcCosine(Angle.Cosine(v1)))
	ass.Equal(t, v2, Angle.ArcCosine(Angle.Cosine(v2)))
	ass.Equal(t, v3, Angle.ArcCosine(Angle.Cosine(v3)))
	ass.Equal(t, v4, Angle.ArcCosine(Angle.Cosine(v4)))
	ass.Equal(t, v0, Angle.ArcCosine(Angle.Cosine(v8)))

	ass.Equal(t, v0, Angle.ArcSine(Angle.Sine(v0)))
	ass.Equal(t, v1, Angle.ArcSine(Angle.Sine(v1)))
	ass.Equal(t, v2, Angle.ArcSine(Angle.Sine(v2)))
	ass.Equal(t, v6, Angle.ArcSine(Angle.Sine(v6)))
	ass.Equal(t, v7, Angle.ArcSine(Angle.Sine(v7)))
	ass.Equal(t, v0, Angle.ArcSine(Angle.Sine(v8)))

	ass.Equal(t, v0, Angle.ArcTangent(Angle.Cosine(v0), Angle.Sine(v0)))
	ass.Equal(t, v1, Angle.ArcTangent(Angle.Cosine(v1), Angle.Sine(v1)))
	ass.Equal(t, v2, Angle.ArcTangent(Angle.Cosine(v2), Angle.Sine(v2)))
	ass.Equal(t, v3, Angle.ArcTangent(Angle.Cosine(v3), Angle.Sine(v3)))
	ass.Equal(t, v4, Angle.ArcTangent(Angle.Cosine(v4), Angle.Sine(v4)))
	ass.Equal(t, v5, Angle.ArcTangent(Angle.Cosine(v5), Angle.Sine(v5)))
	ass.Equal(t, v0, Angle.ArcTangent(Angle.Cosine(v8), Angle.Sine(v8)))
}
