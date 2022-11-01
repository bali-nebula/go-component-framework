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
	fmt "fmt"
)

// SEQUENCE INTERFACES

type KeyLike any

type ItemLike any

type SequenceLike any

// This interface defines the methods supported by all sequences of items.
type Sequential[T ItemLike] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []T
}

// This interface defines the methods supported by all sequences whose items can
// be indexed. The indices of an indexed sequence are ORDINAL rather than ZERO
// based (which is "SO last century"). This allows for positive indices starting
// at the beginning of the sequence, and negative indices starting at the end of
// the sequence as follows:
//
//	    1          2          3            N
//	[item 1] . [item 2] . [item 3] ... [item N]
//	   -N       -(N-1)     -(N-2)         -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Indexed[T ItemLike] interface {
	SetComparer(compare ComparisonFunction)
	GetItem(index int) T
	GetItems(first int, last int) Sequential[T]
	GetIndex(item T) int
}

// This interface defines the methods supported by all searchable sequences of
// items.
type Searchable[T ItemLike] interface {
	ContainsItem(item T) bool
	ContainsAny(items Sequential[T]) bool
	ContainsAll(items Sequential[T]) bool
}

// This interface defines the methods supported by all updatable sequences of
// items.
type Updatable[T ItemLike] interface {
	SetItem(index int, item T)
	SetItems(index int, items Sequential[T])
}

// This interface defines the methods supported by all sequences of items that
// allow their endpoints to be changed.
type Elastic[T ItemLike] interface {
	IsEnumerable() bool
	GetFirst() T
	SetFirst(T)
	GetExtent() Extent
	SetExtent(extent Extent)
	GetLast() T
	SetLast(T)
}

// This interface defines the methods supported by all sequences of items that
// allow items to be added and removed.
type Flexible[T ItemLike] interface {
	SetRanker(rank RankingFunction)
	AddItem(item T)
	AddItems(items Sequential[T])
	RemoveItem(item T)
	RemoveItems(items Sequential[T])
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose items may
// be modified, inserted, removed, or reordered.
type Malleable[T ItemLike] interface {
	AddItem(item T)
	AddItems(items Sequential[T])
	InsertItem(slot int, item T)
	InsertItems(slot int, items Sequential[T])
	RemoveItem(index int) T
	RemoveItems(first int, last int) Sequential[T]
	RemoveAll()
	ShuffleItems()
	SortItems()
	SortItemsWithRanker(ranker RankingFunction)
	ReverseItems()
}

// This interface defines the methods supported by all associative sequences
// whose items consist of key-value pair associations.
type Associative[K KeyLike, V ItemLike] interface {
	AddAssociation(association AssociationLike[K, V])
	AddAssociations(associations Sequential[AssociationLike[K, V]])
	GetKeys() Sequential[K]
	GetValue(key K) V
	GetValues(keys Sequential[K]) Sequential[V]
	SetValue(key K, value V)
	RemoveValue(key K) V
	RemoveValues(keys Sequential[K]) Sequential[V]
	RemoveAll()
	SortAssociations()
	SortAssociationsWithRanker(ranker RankingFunction)
	ReverseAssociations()
}

// This interface defines the methods supported by all sequences whose items
// are accessed using first-in-first-out (FIFO) semantics.
type FIFO[T ItemLike] interface {
	GetCapacity() int
	AddItem(item T)
	AddItems(items Sequential[T])
	RemoveHead() (head T, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences whose items
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[T ItemLike] interface {
	GetCapacity() int
	AddItem(item T)
	AddItems(items Sequential[T])
	GetTop() T
	RemoveTop() T
	RemoveAll()
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by array-like
// sequences.
type ArrayLike[T ItemLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
}

// This interface defines the methods supported by all association-like types.
// An association associates a key with an setable value.
type AssociationLike[K KeyLike, V ItemLike] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This interface consolidates all the interfaces supported by catalog-like
// sequences.
type CatalogLike[K KeyLike, V ItemLike] interface {
	Sequential[AssociationLike[K, V]]
	Indexed[AssociationLike[K, V]]
	Associative[K, V]
}

// This interface consolidates all the interfaces supported by list-like
// sequences.
type ListLike[T ItemLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
	Malleable[T]
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike[T ItemLike] interface {
	Sequential[T]
	FIFO[T]
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike[T ItemLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Flexible[T]
}

// This interface consolidates all the interfaces supported by range-like
// sequences.
type RangeLike[T ItemLike] interface {
	Sequential[T]
	Indexed[T]
	Elastic[T]
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike[T ItemLike] interface {
	Sequential[T]
	LIFO[T]
}

// CONSTANTS

type Extent int

const (
	_ Extent = iota
	INCLUSIVE
	LEFT
	RIGHT
	EXCLUSIVE
)

// FUNCTIONS

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
