/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package temporary

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework"
	mat "math"
	ref "reflect"
)

// REAL IMPLEMENTATION

// This constructor returns the minimum value for a real.
func MinimumReal() Real {
	return Real(mat.Inf(-1))
}

// This constructor returns the maximum value for a real.
func MaximumReal() Real {
	return Real(mat.Inf(1))
}

// This type defines the methods associated with real numbers. It extends the
// native Go float64 type.
type Real float64

// NUMERIC INTERFACE

// This method determines whether or not this real number is discrete.
func (v Real) IsDiscrete() bool {
	return !mat.IsInf(float64(v), 0) && mat.Round(float64(v)) == float64(v)
}

// This method determines whether or not this real number is zero.
func (v Real) IsZero() bool {
	return v == 0
}

// This method determines whether or not this real number is infinite.
func (v Real) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this real number is undefined.
func (v Real) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// This method returns a boolean value for this real number.
func (v Real) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this real number.
func (v Real) AsInteger() int {
	if v.IsInfinite() {
		return mat.MaxInt
	}
	if v.IsUndefined() {
		return mat.MinInt
	}
	return int(mat.Round(float64(v)))
}

// This method returns a real value for this real number.
func (v Real) AsReal() float64 {
	return float64(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this real number is negative.
func (v Real) IsNegative() bool {
	return v < 0
}

// CONTINUUM IMPLEMENTATION

// This constructor creates a new continuous range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Continuum[V Continuous](first V, extent Extent, last V) ContinuumLike[V] {
	var v = continuum[V]{first: first, extent: extent, last: last}
	v.validateContinuum()
	return &v
}

// This type defines the structure and methods associated with a continuous
// range of values. This type is parameterized as follows:
//   - V is any endpoint type.
type continuum[V Continuous] struct {
	first  V
	extent Extent
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

// ELASTIC INTERFACE

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
func (v *continuum[V]) GetExtent() Extent {
	return v.extent
}

// This method sets the extent for this continuous range.
func (v *continuum[V]) SetExtent(extent Extent) {
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

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *continuum[V]) validateContinuum() {
	// Validate the extent.
	switch v.extent {
	case INCLUSIVE:
	case LEFT:
	case RIGHT:
	case EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid continuous range extent: %v", v.extent))
	}

	// Validate the endpoints.
	if !ref.ValueOf(v.first).IsValid() || !ref.ValueOf(v.last).IsValid() {
		panic("A continuous range requires first and last endpoints.")
	}
	var first = float64(v.first)
	var last = float64(v.last)
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
