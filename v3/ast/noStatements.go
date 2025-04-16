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

func NoStatementsClass() NoStatementsClassLike {
	return noStatementsClass()
}

// Constructor Methods

func (c *noStatementsClass_) NoStatements() NoStatementsLike {
	var instance = &noStatements_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noStatements_) GetClass() NoStatementsClassLike {
	return noStatementsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type noStatements_ struct {
	// Declare the instance attributes.
}

// Class Structure

type noStatementsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noStatementsClass() *noStatementsClass_ {
	return noStatementsClassReference_
}

var noStatementsClassReference_ = &noStatementsClass_{
	// Initialize the class constants.
}
