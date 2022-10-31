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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// COMPONENT IMPLEMENTATION

// This constructor creates a new component.
func Component(entity any) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	return v
}

// This constructor creates a new component with the specified context.
func ComponentWithContext(entity any, context abs.ContextLike) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	v.SetContext(context)
	return v
}

// This type defines the structure and methods associated with a component.
type component struct {
	entity  any
	context abs.ContextLike
	note    string
}

// COMPONENT IMPLEMENTATION

// This method determines whether or not this component is parameterized.
func (v *component) IsGeneric() bool {
	return v.context != nil
}

func (v *component) IsAnnotated() bool {
	return len(v.note) > 0
}

// This method returns the entity for this component.
func (v *component) GetEntity() any {
	return v.entity
}

// This method sets the entity for this component.
func (v *component) SetEntity(entity any) {
	if entity == nil {
		panic("A component requires an entity.")
	}
	v.entity = entity
}

// This method returns the context for this component.
func (v *component) GetContext() abs.ContextLike {
	return v.context
}

// This method sets the context for this component.
func (v *component) SetContext(context abs.ContextLike) {
	v.context = context
}

// This method returns the note for this component.
func (v *component) GetNote() string {
	return v.note
}

// This method sets the note for this component.
func (v *component) SetNote(note string) {
	v.note = note
}
