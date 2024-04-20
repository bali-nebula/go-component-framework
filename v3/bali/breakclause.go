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

var breakClauseClass = &breakClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func BreakClause() BreakClauseClassLike {
	return breakClauseClass
}

// CLASS METHODS

// Target

type breakClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *breakClauseClass_) Make() BreakClauseLike {
	return &breakClause_{}
}

// Functions

// INSTANCE METHODS

// Target

type breakClause_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

// Private
