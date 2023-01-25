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
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestLists(t *tes.T) {
	var context = com.Context()
	var one = com.ComponentWithContext(ele.NumberFromComplex(1), context)
	var two = com.Component(ele.NumberFromComplex(2))
	var three = com.Component(ele.NumberFromComplex(3))
	var list = col.List()
	list.SortValues()
	ass.True(t, list.IsEmpty())
	ass.Equal(t, 0, list.GetSize())
	ass.False(t, list.ContainsValue(one))
	list.AddValue(one)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 1, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.False(t, list.ContainsValue(two))
	list.AddValue(two)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 2, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.True(t, list.ContainsValue(two))
	list.AddValue(three)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 3, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.True(t, list.ContainsValue(two))
	ass.True(t, list.ContainsValue(three))
	list.RemoveAll()
	ass.True(t, list.IsEmpty())
}
