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

// NOTARIZE CLAUSE IMPLEMENTATION

// This constructor creates a new notarize clause.
func Notarize(draft, moniker any) abs.NotarizeLike {
	var v = &notarizeClause{}
	// Perform argument validation.
	v.SetDraft(draft)
	v.SetMoniker(moniker)
	return v
}

// This type defines the structure and methods associated with a notarize
// clause.
type notarizeClause struct {
	draft   any
	moniker any
}

// This method returns the draft expression for this notarize clause.
func (v *notarizeClause) GetDraft() any {
	return v.draft
}

// This method sets the draft expression for this notarize clause.
func (v *notarizeClause) SetDraft(draft any) {
	if draft == nil {
		panic("A notarize clause requires a draft expression.")
	}
	v.draft = draft
}

// This method returns the citation moniker for this notarize clause.
func (v *notarizeClause) GetMoniker() any {
	return v.moniker
}

// This method sets the citation moniker for this notarize clause.
func (v *notarizeClause) SetMoniker(moniker any) {
	if moniker == nil {
		panic("A notarize clause requires a citation moniker.")
	}
	v.moniker = moniker
}
