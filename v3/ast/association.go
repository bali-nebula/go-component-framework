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

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AssociationClass() AssociationClassLike {
	return associationClass()
}

// Constructor Methods

func (c *associationClass_) Association(
	symbol string,
	colon string,
	component ComponentLike,
) AssociationLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(colon) {
		panic("The \"colon\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &association_{
		// Initialize the instance attributes.
		symbol_:    symbol,
		colon_:     colon,
		component_: component,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *association_) GetClass() AssociationClassLike {
	return associationClass()
}

// Attribute Methods

func (v *association_) GetSymbol() string {
	return v.symbol_
}

func (v *association_) GetColon() string {
	return v.colon_
}

func (v *association_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Instance Structure

type association_ struct {
	// Declare the instance attributes.
	symbol_    string
	colon_     string
	component_ ComponentLike
}

// Class Structure

type associationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func associationClass() *associationClass_ {
	return associationClassReference_
}

var associationClassReference_ = &associationClass_{
	// Initialize the class constants.
}
