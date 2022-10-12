/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIteratorsWithLists(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5}
	var list = collections.List[int]()
	list.AddItems(array)
	var iterator = agents.Iterator[int](list)
	assert.False(t, iterator.HasPrevious())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.GetNext())
	assert.True(t, iterator.HasPrevious())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.GetPrevious())
	iterator.ToSlot(2)
	assert.True(t, iterator.HasPrevious())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 3, iterator.GetNext())
	iterator.ToEnd()
	assert.True(t, iterator.HasPrevious())
	assert.False(t, iterator.HasNext())
	assert.Equal(t, 5, iterator.GetPrevious())
	iterator.ToStart()
	assert.False(t, iterator.HasPrevious())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.GetNext())
}
