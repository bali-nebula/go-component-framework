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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// ON CLAUSE IMPLEMENTATION

// This constructor creates a new on clause.
func OnClause(failure abs.SymbolLike, blocks abs.Sequential[abs.BlockLike]) abs.OnClauseLike {
	var v = &onClause{}
	// Perform argument validation.
	v.SetFailure(failure)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with an on clause.
type onClause struct {
	failure abs.SymbolLike
	blocks  abs.Sequential[abs.BlockLike]
}

// This method returns the symbol for the failure for this on clause.
func (v *onClause) GetFailure() abs.SymbolLike {
	return v.failure
}

// This method sets the symbol for the failure for this on clause.
func (v *onClause) SetFailure(failure abs.SymbolLike) {
	if failure == nil {
		panic("An on clause requires a symbol for the failure.")
	}
	v.failure = failure
}

// This method returns the list of blocks for this on clause.
func (v *onClause) GetBlocks() abs.Sequential[abs.BlockLike] {
	return v.blocks
}

// This method sets the list of blocks for this on clause.
func (v *onClause) SetBlocks(blocks abs.Sequential[abs.BlockLike]) {
	if blocks == nil || blocks.IsEmpty() {
		panic("An on clause requires at least one block.")
	}
	v.blocks = blocks
}
