/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// COMPONENT IMPLEMENTATION

// This constructor creates a new component.
func Component(entity abs.Entity) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	return v
}

// This constructor creates a new component with the specified context.
func ComponentWithContext(entity abs.Entity, context abs.ContextLike) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	v.SetContext(context)
	return v
}

// This type defines the structure and methods associated with a component.
type component struct {
	entity  abs.Entity
	context abs.ContextLike
	note    abs.NoteLike
}

// COMPONENT IMPLEMENTATION

// This method returns the entity for this component.
func (v *component) GetEntity() abs.Entity {
	return v.entity
}

// This method sets the entity for this component.
func (v *component) SetEntity(entity abs.Entity) {
	if entity == nil {
		panic("A component requires an entity.")
	}
	v.entity = entity
}

// This method determines whether or not this component is parameterized.
func (v *component) IsGeneric() bool {
	return v.context != nil
}

// This method returns the context for this component.
func (v *component) GetContext() abs.ContextLike {
	return v.context
}

// This method sets the context for this component.
func (v *component) SetContext(context abs.ContextLike) {
	v.context = context
}

// This method determines whether or not this component is annotated.
func (v *component) IsAnnotated() bool {
	return v.note != nil
}

// This method returns the note for this component.
func (v *component) GetNote() abs.NoteLike {
	return v.note
}

// This method sets the note for this component.
func (v *component) SetNote(note abs.NoteLike) {
	v.note = note
}
