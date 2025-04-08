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

func RecipientClass() RecipientClassLike {
	return recipientClass()
}

// Constructor Methods

func (c *recipientClass_) Recipient(
	any_ any,
) RecipientLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &recipient_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *recipient_) GetClass() RecipientClassLike {
	return recipientClass()
}

// Attribute Methods

func (v *recipient_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type recipient_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type recipientClass_ struct {
	// Declare the class constants.
}

// Class Reference

func recipientClass() *recipientClass_ {
	return recipientClassReference_
}

var recipientClassReference_ = &recipientClass_{
	// Initialize the class constants.
}
