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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
)

// CONTEXT IMPLEMENTATION

// This constructor creates a new component context.
func Context(parameters abs.Parameters) abs.ContextLike {
	var v = &context{}
	// Perform argument validation.
	v.SetParameters(parameters)
	return v
}

// This type defines the structure and methods associated with a component
// context.
type context struct {
	parameters abs.Parameters
}

// This method returns the catalog of parameters for this component context.
func (v *context) GetParameters() abs.Parameters {
	return v.parameters
}

// This method sets the catalog of parameters for this component context.
func (v *context) SetParameters(parameters abs.Parameters) {
	if parameters == nil || parameters.IsEmpty() {
		panic("A component context requires at least one parameter.")
	}
	v.parameters = parameters
}
