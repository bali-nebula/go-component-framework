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

func ColonEqualsClass() ColonEqualsClassLike {
	return colonEqualsClass()
}

// Constructor Methods

func (c *colonEqualsClass_) ColonEquals() ColonEqualsLike {
	var instance = &colonEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *colonEquals_) GetClass() ColonEqualsClassLike {
	return colonEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type colonEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type colonEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func colonEqualsClass() *colonEqualsClass_ {
	return colonEqualsClassReference_
}

var colonEqualsClassReference_ = &colonEqualsClass_{
	// Initialize the class constants.
}
