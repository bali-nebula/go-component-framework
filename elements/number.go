/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"math"
	"math/cmplx"
	"strconv"
)

// NUMBER INTERFACE

// This constructor creates a new number that is mapped to the Riemann Sphere.
//   - https://en.wikipedia.org/wiki/Riemann_sphere
//
// This mapping removes many of the idiosyncrasies associated with the normal
// complex plane. There is only one value for zero and one value for infinity.
// This simplifies a lot of the mathematical operations:
//
//	z + zero => zero
//	z + infinity => infinity
//
//	z - infinity => infinity  {z != infinity}
//	infinity - z => infinity  {z != infinity}
//
//	z * zero => zero          {z != infinity}
//	z * infinity => infinity  {z != zero}
//
//	zero / z => zero          {z != zero}
//	z / zero => infinity      {z != zero}
//
//	z / infinity => zero      {z != infinity}
//	infinity / z => infinity  {z != infinity}
//
//	z ^ zero => one           {definition}
//	zero ^ z => zero          {z != zero}
//
//	one ^ z => one
//	z ^ one => z
//
//	z ^ infinity => infinity  {z != zero}
//	infinity ^ z => infinity  {z != zero}
//
// This leaves only the following operations undefined:
//
//	infinity - infinity => undefined
//	zero * infinity => undefined
//	zero / zero => undefined
//	infinity / infinity => undefined
//	-one ^ z => undefined       {z != one}
//	-i ^ z => undefined       {z != one}
//	+i ^ z => undefined       {z != one}
//
// The resulting number system is easier to use for most applications. For
// numerical analysis the ANSI plus and minus zero values are often used as a
// crutch that leads to misleading convergence information for oscillating
// functions. In the case of numerical analysis it is probably better to track
// the course of the function as it converges than to look at the final value.
func NumberFromComplex(v complex128) Number {
	switch {
	case cmplx.Abs(v) == 0:
		// Normalize all versions of zero.
		return Zero
	case cmplx.IsInf(v):
		// Normalize any negative infinities or infinite i's.
		return Infinity
	case cmplx.IsNaN(v):
		// Normalize any NaN's mixed with valid numbers.
		return Undefined
	default:
		// Normalize real or imaginary negative zeros.
		var r = real(v)
		var i = imag(v)
		if r == 0 {
			v = complex(0, i)
		}
		if i == 0 {
			v = complex(r, 0)
		}
		return Number(v)
	}
}

// This constructor attempts to create a new number from the specified formatted
// string. It returns an number value and whether or not the string
// contained a valid number. Some common examples of numbers in this format
// include:
//   - 0
//   - -1
//   - 5i
//   - 1.23E-45
//   - (-3.0, 4.0i) - rectangular form
//   - (1.0, ~π)    - polar form
//   - ∞
//   - undefined
//
// For valid string formats for this type see `../abstractions/language.go`.
func NumberFromString(v string) (Number, bool) {
	var number, ok = stringToNumber(v)
	return Number(number), ok
}

// This type defines the methods associated with number elements. It extends the
// native Go complex128 type and may represent an integer, real, or complex
// number.
type Number complex128

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Number) AsString() string {
	if v.IsInfinite() {
		return "∞"
	}
	if v.IsUndefined() {
		return "undefined"
	}
	var realPart = strconv.FormatFloat(real(v), 'G', -1, 64)
	var imagPart = strconv.FormatFloat(imag(v), 'G', -1, 64) + "i"
	if imag(v) == 0 {
		return realPart
	}
	if real(v) == 0 {
		return imagPart
	}
	return "(" + realPart + ", " + imagPart + ")"
}

// This method implements the standard Go Stringer interface.
func (v Number) String() string {
	return v.AsString()
}

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v Number) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v Number) AsInteger() int {
	if v.IsInfinite() {
		return math.MaxInt
	}
	if v.IsUndefined() {
		return math.MinInt
	}
	return int(math.Round(v.AsReal()))
}

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous component.
func (v Number) AsReal() float64 {
	return real(v)
}

// COMPLEX INTERFACE

// This method determines whether or not this number is infinite.
func (v Number) IsZero() bool {
	return cmplx.Abs(complex128(v)) == 0
}

// This method determines whether or not this number is infinite.
func (v Number) IsInfinite() bool {
	return cmplx.IsInf(complex128(v))
}

// This method determines whether or not this number is undefined.
func (v Number) IsUndefined() bool {
	return cmplx.IsNaN(complex128(v))
}

// This method returns the real part of this complex component.
func (v Number) GetReal() float64 {
	return real(v)
}

// This method returns the imaginary part of this complex component.
func (v Number) GetImaginary() float64 {
	return imag(v)
}

// This method returns the magnitude of this complex component.
func (v Number) GetMagnitude() float64 {
	return cmplx.Abs(complex128(v))
}

// This method returns the phase angle of this complex component.
func (v Number) GetPhase() Angle {
	return Angle(cmplx.Phase(complex128(v)))
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Number) IsNegative() bool {
	return real(v) < 0
}

// NUMBERS LIBRARY

var zero = complex(0, 0)
var infinity = cmplx.Inf()
var undefined = cmplx.NaN()

