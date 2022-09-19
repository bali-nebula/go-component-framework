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

func TestStringTags(t *testing.T) {
	for i := 1; i < 33; i++ {
		var t1 = elements.TagOfSize(i)
		assert.Equal(t, len(t1)-1, int(math.Ceil(float64(i)*8.0/5.0)))
		var s1 = t1.AsString()
		var t2, ok = elements.TagFromString(s1)
		assert.True(t, ok)
		assert.Equal(t, t1, t2)
		var s2 = t2.AsString()
		assert.Equal(t, s1, s2)
	}
}

func TestBadTags(t *testing.T) {
	var v, ok = elements.TagFromString("#")
	assert.False(t, ok)
	assert.Equal(t, "", string(v))

	v, ok = elements.TagFromString("BAD")
	assert.False(t, ok)
	assert.Equal(t, "", string(v))
}
