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

var postClauseClass = &postClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func PostClause() PostClauseClassLike {
	return postClauseClass
}

// CLASS METHODS

// Target

type postClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *postClauseClass_) MakeWithAttributes(
	message MessageLike,
	bag BagLike,
) PostClauseLike {
	return &postClause_{
		message_: message,
		bag_:     bag,
	}
}

// Functions

// INSTANCE METHODS

// Target

type postClause_ struct {
	message_ MessageLike
	bag_     BagLike
}

// Attributes

func (v *postClause_) GetMessage() MessageLike {
	return v.message_
}

func (v *postClause_) GetBag() BagLike {
	return v.bag_
}

// Public

// Private
