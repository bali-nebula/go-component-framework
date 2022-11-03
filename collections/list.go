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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	uti "github.com/craterdog-bali/go-bali-document-notation/utilities"
)

// LIST IMPLEMENTATION

// This constructor creates a new empty list that uses the canonical compare
// function.
func List[V abs.ValueLike]() abs.ListLike[V] {
	var capacity = 4 // The minimum value.
	var values = make([]V, 0, capacity)
	var compare = age.CompareValues
	return &list[V]{values, compare}
}

// This constructor creates a new list from the specified array that uses the
// canonical compare function.
func ListFromArray[V abs.ValueLike](array []V) abs.ListLike[V] {
	var v = List[V]()
	for _, value := range array {
		v.AddValue(value)
	}
	return v
}

// This type defines the structure and methods associated with a list of values.
// Each value is associated with an implicit positive integer index. The list
// uses ORDINAL based indexing rather than ZERO based indexing (see the
// description of what this means in the Sequential interface definition).
// This type is parameterized as follows:
//   - V is any type of value.
type list[V abs.ValueLike] struct {
	values  []V
	compare abs.ComparisonFunction
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this list is empty.
func (v *list[V]) IsEmpty() bool {
	return len(v.values) == 0
}

// This method returns the number of values contained in this list.
func (v *list[V]) GetSize() int {
	return len(v.values)
}

// This method returns all the values in this list. The values retrieved are in
// the same order as they are in the list.
func (v *list[V]) AsArray() []V {
	var length = len(v.values)
	var result = make([]V, length)
	copy(result, v.values)
	return result
}

// INDEXED INTERFACE

// This method sets the comparer function for this list.
func (v *list[V]) SetComparer(compare abs.ComparisonFunction) {
	if compare == nil {
		compare = age.CompareValues
	}
	v.compare = compare
}

// This method retrieves from this list the value that is associated with the
// specified index.
func (v *list[V]) GetValue(index int) V {
	var length = len(v.values)
	index = abs.NormalizedIndex(index, length)
	return v.values[index]
}

// This method retrieves from this list all values from the first index through
// the last index (inclusive).
func (v *list[V]) GetValues(first int, last int) abs.Sequential[V] {
	var length = len(v.values)
	first = abs.NormalizedIndex(first, length)
	last = abs.NormalizedIndex(last, length)
	var result = ListFromArray[V](v.values[first : last+1])
	return result
}

// This method returns the index of the FIRST occurence of the specified value in
// this list, or zero if this list does not contain the value.
func (v *list[V]) GetIndex(value V) int {
	for index, candidate := range v.values {
		if v.compare(candidate, value) {
			// Found the value.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The value was not found.
	return 0
}

// SEARCHABLE INTERFACE

// This method determines whether or not this list contains the specified value.
func (v *list[V]) ContainsValue(value V) bool {
	return v.GetIndex(value) > 0
}

// This method determines whether or not this list contains ANY of the specified
// values.
func (v *list[V]) ContainsAny(values abs.Sequential[V]) bool {
	var iterator = age.Iterator[V](values)
	for iterator.HasNext() {
		var candidate = iterator.GetNext()
		if v.GetIndex(candidate) > 0 {
			// Found one of the values.
			return true
		}
	}
	// Did not find any of the values.
	return false
}

// This method determines whether or not this list contains ALL of the specified
// values.
func (v *list[V]) ContainsAll(values abs.Sequential[V]) bool {
	var iterator = age.Iterator[V](values)
	for iterator.HasNext() {
		var candidate = iterator.GetNext()
		if v.GetIndex(candidate) == 0 {
			// Didn't find one of the values.
			return false
		}
	}
	// Found all of the values.
	return true
}

// MALLEABLE INTERFACE

// This method appends the specified value to the end of this list.
func (v *list[V]) AddValue(value V) {
	// Add space for the new value.
	var index = len(v.values)
	var length = index + 1
	v.resize(length)

	// Append the new value.
	v.values[index] = value
}

// This method appends the specified values to the end of this list.
func (v *list[V]) AddValues(values abs.Sequential[V]) {
	// Add space for the new values.
	var index = len(v.values)
	var length = index + values.GetSize()
	v.resize(length)

	// Append the new values.
	copy(v.values[index:], values.AsArray())
}

// This method sets the value in this list that is associated with the specified
// index to be the specified value.
func (v *list[V]) SetValue(index int, value V) {
	var length = len(v.values)
	index = abs.NormalizedIndex(index, length)
	v.values[index] = value
}

// This method sets the values in this list starting with the specified index
// to the specified values.
func (v *list[V]) SetValues(index int, values abs.Sequential[V]) {
	var length = len(v.values)
	index = abs.NormalizedIndex(index, length)
	copy(v.values[index:], values.AsArray())
}

// This method inserts the specified value into this list in the specified
// slot between existing values.
func (v *list[V]) InsertValue(slot int, value V) {
	// Add space for the new value.
	var length = len(v.values) + 1
	v.resize(length)

	// Insert the new value.
	copy(v.values[slot+1:], v.values[slot:])
	v.values[slot] = value
}

// This method inserts the specified values into this list in the specified
// slot between existing values.
func (v *list[V]) InsertValues(slot int, values abs.Sequential[V]) {
	// Add space for the new values.
	var size = values.GetSize()
	var length = len(v.values) + size
	v.resize(length)

	// Insert the new values.
	copy(v.values[slot+size:], v.values[slot:])
	copy(v.values[slot:], values.AsArray())
}

// This method removes the value at the specified index from this list. The
// removed value is returned.
func (v *list[V]) RemoveValue(index int) V {
	// Remove the old value.
	var length = len(v.values)
	index = abs.NormalizedIndex(index, length)
	var old = v.values[index]
	copy(v.values[index:], v.values[index+1:])

	// Remove extra space.
	v.resize(length - 1)
	return old
}

// This method removes the values in the specified index range from this list.
// The removed values are returned.
func (v *list[V]) RemoveValues(first int, last int) abs.Sequential[V] {
	// Remove the specified values.
	var length = len(v.values)
	first = abs.NormalizedIndex(first, length)
	last = abs.NormalizedIndex(last, length)
	var result = ListFromArray[V](v.values[first : last+1])
	copy(v.values[first:], v.values[last+1:])

	// Remove the extra space.
	var size = last - first + 1
	v.resize(length - size)
	return result
}

// This method removes all values from this list.
func (v *list[V]) RemoveAll() {
	v.values = make([]V, 0, 4)
}

// This method pseudo-randomly shuffles the values in this list.
func (v *list[V]) ShuffleValues() {
	uti.ShuffleArray[V](v.values)
}

// This method sorts the values in this list using the canonical rank function.
func (v *list[V]) SortValues() {
	v.SortValuesWithRanker(nil)
}

// This method sorts the values in this list using the specified rank function.
func (v *list[V]) SortValuesWithRanker(rank abs.RankingFunction) {
	if rank == nil {
		rank = age.RankValues
	}
	if len(v.values) > 1 {
		age.SortArray[V](v.values, rank)
	}
}

// This method reverses the order of all values in this list.
func (v *list[V]) ReverseValues() {
	// Allocate a new array/slice.
	var length = len(v.values)
	var capacity = cap(v.values)
	var reversed = make([]V, length, capacity)

	// Copy the values into the new array in reverse.
	for i, _ := range v.values {
		reversed[i] = v.values[length-i-1]
	}
	v.values = reversed
}

// PRIVATE INTERFACE

// This method resizes this list and adjusts the capacity of this list if
// necessary to make it more efficient. Note: Any additional values that are
// added to the length of the list are NOT zeroed out.
func (v *list[V]) resize(length int) {
	var capacity = cap(v.values)
	for length > capacity {
		capacity *= 2
	}
	for length < capacity/4 {
		capacity /= 2
	}
	if capacity != cap(v.values) {
		// Adjust the capacity accordingly.
		var values = make([]V, length, capacity)
		copy(values, v.values)
		v.values = values
	}
	v.values = v.values[:length] // A change the length of the slice.
}

// LISTS LIBRARY

// This constructor creates a new lists library for the specified generic
// value type.
func Lists[V abs.ValueLike]() *lists[V] {
	return &lists[V]{}
}

// This type defines the library functions that operate on lists. Since
// lists have a parameterized value type this library type is also
// parameterized as follows:
//   - V is any type of value.
type lists[V abs.ValueLike] struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified lists.
func (l *lists[V]) Concatenate(first, second abs.ListLike[V]) abs.ListLike[V] {
	var result = List[V]()
	result.AddValues(first)
	result.AddValues(second)
	return result
}
