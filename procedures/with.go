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
func WithClause(value abs.Symbolic, sequence abs.ExpressionLike, statements abs.ProcedureLike) abs.WithClauseLike {
	var v = &withClause{}
	// Perform argument validation.
	v.SetValue(value)
	v.SetSequence(sequence)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a with clause.
type withClause struct {
	value      abs.Symbolic
	sequence   abs.ExpressionLike
	statements abs.ProcedureLike
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

// This method returns the sequence expression for this with clause.
func (v *withClause) GetSequence() abs.ExpressionLike {
	return v.sequence
}

// This method sets the sequence expression for this with clause.
func (v *withClause) SetSequence(sequence abs.ExpressionLike) {
	if sequence == nil {
		panic("A with clause requires a sequence expression.")
	}
	v.sequence = sequence
}

// This method returns the statement at the specified index from this while
// clause.
func (v *withClause) GetStatement(index int) abs.StatementLike {
	return v.statements.GetItem(index)
}

// This method sets the statement at the specified index for this with clause.
func (v *withClause) SetStatement(index int, statement abs.StatementLike) {
	if statement == nil {
		panic("Each index in a with clause requires a statement.")
	}
	v.statements.SetItem(index, statement)
}

// This method returns the list of statements for this with clause.
func (v *withClause) GetStatements() abs.ProcedureLike {
	return v.statements
}

// This method sets the list of statements for this with clause.
func (v *withClause) SetStatements(statements abs.ProcedureLike) {
	if statements == nil {
		panic("A with clause requires a list of statements.")
	}
	v.statements = statements
}
