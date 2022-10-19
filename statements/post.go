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

// POST CLAUSE IMPLEMENTATION

// This constructor creates a new post clause.
func PostClause(message, bag any) abs.PostClauseLike {
	var v = &postClause{}
	// Perform argument validation.
	v.SetMessage(message)
	v.SetBag(bag)
	return v
}

// This type defines the structure and methods associated with a post
// clause.
type postClause struct {
	message any
	bag     any
}

// This method returns the message expression for this post clause.
func (v *postClause) GetMessage() any {
	return v.message
}

// This method sets the message expression for this post clause.
func (v *postClause) SetMessage(message any) {
	if message == nil {
		panic("A post clause requires a message.")
	}
	v.message = message
}

// This method returns the message bag expression for this post clause.
func (v *postClause) GetBag() any {
	return v.bag
}

// This method sets the message bag expression for this post clause.
func (v *postClause) SetBag(bag any) {
	if bag == nil {
		panic("A post clause requires a message bag.")
	}
	v.bag = bag
}
