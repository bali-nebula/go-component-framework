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
	uni "unicode"
)

// RUNE IMPLEMENTATION

// This constructor returns the minimum value for a rune.
func MinimumRune() Rune {
	return Rune(0)
}

// This constructor returns the maximum value for a rune.
func MaximumRune() Rune {
	return Rune(mat.MaxInt32)
}

// This type defines the methods associated with runes. It extends the
// native Go int32 type.
type Rune int32

// NUMERIC INTERFACE

// This method determines whether or not this rune is discrete.
func (v Rune) IsDiscrete() bool {
	return true
}

// This method determines whether or not this rune is zero.
func (v Rune) IsZero() bool {
	return v == 0
}

// This method determines whether or not this rune is infinite.
func (v Rune) IsInfinite() bool {
	return false
}

// This method determines whether or not this rune is undefined.
func (v Rune) IsUndefined() bool {
	return !uni.IsPrint(rune(v))
}

// This method returns a boolean value for this rune.
func (v Rune) AsBoolean() bool {
	return v > -1
}

// This method returns an integer value for this rune.
func (v Rune) AsInteger() int {
	return int(v)
}

// This method returns a real value for this rune.
func (v Rune) AsReal() float64 {
	return float64(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this rune is negative.
func (v Rune) IsNegative() bool {
	return v < 0
}

// INTERVAL IMPLEMENTATION

// This constructor creates a new interval range of values covering the
// specified endpoints. Note that at least one of the endpoints must be non-nil
// so that the endpoint type may be determined.
func Interval[V Discrete](first V, extent Extent, last V) IntervalLike[V] {
	var v = interval[V]{first: first, extent: extent, last: last}
	v.validateInterval()
	return &v
}

// This type defines the structure and methods associated with an interval range
// of values. This type is parameterized as follows:
//   - V is any endpoint type.
type interval[V Discrete] struct {
	first  V
	extent Extent
	last   V
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this interval range is empty.
func (v *interval[V]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this interval range.
func (v *interval[V]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this interval range. The
// values retrieved are in the same order as they are in the interval range.
func (v *interval[V]) AsArray() []V {
	var size = v.size
	if size > 256 {
		size = 256
	}
	var array = make([]V, size)
	if size > 0 {
		var first = v.firstIndex()
		for i := 0; i < size; i++ {
			var value = V(first + i)
			array[i] = value
		}
	}
	return array
}

// INDEXED INTERFACE

// This method retrieves from this interval range the value that is associated
// with the specified index.
func (v *interval[V]) GetValue(index int) V {
	var offset = v.GoIndex(index)
	if offset < 0 {
		panic("The index is outside the interval range of values.")
	}
	var first = v.firstIndex()
	var value = V(first + offset)
	return value
}

// This method retrieves from this interval range all values from the first
// index through the last index (inclusive).
func (v *interval[V]) GetValues(first int, last int) col.Sequential[V] {
	var values = col.List[V]()
	for index := first; index <= last; index++ {
		var value V = V(index)
		values.AddValue(value)
	}
	return values
}

// This method returns the index of the FIRST occurence of the specified value
// in this interval range, or zero if this interval range does not contain the
// value.
func (v *interval[V]) GetIndex(value V) int {
	var index = int(value)
	var first = v.firstIndex()
	var offset = index - first + 1
	if offset < 0 || offset > v.size {
		return 0
	}
	return offset
}

// ELASTIC INTERFACE

// This method returns the first value in this interval range.
func (v *interval[V]) GetFirst() V {
	return v.first
}

// This method sets the first value in this interval range.
func (v *interval[V]) SetFirst(value V) {
	v.first = value
	v.validateInterval()
}

// This method returns the extent for this interval range.
func (v *interval[V]) GetExtent() Extent {
	return v.extent
}

// This method sets the extent for this interval range.
func (v *interval[V]) SetExtent(extent Extent) {
	v.extent = extent
	v.validateInterval()
}

// This method returns the last value in this interval range.
func (v *interval[V]) GetLast() V {
	return v.last
}

// This method sets the last value in this interval range.
func (v *interval[V]) SetLast(value V) {
	v.last = value
	v.validateInterval()
}

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *interval[V]) validateInterval() {
	// Validate the extent.
	switch v.extent {
	case INCLUSIVE:
	case LEFT:
	case RIGHT:
	case EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid interval range extent: %v", v.extent))
	}

	// Validate the endpoints.
	var first = v.effectiveFirst()
	var last = v.effectiveLast()
	v.size = last - first + 1
	if col.RankValues(first, last) > 0 {
		panic("The effective first value in the interval range must not be more than the effective last value.")
	}

}

// This method returns the effective first value in the interval range based on
// its extent type. It can only be called when the parameterized value type V is
// an integer type.
func (v *interval[V]) firstIndex() int {
	var firstIndex = v.effectiveFirst()
	return firstIndex
}

// This method returns the effective first value in the interval range.
func (v *interval[V]) effectiveFirst() int {
	var first = int(v.first)
	switch v.extent {
	case RIGHT, EXCLUSIVE:
		first++
	}
	return first
}

// This method returns the effective last value in the interval range.
func (v *interval[V]) effectiveLast() int {
	var last = int(v.last)
	switch v.extent {
	case LEFT, EXCLUSIVE:
		last--
	}
	return last
}

// This method normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func (v *interval[V]) GoIndex(index int) int {
	var length = v.size
	switch {
	case index < -length || index == 0 || index > length:
		// The index is outside the bounds of the specified interval range.
		return -1
	case index < 0:
		// Convert a negative index.
		return index + length
	case index > 0:
		// Convert a positive index.
		return index - 1
	default:
		// This should never happen so time to panic...
		panic(fmt.Sprintf("Compiler problem, unexpected index value: %v", index))
	}
}
