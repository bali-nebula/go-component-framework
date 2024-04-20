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

var discardClauseClass = &discardClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func DiscardClause() DiscardClauseClassLike {
	return discardClauseClass
}

// CLASS METHODS

// Target

type discardClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *discardClauseClass_) MakeWithDraft(draft DraftLike) DiscardClauseLike {
	return &discardClause_{
		draft_: draft,
	}
}

// Functions

// INSTANCE METHODS

// Target

type discardClause_ struct {
	draft_ DraftLike
}

// Attributes

func (v *discardClause_) GetDraft() DraftLike {
	return v.draft_
}

// Public

// Private
