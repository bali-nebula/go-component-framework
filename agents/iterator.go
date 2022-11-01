/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an iterator that can be used to
// traverse the values in the specified array.
func Iterator[T abs.ValueLike](sequence abs.Sequential[T]) abs.IteratorLike[T] {
	var values = sequence.AsArray() // The returned array is immutable.
	var size = len(values)
	var slot = 0
	return &iterator[T]{values, size, slot}
}

// This type defines the structure and methods for a sequence iterator. The
// iterator operates on a sequence of values and is backed by a native Go array.
type iterator[T abs.ValueLike] struct {
	values []T // The array of values is immutable.
	size   int // So we can safely cache the size.
	slot   int // The default slot is zero.
}

// This method returns the current slot between values that this iterator is
// currently locked into.
func (v *iterator[T]) GetSlot() int {
	return v.slot
}

// This method moves this iterator to the specified slot between values.
func (v *iterator[T]) ToSlot(slot int) {
	if slot > v.size {
		slot = v.size
	}
	if slot < -v.size {
		slot = -v.size
	}
	if slot < 0 {
		slot = slot + v.size + 1
	}
	v.slot = slot
}

// This method moves this iterator to the slot before the first value.
func (v *iterator[T]) ToStart() {
	v.slot = 0
}

// This method moves this iterator to the slot after the last value.
func (v *iterator[T]) ToEnd() {
	v.slot = v.size
}

// This method determines whether or not there is an value before the current
// slot.
func (v *iterator[T]) HasPrevious() bool {
	return v.slot > 0
}

// This method retrieves the value before the current slot.
func (v *iterator[T]) GetPrevious() T {
	var result T
	if v.slot > 0 {
		result = v.values[v.slot-1] // convert to ZERO based indexing
		v.slot = v.slot - 1
	}
	return result
}

// This method determines whether or not there is an value after the current
// slot.
func (v *iterator[T]) HasNext() bool {
	return v.slot < v.size
}

// This method retrieves the value after the current slot.
func (v *iterator[T]) GetNext() T {
	var result T
	if v.slot < v.size {
		v.slot = v.slot + 1
		result = v.values[v.slot-1] // convert to ZERO based indexing
	}
	return result
}
