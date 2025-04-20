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

func DashClass() DashClassLike {
	return dashClass()
}

// Constructor Methods

func (c *dashClass_) Dash() DashLike {
	var instance = &dash_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *dash_) GetClass() DashClassLike {
	return dashClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type dash_ struct {
	// Declare the instance attributes.
}

// Class Structure

type dashClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dashClass() *dashClass_ {
	return dashClassReference_
}

var dashClassReference_ = &dashClass_{
	// Initialize the class constants.
}
