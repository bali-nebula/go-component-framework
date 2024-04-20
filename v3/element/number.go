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

var numberClass = &numberClass_{
	// This class has no private constants to initialize.
}

// Function

func Number() NumberClassLike {
	return numberClass
}

// CLASS METHODS

// Target

type numberClass_ struct {
	minimumValue_ NumberLike
	maximumValue_ NumberLike
	zero_         NumberLike
	one_          NumberLike
	i_            NumberLike
	e_            NumberLike
	pi_           NumberLike
	phi_          NumberLike
	tau_          NumberLike
	infinity_     NumberLike
	undefined_    NumberLike
}

// Constants

func (c *numberClass_) MinimumValue() NumberLike {
	return c.minimumValue_
}

func (c *numberClass_) MaximumValue() NumberLike {
	return c.maximumValue_
}

func (c *numberClass_) Zero() NumberLike {
	return c.zero_
}

func (c *numberClass_) One() NumberLike {
	return c.one_
}

func (c *numberClass_) I() NumberLike {
	return c.i_
}

func (c *numberClass_) E() NumberLike {
	return c.e_
}

func (c *numberClass_) Pi() NumberLike {
	return c.pi_
}

func (c *numberClass_) Phi() NumberLike {
	return c.phi_
}

func (c *numberClass_) Tau() NumberLike {
	return c.tau_
}

func (c *numberClass_) Infinity() NumberLike {
	return c.infinity_
}

func (c *numberClass_) Undefined() NumberLike {
	return c.undefined_
}

// Constructors

func (c *numberClass_) MakeFromComplex(complex_ complex128) NumberLike {
	return &number_{}
}

func (c *numberClass_) MakeFromPolar(
	magnitude float64,
	phase float64,
) NumberLike {
	return &number_{}
}

func (c *numberClass_) MakeFromString(string_ string) NumberLike {
	return &number_{}
}

// Functions

func (c *numberClass_) Inverse(number NumberLike) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Sum(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Difference(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Scaled(
	number NumberLike,
	factor float64,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Reciprocal(number NumberLike) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Conjugate(number NumberLike) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Product(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Quotient(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Remainder(
	first NumberLike,
	second NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Power(
	base NumberLike,
	exponent NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

func (c *numberClass_) Logarithm(
	base NumberLike,
	number NumberLike,
) NumberLike {
	var result_ NumberLike
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type number_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Complex

func (v *number_) AsComplex() complex128 {
	var result_ complex128
	// TBA - Implement the method.
	return result_
}

func (v *number_) GetReal() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *number_) GetImaginary() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *number_) GetMagnitude() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *number_) GetPhase() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Continuous

func (v *number_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v *number_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *number_) IsInfinite() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *number_) IsUndefined() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Lexical

func (v *number_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Polarized

func (v *number_) IsNegative() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// PACKAGE FUNCTIONS

// Private
