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
)

// SET IMPLEMENTATION

// This constructor creates a new empty set that uses the canonical rank
// function.
func Set[T abs.ValueLike]() abs.SetLike[T] {
	var rank = age.RankValues
	var values = List[T]()
	return &set[T]{values, values, values, rank}
}

// This constructor creates a new set from the specified array that uses the
// canonical rank function.
func SetFromArray[T abs.ValueLike](array []T) abs.SetLike[T] {
	var v = Set[T]()
	for _, value := range array {
		v.AddItem(value)
	}
	return v
}

// This type defines the structure and methods associated with a set of values.
// The set uses ORDINAL based indexing rather than ZERO based indexing (see
// the description of what this means in the Sequential interface definition).
// This type is parameterized as follows:
//   - T is any type of value.
type set[T abs.ValueLike] struct {
	// Note: The delegated methods don't see the real collection type.
	abs.Sequential[T]
	abs.Indexed[T]
	values abs.ListLike[T]
	rank   abs.RankingFunction
}

// SEARCHABLE INTERFACE

// This method determines whether or not this set contains the specified value.
func (v *set[T]) ContainsItem(value T) bool {
	var _, found = v.search(value)
	return found
}

// This method determines whether or not this set contains ANY of the
// specified values.
func (v *set[T]) ContainsAny(values abs.Sequential[T]) bool {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsItem(value) {
			// This set contains at least one of the values.
			return true
		}
	}
	// This set does not contain any of the values.
	return false
}

// This method determines whether or not this set contains ALL of the
// specified values.
func (v *set[T]) ContainsAll(values abs.Sequential[T]) bool {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsItem(value) {
			// This set is missing at least one of the values.
			return false
		}
	}
	// This set does contains all of the values.
	return true
}

// FLEXIBLE INTERFACE

// This method sets the ranker function for this set.
func (v *set[T]) SetRanker(rank abs.RankingFunction) {
	if rank == nil {
		rank = age.RankValues
	}
	v.rank = rank
}

// This method adds the specified value to this set if it is not already a
// member of the set.
func (v *set[T]) AddItem(value T) {
	var slot, found = v.search(value)
	if !found {
		// The value is not already a member, so add it.
		v.values.InsertItem(slot, value)
	}
}

// This method adds the specified values to this set if they are not already
// members of the set.
func (v *set[T]) AddItems(values abs.Sequential[T]) {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.AddItem(value)
	}
}

// This method removes the specified value from this set. It returns true if the
// value was in the set and false otherwise.
func (v *set[T]) RemoveItem(value T) {
	var index, found = v.search(value)
	if found {
		// The value is a member, so remove it.
		v.values.RemoveItem(index)
	}
}

// This method removes the specified values from this set. It returns the number
// of values that were removed.
func (v *set[T]) RemoveItems(values abs.Sequential[T]) {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.RemoveItem(value)
	}
}

// This method removes all values from this set.
func (v *set[T]) RemoveAll() {
	v.values.RemoveAll()
}

// PRIVATE INTERFACE

// This private method performs a binary search of the set for the specified
// value. It returns two results:
//   - index: The index of the value, or if not found, the index of the value
//     before which it could be inserted in the underlying list.
//   - found: A boolean stating whether or not the value was found.
//
// The algoritm performs a true O[log(n)] worst case search.
func (v *set[T]) search(value T) (index int, found bool) {
	// We use iteration instead of recursion for better performance.
	//    start        first      middle       last          end
	//    |-------------||----------||----------||-------------|
	//                  |<-- size -------------->|
	//
	var first = 1          // Start at the beginning.
	var last = v.GetSize() // End at the end.
	var size = last        // Initially all values are candidates.
	for size > 0 {
		var middle = first + size/2 // Rounds down to the nearest integer.
		var candidate = v.GetItem(middle)
		switch v.rank(value, candidate) {
		case -1:
			// The index of the value is less than the middle
			// index so the first index stays the same.
			last = middle - 1 // We already tried the middle index.
			size = middle - first
		case 0:
			// The index of the value is the middle index.
			return middle, true
		case 1:
			// The index of the value is greater than the middle
			// index so the last index stays the same.
			first = middle + 1 // We already tried the middle index.
			size = last - middle
		}
	}
	// The value was not found, the last index represents the SLOT where it
	// would be inserted. Note that since the value was not found, the
	// indexes are inverted: last < first (i.e. last = first - 1).
	return last, false
}

// SETS LIBRARY

// This constructor creates a new sets library for the specified generic
// value type.
func Sets[T abs.ValueLike]() *sets[T] {
	return &sets[T]{}
}

// This type defines the library functions that operate on sets. Since
// sets have a parameterized value type this library type is also
// parameterized as follows:
//   - T is any type of value.
type sets[T abs.ValueLike] struct{}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified set.
func (l *sets[T]) Not(set abs.SetLike[T]) abs.SetLike[T] {
	panic("Not(set) is meaningless, use Sans(fullSet, set) instead.")
}

// This library function returns the logical conjunction of the specified sets.
func (l *sets[T]) And(first, second abs.SetLike[T]) abs.SetLike[T] {
	var result = Set[T]()
	var iterator = age.Iterator[T](first)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if second.ContainsItem(value) {
			result.AddItem(value)
		}
	}
	return result
}

// This library function returns the logical material non-implication of the
// specified sets.
func (l *sets[T]) Sans(first, second abs.SetLike[T]) abs.SetLike[T] {
	var result = Set[T]()
	result.AddItems(first)
	result.RemoveItems(second)
	return result
}

// This library function returns the logical disjunction of the specified sets.
func (l *sets[T]) Or(first, second abs.SetLike[T]) abs.SetLike[T] {
	var result = Set[T]()
	result.AddItems(first)
	result.AddItems(second)
	return result
}

// This library function returns the logical exclusive disjunction of the
// specified sets.
func (l *sets[T]) Xor(first, second abs.SetLike[T]) abs.SetLike[T] {
	var result = l.Or(l.Sans(first, second), l.Sans(second, first))
	return result
}
