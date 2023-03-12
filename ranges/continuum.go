/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package ranges

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// CONTINUUM IMPLEMENTATION

// This constructor creates a new continuous range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Continuum[V abs.Continuous](first V, extent abs.Extent, last V) abs.ContinuumLike[V] {
	var v = continuum[V]{first: first, extent: extent, last: last}
	v.validateContinuum()
	return &v
}

// This type defines the structure and methods associated with a continuous
// range of values. This type is parameterized as follows:
//   - V is any endpoint type.
type continuum[V abs.Continuous] struct {
	first  V
	extent abs.Extent
	last   V
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this continuous range is empty.
func (v *continuum[V]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this continuous range.
func (v *continuum[V]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this continuous range. The
// values retrieved are in the same order as they are in the continuous range.
func (v *continuum[V]) AsArray() []V {
	var array = make([]V, 0)
	return array
}

// BOUNDED INTERFACE

// This method returns the first value in this continuous range.
func (v *continuum[V]) GetFirst() V {
	return v.first
}

// This method sets the first value in this continuous range.
func (v *continuum[V]) SetFirst(value V) {
	v.first = value
	v.validateContinuum()
}

// This method returns the extent for this continuous range.
func (v *continuum[V]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this continuous range.
func (v *continuum[V]) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateContinuum()
}

// This method returns the last value in this continuous range.
func (v *continuum[V]) GetLast() V {
	return v.last
}

// This method sets the last value in this continuous range.
func (v *continuum[V]) SetLast(value V) {
	v.last = value
	v.validateContinuum()
}

// SEARCHABLE INTERFACE

// This method always returns zero since a continuous range is not indexable.
func (v *continuum[V]) GetIndex(value V) int {
	return 0
}

// This method determines whether or not the specified value is included in this
// continuous range.
func (v *continuum[V]) ContainsValue(value V) bool {
	var first = v.first.AsReal()
	var candidate = value.AsReal()
	var last = v.last.AsReal()
	switch v.extent {
	case abs.INCLUSIVE:
		return col.RankValues(first, candidate) <= 0 && col.RankValues(candidate, last) <= 0
	case abs.LEFT:
		return col.RankValues(first, candidate) <= 0 && col.RankValues(candidate, last) < 0
	case abs.RIGHT:
		return col.RankValues(first, candidate) < 0 && col.RankValues(candidate, last) <= 0
	case abs.EXCLUSIVE:
		return col.RankValues(first, candidate) < 0 && col.RankValues(candidate, last) < 0
	default:
		var message = fmt.Sprintf("Received an invalid continuous range extent: %v", v.extent)
		panic(message)
	}
}

// This method determines whether or not this continuum contains ANY of the
// specified values.
func (v *continuum[V]) ContainsAny(values abs.Sequential[V]) bool {
	var iterator = col.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This continuum contains at least one of the values.
			return true
		}
	}
	// This continuum does not contain any of the values.
	return false
}

// This method determines whether or not this continuum contains ALL of the
// specified values.
func (v *continuum[V]) ContainsAll(values abs.Sequential[V]) bool {
	var iterator = col.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This continuum is missing at least one of the values.
			return false
		}
	}
	// This continuum does contains all of the values.
	return true
}

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *continuum[V]) validateContinuum() {
	// Validate the extent.
	switch v.extent {
	case abs.INCLUSIVE:
	case abs.LEFT:
	case abs.RIGHT:
	case abs.EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid continuous range extent: %v", v.extent))
	}

	// Validate the endpoints.
	var first = v.first.AsReal()
	var last = v.last.AsReal()
	var rank = col.RankValues(first, last)
	switch {
	case rank < 0:
		v.size = -1 // The size of a continuum is infinite.
	case rank == 0:
		v.size = 0
	default:
		panic("The first value in the continuous range must not be more than the last value.")
	}
}
