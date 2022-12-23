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

// RANGE IMPLEMENTATION

// This constructor creates a new range of values covering the specified
// endpoints. Note that at least one of the endpoints must be non-nil so that
// the numeric type may be determined.
func Range[V abs.Numeric](first V, extent abs.Extent, last V) abs.RangeLike[V] {
	var v = ranje[V]{first: first, extent: extent, last: last}
	v.validateRange()
	return &v
}

// This type defines the structure and methods associated with a range of values.
// This type is parameterized as follows:
//   - V is any primitive type.
type ranje[V abs.Numeric] struct {
	first  V
	extent abs.Extent
	last   V
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this range is empty.
func (v *ranje[V]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this range.
func (v *ranje[V]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this range. The values
// retrieved are in the same order as they are in the range.
func (v *ranje[V]) AsArray() []V {
	var size = 0
	if v.IsEnumerable() {
		size = v.size
		if size > 256 {
			size = 256
		}
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

// This method retrieves from this range the value that is associated with the
// specified index.
func (v *ranje[V]) GetValue(index int) V {
	if v.IsEnumerable() {
		var offset = v.GoIndex(index)
		if offset < 0 {
			panic("The index is outside the range of values.")
		}
		var first = v.firstIndex()
		var value = v.indexToValue(first + offset)
		return value
	}
	panic("Cannot retrieve a value from a continuous range.")
}

// This method retrieves from this range all values from the first index through
// the last index (inclusive).
func (v *ranje[V]) GetValues(first int, last int) col.Sequential[V] {
	if v.IsEnumerable() {
		var values = col.List[V]()
		for index := first; index <= last; index++ {
			var value V = v.indexToValue(index)
			values.AddValue(value)
		}
		return values
	}
	panic("Cannot retrieve values from a continuous range.")
}

// This method returns the index of the FIRST occurence of the specified value in
// this range, or zero if this range does not contain the value.
func (v *ranje[V]) GetIndex(value V) int {
	if v.IsEnumerable() {
		var index = v.valueToIndex(value)
		var first = v.firstIndex()
		var offset = index - first + 1
		if offset < 0 || offset > v.size {
			return 0
		}
		return offset
	}
	panic("Cannot retrieve the index of a value from a continuous range.")
}

// ELASTIC INTERFACE

// This method determines whether or not this range can be enumerated over.
func (v *ranje[V]) IsEnumerable() bool {
	return v.size > 0
}

// This method returns the first value in this range.
func (v *ranje[V]) GetFirst() V {
	return v.first
}

// This method sets the first value in this range.
func (v *ranje[V]) SetFirst(value V) {
	v.first = value
	v.validateRange()
}

// This method returns the extent for this range.
func (v *ranje[V]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this range.
func (v *ranje[V]) SetExtent(extent abs.Extent) {
	v.extent = extent
	v.validateRange()
}

// This method returns the last value in this range.
func (v *ranje[V]) GetLast() V {
	return v.last
}

// This method sets the last value in this range.
func (v *ranje[V]) SetLast(value V) {
	v.last = value
	v.validateRange()
}

// PRIVATE INTERFACE

// This method determines whether or not the first and last endpoints are
// invalid.
func (v *ranje[V]) validateRange() {
	switch v.extent {
	case abs.INCLUSIVE:
	case abs.LEFT:
	case abs.RIGHT:
	case abs.EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid range extent: %v", v.extent))
	}
	if !ref.ValueOf(v.first).IsValid() {
		v.first = v.GetMinimum()
	}
	if !ref.ValueOf(v.last).IsValid() {
		v.last = v.GetMaximum()
	}
	if v.first.IsDiscrete() {
		var firstIndex = v.firstIndex()
		var lastIndex = v.lastIndex()
		if col.RankValues(firstIndex, lastIndex) > 0 {
			fmt.Printf("first: %v, last: %v\n", firstIndex, lastIndex)
			panic("The effective first value in the range must not be more than the effective last value.")
		}
		v.size = lastIndex - firstIndex + 1
	} else {
		var firstValue = v.firstValue()
		var lastValue = v.lastValue()
		if col.RankValues(firstValue, lastValue) > 0 {
			fmt.Printf("first: %v, last: %v\n", firstValue, lastValue)
			panic("The first value in the range must not be more than the last value.")
		}
		v.size = -1
	}
}

// This method converts the type of the specified index to the parameterized
// value type.
func (v *ranje[V]) indexToValue(index int) V {
	var template V
	var value abs.Numeric = template
	switch value.(type) {
	case ele.Angle:
		value = ele.Angle(index)
	case ele.Duration:
		value = ele.Duration(index)
	case ele.Moment:
		value = ele.Moment(index)
	case Rune:
		value = Rune(index)
	default:
		panic(fmt.Sprintf("The value is not a numeric type: %T", template))
	}
	return value.(V)
}

// This method converts the type of the specified value to an index.
func (v *ranje[V]) valueToIndex(value V) int {
	var index = value.AsInteger()
	return index
}

// This method returns the effective first value in the range based on its
// extent type. It can only be called when the parameterized value type V is
// an integer type.
func (v *ranje[V]) firstIndex() int {
	var firstIndex = v.valueToIndex(v.first)
	switch v.extent {
	case abs.RIGHT:
		firstIndex++
	case abs.EXCLUSIVE:
		firstIndex++
	}
	return firstIndex
}

// This method returns the effective last value in the range based on its
// extent type. It can only be called when the parameterized value type V is
// an integer type.
func (v *ranje[V]) lastIndex() int {
	var lastIndex = v.valueToIndex(v.last)
	switch v.extent {
	case abs.LEFT:
		lastIndex--
	case abs.EXCLUSIVE:
		lastIndex--
	}
	return lastIndex
}

// This method returns the first value in the continuous range as a float64 type.
func (v *ranje[V]) firstValue() float64 {
	return v.first.AsReal()
}

// This method returns the last value in the continuous range as a float64 type.
func (v *ranje[V]) lastValue() float64 {
	return v.last.AsReal()
}

// This method normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func (v *ranje[V]) GoIndex(index int) int {
	if v.IsEnumerable() {
		var length = v.size
		switch {
		case index < -length || index == 0 || index > length:
			// The index is outside the bounds of the specified range.
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
	panic("Cannot retrieve the index of a value from a continuous range.")
}

func (v *ranje[V]) GetMinimum() V {
	var value abs.Numeric = v.last
	switch value.(type) {
	case ele.Angle:
		value = ele.MinimumAngle()
	case ele.Duration:
		value = ele.MinimumDuration()
	case ele.Moment:
		value = ele.MinimumMoment()
	case ele.Percentage:
		value = ele.MinimumPercentage()
	case ele.Probability:
		value = ele.MinimumProbability()
	case Real:
		value = MinimumReal()
	case Rune:
		value = MinimumRune()
	default:
		panic(fmt.Sprintf("The value is not a numeric type: %T", value))
	}
	return value.(V)
}

func (v *ranje[V]) GetMaximum() V {
	var value abs.Numeric = v.first
	switch value.(type) {
	case ele.Angle:
		value = ele.MaximumAngle()
	case ele.Duration:
		value = ele.MaximumDuration()
	case ele.Moment:
		value = ele.MaximumMoment()
	case ele.Percentage:
		value = ele.MaximumPercentage()
	case ele.Probability:
		value = ele.MaximumProbability()
	case Real:
		value = MaximumReal()
	case Rune:
		value = MaximumRune()
	default:
		panic(fmt.Sprintf("The value is not a numeric type: %T", value))
	}
	return value.(V)
}
