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
	"fmt"
	"sync"
)

// STRING INTERFACES

// This interface defines the methods supported by all sequential strings of
// items.
type Sequential[T any] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []T
}

// This interface defines the methods supported by all strings of items whose
// items can be indexed. The indices of an indexed string are ORDINAL rather
// than ZERO based (which is "SO last century"). This allows for positive
// indices starting at the beginning of the string, and negative indices
// starting at the end of the string as follows:
//
//	    1          2          3            N
//	[item 1] . [item 2] . [item 3] ... [item N]
//	   -N       -(N-1)     -(N-2)         -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Indexed[T any] interface {
	GetItem(index int) T
	GetItems(first int, last int) []T
	GetIndex(item T) int
}

// This function normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func NormalizedIndex(index int, length int) int {
	switch {
	case index < -length || index == 0 || index > length:
		// The index is outside the bounds of the specified range.
		panic(fmt.Sprintf(
			"The specified index is outside the allowed ranges [-%v..-1] and [1..%v]: %v",
			length,
			length,
			index))
	case index < 0:
		// Convert a negative index.
		return index + length
	case index > 0:
		// Convert a positive index.
		return index - 1
	default:
		// This should never happen so time to panic...
		panic(fmt.Sprintf("Compiler problem, unexpected index value: %v", index))
	}
}

// COLLECTION INTERFACES

// This interface defines the methods supported by all searchable collections of
// items.
type Searchable[T any] interface {
	ContainsItem(item T) bool
	ContainsAny(items []T) bool
	ContainsAll(items []T) bool
}

// This interface defines the methods supported by all collections of items that
// allow their endpoints to be changed.
type Elastic[T any] interface {
	GetFirst() T
	SetFirst(T) T
	GetConnector() string
	SetConnector(connector string) string
	GetLast() T
	SetLast(T) T
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
	SetItem(index int, item T) T
	SetItems(index int, items []T) []T
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
	SetValue(key K, value V) V
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

// This library interface defines the functions supported by all libraries of
// chainable items.
type Chainable[T any] interface {
	Concatenate(first Sequential[T], second Sequential[T]) Sequential[T]
}

// This library interface defines the functions supported by all libraries that
// can apply an elastic slice to another collection.
type Spliceable[T any] interface {
	Cut(collection StringLike[T], slice Elastic[T]) []T
	Splice(collection ListLike[T], slice Elastic[T], items []T) []T
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

// This interface defines the methods supported by all association-like types.
// An association associates a key with an setable value.
type AssociationLike[K any, V any] interface {
	GetKey() K
	GetValue() V
	SetValue(value V) V
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

// This interface consolidates all the interfaces supported by slice-like
// collections.
type SliceLike[T any] interface {
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

// This interface consolidates all the interfaces supported by string-like
// collections.
type StringLike[T any] interface {
	Sequential[T]
	Indexed[T]
}
