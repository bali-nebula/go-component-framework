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

func AttributeClass() AttributeClassLike {
	return attributeClass()
}

// Constructor Methods

func (c *attributeClass_) Attribute(
	identifier string,
	optionalSubcomponent SubcomponentLike,
) AttributeLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &attribute_{
		// Initialize the instance attributes.
		identifier_:           identifier,
		optionalSubcomponent_: optionalSubcomponent,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *attribute_) GetClass() AttributeClassLike {
	return attributeClass()
}

// Attribute Methods

func (v *attribute_) GetIdentifier() string {
	return v.identifier_
}

func (v *attribute_) GetOptionalSubcomponent() SubcomponentLike {
	return v.optionalSubcomponent_
}

// PROTECTED INTERFACE

// Instance Structure

type attribute_ struct {
	// Declare the instance attributes.
	identifier_           string
	optionalSubcomponent_ SubcomponentLike
}

// Class Structure

type attributeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributeClass() *attributeClass_ {
	return attributeClassReference_
}

var attributeClassReference_ = &attributeClass_{
	// Initialize the class constants.
}
