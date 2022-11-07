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
func EvaluateClauseWithRecipient(recipient abs.RecipientLike, operator abs.Operator, expression abs.ExpressionLike) abs.EvaluateClauseLike {
	var v = &evaluateClause{}
	// Perform argument validation.
	v.SetRecipient(recipient, operator)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with an evaluate
// clause.
type evaluateClause struct {
	recipient  abs.RecipientLike
	operator   abs.Operator
	expression abs.ExpressionLike
}

// This method is a dummy method that always returns true.
func (v *evaluateClause) IsEvaluateClause() bool {
	return true
}

// This method is a dummy method that always returns true.
func (v *evaluateClause) HasRecipient() bool {
	return v.recipient != nil
}

// This method returns the recipient for this evaluate clause.
func (v *evaluateClause) GetRecipient() (abs.RecipientLike, abs.Operator) {
	return v.recipient, v.operator
}

// This method sets the recipient for this evaluate clause.
func (v *evaluateClause) SetRecipient(recipient abs.RecipientLike, operator abs.Operator) {
	if recipient != nil && operator >= abs.ASSIGN && operator <= abs.QUOTIENT {
		v.recipient = recipient
		v.operator = operator
	} else {
		v.recipient = nil
		v.operator = 0
	}
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
