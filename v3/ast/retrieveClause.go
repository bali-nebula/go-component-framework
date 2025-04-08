/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func RetrieveClauseClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Constructor Methods

func (c *retrieveClauseClass_) RetrieveClause(
	recipient RecipientLike,
	bag BagLike,
) RetrieveClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &retrieveClause_{
		// Initialize the instance attributes.
		recipient_: recipient,
		bag_:       bag,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *retrieveClause_) GetClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Attribute Methods

func (v *retrieveClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *retrieveClause_) GetBag() BagLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Instance Structure

type retrieveClause_ struct {
	// Declare the instance attributes.
	recipient_ RecipientLike
	bag_       BagLike
}

// Class Structure

type retrieveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func retrieveClauseClass() *retrieveClauseClass_ {
	return retrieveClauseClassReference_
}

var retrieveClauseClassReference_ = &retrieveClauseClass_{
	// Initialize the class constants.
}
