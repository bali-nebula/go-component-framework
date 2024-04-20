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

var procedureClass = &procedureClass_{
	// This class has no private constants to initialize.
}

// Function

func Procedure() ProcedureClassLike {
	return procedureClass
}

// CLASS METHODS

// Target

type procedureClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *procedureClass_) MakeWithLines(lines LinesLike) ProcedureLike {
	return &procedure_{
		lines_: lines,
	}
}

// Functions

// INSTANCE METHODS

// Target

type procedure_ struct {
	lines_ LinesLike
}

// Attributes

func (v *procedure_) GetLines() LinesLike {
	return v.lines_
}

// Public

// Private
