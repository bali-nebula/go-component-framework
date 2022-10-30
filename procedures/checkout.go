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

// CHECKOUT CLAUSE IMPLEMENTATION

// This constructor creates a new checkout clause.
func CheckoutClause(recipient, level, moniker any) abs.CheckoutClauseLike {
	var v = &checkoutClause{}
	// Perform argument validation.
	v.SetRecipient(recipient)
	v.SetLevel(level)
	v.SetMoniker(moniker)
	return v
}

// This type defines the structure and methods associated with a checkout
// clause.
type checkoutClause struct {
	recipient any
	level     any // The version level to be incremented (optional).
	moniker   any // A moniker to the citation for the document to be checked out.
}

// This method returns the recipient for this checkout clause.
func (v *checkoutClause) GetRecipient() any {
	return v.recipient
}

// This method sets the recipient for this checkout clause.
func (v *checkoutClause) SetRecipient(recipient any) {
	if recipient == nil {
		panic("A checkout clause requires a recipient.")
	}
	v.recipient = recipient
}

// This method returns the level for this checkout clause.
func (v *checkoutClause) GetLevel() any {
	return v.level
}

// This method sets the level for this checkout clause.
func (v *checkoutClause) SetLevel(level any) {
	v.level = level
}

// This method returns the citation moniker for this checkout clause.
func (v *checkoutClause) GetMoniker() any {
	return v.moniker
}

// This method sets the citation moniker for this checkout clause.
func (v *checkoutClause) SetMoniker(moniker any) {
	if moniker == nil {
		panic("A checkout clause requires a citation moniker.")
	}
	v.moniker = moniker
}
