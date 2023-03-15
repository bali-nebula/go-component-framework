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

// SPECTRUM IMPLEMENTATION

// This constructor creates a new lexical range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Spectrum[V abs.Lexical](first V, extent abs.Extent, last V) abs.SpectrumLike[V] {
	var v = spectrum[V]{first: first, extent: extent, last: last}
	v.validateSpectrum()
	return &v
}

// This type defines the structure and methods associated with a lexical
// range of values. This type is parameterized as follows:
//   - V is any endpoint type.
type spectrum[V abs.Lexical] struct {
	first  V
	extent abs.Extent
	last   V
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this lexical range is empty.
func (v *spectrum[V]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this lexical range.
func (v *spectrum[V]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this lexical range. The
// values retrieved are in the same order as they are in the lexical range.
func (v *spectrum[V]) AsArray() []V {
	var array = make([]V, 0)
	return array
}

// BOUNDED INTERFACE

// This method returns the first value in this lexical range.
func (v *spectrum[V]) GetFirst() V {
	return v.first
}

// This method sets the first value in this lexical range.
func (v *spectrum[V]) SetFirst(value V) {
	v.first = value
	v.validateSpectrum()
}

// This method returns the extent for this lexical range.
func (v *spectrum[V]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this lexical range.
func (v *spectrum[V]) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateSpectrum()
}

// This method returns the last value in this lexical range.
func (v *spectrum[V]) GetLast() V {
	return v.last
}

// This method sets the last value in this lexical range.
func (v *spectrum[V]) SetLast(value V) {
	v.last = value
	v.validateSpectrum()
}

// SEARCHABLE INTERFACE

// This method always returns zero since a lexical range is not indexable.
func (v *spectrum[V]) GetIndex(value V) int {
	return 0
}

// This method determines whether or not the specified value is included in this
// spectrum range.
func (v *spectrum[V]) ContainsValue(value V) bool {
	var first = v.first.AsString()
	var candidate = value.AsString()
	var last = v.last.AsString()
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
		var message = fmt.Sprintf("Received an invalid spectrum range extent: %v", v.extent)
		panic(message)
	}
}

// This method determines whether or not this spectrum contains ANY of the
// specified values.
func (v *spectrum[V]) ContainsAny(values abs.Sequential[V]) bool {
	var iterator = col.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This spectrum contains at least one of the values.
			return true
		}
	}
	// This spectrum does not contain any of the values.
	return false
}

// This method determines whether or not this spectrum contains ALL of the
// specified values.
func (v *spectrum[V]) ContainsAll(values abs.Sequential[V]) bool {
	var iterator = col.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This spectrum is missing at least one of the values.
			return false
		}
	}
	// This spectrum does contains all of the values.
	return true
}

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *spectrum[V]) validateSpectrum() {
	// Validate the extent.
	switch v.extent {
	case abs.INCLUSIVE:
	case abs.LEFT:
	case abs.RIGHT:
	case abs.EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid lexical range extent: %v", v.extent))
	}

	// Validate the endpoints.
	var first = v.first.AsString()
	var last = v.last.AsString()
	var rank = col.RankValues(first, last)
	switch {
	case rank < 0:
		v.size = -1 // The size of a spectrum is infinite.
	case rank == 0:
		v.size = 0
	default:
		panic("The first value in the lexical range must not be more than the last value.")
	}
}
