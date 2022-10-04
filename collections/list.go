/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	//"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"math/rand"
	"time"
)

// LIST IMPLEMENTATION

// This constructor creates a new empty list that uses the canonical compare
// function.
func List[T any]() abstractions.ListLike[T] {
	return ListWithComparer[T](nil)
}

// This constructor creates a new empty list that uses the specified compare
// function.
func ListWithComparer[T any](compare abstractions.ComparisonFunction) abstractions.ListLike[T] {
	// Groom the arguments.
	var capacity = 4 // The minimum value.
	var items = make([]T, 0, capacity)
	if compare == nil {
		compare = agents.CompareValues
	}

	// Return a new list.
	return &list[T]{items, compare}
}

// This type defines the structure and methods associated with a list of items.
// Each item is associated with an implicit positive integer index. The list
// uses ORDINAL based indexing rather than ZERO based indexing (see the
// description of what this means in the Sequential interface definition).
// This type is parameterized as follows:
//   - T is any type of item.
type list[T any] struct {
	items   []T
	compare abstractions.ComparisonFunction
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this list is empty.
func (v *list[T]) IsEmpty() bool {
	return len(v.items) == 0
}

// This method returns the number of items contained in this list.
func (v *list[T]) GetSize() int {
	return len(v.items)
}

// This method returns all the items in this list. The items retrieved are in
// the same order as they are in the list.
func (v *list[T]) AsArray() []T {
	var length = len(v.items)
	var result = make([]T, length)
	copy(result, v.items)
	return result
}

// INDEXED INTERFACE

// This method retrieves from this list the item that is associated with the
// specified index.
func (v *list[T]) GetItem(index int) T {
	var length = len(v.items)
	index = abstractions.NormalizedIndex(index, length)
	return v.items[index]
}

// This method retrieves from this list all items from the first index through
// the last index (inclusive).
func (v *list[T]) GetItems(first int, last int) []T {
	var length = len(v.items)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	var size = last - first + 1
	var result = make([]T, size)
	copy(result, v.items[first:])
	return result
}

// This method returns the index of the FIRST occurence of the specified item in
// this list, or zero if this list does not contain the item.
func (v *list[T]) GetIndex(item T) int {
	for index, candidate := range v.items {
		if v.compare(candidate, item) {
			// Found the item.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The item was not found.
	return 0
}

// SEARCHABLE INTERFACE

// This method determines whether or not this list contains the specified item.
func (v *list[T]) ContainsItem(item T) bool {
	return v.GetIndex(item) > 0
}

// This method determines whether or not this list contains ANY of the specified
// items.
func (v *list[T]) ContainsAny(items []T) bool {
	for _, item := range items {
		if v.GetIndex(item) > 0 {
			// Found one of the items.
			return true
		}
	}
	// Did not find any of the items.
	return false
}

// This method determines whether or not this list contains ALL of the specified
// items.
func (v *list[T]) ContainsAll(items []T) bool {
	for _, item := range items {
		if v.GetIndex(item) == 0 {
			// Didn't find one of the items.
			return false
		}
	}
	// Found all of the items.
	return true
}

// MALLEABLE INTERFACE

// This method appends the specified item to the end of this list.
func (v *list[T]) AddItem(item T) {
	// Add space for the new item.
	var index = len(v.items)
	var length = index + 1
	v.resize(length)

	// Append the new item.
	v.items[index] = item
}

// This method appends the specified items to the end of this list.
func (v *list[T]) AddItems(items []T) {
	// Add space for the new items.
	var index = len(v.items)
	var length = index + len(items)
	v.resize(length)

	// Append the new items.
	copy(v.items[index:], items)
}

// This method sets the item in this list that is associated with the specified
// index to be the specified item. It returns the old value of that item.
func (v *list[T]) SetItem(index int, item T) T {
	var length = len(v.items)
	index = abstractions.NormalizedIndex(index, length)
	var old = v.items[index]
	v.items[index] = item
	return old
}

// This method sets the items in this list starting with the specified index
// to the specified items. It returns the old values of those items.
func (v *list[T]) SetItems(index int, items []T) []T {
	var length = len(v.items)
	index = abstractions.NormalizedIndex(index, length)
	var size = len(items)
	var old = make([]T, size)
	copy(old, v.items[index:])
	copy(v.items[index:], items)
	return old
}

// This method inserts the specified item into this list in the specified
// slot between existing items.
func (v *list[T]) InsertItem(slot int, item T) {
	// Add space for the new item.
	var length = len(v.items) + 1
	v.resize(length)

	// Insert the new item.
	copy(v.items[slot+1:], v.items[slot:])
	v.items[slot] = item
}

// This method inserts the specified items into this list in the specified
// slot between existing items.
func (v *list[T]) InsertItems(slot int, items []T) {
	// Add space for the new items.
	var length = len(v.items) + len(items)
	v.resize(length)

	// Insert the new items.
	copy(v.items[slot+len(items):], v.items[slot:])
	copy(v.items[slot:], items)
}

// This method removes the item at the specified index from this list. The
// removed item is returned.
func (v *list[T]) RemoveItem(index int) T {
	// Remove the old item.
	var length = len(v.items)
	index = abstractions.NormalizedIndex(index, length)
	var old = v.items[index]
	copy(v.items[index:], v.items[index+1:])

	// Remove extra space.
	v.resize(length - 1)
	return old
}

// This method removes the items in the specified index range from this list.
// The removed items are returned.
func (v *list[T]) RemoveItems(first int, last int) []T {
	// Remove the old items.
	var length = len(v.items)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	var size = last - first + 1
	var old = make([]T, size)
	copy(old, v.items[first:])
	copy(v.items[first:], v.items[last+1:])

	// Remove extra space.
	v.resize(length - size)
	return old
}

// This method removes all items from this list.
func (v *list[T]) RemoveAll() {
	v.items = make([]T, 0, 4)
}

// This method pseudo-randomly shuffles the items in this list.
func (v *list[T]) ShuffleItems() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(v.items), func(i, j int) {
		// Swap the two items.
		v.items[i], v.items[j] = v.items[j], v.items[i]
	})
}

// This method sorts the items in this list using the canonical rank function.
func (v *list[T]) SortItems() {
	v.SortItemsWithRanker(nil)
}

// This method sorts the items in this list using the specified rank function.
func (v *list[T]) SortItemsWithRanker(rank abstractions.RankingFunction) {
	if rank == nil {
		rank = agents.RankValues
	}
	if len(v.items) > 1 {
		agents.SortArray[T](v.items, rank)
	}
}

// This method reverses the order of all items in this list.
func (v *list[T]) ReverseItems() {
	// Allocate a new array/slice.
	var length = len(v.items)
	var capacity = cap(v.items)
	var reversed = make([]T, length, capacity)

	// Copy the items into the new array in reverse.
	for i, _ := range v.items {
		reversed[i] = v.items[length-i-1]
	}
	v.items = reversed
}

// PRIVATE INTERFACE

// This method resizes this list and adjusts the capacity of this list if
// necessary to make it more efficient. Note: Any additional items that are
// added to the length of the list are NOT zeroed out.
func (v *list[T]) resize(length int) {
	var capacity = cap(v.items)
	for length > capacity {
		capacity *= 2
	}
	for length < capacity/4 {
		capacity /= 2
	}
	if capacity != cap(v.items) {
		// Adjust the capacity accordingly.
		var items = make([]T, length, capacity)
		copy(items, v.items)
		v.items = items
	}
	v.items = v.items[:length] // A change the length of the slice.
}

// LISTS LIBRARY

// This constructor creates a new lists library for the specified generic
// item type.
func Lists[T any]() *lists[T] {
	return &lists[T]{}
}

// This type defines the library functions that operate on lists. Since
// lists have a parameterized item type this library type is also
// parameterized as follows:
//   - T is any type of item.
type lists[T any] struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified lists.
func (l *lists[T]) Concatenate(first, second abstractions.ListLike[T]) abstractions.ListLike[T] {
	var result = List[T]()
	result.AddItems(first.AsArray())
	result.AddItems(second.AsArray())
	return result
}
