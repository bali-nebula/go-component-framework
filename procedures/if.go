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
func IfClause(block abs.BlockLike) abs.IfClauseLike {
	var v = &ifClause{}
	// Perform argument validation.
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with an if clause.
type ifClause struct {
	block abs.BlockLike
}

// This method returns the block for this if clause.
func (v *ifClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block for this if clause.
func (v *ifClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("An if clause requires a block.")
	}
	v.block = block
}

// This method returns the condition expression for this if clause.
func (v *ifClause) GetCondition() abs.ExpressionLike {
	return v.block.GetExpression()
}

// This method sets the condition expression for this if clause.
func (v *ifClause) SetCondition(condition abs.ExpressionLike) {
	v.block.SetExpression(condition)
}

// This method returns the statement at the specified index from this if clause.
func (v *ifClause) GetStatement(index int) abs.StatementLike {
	return v.block.GetStatement(index)
}

// This method sets the statement at the specified index for this if clause.
func (v *ifClause) SetStatement(index int, statement abs.StatementLike) {
	v.block.SetStatement(index, statement)
}

// This method returns the list of statements for this if clause.
func (v *ifClause) GetProcedure() abs.ProcedureLike {
	return v.block.GetProcedure()
}

// This method sets the list of statements for this if clause.
func (v *ifClause) SetProcedure(statements abs.ProcedureLike) {
	v.block.SetProcedure(statements)
}
