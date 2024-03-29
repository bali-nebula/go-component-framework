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

// WHILE CLAUSE IMPLEMENTATION

// This constructor creates a new while clause.
func WhileClause(block abs.BlockLike) abs.WhileClauseLike {
	var v = &whileClause{}
	// Perform argument validation.
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with a while clause.
type whileClause struct {
	block abs.BlockLike
}

// This method returns the block for this while clause.
func (v *whileClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block for this while clause.
func (v *whileClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("A while clause requires a block.")
	}
	v.block = block
}
