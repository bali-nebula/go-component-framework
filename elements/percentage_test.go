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
	ele "github.com/bali-nebula/go-component-framework/elements"
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
