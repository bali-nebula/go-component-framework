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

var onClauseClass = &onClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func OnClause() OnClauseClassLike {
	return onClauseClass
}

// CLASS METHODS

// Target

type onClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *onClauseClass_) MakeWithAttributes(
	failure FailureLike,
	template TemplateLike,
	procedure ProcedureLike,
) OnClauseLike {
	return &onClause_{
		failure_:   failure,
		template_:  template,
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type onClause_ struct {
	failure_   FailureLike
	template_  TemplateLike
	procedure_ ProcedureLike
}

// Attributes

func (v *onClause_) GetFailure() FailureLike {
	return v.failure_
}

func (v *onClause_) GetTemplate() TemplateLike {
	return v.template_
}

func (v *onClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
