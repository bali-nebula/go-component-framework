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

// REJECT CLAUSE IMPLEMENTATION

// This constructor creates a new reject clause.
func RejectClause(message abs.Expression) abs.RejectClauseLike {
	var v = &rejectClause{}
	// Perform argument validation.
	v.SetMessage(message)
	return v
}

// This type defines the structure and methods associated with a reject
// clause.
type rejectClause struct {
	message abs.Expression
}

// This method returns the message expression for this reject clause.
func (v *rejectClause) GetMessage() abs.Expression {
	return v.message
}

// This method sets the message expression for this reject clause.
func (v *rejectClause) SetMessage(message abs.Expression) {
	if message == nil {
		panic("An reject clause requires a message.")
	}
	v.message = message
}
