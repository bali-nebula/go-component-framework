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

func ItemClass() ItemClassLike {
	return itemClass()
}

// Constructor Methods

func (c *itemClass_) Item(
	symbol string,
) ItemLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &item_{
		// Initialize the instance attributes.
		symbol_: symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *item_) GetClass() ItemClassLike {
	return itemClass()
}

// Attribute Methods

func (v *item_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type item_ struct {
	// Declare the instance attributes.
	symbol_ string
}

// Class Structure

type itemClass_ struct {
	// Declare the class constants.
}

// Class Reference

func itemClass() *itemClass_ {
	return itemClassReference_
}

var itemClassReference_ = &itemClass_{
	// Initialize the class constants.
}
