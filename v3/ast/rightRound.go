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

func RightRoundClass() RightRoundClassLike {
	return rightRoundClass()
}

// Constructor Methods

func (c *rightRoundClass_) RightRound() RightRoundLike {
	var instance = &rightRound_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *rightRound_) GetClass() RightRoundClassLike {
	return rightRoundClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type rightRound_ struct {
	// Declare the instance attributes.
}

// Class Structure

type rightRoundClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rightRoundClass() *rightRoundClass_ {
	return rightRoundClassReference_
}

var rightRoundClassReference_ = &rightRoundClass_{
	// Initialize the class constants.
}
