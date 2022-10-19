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
func WithClause(item abs.Symbolic, block abs.BlockLike) abs.WithClauseLike {
	var v = &withClause{}
	// Perform argument validation.
	v.SetItem(item)
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with a with clause.
type withClause struct {
	item  abs.Symbolic
	block abs.BlockLike
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

// This method returns the block at the specified index from this with clause.
func (v *withClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block at the specified index for this with clause.
func (v *withClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("The block in a with clause requires a value.")
	}
	v.block = block
}
