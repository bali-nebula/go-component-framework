/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	Primitive  any
	Collection any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all updatable sequences of
// values.
type Updatable[V Value] interface {
	SetValue(index int, value V)
	SetValues(index int, values Sequential[V])
}

// This interface defines the methods supported by all searchable sequences of
// values.
type Searchable[V Value] interface {
	GetIndex(value V) int
	ContainsValue(value V) bool
	ContainsAny(values Sequential[V]) bool
	ContainsAll(values Sequential[V]) bool
}

// This interface defines the methods supported by all sequences of values that
// allow values to be added and removed.
type Flexible[V Value] interface {
	AddValue(value V)
	AddValues(values Sequential[V])
	RemoveValue(value V)
	RemoveValues(values Sequential[V])
	RemoveAll()
}

// This interface defines the methods supported by all indexed sequences whose
// values may be added, inserted, or removed.
type Malleable[V Value] interface {
	AddValue(value V)
	AddValues(values Sequential[V])
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
	AddValue(value V)
	RemoveHead() (head V, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences whose values
// are accessed using last-in-first-out (LIFO) semantics.
type LIFO[V Value] interface {
	GetCapacity() int
	AddValue(value V)
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

// This interface consolidates all the interfaces supported by mapping-like
// sequences.
type MappingLike interface {
	Sequential[AssociationLike]
}

// This interface defines the methods supported by all association-iterator-like
// types.
type AssociationIteratorLike interface {
	Ratcheted[AssociationLike]
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
	Accessible[ComponentLike]
	Updatable[ComponentLike]
	Searchable[ComponentLike]
	Malleable[ComponentLike]
	Sortable[ComponentLike]
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike interface {
	Sequential[ComponentLike]
	FIFO[ComponentLike]
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike interface {
	Sequential[ComponentLike]
	Accessible[ComponentLike]
	Searchable[ComponentLike]
	Flexible[ComponentLike]
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike interface {
	Sequential[ComponentLike]
	LIFO[ComponentLike]
}
