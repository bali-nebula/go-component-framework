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
	ele "github.com/bali-nebula/go-component-framework/elements"
	cox "github.com/craterdog/go-collection-framework"
	mat "math"
)

// INTERVAL IMPLEMENTATION

// This constructor creates a new interval range of values covering the
// specified endpoints.
func Interval[V abs.Discrete](first V, extent abs.Extent, last V) abs.IntervalLike[V] {
	var v = interval[V]{first: first, extent: extent, last: last}
	v.validateInterval()
	return &v
}

// This type defines the structure and methods associated with an interval range
// of values. This type is parameterized as follows:
//   - V is any endpoint type.
type interval[V abs.Discrete] struct {
	first  V
	extent abs.Extent
	last   V
	size   int
}

// BOUNDED INTERFACE

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
func (v *interval[V]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this interval range.
func (v *interval[V]) SetExtent(extent abs.Extent) {
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
			var value = v.indexToValue(first + i)
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
	var value = v.indexToValue(first + offset)
	return value
}

// This method retrieves from this interval range all values from the first
// index through the last index (inclusive).
func (v *interval[V]) GetValues(first int, last int) abs.Sequential[V] {
	var values = cox.List[V]()
	for index := first; index <= last; index++ {
		var value V = v.indexToValue(index)
		values.AddValue(value)
	}
	return values
}

// This method returns the index of the FIRST occurence of the specified value
// in this interval range, or zero if this interval range does not contain the
// value.
func (v *interval[V]) GetIndex(value V) int {
	var index = value.AsInteger()
	var first = v.firstIndex()
	var offset = index - first + 1
	if offset < 0 || offset > v.size {
		return 0
	}
	return offset
}

// SEARCHABLE INTERFACE

// This method determines whether or not the specified value is included in this
// interval range.
func (v *interval[V]) ContainsValue(value V) bool {
	var first = v.first.AsInteger()
	var candidate = value.AsInteger()
	var last = v.last.AsInteger()
	switch v.extent {
	case abs.INCLUSIVE:
		return cox.RankValues(first, candidate) <= 0 && cox.RankValues(candidate, last) <= 0
	case abs.LEFT:
		return cox.RankValues(first, candidate) <= 0 && cox.RankValues(candidate, last) < 0
	case abs.RIGHT:
		return cox.RankValues(first, candidate) < 0 && cox.RankValues(candidate, last) <= 0
	case abs.EXCLUSIVE:
		return cox.RankValues(first, candidate) < 0 && cox.RankValues(candidate, last) < 0
	default:
		var message = fmt.Sprintf("Received an invalid interval range extent: %v", v.extent)
		panic(message)
	}
}

// This method determines whether or not this interval contains ANY of the
// specified values.
func (v *interval[V]) ContainsAny(values abs.Sequential[V]) bool {
	var iterator = cox.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This interval contains at least one of the values.
			return true
		}
	}
	// This interval does not contain any of the values.
	return false
}

// This method determines whether or not this interval contains ALL of the
// specified values.
func (v *interval[V]) ContainsAll(values abs.Sequential[V]) bool {
	var iterator = cox.Iterator[V](values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This interval is missing at least one of the values.
			return false
		}
	}
	// This interval does contains all of the values.
	return true
}

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *interval[V]) validateInterval() {
	// Validate the extent.
	switch v.extent {
	case abs.INCLUSIVE:
	case abs.LEFT:
	case abs.RIGHT:
	case abs.EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid interval range extent: %v", v.extent))
	}

	// Validate the endpoints.
	var first = v.effectiveFirst()
	var last = v.effectiveLast()
	v.size = last - first + 1
	if cox.RankValues(first, last) > 0 {
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
	var first = v.first.AsInteger()
	switch v.extent {
	case abs.RIGHT, abs.EXCLUSIVE:
		first++
	}
	return first
}

// This method returns the effective last value in the interval range.
func (v *interval[V]) effectiveLast() int {
	var last = v.last.AsInteger()
	switch v.extent {
	case abs.LEFT, abs.EXCLUSIVE:
		last--
	}
	return last
}

// This method returns the value associated with the specified index.
func (v *interval[V]) indexToValue(index int) V {
	var template V
	var primitive abs.Primitive = template
	switch primitive.(type) {
	case ele.Duration:
		primitive = ele.Duration(index)
	case ele.Moment:
		primitive = ele.Moment(index)
	case Rune:
		primitive = Rune(index)
	}
	return primitive.(V)
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

// DISCRETE INTERFACE

// This method returns a boolean value for this rune.
func (v Rune) AsBoolean() bool {
	return v > -1
}

// This method returns an integer value for this rune.
func (v Rune) AsInteger() int {
	return int(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this rune is negative.
func (v Rune) IsNegative() bool {
	return v < 0
}
