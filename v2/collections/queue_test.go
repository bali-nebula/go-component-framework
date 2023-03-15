/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	syn "sync"
	tes "testing"
)

func TestQueueWithConcurrency(t *tes.T) {
	var one = com.Component(ele.NumberFromComplex(1))
	var two = com.Component(ele.NumberFromComplex(2))
	var three = com.Component(ele.NumberFromComplex(3))

	// Create a wait group for synchronization.
	var wg = new(syn.WaitGroup)
	defer wg.Wait()

	// Create a new queue with a specific capacity.
	var queue = col.QueueWithCapacity(12)
	ass.Equal(t, 12, queue.GetCapacity())
	ass.True(t, queue.IsEmpty())
	ass.Equal(t, 0, queue.GetSize())

	// Add some values to the queue.
	queue.AddValue(one)
	queue.AddValue(two)
	ass.Equal(t, 2, queue.GetSize())

	// Remove values from the queue in the background.
	wg.Add(1)
	go func() {
		defer wg.Done()
		var value abs.ComponentLike
		var ok = true
		value, ok = queue.RemoveHead()
		ass.True(t, ok)
		ass.Equal(t, one, value)
		value, ok = queue.RemoveHead()
		ass.True(t, ok)
		ass.Equal(t, two, value)
		value, ok = queue.RemoveHead()
		ass.True(t, ok)
		ass.Equal(t, three, value)
	}()

	// Add some more values to the queue.
	queue.AddValue(three)
	queue.CloseQueue()
}
