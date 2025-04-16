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

func ExclusiveQuotesClass() ExclusiveQuotesClassLike {
	return exclusiveQuotesClass()
}

// Constructor Methods

func (c *exclusiveQuotesClass_) ExclusiveQuotes(
	quote1 string,
	quote2 string,
) ExclusiveQuotesLike {
	if uti.IsUndefined(quote1) {
		panic("The \"quote1\" attribute is required by this class.")
	}
	if uti.IsUndefined(quote2) {
		panic("The \"quote2\" attribute is required by this class.")
	}
	var instance = &exclusiveQuotes_{
		// Initialize the instance attributes.
		quote1_: quote1,
		quote2_: quote2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveQuotes_) GetClass() ExclusiveQuotesClassLike {
	return exclusiveQuotesClass()
}

// Attribute Methods

func (v *exclusiveQuotes_) GetQuote1() string {
	return v.quote1_
}

func (v *exclusiveQuotes_) GetQuote2() string {
	return v.quote2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveQuotes_ struct {
	// Declare the instance attributes.
	quote1_ string
	quote2_ string
}

// Class Structure

type exclusiveQuotesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveQuotesClass() *exclusiveQuotesClass_ {
	return exclusiveQuotesClassReference_
}

var exclusiveQuotesClassReference_ = &exclusiveQuotesClass_{
	// Initialize the class constants.
}
