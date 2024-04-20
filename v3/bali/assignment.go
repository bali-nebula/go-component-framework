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

var assignmentClass = &assignmentClass_{
	// This class has no private constants to initialize.
}

// Function

func Assignment() AssignmentClassLike {
	return assignmentClass
}

// CLASS METHODS

// Target

type assignmentClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *assignmentClass_) MakeWithLetClause(letClause LetClauseLike) AssignmentLike {
	return &assignment_{
		letClause_: letClause,
	}
}

// Functions

// INSTANCE METHODS

// Target

type assignment_ struct {
	letClause_ LetClauseLike
}

// Attributes

func (v *assignment_) GetLetClause() LetClauseLike {
	return v.letClause_
}

// Public

// Private
