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

// IF CLAUSE IMPLEMENTATION

// This constructor creates a new if clause.
func IfClause(condition any, statements abs.ListLike[any]) abs.IfClauseLike {
	var v = &ifClause{}
	// Perform argument validation.
	v.SetCondition(condition)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with an if clause.
type ifClause struct {
	condition  any
	statements abs.ListLike[any]
}

// This method returns the condition expression for this if clause.
func (v *ifClause) GetCondition() any {
	return v.condition
}

// This method sets the condition expression for this if clause.
func (v *ifClause) SetCondition(condition any) {
	if condition == nil {
		panic("An if clause requires a condition expression.")
	}
	v.condition = condition
}

// This method returns the statement at the specified index from this if clause.
func (v *ifClause) GetStatement(index int) any {
	return v.statements.GetItem(index)
}

// This method sets the statement at the specified index for this if clause.
func (v *ifClause) SetStatement(index int, statement any) {
	if statement == nil {
		panic("Each statement in an if clause requires a value.")
	}
	v.statements.SetItem(index, statement)
}

// This method returns the list of statements for this if clause.
func (v *ifClause) GetStatements() abs.ListLike[any] {
	return v.statements
}

// This method sets the list of statements for this if clause.
func (v *ifClause) SetStatements(statements abs.ListLike[any]) {
	if statements == nil {
		panic("An if clause requires a list of statements.")
	}
	v.statements = statements
}
