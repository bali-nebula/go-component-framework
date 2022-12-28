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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	ele "github.com/craterdog-bali/go-component-framework/elements"
	col "github.com/craterdog/go-collection-framework"
	mat "math"
	ref "reflect"
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
// so that the primitive type may be determined.
func Interval[V abs.Primitive](first V, extent abs.Extent, last V) abs.IntervalLike[V] {
	var v = interval[V]{first: first, extent: extent, last: last}
	v.validateInterval()
	return &v
}

// This type defines the structure and methods associated with an interval range
// of values. This type is parameterized as follows:
//   - V is any primitive type.
type interval[V abs.Primitive] struct {
	first  V
	extent abs.Extent
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
func (v *interval[V]) GetValues(first int, last int) col.Sequential[V] {
	var values = col.List[V]()
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
	var index = v.valueToIndex(value)
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

// PRIVATE INTERFACE

// This method returns the minimum possible endpoint value for this interval.
func (v *interval[V]) getMinimum() V {
	var value abs.Primitive = v.last
	switch value.(type) {
	case ele.Duration:
		value = ele.MinimumDuration()
	case ele.Moment:
		value = ele.MinimumMoment()
	case Rune:
		value = MinimumRune()
	default:
		panic(fmt.Sprintf("The value is not a supported type: %T", value))
	}
	return value.(V)
}

// This method returns the maximum possible endpoint value for this interval.
func (v *interval[V]) getMaximum() V {
	var value abs.Primitive = v.first
	switch value.(type) {
	case ele.Duration:
		value = ele.MaximumDuration()
	case ele.Moment:
		value = ele.MaximumMoment()
	case Rune:
		value = MaximumRune()
	default:
		panic(fmt.Sprintf("The value is not a supported type: %T", value))
	}
	return value.(V)
}

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

	// Set default endpoints if necessary.
	if !ref.ValueOf(v.first).IsValid() {
		v.first = v.getMinimum()
	}
	if !ref.ValueOf(v.last).IsValid() {
		v.last = v.getMaximum()
	}

	// Calculate the effective size.
	var first = v.effectiveFirst()
	var last = v.effectiveLast()
	switch first.(type) {
	case int:
		v.size = last.(int) - first.(int) + 1
	case float64, string:
		v.size = -1
	}

	// Validate the endpoints.
	if col.RankValues(first, last) > 0 {
		fmt.Printf("first: %v, last: %v\n", first, last)
		panic("The effective first value in the interval range must not be more than the effective last value.")
	}

}

// This method converts the type of the specified index to the parameterized
// value type.
func (v *interval[V]) indexToValue(index int) V {
	var template V
	var value abs.Primitive = template
	switch value.(type) {
	case ele.Duration:
		value = ele.Duration(index)
	case ele.Moment:
		value = ele.Moment(index)
	case Rune:
		value = Rune(index)
	}
	return value.(V)
}

// This method converts the type of the specified value to an index.
func (v *interval[V]) valueToIndex(value V) int {
	var r any = v.asNative(value)
	var index = r.(int)
	return index
}

// This method returns the effective first value in the interval range based on
// its extent type. It can only be called when the parameterized value type V is
// an integer type.
func (v *interval[V]) firstIndex() int {
	var firstIndex = v.effectiveFirst()
	return firstIndex.(int)
}

// This method returns the effective first value in the interval range.
func (v *interval[V]) effectiveFirst() any {
	var first = v.asNative(v.first)
	switch first.(type) {
	case int:
		switch v.extent {
		case abs.RIGHT, abs.EXCLUSIVE:
			first = first.(int) + 1
		}
	}
	return first
}

// This method returns the effective last value in the interval range.
func (v *interval[V]) effectiveLast() any {
	var last = v.asNative(v.last)
	switch last.(type) {
	case int:
		switch v.extent {
		case abs.LEFT, abs.EXCLUSIVE:
			last = last.(int) - 1
		}
	}
	return last
}

// This method returns the Go native value for the specified primitive value.
func (v *interval[V]) asNative(primitive abs.Primitive) any {
	var native any
	switch value := primitive.(type) {
	case ele.Duration:
		native = int(ele.Duration(value))
	case ele.Moment:
		native = int(ele.Moment(value))
	case Rune:
		native = int(Rune(value))
	default:
		panic(fmt.Sprintf("The endpoint is not a supported type: %T", value))
	}
	return native
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
