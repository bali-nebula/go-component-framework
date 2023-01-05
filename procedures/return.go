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

// RETURN CLAUSE IMPLEMENTATION

// This constructor creates a new return clause.
func ReturnClause(result abs.Expression) abs.ReturnClauseLike {
	var v = &returnClause{}
	// Perform argument validation.
	v.SetResult(result)
	return v
}

// This type defines the structure and methods associated with a return
// clause.
type returnClause struct {
	result abs.Expression
}

// This method returns the result expression for this return clause.
func (v *returnClause) GetResult() abs.Expression {
	return v.result
}

// This method sets the result expression for this return clause.
func (v *returnClause) SetResult(result abs.Expression) {
	if result == nil {
		panic("A return clause requires a result.")
	}
	v.result = result
}
