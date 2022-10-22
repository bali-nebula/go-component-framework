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
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

func TestZero(t *tes.T) {
	var v = ele.NumberFromComplex(0 + 0i)
	ass.Equal(t, 0+0i, complex128(v))
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "0", lan.FormatValue(v))
}

func TestInfinity(t *tes.T) {
	var v = ele.NumberFromComplex(cmp.Inf())
	ass.Equal(t, cmp.Inf(), complex128(v))
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, mat.MaxInt, v.AsInteger())
	ass.Equal(t, mat.Inf(1), v.AsReal())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
	ass.Equal(t, "∞", lan.FormatValue(v))
}

func TestUndefined(t *tes.T) {
	var v = ele.NumberFromComplex(cmp.NaN())
	ass.True(t, cmp.IsNaN(complex128(v)))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, mat.MinInt, v.AsInteger())
	ass.True(t, mat.IsNaN(v.AsReal()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
	ass.Equal(t, "undefined", lan.FormatValue(v))
}

func TestPositiveReals(t *tes.T) {
	var ok bool
	var v ele.Number

	v = ele.NumberFromComplex(0.25)
	ass.Equal(t, 0.25+0i, complex128(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.25, v.AsReal())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "0.25", lan.FormatValue(v))

	v, ok = ele.NumberFromString("pi")
	ass.True(t, ok)
	ass.Equal(t, mat.Pi, real(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 3, v.AsInteger())
	ass.Equal(t, mat.Pi, v.AsReal())
	ass.Equal(t, mat.Pi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "π", lan.FormatValue(v))

	v, ok = ele.NumberFromString("phi")
	ass.True(t, ok)
	ass.Equal(t, mat.Phi, real(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 2, v.AsInteger())
	ass.Equal(t, mat.Phi, v.AsReal())
	ass.Equal(t, mat.Phi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "φ", lan.FormatValue(v))

	v, ok = ele.NumberFromString("tau")
	ass.True(t, ok)
	ass.Equal(t, 2.0*mat.Pi, real(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 6, v.AsInteger())
	ass.Equal(t, 2.0*mat.Pi, v.AsReal())
	ass.Equal(t, 2.0*mat.Pi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "τ", lan.FormatValue(v))
}

func TestPositiveImaginaries(t *tes.T) {
	var ok bool
	var v ele.Number

	v = ele.NumberFromComplex(0.25i)
	ass.Equal(t, 0+0.25i, complex128(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
	ass.Equal(t, "0.25i", lan.FormatValue(v))

	v, ok = ele.NumberFromString("pii")
	ass.True(t, ok)
	ass.Equal(t, mat.Pi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, mat.Pi, v.GetImaginary())
	ass.Equal(t, "πi", lan.FormatValue(v))

	v, ok = ele.NumberFromString("phii")
	ass.True(t, ok)
	ass.Equal(t, mat.Phi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, mat.Phi, v.GetImaginary())
	ass.Equal(t, "φi", lan.FormatValue(v))

	v, ok = ele.NumberFromString("taui")
	ass.True(t, ok)
	ass.Equal(t, 2.0*mat.Pi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 2.0*mat.Pi, v.GetImaginary())
	ass.Equal(t, "τi", lan.FormatValue(v))
}

func TestNegativeReals(t *tes.T) {
	var ok bool
	var v ele.Number

	v = ele.NumberFromComplex(-0.75)
	ass.Equal(t, -0.75+0i, complex128(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1, v.AsInteger())
	ass.Equal(t, -0.75, v.AsReal())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "-0.75", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-pi")
	ass.True(t, ok)
	ass.Equal(t, -mat.Pi, real(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -3, v.AsInteger())
	ass.Equal(t, -mat.Pi, v.AsReal())
	ass.Equal(t, -mat.Pi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "-π", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-phi")
	ass.True(t, ok)
	ass.Equal(t, -mat.Phi, real(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -2, v.AsInteger())
	ass.Equal(t, -mat.Phi, v.AsReal())
	ass.Equal(t, -mat.Phi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "-φ", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-tau")
	ass.True(t, ok)
	ass.Equal(t, -2.0*mat.Pi, real(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -6, v.AsInteger())
	ass.Equal(t, -2.0*mat.Pi, v.AsReal())
	ass.Equal(t, -2.0*mat.Pi, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "-τ", lan.FormatValue(v))
}

func TestNegativeImaginaries(t *tes.T) {
	var ok bool
	var v ele.Number

	v = ele.NumberFromComplex(-0.75i)
	ass.Equal(t, 0-0.75i, complex128(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
	ass.Equal(t, "-0.75i", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-pii")
	ass.True(t, ok)
	ass.Equal(t, -mat.Pi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -mat.Pi, v.GetImaginary())
	ass.Equal(t, "-πi", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-phii")
	ass.True(t, ok)
	ass.Equal(t, -mat.Phi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -mat.Phi, v.GetImaginary())
	ass.Equal(t, "-φi", lan.FormatValue(v))

	v, ok = ele.NumberFromString("-taui")
	ass.True(t, ok)
	ass.Equal(t, -2.0*mat.Pi, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -2.0*mat.Pi, v.GetImaginary())
	ass.Equal(t, "-τi", lan.FormatValue(v))
}

func TestPolarNumbers(t *tes.T) {
	var ok bool
	var v ele.Number

	v, ok = ele.NumberFromString("(5e^~0.9272952180016123i)")
	ass.True(t, ok)
	ass.Equal(t, 3.0, real(v))
	ass.Equal(t, 4.0, imag(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 3, v.AsInteger())
	ass.Equal(t, 3.0, v.AsReal())
	ass.Equal(t, 3.0, v.GetReal())
	ass.Equal(t, 4.0, v.GetImaginary())
	ass.Equal(t, "(3, 4i)", lan.FormatValue(v))
	ass.Equal(t, "(5e^~0.9272952180016122i)", v.AsPolar())

	v, ok = ele.NumberFromString("(1e^~πi)")
	ass.True(t, ok)
	ass.Equal(t, -1.0, real(v))
	ass.Equal(t, 0.0, imag(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1, v.AsInteger())
	ass.Equal(t, -1.0, v.AsReal())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, "-1", lan.FormatValue(v))
	ass.Equal(t, "(1e^~πi)", v.AsPolar())
}

func TestRoundtripNumbers(t *tes.T) {
	for _, s := range numbers {
		var v, ok = ele.NumberFromString(s)
		ass.True(t, ok)
		ass.Equal(t, s, lan.FormatValue(v))
	}
}

func TestNumbersLibrary(t *tes.T) {
	var zero = ele.Zero
	var i = ele.I
	var half = ele.NumberFromComplex(0.5)
	var one = ele.One
	var two = ele.One * 2.0
	var infinity = ele.Infinity
	var undefined = ele.Undefined

	//	-z
	ass.Equal(t, zero, ele.Numbers.Inverse(zero))
	ass.Equal(t, -half, ele.Numbers.Inverse(half))
	ass.Equal(t, -one, ele.Numbers.Inverse(one))
	ass.Equal(t, -i, ele.Numbers.Inverse(i))
	ass.Equal(t, infinity, ele.Numbers.Inverse(infinity))
	ass.True(t, ele.Numbers.Inverse(undefined).IsUndefined())

	//	z + zero => z
	ass.Equal(t, -i, ele.Numbers.Sum(-i, zero))
	ass.Equal(t, -one, ele.Numbers.Sum(-one, zero))
	ass.Equal(t, zero, ele.Numbers.Sum(zero, zero))
	ass.Equal(t, one, ele.Numbers.Sum(one, zero))
	ass.Equal(t, i, ele.Numbers.Sum(i, zero))
	ass.Equal(t, infinity, ele.Numbers.Sum(infinity, zero))
	ass.True(t, ele.Numbers.Sum(undefined, zero).IsUndefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, ele.Numbers.Sum(-i, infinity))
	ass.Equal(t, infinity, ele.Numbers.Sum(-one, infinity))
	ass.Equal(t, infinity, ele.Numbers.Sum(zero, infinity))
	ass.Equal(t, infinity, ele.Numbers.Sum(one, infinity))
	ass.Equal(t, infinity, ele.Numbers.Sum(i, infinity))
	ass.Equal(t, infinity, ele.Numbers.Sum(infinity, infinity))
	ass.True(t, ele.Numbers.Sum(undefined, infinity).IsUndefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Numbers.Difference(-i, infinity))
	ass.Equal(t, infinity, ele.Numbers.Difference(-one, infinity))
	ass.Equal(t, infinity, ele.Numbers.Difference(zero, infinity))
	ass.Equal(t, infinity, ele.Numbers.Difference(one, infinity))
	ass.Equal(t, infinity, ele.Numbers.Difference(i, infinity))
	ass.True(t, ele.Numbers.Difference(infinity, infinity).IsUndefined())
	ass.True(t, ele.Numbers.Difference(undefined, infinity).IsUndefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Numbers.Difference(infinity, -i))
	ass.Equal(t, infinity, ele.Numbers.Difference(infinity, -one))
	ass.Equal(t, infinity, ele.Numbers.Difference(infinity, zero))
	ass.Equal(t, infinity, ele.Numbers.Difference(infinity, one))
	ass.Equal(t, infinity, ele.Numbers.Difference(infinity, i))
	ass.True(t, ele.Numbers.Difference(infinity, undefined).IsUndefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, ele.Numbers.Difference(-i, -i))
	ass.Equal(t, zero, ele.Numbers.Difference(-one, -one))
	ass.Equal(t, zero, ele.Numbers.Difference(zero, zero))
	ass.Equal(t, zero, ele.Numbers.Difference(one, one))
	ass.Equal(t, zero, ele.Numbers.Difference(i, i))
	ass.True(t, ele.Numbers.Difference(infinity, infinity).IsUndefined())
	ass.True(t, ele.Numbers.Difference(undefined, undefined).IsUndefined())

	//	z * r
	ass.Equal(t, -i, ele.Numbers.Scaled(-i, 1.0))
	ass.Equal(t, -half, ele.Numbers.Scaled(-one, 0.5))
	ass.Equal(t, zero, ele.Numbers.Scaled(zero, 5.0))
	ass.Equal(t, half, ele.Numbers.Scaled(one, 0.5))
	ass.Equal(t, i, ele.Numbers.Scaled(i, 1.0))
	ass.Equal(t, infinity, ele.Numbers.Scaled(infinity, 5.0))
	ass.True(t, ele.Numbers.Scaled(undefined, 5.0).IsUndefined())

	//	/z
	ass.Equal(t, infinity, ele.Numbers.Reciprocal(zero))
	ass.Equal(t, two, ele.Numbers.Reciprocal(half))
	ass.Equal(t, one, ele.Numbers.Reciprocal(one))
	ass.Equal(t, -half, ele.Numbers.Reciprocal(-two))
	ass.Equal(t, -i, ele.Numbers.Reciprocal(i))
	ass.Equal(t, zero, ele.Numbers.Reciprocal(infinity))
	ass.True(t, ele.Numbers.Reciprocal(undefined).IsUndefined())

	//	*z
	ass.Equal(t, zero, ele.Numbers.Conjugate(zero))
	ass.Equal(t, one, ele.Numbers.Conjugate(one))
	ass.Equal(t, -i, ele.Numbers.Conjugate(i))
	ass.Equal(t, i, ele.Numbers.Conjugate(-i))
	ass.True(t, ele.Numbers.Conjugate(undefined).IsUndefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, ele.Numbers.Product(zero, zero))
	ass.Equal(t, zero, ele.Numbers.Product(one, zero))
	ass.Equal(t, zero, ele.Numbers.Product(i, zero))
	ass.True(t, ele.Numbers.Product(infinity, zero).IsUndefined())
	ass.True(t, ele.Numbers.Product(undefined, zero).IsUndefined())

	//	z * one => z
	ass.Equal(t, zero, ele.Numbers.Product(zero, one))
	ass.Equal(t, one, ele.Numbers.Product(one, one))
	ass.Equal(t, i, ele.Numbers.Product(i, one))
	ass.Equal(t, infinity, ele.Numbers.Product(infinity, one))
	ass.True(t, ele.Numbers.Product(undefined, one).IsUndefined())

	//	z * infinity => infinity  {z != zero}
	ass.True(t, ele.Numbers.Product(zero, infinity).IsUndefined())
	ass.Equal(t, infinity, ele.Numbers.Product(one, infinity))
	ass.Equal(t, infinity, ele.Numbers.Product(i, infinity))
	ass.Equal(t, infinity, ele.Numbers.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.True(t, ele.Numbers.Quotient(zero, zero).IsUndefined())
	ass.Equal(t, zero, ele.Numbers.Quotient(zero, one))
	ass.Equal(t, zero, ele.Numbers.Quotient(zero, i))
	ass.Equal(t, zero, ele.Numbers.Quotient(zero, infinity))
	ass.True(t, ele.Numbers.Quotient(zero, undefined).IsUndefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, ele.Numbers.Quotient(one, zero))
	ass.Equal(t, infinity, ele.Numbers.Quotient(i, zero))
	ass.Equal(t, infinity, ele.Numbers.Quotient(infinity, zero))
	ass.True(t, ele.Numbers.Quotient(undefined, zero).IsUndefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, ele.Numbers.Quotient(one, infinity))
	ass.Equal(t, zero, ele.Numbers.Quotient(i, infinity))
	ass.True(t, ele.Numbers.Quotient(infinity, infinity).IsUndefined())
	ass.True(t, ele.Numbers.Quotient(undefined, infinity).IsUndefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Numbers.Quotient(infinity, zero))
	ass.Equal(t, infinity, ele.Numbers.Quotient(infinity, one))
	ass.Equal(t, infinity, ele.Numbers.Quotient(infinity, i))
	ass.True(t, ele.Numbers.Quotient(infinity, undefined).IsUndefined())

	//	y / z
	ass.Equal(t, one, ele.Numbers.Quotient(one, one))
	ass.Equal(t, one, ele.Numbers.Quotient(i, i))
	ass.Equal(t, i, ele.Numbers.Quotient(i, one))
	ass.Equal(t, two, ele.Numbers.Quotient(one, half))
	ass.Equal(t, one, ele.Numbers.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, ele.Numbers.Power(-i, zero))
	ass.Equal(t, one, ele.Numbers.Power(-one, zero))
	ass.Equal(t, one, ele.Numbers.Power(zero, zero))
	ass.Equal(t, one, ele.Numbers.Power(one, zero))
	ass.Equal(t, one, ele.Numbers.Power(i, zero))
	ass.Equal(t, one, ele.Numbers.Power(infinity, zero))
	ass.True(t, ele.Numbers.Power(undefined, zero).IsUndefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, ele.Numbers.Power(zero, one))
	ass.Equal(t, zero, ele.Numbers.Power(zero, i))
	ass.Equal(t, zero, ele.Numbers.Power(zero, infinity))
	ass.True(t, ele.Numbers.Power(zero, undefined).IsUndefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, ele.Numbers.Power(-two, infinity))
	ass.Equal(t, one, ele.Numbers.Power(-i, infinity))
	ass.Equal(t, one, ele.Numbers.Power(-one, infinity))
	ass.Equal(t, zero, ele.Numbers.Power(-half, infinity))
	ass.Equal(t, zero, ele.Numbers.Power(half, infinity))
	ass.Equal(t, one, ele.Numbers.Power(one, infinity))
	ass.Equal(t, one, ele.Numbers.Power(i, infinity))
	ass.Equal(t, infinity, ele.Numbers.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, ele.Numbers.Power(infinity, zero))
	ass.Equal(t, infinity, ele.Numbers.Power(infinity, one))
	ass.Equal(t, infinity, ele.Numbers.Power(infinity, i))
	ass.Equal(t, infinity, ele.Numbers.Power(infinity, infinity))
	ass.True(t, ele.Numbers.Power(infinity, undefined).IsUndefined())

	//	one ^ z => one
	ass.Equal(t, one, ele.Numbers.Power(one, one))
	ass.Equal(t, one, ele.Numbers.Power(one, i))
	ass.Equal(t, one, ele.Numbers.Power(one, -one))
	ass.Equal(t, one, ele.Numbers.Power(one, -i))

	//	log(zero, z) => zero
	ass.True(t, ele.Numbers.Logarithm(zero, zero).IsUndefined())
	ass.Equal(t, zero, ele.Numbers.Logarithm(zero, i))
	ass.Equal(t, zero, ele.Numbers.Logarithm(zero, one))
	ass.True(t, ele.Numbers.Logarithm(zero, infinity).IsUndefined())
	ass.True(t, ele.Numbers.Logarithm(zero, undefined).IsUndefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, ele.Numbers.Logarithm(one, zero))
	ass.True(t, ele.Numbers.Logarithm(one, one).IsUndefined())
	ass.Equal(t, infinity, ele.Numbers.Logarithm(one, infinity))
	ass.True(t, ele.Numbers.Logarithm(one, undefined).IsUndefined())

	//	log(infinity, z) => zero
	ass.True(t, ele.Numbers.Logarithm(infinity, zero).IsUndefined())
	ass.Equal(t, zero, ele.Numbers.Logarithm(infinity, one))
	ass.True(t, ele.Numbers.Logarithm(infinity, infinity).IsUndefined())
	ass.True(t, ele.Numbers.Logarithm(infinity, undefined).IsUndefined())
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
