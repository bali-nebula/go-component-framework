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

func SANClass() SANClassLike {
	return sANClass()
}

// Constructor Methods

func (c *sANClass_) SAN() SANLike {
	var instance = &sAN_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *sAN_) GetClass() SANClassLike {
	return sANClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type sAN_ struct {
	// Declare the instance attributes.
}

// Class Structure

type sANClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sANClass() *sANClass_ {
	return sANClassReference_
}

var sANClassReference_ = &sANClass_{
	// Initialize the class constants.
}
