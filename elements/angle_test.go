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
	"math"
	"testing"
)

func TestZeroAngles(t *testing.T) {
	var v = elements.AngleFromFloat(0)
	assert.Equal(t, 0.0, float64(v))
	assert.Equal(t, "~0", v.AsString())
	assert.Equal(t, 0.0, v.AsReal())
}

func TestPositiveAngles(t *testing.T) {
	var v = elements.AngleFromFloat(math.Pi)
	assert.Equal(t, math.Pi, float64(v))
	assert.Equal(t, "~π", v.AsString())
	assert.Equal(t, math.Pi, v.AsReal())
}

func TestNegativeAngles(t *testing.T) {
	var v1 = elements.AngleFromFloat(-math.Pi)
	assert.Equal(t, math.Pi, float64(v1))
	assert.Equal(t, "~π", v1.AsString())
	assert.Equal(t, math.Pi, v1.AsReal())

	var v2 = elements.AngleFromFloat(-math.Pi / 2.0)
	assert.Equal(t, 1.5*math.Pi, float64(v2))
	assert.Equal(t, "~4.71238898038469", v2.AsString())
	assert.Equal(t, 1.5*math.Pi, v2.AsReal())
}

func TestStringAngles(t *testing.T) {
	var v, ok = elements.AngleFromString("~π")
	assert.True(t, ok)
	assert.Equal(t, math.Pi, float64(v))
	assert.Equal(t, "~π", v.AsString())
	assert.Equal(t, math.Pi, v.AsReal())

	v, ok = elements.AngleFromString("bad")
	assert.False(t, ok)
	assert.Equal(t, 0, int(v))
}

func TestAnglesLibrary(t *testing.T) {
	var v0 = elements.AngleFromFloat(0)
	var v1 = elements.AngleFromFloat(math.Pi * 0.25)
	var v2 = elements.AngleFromFloat(math.Pi * 0.5)
	var v3 = elements.AngleFromFloat(math.Pi * 0.75)
	var v4 = elements.AngleFromFloat(math.Pi)
	var v5 = elements.AngleFromFloat(math.Pi * 1.25)
	var v6 = elements.AngleFromFloat(math.Pi * 1.5)
	var v7 = elements.AngleFromFloat(math.Pi * 1.75)

	assert.Equal(t, v4, elements.Angles.Inverse(v0))
	assert.Equal(t, v5, elements.Angles.Inverse(v1))
	assert.Equal(t, v6, elements.Angles.Inverse(v2))
	assert.Equal(t, v7, elements.Angles.Inverse(v3))
	assert.Equal(t, v0, elements.Angles.Inverse(v4))

	assert.Equal(t, v1, elements.Angles.Sum(v0, v1))
	assert.Equal(t, v0, elements.Angles.Difference(v1, v1))
	assert.Equal(t, v3, elements.Angles.Sum(v1, v2))
	assert.Equal(t, v1, elements.Angles.Difference(v3, v2))
	assert.Equal(t, v5, elements.Angles.Sum(v2, v3))
	assert.Equal(t, v2, elements.Angles.Difference(v5, v3))
	assert.Equal(t, v7, elements.Angles.Sum(v3, v4))
	assert.Equal(t, v3, elements.Angles.Difference(v7, v4))

	assert.Equal(t, v3, elements.Angles.Scaled(v1, 3.0))
	assert.Equal(t, v0, elements.Angles.Scaled(v4, 2.0))
	assert.Equal(t, v4, elements.Angles.Scaled(v4, -1.0))

	assert.Equal(t, v0, elements.Angles.ArcCosine(elements.Angles.Cosine(v0)))
	assert.Equal(t, v1, elements.Angles.ArcCosine(elements.Angles.Cosine(v1)))
	assert.Equal(t, v2, elements.Angles.ArcCosine(elements.Angles.Cosine(v2)))
	assert.Equal(t, v3, elements.Angles.ArcCosine(elements.Angles.Cosine(v3)))
	assert.Equal(t, v4, elements.Angles.ArcCosine(elements.Angles.Cosine(v4)))

	assert.Equal(t, v0, elements.Angles.ArcSine(elements.Angles.Sine(v0)))
	assert.Equal(t, v1, elements.Angles.ArcSine(elements.Angles.Sine(v1)))
	assert.Equal(t, v2, elements.Angles.ArcSine(elements.Angles.Sine(v2)))
	assert.Equal(t, v6, elements.Angles.ArcSine(elements.Angles.Sine(v6)))
	assert.Equal(t, v7, elements.Angles.ArcSine(elements.Angles.Sine(v7)))

	assert.Equal(t, v0, elements.Angles.ArcTangent(elements.Angles.Cosine(v0), elements.Angles.Sine(v0)))
	assert.Equal(t, v1, elements.Angles.ArcTangent(elements.Angles.Cosine(v1), elements.Angles.Sine(v1)))
	assert.Equal(t, v2, elements.Angles.ArcTangent(elements.Angles.Cosine(v2), elements.Angles.Sine(v2)))
	assert.Equal(t, v3, elements.Angles.ArcTangent(elements.Angles.Cosine(v3), elements.Angles.Sine(v3)))
	assert.Equal(t, v4, elements.Angles.ArcTangent(elements.Angles.Cosine(v4), elements.Angles.Sine(v4)))
	assert.Equal(t, v5, elements.Angles.ArcTangent(elements.Angles.Cosine(v5), elements.Angles.Sine(v5)))
}
