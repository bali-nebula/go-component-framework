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

func IORClass() IORClassLike {
	return iORClass()
}

// Constructor Methods

func (c *iORClass_) IOR() IORLike {
	var instance = &iOR_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *iOR_) GetClass() IORClassLike {
	return iORClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type iOR_ struct {
	// Declare the instance attributes.
}

// Class Structure

type iORClass_ struct {
	// Declare the class constants.
}

// Class Reference

func iORClass() *iORClass_ {
	return iORClassReference_
}

var iORClassReference_ = &iORClass_{
	// Initialize the class constants.
}
