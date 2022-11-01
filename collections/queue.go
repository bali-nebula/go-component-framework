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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	syn "sync"
)

// QUEUE IMPLEMENTATION

// This constructor creates a new empty queue with the default capacity.
// The default capacity is 16 values.
func Queue[T abs.ValueLike]() abs.QueueLike[T] {
	return QueueWithCapacity[T](0)
}

// This constructor creates a new empty queue with the specified capacity.
func QueueWithCapacity[T abs.ValueLike](capacity int) abs.QueueLike[T] {
	// Groom the arguments.
	if capacity < 1 {
		capacity = 16 // The default value.
	}

	// Return an empty queue.
	var available = make(chan bool, capacity)
	var values = List[T]()
	return &queue[T]{available: available, values: values}
}

// This type defines the structure and methods associated with a queue of values.
// A queue implements first-in-first-out semantics. It is generally used by
// multiple goroutines at the same time and therefore enforces synchronized
// access. This type is parameterized as follows:
//   - T is any type of value.
//
// If the go chan type ever supports snapshots of its state, the underlying list
// can be removed and the channel modified to pass the values instead of the
// availability. Currently, the underlying list is only required by the
// AsArray() method.
type queue[T abs.ValueLike] struct {
	available chan bool
	values    abs.ListLike[T]
	mutex     syn.Mutex
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this queue is empty.
func (v *queue[T]) IsEmpty() bool {
	v.mutex.Lock()
	var result = len(v.available) == 0
	v.mutex.Unlock()
	return result
}

// This method returns the number of values contained in this queue.
func (v *queue[T]) GetSize() int {
	v.mutex.Lock()
	var result = len(v.available)
	v.mutex.Unlock()
	return result
}

// This method returns all the values in this queue. The values retrieved are in
// the same order as they are in the queue.
func (v *queue[T]) AsArray() []T {
	v.mutex.Lock()
	var result = v.values.AsArray()
	v.mutex.Unlock()
	return result
}

// FIFO INTERFACE

// This method retrieves the capacity of this queue.
func (v *queue[T]) GetCapacity() int {
	return cap(v.available) // The channel capacity is static.
}

// This method adds the specified value to the end of this queue.
func (v *queue[T]) AddValue(value T) {
	v.mutex.Lock()
	v.values.AddValue(value)
	v.mutex.Unlock()
	v.available <- true // Will block if at capacity.
}

// This method adds the specified values to the top of this queue.
func (v *queue[T]) AddValues(values abs.Sequential[T]) {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.mutex.Lock()
		v.values.AddValue(value)
		v.mutex.Unlock()
		v.available <- true // Will block if at capacity.
	}
}

// This method removes from this queue the value that is at the head of it. It
// returns the removed value and a "comma ok" value as the result.
func (v *queue[T]) RemoveHead() (head T, ok bool) {
	// Default the return value to the zero value for type T.

	// Remove the head value from the queue if one exists.
	_, ok = <-v.available // Will block until an value is available.
	if ok {
		v.mutex.Lock()
		head = v.values.RemoveValue(1)
		v.mutex.Unlock()
	}

	// Return the results
	return
}

// This method closes the queue so no more values can be placed on it.
func (v *queue[T]) CloseQueue() {
	v.mutex.Lock()
	close(v.available) // No more values can be placed on the queue.
	v.mutex.Unlock()
}

// QUEUES LIBRARY

// This constructor creates a new queues library for the specified generic
// value type.
func Queues[T abs.ValueLike]() *queues[T] {
	return &queues[T]{}
}

// This type defines the library functions that operate on queues. Since
// queues have a parameterized value type this library type is also
// parameterized as follows:
//   - T is any type of value.
type queues[T abs.ValueLike] struct{}

// ROUTEABLE INTERFACE

