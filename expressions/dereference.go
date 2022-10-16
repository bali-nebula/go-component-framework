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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// DEREFERENCE EXPRESSION IMPLEMENTATION

// This constructor creates a new dereference expression.
func Dereference(expression any) abstractions.DereferenceLike {
	var v = &dereferenceExpression{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a dereference
// expression.
type dereferenceExpression struct {
	expression any
}

// This method returns the expression to be dereference by this dereference
// expression.
func (v *dereferenceExpression) GetExpression() any {
	return v.expression
}

// This method sets the expression to be dereferenced by this dereference
// expression to the specified value.
func (v *dereferenceExpression) SetExpression(expression any) {
	if expression == nil {
		panic("The expression to be dereferenced cannot be nil.")
	}
	v.expression = expression
}
