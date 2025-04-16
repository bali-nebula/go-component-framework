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

func InclusiveSymbolsClass() InclusiveSymbolsClassLike {
	return inclusiveSymbolsClass()
}

// Constructor Methods

func (c *inclusiveSymbolsClass_) InclusiveSymbols(
	symbol1 string,
	symbol2 string,
) InclusiveSymbolsLike {
	if uti.IsUndefined(symbol1) {
		panic("The \"symbol1\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol2) {
		panic("The \"symbol2\" attribute is required by this class.")
	}
	var instance = &inclusiveSymbols_{
		// Initialize the instance attributes.
		symbol1_: symbol1,
		symbol2_: symbol2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveSymbols_) GetClass() InclusiveSymbolsClassLike {
	return inclusiveSymbolsClass()
}

// Attribute Methods

func (v *inclusiveSymbols_) GetSymbol1() string {
	return v.symbol1_
}

func (v *inclusiveSymbols_) GetSymbol2() string {
	return v.symbol2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveSymbols_ struct {
	// Declare the instance attributes.
	symbol1_ string
	symbol2_ string
}

// Class Structure

type inclusiveSymbolsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveSymbolsClass() *inclusiveSymbolsClass_ {
	return inclusiveSymbolsClassReference_
}

var inclusiveSymbolsClassReference_ = &inclusiveSymbolsClass_{
	// Initialize the class constants.
}
