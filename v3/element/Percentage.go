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

var percentageClass = &percentageClass_{
	// This class has no private constants to initialize.
}

// Function

func Percentage() PercentageClassLike {
	return percentageClass
}

// CLASS METHODS

// Target

type percentageClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *percentageClass_) MakeFromInteger(integer int64) PercentageLike {
	return &percentage_{}
}

func (c *percentageClass_) MakeFromFloat(float float64) PercentageLike {
	return &percentage_{}
}

func (c *percentageClass_) MakeFromString(string_ string) PercentageLike {
	return &percentage_{}
}

// Functions

// INSTANCE METHODS

// Target

type percentage_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Continuous

func (v *percentage_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *percentage_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *percentage_) IsInfinite() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *percentage_) IsUndefined() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Discrete

func (v *percentage_) AsBoolean() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *percentage_) AsInteger() int64 {
	var result_ int64
	// TBA - Implement the method.
	return result_
}

// Lexical

func (v *percentage_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Polarized

func (v *percentage_) IsNegative() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

// Private
