/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	"sync"
)

// COLLECTION INTERFACES

// This interface defines the methods supported by all searchable collections of
// items.
type Searchable[T any] interface {
	ContainsItem(item T) bool
	ContainsAny(items []T) bool
	ContainsAll(items []T) bool
}

// This interface defines the methods supported by all updatable collections of
// items.
type Updatable[T any] interface {
	SetItem(index int, item T)
	SetItems(index int, items []T)
}

// This interface defines the methods supported by all collections of items that
// allow their endpoints to be changed.
type Elastic[T any] interface {
	GetFirst() T
	SetFirst(T)
	GetConnector() string
	SetConnector(connector string)
	GetLast() T
	SetLast(T)
	IsEnumerable() bool
}

// This interface defines the methods supported by all collections of items that
// allow items to be added and removed.
type Flexible[T any] interface {
	AddItem(item T)
	AddItems(items []T)
	RemoveItem(item T)
	RemoveItems(items []T)
	RemoveAll()
}

// This interface defines the methods supported by all collections whose items
// may be modified, inserted, removed, or reordered.
type Malleable[T any] interface {
	AddItem(item T)
	AddItems(items []T)
	InsertItem(slot int, item T)
	InsertItems(slot int, items []T)
	RemoveItem(index int) T
	RemoveItems(first int, last int) []T
	RemoveAll()
	ShuffleItems()
	SortItems()
	SortItemsWithRanker(ranker RankingFunction)
	ReverseItems()
}

// This interface defines the methods supported by all associative collections
// whose items consist of key-value pair associations.
type Associative[K any, V any] interface {
	AddAssociation(association AssociationLike[K, V])
	AddAssociations(associations []AssociationLike[K, V])
	GetKeys() []K
	GetValue(key K) V
	GetValues(keys []K) []V
	SetValue(key K, value V)
	RemoveValue(key K) V
	RemoveValues(keys []K) []V
	RemoveAll()
	SortAssociations()
	SortAssociationsWithRanker(ranker RankingFunction)
	ReverseAssociations()
}

// This interface defines the methods supported by all collections whose items
// are accessed using first-in-first-out (FIFO) semantics.
type FIFO[T any] interface {
	GetCapacity() int
	AddItem(item T)
	AddItems(items []T)
	RemoveHead() (head T, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all collections whose items
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[T any] interface {
	GetCapacity() int
	AddItem(item T)
	AddItems(items []T)
	GetTop() T
	RemoveTop() T
	RemoveAll()
}

// LIBRARY INTERFACES

// This library interface defines the functions supported by all libraries that
// can apply an elastic range to another collection.
type Spliceable[T any] interface {
	Cut(collection StringLike[T], rng Elastic[T]) []T
	Splice(collection ListLike[T], rng Elastic[T], items []T) []T
}

// This library interface defines the functions supported by all libraries that
// can disect and combine associative collections.
type Blendable[K any, V any] interface {
	Merge(first, second Associative[K, V]) Associative[K, V]
	Extract(catalog Associative[K, V], keys []K) Associative[K, V]
}

// This library interface defines the functions supported by all libraries that
// can combine queue-like collections in interesting ways.
type Routeable[T any] interface {
	Fork(wg *sync.WaitGroup, input FIFO[T], size int) []FIFO[T]
	Split(wg *sync.WaitGroup, input FIFO[T], size int) []FIFO[T]
	Join(wg *sync.WaitGroup, inputs []FIFO[T]) FIFO[T]
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by array-like
// collections.
type ArrayLike[T any] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
}

// This interface defines the methods supported by all association-like types.
// An association associates a key with an setable value.
type AssociationLike[K any, V any] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This interface consolidates all the interfaces supported by catalog-like
// collections.
type CatalogLike[K any, V any] interface {
	Sequential[AssociationLike[K, V]]
	Indexed[AssociationLike[K, V]]
	Associative[K, V]
}

// This interface consolidates all the interfaces supported by list-like
// collections.
type ListLike[T any] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
	Malleable[T]
}

// This interface consolidates all of the interfaces supported by queue-like
// collections.
type QueueLike[T any] interface {
	Sequential[T]
	FIFO[T]
}

// This interface consolidates all the interfaces supported by set-like
// collections.
type SetLike[T any] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Flexible[T]
}

// This interface consolidates all the interfaces supported by range-like
// collections.
type RangeLike[T any] interface {
	Sequential[T]
	Indexed[T]
	Elastic[T]
}

// This interface consolidates all the interfaces supported by stack-like
// collections.
type StackLike[T any] interface {
	Sequential[T]
	LIFO[T]
}
