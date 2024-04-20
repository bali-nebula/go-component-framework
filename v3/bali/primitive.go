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

var primitiveClass = &primitiveClass_{
	// This class has no private constants to initialize.
}

// Function

func Primitive() PrimitiveClassLike {
	return primitiveClass
}

// CLASS METHODS

// Target

type primitiveClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *primitiveClass_) MakeWithElement(element ElementLike) PrimitiveLike {
	return &primitive_{
		element_: element,
	}
}

func (c *primitiveClass_) MakeWithString(string_ StringLike) PrimitiveLike {
	return &primitive_{
		string_: string_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type primitive_ struct {
	element_ ElementLike
	string_  StringLike
}

// Attributes

func (v *primitive_) GetElement() ElementLike {
	return v.element_
}

func (v *primitive_) GetString() StringLike {
	return v.string_
}

// Public

// Private
