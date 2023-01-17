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

// DISCARD CLAUSE IMPLEMENTATION

// This constructor creates a new discard clause.
func DiscardClause(document abs.Expression) abs.DiscardClauseLike {
	var v = &discardClause{}
	// Perform argument validation.
	v.SetDocument(document)
	return v
}

// This type defines the structure and methods associated with a discard
// clause.
type discardClause struct {
	document abs.Expression
}

// This method returns the document expression for this discard clause.
func (v *discardClause) GetDocument() abs.Expression {
	return v.document
}

// This method sets the document expression for this discard clause.
func (v *discardClause) SetDocument(document abs.Expression) {
	if document == nil {
		panic("A discard clause requires a document.")
	}
	v.document = document
}
