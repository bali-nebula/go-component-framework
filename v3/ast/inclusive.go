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

func InclusiveClass() InclusiveClassLike {
	return inclusiveClass()
}

// Constructor Methods

func (c *inclusiveClass_) Inclusive() InclusiveLike {
	var instance = &inclusive_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusive_) GetClass() InclusiveClassLike {
	return inclusiveClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type inclusive_ struct {
	// Declare the instance attributes.
}

// Class Structure

type inclusiveClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveClass() *inclusiveClass_ {
	return inclusiveClassReference_
}

var inclusiveClassReference_ = &inclusiveClass_{
	// Initialize the class constants.
}
