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

var rejectClauseClass = &rejectClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func RejectClause() RejectClauseClassLike {
	return rejectClauseClass
}

// CLASS METHODS

// Target

type rejectClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *rejectClauseClass_) MakeWithMessage(message MessageLike) RejectClauseLike {
	return &rejectClause_{
		message_: message,
	}
}

// Functions

// INSTANCE METHODS

// Target

type rejectClause_ struct {
	message_ MessageLike
}

// Attributes

func (v *rejectClause_) GetMessage() MessageLike {
	return v.message_
}

// Public

// Private
