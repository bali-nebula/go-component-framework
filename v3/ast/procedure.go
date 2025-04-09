/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import ()

// CLASS INTERFACE

// Access Function

func ProcedureClass() ProcedureClassLike {
	return procedureClass()
}

// Constructor Methods

func (c *procedureClass_) Procedure(
	optionalStatements StatementsLike,
) ProcedureLike {
	var instance = &procedure_{
		// Initialize the instance attributes.
		optionalStatements_: optionalStatements,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *procedure_) GetClass() ProcedureClassLike {
	return procedureClass()
}

// Attribute Methods

func (v *procedure_) GetOptionalStatements() StatementsLike {
	return v.optionalStatements_
}

// PROTECTED INTERFACE

// Instance Structure

type procedure_ struct {
	// Declare the instance attributes.
	optionalStatements_ StatementsLike
}

// Class Structure

type procedureClass_ struct {
	// Declare the class constants.
}

// Class Reference

func procedureClass() *procedureClass_ {
	return procedureClassReference_
}

var procedureClassReference_ = &procedureClass_{
	// Initialize the class constants.
}
