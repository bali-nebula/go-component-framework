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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// SAVE CLAUSE IMPLEMENTATION

// This constructor creates a new save clause.
func SaveClause(document abs.Expression, recipient abs.Recipient) abs.SaveClauseLike {
	var v = &saveClause{}
	// Perform argument validation.
	v.SetDocument(document)
	v.SetRecipient(recipient)
	return v
}

// This type defines the structure and methods associated with a save
// clause.
type saveClause struct {
	document  abs.Expression
	recipient abs.Recipient
}

// This method is a dummy method that always returns true.
func (v *saveClause) IsSaveClause() bool {
	return true
}

// This method returns the document expression for this save clause.
func (v *saveClause) GetDocument() abs.Expression {
	return v.document
}

// This method sets the document expression for this save clause.
func (v *saveClause) SetDocument(document abs.Expression) {
	if document == nil {
		panic("A save clause requires a document.")
	}
	v.document = document
}

// This method returns the document citation recipient for this save clause.
func (v *saveClause) GetRecipient() abs.Recipient {
	return v.recipient
}

// This method sets the document citation recipient for this save clause.
func (v *saveClause) SetRecipient(recipient abs.Recipient) {
	if recipient == nil {
		panic("A save clause requires a document citation recipient.")
	}
	v.recipient = recipient
}
