/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import ()

// CLASS ACCESS

// Reference

var checkoutClauseClass = &checkoutClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func CheckoutClause() CheckoutClauseClassLike {
	return checkoutClauseClass
}

// CLASS METHODS

// Target

type checkoutClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *checkoutClauseClass_) MakeWithAttributes(
	recipient RecipientLike,
	level LevelLike,
	citation CitationLike,
) CheckoutClauseLike {
	return &checkoutClause_{
		recipient_: recipient,
		level_:     level,
		citation_:  citation,
	}
}

// Functions

// INSTANCE METHODS

// Target

type checkoutClause_ struct {
	recipient_ RecipientLike
	level_     LevelLike
	citation_  CitationLike
}

// Attributes

func (v *checkoutClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *checkoutClause_) GetLevel() LevelLike {
	return v.level_
}

func (v *checkoutClause_) GetCitation() CitationLike {
	return v.citation_
}

// Public

// Private
