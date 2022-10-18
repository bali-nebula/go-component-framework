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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetsWithStrings(t *testing.T) {
	var empty = []string{}
	var bazbar = collections.ListFromArray([]string{"baz", "bar"})
	var bazfoo = collections.ListFromArray([]string{"baz", "foo"})
	var baxbaz = collections.ListFromArray([]string{"bax", "baz"})
	var baxbez = collections.ListFromArray([]string{"bax", "bez"})
	var barbaz = collections.ListFromArray([]string{"bar", "baz"})
	var bar = collections.ListFromArray([]string{"bar"})
	var set = collections.Set[string]()                             // [ ]
	assert.True(t, set.IsEmpty())                                   // [ ]
	assert.Equal(t, 0, set.GetSize())                               // [ ]
	assert.False(t, set.ContainsItem("bax"))                        // [ ]
	assert.Equal(t, empty, set.AsArray())                           // [ ]
	var iterator = agents.Iterator[string](set)                     // [ ]
	assert.False(t, iterator.HasNext())                             // [ ]
	assert.False(t, iterator.HasPrevious())                         // [ ]
	iterator.ToStart()                                              // [ ]
	iterator.ToEnd()                                                // [ ]
	set.RemoveAll()                                                 // [ ]
	set.RemoveItem("foo")                                           // [ ]
	set.AddItem("foo")                                              // ["foo"]
	assert.False(t, set.IsEmpty())                                  // ["foo"]
	assert.Equal(t, 1, set.GetSize())                               // ["foo"]
	assert.Equal(t, "foo", string(set.GetItem(1)))                  // ["foo"]
	assert.True(t, set.ContainsItem("foo"))                         // ["foo"]
	assert.False(t, set.ContainsItem("bax"))                        // ["foo"]
	set.AddItems(bazbar)                                            // ["bar", "baz", "foo"]
	assert.Equal(t, 3, set.GetSize())                               // ["bar", "baz", "foo"]
	assert.Equal(t, "bar", string(set.GetItem(1)))                  // ["bar", "baz", "foo"]
	assert.Equal(t, bazfoo.AsArray(), set.GetItems(2, 3).AsArray()) // ["bar", "baz", "foo"]
	assert.Equal(t, bar.AsArray(), set.GetItems(1, 1).AsArray())    // ["bar", "baz", "foo"]
	iterator = agents.Iterator[string](set)                         // ["bar", "baz", "foo"]
	assert.True(t, iterator.HasNext())                              // ["bar", "baz", "foo"]
	assert.False(t, iterator.HasPrevious())                         // ["bar", "baz", "foo"]
	assert.Equal(t, "bar", string(iterator.GetNext()))              // ["bar", "baz", "foo"]
	assert.True(t, iterator.HasPrevious())                          // ["bar", "baz", "foo"]
	iterator.ToEnd()                                                // ["bar", "baz", "foo"]
	assert.False(t, iterator.HasNext())                             // ["bar", "baz", "foo"]
	assert.True(t, iterator.HasPrevious())                          // ["bar", "baz", "foo"]
	assert.Equal(t, "foo", string(iterator.GetPrevious()))          // ["bar", "baz", "foo"]
	assert.True(t, iterator.HasNext())                              // ["bar", "baz", "foo"]
	assert.True(t, set.ContainsItem("baz"))                         // ["bar", "baz", "foo"]
	assert.False(t, set.ContainsItem("bax"))                        // ["bar", "baz", "foo"]
	assert.True(t, set.ContainsAny(baxbaz))                         // ["bar", "baz", "foo"]
	assert.False(t, set.ContainsAny(baxbez))                        // ["bar", "baz", "foo"]
	assert.True(t, set.ContainsAll(barbaz))                         // ["bar", "baz", "foo"]
	assert.False(t, set.ContainsAll(baxbaz))                        // ["bar", "baz", "foo"]
	set.RemoveAll()                                                 // [ ]
	assert.True(t, set.IsEmpty())                                   // [ ]
	assert.Equal(t, 0, set.GetSize())                               // [ ]
}

func TestSetsWithIntegers(t *testing.T) {
	var list = collections.ListFromArray([]int{3, 1, 4, 5, 9, 2})
	var set = collections.Set[int]()         // [ ]
	set.AddItems(list)                       // [1,2,3,4,5,9]
	assert.False(t, set.IsEmpty())           // [1,2,3,4,5,9]
	assert.Equal(t, 6, set.GetSize())        // [1,2,3,4,5,9]
	assert.Equal(t, 1, int(set.GetItem(1)))  // [1,2,3,4,5,9]
	assert.Equal(t, 9, int(set.GetItem(-1))) // [1,2,3,4,5,9]
	set.RemoveItem(6)                        // [1,2,3,4,5,9]
	assert.Equal(t, 6, set.GetSize())        // [1,2,3,4,5,9]
	set.RemoveItem(3)                        // [1,2,3,4,5,9]
	assert.Equal(t, 5, set.GetSize())        // [1,2,3,4,5,9]
	assert.Equal(t, 4, int(set.GetItem(3)))  // [1,2,3,4,5,9]
}

