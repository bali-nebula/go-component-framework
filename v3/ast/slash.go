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

func SlashClass() SlashClassLike {
	return slashClass()
}

// Constructor Methods

func (c *slashClass_) Slash() SlashLike {
	var instance = &slash_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *slash_) GetClass() SlashClassLike {
	return slashClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type slash_ struct {
	// Declare the instance attributes.
}

// Class Structure

type slashClass_ struct {
	// Declare the class constants.
}

// Class Reference

func slashClass() *slashClass_ {
	return slashClassReference_
}

var slashClassReference_ = &slashClass_{
	// Initialize the class constants.
}
