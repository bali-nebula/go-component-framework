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

var ifClauseClass = &ifClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func IfClause() IfClauseClassLike {
	return ifClauseClass
}

// CLASS METHODS

// Target

type ifClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *ifClauseClass_) MakeWithAttributes(
	condition ConditionLike,
	procedure ProcedureLike,
) IfClauseLike {
	return &ifClause_{
		condition_: condition,
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type ifClause_ struct {
	condition_ ConditionLike
	procedure_ ProcedureLike
}

// Attributes

func (v *ifClause_) GetCondition() ConditionLike {
	return v.condition_
}

func (v *ifClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
