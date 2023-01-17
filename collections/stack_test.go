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

func TestStacks(t *tes.T) {
	var one = com.Component(ele.Number(1))
	var two = com.Component(ele.Number(2))
	var three = com.Component(ele.Number(3))
	var stack = col.Stack()
	ass.True(t, stack.IsEmpty())
	ass.Equal(t, 0, stack.GetSize())
	stack.AddValue(one)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 1, stack.GetSize())
	stack.AddValue(two)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 2, stack.GetSize())
	stack.AddValue(three)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 3, stack.GetSize())
	ass.Equal(t, three, stack.RemoveTop())
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 2, stack.GetSize())
	ass.Equal(t, two, stack.RemoveTop())
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 1, stack.GetSize())
	ass.Equal(t, one, stack.RemoveTop())
	ass.True(t, stack.IsEmpty())
	ass.Equal(t, 0, stack.GetSize())
	stack.RemoveAll()
	ass.True(t, stack.IsEmpty())
}
