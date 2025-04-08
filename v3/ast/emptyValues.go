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

func EmptyValuesClass() EmptyValuesClassLike {
	return emptyValuesClass()
}

// Constructor Methods

func (c *emptyValuesClass_) EmptyValues() EmptyValuesLike {
	var instance = &emptyValues_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *emptyValues_) GetClass() EmptyValuesClassLike {
	return emptyValuesClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type emptyValues_ struct {
	// Declare the instance attributes.
}

// Class Structure

type emptyValuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func emptyValuesClass() *emptyValuesClass_ {
	return emptyValuesClassReference_
}

var emptyValuesClassReference_ = &emptyValuesClass_{
	// Initialize the class constants.
}
