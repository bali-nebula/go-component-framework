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

func MatchesClass() MatchesClassLike {
	return matchesClass()
}

// Constructor Methods

func (c *matchesClass_) Matches() MatchesLike {
	var instance = &matches_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *matches_) GetClass() MatchesClassLike {
	return matchesClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type matches_ struct {
	// Declare the instance attributes.
}

// Class Structure

type matchesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func matchesClass() *matchesClass_ {
	return matchesClassReference_
}

var matchesClassReference_ = &matchesClass_{
	// Initialize the class constants.
}
