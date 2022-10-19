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

// BLOCK IMPLEMENTATION

// This constructor creates a new block.
func Block(pattern any, statements abs.ListLike[any]) abs.BlockLike {
	var v = &block{}
	// Perform argument validation.
	v.SetPattern(pattern)
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a block of statements.
type block struct {
	pattern    any
	statements abs.ListLike[any]
}

// This method returns the pattern expression for this block.
func (v *block) GetPattern() any {
	return v.pattern
}

// This method sets the pattern expression for this block.
func (v *block) SetPattern(pattern any) {
	if pattern == nil {
		panic("A block requires a pattern expression.")
	}
	v.pattern = pattern
}

// This method returns the statement at the specified index from this block.
func (v *block) GetStatement(index int) any {
	return v.statements.GetItem(index)
}

// This method sets the statement at the specified index for this block.
func (v *block) SetStatement(index int, statement any) {
	if statement == nil {
		panic("Each statement in a block requires a value.")
	}
	v.statements.SetItem(index, statement)
}

// This method returns the list of statements for this block.
func (v *block) GetStatements() abs.ListLike[any] {
	return v.statements
}

// This method sets the list of statements for this block.
func (v *block) SetStatements(statements abs.ListLike[any]) {
	if statements == nil || statements.IsEmpty() {
		panic("A block requires at least one statement.")
	}
	v.statements = statements
}
