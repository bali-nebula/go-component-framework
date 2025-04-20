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

func CaretClass() CaretClassLike {
	return caretClass()
}

// Constructor Methods

func (c *caretClass_) Caret() CaretLike {
	var instance = &caret_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *caret_) GetClass() CaretClassLike {
	return caretClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type caret_ struct {
	// Declare the instance attributes.
}

// Class Structure

type caretClass_ struct {
	// Declare the class constants.
}

// Class Reference

func caretClass() *caretClass_ {
	return caretClassReference_
}

var caretClassReference_ = &caretClass_{
	// Initialize the class constants.
}
