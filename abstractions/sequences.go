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

type ValueLike any

type SequenceLike any

// This interface defines the methods supported by all sequences of values.
type Sequential[T ValueLike] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []T
}

// This interface defines the methods supported by all sequences whose values can
// be indexed. The indices of an indexed sequence are ORDINAL rather than ZERO
// based (which is "SO last century"). This allows for positive indices starting
// at the beginning of the sequence, and negative indices starting at the end of
// the sequence as follows:
//
//	    1           2           3             N
//	[value 1] . [value 2] . [value 3] ... [value N]
//	   -N        -(N-1)      -(N-2)          -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Indexed[T ValueLike] interface {
	SetComparer(compare ComparisonFunction)
	GetValue(index int) T
	GetValues(first int, last int) Sequential[T]
	GetIndex(value T) int
}

// This interface defines the methods supported by all searchable sequences of
// values.
type Searchable[T ValueLike] interface {
	ContainsValue(value T) bool
	ContainsAny(values Sequential[T]) bool
	ContainsAll(values Sequential[T]) bool
}

// This interface defines the methods supported by all updatable sequences of
// values.
type Updatable[T ValueLike] interface {
	SetValue(index int, value T)
	SetValues(index int, values Sequential[T])
}

// This interface defines the methods supported by all sequences of values that
// allow their endpoints to be changed.
type Elastic[T ValueLike] interface {
	IsEnumerable() bool
	GetFirst() T
	SetFirst(T)
	GetExtent() Extent
	SetExtent(extent Extent)
	GetLast() T
	SetLast(T)
}

// This interface defines the methods supported by all sequences of values that
// allow values to be added and removed.
type Flexible[T ValueLike] interface {
	SetRanker(rank RankingFunction)
	AddValue(value T)
	AddValues(values Sequential[T])
	RemoveValue(value T)
	RemoveValues(values Sequential[T])
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose values may
// be modified, inserted, removed, or reordered.
type Malleable[T ValueLike] interface {
	AddValue(value T)
	AddValues(values Sequential[T])
	InsertValue(slot int, value T)
	InsertValues(slot int, values Sequential[T])
	RemoveValue(index int) T
	RemoveValues(first int, last int) Sequential[T]
	RemoveAll()
	ShuffleValues()
	SortValues()
	SortValuesWithRanker(ranker RankingFunction)
	ReverseValues()
}

// This interface defines the methods supported by all associative sequences
// whose values consist of key-value pair associations.
type Associative[K KeyLike, V ValueLike] interface {
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

// This interface defines the methods supported by all sequences whose values
// are accessed using first-in-first-out (FIFO) semantics.
type FIFO[T ValueLike] interface {
	GetCapacity() int
	AddValue(value T)
	AddValues(values Sequential[T])
	RemoveHead() (head T, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences whose values
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[T ValueLike] interface {
	GetCapacity() int
	AddValue(value T)
	AddValues(values Sequential[T])
	GetTop() T
	RemoveTop() T
	RemoveAll()
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by array-like
// sequences.
type ArrayLike[T ValueLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
}

// This interface defines the methods supported by all association-like types.
// An association associates a key with an setable value.
type AssociationLike[K KeyLike, V ValueLike] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This interface consolidates all the interfaces supported by catalog-like
// sequences.
type CatalogLike[K KeyLike, V ValueLike] interface {
	Sequential[AssociationLike[K, V]]
	Associative[K, V]
}

// This interface consolidates all the interfaces supported by list-like
// sequences.
type ListLike[T ValueLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Updatable[T]
	Malleable[T]
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike[T ValueLike] interface {
	Sequential[T]
	FIFO[T]
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike[T ValueLike] interface {
	Sequential[T]
	Indexed[T]
	Searchable[T]
	Flexible[T]
}

// This interface consolidates all the interfaces supported by range-like
// sequences.
type RangeLike[T ValueLike] interface {
	Sequential[T]
	Indexed[T]
	Elastic[T]
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike[T ValueLike] interface {
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
