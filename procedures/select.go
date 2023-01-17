/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// SELECT CLAUSE IMPLEMENTATION

// This constructor creates a new select clause.
func SelectClause(target abs.Expression, blocks abs.Sequential[abs.BlockLike]) abs.SelectClauseLike {
	var v = &selectClause{}
	// Perform argument validation.
	v.SetTarget(target)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with a select clause.
type selectClause struct {
	target abs.Expression
	blocks abs.Sequential[abs.BlockLike]
}

// This method returns the target expression for this select clause.
func (v *selectClause) GetTarget() abs.Expression {
	return v.target
}

// This method sets the target expression for this select clause.
func (v *selectClause) SetTarget(target abs.Expression) {
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
func (v *selectClause) SetBlocks(blocks abs.Sequential[abs.BlockLike]) {
	if blocks == nil || blocks.IsEmpty() {
		panic("A select clause requires at least one block.")
	}
	v.blocks = blocks
}
