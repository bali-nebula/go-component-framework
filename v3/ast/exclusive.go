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

func ExclusiveClass() ExclusiveClassLike {
	return exclusiveClass()
}

// Constructor Methods

func (c *exclusiveClass_) Exclusive() ExclusiveLike {
	var instance = &exclusive_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusive_) GetClass() ExclusiveClassLike {
	return exclusiveClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type exclusive_ struct {
	// Declare the instance attributes.
}

// Class Structure

type exclusiveClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveClass() *exclusiveClass_ {
	return exclusiveClassReference_
}

var exclusiveClassReference_ = &exclusiveClass_{
	// Initialize the class constants.
}
