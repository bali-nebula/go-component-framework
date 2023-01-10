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

// TYPE DEFINITIONS

type (
	Key        any
	Value      any
	Primitive  any
	Collection any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all sequences of values.
type Sequential[V Value] interface {
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
type Indexed[V Value] interface {
	GetValue(index int) V
	GetValues(first int, last int) Sequential[V]
	GetIndex(value V) int
}

// This interface defines the methods supported by all ratcheted agents that
// are capable of moving forward and backward over the values in a sequence. It
// is used to implement the GoF Iterator Pattern:
//   - https://en.wikipedia.org/wiki/Iterator_pattern
//
// A ratcheted agent locks into the slots that reside between each value in the
// sequence:
//
//	    [value 1] . [value 2] . [value 3] ... [value N]
//	  ^           ^           ^                         ^
//	slot 0      slot 1      slot 2                    slot N
//
// It moves from slot to slot and has access to the values (if they exist) on
// each side of the slot.
type Ratcheted[V Value] interface {
	GetSlot() int
	ToSlot(slot int)
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V
}

// This interface defines the methods supported by all searchable sequences of
// values.
type Searchable[V Value] interface {
	ContainsValue(value V) bool
	ContainsAny(values Sequential[V]) bool
	ContainsAll(values Sequential[V]) bool
}

// This interface defines the methods supported by all sequences of values that
// allow values to be added and removed.
type Expandable[V Value] interface {
	AddValue(value V)
	AddValues(values Sequential[V])
}

// This interface defines the methods supported by all sequences of values that
// allow values to be added and removed.
type Shrinkable[V Value] interface {
	RemoveValue(value V)
	RemoveValues(values Sequential[V])
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose values may
// be modified, inserted, removed, or reordered.
type Malleable[V Value] interface {
	SetValue(index int, value V)
	SetValues(index int, values Sequential[V])
	InsertValue(slot int, value V)
	InsertValues(slot int, values Sequential[V])
	RemoveValue(index int) V
	RemoveValues(first int, last int) Sequential[V]
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose values may
// be sorted using various sorting algorithms.
type Sortable[V Value] interface {
	SortValues()
	ReverseValues()
	ShuffleValues()
}

// This interface defines the methods supported by all binding types.
// It binds a readonly key with a setable value.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

// This interface defines the methods supported by all associative sequences
// whose values consist of key-value pair associations.
type Associative[K Key, V Value] interface {
	GetKeys() Sequential[K]
	GetValues(keys Sequential[K]) Sequential[V]
	GetValue(key K) V
	SetValue(key K, value V)
	RemoveValue(key K) V
	RemoveValues(keys Sequential[K]) Sequential[V]
	RemoveAll()
}

// This interface defines the methods supported by all sequences whose values
// are accessed using first-in-first-out (FIFO) semantics.
type FIFO[V Value] interface {
	GetCapacity() int
	RemoveHead() (head V, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences whose values
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[V Value] interface {
	GetCapacity() int
	GetTop() V
	RemoveTop() V
	RemoveAll()
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by series-like
// sequences.
type SeriesLike interface {
	Sequential[ComponentLike]
}

// This interface defines the methods supported by all component-iterator-like
// types.
type ComponentIteratorLike interface {
	Ratcheted[ComponentLike]
}

// This interface consolidates all the interfaces supported by association-like
// types.
type AssociationLike interface {
	Binding[Primitive, ComponentLike]
}

// This interface defines the methods supported by all association-iterator-like
// types.
type AssociationIteratorLike interface {
	Ratcheted[AssociationLike]
}

// This interface consolidates all the interfaces supported by mapping-like
// sequences.
type MappingLike interface {
	Sequential[AssociationLike]
}

// This interface consolidates all the interfaces supported by catalog-like
// sequences.
type CatalogLike interface {
	Sequential[AssociationLike]
	Associative[Primitive, ComponentLike]
	Sortable[AssociationLike]
}

// This interface consolidates all the interfaces supported by list-like
// sequences.
type ListLike interface {
	Sequential[ComponentLike]
	Indexed[ComponentLike]
	Searchable[ComponentLike]
	Expandable[ComponentLike]
	Malleable[ComponentLike]
	Sortable[ComponentLike]
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike interface {
	Sequential[ComponentLike]
	Expandable[ComponentLike]
	FIFO[ComponentLike]
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike interface {
	Sequential[ComponentLike]
	Indexed[ComponentLike]
	Searchable[ComponentLike]
	Expandable[ComponentLike]
	Shrinkable[ComponentLike]
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike interface {
	Sequential[ComponentLike]
	Expandable[ComponentLike]
	LIFO[ComponentLike]
}
