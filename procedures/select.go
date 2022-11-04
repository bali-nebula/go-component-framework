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

// SELECT CLAUSE IMPLEMENTATION

// This constructor creates a new select clause.
func SelectClause(target abs.ExpressionLike, blocks abs.ListLike[abs.BlockLike]) abs.SelectClauseLike {
	var v = &selectClause{}
	// Perform argument validation.
	v.SetTarget(target)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with a select clause.
type selectClause struct {
	target abs.ExpressionLike
	blocks abs.ListLike[abs.BlockLike]
}

// This method returns the target expression for this select clause.
func (v *selectClause) GetTarget() abs.ExpressionLike {
	return v.target
}

// This method sets the target expression for this select clause.
func (v *selectClause) SetTarget(target abs.ExpressionLike) {
	if target == nil {
		panic("A select clause requires a target expression.")
	}
	v.target = target
}

// This method returns the list of blocks for this select clause.
func (v *selectClause) GetBlocks() abs.Sequential[abs.BlockLike] {
	return v.blocks
}

// This method sets the list of blocks for this select clause.
func (v *selectClause) SetBlocks(blocks abs.ListLike[abs.BlockLike]) {
	if blocks == nil || blocks.IsEmpty() {
		panic("A select clause requires at least one block.")
	}
	v.blocks = blocks
}

// This method returns the block at the specified index from this select clause.
func (v *selectClause) GetBlock(index int) abs.BlockLike {
	return v.blocks.GetValue(index)
}

// This method sets the block at the specified index for this select clause.
func (v *selectClause) SetBlock(index int, block abs.BlockLike) {
	if block == nil {
		panic("Each index in a select clause requires a block block.")
	}
	v.blocks.SetValue(index, block)
}
