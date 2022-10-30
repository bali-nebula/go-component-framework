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

// CONTEXT IMPLEMENTATION

// This constructor creates a new component context.
func Context(parameters abs.CatalogLike[abs.Symbolic, abs.ComponentLike]) abs.ContextLike {
	var v = &context{}
	// Perform argument validation.
	v.SetParameters(parameters)
	return v
}

// This type defines the structure and methods associated with a component
// context.
type context struct {
	parameters abs.CatalogLike[abs.Symbolic, abs.ComponentLike]
}

// This method returns the list of parameter names from this component context.
func (v *context) GetNames() abs.Sequential[abs.Symbolic] {
	return v.parameters.GetKeys()
}

// This method returns the parameter value with the specified name from this
// component context.
func (v *context) GetValue(name abs.Symbolic) abs.ComponentLike {
	return v.parameters.GetValue(name)
}

// This method sets the parameter with the specified name to the specified value
// for this component context.
func (v *context) SetValue(name abs.Symbolic, value abs.ComponentLike) {
	if value == nil {
		panic("Each parameter for a component context requires a value.")
	}
	v.parameters.SetValue(name, value)
}

// This method returns the catalog of parameters for this component context.
func (v *context) GetParameters() abs.CatalogLike[abs.Symbolic, abs.ComponentLike] {
	return v.parameters
}

// This method sets the catalog of parameters for this component context.
func (v *context) SetParameters(parameters abs.CatalogLike[abs.Symbolic, abs.ComponentLike]) {
	if parameters == nil || parameters.IsEmpty() {
		panic("A component context requires at least one parameter.")
	}
	v.parameters = parameters
}
