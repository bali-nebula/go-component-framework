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

var selectClauseClass = &selectClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func SelectClause() SelectClauseClassLike {
	return selectClauseClass
}

// CLASS METHODS

// Target

type selectClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *selectClauseClass_) MakeWithAttributes(
	target TargetLike,
	template TemplateLike,
	procedure ProcedureLike,
) SelectClauseLike {
	return &selectClause_{
		target_:    target,
		template_:  template,
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type selectClause_ struct {
	target_    TargetLike
	template_  TemplateLike
	procedure_ ProcedureLike
}

// Attributes

func (v *selectClause_) GetTarget() TargetLike {
	return v.target_
}

func (v *selectClause_) GetTemplate() TemplateLike {
	return v.template_
}

func (v *selectClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
