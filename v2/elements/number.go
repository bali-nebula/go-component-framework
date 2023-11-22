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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	mat "math"
	cmp "math/cmplx"
	sts "strings"
)

// NUMBER CLASS DEFINITION

// This public singleton creates a unique name space for the number class.
var Number = &numbers_{
	number_(complex(0, 0)),
	number_(complex(1, 0)),
	number_(complex(0, 1)),
	number_(complex(mat.E, 0)),
	number_(complex(mat.Pi, 0)),
	number_(complex(mat.Phi, 0)),
	number_(complex(2.0*mat.Pi, 0)),
	number_(cmp.Inf()),
	number_(cmp.NaN()),
}

// This private type defines the structure associated with the class constants
// and class methods for the number elements.
type numbers_ struct {
	zero      abs.NumberLike
	one       abs.NumberLike
	i         abs.NumberLike
	e         abs.NumberLike
	pi        abs.NumberLike
	phi       abs.NumberLike
	tau       abs.NumberLike
	infinity  abs.NumberLike
	undefined abs.NumberLike
}

// NUMBER CLASS CONSTANTS

// This class constant represents the number zero.
func (t *numbers_) Zero() abs.NumberLike {
	return t.zero
}

// This class constant represents the number one.
func (t *numbers_) One() abs.NumberLike {
	return t.one
}

// This class constant represents the number i.
func (t *numbers_) I() abs.NumberLike {
	return t.i
}

// This class constant represents the number e.
func (t *numbers_) E() abs.NumberLike {
	return t.e
}

// This class constant represents the number pi.
func (t *numbers_) Pi() abs.NumberLike {
	return t.pi
}

// This class constant represents the number phi.
func (t *numbers_) Phi() abs.NumberLike {
	return t.phi
}

// This class constant represents the number tau.
func (t *numbers_) Tau() abs.NumberLike {
	return t.tau
}

// This class constant represents an infinite number.
func (t *numbers_) Infinity() abs.NumberLike {
	return t.infinity
}

// This class constant represents an undefined number.
func (t *numbers_) Undefined() abs.NumberLike {
	return t.undefined
}

// NUMBER CLASS METHODS

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
func (t *numbers_) FromComplex(complex_ complex128) abs.NumberLike {
	var number abs.NumberLike
	switch {
	case cmp.Abs(complex_) == 0:
		// Normalize all versions of zero.
		number = t.zero
	case cmp.IsInf(complex_):
		// Normalize any negative infinities or infinite i's.
		number = t.infinity
	case cmp.IsNaN(complex_):
		// Normalize any NaN's mixed with valid numbers.
		number = t.undefined
	default:
		// Lock onto 0, -1, 1, -i, i, and ∞ if necessary.
		var realPart = lockMagnitude(real(complex_))
		var imaginaryPart = lockMagnitude(imag(complex_))
		number = number_(complex(realPart, imaginaryPart))
	}
	return number
}

// This constructor creates a new number from the specified polar values.
func (t *numbers_) FromPolar(magnitude float64, phase float64) abs.NumberLike {
	var complex_ = cmp.Rect(magnitude, phase)
	var number = t.FromComplex(complex_)
	return number
}

// This constructor creates a new number from the specified string value.
func (t *numbers_) FromString(string_ string) abs.NumberLike {
	var matches = uti.NumberMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a number from an invalid string: %v", string_)
		panic(message)
	}
	var complex_ complex128
	switch {
	case len(matches[2]) > 0:
		// This is a complex number in rectangular form.
		var realPart = floatFromString(matches[1])
		var imaginaryPart = floatFromString(matches[3][:len(matches[3])-1])
		complex_ = complex(realPart, imaginaryPart)
	case len(matches[5]) > 0:
		// This is a complex number in polar form.
		var magnitude = floatFromString(matches[4])
		var phase = floatFromString(matches[7])
		complex_ = cmp.Rect(magnitude, phase)
	default:
		// This is a pure (non-complex) number.
		switch matches[0] {
		case "undefined":
			complex_ = cmp.NaN()
		case "infinity", "∞":
			complex_ = cmp.Inf()
		case "0":
			complex_ = complex(0, 0)
		case "+i", "i":
			complex_ = complex(0, 1)
		case "-i":
			complex_ = complex(0, -1)
		case "+pi", "pi", "-pi", "+phi", "phi", "-phi":
			// We must handle the constants that end in "i" separately.
			complex_ = complex(floatFromString(matches[1]), 0)
		default:
			if sts.HasSuffix(matches[0], "i") {
				// This is a pure imaginary number.
				complex_ = complex(0, floatFromString(matches[0][:len(matches[0])-1]))
			} else {
				// This is a pure real number.
				complex_ = complex(floatFromString(matches[0]), 0)
			}
		}
	}
	var number = t.FromComplex(complex_)
	return number
}

