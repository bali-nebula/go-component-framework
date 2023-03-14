/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	col "github.com/bali-nebula/go-component-framework/v1/collections"
	com "github.com/bali-nebula/go-component-framework/v1/components"
	ele "github.com/bali-nebula/go-component-framework/v1/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSets(t *tes.T) {
	var one = com.Component(ele.NumberFromComplex(1))
	var two = com.Component(ele.NumberFromComplex(2))
	var three = com.Component(ele.NumberFromComplex(3))
	var set = col.Set()
	ass.True(t, set.IsEmpty())
	ass.Equal(t, 0, set.GetSize())
	ass.False(t, set.ContainsValue(one))
	set.AddValue(one)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 1, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.False(t, set.ContainsValue(two))
	set.AddValue(two)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 2, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.True(t, set.ContainsValue(two))
	set.AddValue(three)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 3, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.True(t, set.ContainsValue(two))
	ass.True(t, set.ContainsValue(three))
	set.RemoveAll()
	ass.True(t, set.IsEmpty())
}
