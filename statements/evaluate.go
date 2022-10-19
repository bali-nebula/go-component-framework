/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package statements

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// EVALUATE CLAUSE IMPLEMENTATION

// This constructor creates a new evaluate clause.
func Evaluate(recipient any, assignment Assignment, expression any) abs.EvaluateLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetRecipient(recipient)
	v.SetAssignment(assignment)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with an evaluate
// clause.
type evaluateClause struct {
	recipient  any
	assignment any
	expression any
}

// This method returns the recipient for this evaluate clause.
func (v *evaluateClause) GetRecipient() any {
	return v.recipient
}

// This method sets the recipient for this evaluate clause.
func (v *evaluateClause) SetRecipient(recipient any) {
	if recipient == nil {
		panic("An evaluate clause requires a recipient.")
	}
	v.recipient = recipient
}

// This method returns the type of assignment for this evaluate clause.
func (v *evaluateClause) GetRecipient() any {
	return v.assignment
}

// This method sets the type of assignment for this evaluate clause.
func (v *evaluateClause) SetRecipient(assignment any) {
	v.assignment = assignment
}

// This method returns the expression for this evaluate clause.
func (v *evaluateClause) GetExpression() any {
	return v.expression
}

// This method sets the expression for this evaluate clause.
func (v *evaluateClause) SetExpression(expression any) {
	if expression == nil {
		panic("An evaluate clause requires an expression.")
	}
	v.expression = expression
}
