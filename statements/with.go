/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package statements

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// WITH CLAUSE IMPLEMENTATION

// This constructor creates a new with clause.
func WithClause(item abs.Symbolic, sequence any, statements abs.ProcedureLike) abs.WithClauseLike {
	var v = &withClause{}
	// Perform argument validation.
	v.SetItem(item)
	v.SetSequence(sequence)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a with clause.
type withClause struct {
	item       abs.Symbolic
	sequence   any
	statements abs.ProcedureLike
}

// This method returns the item symbol for this with clause.
func (v *withClause) GetItem() abs.Symbolic {
	return v.item
}

// This method sets the item symbol for this with clause.
func (v *withClause) SetItem(item abs.Symbolic) {
	if item == nil {
		panic("A with clause requires an item symbol.")
	}
	v.item = item
}

// This method returns the sequence expression for this with clause.
func (v *withClause) GetSequence() any {
	return v.sequence
}

// This method sets the sequence expression for this with clause.
func (v *withClause) SetSequence(sequence any) {
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
		panic("Each statement in a with clause requires a value.")
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
