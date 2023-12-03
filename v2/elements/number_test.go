/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

var Number = ele.Number()

func TestZero(t *tes.T) {
	var v = Number.FromComplex(0 + 0i)
	ass.Equal(t, 0+0i, v.AsComplex())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestInfinity(t *tes.T) {
	var v = Number.FromComplex(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.AsComplex())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, mat.Inf(1), v.AsFloat())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
}

func TestUndefined(t *tes.T) {
	var v = Number.FromComplex(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.AsComplex()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsFloat()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositivePureReals(t *tes.T) {
	var v = Number.FromComplex(0.25)
	ass.Equal(t, 0.25+0i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = Number.FromComplex(0.25i)
	ass.Equal(t, 0+0.25i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = Number.FromComplex(-0.75)
	ass.Equal(t, -0.75+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = Number.FromComplex(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumberFromPolar(t *tes.T) {
	var v = Number.FromPolar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())
}

func TestNumberFromString(t *tes.T) {
	var v = Number.FromString("(1e^~πi)")
	ass.Equal(t, -1.0+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())
	_ = Number.FromString("(1.2, -3.4i)")
	_ = Number.FromString("undefined")
	_ = Number.FromString("infinity")
	_ = Number.FromString("∞")
	_ = Number.FromString("+1")
	_ = Number.FromString("1")
	_ = Number.FromString("-π")
	_ = Number.FromString("+i")
	_ = Number.FromString("i")
	_ = Number.FromString("-i")
}

func TestNumberLibrary(t *tes.T) {
	var zero = Number.Zero()
	var i = Number.I()
	var minusi = Number.FromComplex(-1i)
	var half = Number.FromComplex(0.5)
	var minushalf = Number.FromComplex(-0.5)
	var one = Number.One()
	var minusone = Number.FromComplex(-1)
	var two = Number.FromComplex(2.0)
	var minustwo = Number.FromComplex(-2.0)
	var infinity = Number.Infinity()
	var undefined = Number.Undefined()

	//	-z
	ass.Equal(t, zero, Number.Inverse(zero))
	ass.Equal(t, minushalf, Number.Inverse(half))
	ass.Equal(t, minusone, Number.Inverse(one))
	ass.Equal(t, minusi, Number.Inverse(i))
	ass.Equal(t, infinity, Number.Inverse(infinity))
	ass.True(t, Number.Inverse(undefined).IsUndefined())

	//	z + zero => z
	ass.Equal(t, minusi, Number.Sum(minusi, zero))
	ass.Equal(t, minusone, Number.Sum(minusone, zero))
	ass.Equal(t, zero, Number.Sum(zero, zero))
	ass.Equal(t, one, Number.Sum(one, zero))
	ass.Equal(t, i, Number.Sum(i, zero))
	ass.Equal(t, infinity, Number.Sum(infinity, zero))
	ass.True(t, Number.Sum(undefined, zero).IsUndefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, Number.Sum(minusi, infinity))
	ass.Equal(t, infinity, Number.Sum(minusone, infinity))
	ass.Equal(t, infinity, Number.Sum(zero, infinity))
	ass.Equal(t, infinity, Number.Sum(one, infinity))
	ass.Equal(t, infinity, Number.Sum(i, infinity))
	ass.Equal(t, infinity, Number.Sum(infinity, infinity))
	ass.True(t, Number.Sum(undefined, infinity).IsUndefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, Number.Difference(minusi, infinity))
	ass.Equal(t, infinity, Number.Difference(minusone, infinity))
	ass.Equal(t, infinity, Number.Difference(zero, infinity))
	ass.Equal(t, infinity, Number.Difference(one, infinity))
	ass.Equal(t, infinity, Number.Difference(i, infinity))
	ass.True(t, Number.Difference(infinity, infinity).IsUndefined())
	ass.True(t, Number.Difference(undefined, infinity).IsUndefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, Number.Difference(infinity, minusi))
	ass.Equal(t, infinity, Number.Difference(infinity, minusone))
	ass.Equal(t, infinity, Number.Difference(infinity, zero))
	ass.Equal(t, infinity, Number.Difference(infinity, one))
	ass.Equal(t, infinity, Number.Difference(infinity, i))
	ass.True(t, Number.Difference(infinity, undefined).IsUndefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, Number.Difference(minusi, minusi))
	ass.Equal(t, zero, Number.Difference(minusone, minusone))
	ass.Equal(t, zero, Number.Difference(zero, zero))
	ass.Equal(t, zero, Number.Difference(one, one))
	ass.Equal(t, zero, Number.Difference(i, i))
	ass.True(t, Number.Difference(infinity, infinity).IsUndefined())
	ass.True(t, Number.Difference(undefined, undefined).IsUndefined())

	//	z * r
	ass.Equal(t, minusi, Number.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, Number.Scaled(minusone, 0.5))
	ass.Equal(t, zero, Number.Scaled(zero, 5.0))
	ass.Equal(t, half, Number.Scaled(one, 0.5))
	ass.Equal(t, i, Number.Scaled(i, 1.0))
	ass.Equal(t, infinity, Number.Scaled(infinity, 5.0))
	ass.True(t, Number.Scaled(undefined, 5.0).IsUndefined())

	//	/z
	ass.Equal(t, infinity, Number.Reciprocal(zero))
	ass.Equal(t, two, Number.Reciprocal(half))
	ass.Equal(t, one, Number.Reciprocal(one))
	ass.Equal(t, minushalf, Number.Reciprocal(minustwo))
	ass.Equal(t, minusi, Number.Reciprocal(i))
	ass.Equal(t, zero, Number.Reciprocal(infinity))
	ass.True(t, Number.Reciprocal(undefined).IsUndefined())

	//	*z
	ass.Equal(t, zero, Number.Conjugate(zero))
	ass.Equal(t, one, Number.Conjugate(one))
	ass.Equal(t, minusi, Number.Conjugate(i))
	ass.Equal(t, i, Number.Conjugate(minusi))
	ass.True(t, Number.Conjugate(undefined).IsUndefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, Number.Product(zero, zero))
	ass.Equal(t, zero, Number.Product(one, zero))
	ass.Equal(t, zero, Number.Product(i, zero))
	ass.True(t, Number.Product(infinity, zero).IsUndefined())
	ass.True(t, Number.Product(undefined, zero).IsUndefined())

	//	z * one => z
	ass.Equal(t, zero, Number.Product(zero, one))
	ass.Equal(t, one, Number.Product(one, one))
	ass.Equal(t, i, Number.Product(i, one))
	ass.Equal(t, infinity, Number.Product(infinity, one))
	ass.True(t, Number.Product(undefined, one).IsUndefined())

	//	z * infinity => infinity  {z != zero}
	ass.True(t, Number.Product(zero, infinity).IsUndefined())
	ass.Equal(t, infinity, Number.Product(one, infinity))
	ass.Equal(t, infinity, Number.Product(i, infinity))
	ass.Equal(t, infinity, Number.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.True(t, Number.Quotient(zero, zero).IsUndefined())
	ass.Equal(t, zero, Number.Quotient(zero, one))
	ass.Equal(t, zero, Number.Quotient(zero, i))
	ass.Equal(t, zero, Number.Quotient(zero, infinity))
	ass.True(t, Number.Quotient(zero, undefined).IsUndefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, Number.Quotient(one, zero))
	ass.Equal(t, infinity, Number.Quotient(i, zero))
	ass.Equal(t, infinity, Number.Quotient(infinity, zero))
	ass.True(t, Number.Quotient(undefined, zero).IsUndefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, Number.Quotient(one, infinity))
	ass.Equal(t, zero, Number.Quotient(i, infinity))
	ass.True(t, Number.Quotient(infinity, infinity).IsUndefined())
	ass.True(t, Number.Quotient(undefined, infinity).IsUndefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, Number.Quotient(infinity, zero))
	ass.Equal(t, infinity, Number.Quotient(infinity, one))
	ass.Equal(t, infinity, Number.Quotient(infinity, i))
	ass.True(t, Number.Quotient(infinity, undefined).IsUndefined())

	//	y / z
	ass.Equal(t, one, Number.Quotient(one, one))
	ass.Equal(t, one, Number.Quotient(i, i))
	ass.Equal(t, i, Number.Quotient(i, one))
	ass.Equal(t, two, Number.Quotient(one, half))
	ass.Equal(t, one, Number.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, Number.Power(minusi, zero))
	ass.Equal(t, one, Number.Power(minusone, zero))
	ass.Equal(t, one, Number.Power(zero, zero))
	ass.Equal(t, one, Number.Power(one, zero))
	ass.Equal(t, one, Number.Power(i, zero))
	ass.Equal(t, one, Number.Power(infinity, zero))
	ass.True(t, Number.Power(undefined, zero).IsUndefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, Number.Power(zero, one))
	ass.Equal(t, zero, Number.Power(zero, i))
	ass.Equal(t, zero, Number.Power(zero, infinity))
	ass.True(t, Number.Power(zero, undefined).IsUndefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, Number.Power(minustwo, infinity))
	ass.Equal(t, one, Number.Power(minusi, infinity))
	ass.Equal(t, one, Number.Power(minusone, infinity))
	ass.Equal(t, zero, Number.Power(minushalf, infinity))
	ass.Equal(t, zero, Number.Power(half, infinity))
	ass.Equal(t, one, Number.Power(one, infinity))
	ass.Equal(t, one, Number.Power(i, infinity))
	ass.Equal(t, infinity, Number.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, Number.Power(infinity, zero))
	ass.Equal(t, infinity, Number.Power(infinity, one))
	ass.Equal(t, infinity, Number.Power(infinity, i))
	ass.Equal(t, infinity, Number.Power(infinity, infinity))
	ass.True(t, Number.Power(infinity, undefined).IsUndefined())

	//	one ^ z => one
	ass.Equal(t, one, Number.Power(one, one))
	ass.Equal(t, one, Number.Power(one, i))
	ass.Equal(t, one, Number.Power(one, minusone))
	ass.Equal(t, one, Number.Power(one, minusi))

	//	log(zero, z) => zero
	ass.True(t, Number.Logarithm(zero, zero).IsUndefined())
	ass.Equal(t, zero, Number.Logarithm(zero, i))
	ass.Equal(t, zero, Number.Logarithm(zero, one))
	ass.True(t, Number.Logarithm(zero, infinity).IsUndefined())
	ass.True(t, Number.Logarithm(zero, undefined).IsUndefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, Number.Logarithm(one, zero))
	ass.True(t, Number.Logarithm(one, one).IsUndefined())
	ass.Equal(t, infinity, Number.Logarithm(one, infinity))
	ass.True(t, Number.Logarithm(one, undefined).IsUndefined())

	//	log(infinity, z) => zero
	ass.True(t, Number.Logarithm(infinity, zero).IsUndefined())
	ass.Equal(t, zero, Number.Logarithm(infinity, one))
	ass.True(t, Number.Logarithm(infinity, infinity).IsUndefined())
	ass.True(t, Number.Logarithm(infinity, undefined).IsUndefined())
}
