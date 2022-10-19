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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// COMPLEMENT EXPRESSION IMPLEMENTATION

// This constructor creates a new complement expression.
func Complement(expression any) abs.ComplementLike {
	var v = &complementExpression{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a complement
// expression.
type complementExpression struct {
	expression any
}

// This method returns the expression to be operated on by this complement
// expression.
func (v *complementExpression) GetExpression() any {
	return v.expression
}

// This method sets the expression to be operated on by this complement
// expression to the specified value.
func (v *complementExpression) SetExpression(expression any) {
	if expression == nil {
		panic("The expression to be operated on cannot be nil.")
	}
	v.expression = expression
}
