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

// IF CLAUSE IMPLEMENTATION

// This constructor creates a new if clause.
func IfClause(condition abs.ExpressionLike, statements abs.ProcedureLike) abs.IfClauseLike {
	var v = &ifClause{}
	// Perform argument validation.
	v.SetCondition(condition)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with an if clause.
type ifClause struct {
	condition  abs.ExpressionLike
	statements abs.ProcedureLike
}

// This method returns the condition expression for this if clause.
func (v *ifClause) GetCondition() abs.ExpressionLike {
	return v.condition
}

// This method sets the condition expression for this if clause.
func (v *ifClause) SetCondition(condition abs.ExpressionLike) {
	if condition == nil {
		panic("An if clause requires a condition expression.")
	}
	v.condition = condition
}

// This method returns the statement at the specified index from this if clause.
func (v *ifClause) GetStatement(index int) abs.StatementLike {
	return v.statements.GetItem(index)
}

// This method sets the statement at the specified index for this if clause.
func (v *ifClause) SetStatement(index int, statement abs.StatementLike) {
	if statement == nil {
		panic("Each index in an if clause requires a statement.")
	}
	v.statements.SetItem(index, statement)
}

// This method returns the list of statements for this if clause.
func (v *ifClause) GetStatements() abs.ProcedureLike {
	return v.statements
}

// This method sets the list of statements for this if clause.
func (v *ifClause) SetStatements(statements abs.ProcedureLike) {
	if statements == nil {
		panic("An if clause requires a list of statements.")
	}
	v.statements = statements
}
