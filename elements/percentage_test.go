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

func TestZeroPercentages(t *testing.T) {
	var v = elements.Percentage(0.0)
	assert.Equal(t, 0.0, float64(v))
	assert.Equal(t, 0.0, v.AsReal())
}

func TestPositivePercentages(t *testing.T) {
	var v = elements.Percentage(25)
	assert.Equal(t, 25, int(v))
	assert.Equal(t, 0.25, v.AsReal())
}

func TestNegativePercentages(t *testing.T) {
	var v = elements.Percentage(-75)
	assert.Equal(t, -75, int(v))
	assert.Equal(t, -0.75, v.AsReal())
}

func TestStringPercentages(t *testing.T) {
	var v, ok = elements.PercentageFromString("50%")
	assert.True(t, ok)
	assert.Equal(t, 50, int(v))
	assert.Equal(t, "50%", v.AsString())
	assert.Equal(t, 0.5, v.AsReal())

	v, ok = elements.PercentageFromString("100")
	assert.False(t, ok)
	assert.Equal(t, 0, int(v))
}

func TestPercentagesLibrary(t *testing.T) {
	var v0 = elements.Percentage(0)
	var v1 = elements.Percentage(0.25)
	var v2 = elements.Percentage(0.5)
	var v3 = elements.Percentage(1)

	assert.Equal(t, v0, elements.Percentages.Inverse(v0))
	assert.Equal(t, -v1, elements.Percentages.Inverse(v1))
	assert.Equal(t, -v2, elements.Percentages.Inverse(v2))
	assert.Equal(t, -v3, elements.Percentages.Inverse(v3))

	assert.Equal(t, v1, elements.Percentages.Sum(v0, v1))
	assert.Equal(t, v0, elements.Percentages.Difference(v1, v1))
	assert.Equal(t, v2, elements.Percentages.Sum(v1, v1))
	assert.Equal(t, v1, elements.Percentages.Difference(v2, v1))
	assert.Equal(t, v3, elements.Percentages.Sum(v2, v2))
	assert.Equal(t, v2, elements.Percentages.Difference(v3, v2))

	assert.Equal(t, v0, elements.Percentages.Scaled(v0, 3.0))
	assert.Equal(t, v3, elements.Percentages.Scaled(v1, 4.0))
	assert.Equal(t, -v2, elements.Percentages.Scaled(v2, -1.0))
}
