/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// EVALUATE CLAUSE IMPLEMENTATION

// This constructor creates a new evaluate clause.
func EvaluateClause(expression abs.Expression) abs.EvaluateClauseLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This constructor creates a new evaluate clause.
func EvaluateClauseWithRecipient(recipient any, assignment abs.Assignment, expression abs.Expression) abs.EvaluateClauseLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetRecipient(recipient, assignment)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with an evaluate
// clause.
type evaluateClause struct {
	recipient  any
	assignment abs.Assignment
	expression abs.Expression
}

// This method returns the recipient (with assignment type) for this evaluate
// clause.
func (v *evaluateClause) GetRecipient() (recipient any, assignment abs.Assignment) {
	return v.recipient, v.assignment
}

// This method sets the recipient (with assignment type) for this evaluate
// clause.
func (v *evaluateClause) SetRecipient(recipient any, assignment abs.Assignment) {
	if recipient == nil || assignment < abs.REGULAR || assignment > abs.MINUS {
		panic("An evaluate clause requires a recipient and valid assignment type.")
	}
	v.recipient = recipient
	v.assignment = assignment
}

// This method returns the expression for this evaluate clause.
func (v *evaluateClause) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression for this evaluate clause.
func (v *evaluateClause) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("An evaluate clause requires an expression.")
	}
	v.expression = expression
}
