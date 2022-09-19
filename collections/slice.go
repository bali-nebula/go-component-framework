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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"reflect"
)

// SLICE IMPLEMENTATION

// This constructor creates a new slice of items covering the specified endpoints.
func Slice[T any](first T, connector string, last T) abstractions.SliceLike[T] {
	switch connector {
	case "..":
	case "<..":
	case "<..<":
	case "..<":
	default:
		panic(fmt.Sprintf("Received an invalid slice connector: %v", connector))
	}
	var v = slice[T]{first: first, connector: connector, last: last}
	v.size = v.calculateSize()
	return &v
}

// This type defines the structure and methods associated with a slice of items.
// This type is parameterized as follows:
//   - T is any type of primitive item.
type slice[T any] struct {
	first     T
	connector string
	last      T
	size      int
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this slice is empty.
func (v *slice[T]) IsEmpty() bool {
	return v.size == 0
}

// This method returns the number of items contained in this slice.
func (v *slice[T]) GetSize() int {
	return v.size
}

// This method returns up to the first 256 items in this slice. The items
// retrieved are in the same order as they are in the slice.
func (v *slice[T]) AsArray() []T {
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

// This method retrieves from this slice the item that is associated with the
// specified index.
func (v *slice[T]) GetItem(index int) T {
	var offset = abstractions.NormalizedIndex(index, v.size)
	var first = v.effectiveFirst()
	var item = v.indexToItem(first + offset)
	return item
}

// This method retrieves from this slice all items from the first index through
// the last index (inclusive).
func (v *slice[T]) GetItems(first int, last int) []T {
	var items []T
	for index := first; index <= last; index++ {
		var item = v.indexToItem(index)
		items = append(items, item)
	}
	return items
}

// This method returns the index of the FIRST occurence of the specified item in
// this slice, or zero if this slice does not contain the item.
func (v *slice[T]) GetIndex(item T) int {
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

// This method returns the first item in this slice.
func (v *slice[T]) GetFirst() T {
	return v.first
}

// This method sets the first item in this slice and returns the previous value.
func (v *slice[T]) SetFirst(item T) T {
	var old = v.first
	v.first = item
	return old
}

// This method returns the connector for this slice.
func (v *slice[T]) GetConnector() string {
	return v.connector
}

// This method sets the connector for this slice and returns the previous value.
func (v *slice[T]) SetConnector(connector string) string {
	var old = v.connector
	v.connector = connector
	return old
}

// This method returns the last item in this slice.
func (v *slice[T]) GetLast() T {
	return v.last
}

// This method sets the last item in this slice and returns the previous value.
func (v *slice[T]) SetLast(item T) T {
	var old = v.last
	v.last = item
	return old
}

// This method determines whether or not this slice can be enumerated over.
func (v *slice[T]) IsEnumerable() bool {
	return v.size > 0
}

// PRIVATE INTERFACE

// This function converts the type of the specified index to the parameterized
// type T.
func (v *slice[T]) indexToItem(index int) T {
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
func (v *slice[T]) itemToIndex(item T) int {
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
	case elements.Number:
		return int(real(index))
	default:
		panic(fmt.Sprintf("The item is not an integer type: %T", item))
	}
}

// This function returns the effective first item in the slice based on its
// connector type. It can only be called when the parameterized item type T is
// an integer type.
func (v *slice[T]) effectiveFirst() int {
	var first = v.itemToIndex(v.first)
	switch v.connector {
	case "<..":
		first++
	case "<..<":
		first++
	}
	return first
}

// This function returns the effective last item in the slice based on its
// connector type. It can only be called when the parameterized item type T is
// an integer type.
func (v *slice[T]) effectiveLast() int {
	var last = v.itemToIndex(v.last)
	switch v.connector {
	case "..<":
		last--
	case "<..<":
		last--
	}
	return last
}

// This function calulates the effective size of the slice assuming the
// parameterized type T is an integer type. If it is not an integer type
// the effective size is zero.
func (v *slice[T]) calculateSize() int {
	switch reflect.ValueOf(v.first).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Complex64, reflect.Complex128:
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

// SLICES LIBRARY

// This constructor creates a new slices library for the specified generic
// item type.
func Slices[T any]() *slices[T] {
	return &slices[T]{}
}

// This type defines the library functions that operate on slices. Since
// slices have a parameterized item type this library type is also
// parameterized as follows:
//   - T is any type of primitive item.
type slices[T any] struct{}

// SPLICEABLE INTERFACE

// This library function removes the specified slice of items from the specified
// list and returns it as an array of items.
func (l *slices[T]) Cut(collection abstractions.ListLike[T], slice abstractions.Elastic[T]) []T {
	var firstIndex, lastIndex = l.effectiveIndices(collection, slice)
	var items = collection.RemoveItems(firstIndex, lastIndex)
	return items
}

// This library function replaces the specified splice of items in the specified
// list with the specified items.
func (l *slices[T]) Splice(collection abstractions.ListLike[T], slice abstractions.Elastic[T], items []T) []T {
	var firstIndex, lastIndex = l.effectiveIndices(collection, slice)
	var removed = collection.RemoveItems(firstIndex, lastIndex)
	var slot = firstIndex - 1 // This is the slot before the first index.
	collection.InsertItems(slot, items)
	return removed
}

// PRIVATE INTERFACE

// This library function returns the effective first and last indices for the
// specified slice within the specified list.
func (l *slices[T]) effectiveIndices(c abstractions.ListLike[T], s abstractions.Elastic[T]) (first, last int) {
	var firstItem = s.GetFirst()
	var lastItem = s.GetLast()
	var firstIndex = c.GetIndex(firstItem)
	var lastIndex = c.GetIndex(lastItem)
	var connector = s.GetConnector()
	switch connector {
	case "<..":
		firstIndex++
	case "<..<":
		firstIndex++
		lastIndex--
	case "..<":
		lastIndex--
	default:
		// No changes needed.
	}
	return firstIndex, lastIndex
}
