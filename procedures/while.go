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

// WHILE CLAUSE IMPLEMENTATION

// This constructor creates a new while clause.
func WhileClause(condition abs.ExpressionLike, statements abs.ProcedureLike) abs.WhileClauseLike {
	var v = &whileClause{}
	// Perform argument validation.
	v.SetCondition(condition)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a while clause.
type whileClause struct {
	condition  abs.ExpressionLike
	statements abs.ProcedureLike
}

// This method returns the condition expression for this while clause.
func (v *whileClause) GetCondition() abs.ExpressionLike {
	return v.condition
}

// This method sets the condition expression for this while clause.
func (v *whileClause) SetCondition(condition abs.ExpressionLike) {
	if condition == nil {
		panic("A while clause requires a condition expression.")
	}
	v.condition = condition
}

// This method returns the statement at the specified index from this while
// clause.
func (v *whileClause) GetStatement(index int) abs.StatementLike {
	return v.statements.GetValue(index)
}

// This method sets the statement at the specified index for this while clause.
func (v *whileClause) SetStatement(index int, statement abs.StatementLike) {
	if statement == nil {
		panic("Each index in a while clause requires a statement.")
	}
	v.statements.SetValue(index, statement)
}

// This method returns the list of statements for this while clause.
func (v *whileClause) GetStatements() abs.ProcedureLike {
	return v.statements
}

// This method sets the list of statements for this while clause.
func (v *whileClause) SetStatements(statements abs.ProcedureLike) {
	if statements == nil {
		panic("A while clause requires a list of statements.")
	}
	v.statements = statements
}
