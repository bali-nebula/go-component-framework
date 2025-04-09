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

func DefaultEqualsClass() DefaultEqualsClassLike {
	return defaultEqualsClass()
}

// Constructor Methods

func (c *defaultEqualsClass_) DefaultEquals() DefaultEqualsLike {
	var instance = &defaultEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *defaultEquals_) GetClass() DefaultEqualsClassLike {
	return defaultEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type defaultEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type defaultEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func defaultEqualsClass() *defaultEqualsClass_ {
	return defaultEqualsClassReference_
}

var defaultEqualsClassReference_ = &defaultEqualsClass_{
	// Initialize the class constants.
}
