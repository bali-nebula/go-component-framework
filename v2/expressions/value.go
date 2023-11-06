/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// VALUE EXPRESSION IMPLEMENTATION

// This constructor creates a new value expression.
func Value(component abs.ComponentLike) abs.ValueLike {
	var v = &valueExpression{}
	// Perform argument validation.
	v.SetComponent(component)
	return v
}

// This type defines the structure and methods associated with a value
// expression.
type valueExpression struct {
	component abs.ComponentLike
}

// This method returns the component for this value expression.
func (v *valueExpression) GetComponent() abs.ComponentLike {
	return v.component
}

// This method sets the component for this value expression.
func (v *valueExpression) SetComponent(component abs.ComponentLike) {
	if component == nil {
		panic("A value expression requires a component.")
	}
	v.component = component
}
