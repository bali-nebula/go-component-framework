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
	mat "math"
	tes "testing"
)

func TestStringTags(t *tes.T) {
	for i := 1; i < 33; i++ {
		var t1 = ele.TagOfSize(i)
		ass.Equal(t, len(t1)-1, int(mat.Ceil(float64(i)*8.0/5.0)))
		var s1 = t1.AsString()
		var t2, ok = ele.TagFromString(s1)
		ass.True(t, ok)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsString()
		ass.Equal(t, s1, s2)
	}
}

func TestBadTags(t *tes.T) {
	var v, ok = ele.TagFromString("#")
	ass.False(t, ok)
	ass.Equal(t, "", string(v))

	v, ok = ele.TagFromString("BAD")
	ass.False(t, ok)
	ass.Equal(t, "", string(v))
}
