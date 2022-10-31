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
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestListsWithStrings(t *tes.T) {
	var foo = col.ListFromArray([]string{"foo"})
	var bar = col.ListFromArray([]string{"bar"})
	var baz = col.ListFromArray([]string{"baz"})
	var foz = col.ListFromArray([]string{"foz"})
	var barbaz = col.ListFromArray([]string{"bar", "baz"})
	var bazbaz = col.ListFromArray([]string{"baz", "baz"})
	var foobar = col.ListFromArray([]string{"foo", "bar"})
	var baxbaz = col.ListFromArray([]string{"bax", "baz"})
	var baxbez = col.ListFromArray([]string{"bax", "bez"})
	var barfoobax = col.ListFromArray([]string{"bar", "foo", "bax"})
	var foobazbar = col.ListFromArray([]string{"foo", "baz", "bar"})
	var foobarbaz = col.ListFromArray([]string{"foo", "bar", "baz"})
	var barbazfoo = col.ListFromArray([]string{"bar", "baz", "foo"})
	var list = col.List[string]()
	ass.True(t, list.IsEmpty())
	ass.Equal(t, 0, list.GetSize())
	ass.False(t, list.ContainsItem("bax"))
	ass.Equal(t, []string{}, list.AsArray())
	var iterator = age.Iterator[string](list)
	ass.False(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	list.ShuffleItems()
	list.SortItems()
	list.RemoveAll()                             //        [ ]
	list.AddItem("foo")                          //        ["foo"]
	ass.False(t, list.IsEmpty())                 //        ["foo"]
	ass.Equal(t, 1, list.GetSize())              //        ["foo"]
	ass.Equal(t, "foo", string(list.GetItem(1))) //        ["foo"]
	list.AddItems(barbaz)                        //        ["foo", "bar", "baz"]
	ass.Equal(t, 3, list.GetSize())              //        ["foo", "bar", "baz"]
	ass.Equal(t, "foo", string(list.GetItem(1))) //        ["foo", "bar", "baz"]
	ass.Equal(t, barbaz.AsArray(), list.GetItems(2, 3).AsArray())
	ass.Equal(t, foo.AsArray(), list.GetItems(1, 1).AsArray())
	iterator = age.Iterator[string](list)               // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())                     // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasPrevious())                // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", string(iterator.GetNext()))     // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())                 // ["foo", "bar", "baz"]
	iterator.ToEnd()                                    // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasNext())                    // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())                 // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", string(iterator.GetPrevious())) // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())                     // ["foo", "bar", "baz"]
	list.ShuffleItems()                                 // [ ?, ?, ? ]
	list.RemoveAll()                                    // [ ]
	ass.True(t, list.IsEmpty())                         // [ ]
	ass.Equal(t, 0, list.GetSize())                     // [ ]
	list.InsertItem(0, "baz")                           // ["baz"]
	ass.Equal(t, 1, list.GetSize())                     // ["baz"]
	ass.Equal(t, "baz", string(list.GetItem(-1)))       // ["baz"]
	list.InsertItems(0, foobar)                         // ["foo", "bar", "baz"]
	ass.Equal(t, 3, list.GetSize())                     // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", string(list.GetItem(-3)))       // ["foo", "bar", "baz"]
	ass.Equal(t, "bar", string(list.GetItem(-2)))       // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", string(list.GetItem(-1)))       // ["foo", "bar", "baz"]
	list.ReverseItems()                                 // ["baz", "bar", "foo"]
	ass.Equal(t, "foo", string(list.GetItem(-1)))       // ["baz", "bar", "foo"]
	ass.Equal(t, "bar", string(list.GetItem(-2)))       // ["baz", "bar", "foo"]
	ass.Equal(t, "baz", string(list.GetItem(-3)))       // ["baz", "bar", "foo"]
	list.ReverseItems()                                 // ["foo", "bar", "baz"]
	ass.Equal(t, 0, list.GetIndex("foz"))               // ["foo", "bar", "baz"]
	ass.Equal(t, 3, list.GetIndex("baz"))               // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsItem("baz"))               // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsItem("bax"))              // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAny(baxbaz))               // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAny(baxbez))              // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAll(barbaz))               // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAll(baxbaz))              // ["foo", "bar", "baz"]
	list.SetItem(3, "bax")                              // ["foo", "bar", "bax"]
	list.InsertItems(3, baz)                            // ["foo", "bar", "bax", "baz"]
	ass.Equal(t, 4, list.GetSize())                     // ["foo", "bar", "bax", "baz"]
	ass.Equal(t, "baz", string(list.GetItem(4)))        // ["foo", "bar", "bax", "baz"]
	list.InsertItem(4, "bar")                           // ["foo", "bar", "bax", "baz", "bar"]
	ass.Equal(t, 5, list.GetSize())                     // ["foo", "bar", "bax", "baz", "bar"]
	ass.Equal(t, "bar", string(list.GetItem(5)))        // ["foo", "bar", "bax", "baz", "bar"]
	list.InsertItem(2, "foo")                           // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, 6, list.GetSize())                     // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bar", string(list.GetItem(2)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "foo", string(list.GetItem(3)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bax", string(list.GetItem(4)))        // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, bar.AsArray(), list.GetItems(6, 6).AsArray())
	list.InsertItems(5, baz)                     //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, 7, list.GetSize())              //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "bax", string(list.GetItem(4))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", string(list.GetItem(5))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", string(list.GetItem(6))) //  ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, barfoobax.AsArray(), list.GetItems(2, -4).AsArray())
	list.SetItems(2, foobazbar) //                      ["foo", "foo", "baz", "bar", "baz", "baz", "bar"]
	ass.Equal(t, foobazbar.AsArray(), list.GetItems(2, -4).AsArray())
	list.SetItems(-1, foz)
	ass.Equal(t, "foz", string(list.GetItem(-1))) // ["foo", "foo", "baz", "bar", "baz", "baz", "foz"]
	list.SortItems()                              // ["bar", "baz", "baz", "baz", "foo", "foo", "foz"]

	ass.Equal(t, bazbaz.AsArray(), list.RemoveItems(2, -5).AsArray()) // ["bar", "baz", "foo", "foo", "foz"]
	ass.Equal(t, barbaz.AsArray(), list.RemoveItems(1, 2).AsArray())  // ["foo", "foo", "foz"]
	ass.Equal(t, "foz", string(list.RemoveItem(-1)))                  // ["foo", "foo"]
	ass.Equal(t, 2, list.GetSize())                                   // ["foo", "foo"]
	list.RemoveAll()                                                  // [ ]
	ass.Equal(t, 0, list.GetSize())                                   // [ ]
	list.SortItems()                                                  // [ ]
	list.AddItems(foobarbaz)                                          // ["foo", "bar", "baz"]
	list.SortItems()                                                  // ["bar", "baz", "foo"]
	ass.Equal(t, barbazfoo.AsArray(), list.AsArray())                 // ["bar", "baz", "foo"]
	list.RemoveAll()                                                  // [ ]
	list.AddItem("foo")                                               // ["foo"]
	list.SortItems()                                                  // ["foo"]
	ass.Equal(t, 1, list.GetSize())                                   // ["foo"]
	ass.Equal(t, "foo", string(list.GetItem(1)))                      // ["foo"]
	list.AddItem("bar")                                               // ["foo", "bar"]
	list.SortItems()                                                  // ["bar", "foo"]
	ass.Equal(t, 2, list.GetSize())                                   // ["bar", "foo"]
	ass.Equal(t, "bar", string(list.GetItem(1)))                      // ["bar", "foo"]
}

func TestListsWithConcatenate(t *tes.T) {
	var lists = col.Lists[int]()
	var list1 = col.List[int]()
	var onetwothree = col.ListFromArray([]int{1, 2, 3})
	var fourfivesix = col.ListFromArray([]int{4, 5, 6})
	var onethrusix = col.ListFromArray([]int{1, 2, 3, 4, 5, 6})
	list1.AddItems(onetwothree)
	var list2 = col.List[int]()
	list2.AddItems(fourfivesix)
	var list3 = lists.Concatenate(list1, list2)
	var list4 = col.List[int]()
	list4.AddItems(onethrusix)
	ass.True(t, age.CompareItems(list3, list4))
}

func TestListsWithEmptyLists(t *tes.T) {
	var lists = col.Lists[int]()
	var list1 = col.List[int]()
	var list2 = col.List[int]()
	var list3 = lists.Concatenate(list1, list2)
	ass.True(t, age.CompareItems(list1, list2))
	ass.True(t, age.CompareItems(list2, list3))
	ass.True(t, age.CompareItems(list3, list1))
}
