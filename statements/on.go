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

// ON CLAUSE IMPLEMENTATION

// This constructor creates a new on clause.
func On(exception abs.Symbolic, blocks abs.ListLike[abs.BlockLike]) abs.OnLike {
	var v = &onClause{}
	// Perform argument validation.
	v.SetException(exception)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with an on clause.
type onClause struct {
	exception abs.Symbolic
	blocks    abs.ListLike[abs.BlockLike]
}

// This method returns the expression for this on clause.
func (v *onClause) GetException() string {
	return v.variable
}

// This method sets the expression for this on clause.
func (v *onClause) SetException(exception abs.Symbolic) {
	if expression == nil {
		panic("An on clause requires an expression.")
	}
	v.expression = expression
}

// This method returns the block at the specified index from this on clause.
func (v *onClause) GetBlock(index int) abs.BlockLike {
	return v.blocks.GetItem(index)
}

// This method sets the expression at the specified index for this on clause.
func (v *onClause) SetBlock(index int, block abs.BlockLike) {
	if block == nil {
		panic("Each block in an on clause requires a value.")
	}
	v.blocks.SetItem(index, block)
}

// This method returns the list of blocks for this on clause.
func (v *onClause) GetBlocks() abs.ListLike[abs.BlockLike] {
	return v.blocks
}

// This method sets the list of blocks for this on clause.
func (v *onClause) SetBlocks(blocks abs.ListLike[abs.BlockLike]) {
	if blocks == nil || blocks.IsEmpty() {
		panic("An on clause requires at least one block.")
	}
	v.blocks = blocks
}
