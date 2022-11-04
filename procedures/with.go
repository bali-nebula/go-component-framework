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

// WITH CLAUSE IMPLEMENTATION

// This constructor creates a new with clause.
func WithClause(value abs.Symbolic, block abs.BlockLike) abs.WithClauseLike {
	var v = &withClause{}
	// Perform argument validation.
	v.SetValue(value)
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with a with clause.
type withClause struct {
	value abs.Symbolic
	block abs.BlockLike
}

// This method returns the value symbol for this with clause.
func (v *withClause) GetValue() abs.Symbolic {
	return v.value
}

// This method sets the value symbol for this with clause.
func (v *withClause) SetValue(value abs.Symbolic) {
	if value == nil {
		panic("A with clause requires an value symbol.")
	}
	v.value = value
}

// This method returns the block for this with clause.
func (v *withClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block for this with clause.
func (v *withClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("A with clause requires a block.")
	}
	v.block = block
}

// This method returns the sequence expression for this with clause.
func (v *withClause) GetSequence() abs.ExpressionLike {
	return v.block.GetExpression()
}

// This method sets the sequence expression for this with clause.
func (v *withClause) SetSequence(sequence abs.ExpressionLike) {
	v.block.SetExpression(sequence)
}

// This method returns the statement at the specified index from this with clause.
func (v *withClause) GetStatement(index int) abs.StatementLike {
	return v.block.GetStatement(index)
}

// This method sets the statement at the specified index for this with clause.
func (v *withClause) SetStatement(index int, statement abs.StatementLike) {
	v.block.SetStatement(index, statement)
}

// This method returns the list of statements for this with clause.
func (v *withClause) GetProcedure() abs.ProcedureLike {
	return v.block.GetProcedure()
}

// This method sets the list of statements for this with clause.
func (v *withClause) SetProcedure(statements abs.ProcedureLike) {
	v.block.SetProcedure(statements)
}
