/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// LET CLAUSE IMPLEMENTATION

// This constructor creates a new let clause.
func LetClause(expression abs.Expression) abs.LetClauseLike {
	var v = &letClause{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This constructor creates a new let clause.
func LetClauseWithRecipient(recipient abs.Recipient, operator abs.Operator, expression abs.Expression) abs.LetClauseLike {
	var v = &letClause{}
	// Perform argument validation.
	v.SetRecipient(recipient, operator)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a let clause.
type letClause struct {
	recipient  abs.Recipient
	operator   abs.Operator
	expression abs.Expression
}

// This method determines whether or not this clause has a recipient.
func (v *letClause) HasRecipient() bool {
	return v.recipient != nil
}

// This method returns the recipient for this let clause.
func (v *letClause) GetRecipient() (abs.Recipient, abs.Operator) {
	return v.recipient, v.operator
}

// This method sets the recipient for this let clause.
func (v *letClause) SetRecipient(recipient abs.Recipient, operator abs.Operator) {
	if recipient != nil && operator >= abs.ASSIGN && operator <= abs.QUOTIENT {
		v.recipient = recipient
		v.operator = operator
	} else {
		v.recipient = nil
		v.operator = 0
	}
}

// This method returns the expression for this let clause.
func (v *letClause) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression for this let clause.
func (v *letClause) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("A let clause requires an expression.")
	}
	v.expression = expression
}
