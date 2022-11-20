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
	col "github.com/craterdog/go-collection-framework"
)

// CONSTANTS

type Extent int

const (
	_ Extent = iota
	INCLUSIVE
	LEFT
	RIGHT
	EXCLUSIVE
)

// TYPE ALIASES

// These type assignments hide the dependencies on the package used to implement
// the collection types. It preserves the collection interfaces in a way that
// will allow them to evolve separately as needed. Currently, the interfaces are
// synchronized.
type (
	Primitive       = col.Key
	Collection      = col.Value
	Searchable      = col.Searchable[ComponentLike]
	Updatable       = col.Updatable[ComponentLike]
	Flexible        = col.Flexible[ComponentLike]
	Malleable       = col.Malleable[ComponentLike]
	Binding         = col.Binding[Primitive, ComponentLike]
	Associative     = col.Associative[Primitive, ComponentLike]
	FIFO            = col.FIFO[ComponentLike]
	LIFO            = col.LIFO[ComponentLike]
	AssociationLike = col.AssociationLike[Primitive, ComponentLike]
	CatalogLike     = col.CatalogLike[Primitive, ComponentLike]
	ListLike        = col.ListLike[ComponentLike]
	QueueLike       = col.QueueLike[ComponentLike]
	SetLike         = col.SetLike[ComponentLike]
	StackLike       = col.StackLike[ComponentLike]
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all sequences of primitive
// values that allow their endpoints to be changed. The type is parameterized
// to force the first and last primitive values to be the same type.
type Elastic[V Primitive] interface {
	IsEnumerable() bool
	GetFirst() V
	SetFirst(V)
	GetExtent() Extent
	SetExtent(extent Extent)
	GetLast() V
	SetLast(V)
}

// CONSOLIDATED INTERFACES

// This interface consolidates all of the interfaces supported by range-like
// sequences.
type RangeLike[V Primitive] interface {
	col.Sequential[V]
	col.Indexed[V]
	Elastic[V]
}
