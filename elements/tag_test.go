/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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
	mat "math"
	tes "testing"
)

func TestStringTags(t *tes.T) {
	for i := 1; i < 33; i++ {
		var t1 = ele.TagOfSize(i)
		ass.Equal(t, len(t1), int(mat.Ceil(float64(i)*8.0/5.0)))
		var s1 = string(t1)
		var t2 = ele.Tag(s1)
		ass.Equal(t, t1, t2)
		var s2 = string(t2)
		ass.Equal(t, s1, s2)
	}
}
