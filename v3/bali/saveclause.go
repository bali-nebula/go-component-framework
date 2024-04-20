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

var saveClauseClass = &saveClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func SaveClause() SaveClauseClassLike {
	return saveClauseClass
}

// CLASS METHODS

// Target

type saveClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *saveClauseClass_) MakeWithAttributes(
	draft DraftLike,
	citation CitationLike,
) SaveClauseLike {
	return &saveClause_{
		draft_:    draft,
		citation_: citation,
	}
}

// Functions

// INSTANCE METHODS

// Target

type saveClause_ struct {
	draft_    DraftLike
	citation_ CitationLike
}

// Attributes

func (v *saveClause_) GetDraft() DraftLike {
	return v.draft_
}

func (v *saveClause_) GetCitation() CitationLike {
	return v.citation_
}

// Public

// Private
