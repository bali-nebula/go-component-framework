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

type PrimitiveLike any

type ValueLike any

type StringLike any

type CollectionLike any

// This interface defines the methods supported by all sequences of values.
type Sequential[V ValueLike] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []V
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
type Indexed[V ValueLike] interface {
	SetComparer(compare ComparisonFunction)
	GetValue(index int) V
	GetValues(first int, last int) Sequential[V]
	GetIndex(value V) int
}

// This interface defines the methods supported by all searchable sequences of
// values.
type Searchable[V ValueLike] interface {
	ContainsValue(value V) bool
	ContainsAny(values Sequential[V]) bool
	ContainsAll(values Sequential[V]) bool
}

// This interface defines the methods supported by all updatable sequences of
// values.
type Updatable[V ValueLike] interface {
	SetValue(index int, value V)
	SetValues(index int, values Sequential[V])
}

// This interface defines the methods supported by all sequences of values that
// allow their endpoints to be changed.
type Elastic[V ValueLike] interface {
	IsEnumerable() bool
	GetFirst() V
	SetFirst(V)
	GetExtent() Extent
	SetExtent(extent Extent)
	GetLast() V
	SetLast(V)
}

// This interface defines the methods supported by all sequences of values that
// allow values to be added and removed.
type Flexible[V ValueLike] interface {
	SetRanker(rank RankingFunction)
	AddValue(value V)
	AddValues(values Sequential[V])
	RemoveValue(value V)
	RemoveValues(values Sequential[V])
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose values may
// be modified, inserted, removed, or reordered.
type Malleable[V ValueLike] interface {
	AddValue(value V)
	AddValues(values Sequential[V])
	InsertValue(slot int, value V)
	InsertValues(slot int, values Sequential[V])
	RemoveValue(index int) V
	RemoveValues(first int, last int) Sequential[V]
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
type FIFO[V ValueLike] interface {
	GetCapacity() int
	AddValue(value V)
	AddValues(values Sequential[V])
	RemoveHead() (head V, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences whose values
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[V ValueLike] interface {
	GetCapacity() int
	AddValue(value V)
	AddValues(values Sequential[V])
	GetTop() V
	RemoveTop() V
	RemoveAll()
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by array-like
// sequences.
type ArrayLike[V ValueLike] interface {
	Sequential[V]
	Indexed[V]
	Searchable[V]
	Updatable[V]
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
type ListLike[V ValueLike] interface {
	Sequential[V]
	Indexed[V]
	Searchable[V]
	Updatable[V]
	Malleable[V]
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike[V ValueLike] interface {
	Sequential[V]
	FIFO[V]
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike[V ValueLike] interface {
	Sequential[V]
	Indexed[V]
	Searchable[V]
	Flexible[V]
}

// This interface consolidates all the interfaces supported by range-like
// sequences.
type RangeLike[V ValueLike] interface {
	Sequential[V]
	Indexed[V]
	Elastic[V]
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike[V ValueLike] interface {
	Sequential[V]
	LIFO[V]
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
