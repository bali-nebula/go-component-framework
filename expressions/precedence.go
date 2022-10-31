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

// PRECEDENCE EXPRESSION IMPLEMENTATION

// This constructor creates a new precedence expression.
func Precedence(expression abs.ExpressionLike) abs.PrecedenceLike {
	var v = &precedenceExpression{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a precedence
// expression.
type precedenceExpression struct {
	expression abs.ExpressionLike
}

// This method returns the expression wrapped by this precedence expression.
func (v *precedenceExpression) GetExpression() abs.ExpressionLike {
	return v.expression
}

// This method sets the expression to be wrapped by this precedence expression
// to the specified value.
func (v *precedenceExpression) SetExpression(expression abs.ExpressionLike) {
	if expression == nil {
		panic("The expression wrapped by a precedence expression cannot be nil.")
	}
	v.expression = expression
}
