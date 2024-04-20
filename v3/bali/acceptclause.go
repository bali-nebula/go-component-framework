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

var acceptClauseClass = &acceptClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func AcceptClause() AcceptClauseClassLike {
	return acceptClauseClass
}

// CLASS METHODS

// Target

type acceptClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *acceptClauseClass_) MakeWithMessage(message MessageLike) AcceptClauseLike {
	return &acceptClause_{
		message_: message,
	}
}

// Functions

// INSTANCE METHODS

// Target

type acceptClause_ struct {
	message_ MessageLike
}

// Attributes

func (v *acceptClause_) GetMessage() MessageLike {
	return v.message_
}

// Public

// Private
