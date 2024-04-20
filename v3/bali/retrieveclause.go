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

var retrieveClauseClass = &retrieveClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func RetrieveClause() RetrieveClauseClassLike {
	return retrieveClauseClass
}

// CLASS METHODS

// Target

type retrieveClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *retrieveClauseClass_) MakeWithAttributes(
	recipient RecipientLike,
	bag BagLike,
) RetrieveClauseLike {
	return &retrieveClause_{
		recipient_: recipient,
		bag_:       bag,
	}
}

// Functions

// INSTANCE METHODS

// Target

type retrieveClause_ struct {
	recipient_ RecipientLike
	bag_       BagLike
}

// Attributes

func (v *retrieveClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *retrieveClause_) GetBag() BagLike {
	return v.bag_
}

// Public

// Private
