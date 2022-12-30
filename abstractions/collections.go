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
	tem "github.com/craterdog-bali/go-component-framework/temporary"
	col "github.com/craterdog/go-collection-framework"
)

// TYPE ALIASES

// These type assignments hide the dependencies on the package used to implement
// the collection types. It preserves the collection interfaces in a way that
// will allow them to evolve separately as needed. Currently, the interfaces are
// synchronized.
type (
	Primitive       = col.Key
	Collection      = col.Value
	Extent          = tem.Extent
	Sequential      = col.Sequential[ComponentLike]
	Searchable      = col.Searchable[ComponentLike]
	Updatable       = col.Updatable[ComponentLike]
	Flexible        = col.Flexible[ComponentLike]
	Malleable       = col.Malleable[ComponentLike]
	Binding         = col.Binding[Primitive, ComponentLike]
	Associative     = col.Associative[Primitive, ComponentLike]
	Structural      = col.Sequential[Binding]
	FIFO            = col.FIFO[ComponentLike]
	LIFO            = col.LIFO[ComponentLike]
	AssociationLike = col.AssociationLike[Primitive, ComponentLike]
	CatalogLike     = col.CatalogLike[Primitive, ComponentLike]
	ListLike        = col.ListLike[ComponentLike]
	QueueLike       = col.QueueLike[ComponentLike]
	SetLike         = col.SetLike[ComponentLike]
	StackLike       = col.StackLike[ComponentLike]
	IntervalLike    = tem.IntervalLike[Discrete]
	SpectrumLike    = tem.SpectrumLike[Spectral]
	ContinuumLike   = tem.ContinuumLike[Continuous]
)

// INDIVIDUAL TYPES

type Discrete int
type Spectral string
type Continuous float64
