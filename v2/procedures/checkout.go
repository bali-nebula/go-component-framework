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

// CHECKOUT CLAUSE IMPLEMENTATION

// This constructor creates a new checkout clause.
func CheckoutClause(recipient abs.Recipient, level abs.Expression, name abs.Expression) abs.CheckoutClauseLike {
	var v = &checkoutClause{}
	// Perform argument validation.
	v.SetRecipient(recipient)
	v.SetLevel(level)
	v.SetName(name)
	return v
}

// This type defines the structure and methods associated with a checkout
// clause.
type checkoutClause struct {
	recipient abs.Recipient
	level     abs.Expression // The version level to be incremented (optional).
	name      abs.Expression // A name to the citation for the document to be checked out.
}

// This method returns the recipient for this checkout clause.
func (v *checkoutClause) GetRecipient() abs.Recipient {
	return v.recipient
}

// This method sets the recipient for this checkout clause.
func (v *checkoutClause) SetRecipient(recipient abs.Recipient) {
	if recipient == nil {
		panic("A checkout clause requires a recipient.")
	}
	v.recipient = recipient
}

// This method returns the level for this checkout clause.
func (v *checkoutClause) GetLevel() abs.Expression {
	return v.level
}

// This method sets the level for this checkout clause.
func (v *checkoutClause) SetLevel(level abs.Expression) {
	v.level = level
}

// This method returns the citation name for this checkout clause.
func (v *checkoutClause) GetName() abs.Expression {
	return v.name
}

// This method sets the citation name for this checkout clause.
func (v *checkoutClause) SetName(name abs.Expression) {
	if name == nil {
		panic("A checkout clause requires a citation name.")
	}
	v.name = name
}
