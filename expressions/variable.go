/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// VARIABLE EXPRESSION IMPLEMENTATION

// This constructor creates a new variable expression.
func Variable(identifier string) abs.VariableLike {
	var v = &variableExpression{}
	// Perform argument validation.
	v.SetIdentifier(identifier)
	return v
}

// This type defines the structure and methods associated with a variable
// expression.
type variableExpression struct {
	identifier string
}

// This method is a dummy method that always returns true.
func (v *variableExpression) IsVariable() bool {
	return true
}

// This method returns the identifier for this variable expression.
func (v *variableExpression) GetIdentifier() string {
	return v.identifier
}

// This method sets the identifier for this variable expression.
func (v *variableExpression) SetIdentifier(identifier string) {
	if len(identifier) == 0 {
		panic("A variable expression requires an identifier.")
	}
	v.identifier = identifier
}
