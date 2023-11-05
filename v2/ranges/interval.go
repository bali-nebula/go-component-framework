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
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	col "github.com/craterdog/go-collection-framework/v2"
)

// INTERVAL IMPLEMENTATION

// This constructor creates a new interval range of values covering the
// specified endpoints.
func Interval(first abs.Discrete, extent abs.Extent, last abs.Discrete) abs.IntervalLike {
	var v = interval{first: first, extent: extent, last: last}
	v.validateInterval()
	return &v
}

// This type defines the structure and methods associated with an interval range
// of values.
type interval struct {
	first  abs.Discrete
	extent abs.Extent
	last   abs.Discrete
	size   int
}

// BOUNDED INTERFACE

// This method returns the first value in this interval range.
func (v *interval) GetFirst() abs.Discrete {
	return v.first
}

// This method sets the first value in this interval range.
func (v *interval) SetFirst(value abs.Discrete) {
	v.first = value
	v.validateInterval()
}

// This method returns the extent for this interval range.
func (v *interval) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this interval range.
func (v *interval) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateInterval()
}

// This method returns the last value in this interval range.
func (v *interval) GetLast() abs.Discrete {
	return v.last
}

// This method sets the last value in this interval range.
func (v *interval) SetLast(value abs.Discrete) {
	v.last = value
	v.validateInterval()
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this interval range is empty.
func (v *interval) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this interval range.
func (v *interval) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this interval range. The
// values retrieved are in the same order as they are in the interval range.
func (v *interval) AsArray() []abs.Discrete {
	var size = v.size
	if size > 256 {
		size = 256
	}
	var array = make([]abs.Discrete, size)
	if size > 0 {
		var first = v.firstIndex()
		for i := 0; i < size; i++ {
			var value = v.indexToValue(first + i)
			array[i] = value
		}
	}
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this interval range the value that is associated
// with the specified index.
func (v *interval) GetValue(index int) abs.Discrete {
	var offset = v.goIndex(index)
	if offset < 0 {
		panic("The index is outside the interval range of values.")
	}
	var first = v.firstIndex()
	var value = v.indexToValue(first + offset)
	return value
}

// This method retrieves from this interval range all values from the first
// index through the last index (inclusive).
func (v *interval) GetValues(first int, last int) abs.Sequential[abs.Discrete] {
	var values = col.List[abs.Discrete]()
	for index := first; index <= last; index++ {
		var value = v.indexToValue(index)
		values.AddValue(value)
	}
	return values
}

// SEARCHABLE INTERFACE

// This method returns the index of the FIRST occurence of the specified value
// in this interval range, or zero if this interval range does not contain the
// value.
func (v *interval) GetIndex(value abs.Discrete) int {
	var index = value.AsInteger()
	var first = v.firstIndex()
	var offset = index - first + 1
	if offset < 0 || offset > v.size {
		return 0
	}
	return offset
}

// This method determines whether or not the specified value is included in this
// interval range.
func (v *interval) ContainsValue(value abs.Discrete) bool {
	var first = v.first.AsInteger()
	var candidate = value.AsInteger()
	var last = v.last.AsInteger()
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
		var message = fmt.Sprintf("Received an invalid interval range extent: %v", v.extent)
		panic(message)
	}
}

// This method determines whether or not this interval contains ANY of the
// specified values.
func (v *interval) ContainsAny(values abs.Sequential[abs.Discrete]) bool {
	var iterator = col.Iterator[abs.Discrete](values)
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
func (v *interval) ContainsAll(values abs.Sequential[abs.Discrete]) bool {
	var iterator = col.Iterator[abs.Discrete](values)
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
func (v *interval) validateInterval() {
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
	if col.RankValues(first, last) > 0 {
		panic("The effective first value in the interval range must not be more than the effective last value.")
	}

}

// This method returns the effective first value in the interval range based on
// its extent type.
func (v *interval) firstIndex() int {
	var firstIndex = v.effectiveFirst()
	return firstIndex
}

// This method returns the effective first value in the interval range.
func (v *interval) effectiveFirst() int {
	var first = v.first.AsInteger()
	switch v.extent {
	case abs.RIGHT, abs.EXCLUSIVE:
		first++
	}
	return first
}

// This method returns the effective last value in the interval range.
func (v *interval) effectiveLast() int {
	var last = v.last.AsInteger()
	switch v.extent {
	case abs.LEFT, abs.EXCLUSIVE:
		last--
	}
	return last
}

// This method returns the value associated with the specified index.
func (v *interval) indexToValue(index int) abs.Discrete {
	var discrete abs.Discrete = v.first
	switch discrete.(type) {
	case abs.DurationLike:
		discrete = ele.DurationFromInt(index)
	case abs.MomentLike:
		discrete = ele.MomentFromInt(index)
	case abs.IntegerLike:
		discrete = ele.Integer(index)
	case abs.CharacterLike:
		discrete = ele.Character(index)
	default:
		var message = fmt.Sprintf("The discrete type was not found: %T", discrete)
		panic(message)
	}
	return discrete
}

// This method normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func (v *interval) goIndex(index int) int {
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
