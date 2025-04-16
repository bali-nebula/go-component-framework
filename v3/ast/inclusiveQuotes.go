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

func InclusiveQuotesClass() InclusiveQuotesClassLike {
	return inclusiveQuotesClass()
}

// Constructor Methods

func (c *inclusiveQuotesClass_) InclusiveQuotes(
	quote1 string,
	quote2 string,
) InclusiveQuotesLike {
	if uti.IsUndefined(quote1) {
		panic("The \"quote1\" attribute is required by this class.")
	}
	if uti.IsUndefined(quote2) {
		panic("The \"quote2\" attribute is required by this class.")
	}
	var instance = &inclusiveQuotes_{
		// Initialize the instance attributes.
		quote1_: quote1,
		quote2_: quote2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveQuotes_) GetClass() InclusiveQuotesClassLike {
	return inclusiveQuotesClass()
}

// Attribute Methods

func (v *inclusiveQuotes_) GetQuote1() string {
	return v.quote1_
}

func (v *inclusiveQuotes_) GetQuote2() string {
	return v.quote2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveQuotes_ struct {
	// Declare the instance attributes.
	quote1_ string
	quote2_ string
}

// Class Structure

type inclusiveQuotesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveQuotesClass() *inclusiveQuotesClass_ {
	return inclusiveQuotesClassReference_
}

var inclusiveQuotesClassReference_ = &inclusiveQuotesClass_{
	// Initialize the class constants.
}
