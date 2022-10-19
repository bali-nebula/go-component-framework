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

// IF CLAUSE IMPLEMENTATION

// This constructor creates a new if clause.
func If(block BlockLike) abs.IfLike {
	var v = &ifClause{}
	// Perform argument validation.
	v.SetBlock(block)
	return v
}

// This type defines the structure and methods associated with an if clause.
type ifClause struct {
	block BlockLike
}

// This method returns the statement block for this if clause.
func (v *ifClause) GetBlock() any {
	return v.block
}

// This method sets the statement block for this if clause.
func (v *ifClause) SetBlock(block any) {
	if block == nil {
		panic("An if clause requires a statement block.")
	}
	v.block = block
}
