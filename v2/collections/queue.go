/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// QUEUE IMPLEMENTATION

// This constructor creates a new empty value queue with the default
// capacity.
func Queue() abs.QueueLike {
	var v = col.Queue[abs.ComponentLike]()
	return &queue{v}
}

// This constructor creates a new empty value queue with the specified
// capacity.
func QueueWithCapacity(capacity int) abs.QueueLike {
	var v = col.QueueWithCapacity[abs.ComponentLike](capacity)
	return &queue{v}
}

// This constructor creates a new value queue from the specified sequence.
func QueueFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.QueueLike {
	var v = col.QueueWithCapacity[abs.ComponentLike](sequence.GetSize() * 2)
	var iterator = col.Iterator[abs.ComponentLike](sequence)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.AddValue(value)
	}
	return &queue{v}
}

// This type defines the structure and methods associated with a value
// queue.
type queue struct {
	values col.QueueLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this queue is empty.
func (v *queue) IsEmpty() bool {
	return v.values.IsEmpty()
}

// This method returns the number of values contained in this queue.
func (v *queue) GetSize() int {
	return v.values.GetSize()
}

// This method returns all the values in this queue. The values retrieved are in
// the same order as they are in the queue.
func (v *queue) AsArray() []abs.ComponentLike {
	return v.values.AsArray()
}

// FIFO INTERFACE

// This method retrieves the capacity of this queue.
func (v *queue) GetCapacity() int {
	return v.values.GetCapacity()
}

// This method appends the specified value to the end of this queue.
func (v *queue) AddValue(value abs.ComponentLike) {
	v.values.AddValue(value)
}

// This method removes from this queue the value that is at the head of it. It
// returns the removed value and a "comma ok" value as the result.
func (v *queue) RemoveHead() (abs.ComponentLike, bool) {
	return v.values.RemoveHead()
}

// This method closes the queue so no more values can be placed on it.
func (v *queue) CloseQueue() {
	v.values.CloseQueue()
}
