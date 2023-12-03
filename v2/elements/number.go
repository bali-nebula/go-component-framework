/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	mat "math"
	cmp "math/cmplx"
)

// CLASS DEFINITIONS

// This private type implements the NumberLike interface.  It extends the native
// Go `complex128` type and may represent an integer, floating point, or complex
// number.
type number_ complex128

// This private type defines the structure associated with the class constants
// and class functions for the number elements.
type numberClass_ struct {
	zero      NumberLike
	one       NumberLike
	i         NumberLike
	e         NumberLike
	pi        NumberLike
	phi       NumberLike
	tau       NumberLike
	infinity  NumberLike
	undefined NumberLike
}

// CLASS CONSTANTS

// This class constant represents the minimum value for a number.
func (c *numberClass_) MinimumValue() NumberLike {
	return c.zero
}

// This class constant represents the maximum value for a number.
func (c *numberClass_) MaximumValue() NumberLike {
	return c.infinity
}

// This class constant represents the number zero.
func (c *numberClass_) Zero() NumberLike {
	return c.zero
}

// This class constant represents the number one.
func (c *numberClass_) One() NumberLike {
	return c.one
}

// This class constant represents the number i.
func (c *numberClass_) I() NumberLike {
	return c.i
}

// This class constant represents the number e.
func (c *numberClass_) E() NumberLike {
	return c.e
}

// This class constant represents the number pi.
func (c *numberClass_) Pi() NumberLike {
	return c.pi
}

// This class constant represents the number phi.
func (c *numberClass_) Phi() NumberLike {
	return c.phi
}

// This class constant represents the number tau.
func (c *numberClass_) Tau() NumberLike {
	return c.tau
}

// This class constant represents an infinite number.
func (c *numberClass_) Infinity() NumberLike {
	return c.infinity
}

// This class constant represents an undefined number.
func (c *numberClass_) Undefined() NumberLike {
	return c.undefined
}

// CLASS CONSTRUCTORS

// This constructor creates a new number element from the specified complex
// number.  Each number element is mapped to the Riemann Sphere.
//   - https://en.wikipedia.org/wiki/Riemann_sphere
//
// This mapping removes many of the idiosyncrasies associated with the normal
// complex plane. There is only one value for zero and one value for infinity.
// This simplifies a lot of the mathematical operations:
//
//	z + zero => z
//	z + infinity => infinity
//
//	z - infinity => infinity     {z != infinity}
//	infinity - z => infinity     {z != infinity}
//
//	z * zero => zero             {z != infinity}
//	z * infinity => infinity     {z != zero}
//
//	z / zero => infinity         {z != zero}
//	zero / z => zero             {z != zero}
//
//	z / infinity => zero         {z != infinity}
//	infinity / z => infinity     {z != infinity}
//
//	z ^ zero => one              {by definition}
//	zero ^ z => zero             {z != zero}
//
//	z ^ infinity => zero         {|z| < one}
//	z ^ infinity => one          {|z| = one}
//	z ^ infinity => infinity     {|z| > one}
//	infinity ^ z => infinity     {z != zero}
//
//	log(z, zero) => infinity     {zero < z < infinity}
//	log(zero, z) => zero         {zero < z < infinity}
//
//	log(z, one) => zero          {zero < z}
//	log(one, z) => zero          {zero < z}
//
//	log(z, infinity) => infinity {zero < z < infinity}
//	log(infinity, z) => zero     {zero < z < infinity}
//
// This leaves only the following operations undefined:
//
//	infinity - infinity => undefined
//	zero * infinity => undefined
//	zero / zero => undefined
//	infinity / infinity => undefined
//	log(zero, zero) => undefined
//	log(zero, infinity) => undefined
//	log(infinity, zero) => undefined
//	log(infinity, infinity) => undefined
//
// The resulting number system is easier to use for most applications. For
// numerical analysis the ANSI plus and minus zero values are often used as a
// crutch that leads to misleading convergence information for oscillating
// functions. In the case of numerical analysis it is probably better to track
// the course of the function as it converges than to look at the final value.
func (c *numberClass_) FromComplex(complex_ complex128) NumberLike {
	var number NumberLike
	switch {
	case cmp.Abs(complex_) == 0:
		// Normalize all versions of zero.
		number = c.zero
	case cmp.IsInf(complex_):
		// Normalize any negative infinities or infinite i's.
		number = c.infinity
	case cmp.IsNaN(complex_):
		// Normalize any NaN's mixed with valid numbers.
		number = c.undefined
	default:
		// Lock onto 0, -1, 1, -i, i, and ∞ if necessary.
		var realPart = lockMagnitude(real(complex_))
		var imaginaryPart = lockMagnitude(imag(complex_))
		number = number_(complex(realPart, imaginaryPart))
	}
	return number
}

