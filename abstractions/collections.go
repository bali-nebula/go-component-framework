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

// TYPE DEFINITIONS

type (
	Primitive  any
	Collection any
)

// TYPE ALIASES

// These type assignments hide the dependencies on the package used to implement
// the collection types. It preserves the collection interfaces in a way that
// will allow them to evolve separately as needed. Currently, the interfaces are
// synchronized.
type (
	Searchable         = col.Searchable[ComponentLike]
	Updatable          = col.Updatable[ComponentLike]
	Flexible           = col.Flexible[ComponentLike]
	Malleable          = col.Malleable[ComponentLike]
	Binding            = col.Binding[Primitive, ComponentLike]
	Associative        = col.Associative[Primitive, ComponentLike]
	FIFO               = col.FIFO[ComponentLike]
	LIFO               = col.LIFO[ComponentLike]
	ComparisonFunction = col.ComparisonFunction
	RankingFunction    = col.RankingFunction
	SeriesLike         = col.ArrayLike[ComponentLike]
	ListLike           = col.ListLike[ComponentLike]
	QueueLike          = col.QueueLike[ComponentLike]
	SetLike            = col.SetLike[ComponentLike]
	StackLike          = col.StackLike[ComponentLike]
	AssociationLike    = col.AssociationLike[Primitive, ComponentLike]
	StructureLike      = col.MapLike[Primitive, ComponentLike]
	CatalogLike        = col.CatalogLike[Primitive, ComponentLike]
)
