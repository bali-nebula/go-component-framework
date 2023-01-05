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

// POST CLAUSE IMPLEMENTATION

// This constructor creates a new post clause.
func PostClause(message, bag abs.Expression) abs.PostClauseLike {
	var v = &postClause{}
	// Perform argument validation.
	v.SetMessage(message)
	v.SetBag(bag)
	return v
}

// This type defines the structure and methods associated with a post
// clause.
type postClause struct {
	message abs.Expression
	bag     abs.Expression
}

// This method returns the message expression for this post clause.
func (v *postClause) GetMessage() abs.Expression {
	return v.message
}

// This method sets the message expression for this post clause.
func (v *postClause) SetMessage(message abs.Expression) {
	if message == nil {
		panic("A post clause requires a message.")
	}
	v.message = message
}

// This method returns the message bag expression for this post clause.
func (v *postClause) GetBag() abs.Expression {
	return v.bag
}

// This method sets the message bag expression for this post clause.
func (v *postClause) SetBag(bag abs.Expression) {
	if bag == nil {
		panic("A post clause requires a message bag.")
	}
	v.bag = bag
}
