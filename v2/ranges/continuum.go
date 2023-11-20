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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// CONTINUUM IMPLEMENTATION

// This constructor creates a new continuous range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Continuum(first abs.Continuous, extent abs.Extent, last abs.Continuous) abs.ContinuumLike {
	var v = continuum{first: first, extent: extent, last: last}
	v.validateContinuum()
	return &v
}

// This type defines the structure and methods associated with a continuous
// range of values.
type continuum struct {
	first  abs.Continuous
	extent abs.Extent
	last   abs.Continuous
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this continuous range is empty.
func (v *continuum) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this continuous range.
func (v *continuum) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this continuous range. The
// values retrieved are in the same order as they are in the continuous range.
func (v *continuum) AsArray() []abs.Continuous {
	var array = make([]abs.Continuous, 0)
	return array
}

// BOUNDED INTERFACE

// This method returns the first value in this continuous range.
func (v *continuum) GetFirst() abs.Continuous {
	return v.first
}

// This method sets the first value in this continuous range.
func (v *continuum) SetFirst(value abs.Continuous) {
	v.first = value
	v.validateContinuum()
}

// This method returns the extent for this continuous range.
func (v *continuum) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this continuous range.
func (v *continuum) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateContinuum()
}

// This method returns the last value in this continuous range.
func (v *continuum) GetLast() abs.Continuous {
	return v.last
}

// This method sets the last value in this continuous range.
func (v *continuum) SetLast(value abs.Continuous) {
	v.last = value
	v.validateContinuum()
}

// SEARCHABLE INTERFACE

// This method always returns zero since a continuous range is not indexable.
func (v *continuum) GetIndex(value abs.Continuous) int {
	return 0
}

// This method determines whether or not the specified value is included in this
// continuous range.
func (v *continuum) ContainsValue(value abs.Continuous) bool {
	var first = v.first.AsFloat()
	var candidate = value.AsFloat()
	var last = v.last.AsFloat()
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
func (v *continuum) ContainsAny(values abs.Sequential[abs.Continuous]) bool {
	var iterator = col.Iterator[abs.Continuous](values)
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
func (v *continuum) ContainsAll(values abs.Sequential[abs.Continuous]) bool {
	var iterator = col.Iterator[abs.Continuous](values)
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
func (v *continuum) validateContinuum() {
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
	var first = v.first.AsFloat()
	var last = v.last.AsFloat()
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
