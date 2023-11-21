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

// NUMBER CONSTANTS

var zero = complex(0, 0)
var infinity = cmp.Inf()
var undefined = cmp.NaN()

var I abs.NumberLike = number_(complex(0.0, 1.0))
var One abs.NumberLike = number_(complex(1.0, 0.0))
var Phi abs.NumberLike = number_(complex(mat.Phi, 0.0))
var Zero abs.NumberLike = number_(zero)
var Infinity abs.NumberLike = number_(infinity)
var Undefined abs.NumberLike = number_(undefined)

// NUMBER CONSTRUCTORS

// This constructor creates a new number that is mapped to the Riemann Sphere.
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
func NumberFromComplex(complex_ complex128) abs.NumberLike {
	switch {
	case cmp.Abs(complex_) == 0:
		// Normalize all versions of zero.
		return Zero
	case cmp.IsInf(complex_):
		// Normalize any negative infinities or infinite i's.
		return Infinity
	case cmp.IsNaN(complex_):
		// Normalize any NaN's mixed with valid numbers.
		return Undefined
	default:
		// Lock onto 0, -1, 1, -i, i, and ∞ if necessary.
		var r = lockMagnitude(real(complex_))
		var i = lockMagnitude(imag(complex_))
		return number_(complex(r, i))
	}
}

// This constructor creates a new number from the specified polar values.
func NumberFromPolar(magnitude float64, phase float64) abs.NumberLike {
	var complex_ = cmp.Rect(magnitude, phase)
	return NumberFromComplex(complex_)
}

// This constructor creates a new number from the specified string value.
func NumberFromString(string_ string) abs.NumberLike {
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
	var number = NumberFromComplex(complex_)
	return number
}

// This constructor returns the minimum value for a number.
func MinimumNumber() abs.NumberLike {
	return Zero
}

// This constructor returns the maximum value for a number.
func MaximumNumber() abs.NumberLike {
	return Infinity
}

// NUMBER METHODS

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
			string_ = imaginaryFromFloat(imagPart)
		default:
			string_ += "("
			string_ += stringFromFloat(realPart)
			string_ += ", "
			string_ += imaginaryFromFloat(imagPart)
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

// NUMBER LIBRARY

// This singleton creates a unique name space for the library functions for
// number elements.
var Number = &numbers_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for number elements.
type numbers_ struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified number.
func (l *numbers_) Inverse(number abs.NumberLike) abs.NumberLike {
	return NumberFromComplex(-number.AsComplex())
}

// This library function returns the sum of the specified numbers.
func (l *numbers_) Sum(first, second abs.NumberLike) abs.NumberLike {
	return NumberFromComplex(first.AsComplex() + second.AsComplex())
}

// This library function returns the difference of the specified numbers.
func (l *numbers_) Difference(first, second abs.NumberLike) abs.NumberLike {
	return NumberFromComplex(first.AsComplex() - second.AsComplex())
}

// This library function returns the specified number scaled by the specified
// factor.
func (l *numbers_) Scaled(number abs.NumberLike, factor float64) abs.NumberLike {
	return l.Product(number, NumberFromComplex(complex(factor, 0)))
}

// NUMERICAL INTERFACE

// This library function returns the reciprocal of the specified number.
func (l *numbers_) Reciprocal(number abs.NumberLike) abs.NumberLike {
	return NumberFromComplex(1.0 / number.AsComplex())
}

// This library function returns the complex conjugate of the specified number.
func (l *numbers_) Conjugate(number abs.NumberLike) abs.NumberLike {
	return NumberFromComplex(cmp.Conj(number.AsComplex()))
}

// This library function returns the product of the specified numbers.
func (l *numbers_) Product(first, second abs.NumberLike) abs.NumberLike {
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		return Undefined
	case first.IsInfinite() && !second.IsZero():
		// Infinity times anything other than zero is infinite.
		return Infinity
	case second.IsInfinite() && !first.IsZero():
		// Anything other than zero times infinity is infinite.
		return Infinity
	default:
		return NumberFromComplex(first.AsComplex() * second.AsComplex())
	}
}

