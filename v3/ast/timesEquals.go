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

func TimesEqualsClass() TimesEqualsClassLike {
	return timesEqualsClass()
}

// Constructor Methods

func (c *timesEqualsClass_) TimesEquals() TimesEqualsLike {
	var instance = &timesEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *timesEquals_) GetClass() TimesEqualsClassLike {
	return timesEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type timesEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type timesEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func timesEqualsClass() *timesEqualsClass_ {
	return timesEqualsClassReference_
}

var timesEqualsClassReference_ = &timesEqualsClass_{
	// Initialize the class constants.
}
