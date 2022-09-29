/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/stretchr/testify/assert"
	"math"
	"math/cmplx"
	"testing"
)

func TestZero(t *testing.T) {
	var v = elements.NumberFromComplex(0 + 0i)
	assert.Equal(t, 0+0i, complex128(v))
	assert.True(t, v.IsZero())
	assert.False(t, v.IsInfinite())
	assert.False(t, v.IsUndefined())
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "0", v.AsString())
}

func TestInfinity(t *testing.T) {
	var v = elements.NumberFromComplex(cmplx.Inf())
	assert.Equal(t, cmplx.Inf(), complex128(v))
	assert.False(t, v.IsZero())
	assert.True(t, v.IsInfinite())
	assert.False(t, v.IsUndefined())
	assert.False(t, v.IsNegative())
	assert.Equal(t, math.MaxInt, v.AsInteger())
	assert.Equal(t, math.Inf(1), v.AsReal())
	assert.Equal(t, math.Inf(1), v.GetReal())
	assert.Equal(t, math.Inf(1), v.GetImaginary())
	assert.Equal(t, "∞", v.AsString())
}

func TestUndefined(t *testing.T) {
	var v = elements.NumberFromComplex(cmplx.NaN())
	assert.True(t, cmplx.IsNaN(complex128(v)))
	assert.False(t, v.IsZero())
	assert.False(t, v.IsInfinite())
	assert.True(t, v.IsUndefined())
	assert.False(t, v.IsNegative())
	assert.Equal(t, math.MinInt, v.AsInteger())
	assert.True(t, math.IsNaN(v.AsReal()))
	assert.True(t, math.IsNaN(v.GetReal()))
	assert.True(t, math.IsNaN(v.GetImaginary()))
	assert.Equal(t, "undefined", v.AsString())
}