// This constructor creates a new number from the specified polar values.
func (c *numberClass_) FromPolar(magnitude float64, phase float64) NumberLike {
	var complex_ = cmp.Rect(magnitude, phase)
	var number = c.FromComplex(complex_)
	return number
}

// This constructor creates a new number from the specified string value.
func (c *numberClass_) FromString(string_ string) NumberLike {
	var matches = uti.NumberMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a number from an invalid string: %v", string_)
		panic(message)
	}
	var complex_ = complexFromMatches(matches)
	var number = c.FromComplex(complex_)
	return number
}

// CLASS METHODS

// Complex Interface

// This method returns a native complex value for this continuous component.
func (v number_) AsComplex() complex128 {
	return complex128(v)
}

// This method returns the real part of this complex component.
func (v number_) GetReal() float64 {
	return real(v)
}

// This method returns the imaginary part of this complex component.
func (v number_) GetImaginary() float64 {
	return imag(v)
}

// This method returns the magnitude of this complex component.
func (v number_) GetMagnitude() float64 {
	return lockMagnitude(cmp.Abs(complex128(v)))
}

// This method returns the phase angle of this complex component.
func (v number_) GetPhase() float64 {
	return cmp.Phase(complex128(v))
}

// Continuous Interface

// This method returns a real value for this continuous component.
func (v number_) AsFloat() float64 {
	return real(v)
}

// This method determines whether or not this number is zero.
func (v number_) IsZero() bool {
	return real(v) == 0 && imag(v) == 0
}

// This method determines whether or not this number is infinite.
func (v number_) IsInfinite() bool {
	return mat.IsInf(real(v), 0) || mat.IsInf(imag(v), 0)
}

// This method determines whether or not this number is undefined.
func (v number_) IsUndefined() bool {
	return mat.IsNaN(real(v)) || mat.IsNaN(imag(v))
}

// Lexical Interface

// This method returns a string value for this lexical element.
func (v number_) AsString() string {
	var string_ string
	switch {
	case v.IsZero():
		string_ = "0"
	case v.IsInfinite():
		string_ = "∞"
	case v.IsUndefined():
		string_ = "undefined"
	default:
		var realPart = v.GetReal()
		var imagPart = v.GetImaginary()
		switch {
		case imagPart == 0:
			string_ = stringFromFloat(realPart)
		case realPart == 0:
			string_ = stringFromImaginary(imagPart)
		default:
			string_ += "("
			string_ += stringFromFloat(realPart)
			string_ += ", "
			string_ += stringFromImaginary(imagPart)
			string_ += ")"
		}
	}
	return string_
}

// Polarized Interface

// This method determines whether or not this polarized component is negative.
func (v number_) IsNegative() bool {
	return real(v) < 0
}

// CLASS FUNCTIONS

// This library function returns the inverse of the specified number.
func (c *numberClass_) Inverse(number NumberLike) NumberLike {
	number = c.FromComplex(-number.AsComplex())
	return number
}

// This library function returns the sum of the specified numbers.
func (c *numberClass_) Sum(first, second NumberLike) NumberLike {
	var number = c.FromComplex(first.AsComplex() + second.AsComplex())
	return number
}

// This library function returns the difference of the specified numbers.
func (c *numberClass_) Difference(first, second NumberLike) NumberLike {
	var number = c.FromComplex(first.AsComplex() - second.AsComplex())
	return number
}

// This library function returns the specified number scaled by the specified
// factor.
func (c *numberClass_) Scaled(number NumberLike, factor float64) NumberLike {
	number = c.Product(number, c.FromComplex(complex(factor, 0)))
	return number
}

