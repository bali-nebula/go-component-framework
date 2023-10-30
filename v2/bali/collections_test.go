/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	ass "github.com/stretchr/testify/assert"
	syn "sync"
	tes "testing"
)

func TestCatalogs(t *tes.T) {
	var foo = bal.Symbol("$foo")
	var bar = bal.Symbol("$bar")
	var baz = bal.Symbol("$baz")
	var one = bal.Component(1)
	var two = bal.Component(2)
	var three = bal.Component(3)
	var catalog = col.Catalog()
	ass.True(t, catalog.IsEmpty())
	ass.Equal(t, 0, catalog.GetSize())
	catalog.SetValue(foo, one)
	ass.False(t, catalog.IsEmpty())
	ass.Equal(t, 1, catalog.GetSize())
	ass.Equal(t, one, catalog.GetValue(foo))
	catalog.SetValue(bar, two)
	ass.Equal(t, 2, catalog.GetSize())
	ass.Equal(t, two, catalog.GetValue(bar))
	catalog.SetValue(baz, three)
	ass.Equal(t, 3, catalog.GetSize())
	ass.Equal(t, three, catalog.GetValue(baz))
	ass.Equal(t, []abs.Primitive{foo, bar, baz}, catalog.GetKeys().AsArray())
	catalog.SortValues()
	catalog.ShuffleValues()
	catalog.RemoveAll()
}

func TestLists(t *tes.T) {
	var one = bal.Component(1)
	var two = bal.Component(2)
	var three = bal.Component(3)
	var list = col.List()
	list.SortValues()
	ass.True(t, list.IsEmpty())
	ass.Equal(t, 0, list.GetSize())
	ass.False(t, list.ContainsValue(one))
	list.AddValue(one)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 1, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.False(t, list.ContainsValue(two))
	list.AddValue(two)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 2, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.True(t, list.ContainsValue(two))
	list.AddValue(three)
	ass.False(t, list.IsEmpty())
	ass.Equal(t, 3, list.GetSize())
	ass.True(t, list.ContainsValue(one))
	ass.True(t, list.ContainsValue(two))
	ass.True(t, list.ContainsValue(three))
	list.RemoveAll()
	ass.True(t, list.IsEmpty())
}

func TestQueueWithConcurrency(t *tes.T) {
	var one = bal.Component(1)
	var two = bal.Component(2)
	var three = bal.Component(3)

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
		var ok bool
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

func TestSets(t *tes.T) {
	var one = bal.Component(1)
	var two = bal.Component(2)
	var three = bal.Component(3)
	var set = col.Set()
	ass.True(t, set.IsEmpty())
	ass.Equal(t, 0, set.GetSize())
	ass.False(t, set.ContainsValue(one))
	set.AddValue(one)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 1, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.False(t, set.ContainsValue(two))
	set.AddValue(two)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 2, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.True(t, set.ContainsValue(two))
	set.AddValue(three)
	ass.False(t, set.IsEmpty())
	ass.Equal(t, 3, set.GetSize())
	ass.True(t, set.ContainsValue(one))
	ass.True(t, set.ContainsValue(two))
	ass.True(t, set.ContainsValue(three))
	set.RemoveAll()
	ass.True(t, set.IsEmpty())
}

func TestStacks(t *tes.T) {
	var one = bal.Component(1)
	var two = bal.Component(1)
	var three = bal.Component(1)
	var stack = col.Stack()
	ass.True(t, stack.IsEmpty())
	ass.Equal(t, 0, stack.GetSize())
	stack.AddValue(one)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 1, stack.GetSize())
	stack.AddValue(two)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 2, stack.GetSize())
	stack.AddValue(three)
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 3, stack.GetSize())
	ass.Equal(t, three, stack.RemoveTop())
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 2, stack.GetSize())
	ass.Equal(t, two, stack.RemoveTop())
	ass.False(t, stack.IsEmpty())
	ass.Equal(t, 1, stack.GetSize())
	ass.Equal(t, one, stack.RemoveTop())
	ass.True(t, stack.IsEmpty())
	ass.Equal(t, 0, stack.GetSize())
	stack.RemoveAll()
	ass.True(t, stack.IsEmpty())
}
