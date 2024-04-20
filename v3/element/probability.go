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

var probabilityClass = &probabilityClass_{
	// This class has no private constants to initialize.
}

// Function

func Probability() ProbabilityClassLike {
	return probabilityClass
}

// CLASS METHODS

// Target

type probabilityClass_ struct {
	minimumValue_ ProbabilityLike
	maximumValue_ ProbabilityLike
}

// Constants

func (c *probabilityClass_) MinimumValue() ProbabilityLike {
	return c.minimumValue_
}

func (c *probabilityClass_) MaximumValue() ProbabilityLike {
	return c.maximumValue_
}

// Constructors

func (c *probabilityClass_) MakeFromFloat(float float64) ProbabilityLike {
	return &probability_{}
}

func (c *probabilityClass_) MakeFromBoolean(boolean bool) ProbabilityLike {
	return &probability_{}
}

func (c *probabilityClass_) MakeFromString(string_ string) ProbabilityLike {
	return &probability_{}
}

// Functions

func (c *probabilityClass_) Random() ProbabilityLike {
	var result_ ProbabilityLike
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type probability_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Continuous

func (v *probability_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *probability_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *probability_) IsInfinite() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *probability_) IsUndefined() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Discrete

func (v *probability_) AsBoolean() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *probability_) AsInteger() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Lexical

func (v *probability_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Public

// Private