// This library function returns the reciprocal of the specified number.
func (c *numberClass_) Reciprocal(number NumberLike) NumberLike {
	number = c.FromComplex(1.0 / number.AsComplex())
	return number
}

// This library function returns the complex conjugate of the specified number.
func (c *numberClass_) Conjugate(number NumberLike) NumberLike {
	number = c.FromComplex(cmp.Conj(number.AsComplex()))
	return number
}

// This library function returns the product of the specified numbers.
func (c *numberClass_) Product(first, second NumberLike) NumberLike {
	var number NumberLike
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined
	case first.IsInfinite() && !second.IsZero():
		// Infinity times anything other than zero is infinite.
		number = c.infinity
	case second.IsInfinite() && !first.IsZero():
		// Anything other than zero times infinity is infinite.
		number = c.infinity
	default:
		number = c.FromComplex(first.AsComplex() * second.AsComplex())
	}
	return number
}

// This library function returns the quotient of the specified numbers.
func (c *numberClass_) Quotient(first, second NumberLike) NumberLike {
	var number NumberLike
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined
	case first.IsZero() && second.IsZero():
		// Zero divided by zero is undefined.
		number = c.undefined
	case first.IsInfinite() && second.IsInfinite():
		// Infinity divided by infinity is undefined.
		number = c.undefined
	case first.IsZero() && !second.IsZero():
		// Zero divided by anything other than zero is zero.
		number = c.zero
	case second.IsZero() && !first.IsZero():
		// Anything other than zero divided by zero is infinite.
		number = c.infinity
	case first.IsInfinite() && !second.IsInfinite():
		// Zero divided by anything other than zero is zero.
		number = c.infinity
	case second.IsInfinite() && !first.IsInfinite():
		// Anything other than zero divided by zero is infinite.
		number = c.zero
	default:
		number = c.FromComplex(first.AsComplex() / second.AsComplex())
	}
	return number
}

// This library function returns the remainder of the specified numbers.
func (c *numberClass_) Remainder(first, second NumberLike) NumberLike {
	var m1 = cmp.Abs(first.AsComplex())
	var p1 = cmp.Phase(first.AsComplex())
	var m2 = cmp.Abs(second.AsComplex())
	var p2 = cmp.Phase(second.AsComplex())
	var magnitude = lockMagnitude(mat.Remainder(m1, m2))
	var phase = lockPhase(p2 - p1)
	var number = c.FromPolar(magnitude, phase)
	return number
}

// This library function returns the result of raising the specified base to the
// specified exponent.
func (c *numberClass_) Power(base, exponent NumberLike) NumberLike {
	var number NumberLike
	switch {
	case base.IsUndefined() || exponent.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = c.undefined
	case exponent.IsZero():
		// Anything to the zero power is 1 by definition.
		number = c.one
	case base.IsZero():
		// Zero to any power other than zero is still zero.
		number = c.zero
	case base.IsInfinite():
		// Infinity to any power other than zero is infinite.
		number = c.infinity
	case exponent.IsInfinite():
		var magnitude = base.GetMagnitude()
		switch {
		case magnitude < 1:
			// A magnitude less than one to an infinite power is zero.
			number = c.zero
		case magnitude == 1:
			// A magnitude equal to one to an infinite power is one.
			number = c.one
		case magnitude > 1:
			// A magnitude greater than one to an infinite power is infinite.
			number = c.infinity
		default:
			panic(fmt.Sprintf("An impossible magnitude was encountered: %v", magnitude))
		}
	default:
		number = c.FromComplex(cmp.Pow(base.AsComplex(), exponent.AsComplex()))
	}
	return number
}

// This library function returns the result of taking the logarithm using the
// specified base of the specified number.
func (c *numberClass_) Logarithm(base, number NumberLike) NumberLike {
	// logB(z) => ln(z) / ln(b)
	var lnB = cmp.Log(base.AsComplex())
	var lnZ = cmp.Log(number.AsComplex())
	var logB = lnZ / lnB
	number = c.FromComplex(logB)
	return number
}
