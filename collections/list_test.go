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
	var foo = collections.ListFromArray([]string{"foo"})
	var bar = collections.ListFromArray([]string{"bar"})
	var baz = collections.ListFromArray([]string{"baz"})
	var foz = collections.ListFromArray([]string{"foz"})
	var barbaz = collections.ListFromArray([]string{"bar", "baz"})
	var bazbaz = collections.ListFromArray([]string{"baz", "baz"})
	var foobar = collections.ListFromArray([]string{"foo", "bar"})
	var baxbaz = collections.ListFromArray([]string{"bax", "baz"})
	var baxbez = collections.ListFromArray([]string{"bax", "bez"})
	var barfoobax = collections.ListFromArray([]string{"bar", "foo", "bax"})
	var foobazbar = collections.ListFromArray([]string{"foo", "baz", "bar"})
	var foobarbaz = collections.ListFromArray([]string{"foo", "bar", "baz"})
	var barbazfoo = collections.ListFromArray([]string{"bar", "baz", "foo"})
	var list = collections.List[string]()
	assert.True(t, list.IsEmpty())
	assert.Equal(t, 0, list.GetSize())
	assert.False(t, list.ContainsItem("bax"))
	assert.Equal(t, []string{}, list.AsArray())
	var iterator = agents.Iterator[string](list)
	assert.False(t, iterator.HasNext())
	assert.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	list.ShuffleItems()
	list.SortItems()
	list.RemoveAll()                                //        [ ]
	list.AddItem("foo")                             //        ["foo"]
	assert.False(t, list.IsEmpty())                 //        ["foo"]
	assert.Equal(t, 1, list.GetSize())              //        ["foo"]
	assert.Equal(t, "foo", string(list.GetItem(1))) //        ["foo"]
	list.AddItems(barbaz)                           //        ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetSize())              //        ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(list.GetItem(1))) //        ["foo", "bar", "baz"]
	assert.Equal(t, barbaz.AsArray(), list.GetItems(2, 3).AsArray())
	assert.Equal(t, foo.AsArray(), list.GetItems(1, 1).AsArray())
	iterator = agents.Iterator[string](list)               // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasNext())                     // ["foo", "bar", "baz"]
	assert.False(t, iterator.HasPrevious())                // ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(iterator.GetNext()))     // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasPrevious())                 // ["foo", "bar", "baz"]
	iterator.ToEnd()                                       // ["foo", "bar", "baz"]
	assert.False(t, iterator.HasNext())                    // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasPrevious())                 // ["foo", "bar", "baz"]
	assert.Equal(t, "baz", string(iterator.GetPrevious())) // ["foo", "bar", "baz"]
	assert.True(t, iterator.HasNext())                     // ["foo", "bar", "baz"]
	list.ShuffleItems()                                    // [ ?, ?, ? ]
	list.RemoveAll()                                       // [ ]
	assert.True(t, list.IsEmpty())                         // [ ]
	assert.Equal(t, 0, list.GetSize())                     // [ ]
	list.InsertItem(0, "baz")                              // ["baz"]
	assert.Equal(t, 1, list.GetSize())                     // ["baz"]
	assert.Equal(t, "baz", string(list.GetItem(-1)))       // ["baz"]
	list.InsertItems(0, foobar)                            // ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetSize())                     // ["foo", "bar", "baz"]
	assert.Equal(t, "foo", string(list.GetItem(-3)))       // ["foo", "bar", "baz"]
	assert.Equal(t, "bar", string(list.GetItem(-2)))       // ["foo", "bar", "baz"]
	assert.Equal(t, "baz", string(list.GetItem(-1)))       // ["foo", "bar", "baz"]
	list.ReverseItems()                                    // ["baz", "bar", "foo"]
	assert.Equal(t, "foo", string(list.GetItem(-1)))       // ["baz", "bar", "foo"]
	assert.Equal(t, "bar", string(list.GetItem(-2)))       // ["baz", "bar", "foo"]
	assert.Equal(t, "baz", string(list.GetItem(-3)))       // ["baz", "bar", "foo"]
	list.ReverseItems()                                    // ["foo", "bar", "baz"]
	assert.Equal(t, 0, list.GetIndex("foz"))               // ["foo", "bar", "baz"]
	assert.Equal(t, 3, list.GetIndex("baz"))               // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsItem("baz"))               // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsItem("bax"))              // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsAny(baxbaz))               // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsAny(baxbez))              // ["foo", "bar", "baz"]
	assert.True(t, list.ContainsAll(barbaz))               // ["foo", "bar", "baz"]
	assert.False(t, list.ContainsAll(baxbaz))              // ["foo", "bar", "baz"]
	list.SetItem(3, "bax")                                 // ["foo", "bar", "bax"]
	list.InsertItems(3, baz)                               // ["foo", "bar", "bax", "baz"]
	assert.Equal(t, 4, list.GetSize())                     // ["foo", "bar", "bax", "baz"]
	assert.Equal(t, "baz", string(list.GetItem(4)))        // ["foo", "bar", "bax", "baz"]
	list.InsertItem(4, "bar")                              // ["foo", "bar", "bax", "baz", "bar"]
	assert.Equal(t, 5, list.GetSize())                     // ["foo", "bar", "bax", "baz", "bar"]
	assert.Equal(t, "bar", string(list.GetItem(5)))        // ["foo", "bar", "bax", "baz", "bar"]
	list.InsertItem(2, "foo")                              // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, 6, list.GetSize())                     // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "bar", string(list.GetItem(2)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "foo", string(list.GetItem(3)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, "bax", string(list.GetItem(4)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	assert.Equal(t, bar.AsArray(), list.GetItems(6, 6).AsArray())
	list.InsertItems(5, baz)                        //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, 7, list.GetSize())              //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "bax", string(list.GetItem(4))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "baz", string(list.GetItem(5))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, "baz", string(list.GetItem(6))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	assert.Equal(t, barfoobax.AsArray(), list.GetItems(2, -4).AsArray())
	list.SetItems(2, foobazbar) //                      ["foo", "foo", "baz", "bar", "baz", "baz", "bar"]
	assert.Equal(t, foobazbar.AsArray(), list.GetItems(2, -4).AsArray())
	list.SetItems(-1, foz)
	assert.Equal(t, "foz", string(list.GetItem(-1))) // ["foo", "foo", "baz", "bar", "baz", "baz", "foz"]
	list.SortItems()                                 // ["bar", "baz", "baz", "baz", "foo", "foo", "foz"]

	assert.Equal(t, bazbaz.AsArray(), list.RemoveItems(2, -5).AsArray()) // ["bar", "baz", "foo", "foo", "foz"]
	assert.Equal(t, barbaz.AsArray(), list.RemoveItems(1, 2).AsArray())  // ["foo", "foo", "foz"]
	assert.Equal(t, "foz", string(list.RemoveItem(-1)))                  // ["foo", "foo"]
	assert.Equal(t, 2, list.GetSize())                                   // ["foo", "foo"]
	list.RemoveAll()                                                     // [ ]
	assert.Equal(t, 0, list.GetSize())                                   // [ ]
	list.SortItems()                                                     // [ ]
	list.AddItems(foobarbaz)                                             // ["foo", "bar", "baz"]
	list.SortItems()                                                     // ["bar", "baz", "foo"]
	assert.Equal(t, barbazfoo.AsArray(), list.AsArray())                 // ["bar", "baz", "foo"]
	list.RemoveAll()                                                     // [ ]
	list.AddItem("foo")                                                  // ["foo"]
	list.SortItems()                                                     // ["foo"]
	assert.Equal(t, 1, list.GetSize())                                   // ["foo"]
	assert.Equal(t, "foo", string(list.GetItem(1)))                      // ["foo"]
	list.AddItem("bar")                                                  // ["foo", "bar"]
	list.SortItems()                                                     // ["bar", "foo"]
	assert.Equal(t, 2, list.GetSize())                                   // ["bar", "foo"]
	assert.Equal(t, "bar", string(list.GetItem(1)))                      // ["bar", "foo"]
}

func TestListsWithConcatenate(t *testing.T) {
	var lists = collections.Lists[int]()
	var list1 = collections.List[int]()
	var onetwothree = collections.ListFromArray([]int{1, 2, 3})
	var fourfivesix = collections.ListFromArray([]int{4, 5, 6})
	var onethrusix = collections.ListFromArray([]int{1, 2, 3, 4, 5, 6})
	list1.AddItems(onetwothree)
	var list2 = collections.List[int]()
	list2.AddItems(fourfivesix)
	var list3 = lists.Concatenate(list1, list2)
	var list4 = collections.List[int]()
	list4.AddItems(onethrusix)
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
