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

func CheckoutClauseClass() CheckoutClauseClassLike {
	return checkoutClauseClass()
}

// Constructor Methods

func (c *checkoutClauseClass_) CheckoutClause(
	recipient RecipientLike,
	optionalAtLevel AtLevelLike,
	citation CitationLike,
) CheckoutClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(citation) {
		panic("The \"citation\" attribute is required by this class.")
	}
	var instance = &checkoutClause_{
		// Initialize the instance attributes.
		recipient_:       recipient,
		optionalAtLevel_: optionalAtLevel,
		citation_:        citation,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *checkoutClause_) GetClass() CheckoutClauseClassLike {
	return checkoutClauseClass()
}

// Attribute Methods

func (v *checkoutClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *checkoutClause_) GetOptionalAtLevel() AtLevelLike {
	return v.optionalAtLevel_
}

func (v *checkoutClause_) GetCitation() CitationLike {
	return v.citation_
}

// PROTECTED INTERFACE

// Instance Structure

type checkoutClause_ struct {
	// Declare the instance attributes.
	recipient_       RecipientLike
	optionalAtLevel_ AtLevelLike
	citation_        CitationLike
}

// Class Structure

type checkoutClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func checkoutClauseClass() *checkoutClauseClass_ {
	return checkoutClauseClassReference_
}

var checkoutClauseClassReference_ = &checkoutClauseClass_{
	// Initialize the class constants.
}
