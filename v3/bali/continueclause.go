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

var continueClauseClass = &continueClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func ContinueClause() ContinueClauseClassLike {
	return continueClauseClass
}

// CLASS METHODS

// Target

type continueClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *continueClauseClass_) Make() ContinueClauseLike {
	return &continueClause_{}
}

// Functions

// INSTANCE METHODS

// Target

type continueClause_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

// Private
