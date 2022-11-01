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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	ref "reflect"
)

// RANGE IMPLEMENTATION

// This constructor creates a new range of values covering the specified endpoints.
func Range[T abs.ValueLike](first T, extent abs.Extent, last T) abs.RangeLike[T] {
	switch extent {
	case abs.INCLUSIVE:
	case abs.LEFT:
	case abs.RIGHT:
	case abs.EXCLUSIVE:
	default:
		panic(fmt.Sprintf("Received an invalid range extent: %v", extent))
	}
	var v = ranje[T]{first: first, extent: extent, last: last}
	v.size = v.calculateSize()
	return &v
}

// This type defines the structure and methods associated with a range of values.
// This type is parameterized as follows:
//   - T is any primitive type.
type ranje[T abs.ValueLike] struct {
	first  T
	extent abs.Extent
	last   T
	size   int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this range is empty.
func (v *ranje[T]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of values contained in this range.
func (v *ranje[T]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 values in this range. The values
// retrieved are in the same order as they are in the range.
func (v *ranje[T]) AsArray() []T {
	var size = v.size
	if size > 256 {
		size = 256
	}
	var array = make([]T, size)
	if size > 0 {
		var first = v.effectiveFirst()
		for i := 0; i < size; i++ {
			var value = v.indexToValue(first + i)
			array[i] = value
		}
	}
	return array
}

// INDEXED INTERFACE

// This method sets the comparer function for this list.
func (v *ranje[T]) SetComparer(compare abs.ComparisonFunction) {
	// This method does nothing since range values are either sequential or only
	// have meaning within the context of a set. We may revisit this...
}

// This method retrieves from this range the value that is associated with the
// specified index.
func (v *ranje[T]) GetValue(index int) T {
	var offset = abs.NormalizedIndex(index, v.size)
	var first = v.effectiveFirst()
	var value = v.indexToValue(first + offset)
	return value
}

// This method retrieves from this range all values from the first index through
// the last index (inclusive).
func (v *ranje[T]) GetValues(first int, last int) abs.Sequential[T] {
	var values = List[T]()
	for index := first; index <= last; index++ {
		var value = v.indexToValue(index)
		values.AddValue(value)
	}
	return values
}

// This method returns the index of the FIRST occurence of the specified value in
// this range, or zero if this range does not contain the value.
func (v *ranje[T]) GetIndex(value T) int {
	if v.size == 0 {
		return 0
	}
	var index = v.valueToIndex(value)
	var first = v.effectiveFirst()
	var offset = index - first + 1
	if offset < 0 || offset > v.size {
		return 0
	}
	return offset
}

// ELASTIC INTERFACE

// This method determines whether or not this range can be enumerated over.
func (v *ranje[T]) IsEnumerable() bool {
	return v.size > 0
}

// This method returns the first value in this range.
func (v *ranje[T]) GetFirst() T {
	return v.first
}

// This method sets the first value in this range.
func (v *ranje[T]) SetFirst(value T) {
	v.first = value
}

// This method returns the extent for this range.
func (v *ranje[T]) GetExtent() abs.Extent {
	return v.extent
}

// This method sets the extent for this range.
func (v *ranje[T]) SetExtent(extent abs.Extent) {
	v.extent = extent
}

// This method returns the last value in this range.
func (v *ranje[T]) GetLast() T {
	return v.last
}

// This method sets the last value in this range.
func (v *ranje[T]) SetLast(value T) {
	v.last = value
}

// PRIVATE INTERFACE

// This function converts the type of the specified index to the parameterized
// type T.
func (v *ranje[T]) indexToValue(index int) T {
	var template T
	var value abs.ValueLike = template
	switch value.(type) {
	case uint8:
		value = uint8(index)
	case uint16:
		value = uint16(index)
	case uint32:
		value = uint32(index)
	case int8:
		value = int8(index)
	case int16:
		value = int16(index)
	case int32:
		value = int32(index)
	case int:
		value = int(index)
	default:
		panic(fmt.Sprintf("The value is not an integer type: %T", template))
	}
	return value.(T)
}

// This function converts the type of the specified value to an integer type.
func (v *ranje[T]) valueToIndex(value T) int {
	var template abs.ValueLike = value
	switch index := template.(type) {
	case uint8:
		return int(index)
	case uint16:
		return int(index)
	case uint32:
		return int(index)
	case int8:
		return int(index)
	case int16:
		return int(index)
	case int32:
		return int(index)
	case int64:
		return int(index)
	case int:
		return int(index)
	case complex64:
		return int(real(index))
	case complex128:
		return int(real(index))
	case ele.Number:
		return int(real(index))
	default:
		panic(fmt.Sprintf("The value is not an integer type: %T", value))
	}
}

// This function returns the effective first value in the range based on its
// extent type. It can only be called when the parameterized value type T is
// an integer type.
func (v *ranje[T]) effectiveFirst() int {
	var first = v.valueToIndex(v.first)
	switch v.extent {
	case abs.RIGHT:
		first++
	case abs.EXCLUSIVE:
		first++
	}
	return first
}

// This function returns the effective last value in the range based on its
// extent type. It can only be called when the parameterized value type T is
// an integer type.
func (v *ranje[T]) effectiveLast() int {
	var last = v.valueToIndex(v.last)
	switch v.extent {
	case abs.LEFT:
		last--
	case abs.EXCLUSIVE:
		last--
	}
	return last
}

// This function calulates the effective size of the range assuming the
// parameterized type T is an integer type. If it is not an integer type
// the effective size is zero.
func (v *ranje[T]) calculateSize() int {
	switch ref.ValueOf(v.first).Kind() {
	case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32,
		ref.Int, ref.Int8, ref.Int16, ref.Int32,
		ref.Complex64, ref.Complex128:
		var first = v.effectiveFirst()
		var last = v.effectiveLast()
		var size = last - first + 1
		if size < 0 {
			size = 0
		}
		return size
	default:
		return 0
	}
}
