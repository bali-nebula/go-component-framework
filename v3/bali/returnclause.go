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

var returnClauseClass = &returnClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func ReturnClause() ReturnClauseClassLike {
	return returnClauseClass
}

// CLASS METHODS

// Target

type returnClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *returnClauseClass_) MakeWithResult(result ResultLike) ReturnClauseLike {
	return &returnClause_{
		result_: result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type returnClause_ struct {
	result_ ResultLike
}

// Attributes

func (v *returnClause_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private
