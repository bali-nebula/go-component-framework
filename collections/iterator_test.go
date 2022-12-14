/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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

func TestComponentIterators(t *tes.T) {
	var foo = com.Component(ele.Symbol("foo"))
	var bar = com.Component(ele.Symbol("bar"))
	var baz = com.Component(ele.Symbol("baz"))
	var list = col.List()
	list.AddValue(foo)
	list.AddValue(bar)
	list.AddValue(baz)
	var iterator = col.ComponentIterator(list)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, foo, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, foo, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, baz, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, baz, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, foo, iterator.GetNext())
	ass.Equal(t, bar, iterator.GetNext())
}
