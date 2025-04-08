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

func ContinueClauseClass() ContinueClauseClassLike {
	return continueClauseClass()
}

// Constructor Methods

func (c *continueClauseClass_) ContinueClause() ContinueClauseLike {
	var instance = &continueClause_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *continueClause_) GetClass() ContinueClauseClassLike {
	return continueClauseClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type continueClause_ struct {
	// Declare the instance attributes.
}

// Class Structure

type continueClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func continueClauseClass() *continueClauseClass_ {
	return continueClauseClassReference_
}

var continueClauseClassReference_ = &continueClauseClass_{
	// Initialize the class constants.
}
