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

// This constructor creates a new range of items covering the specified endpoints.
func Range[T any](first T, connector string, last T) abs.RangeLike[T] {
	switch connector {
	case "..":
	case "<..":
	case "<..<":
	case "..<":
	default:
		panic(fmt.Sprintf("Received an invalid range connector: %v", connector))
	}
	var v = ranje[T]{first: first, connector: connector, last: last}
	v.size = v.calculateSize()
	return &v
}

// This type defines the structure and methods associated with a range of items.
// This type is parameterized as follows:
//   - T is any type of primitive item.
type ranje[T any] struct {
	first     T
	connector string
	last      T
	size      int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this range is empty.
func (v *ranje[T]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of items contained in this range.
func (v *ranje[T]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 items in this range. The items
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
			var item = v.indexToItem(first + i)
			array[i] = item
		}
	}
	return array
}

// INDEXED INTERFACE

// This method sets the comparer function for this list.
func (v *ranje[T]) SetComparer(compare abs.ComparisonFunction) {
	// This method does nothing since range items are either sequential or only
	// have meaning within the context of a set. We may revisit this...
}

// This method retrieves from this range the item that is associated with the
// specified index.
func (v *ranje[T]) GetItem(index int) T {
	var offset = abs.NormalizedIndex(index, v.size)
	var first = v.effectiveFirst()
	var item = v.indexToItem(first + offset)
	return item
}

// This method retrieves from this range all items from the first index through
// the last index (inclusive).
func (v *ranje[T]) GetItems(first int, last int) abs.Sequential[T] {
	var items = List[T]()
	for index := first; index <= last; index++ {
		var item = v.indexToItem(index)
		items.AddItem(item)
	}
	return items
}

// This method returns the index of the FIRST occurence of the specified item in
// this range, or zero if this range does not contain the item.
func (v *ranje[T]) GetIndex(item T) int {
	if v.size == 0 {
		return 0
	}
	var index = v.itemToIndex(item)
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

// This method returns the first item in this range.
func (v *ranje[T]) GetFirst() T {
	return v.first
}

// This method sets the first item in this range.
func (v *ranje[T]) SetFirst(item T) {
	v.first = item
}

// This method returns the connector for this range.
func (v *ranje[T]) GetConnector() string {
	return v.connector
}

// This method sets the connector for this range.
func (v *ranje[T]) SetConnector(connector string) {
	v.connector = connector
}

// This method returns the last item in this range.
func (v *ranje[T]) GetLast() T {
	return v.last
}

// This method sets the last item in this range.
func (v *ranje[T]) SetLast(item T) {
	v.last = item
}

// PRIVATE INTERFACE

// This function converts the type of the specified index to the parameterized
// type T.
func (v *ranje[T]) indexToItem(index int) T {
	var template T
	var value any = template
	switch value.(type) {
	case uint8:
		var item any = uint8(index)
		return item.(T)
	case uint16:
		var item any = uint16(index)
		return item.(T)
	case uint32:
		var item any = uint32(index)
		return item.(T)
	case int8:
		var item any = int8(index)
		return item.(T)
	case int16:
		var item any = int16(index)
		return item.(T)
	case int32:
		var item any = int32(index)
		return item.(T)
	case int:
		var item any = int(index)
		return item.(T)
	default:
		panic(fmt.Sprintf("The item is not an integer type: %T", template))
	}
}

// This function converts the type of the specified item to an integer type.
func (v *ranje[T]) itemToIndex(item T) int {
	var value any = item
	switch index := value.(type) {
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
		panic(fmt.Sprintf("The item is not an integer type: %T", item))
	}
}

// This function returns the effective first item in the range based on its
// connector type. It can only be called when the parameterized item type T is
// an integer type.
func (v *ranje[T]) effectiveFirst() int {
	var first = v.itemToIndex(v.first)
	switch v.connector {
	case "<..":
		first++
	case "<..<":
		first++
	}
	return first
}

// This function returns the effective last item in the range based on its
// connector type. It can only be called when the parameterized item type T is
// an integer type.
func (v *ranje[T]) effectiveLast() int {
	var last = v.itemToIndex(v.last)
	switch v.connector {
	case "..<":
		last--
	case "<..<":
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
