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

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	optionalNotice NoticeLike,
	component ComponentLike,
) DocumentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		optionalNotice_: optionalNotice,
		component_:      component,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

// Attribute Methods

func (v *document_) GetOptionalNotice() NoticeLike {
	return v.optionalNotice_
}

func (v *document_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	optionalNotice_ NoticeLike
	component_      ComponentLike
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentClass() *documentClass_ {
	return documentClassReference_
}

var documentClassReference_ = &documentClass_{
	// Initialize the class constants.
}
