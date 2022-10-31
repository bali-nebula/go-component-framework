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

// RETRIEVE CLAUSE IMPLEMENTATION

// This constructor creates a new retrieve clause.
func RetrieveClause(recipient abs.Recipient, bag abs.Expression) abs.RetrieveClauseLike {
	var v = &retrieveClause{}
	// Perform argument validation.
	v.SetRecipient(recipient)
	v.SetBag(bag)
	return v
}

// This type defines the structure and methods associated with a retrieve
// clause.
type retrieveClause struct {
	recipient abs.Recipient
	bag       abs.Expression
}

// This method returns the recipient expression for this retrieve clause.
func (v *retrieveClause) GetRecipient() abs.Recipient {
	return v.recipient
}

// This method sets the recipient expression for this retrieve clause.
func (v *retrieveClause) SetRecipient(recipient abs.Recipient) {
	if recipient == nil {
		panic("A retrieve clause requires a recipient.")
	}
	v.recipient = recipient
}

// This method returns the message bag expression for this retrieve clause.
func (v *retrieveClause) GetBag() abs.Expression {
	return v.bag
}

// This method sets the message bag expression for this retrieve clause.
func (v *retrieveClause) SetBag(bag abs.Expression) {
	if bag == nil {
		panic("A retrieve clause requires a message bag.")
	}
	v.bag = bag
}
