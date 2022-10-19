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

// WHILE CLAUSE IMPLEMENTATION

// This constructor creates a new while clause.
func WhileClause(condition any, statements abs.ListLike[any]) abs.WhileClauseLike {
	var v = &whileClause{}
	// Perform argument validation.
	v.SetCondition(condition)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a while clause.
type whileClause struct {
	condition  any
	statements abs.ListLike[any]
}

// This method returns the condition expression for this while clause.
func (v *whileClause) GetCondition() any {
	return v.condition
}

// This method sets the condition expression for this while clause.
func (v *whileClause) SetCondition(condition any) {
	if condition == nil {
		panic("A while clause requires a condition expression.")
	}
	v.condition = condition
}

// This method returns the statement at the specified index from this while
// clause.
func (v *whileClause) GetStatement(index int) any {
	return v.statements.GetItem(index)
}

// This method sets the statement at the specified index for this while clause.
func (v *whileClause) SetStatement(index int, statement any) {
	if statement == nil {
		panic("Each statement in a while clause requires a value.")
	}
	v.statements.SetItem(index, statement)
}

// This method returns the list of statements for this while clause.
func (v *whileClause) GetStatements() abs.ListLike[any] {
	return v.statements
}

// This method sets the list of statements for this while clause.
func (v *whileClause) SetStatements(statements abs.ListLike[any]) {
	if statements == nil {
		panic("A while clause requires a list of statements.")
	}
	v.statements = statements
}
