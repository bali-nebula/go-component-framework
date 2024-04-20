/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import ()

// CLASS ACCESS

// Reference

var itemClass = &itemClass_{
	// This class has no private constants to initialize.
}

// Function

func Item() ItemClassLike {
	return itemClass
}

// CLASS METHODS

// Target

type itemClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *itemClass_) MakeWithSymbol(symbol string) ItemLike {
	return &item_{
		symbol_: symbol,
	}
}

// Functions

// INSTANCE METHODS

// Target

type item_ struct {
	symbol_ string
}

// Attributes

func (v *item_) GetSymbol() string {
	return v.symbol_
}

// Public

// Private
