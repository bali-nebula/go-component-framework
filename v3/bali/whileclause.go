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

var whileClauseClass = &whileClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func WhileClause() WhileClauseClassLike {
	return whileClauseClass
}

// CLASS METHODS

// Target

type whileClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *whileClauseClass_) MakeWithAttributes(
	condition ConditionLike,
	procedure ProcedureLike,
) WhileClauseLike {
	return &whileClause_{
		condition_: condition,
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type whileClause_ struct {
	condition_ ConditionLike
	procedure_ ProcedureLike
}

// Attributes

func (v *whileClause_) GetCondition() ConditionLike {
	return v.condition_
}

func (v *whileClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
