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

func TestListsWithStrings(t *testing.T) {
	var list = collections.List[string]()                        // [ ]
	assert.True(t, list.IsEmpty())                               // [ ]
	assert.Equal(t, 0, list.GetSize())                           // [ ]
	assert.False(t, list.ContainsItem("bax"))                    // [ ]
	assert.Equal(t, []string{}, list.AsArray())                  // [ ]
	var iterator = agents.Iterator[string](list)                 // [ ]
	assert.False(t, iterator.HasNext())                          // [ ]
	assert.False(t, iterator.HasPrevious())                      // [ ]
	iterator.ToStart()                                           // [ ]
	iterator.ToEnd()                                             // [ ]
	list.ShuffleItems()                                          // [ ]
	list.SortItems()                                             // [ ]
	list.RemoveAll()                                             // [ ]
	list.AddItem("foo")                                          // ["foo"]
	assert.False(t, list.IsEmpty())                              // ["foo"]
	assert.Equal(t, 1, list.GetSize())                           // ["foo"]
	assert.Equal(t, "foo", string(list.GetItem(1)))              // ["foo"]
	list.AddItems([]string{"bar", "baz"})                        // ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetSize())                           // ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(list.GetItem(1)))              // ["foo", "bar", "baz"]
	assert.Equal(t, []string{"bar", "baz"}, list.GetItems(2, 3)) // ["foo", "bar", "baz"]
	assert.Equal(t, []string{"foo"}, list.GetItems(1, 1))        // ["foo", "bar", "baz"]
	iterator = agents.Iterator[string](list)                     // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasNext())                           // ["foo", "bar", "baz"]
	assert.False(t, iterator.HasPrevious())                      // ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(iterator.GetNext()))           // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasPrevious())                       // ["foo", "bar", "baz"]
	iterator.ToEnd()                                             // ["foo", "bar", "baz"]
	assert.False(t, iterator.HasNext())                          // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasPrevious())                       // ["foo", "bar", "baz"]
	assert.Equal(t, "baz", string(iterator.GetPrevious()))       // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasNext())                           // ["foo", "bar", "baz"]
	list.ShuffleItems()                                          // [ ?, ?, ? ]
	list.RemoveAll()                                             // [ ]
	assert.True(t, list.IsEmpty())                               // [ ]
	assert.Equal(t, 0, list.GetSize())                           // [ ]
	list.InsertItem(0, "baz")                                    // ["baz"]
	assert.Equal(t, 1, list.GetSize())                           // ["baz"]
	assert.Equal(t, "baz", string(list.GetItem(-1)))             // ["baz"]
	list.InsertItems(0, []string{"foo", "bar"})                  // ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetSize())                           // ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(list.GetItem(-3)))             // ["foo", "bar", "baz"]
	assert.Equal(t, "bar", string(list.GetItem(-2)))             // ["foo", "bar", "baz"]
	assert.Equal(t, "baz", string(list.GetItem(-1)))             // ["foo", "bar", "baz"]
	list.ReverseItems()                                          // ["baz", "bar", "foo"]
	assert.Equal(t, "foo", string(list.GetItem(-1)))             // ["baz", "bar", "foo"]
	assert.Equal(t, "bar", string(list.GetItem(-2)))             // ["baz", "bar", "foo"]
	assert.Equal(t, "baz", string(list.GetItem(-3)))             // ["baz", "bar", "foo"]
	list.ReverseItems()                                          // ["foo", "bar", "baz"]
	assert.Equal(t, 0, list.GetIndex("foz"))                     // ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetIndex("baz"))                     // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsItem("baz"))                     // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsItem("bax"))                    // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsAny([]string{"bax", "baz"}))     // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsAny([]string{"bax", "bez"}))    // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsAll([]string{"bar", "baz"}))     // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsAll([]string{"bax", "baz"}))    // ["foo", "bar", "baz"]
	list.SetItem(3, "bax")                                       // ["foo", "bar", "bax"]

	list.InsertItems(3, []string{"baz"})                  // ["foo", "bar", "bax", "baz"]
	assert.Equal(t, 4, list.GetSize())                    // ["foo", "bar", "bax", "baz"]
	assert.Equal(t, "baz", string(list.GetItem(4)))       // ["foo", "bar", "bax", "baz"]
	list.InsertItem(4, "bar")                             // ["foo", "bar", "bax", "baz", "bar"]
	assert.Equal(t, 5, list.GetSize())                    // ["foo", "bar", "bax", "baz", "bar"]
	assert.Equal(t, "bar", string(list.GetItem(5)))       // ["foo", "bar", "bax", "baz", "bar"]
	list.InsertItem(2, "foo")                             // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, 6, list.GetSize())                    // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "bar", string(list.GetItem(2)))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "foo", string(list.GetItem(3)))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "bax", string(list.GetItem(4)))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, []string{"bar"}, list.GetItems(6, 6)) // ["foo", "bar", "foo", "bax", "baz", "bar"]
	list.InsertItems(5, []string{"baz"})                  // ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, 7, list.GetSize())                    // ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "bax", string(list.GetItem(4)))       // ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "baz", string(list.GetItem(5)))       // ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "baz", string(list.GetItem(6)))       // ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, []string{"bar", "foo", "bax"}, list.GetItems(2, -4))
	list.SetItems(2, []string{"foo", "baz", "bar"}) //       ["foo", "foo", "baz", "bar", "baz", "baz", "bar"]
	assert.Equal(t, []string{"foo", "baz", "bar"}, list.GetItems(2, -4))
	list.SetItems(-1, []string{"foz"})
	assert.Equal(t, "foz", string(list.GetItem(-1))) //      ["foo", "foo", "baz", "bar", "baz", "baz", "foz"]
	list.SortItems()                                 //      ["bar", "baz", "baz", "baz", "foo", "foo", "foz"]

	assert.Equal(t, []string{"baz", "baz"}, list.RemoveItems(2, -5)) // ["bar", "baz", "foo", "foo", "foz"]
	assert.Equal(t, []string{"bar", "baz"}, list.RemoveItems(1, 2))  // ["foo", "foo", "foz"]
	assert.Equal(t, "foz", string(list.RemoveItem(-1)))              // ["foo", "foo"]
	assert.Equal(t, 2, list.GetSize())                               // ["foo", "foo"]
	list.RemoveAll()                                                 // [ ]
	assert.Equal(t, 0, list.GetSize())                               // [ ]
	list.SortItems()                                                 // [ ]
	list.AddItems([]string{"foo", "bar", "baz"})                     // ["foo", "bar", "baz"]
	list.SortItems()                                                 // ["bar", "baz", "foo"]
	assert.Equal(t, []string{"bar", "baz", "foo"}, list.AsArray())   // ["bar", "baz", "foo"]
	list.RemoveAll()                                                 // [ ]
	list.AddItem("foo")                                              // ["foo"]
	list.SortItems()                                                 // ["foo"]
	assert.Equal(t, 1, list.GetSize())                               // ["foo"]
	assert.Equal(t, "foo", string(list.GetItem(1)))                  // ["foo"]
	list.AddItem("bar")                                              // ["foo", "bar"]
	list.SortItems()                                                 // ["bar", "foo"]
	assert.Equal(t, 2, list.GetSize())                               // ["bar", "foo"]
	assert.Equal(t, "bar", string(list.GetItem(1)))                  // ["bar", "foo"]
}

func TestListsWithConcatenate(t *testing.T) {
	var lists = collections.Lists[int]()
	var list1 = collections.List[int]()
	list1.AddItems([]int{1, 2, 3})
	var list2 = collections.List[int]()
	list2.AddItems([]int{4, 5, 6})
	var list3 = lists.Concatenate(list1, list2)
	var list4 = collections.List[int]()
	list4.AddItems([]int{1, 2, 3, 4, 5, 6})
	assert.True(t, agents.CompareValues(list3, list4))
}

func TestListsWithEmptyLists(t *testing.T) {
	var lists = collections.Lists[int]()
	var list1 = collections.List[int]()
	var list2 = collections.List[int]()
	var list3 = lists.Concatenate(list1, list2)
	assert.True(t, agents.CompareValues(list1, list2))
	assert.True(t, agents.CompareValues(list2, list3))
	assert.True(t, agents.CompareValues(list3, list1))
}
