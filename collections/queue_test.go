/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestQueueWithConcurrency(t *testing.T) {
	// Create a wait group for synchronization.
	var wg = new(sync.WaitGroup)
	defer wg.Wait()

	// Create a new queue with a specific capacity.
	var queue = collections.QueueWithCapacity[int](12)
	assert.Equal(t, 12, queue.GetCapacity())
	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.GetSize())

	// Add items to the queue in bulk.
	for i := 1; i < 10; i++ {
		queue.AddItem(i)
	}
	assert.Equal(t, 9, queue.GetSize())

	// Remove items from the queue in the background.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			var item, _ = queue.RemoveHead()
			assert.Equal(t, i, item)
		}
	}()

	// Add more items to the queue.
	for i := 10; i < 101; i++ {
		queue.AddItem(i)
	}
}

func TestQueueWithFork(t *testing.T) {
	var queues = collections.Queues[int]()

	// Create a wait group for synchronization.
	var wg = new(sync.WaitGroup)
	defer wg.Wait()

	// Create a new queue with a fan out of two.
	var input = collections.QueueWithCapacity[int](3)
	var outputs = queues.Fork(wg, input, 2)

	// Remove items from the output queues in the background.
	var readOutput = func(output abstractions.FIFO[int], name string) {
		defer wg.Done()
		var item, ok = output.RemoveHead()
		for i := 1; ok; i++ {
			assert.Equal(t, i, item)
			item, ok = output.RemoveHead()
		}
	}
	wg.Add(2)
	var iterator = agents.Iterator(outputs)
	for iterator.HasNext() {
		var output = iterator.GetNext()
		go readOutput(output, "output")
	}

	// Add items to the input queue.
	for i := 1; i < 11; i++ {
		input.AddItem(i)
	}
	input.CloseQueue()
}

func TestQueueWithSplitAndJoin(t *testing.T) {
	var queues = collections.Queues[int]()

	// Create a wait group for synchronization.
	var wg = new(sync.WaitGroup)
	defer wg.Wait()

	// Create a new queue with a split of five outputs and a join back to one.
	var input = collections.QueueWithCapacity[int](3)
	var split = queues.Split(wg, input, 5)
	var output = queues.Join(wg, split)

	// Remove items from the output queue in the background.
	wg.Add(1)
	go func() {
		defer wg.Done()
		var item, ok = output.RemoveHead()
		for i := 1; ok; i++ {
			assert.Equal(t, i, item)
			item, ok = output.RemoveHead()
		}
	}()

	// Add items to the input queue.
	for i := 1; i < 21; i++ {
		input.AddItem(i)
	}
	input.CloseQueue()
}
