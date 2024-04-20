/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import ()

// CLASS ACCESS

// Reference

var integerClass = &integerClass_{
	// This class has no private constants to initialize.
}

// Function

func Integer() IntegerClassLike {
	return integerClass
}

// CLASS METHODS

// Target

type integerClass_ struct {
	minimumValue_ IntegerLike
	maximumValue_ IntegerLike
}

// Constants

func (c *integerClass_) MinimumValue() IntegerLike {
	return c.minimumValue_
}

func (c *integerClass_) MaximumValue() IntegerLike {
	return c.maximumValue_
}

// Constructors

func (c *integerClass_) MakeFromInteger(integer int) IntegerLike {
	return &integer_{}
}

func (c *integerClass_) MakeFromString(string_ string) IntegerLike {
	return &integer_{}
}

// Functions

// INSTANCE METHODS

// Target

type integer_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Discrete

func (v *integer_) AsBoolean() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *integer_) AsInteger() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Lexical

func (v *integer_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Polarized

func (v *integer_) IsNegative() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

// Private
