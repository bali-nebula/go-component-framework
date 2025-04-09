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

func ModuloClass() ModuloClassLike {
	return moduloClass()
}

// Constructor Methods

func (c *moduloClass_) Modulo() ModuloLike {
	var instance = &modulo_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *modulo_) GetClass() ModuloClassLike {
	return moduloClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type modulo_ struct {
	// Declare the instance attributes.
}

// Class Structure

type moduloClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moduloClass() *moduloClass_ {
	return moduloClassReference_
}

var moduloClassReference_ = &moduloClass_{
	// Initialize the class constants.
}
