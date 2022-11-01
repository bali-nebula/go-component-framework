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

// COMPONENT INTERFACES

type EntityLike any

// This interface defines the methods supported by all component-like types.
type ComponentLike interface {
	IsGeneric() bool
	IsAnnotated() bool
	GetEntity() EntityLike
	GetContext() ContextLike
	GetNote() string
	SetNote(note string)
}

// This interface defines the methods supported by all context-like types.
type ContextLike interface {
	GetNames() Sequential[Symbolic]
	GetValue(name Symbolic) ComponentLike
	SetValue(name Symbolic, value ComponentLike)
	GetParameters() CatalogLike[Symbolic, ComponentLike]
	SetParameters(parameters CatalogLike[Symbolic, ComponentLike])
}

// This interface defines the methods supported by all symbolic types.
type Symbolic interface {
	GetIdentifier() string
}
