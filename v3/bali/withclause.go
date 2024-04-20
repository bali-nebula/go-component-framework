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

var withClauseClass = &withClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func WithClause() WithClauseClassLike {
	return withClauseClass
}

// CLASS METHODS

// Target

type withClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *withClauseClass_) MakeWithAttributes(
	item ItemLike,
	sequence SequenceLike,
	procedure ProcedureLike,
) WithClauseLike {
	return &withClause_{
		item_:      item,
		sequence_:  sequence,
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type withClause_ struct {
	item_      ItemLike
	sequence_  SequenceLike
	procedure_ ProcedureLike
}

// Attributes

func (v *withClause_) GetItem() ItemLike {
	return v.item_
}

func (v *withClause_) GetSequence() SequenceLike {
	return v.sequence_
}

func (v *withClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
