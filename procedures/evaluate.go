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
func EvaluateClause(expression abs.ExpressionLike) abs.EvaluateClauseLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This constructor creates a new evaluate clause.
func EvaluateClauseWithRecipient(recipient abs.RecipientLike, assignment abs.Assignment, expression abs.ExpressionLike) abs.EvaluateClauseLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetRecipient(recipient, assignment)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with an evaluate
// clause.
type evaluateClause struct {
	recipient  abs.RecipientLike
	assignment abs.Assignment
	expression abs.ExpressionLike
}

// This method returns the recipient (with assignment type) for this evaluate
// clause.
func (v *evaluateClause) GetRecipient() (recipient abs.RecipientLike, assignment abs.Assignment) {
	return v.recipient, v.assignment
}

// This method sets the recipient (with assignment type) for this evaluate
// clause.
func (v *evaluateClause) SetRecipient(recipient abs.RecipientLike, assignment abs.Assignment) {
	if recipient == nil || assignment < abs.REGULAR || assignment > abs.MINUS {
		panic("An evaluate clause requires a recipient and valid assignment type.")
	}
	v.recipient = recipient
	v.assignment = assignment
}

// This method returns the expression for this evaluate clause.
func (v *evaluateClause) GetExpression() abs.ExpressionLike {
	return v.expression
}

// This method sets the expression for this evaluate clause.
func (v *evaluateClause) SetExpression(expression abs.ExpressionLike) {
	if expression == nil {
		panic("An evaluate clause requires an expression.")
	}
	v.expression = expression
}
