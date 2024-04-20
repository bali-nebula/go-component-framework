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

var recipientClass = &recipientClass_{
	// This class has no private constants to initialize.
}

// Function

func Recipient() RecipientClassLike {
	return recipientClass
}

// CLASS METHODS

// Target

type recipientClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *recipientClass_) MakeWithSymbol(symbol string) RecipientLike {
	return &recipient_{
		symbol_: symbol,
	}
}

func (c *recipientClass_) MakeWithAttribute(attribute AttributeLike) RecipientLike {
	return &recipient_{
		attribute_: attribute,
	}
}

// Functions

// INSTANCE METHODS

// Target

type recipient_ struct {
	symbol_    string
	attribute_ AttributeLike
}

// Attributes

func (v *recipient_) GetSymbol() string {
	return v.symbol_
}

func (v *recipient_) GetAttribute() AttributeLike {
	return v.attribute_
}

// Public

// Private