var I = Number(complex(0.0, 1.0))
var One = Number(complex(1.0, 0.0))
var Phi = Number(complex(math.Phi, 0.0))
var Zero = Number(zero)
var Infinity = Number(infinity)
var Undefined = Number(undefined)

// This singleton creates a unique name space for the library functions for
// number elements.
var Numbers = &numbers{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for number elements.
type numbers struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified number.
func (l *numbers) Inverse(number Number) Number {
	var result = complex128(-number)
	return NumberFromComplex(result)
}

// This library function returns the sum of the specified numbers.
func (l *numbers) Sum(first, second Number) Number {
	var result = complex128(first + second)
	return NumberFromComplex(result)
}

// This library function returns the difference of the specified numbers.
func (l *numbers) Difference(first, second Number) Number {
	var result = complex128(first - second)
	return NumberFromComplex(result)
}

// This library function returns the specified number scaled by the specified
// factor.
func (l *numbers) Scaled(number Number, factor float64) Number {
	return l.Product(number, NumberFromComplex(complex(factor, 0)))
}

// NUMERICAL INTERFACE

// This library function returns the reciprocal of the specified number.
func (l *numbers) Reciprocal(number Number) Number {
	var result = 1.0 / complex128(number)
	return NumberFromComplex(result)
}

// This library function returns the complex conjugate of the specified number.
func (l *numbers) Conjugate(number Number) Number {
	var result = cmplx.Conj(complex128(number))
	return NumberFromComplex(result)
}

// This library function returns the product of the specified numbers.
func (l *numbers) Product(first, second Number) Number {
	switch {
	case first.IsInfinite() && !second.IsZero():
		// Infinity times anything other than zero is infinite.
		return Infinity
	case second.IsInfinite() && !first.IsZero():
		// Anything other than zero times infinity is infinite.
		return Infinity
	default:
		return NumberFromComplex(complex128(first) * complex128(second))
	}
}

// This library function returns the quotient of the specified numbers.
func (l *numbers) Quotient(first, second Number) Number {
	return l.Product(first, l.Reciprocal(second))
}

// This library function returns the remainder of the specified numbers.
func (l *numbers) Remainder(first, second Number) Number {
	var m1 = cmplx.Abs(complex128(first))
	var p1 = cmplx.Phase(complex128(first))
	var m2 = cmplx.Abs(complex128(second))
	var p2 = cmplx.Phase(complex128(second))
	var magnitude = math.Remainder(m1, m2)
	var phase = p2 - p1
	var result = cmplx.Rect(magnitude, phase)
	return NumberFromComplex(result)
}

// This library function returns the result of raising the specified base to the
// specified power.
func (l *numbers) Exponential(base, power Number) Number {
	switch {
	case base.IsUndefined() || power.IsUndefined():
		// Undefined values always result in undefined results.
		return Undefined
	case power.IsZero():
		// Anything to the zero power is 1 by definition.
		return One
	case base == One:
		// One to any power is still one.
		return One
	case power == One:
		// Anything to the first power is stays the same.
		return base
	case base.GetMagnitude() == 1:
		// The phase for any other magnitude of one is undefined.
		return Undefined
	case base.IsZero():
		// Zero to any power other than zero is still zero.
		return Zero
	case base.IsInfinite():
		// Infinity to any power other than zero is infinite.
		return Infinity
	case power.IsInfinite():
		// Anything other than zero to an infinite power is infinite.
		return Infinity
	default:
		return NumberFromComplex(cmplx.Pow(complex128(base), complex128(power)))
	}
}

// This library function returns the result of taking the logarithm using the
// specified base of the specified number.
func (l *numbers) Logarithm(base, number Number) Number {
	// logB(z) => ln(z) / ln(b)
	return l.Quotient(
		NumberFromComplex(cmplx.Log(complex128(number))),
		NumberFromComplex(cmplx.Log(complex128(base))))
}

// PRIVATE FUNCTIONS

// This function parses a number string and returns the corresponding complex
// number.
func stringToNumber(v string) (complex128, bool) {
	var number complex128
	var ok = true
	var matches = abstractions.ScanNumber([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	case matches[1] == "undefined":
		// The value is undefined.
		number = undefined
	case matches[1] == "infinity" || matches[1] == "∞":
		// The value is infinity.
		number = infinity
	case matches[1][0] == '(':
		// The value is complex.
		var realPart, err = strconv.ParseFloat(matches[2], 64)
		if err != nil {
		}
		if matches[3][0] == '~' {
			// The complex number is in polar form.
			var imaginaryPart, err = strconv.ParseFloat(matches[3][1:], 64)
			if err != nil {
			}
			number = cmplx.Rect(realPart, imaginaryPart)
		} else {
			// The complex number is in rectangular form.
			var imaginaryPart, err = strconv.ParseFloat(matches[3], 64)
			if err != nil {
			}
			number = complex(realPart, imaginaryPart)
		}
	case matches[1][len(matches[1])-1] == 'i':
		// The value is pure imaginary.
		var imaginaryPart, err = strconv.ParseFloat(matches[1][:len(matches[1])-1], 64)
		if err != nil {
		}
		number = complex(0, imaginaryPart)
	default:
		// The value is pure real.
		var realPart, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
		}
		number = complex(realPart, 0)
	}
	return number, ok
}
