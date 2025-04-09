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

func DivideEqualsClass() DivideEqualsClassLike {
	return divideEqualsClass()
}

// Constructor Methods

func (c *divideEqualsClass_) DivideEquals() DivideEqualsLike {
	var instance = &divideEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *divideEquals_) GetClass() DivideEqualsClassLike {
	return divideEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type divideEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type divideEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func divideEqualsClass() *divideEqualsClass_ {
	return divideEqualsClassReference_
}

var divideEqualsClassReference_ = &divideEqualsClass_{
	// Initialize the class constants.
}
