/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
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
	col "github.com/craterdog/go-collection-framework"
)

// SPECTRUM IMPLEMENTATION

// This constructor creates a new spectral range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Spectrum[V abs.Spectral](first V, extent abs.Extent, last V) abs.SpectrumLike[V] {
	var v = spectrum[V]{first: first, extent: extent, last: last}
	v.validateSpectrum()
	return &v
}

// This type defines the structure and methods associated with a spectral
// range of values. This type is parameterized as follows:
//   - V is any endpoint type.
type spectrum[V abs.Spectral] struct {
	first  V
	extent abs.Extent
	last   V
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this spectral range is empty.
func (v *spectrum[V]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this spectral range.
func (v *spectrum[V]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this spectral range. The
// values retrieved are in the same order as they are in the spectral range.
func (v *spectrum[V]) AsArray() []V {
	var array = make([]V, 0)
	return array
}

// ELASTIC INTERFACE

// This method returns the first value in this spectral range.
func (v *spectrum[V]) GetFirst() V {
	return v.first
}

// This method sets the first value in this spectral range.
func (v *spectrum[V]) SetFirst(value V) {
	v.first = value
	v.validateSpectrum()
}

// This method returns the extent for this spectral range.
func (v *spectrum[V]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this spectral range.
func (v *spectrum[V]) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateSpectrum()
}

// This method returns the last value in this spectral range.
func (v *spectrum[V]) GetLast() V {
	return v.last
}

// This method sets the last value in this spectral range.
func (v *spectrum[V]) SetLast(value V) {
	v.last = value
	v.validateSpectrum()
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
		panic(fmt.Sprintf("Received an invalid spectral range extent: %v", v.extent))
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
		panic("The first value in the spectral range must not be more than the last value.")
	}
}
