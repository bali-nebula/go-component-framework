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

// TYPE ALIASES

// These type assignments hide the dependencies on the package used to implement
// the collection types. It preserves the collection interfaces in a way that
// will allow them to evolve separately as needed. Currently, the interfaces are
// synchronized.
type (
	Value      = col.Value
	Runes      = col.Sequential[rune]
	Lines      = col.Sequential[string]
	Parameter  = col.Binding[Symbolic, ComponentLike]
	Parameters = col.Sequential[Parameter]
)

// INDIVIDUAL INTERFACES

type Annotation any

type Entity any

type Sequential[V Value] col.Sequential[Value]

type IteratorLike[V Value] col.IteratorLike[V]

// This interface defines the methods supported by all component-like types.
type Encapsulated interface {
	GetEntity() Entity
	SetEntity(entity Entity)
	IsGeneric() bool
	GetContext() ContextLike
	SetContext(context ContextLike)
	IsAnnotated() bool
	GetNote() NoteLike
	SetNote(note NoteLike)
}

// This interface defines the methods supported by all symbolic types.
type Symbolic interface {
	GetIdentifier() string
}

// This interface defines the methods supported by all generic types.
type Generic interface {
	GetParameters() Parameters
	SetParameters(parameters Parameters)
}

// CONSOLIDATED INTERFACES

// This interface defines the interfaces supported by all component-like types.
type ComponentLike interface {
	Encapsulated
}

// This interface defines the interfaces supported by all context-like types.
type ContextLike interface {
	Generic
}

// This interface defines the interfaces supported by all note-like types.
type NoteLike interface {
	Runes
}

// This interface defines the interfaces supported by all comment-like types.
type CommentLike interface {
	Lines
}
