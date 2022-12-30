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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// IF CLAUSE IMPLEMENTATION

// This constructor creates a new if clause.
func IfClause(block abs.BlockLike) abs.IfClauseLike {
	var v = &ifClause{}
	// Perform argument validation.
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with an if clause.
type ifClause struct {
	block abs.BlockLike
}

// This method is a dummy method that always returns true.
func (v *ifClause) IsIfClause() bool {
	return true
}

// This method returns the block for this if clause.
func (v *ifClause) GetBlock() abs.BlockLike {
	return v.block
}

// This method sets the block for this if clause.
func (v *ifClause) SetBlock(block abs.BlockLike) {
	if block == nil {
		panic("An if clause requires a block.")
	}
	v.block = block
}
