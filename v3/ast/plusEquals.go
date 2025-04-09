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

func PlusEqualsClass() PlusEqualsClassLike {
	return plusEqualsClass()
}

// Constructor Methods

func (c *plusEqualsClass_) PlusEquals() PlusEqualsLike {
	var instance = &plusEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *plusEquals_) GetClass() PlusEqualsClassLike {
	return plusEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type plusEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type plusEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func plusEqualsClass() *plusEqualsClass_ {
	return plusEqualsClassReference_
}

var plusEqualsClassReference_ = &plusEqualsClass_{
	// Initialize the class constants.
}
