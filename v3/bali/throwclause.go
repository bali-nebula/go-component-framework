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

var throwClauseClass = &throwClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func ThrowClause() ThrowClauseClassLike {
	return throwClauseClass
}

// CLASS METHODS

// Target

type throwClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *throwClauseClass_) MakeWithException(exception ExceptionLike) ThrowClauseLike {
	return &throwClause_{
		exception_: exception,
	}
}

// Functions

// INSTANCE METHODS

// Target

type throwClause_ struct {
	exception_ ExceptionLike
}

// Attributes

func (v *throwClause_) GetException() ExceptionLike {
	return v.exception_
}

// Public

// Private
