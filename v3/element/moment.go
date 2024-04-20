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

var momentClass = &momentClass_{
	// This class has no private constants to initialize.
}

// Function

func Moment() MomentClassLike {
	return momentClass
}

// CLASS METHODS

// Target

type momentClass_ struct {
	minimumValue_ MomentLike
	maximumValue_ MomentLike
	epoch_        MomentLike
}

// Constants

func (c *momentClass_) MinimumValue() MomentLike {
	return c.minimumValue_
}

func (c *momentClass_) MaximumValue() MomentLike {
	return c.maximumValue_
}

func (c *momentClass_) Epoch() MomentLike {
	return c.epoch_
}

// Constructors

func (c *momentClass_) Make() MomentLike {
	return &moment_{}
}

func (c *momentClass_) MakeFromMilliseconds(milliseconds int) MomentLike {
	return &moment_{}
}

func (c *momentClass_) MakeFromString(string_ string) MomentLike {
	return &moment_{}
}

// Functions

func (c *momentClass_) Duration(
	first MomentLike,
	second MomentLike,
) DurationLike {
	var result_ DurationLike
	// TBA - Implement the function.
	return result_
}

func (c *momentClass_) Earlier(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	var result_ MomentLike
	// TBA - Implement the function.
	return result_
}

func (c *momentClass_) Later(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	var result_ MomentLike
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type moment_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Discrete

func (v *moment_) AsBoolean() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsInteger() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Lexical

func (v *moment_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Temporal

func (v *moment_) AsMilliseconds() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsSeconds() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsMinutes() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsHours() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsDays() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsWeeks() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsMonths() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *moment_) AsYears() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Public

// Private
