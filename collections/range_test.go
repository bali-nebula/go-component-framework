/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangesWithIntegers(t *testing.T) {
	var s = collections.Range[int](3, "..", 7)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 5, s.GetSize())
	assert.Equal(t, 5, s.GetItem(3))
	assert.Equal(t, 0, s.GetIndex(2))
	assert.Equal(t, 1, s.GetIndex(3))
	assert.Equal(t, 3, s.GetIndex(5))
	assert.Equal(t, 5, s.GetIndex(7))
	assert.Equal(t, 0, s.GetIndex(8))
	assert.Equal(t, []int{3, 4, 5, 6, 7}, s.AsArray())
	var iterator = agents.Iterator[int](s)
	assert.Equal(t, 3, iterator.GetNext())
	iterator.ToEnd()
	assert.Equal(t, 7, iterator.GetPrevious())

	s = collections.Range[int](3, "<..", 7)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 4, s.GetSize())
	assert.Equal(t, 4, s.GetItem(1))
	assert.Equal(t, 0, s.GetIndex(3))
	assert.Equal(t, 1, s.GetIndex(4))
	assert.Equal(t, 2, s.GetIndex(5))
	assert.Equal(t, 4, s.GetIndex(7))
	assert.Equal(t, 0, s.GetIndex(8))
	assert.Equal(t, []int{4, 5, 6, 7}, s.AsArray())

	s = collections.Range[int](3, "<..<", 7)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 3, s.GetSize())
	assert.Equal(t, 4, s.GetItem(1))
	assert.Equal(t, 0, s.GetIndex(3))
	assert.Equal(t, 1, s.GetIndex(4))
	assert.Equal(t, 3, s.GetIndex(6))
	assert.Equal(t, 0, s.GetIndex(7))
	assert.Equal(t, []int{4, 5, 6}, s.AsArray())

	s = collections.Range[int](3, "..<", 7)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 4, s.GetSize())
	assert.Equal(t, 3, s.GetItem(1))
	assert.Equal(t, 1, s.GetIndex(3))
	assert.Equal(t, 4, s.GetIndex(6))
	assert.Equal(t, 0, s.GetIndex(7))
	assert.Equal(t, []int{3, 4, 5, 6}, s.AsArray())
}

func TestRangesWithRunes(t *testing.T) {
	var s = collections.Range[rune]('a', "..", 'z')
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 26, s.GetSize())
	assert.Equal(t, 'd', s.GetItem(4))
	assert.Equal(t, 13, s.GetIndex('m'))

	var a = collections.List[rune]()
	var items = s.AsArray()
	a.AddItems(s)
	assert.Equal(t, items, a.AsArray())
}
