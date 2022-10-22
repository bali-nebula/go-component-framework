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

// This interface defines the methods supported by all component-like types.
type ComponentLike interface {
	IsGeneric() bool
	GetEntity() any
	GetContext() ContextLike
	GetNote() string
	SetNote(note string)
}

// This interface defines the methods supported by all context-like types.
type ContextLike interface {
	GetNames() Sequential[Symbolic]
	GetValue(name Symbolic) any
	SetValue(name Symbolic, value any)
	GetParameters() CatalogLike[Symbolic, any]
	SetParameters(parameters CatalogLike[Symbolic, any])
}