// This library function connects the output of the specified input queue with a
// number of new output queues specified by the size parameter and returns an
// array of the new output queues. Each value added to the input queue will be
// added automatically to ALL of the output queues. This pattern is useful when
// a set of DIFFERENT operations needs to occur for every value and each
// operation can be done in parallel.
func (l *queues[T]) Fork(wg *syn.WaitGroup, input abs.FIFO[T], size int) abs.Sequential[abs.FIFO[T]] {
	// Validate the arguments.
	if size < 1 {
		panic("The fan out size for a queue must be greater than zero.")
	}

	// Create the new output queues.
	var capacity = input.GetCapacity()
	var outputs = List[abs.FIFO[T]]()
	for i := 0; i < size; i++ {
		outputs.AddValue(abs.FIFO[T](QueueWithCapacity[T](capacity)))
	}

	// Connect up the input queue to the output queues in a separate goroutine.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Write each value read from the input queue to each output queue.
		var iterator = age.Iterator[abs.FIFO[T]](outputs)
		for {
			// Read from the input queue.
			var value, ok = input.RemoveHead() // Will block when empty.
			if !ok {
				break // The input queue has been closed.
			}

			// Write to all output queues.
			iterator.ToStart()
			for iterator.HasNext() {
				var output = iterator.GetNext()
				output.AddValue(value) // Will block when full.
			}
		}

		// Close all output queues.
		iterator.ToStart()
		for iterator.HasNext() {
			var output = iterator.GetNext()
			output.CloseQueue()
		}
	}()

	return outputs
}

// This library function connects the output of the specified input queue with
// the number of output queues specified by the size parameter and returns an
// array of the new output queues. Each value added to the input queue will be
// added automatically to ONE of the output queues. This pattern is useful when
// a SINGLE operation needs to occur for each value and the operation can be done
// on the values in parallel. The results can then be consolidated later on using
// the Join() function.
func (l *queues[T]) Split(wg *syn.WaitGroup, input abs.FIFO[T], size int) abs.Sequential[abs.FIFO[T]] {
	// Validate the arguments.
	if size < 1 {
		panic("The size of the split must be greater than zero.")
	}

	// Create the new output queues.
	var capacity = input.GetCapacity()
	var outputs = List[abs.FIFO[T]]()
	for i := 0; i < size; i++ {
		outputs.AddValue(abs.FIFO[T](QueueWithCapacity[T](capacity)))
	}

	// Connect up the input queue to the output queues.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Take turns reading from the input queue and writing to each output queue.
		var iterator = age.Iterator[abs.FIFO[T]](outputs)
		for {
			// Read from the input queue.
			var value, ok = input.RemoveHead() // Will block when empty.
			if !ok {
				break // The input queue has been closed.
			}

			// Write to the next output queue.
			var output = iterator.GetNext()
			output.AddValue(value) // Will block when full.
			if !iterator.HasNext() {
				iterator.ToStart()
			}
		}

		// Close all output queues.
		iterator.ToStart()
		for iterator.HasNext() {
			var output = iterator.GetNext()
			output.CloseQueue()
		}
	}()

	return outputs
}

// This library function connects the outputs of the specified array of input
// queues with a new output queue returns the new output queue. Each value
// removed from each input queue will automatically be added to the output
// queue. This pattern is useful when the results of the processing with a
// Split() function need to be consolicated into a single queue.
func (l *queues[T]) Join(wg *syn.WaitGroup, inputs abs.Sequential[abs.FIFO[T]]) abs.FIFO[T] {
	// Validate the arguments.
	if inputs == nil || inputs.IsEmpty() {
		panic("The number of input queues for a join must be greater than zero.")
	}

	// Create the new output queue.
	var iterator = age.Iterator(inputs)
	var capacity = iterator.GetNext().GetCapacity()
	var output = QueueWithCapacity[T](capacity)

	// Connect up the input queues to the output queue.
	wg.Add(1)
	go func() {
		// Make sure the wait group is decremented on termination.
		defer wg.Done()

		// Take turns reading from each input queue and writing to the output queue.
		iterator.ToStart()
		for {
			var input = iterator.GetNext()
			var value, ok = input.RemoveHead() // Will block when empty.
			if !ok {
				break // The input queue has been closed.
			}
			output.AddValue(value) // Will block when full.
			if !iterator.HasNext() {
				iterator.ToStart()
			}
		}

		// Close the output queue.
		output.CloseQueue()
	}()

	return output
}
