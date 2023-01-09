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
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	cox "github.com/craterdog/go-collection-framework"
)

// SERIES IMPLEMENTATION

// This constructor creates a new empty series of component values.
func Series() abs.SeriesLike {
	var array = make([]abs.ComponentLike, 0)
	return &series{array}
}

// This constructor creates a new series of component values from the specified
// array.
func SeriesFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.SeriesLike {
	var array = sequence.AsArray()
	return &series{array}
}

// This type defines the structure and methods associated with a series collection
// of values.
type series struct {
	array []abs.ComponentLike
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this series collection is empty.
func (v *series) IsEmpty() bool {
	return len(v.array) == 0
}

// This method returns the number of values contained in this series collection.
func (v *series) GetSize() int {
	return len(v.array)
}

// This method returns up to the first 256 values in this series collection. The
// values retrieved are in the same order as they are in the series collection.
func (v *series) AsArray() []abs.ComponentLike {
	var length = len(v.array)
	var array = make([]abs.ComponentLike, length)
	copy(array, v.array)
	return array
}

// INDEXED INTERFACE

// This method retrieves from this series collection the value that is associated
// with the specified index.
func (v *series) GetValue(index int) abs.ComponentLike {
	index = v.goIndex(index)
	return v.array[index]
}

// This method retrieves from this series collection all values from the first
// index through the last index (inclusive).
func (v *series) GetValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	first = v.goIndex(first)
	last = v.goIndex(last)
	return &series{v.AsArray()[first : last+1]}
}

// This method returns the index of the FIRST occurence of the specified value
// in this series collection, or zero if this series collection does not contain the
// value.
func (v *series) GetIndex(value abs.ComponentLike) int {
	for index, candidate := range v.array {
		if cox.CompareValues(candidate, value) {
			// Found the value.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The value was not found.
	return 0
}

// SEARCHABLE INTERFACE

// This method determines whether or not the specified value is included in this
// series collection.
func (v *series) ContainsValue(value abs.ComponentLike) bool {
	return v.GetIndex(value) > 0
}

// This method determines whether or not this series collection contains ANY of the
// specified values.
func (v *series) ContainsAny(values abs.Sequential[abs.ComponentLike]) bool {
	var iterator = cox.Iterator[abs.ComponentLike](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This series collection contains at least one of the values.
			return true
		}
	}
	// This series collection does not contain any of the values.
	return false
}

// This method determines whether or not this series collection contains ALL of the
// specified values.
func (v *series) ContainsAll(values abs.Sequential[abs.ComponentLike]) bool {
	var iterator = cox.Iterator[abs.ComponentLike](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This series collection is missing at least one of the values.
			return false
		}
	}
	// This series collection does contains all of the values.
	return true
}

// PRIVATE INTERFACE

// This method normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func (v *series) goIndex(index int) int {
	var length = len(v.array)
	switch {
	case length == 0:
		// The array is empty.
		panic("Cannot index an empty array.")
	case index == 0:
		// Zero is not an ordinal.
		panic("Indices must be positive or negative ordinals, not zero.")
	case index < -length || index > length:
		// The index is outside the bounds of the specified range.
		panic(fmt.Sprintf(
			"The specified index is outside the allowed ranges [-%v..-1] and [1..%v]: %v",
			length,
			length,
			index))
	case index < 0:
		// Convert a negative index.
		return index + length
	case index > 0:
		// Convert a positive index.
		return index - 1
	default:
		// This should never happen so time to...
		panic(fmt.Sprintf("Go compiler problem, unexpected index value: %v", index))
	}
}
