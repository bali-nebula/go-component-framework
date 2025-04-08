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

func BreakClauseClass() BreakClauseClassLike {
	return breakClauseClass()
}

// Constructor Methods

func (c *breakClauseClass_) BreakClause() BreakClauseLike {
	var instance = &breakClause_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *breakClause_) GetClass() BreakClauseClassLike {
	return breakClauseClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type breakClause_ struct {
	// Declare the instance attributes.
}

// Class Structure

type breakClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func breakClauseClass() *breakClauseClass_ {
	return breakClauseClassReference_
}

var breakClauseClassReference_ = &breakClauseClass_{
	// Initialize the class constants.
}
