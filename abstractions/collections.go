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

// This interface defines the methods supported by all binding types.
// It binds a readonly key with a setable value.
type Binding[K Key, V Value] interface {
	GetKey() K
	GetValue() V
	SetValue(value V)
}

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

// This interface defines the methods supported by all searchable sequences of
// values.
type Searchable[V Value] interface {
	ContainsValue(value V) bool
	ContainsAny(values Sequential[V]) bool
	ContainsAll(values Sequential[V]) bool
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by series-like
// sequences.
type SeriesLike interface {
	Sequential[ComponentLike]
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
