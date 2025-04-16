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

func NoValuesClass() NoValuesClassLike {
	return noValuesClass()
}

// Constructor Methods

func (c *noValuesClass_) NoValues() NoValuesLike {
	var instance = &noValues_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noValues_) GetClass() NoValuesClassLike {
	return noValuesClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type noValues_ struct {
	// Declare the instance attributes.
}

// Class Structure

type noValuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noValuesClass() *noValuesClass_ {
	return noValuesClassReference_
}

var noValuesClassReference_ = &noValuesClass_{
	// Initialize the class constants.
}