// This constructor returns the minimum value for a number.
func (t *numbers_) MinimumValue() abs.NumberLike {
	return t.zero
}

// This constructor returns the maximum value for a number.
func (t *numbers_) MaximumValue() abs.NumberLike {
	return t.infinity
}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified number.
func (t *numbers_) Inverse(number abs.NumberLike) abs.NumberLike {
	number = t.FromComplex(-number.AsComplex())
	return number
}

// This library function returns the sum of the specified numbers.
func (t *numbers_) Sum(first, second abs.NumberLike) abs.NumberLike {
	var number = t.FromComplex(first.AsComplex() + second.AsComplex())
	return number
}

// This library function returns the difference of the specified numbers.
func (t *numbers_) Difference(first, second abs.NumberLike) abs.NumberLike {
	var number = t.FromComplex(first.AsComplex() - second.AsComplex())
	return number
}

// This library function returns the specified number scaled by the specified
// factor.
func (t *numbers_) Scaled(number abs.NumberLike, factor float64) abs.NumberLike {
	number = t.Product(number, t.FromComplex(complex(factor, 0)))
	return number
}

// NUMERICAL INTERFACE

// This library function returns the reciprocal of the specified number.
func (t *numbers_) Reciprocal(number abs.NumberLike) abs.NumberLike {
	number = t.FromComplex(1.0 / number.AsComplex())
	return number
}

// This library function returns the complex conjugate of the specified number.
func (t *numbers_) Conjugate(number abs.NumberLike) abs.NumberLike {
	number = t.FromComplex(cmp.Conj(number.AsComplex()))
	return number
}

// This library function returns the product of the specified numbers.
func (t *numbers_) Product(first, second abs.NumberLike) abs.NumberLike {
	var number abs.NumberLike
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = t.undefined
	case first.IsInfinite() && !second.IsZero():
		// Infinity times anything other than zero is infinite.
		number = t.infinity
	case second.IsInfinite() && !first.IsZero():
		// Anything other than zero times infinity is infinite.
		number = t.infinity
	default:
		number = t.FromComplex(first.AsComplex() * second.AsComplex())
	}
	return number
}

// This library function returns the quotient of the specified numbers.
func (t *numbers_) Quotient(first, second abs.NumberLike) abs.NumberLike {
	var number abs.NumberLike
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = t.undefined
	case first.IsZero() && second.IsZero():
		// Zero divided by zero is undefined.
		number = t.undefined
	case first.IsInfinite() && second.IsInfinite():
		// Infinity divided by infinity is undefined.
		number = t.undefined
	case first.IsZero() && !second.IsZero():
		// Zero divided by anything other than zero is zero.
		number = t.zero
	case second.IsZero() && !first.IsZero():
		// Anything other than zero divided by zero is infinite.
		number = t.infinity
	case first.IsInfinite() && !second.IsInfinite():
		// Zero divided by anything other than zero is zero.
		number = t.infinity
	case second.IsInfinite() && !first.IsInfinite():
		// Anything other than zero divided by zero is infinite.
		number = t.zero
	default:
		number = t.FromComplex(first.AsComplex() / second.AsComplex())
	}
	return number
}

