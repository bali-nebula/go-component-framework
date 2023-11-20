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

// NOTARIZE CLAUSE IMPLEMENTATION

// This constructor creates a new notarize clause.
func NotarizeClause(document, name abs.Expression) abs.NotarizeClauseLike {
	var v = &notarizeClause{}
	// Perform argument validation.
	v.SetDocument(document)
	v.SetName(name)
	return v
}

// This type defines the structure and methods associated with a notarize
// clause.
type notarizeClause struct {
	document abs.Expression
	name     abs.Expression
}

// This method returns the document expression for this notarize clause.
func (v *notarizeClause) GetDocument() abs.Expression {
	return v.document
}

// This method sets the document expression for this notarize clause.
func (v *notarizeClause) SetDocument(document abs.Expression) {
	if document == nil {
		panic("A notarize clause requires a document expression.")
	}
	v.document = document
}

// This method returns the citation name for this notarize clause.
func (v *notarizeClause) GetName() abs.Expression {
	return v.name
}

// This method sets the citation name for this notarize clause.
func (v *notarizeClause) SetName(name abs.Expression) {
	if name == nil {
		panic("A notarize clause requires a citation name.")
	}
	v.name = name
}