// This library function returns the quotient of the specified numbers.
func (l *numbers_) Quotient(first, second abs.NumberLike) abs.NumberLike {
	switch {
	case first.IsUndefined() || second.IsUndefined():
		// Any undefined arguments result in an undefined result.
		return Undefined
	case first.IsZero() && second.IsZero():
		// Zero divided by zero is undefined.
		return Undefined
	case first.IsInfinite() && second.IsInfinite():
		// Infinity divided by infinity is undefined.
		return Undefined
	case first.IsZero() && !second.IsZero():
		// Zero divided by anything other than zero is zero.
		return Zero
	case second.IsZero() && !first.IsZero():
		// Anything other than zero divided by zero is infinite.
		return Infinity
	case first.IsInfinite() && !second.IsInfinite():
		// Zero divided by anything other than zero is zero.
		return Infinity
	case second.IsInfinite() && !first.IsInfinite():
		// Anything other than zero divided by zero is infinite.
		return Zero
	default:
		return NumberFromComplex(first.AsComplex() / second.AsComplex())
	}
}

// This library function returns the remainder of the specified numbers.
func (l *numbers_) Remainder(first, second abs.NumberLike) abs.NumberLike {
	var m1 = cmp.Abs(first.AsComplex())
	var p1 = cmp.Phase(first.AsComplex())
	var m2 = cmp.Abs(second.AsComplex())
	var p2 = cmp.Phase(second.AsComplex())
	var magnitude = lockMagnitude(mat.Remainder(m1, m2))
	var phase = lockPhase(p2 - p1)
	return NumberFromPolar(magnitude, phase)
}

// This library function returns the result of raising the specified base to the
// specified exponent.
func (l *numbers_) Power(base, exponent abs.NumberLike) abs.NumberLike {
	switch {
	case base.IsUndefined() || exponent.IsUndefined():
		// Any undefined arguments result in an undefined result.
		return Undefined
	case exponent.IsZero():
		// Anything to the zero power is 1 by definition.
		return One
	case base.IsZero():
		// Zero to any power other than zero is still zero.
		return Zero
	case base.IsInfinite():
		// Infinity to any power other than zero is infinite.
		return Infinity
	case exponent.IsInfinite():
		var magnitude = base.GetMagnitude()
		switch {
		case magnitude < 1:
			// A magnitude less than one to an infinite power is zero.
			return Zero
		case magnitude == 1:
			// A magnitude equal to one to an infinite power is one.
			return One
		case magnitude > 1:
			// A magnitude greater than one to an infinite power is infinite.
			return Infinity
		default:
			panic(fmt.Sprintf("An impossible magnitude was encountered: %v", magnitude))
		}
	default:
		return NumberFromComplex(cmp.Pow(base.AsComplex(), exponent.AsComplex()))
	}
}

// This library function returns the result of taking the logarithm using the
// specified base of the specified number.
func (l *numbers_) Logarithm(base, number abs.NumberLike) abs.NumberLike {
	// logB(z) => ln(z) / ln(b)
	var lnB = cmp.Log(base.AsComplex())
	var lnZ = cmp.Log(number.AsComplex())
	var logB = lnZ / lnB
	return NumberFromComplex(logB)
}

// PRIVATE FUNCTIONS

// This function uses the single precision floating point range to lock a double
// precision magnitude onto 0, 1, -1, or ∞ if the magnitude falls outside the
// single precision range for these values. Otherwise, the magnitude is returned
// unchanged.
func lockMagnitude(v float64) float64 {
	var v32 float32 = float32(v)
	switch {
	case mat.Abs(v) <= 1.2246467991473515e-16:
		return 0
	case v32 == -1:
		return -1
	case v32 == 1:
		return 1
	case mat.IsInf(v, 0):
		return mat.Inf(1)
	default:
		return v
	}
}

// This function uses the single precision floating point range to lock a double
// precision phase onto 0, π/2, π, or 3π/2 if the phase falls outside the single
// precision range for these values. Otherwise, the phase is returned unchanged.
func lockPhase(v float64) float64 {
	var v32 float32 = float32(v)
	switch {
	case mat.Abs(v) <= 1.2246467991473515e-16:
		return 0
	case v32 == float32(0.5*mat.Pi):
		return 0.5 * mat.Pi
	case v32 == float32(mat.Pi):
		return mat.Pi
	case v32 == float32(1.5*mat.Pi):
		return 1.5 * mat.Pi
	default:
		return v
	}
}

// This returns the imaginary string for the specified floating point number.
func imaginaryFromFloat(imaginary float64) string {
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