func TestPositiveReals(t *testing.T) {
	var ok bool
	var v elements.Number

	v = elements.NumberFromComplex(0.25)
	assert.Equal(t, 0.25+0i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.25, v.AsReal())
	assert.Equal(t, 0.25, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "0.25", v.AsString())

	v, ok = elements.NumberFromString("pi")
	assert.True(t, ok)
	assert.Equal(t, math.Pi, real(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 3, v.AsInteger())
	assert.Equal(t, math.Pi, v.AsReal())
	assert.Equal(t, math.Pi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "π", v.AsString())

	v, ok = elements.NumberFromString("phi")
	assert.True(t, ok)
	assert.Equal(t, math.Phi, real(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 2, v.AsInteger())
	assert.Equal(t, math.Phi, v.AsReal())
	assert.Equal(t, math.Phi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "φ", v.AsString())

	v, ok = elements.NumberFromString("tau")
	assert.True(t, ok)
	assert.Equal(t, 2.0*math.Pi, real(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 6, v.AsInteger())
	assert.Equal(t, 2.0*math.Pi, v.AsReal())
	assert.Equal(t, 2.0*math.Pi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "τ", v.AsString())
}

func TestPositiveImaginaries(t *testing.T) {
	var ok bool
	var v elements.Number

	v = elements.NumberFromComplex(0.25i)
	assert.Equal(t, 0+0.25i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, 0.25, v.GetImaginary())
	assert.Equal(t, "0.25i", v.AsString())

	v, ok = elements.NumberFromString("pii")
	assert.True(t, ok)
	assert.Equal(t, math.Pi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, math.Pi, v.GetImaginary())
	assert.Equal(t, "πi", v.AsString())

	v, ok = elements.NumberFromString("phii")
	assert.True(t, ok)
	assert.Equal(t, math.Phi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, math.Phi, v.GetImaginary())
	assert.Equal(t, "φi", v.AsString())

	v, ok = elements.NumberFromString("taui")
	assert.True(t, ok)
	assert.Equal(t, 2.0*math.Pi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, 2.0*math.Pi, v.GetImaginary())
	assert.Equal(t, "τi", v.AsString())
}

func TestNegativeReals(t *testing.T) {
	var ok bool
	var v elements.Number

	v = elements.NumberFromComplex(-0.75)
	assert.Equal(t, -0.75+0i, complex128(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -1, v.AsInteger())
	assert.Equal(t, -0.75, v.AsReal())
	assert.Equal(t, -0.75, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "-0.75", v.AsString())

	v, ok = elements.NumberFromString("-pi")
	assert.True(t, ok)
	assert.Equal(t, -math.Pi, real(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -3, v.AsInteger())
	assert.Equal(t, -math.Pi, v.AsReal())
	assert.Equal(t, -math.Pi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "-π", v.AsString())

	v, ok = elements.NumberFromString("-phi")
	assert.True(t, ok)
	assert.Equal(t, -math.Phi, real(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -2, v.AsInteger())
	assert.Equal(t, -math.Phi, v.AsReal())
	assert.Equal(t, -math.Phi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "-φ", v.AsString())

	v, ok = elements.NumberFromString("-tau")
	assert.True(t, ok)
	assert.Equal(t, -2.0*math.Pi, real(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -6, v.AsInteger())
	assert.Equal(t, -2.0*math.Pi, v.AsReal())
	assert.Equal(t, -2.0*math.Pi, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "-τ", v.AsString())
}

func TestNegativeImaginaries(t *testing.T) {
	var ok bool
	var v elements.Number

	v = elements.NumberFromComplex(-0.75i)
	assert.Equal(t, 0-0.75i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, -0.75, v.GetImaginary())
	assert.Equal(t, "-0.75i", v.AsString())

	v, ok = elements.NumberFromString("-pii")
	assert.True(t, ok)
	assert.Equal(t, -math.Pi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, -math.Pi, v.GetImaginary())
	assert.Equal(t, "-πi", v.AsString())

	v, ok = elements.NumberFromString("-phii")
	assert.True(t, ok)
	assert.Equal(t, -math.Phi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, -math.Phi, v.GetImaginary())
	assert.Equal(t, "-φi", v.AsString())

	v, ok = elements.NumberFromString("-taui")
	assert.True(t, ok)
	assert.Equal(t, -2.0*math.Pi, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, 0.0, v.GetReal())
	assert.Equal(t, -2.0*math.Pi, v.GetImaginary())
	assert.Equal(t, "-τi", v.AsString())
}

func TestPolarNumbers(t *testing.T) {
	var ok bool
	var v elements.Number

	v, ok = elements.NumberFromString("(5e^~0.9272952180016123i)")
	assert.True(t, ok)
	assert.Equal(t, 3.0, real(v))
	assert.Equal(t, 4.0, imag(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 3, v.AsInteger())
	assert.Equal(t, 3.0, v.AsReal())
	assert.Equal(t, 3.0, v.GetReal())
	assert.Equal(t, 4.0, v.GetImaginary())
	assert.Equal(t, "(3, 4i)", v.AsString())
	assert.Equal(t, "(5e^~0.9272952180016122i)", v.AsPolar())

	v, ok = elements.NumberFromString("(1e^~πi)")
	assert.True(t, ok)
	assert.Equal(t, -1.0, real(v))
	assert.Equal(t, 0.0, imag(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -1, v.AsInteger())
	assert.Equal(t, -1.0, v.AsReal())
	assert.Equal(t, -1.0, v.GetReal())
	assert.Equal(t, 0.0, v.GetImaginary())
	assert.Equal(t, "-1", v.AsString())
	assert.Equal(t, "(1e^~πi)", v.AsPolar())
}

func TestRoundtripNumbers(t *testing.T) {
	for _, s := range numbers {
		var v, ok = elements.NumberFromString(s)
		assert.True(t, ok)
		assert.Equal(t, s, v.AsString())
	}
}

func TestNumbersLibrary(t *testing.T) {
	var zero = elements.Zero
	var i = elements.I
	var one = elements.One
	var five = elements.NumberFromComplex(5)
	var infinity = elements.Infinity
	var undefined = elements.Undefined

	assert.Equal(t, zero, elements.Numbers.Inverse(zero))
	assert.Equal(t, -i, elements.Numbers.Inverse(i))
	assert.Equal(t, -one, elements.Numbers.Inverse(one))
	assert.Equal(t, infinity, elements.Numbers.Inverse(infinity))
	assert.True(t, elements.Numbers.Inverse(undefined).IsUndefined())

	assert.Equal(t, one, elements.Numbers.Sum(zero, one))
	assert.Equal(t, i, elements.Numbers.Sum(zero, i))
	assert.Equal(t, zero, elements.Numbers.Difference(one, one))
	assert.Equal(t, zero, elements.Numbers.Difference(i, i))
	assert.Equal(t, infinity, elements.Numbers.Sum(zero, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(zero, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(i, infinity))
	assert.True(t, elements.Numbers.Sum(zero, undefined).IsUndefined())
	assert.True(t, elements.Numbers.Sum(undefined, i).IsUndefined())
	assert.True(t, elements.Numbers.Difference(undefined, zero).IsUndefined())
	assert.True(t, elements.Numbers.Difference(i, undefined).IsUndefined())

	assert.Equal(t, zero, elements.Numbers.Scaled(zero, 5.0))
	assert.Equal(t, five, elements.Numbers.Scaled(one, 5.0))
	assert.Equal(t, i, elements.Numbers.Scaled(i, 1.0))
	assert.Equal(t, infinity, elements.Numbers.Scaled(infinity, 5.0))
	assert.True(t, elements.Numbers.Scaled(undefined, 5.0).IsUndefined())

	assert.Equal(t, infinity, elements.Numbers.Reciprocal(zero))
	assert.Equal(t, one, elements.Numbers.Reciprocal(one))
	assert.Equal(t, -i, elements.Numbers.Reciprocal(i))
	assert.Equal(t, zero, elements.Numbers.Reciprocal(infinity))
	assert.True(t, elements.Numbers.Reciprocal(undefined).IsUndefined())

	assert.Equal(t, zero, elements.Numbers.Conjugate(zero))
	assert.Equal(t, one, elements.Numbers.Conjugate(one))
	assert.Equal(t, -i, elements.Numbers.Conjugate(i))
	assert.True(t, elements.Numbers.Conjugate(undefined).IsUndefined())

	assert.Equal(t, zero, elements.Numbers.Product(zero, one))
	assert.Equal(t, one, elements.Numbers.Quotient(one, one))
	assert.Equal(t, zero, elements.Numbers.Product(zero, i))
	assert.Equal(t, one, elements.Numbers.Quotient(i, i))
	assert.Equal(t, i, elements.Numbers.Product(one, i))
	assert.Equal(t, i, elements.Numbers.Quotient(i, one))
	assert.Equal(t, five, elements.Numbers.Product(one, five))
	assert.Equal(t, one, elements.Numbers.Quotient(five, five))
	assert.True(t, elements.Numbers.Product(zero, infinity).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Quotient(zero, infinity))
	assert.True(t, elements.Numbers.Product(zero, undefined).IsUndefined())
	assert.True(t, elements.Numbers.Quotient(undefined, infinity).IsUndefined())

	assert.Equal(t, one, elements.Numbers.Exponential(zero, zero))
	assert.Equal(t, zero, elements.Numbers.Exponential(zero, one))
	assert.Equal(t, zero, elements.Numbers.Exponential(zero, infinity))
	assert.Equal(t, one, elements.Numbers.Exponential(one, zero))
	assert.Equal(t, one, elements.Numbers.Exponential(one, one))
	assert.Equal(t, one, elements.Numbers.Exponential(one, infinity))
	assert.Equal(t, one, elements.Numbers.Exponential(infinity, zero))
	assert.Equal(t, infinity, elements.Numbers.Exponential(infinity, one))
	assert.Equal(t, infinity, elements.Numbers.Exponential(infinity, infinity))

	assert.True(t, elements.Numbers.Exponential(undefined, zero).IsUndefined())
	assert.True(t, elements.Numbers.Exponential(zero, undefined).IsUndefined())
	assert.True(t, elements.Numbers.Exponential(undefined, one).IsUndefined())
	assert.True(t, elements.Numbers.Exponential(one, undefined).IsUndefined())
	assert.True(t, elements.Numbers.Exponential(undefined, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Exponential(infinity, undefined).IsUndefined())

	assert.Equal(t, zero, elements.Numbers.Exponential(zero, i))
	assert.Equal(t, one, elements.Numbers.Exponential(one, i))
	assert.Equal(t, infinity, elements.Numbers.Exponential(infinity, i))

	assert.Equal(t, one, elements.Numbers.Exponential(-one, zero))
	assert.Equal(t, -one, elements.Numbers.Exponential(-one, one))
	assert.True(t, elements.Numbers.Exponential(-one, infinity).IsUndefined())

	assert.Equal(t, one, elements.Numbers.Exponential(i, zero))
	assert.Equal(t, i, elements.Numbers.Exponential(i, one))
	assert.True(t, elements.Numbers.Exponential(i, infinity).IsUndefined())

	assert.Equal(t, one, elements.Numbers.Exponential(-i, zero))
	assert.Equal(t, -i, elements.Numbers.Exponential(-i, one))
	assert.True(t, elements.Numbers.Exponential(-i, infinity).IsUndefined())

	assert.Equal(t, zero, elements.Numbers.Logarithm(zero, i))
	assert.True(t, elements.Numbers.Logarithm(zero, zero).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Logarithm(zero, one))
	assert.True(t, elements.Numbers.Logarithm(zero, infinity).IsUndefined())
	assert.Equal(t, infinity, elements.Numbers.Logarithm(one, zero))
	assert.True(t, elements.Numbers.Logarithm(one, one).IsUndefined())
	assert.Equal(t, infinity, elements.Numbers.Logarithm(one, infinity))
	assert.True(t, elements.Numbers.Logarithm(infinity, zero).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Logarithm(infinity, one))
	assert.True(t, elements.Numbers.Logarithm(infinity, infinity).IsUndefined())
}

var numbers = []string{
	"undefined",
	"0",
	"1",
	"1.2",
	"1.2E+30",
	"1.2E-30",
	"-1",
	"-1.2",
	"-1.2E+30",
	"-1.2E-30",
	"i",
	"1.2i",
	"1.2E+30i",
	"1.2E-30i",
	"-i",
	"-1.2i",
	"-1.2E+30i",
	"-1.2E-30i",
	"e",
	"-e",
	"ei",
	"-ei",
	"π",
	"-π",
	"πi",
	"-πi",
	"φ",
	"-φ",
	"φi",
	"-φi",
	"τ",
	"-τ",
	"τi",
	"-τi",
	"∞",
	"(1, i)",
	"(-1, i)",
	"(1, -i)",
	"(-1, -i)",
	"(π, πi)",
	"(-π, πi)",
	"(π, -πi)",
	"(-π, -πi)",
	"(φ, φi)",
	"(-φ, φi)",
	"(φ, -φi)",
	"(-φ, -φi)",
	"(τ, τi)",
	"(-τ, τi)",
	"(τ, -τi)",
	"(-τ, -τi)",
	"(3, 4i)",
	"(-3, 4i)",
	"(3, -4i)",
	"(-3, -4i)",
	"(3.15, 4.15i)",
	"(-3.15, 4.15i)",
	"(3.15, -4.15i)",
	"(-3.15, -4.15i)",
	"(3.15E+12, 4.15E+12i)",
	"(-3.15E+12, 4.15E+12i)",
	"(3.15E+12, -4.15E+12i)",
	"(-3.15E+12, -4.15E+12i)",
	"(3.15E-12, 4.15E-12i)",
	"(-3.15E-12, 4.15E-12i)",
	"(3.15E-12, -4.15E-12i)",
	"(-3.15E-12, -4.15E-12i)",
}
