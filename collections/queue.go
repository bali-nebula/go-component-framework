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
	"sync"
)

// QUEUE IMPLEMENTATION

// This constructor creates a new empty queue with the default capacity.
// The default capacity is 16 items.
func Queue[T any]() abstractions.QueueLike[T] {
	return QueueWithCapacity[T](0)
}

// This constructor creates a new empty queue with the specified capacity.
func QueueWithCapacity[T any](capacity int) abstractions.QueueLike[T] {
	// Groom the arguments.
	if capacity < 1 {
		capacity = 16 // The default value.
	}

	// Return an empty queue.
	var available = make(chan bool, capacity)
	var items = List[T]()
	return &queue[T]{available: available, items: items}
}

// This type defines the structure and methods associated with a queue of items.
// A queue implements first-in-first-out semantics. It is generally used by
// multiple goroutines at the same time and therefore enforces synchronized
// access. This type is parameterized as follows:
//   - T is any type of item.
//
// If the go chan type ever supports snapshots of its state, the underlying list
// can be removed and the channel modified to pass the items instead of the
// availability. Currently, the underlying list is only required by the
// AsArray() method.
type queue[T any] struct {
	available chan bool
	items     abstractions.ListLike[T]
	mutex     sync.Mutex
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this queue is empty.
func (v *queue[T]) IsEmpty() bool {
	v.mutex.Lock()
	var result = len(v.available) == 0
	v.mutex.Unlock()
	return result
}

// This method returns the number of items contained in this queue.
func (v *queue[T]) GetSize() int {
	v.mutex.Lock()
	var result = len(v.available)
	v.mutex.Unlock()
	return result
}

// This method returns all the items in this queue. The items retrieved are in
// the same order as they are in the queue.
func (v *queue[T]) AsArray() []T {
	v.mutex.Lock()
	var result = v.items.AsArray()
	v.mutex.Unlock()
	return result
}

// FIFO INTERFACE

// This method retrieves the capacity of this queue.
func (v *queue[T]) GetCapacity() int {
	return cap(v.available) // The channel capacity is static.
}

// This method adds the specified item to the end of this queue.
func (v *queue[T]) AddItem(item T) {
	v.mutex.Lock()
	v.items.AddItem(item)
	v.mutex.Unlock()
	v.available <- true // Will block if at capacity.
}

// This method adds the specified items to the top of this queue.
func (v *queue[T]) AddItems(items []T) {
	for _, item := range items {
		v.mutex.Lock()
		v.items.AddItem(item)
		v.mutex.Unlock()
		v.available <- true // Will block if at capacity.
	}
}

// This method removes from this queue the item that is at the head of it. It
// returns the removed value and a "comma ok" value as the result.
func (v *queue[T]) RemoveHead() (head T, ok bool) {
	// Default the return value to the zero value for type T.

	// Remove the head item from the queue if one exists.
	_, ok = <-v.available // Will block until an item is available.
	if ok {
		v.mutex.Lock()
		head = v.items.RemoveItem(1)
		v.mutex.Unlock()
	}

	// Return the results
	return
}

// This method closes the queue so no more items can be placed on it.
func (v *queue[T]) CloseQueue() {
	v.mutex.Lock()
	close(v.available) // No more items can be placed on the queue.
	v.mutex.Unlock()
}

// QUEUES LIBRARY

// This constructor creates a new queues library for the specified generic
// item type.
func Queues[T any]() *queues[T] {
	return &queues[T]{}
}

// This type defines the library functions that operate on queues. Since
// queues have a parameterized item type this library type is also
// parameterized as follows:
//   - T is any type of item.
type queues[T any] struct{}

// ROUTEABLE INTERFACE

// This library function connects the output of the specified input queue with a
// number of new output queues specified by the size parameter and returns an
// array of the new output queues. Each item added to the input queue will be
// added automatically to ALL of the output queues. This pattern is useful when
// a set of DIFFERENT operations needs to occur for every item and each
// operation can be done in parallel.
func (l *queues[T]) Fork(wg *sync.WaitGroup, input abstractions.FIFO[T], size int) []abstractions.FIFO[T] {
	// Validate the arguments.
	if size < 1 {
		panic(fmt.Sprintf("The fan out size for a queue must be greater than zero: %v!", size))
	}

	// Create the new output queues.
	var capacity = input.GetCapacity()
	var outputs = make([]abstractions.FIFO[T], size)
	for index, _ := range outputs {
		outputs[index] = abstractions.FIFO[T](QueueWithCapacity[T](capacity))
	}

	// Connect up the input queue to the output queues in a separate goroutine.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Write each item read from the input queue to each output queue.
		var item, ok = input.RemoveHead() // Will block when empty.
		for ok {
			for _, output := range outputs {
				output.AddItem(item) // Will block when full.
			}
			item, ok = input.RemoveHead() // Will block when empty.
		}

		// Close all output queues.
		for _, output := range outputs {
			output.CloseQueue()
		}
	}()

	return outputs
}

// This library function connects the output of the specified input queue with
// the number of output queues specified by the size parameter and returns an
// array of the new output queues. Each item added to the input queue will be
// added automatically to ONE of the output queues. This pattern is useful when
// a SINGLE operation needs to occur for each item and the operation can be done
// on the items in parallel. The results can then be consolidated later on using
// the Join() function.
func (l *queues[T]) Split(wg *sync.WaitGroup, input abstractions.FIFO[T], size int) []abstractions.FIFO[T] {
	// Validate the arguments.
	if size < 1 {
		panic(fmt.Sprintf("The size of the split must be greater than zero: %v!", size))
	}

	// Create the new output queues.
	var capacity = input.GetCapacity()
	var outputs = make([]abstractions.FIFO[T], size)
	for index, _ := range outputs {
		outputs[index] = QueueWithCapacity[T](capacity)
	}

	// Connect up the input queue to the output queues.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Take turns reading from the input queue and writing to each output queue.
		var i = 0
		var item, ok = input.RemoveHead() // Will block when empty.
		for ok {
			outputs[i].AddItem(item) // Will block when full.
			i = (i + 1) % size
			item, ok = input.RemoveHead() // Will block when empty.
		}

		// Close all output queues.
		for _, output := range outputs {
			output.CloseQueue()
		}
	}()

	return outputs
}

// This library function connects the outputs of the specified array of input
// queues with a new output queue returns the new output queue. Each item
// removed from each input queue will automatically be added to the output
// queue. This pattern is useful when the results of the processing with a
// Split() function need to be consolicated into a single queue.
func (l *queues[T]) Join(wg *sync.WaitGroup, inputs []abstractions.FIFO[T]) abstractions.FIFO[T] {
	// Validate the arguments.
	var size = len(inputs)
	if size < 1 {
		panic(fmt.Sprintf("The number of input queues for a join must be greater than zero: %v!", size))
	}

	// Create the new output queue.
	var capacity = inputs[0].GetCapacity()
	var output = QueueWithCapacity[T](capacity)

	// Connect up the input queues to the output queue.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Take turns reading from each input queue and writing to the output queue.
		var i = 0
		var item, ok = inputs[i].RemoveHead() // Will block when empty.
		for ok {
			output.AddItem(item) // Will block when full.
			i = (i + 1) % size
			item, ok = inputs[i].RemoveHead() // Will block when empty.
		}

		// Close the output queue.
		output.CloseQueue()
	}()

	return output
}
