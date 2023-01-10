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
	col "github.com/craterdog/go-collection-framework"
)

// PARAMETER IMPLEMENTATION

// This constructor creates a new parameter name-value pair.
func Parameter(name abs.Symbolic, value abs.ComponentLike) abs.ParameterLike {
	return col.Association(name, value)
}

// CONTEXT IMPLEMENTATION

// This constructor creates a new component context.
func Context() abs.ContextLike {
	var parameters = col.Catalog[abs.Symbolic, abs.ComponentLike]()
	var v = &context{parameters}
	return v
}

// This type defines the structure and methods associated with a component
// context.
type context struct {
	parameters col.CatalogLike[abs.Symbolic, abs.ComponentLike]
}

// This method returns the sequence of parameter names for this context.
func (v *context) GetNames() abs.Sequential[abs.Symbolic] {
	return v.parameters.GetKeys()
}

// This method returns the parameter value associated with the specified name.
func (v *context) GetValue(name abs.Symbolic) abs.ComponentLike {
	return v.parameters.GetValue(name)
}

// This method set the value of the parameter associated with the specified
// name.
func (v *context) SetValue(name abs.Symbolic, value abs.ComponentLike) {
	v.parameters.SetValue(name, value)
}
