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

func ExInclusiveSymbolsClass() ExInclusiveSymbolsClassLike {
	return exInclusiveSymbolsClass()
}

// Constructor Methods

func (c *exInclusiveSymbolsClass_) ExInclusiveSymbols(
	symbol1 string,
	symbol2 string,
) ExInclusiveSymbolsLike {
	if uti.IsUndefined(symbol1) {
		panic("The \"symbol1\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol2) {
		panic("The \"symbol2\" attribute is required by this class.")
	}
	var instance = &exInclusiveSymbols_{
		// Initialize the instance attributes.
		symbol1_: symbol1,
		symbol2_: symbol2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveSymbols_) GetClass() ExInclusiveSymbolsClassLike {
	return exInclusiveSymbolsClass()
}

// Attribute Methods

func (v *exInclusiveSymbols_) GetSymbol1() string {
	return v.symbol1_
}

func (v *exInclusiveSymbols_) GetSymbol2() string {
	return v.symbol2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveSymbols_ struct {
	// Declare the instance attributes.
	symbol1_ string
	symbol2_ string
}

// Class Structure

type exInclusiveSymbolsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveSymbolsClass() *exInclusiveSymbolsClass_ {
	return exInclusiveSymbolsClassReference_
}

var exInclusiveSymbolsClassReference_ = &exInclusiveSymbolsClass_{
	// Initialize the class constants.
}
