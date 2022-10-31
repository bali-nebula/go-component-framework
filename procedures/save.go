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
func SaveClause(draft abs.Expression, recipient any) abs.SaveClauseLike {
	var v = &saveClause{}
	// Perform argument validation.
	v.SetDraft(draft)
	v.SetRecipient(recipient)
	return v
}

// This type defines the structure and methods associated with a save
// clause.
type saveClause struct {
	draft     abs.Expression
	recipient any
}

// This method returns the draft expression for this save clause.
func (v *saveClause) GetDraft() abs.Expression {
	return v.draft
}

// This method sets the draft expression for this save clause.
func (v *saveClause) SetDraft(draft abs.Expression) {
	if draft == nil {
		panic("A save clause requires a draft.")
	}
	v.draft = draft
}

// This method returns the document citation recipient for this save clause.
func (v *saveClause) GetRecipient() any {
	return v.recipient
}

// This method sets the document citation recipient for this save clause.
func (v *saveClause) SetRecipient(recipient any) {
	if recipient == nil {
		panic("A save clause requires a document citation recipient.")
	}
	v.recipient = recipient
}
