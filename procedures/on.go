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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
)

// ON CLAUSE IMPLEMENTATION

// This constructor creates a new on clause.
func OnClause(exception abs.Symbolic, blocks abs.Blocks) abs.OnClauseLike {
	var v = &onClause{}
	// Perform argument validation.
	v.SetException(exception)
	v.SetBlocks(blocks)
	return v
}

// This type defines the structure and methods associated with an on clause.
type onClause struct {
	exception abs.Symbolic
	blocks    abs.Blocks
}

// This method returns the exception symbol for this on clause.
func (v *onClause) GetException() abs.Symbolic {
	return v.exception
}

// This method sets the exception symbol for this on clause.
func (v *onClause) SetException(exception abs.Symbolic) {
	if exception == nil {
		panic("An on clause requires an exception symbol.")
	}
	v.exception = exception
}

// This method returns the list of blocks for this on clause.
func (v *onClause) GetBlocks() abs.Blocks {
	return v.blocks
}

// This method sets the list of blocks for this on clause.
func (v *onClause) SetBlocks(blocks abs.Blocks) {
	if blocks == nil || blocks.IsEmpty() {
		panic("An on clause requires at least one block.")
	}
	v.blocks = blocks
}
