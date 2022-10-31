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

// BLOCK IMPLEMENTATION

// This constructor creates a new block.
func Block(expression abs.ExpressionLike, procedure abs.ProcedureLike) abs.BlockLike {
	var v = &block{}
	// Perform argument validation.
	v.SetExpression(expression)
	v.SetProcedure(procedure)
	return v
}

// This type defines the structure and methods associated with a procedure block.
type block struct {
	expression abs.ExpressionLike
	procedure  abs.ProcedureLike
}

// This method returns the expression for this block.
func (v *block) GetExpression() abs.ExpressionLike {
	return v.expression
}

// This method sets the expression for this block.
func (v *block) SetExpression(expression abs.ExpressionLike) {
	if expression == nil {
		panic("A block requires an expression.")
	}
	v.expression = expression
}

// This method returns the statement at the specified index from this block.
func (v *block) GetStatement(index int) abs.StatementLike {
	return v.procedure.GetItem(index)
}

// This method sets the statement at the specified index for this block.
func (v *block) SetStatement(index int, statement abs.StatementLike) {
	if statement == nil {
		panic("Each statement in a block requires a value.")
	}
	v.procedure.SetItem(index, statement)
}

// This method returns the procedure for this block.
func (v *block) GetProcedure() abs.ProcedureLike {
	return v.procedure
}

// This method sets the procedure for this block.
func (v *block) SetProcedure(procedure abs.ProcedureLike) {
	if procedure == nil || procedure.IsEmpty() {
		panic("A block requires at least one statement.")
	}
	v.procedure = procedure
}
