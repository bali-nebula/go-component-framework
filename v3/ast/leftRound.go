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

func LeftRoundClass() LeftRoundClassLike {
	return leftRoundClass()
}

// Constructor Methods

func (c *leftRoundClass_) LeftRound() LeftRoundLike {
	var instance = &leftRound_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *leftRound_) GetClass() LeftRoundClassLike {
	return leftRoundClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type leftRound_ struct {
	// Declare the instance attributes.
}

// Class Structure

type leftRoundClass_ struct {
	// Declare the class constants.
}

// Class Reference

func leftRoundClass() *leftRoundClass_ {
	return leftRoundClassReference_
}

var leftRoundClassReference_ = &leftRoundClass_{
	// Initialize the class constants.
}
