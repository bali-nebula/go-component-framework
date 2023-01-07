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
	Primitive  any
	Collection any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all updatable sequences of
// components.
type Updatable interface {
	SetValue(index int, value ComponentLike)
}

// This interface defines the methods supported by all sequences of components
// that allow components to be added and removed.
type Flexible interface {
	AddValue(value ComponentLike)
	RemoveValue(value ComponentLike)
	RemoveAll()
}

// This interface defines the methods supported by all sequences of components
// whose components may be modified, inserted, removed, or reordered.
type Malleable interface {
	AddValue(value ComponentLike)
	InsertValue(slot int, value ComponentLike)
	RemoveValue(index int) ComponentLike
	RemoveAll()
}

// This interface defines the methods supported by all sequences of components
// whose components may be sorted using various sorting algorithms.
type Sortable interface {
	SortValues()
	ReverseValues()
	ShuffleValues()
}

// This interface defines the methods supported by all associative sequences
// whose values consist of key-value pair associations.
type Associative interface {
	IsEmpty() bool
	GetSize() int
	GetKeys() Sequential[Primitive]
	GetValue(key Primitive) ComponentLike
	SetValue(key Primitive, value ComponentLike)
	RemoveValue(key Primitive) ComponentLike
	RemoveAll()
}

// This interface defines the methods supported by all sequences of components
// whose components are accessed using first-in-first-out (FIFO) semantics.
type FIFO interface {
	GetCapacity() int
	AddValue(value ComponentLike)
	RemoveHead() (head ComponentLike, ok bool)
	CloseQueue()
}

// This interface defines the methods supported by all sequences of components
// whose components are accessed using last-in-first-out (LIFO) semantics.
type LIFO interface {
	GetCapacity() int
	AddValue(value ComponentLike)
	GetTop() ComponentLike
	RemoveTop() ComponentLike
	RemoveAll()
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by catalog-like
// sequences.
type CatalogLike interface {
	Associative
	Sortable
}

// This interface consolidates all the interfaces supported by list-like
// sequences.
type ListLike interface {
	Sequential[ComponentLike]
	Indexed[ComponentLike]
	Updatable
	Malleable
	Sortable
}

// This interface consolidates all of the interfaces supported by queue-like
// sequences.
type QueueLike interface {
	Sequential[ComponentLike]
	FIFO
}

// This interface consolidates all the interfaces supported by series-like
// sequences.
type SeriesLike interface {
	Sequential[ComponentLike]
	Indexed[ComponentLike]
	Updatable
}

// This interface consolidates all the interfaces supported by set-like
// sequences.
type SetLike interface {
	Sequential[ComponentLike]
	Indexed[ComponentLike]
	Flexible
}

// This interface consolidates all the interfaces supported by stack-like
// sequences.
type StackLike interface {
	Sequential[ComponentLike]
	LIFO
}

// This interface consolidates all the interfaces supported by association-like
// types.
type AssociationLike interface {
	Binding[Primitive, ComponentLike]
}

// This interface consolidates all the interfaces supported by structure-like
// sequences.
type StructureLike interface {
	Associative
}