func TestSetsWithSets(t *testing.T) {
	var list1 = collections.ListFromArray([]int{3, 1, 4, 5, 9, 2})
	var list2 = collections.ListFromArray([]int{7, 1, 4, 5, 9, 2})
	var set1 = collections.Set[int]()
	set1.AddItems(list1)
	var set2 = collections.Set[int]()
	set2.AddItems(list2)
	var set = collections.Set[abstractions.SetLike[int]]()
	var list3 = collections.ListFromArray([]abstractions.SetLike[int]{set1, set2})
	set.AddItems(list3)
	assert.False(t, set.IsEmpty())
	assert.Equal(t, 2, set.GetSize())
	assert.Equal(t, set1, set.GetItem(1))
	assert.Equal(t, set2, set.GetItem(-1))
	set.RemoveItem(set1)
	assert.Equal(t, 1, set.GetSize())
	set.RemoveAll()
	assert.Equal(t, 0, set.GetSize())
}

func TestSetsWithAnd(t *testing.T) {
	var list1 = collections.ListFromArray([]int{3, 1, 2})
	var list2 = collections.ListFromArray([]int{3, 2, 4})
	var list3 = collections.ListFromArray([]int{3, 2})
	var sets = collections.Sets[int]()
	var set1 = collections.Set[int]()
	set1.AddItems(list1)
	var set2 = collections.Set[int]()
	set2.AddItems(list2)
	var set3 = sets.And(set1, set2)
	var set4 = collections.Set[int]()
	set4.AddItems(list3)
	assert.True(t, agents.CompareValues(set3, set4))
}

func TestSetsWithSans(t *testing.T) {
	var list1 = collections.ListFromArray([]int{3, 1, 2})
	var list2 = collections.ListFromArray([]int{3, 2, 4})
	var list3 = collections.ListFromArray([]int{1})
	var sets = collections.Sets[int]()
	var set1 = collections.Set[int]()
	set1.AddItems(list1)
	var set2 = collections.Set[int]()
	set2.AddItems(list2)
	var set3 = sets.Sans(set1, set2)
	var set4 = collections.Set[int]()
	set4.AddItems(list3)
	assert.True(t, agents.CompareValues(set3, set4))
}

func TestSetsWithOr(t *testing.T) {
	var list1 = collections.ListFromArray([]int{3, 1, 5})
	var list2 = collections.ListFromArray([]int{6, 2, 4})
	var list3 = collections.ListFromArray([]int{1, 3, 5, 6, 2, 4})
	var sets = collections.Sets[int]()
	var set1 = collections.Set[int]()
	set1.AddItems(list1)
	var set2 = collections.Set[int]()
	set2.AddItems(list2)
	var set3 = sets.Or(set1, set2)
	assert.True(t, set3.ContainsAll(set1))
	assert.True(t, set3.ContainsAll(set2))
	var set4 = collections.Set[int]()
	set4.AddItems(list3)
	assert.True(t, agents.CompareValues(set3, set4))
}

func TestSetsWithXor(t *testing.T) {
	var list1 = collections.ListFromArray([]int{2, 3, 1, 5})
	var list2 = collections.ListFromArray([]int{6, 2, 5, 4})
	var list3 = collections.ListFromArray([]int{1, 3, 4, 6})
	var sets = collections.Sets[int]()
	var set1 = collections.Set[int]()
	set1.AddItems(list1)
	var set2 = collections.Set[int]()
	set2.AddItems(list2)
	var set3 = sets.Xor(set1, set2)
	var set4 = collections.Set[int]()
	set4.AddItems(list3)
	assert.True(t, agents.CompareValues(set3, set4))
}

func TestSetsWithEmptySets(t *testing.T) {
	var sets = collections.Sets[int]()
	var set1 = collections.Set[int]()
	var set2 = collections.Set[int]()
	var set3 = sets.And(set1, set2)
	var set4 = sets.Sans(set1, set2)
	var set5 = sets.Or(set1, set2)
	var set6 = sets.Xor(set1, set2)
	assert.True(t, agents.CompareValues(set3, set4))
	assert.True(t, agents.CompareValues(set4, set5))
	assert.True(t, agents.CompareValues(set5, set6))
	assert.True(t, agents.CompareValues(set6, set1))
}
