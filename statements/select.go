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

// SELECT CLAUSE IMPLEMENTATION

// This constructor creates a new select clause.
func SelectClause(control any, blocks abs.ListLike[abs.BlockLike]) abs.SelectClauseLike {
	var v = &selectClause{}
	// Perform argument validation.
	v.SetControl(control)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with a select clause.
type selectClause struct {
	control any
	blocks  abs.ListLike[abs.BlockLike]
}

// This method returns the control expression for this select clause.
func (v *selectClause) GetControl() any {
	return v.control
}

// This method sets the control expression for this select clause.
func (v *selectClause) SetControl(control any) {
	if control == nil {
		panic("A select clause requires an control expression.")
	}
	v.control = control
}

// This method returns the block at the specified index from this select clause.
func (v *selectClause) GetBlock(index int) abs.BlockLike {
	return v.blocks.GetItem(index)
}

// This method sets the block at the specified index for this select clause.
func (v *selectClause) SetBlock(index int, block abs.BlockLike) {
	if block == nil {
		panic("Each block in a select clause requires a value.")
	}
	v.blocks.SetItem(index, block)
}

// This method returns the list of blocks for this select clause.
func (v *selectClause) GetBlocks() abs.ListLike[abs.BlockLike] {
	return v.blocks
}

// This method sets the list of blocks for this select clause.
func (v *selectClause) SetBlocks(blocks abs.ListLike[abs.BlockLike]) {
	if blocks == nil || blocks.IsEmpty() {
		panic("A select clause requires at least one block.")
	}
	v.blocks = blocks
}
