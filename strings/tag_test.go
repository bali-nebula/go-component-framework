/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	sts "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestStringTags(t *tes.T) {
	for i := 1; i < 33; i++ {
		var t1 = sts.TagOfSize(i)
		ass.Equal(t, len(t1.AsString()), int(mat.Ceil(float64(i)*8.0/5.0)))
		var s1 = t1.AsString()
		var t2 = sts.TagFromString(s1)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsString()
		ass.Equal(t, s1, s2)
	}
}
