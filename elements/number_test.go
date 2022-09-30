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
	var half = elements.NumberFromComplex(0.5)
	var one = elements.One
	var two = elements.One * 2.0
	var infinity = elements.Infinity
	var undefined = elements.Undefined

	//	-z
	assert.Equal(t, zero, elements.Numbers.Inverse(zero))
	assert.Equal(t, -half, elements.Numbers.Inverse(half))
	assert.Equal(t, -one, elements.Numbers.Inverse(one))
	assert.Equal(t, -i, elements.Numbers.Inverse(i))
	assert.Equal(t, infinity, elements.Numbers.Inverse(infinity))
	assert.True(t, elements.Numbers.Inverse(undefined).IsUndefined())

	//	z + zero => z
	assert.Equal(t, -i, elements.Numbers.Sum(-i, zero))
	assert.Equal(t, -one, elements.Numbers.Sum(-one, zero))
	assert.Equal(t, zero, elements.Numbers.Sum(zero, zero))
	assert.Equal(t, one, elements.Numbers.Sum(one, zero))
	assert.Equal(t, i, elements.Numbers.Sum(i, zero))
	assert.Equal(t, infinity, elements.Numbers.Sum(infinity, zero))
	assert.True(t, elements.Numbers.Sum(undefined, zero).IsUndefined())

	//	z + infinity => infinity
	assert.Equal(t, infinity, elements.Numbers.Sum(-i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(-one, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(zero, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(one, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Sum(infinity, infinity))
	assert.True(t, elements.Numbers.Sum(undefined, infinity).IsUndefined())

	//	z - infinity => infinity  {z != infinity}
	assert.Equal(t, infinity, elements.Numbers.Difference(-i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(-one, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(zero, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(one, infinity))
	assert.Equal(t, infinity, elements.Numbers.Difference(i, infinity))
	assert.True(t, elements.Numbers.Difference(infinity, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Difference(undefined, infinity).IsUndefined())

	//	infinity - z => infinity  {z != infinity}
	assert.Equal(t, infinity, elements.Numbers.Difference(infinity, -i))
	assert.Equal(t, infinity, elements.Numbers.Difference(infinity, -one))
	assert.Equal(t, infinity, elements.Numbers.Difference(infinity, zero))
	assert.Equal(t, infinity, elements.Numbers.Difference(infinity, one))
	assert.Equal(t, infinity, elements.Numbers.Difference(infinity, i))
	assert.True(t, elements.Numbers.Difference(infinity, undefined).IsUndefined())

	//	z - z => zero  {z != infinity}
	assert.Equal(t, zero, elements.Numbers.Difference(-i, -i))
	assert.Equal(t, zero, elements.Numbers.Difference(-one, -one))
	assert.Equal(t, zero, elements.Numbers.Difference(zero, zero))
	assert.Equal(t, zero, elements.Numbers.Difference(one, one))
	assert.Equal(t, zero, elements.Numbers.Difference(i, i))
	assert.True(t, elements.Numbers.Difference(infinity, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Difference(undefined, undefined).IsUndefined())

	//	z * r
	assert.Equal(t, -i, elements.Numbers.Scaled(-i, 1.0))
	assert.Equal(t, -half, elements.Numbers.Scaled(-one, 0.5))
	assert.Equal(t, zero, elements.Numbers.Scaled(zero, 5.0))
	assert.Equal(t, half, elements.Numbers.Scaled(one, 0.5))
	assert.Equal(t, i, elements.Numbers.Scaled(i, 1.0))
	assert.Equal(t, infinity, elements.Numbers.Scaled(infinity, 5.0))
	assert.True(t, elements.Numbers.Scaled(undefined, 5.0).IsUndefined())

	//	/z
	assert.Equal(t, infinity, elements.Numbers.Reciprocal(zero))
	assert.Equal(t, two, elements.Numbers.Reciprocal(half))
	assert.Equal(t, one, elements.Numbers.Reciprocal(one))
	assert.Equal(t, -half, elements.Numbers.Reciprocal(-two))
	assert.Equal(t, -i, elements.Numbers.Reciprocal(i))
	assert.Equal(t, zero, elements.Numbers.Reciprocal(infinity))
	assert.True(t, elements.Numbers.Reciprocal(undefined).IsUndefined())

	//	*z
	assert.Equal(t, zero, elements.Numbers.Conjugate(zero))
	assert.Equal(t, one, elements.Numbers.Conjugate(one))
	assert.Equal(t, -i, elements.Numbers.Conjugate(i))
	assert.Equal(t, i, elements.Numbers.Conjugate(-i))
	assert.True(t, elements.Numbers.Conjugate(undefined).IsUndefined())

	//	z * zero => zero          {z != infinity}
	assert.Equal(t, zero, elements.Numbers.Product(zero, zero))
	assert.Equal(t, zero, elements.Numbers.Product(one, zero))
	assert.Equal(t, zero, elements.Numbers.Product(i, zero))
	assert.True(t, elements.Numbers.Product(infinity, zero).IsUndefined())
	assert.True(t, elements.Numbers.Product(undefined, zero).IsUndefined())

	//	z * one => z
	assert.Equal(t, zero, elements.Numbers.Product(zero, one))
	assert.Equal(t, one, elements.Numbers.Product(one, one))
	assert.Equal(t, i, elements.Numbers.Product(i, one))
	assert.Equal(t, infinity, elements.Numbers.Product(infinity, one))
	assert.True(t, elements.Numbers.Product(undefined, one).IsUndefined())

	//	z * infinity => infinity  {z != zero}
	assert.True(t, elements.Numbers.Product(zero, infinity).IsUndefined())
	assert.Equal(t, infinity, elements.Numbers.Product(one, infinity))
	assert.Equal(t, infinity, elements.Numbers.Product(i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	assert.True(t, elements.Numbers.Quotient(zero, zero).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Quotient(zero, one))
	assert.Equal(t, zero, elements.Numbers.Quotient(zero, i))
	assert.Equal(t, zero, elements.Numbers.Quotient(zero, infinity))
	assert.True(t, elements.Numbers.Quotient(zero, undefined).IsUndefined())

	//	z / zero => infinity      {z != zero}
	assert.Equal(t, infinity, elements.Numbers.Quotient(one, zero))
	assert.Equal(t, infinity, elements.Numbers.Quotient(i, zero))
	assert.Equal(t, infinity, elements.Numbers.Quotient(infinity, zero))
	assert.True(t, elements.Numbers.Quotient(undefined, zero).IsUndefined())

	//	z / infinity => zero      {z != infinity}
	assert.Equal(t, zero, elements.Numbers.Quotient(one, infinity))
	assert.Equal(t, zero, elements.Numbers.Quotient(i, infinity))
	assert.True(t, elements.Numbers.Quotient(infinity, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Quotient(undefined, infinity).IsUndefined())

	//	infinity / z => infinity  {z != infinity}
	assert.Equal(t, infinity, elements.Numbers.Quotient(infinity, zero))
	assert.Equal(t, infinity, elements.Numbers.Quotient(infinity, one))
	assert.Equal(t, infinity, elements.Numbers.Quotient(infinity, i))
	assert.True(t, elements.Numbers.Quotient(infinity, undefined).IsUndefined())

	//	y / z
	assert.Equal(t, one, elements.Numbers.Quotient(one, one))
	assert.Equal(t, one, elements.Numbers.Quotient(i, i))
	assert.Equal(t, i, elements.Numbers.Quotient(i, one))
	assert.Equal(t, two, elements.Numbers.Quotient(one, half))
	assert.Equal(t, one, elements.Numbers.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	assert.Equal(t, one, elements.Numbers.Power(-i, zero))
	assert.Equal(t, one, elements.Numbers.Power(-one, zero))
	assert.Equal(t, one, elements.Numbers.Power(zero, zero))
	assert.Equal(t, one, elements.Numbers.Power(one, zero))
	assert.Equal(t, one, elements.Numbers.Power(i, zero))
	assert.Equal(t, one, elements.Numbers.Power(infinity, zero))
	assert.True(t, elements.Numbers.Power(undefined, zero).IsUndefined())

	//	zero ^ z => zero          {z != zero}
	assert.Equal(t, zero, elements.Numbers.Power(zero, one))
	assert.Equal(t, zero, elements.Numbers.Power(zero, i))
	assert.Equal(t, zero, elements.Numbers.Power(zero, infinity))
	assert.True(t, elements.Numbers.Power(zero, undefined).IsUndefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	assert.Equal(t, infinity, elements.Numbers.Power(-two, infinity))
	assert.Equal(t, one, elements.Numbers.Power(-i, infinity))
	assert.Equal(t, one, elements.Numbers.Power(-one, infinity))
	assert.Equal(t, zero, elements.Numbers.Power(-half, infinity))
	assert.Equal(t, zero, elements.Numbers.Power(half, infinity))
	assert.Equal(t, one, elements.Numbers.Power(one, infinity))
	assert.Equal(t, one, elements.Numbers.Power(i, infinity))
	assert.Equal(t, infinity, elements.Numbers.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	assert.Equal(t, one, elements.Numbers.Power(infinity, zero))
	assert.Equal(t, infinity, elements.Numbers.Power(infinity, one))
	assert.Equal(t, infinity, elements.Numbers.Power(infinity, i))
	assert.Equal(t, infinity, elements.Numbers.Power(infinity, infinity))
	assert.True(t, elements.Numbers.Power(infinity, undefined).IsUndefined())

	//	one ^ z => one
	assert.Equal(t, one, elements.Numbers.Power(one, one))
	assert.Equal(t, one, elements.Numbers.Power(one, i))
	assert.Equal(t, one, elements.Numbers.Power(one, -one))
	assert.Equal(t, one, elements.Numbers.Power(one, -i))

	//	log(zero, z) => zero
	assert.True(t, elements.Numbers.Logarithm(zero, zero).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Logarithm(zero, i))
	assert.Equal(t, zero, elements.Numbers.Logarithm(zero, one))
	assert.True(t, elements.Numbers.Logarithm(zero, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Logarithm(zero, undefined).IsUndefined())

	//	log(one, z) => infinity
	assert.Equal(t, infinity, elements.Numbers.Logarithm(one, zero))
	assert.True(t, elements.Numbers.Logarithm(one, one).IsUndefined())
	assert.Equal(t, infinity, elements.Numbers.Logarithm(one, infinity))
	assert.True(t, elements.Numbers.Logarithm(one, undefined).IsUndefined())

	//	log(infinity, z) => zero
	assert.True(t, elements.Numbers.Logarithm(infinity, zero).IsUndefined())
	assert.Equal(t, zero, elements.Numbers.Logarithm(infinity, one))
	assert.True(t, elements.Numbers.Logarithm(infinity, infinity).IsUndefined())
	assert.True(t, elements.Numbers.Logarithm(infinity, undefined).IsUndefined())
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
