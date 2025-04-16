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

func AndClass() AndClassLike {
	return andClass()
}

// Constructor Methods

func (c *andClass_) And() AndLike {
	var instance = &and_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *and_) GetClass() AndClassLike {
	return andClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type and_ struct {
	// Declare the instance attributes.
}

// Class Structure

type andClass_ struct {
	// Declare the class constants.
}

// Class Reference

func andClass() *andClass_ {
	return andClassReference_
}

var andClassReference_ = &andClass_{
	// Initialize the class constants.
}