// This library function returns the remainder of the specified numbers.
func (t *numbers_) Remainder(first, second abs.NumberLike) abs.NumberLike {
	var m1 = cmp.Abs(first.AsComplex())
	var p1 = cmp.Phase(first.AsComplex())
	var m2 = cmp.Abs(second.AsComplex())
	var p2 = cmp.Phase(second.AsComplex())
	var magnitude = lockMagnitude(mat.Remainder(m1, m2))
	var phase = lockPhase(p2 - p1)
	var number = t.FromPolar(magnitude, phase)
	return number
}

// This library function returns the result of raising the specified base to the
// specified exponent.
func (t *numbers_) Power(base, exponent abs.NumberLike) abs.NumberLike {
	var number abs.NumberLike
	switch {
	case base.IsUndefined() || exponent.IsUndefined():
		// Any undefined arguments result in an undefined result.
		number = t.undefined
	case exponent.IsZero():
		// Anything to the zero power is 1 by definition.
		number = t.one
	case base.IsZero():
		// Zero to any power other than zero is still zero.
		number = t.zero
	case base.IsInfinite():
		// Infinity to any power other than zero is infinite.
		number = t.infinity
	case exponent.IsInfinite():
		var magnitude = base.GetMagnitude()
		switch {
		case magnitude < 1:
			// A magnitude less than one to an infinite power is zero.
			number = t.zero
		case magnitude == 1:
			// A magnitude equal to one to an infinite power is one.
			number = t.one
		case magnitude > 1:
			// A magnitude greater than one to an infinite power is infinite.
			number = t.infinity
		default:
			panic(fmt.Sprintf("An impossible magnitude was encountered: %v", magnitude))
		}
	default:
		number = t.FromComplex(cmp.Pow(base.AsComplex(), exponent.AsComplex()))
	}
	return number
}

// This library function returns the result of taking the logarithm using the
// specified base of the specified number.
func (t *numbers_) Logarithm(base, number abs.NumberLike) abs.NumberLike {
	// logB(z) => ln(z) / ln(b)
	var lnB = cmp.Log(base.AsComplex())
	var lnZ = cmp.Log(number.AsComplex())
	var logB = lnZ / lnB
	number = t.FromComplex(logB)
	return number
}

// NUMBER INSTANCE METHODS

// This private type implements the NumberLike interface.  It extends the native
// Go `complex128` type and may represent an integer, floating point, or complex
// number.
type number_ complex128

// COMPLEX INTERFACE

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

// CONTINUOUS INTERFACE

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

// LEXICAL INTERFACE

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

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v number_) IsNegative() bool {
	return real(v) < 0
}

// PRIVATE FUNCTIONS

// This function uses the single precision floating point range to lock a double
// precision magnitude onto 0, 1, -1, or ∞ if the magnitude falls outside the
// single precision range for these values. Otherwise, the magnitude is returned
// unchanged.
func lockMagnitude(magnitude float64) float64 {
	var magnitude32 = float32(magnitude)
	switch {
	case mat.Abs(magnitude) <= 1.2246467991473515e-16:
		magnitude = 0
	case magnitude32 == -1:
		magnitude = -1
	case magnitude32 == 1:
		magnitude = 1
	case mat.IsInf(magnitude, 0):
		magnitude = mat.Inf(1)
	}
	return magnitude
}

// This function uses the single precision floating point range to lock a double
// precision phase onto 0, π/2, π, or 3π/2 if the phase falls outside the single
// precision range for these values. Otherwise, the phase is returned unchanged.
func lockPhase(phase float64) float64 {
	var phase32 = float32(phase)
	switch {
	case mat.Abs(phase) <= 1.2246467991473515e-16:
		phase = 0
	case phase32 == float32(0.5*mat.Pi):
		phase = 0.5 * mat.Pi
	case phase32 == float32(mat.Pi):
		phase = mat.Pi
	case phase32 == float32(1.5*mat.Pi):
		phase = 1.5 * mat.Pi
	}
	return phase
}

// This returns the string for the specified imaginary floating point number.
func stringFromImaginary(imaginary float64) string {
	var string_ string
	switch imaginary {
	case 1:
		string_ = "i"
	case -1:
		string_ = "-i"
	default:
		string_ = stringFromFloat(imaginary) + "i"
	}
	return string_
}
