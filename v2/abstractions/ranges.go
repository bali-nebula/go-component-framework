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
	Range  any
	Extent int
)

// CONSTANT DEFINITIONS

// This type and its associated constants define whether or not each endpoint in
// a bounded collection is included in the range of possible values.
const (
	_ Extent = iota
	INCLUSIVE
	LEFT
	RIGHT
	EXCLUSIVE
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all ranges of primitive
// values that allow their endpoints to be changed. The type is parameterized to
// force the first and last endpoint values to be the same type.
type Bounded[V Primitive] interface {
	GetFirst() V
	SetFirst(value V)
	GetExtent() Extent
	SetExtent(extent Extent)
	GetLast() V
	SetLast(value V)
}

// CONSOLIDATED INTERFACES

type ContinuumLike interface {
	Bounded[Continuous]
	Searchable[Continuous]
}

type IntervalLike interface {
	Bounded[Discrete]
	Sequential[Discrete]
	Accessible[Discrete]
	Searchable[Discrete]
}

type SpectrumLike interface {
	Bounded[Lexical]
	Searchable[Lexical]
}
