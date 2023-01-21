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
	mat "math"
	cmp "math/cmplx"
)

// NUMBER IMPLEMENTATION

var zero = complex(0, 0)
var infinity = cmp.Inf()
var undefined = cmp.NaN()

var I = Number(complex(0.0, 1.0))
var One = Number(complex(1.0, 0.0))
var Phi = Number(complex(mat.Phi, 0.0))
var Zero = Number(zero)
var Infinity = Number(infinity)
var Undefined = Number(undefined)

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
func NumberFromComplex(v complex128) Number {
	switch {
	case cmp.Abs(v) == 0:
		// Normalize all versions of zero.
		return Zero
	case cmp.IsInf(v):
		// Normalize any negative infinities or infinite i's.
		return Infinity
	case cmp.IsNaN(v):
		// Normalize any NaN's mixed with valid numbers.
		return Undefined
	default:
		// Lock onto 0, -1, 1, -i, i, and ∞ if necessary.
		var r = lockMagnitude(real(v))
		var i = lockMagnitude(imag(v))
		return Number(complex(r, i))
	}
}

func NumberFromPolar(m float64, p Angle) Number {
	var v = cmp.Rect(m, float64(p))
	return NumberFromComplex(v)
}

// This constructor returns the minimum value for a number.
func MinimumNumber() Number {
	return Zero
}

// This constructor returns the maximum value for a number.
func MaximumNumber() Number {
	return Infinity
}

// This type defines the methods associated with number elements. It extends the
// native Go complex128 type and may represent an integer, real, or complex
// number.
type Number complex128

// CONTINUOUS INTERFACE

// This method determines whether or not this number is zero.
func (v Number) IsZero() bool {
	return real(v) == 0 && imag(v) == 0
}

// This method determines whether or not this number is infinite.
func (v Number) IsInfinite() bool {
	return mat.IsInf(real(v), 0) || mat.IsInf(imag(v), 0)
}

// This method determines whether or not this number is undefined.
func (v Number) IsUndefined() bool {
	return mat.IsNaN(real(v)) || mat.IsNaN(imag(v))
}

// This method returns a real value for this continuous component.
func (v Number) AsReal() float64 {
	return real(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Number) IsNegative() bool {
	return real(v) < 0
}

// COMPLEX INTERFACE

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
	return lockMagnitude(cmp.Abs(complex128(v)))
}

// This method returns the phase angle of this complex component.
func (v Number) GetPhase() Angle {
	return AngleFromFloat(cmp.Phase(complex128(v)))
}

// NUMBERS LIBRARY

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
	var result = cmp.Conj(complex128(number))
	return NumberFromComplex(result)
}

// This library function returns the product of the specified numbers.
func (l *numbers) Product(first, second Number) Number {
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
		return NumberFromComplex(complex128(first) * complex128(second))
	}
}

// This library function returns the quotient of the specified numbers.
func (l *numbers) Quotient(first, second Number) Number {
	return l.Product(first, l.Reciprocal(second))
}

// This library function returns the remainder of the specified numbers.
func (l *numbers) Remainder(first, second Number) Number {
	var m1 = cmp.Abs(complex128(first))
	var p1 = cmp.Phase(complex128(first))
	var m2 = cmp.Abs(complex128(second))
	var p2 = cmp.Phase(complex128(second))
	var magnitude = lockMagnitude(mat.Remainder(m1, m2))
	var phase = AngleFromFloat(p2 - p1)
	return NumberFromPolar(magnitude, phase)
}

// This library function returns the result of raising the specified base to the
// specified exponent.
func (l *numbers) Power(base, exponent Number) Number {
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
		return NumberFromComplex(cmp.Pow(complex128(base), complex128(exponent)))
	}
}

// This library function returns the result of taking the logarithm using the
// specified base of the specified number.
func (l *numbers) Logarithm(base, number Number) Number {
	// logB(z) => ln(z) / ln(b)
	return l.Quotient(
		NumberFromComplex(cmp.Log(complex128(number))),
		NumberFromComplex(cmp.Log(complex128(base))))
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
