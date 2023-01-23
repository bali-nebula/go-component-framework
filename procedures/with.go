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

// WITH CLAUSE IMPLEMENTATION

// This constructor creates a new with clause.
func WithClause(item abs.SymbolLike, block abs.BlockLike) abs.WithClauseLike {
	var v = &withClause{}
	// Perform argument validation.
	v.SetItem(item)
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with a with clause.
type withClause struct {
	item  abs.SymbolLike
	block abs.BlockLike
}

// This method returns the symbol for the item for this with clause.
func (v *withClause) GetItem() abs.SymbolLike {
	return v.item
}

// This method sets the symbol for the item for this with clause.
func (v *withClause) SetItem(item abs.SymbolLike) {
	if item == nil {
		panic("A with clause requires a symbol for the item.")
	}
	v.item = item
}

// This method returns the block for this with clause.
func (v *withClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block for this with clause.
func (v *withClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("A with clause requires a block.")
	}
	v.block = block
}
