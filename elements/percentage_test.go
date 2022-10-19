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
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestZeroPercentages(t *tes.T) {
	var v = ele.Percentage(0.0)
	ass.Equal(t, 0.0, float64(v))
	ass.Equal(t, 0.0, v.AsReal())
}

func TestPositivePercentages(t *tes.T) {
	var v = ele.Percentage(25)
	ass.Equal(t, 25, int(v))
	ass.Equal(t, 0.25, v.AsReal())
}

func TestNegativePercentages(t *tes.T) {
	var v = ele.Percentage(-75)
	ass.Equal(t, -75, int(v))
	ass.Equal(t, -0.75, v.AsReal())
}

func TestStringPercentages(t *tes.T) {
	var v, ok = ele.PercentageFromString("50%")
	ass.True(t, ok)
	ass.Equal(t, 50, int(v))
	ass.Equal(t, "50%", v.AsString())
	ass.Equal(t, 0.5, v.AsReal())

	v, ok = ele.PercentageFromString("100")
	ass.False(t, ok)
	ass.Equal(t, 0, int(v))
}

func TestPercentagesLibrary(t *tes.T) {
	var v0 = ele.Percentage(0)
	var v1 = ele.Percentage(0.25)
	var v2 = ele.Percentage(0.5)
	var v3 = ele.Percentage(1)

	ass.Equal(t, v0, ele.Percentages.Inverse(v0))
	ass.Equal(t, -v1, ele.Percentages.Inverse(v1))
	ass.Equal(t, -v2, ele.Percentages.Inverse(v2))
	ass.Equal(t, -v3, ele.Percentages.Inverse(v3))

	ass.Equal(t, v1, ele.Percentages.Sum(v0, v1))
	ass.Equal(t, v0, ele.Percentages.Difference(v1, v1))
	ass.Equal(t, v2, ele.Percentages.Sum(v1, v1))
	ass.Equal(t, v1, ele.Percentages.Difference(v2, v1))
	ass.Equal(t, v3, ele.Percentages.Sum(v2, v2))
	ass.Equal(t, v2, ele.Percentages.Difference(v3, v2))

	ass.Equal(t, v0, ele.Percentages.Scaled(v0, 3.0))
	ass.Equal(t, v3, ele.Percentages.Scaled(v1, 4.0))
	ass.Equal(t, -v2, ele.Percentages.Scaled(v2, -1.0))
}
