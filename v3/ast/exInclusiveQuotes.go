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

func ExInclusiveQuotesClass() ExInclusiveQuotesClassLike {
	return exInclusiveQuotesClass()
}

// Constructor Methods

func (c *exInclusiveQuotesClass_) ExInclusiveQuotes(
	quote1 string,
	quote2 string,
) ExInclusiveQuotesLike {
	if uti.IsUndefined(quote1) {
		panic("The \"quote1\" attribute is required by this class.")
	}
	if uti.IsUndefined(quote2) {
		panic("The \"quote2\" attribute is required by this class.")
	}
	var instance = &exInclusiveQuotes_{
		// Initialize the instance attributes.
		quote1_: quote1,
		quote2_: quote2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveQuotes_) GetClass() ExInclusiveQuotesClassLike {
	return exInclusiveQuotesClass()
}

// Attribute Methods

func (v *exInclusiveQuotes_) GetQuote1() string {
	return v.quote1_
}

func (v *exInclusiveQuotes_) GetQuote2() string {
	return v.quote2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveQuotes_ struct {
	// Declare the instance attributes.
	quote1_ string
	quote2_ string
}

// Class Structure

type exInclusiveQuotesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveQuotesClass() *exInclusiveQuotesClass_ {
	return exInclusiveQuotesClassReference_
}

var exInclusiveQuotesClassReference_ = &exInclusiveQuotesClass_{
	// Initialize the class constants.
}
