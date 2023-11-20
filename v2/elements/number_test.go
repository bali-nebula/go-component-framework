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

func TestZero(t *tes.T) {
	var v = ele.NumberFromComplex(0 + 0i)
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
	var v = ele.NumberFromComplex(cmp.Inf())
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
	var v = ele.NumberFromComplex(cmp.NaN())
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
	var v = ele.NumberFromComplex(0.25)
	ass.Equal(t, 0.25+0i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = ele.NumberFromComplex(0.25i)
	ass.Equal(t, 0+0.25i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = ele.NumberFromComplex(-0.75)
	ass.Equal(t, -0.75+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = ele.NumberFromComplex(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumberFromPolar(t *tes.T) {
	var v = ele.NumberFromPolar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())
}

func TestNumberFromString(t *tes.T) {
	var v = ele.NumberFromString("(1e^~πi)")
	ass.Equal(t, -1.0+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())
	_ = ele.NumberFromString("(1.2, -3.4i)")
	_ = ele.NumberFromString("undefined")
	_ = ele.NumberFromString("infinity")
	_ = ele.NumberFromString("∞")
	_ = ele.NumberFromString("+1")
	_ = ele.NumberFromString("1")
	_ = ele.NumberFromString("-π")
	_ = ele.NumberFromString("+i")
	_ = ele.NumberFromString("i")
	_ = ele.NumberFromString("-i")
}

func TestNumberLibrary(t *tes.T) {
	var zero = ele.Zero
	var i = ele.I
	var minusi = ele.NumberFromComplex(-1i)
	var half = ele.NumberFromComplex(0.5)
	var minushalf = ele.NumberFromComplex(-0.5)
	var one = ele.One
	var minusone = ele.NumberFromComplex(-1)
	var two = ele.NumberFromComplex(2.0)
	var minustwo = ele.NumberFromComplex(-2.0)
	var infinity = ele.Infinity
	var undefined = ele.Undefined

	//	-z
	ass.Equal(t, zero, ele.Number.Inverse(zero))
	ass.Equal(t, minushalf, ele.Number.Inverse(half))
	ass.Equal(t, minusone, ele.Number.Inverse(one))
	ass.Equal(t, minusi, ele.Number.Inverse(i))
	ass.Equal(t, infinity, ele.Number.Inverse(infinity))
	ass.True(t, ele.Number.Inverse(undefined).IsUndefined())

	//	z + zero => z
	ass.Equal(t, minusi, ele.Number.Sum(minusi, zero))
	ass.Equal(t, minusone, ele.Number.Sum(minusone, zero))
	ass.Equal(t, zero, ele.Number.Sum(zero, zero))
	ass.Equal(t, one, ele.Number.Sum(one, zero))
	ass.Equal(t, i, ele.Number.Sum(i, zero))
	ass.Equal(t, infinity, ele.Number.Sum(infinity, zero))
	ass.True(t, ele.Number.Sum(undefined, zero).IsUndefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, ele.Number.Sum(minusi, infinity))
	ass.Equal(t, infinity, ele.Number.Sum(minusone, infinity))
	ass.Equal(t, infinity, ele.Number.Sum(zero, infinity))
	ass.Equal(t, infinity, ele.Number.Sum(one, infinity))
	ass.Equal(t, infinity, ele.Number.Sum(i, infinity))
	ass.Equal(t, infinity, ele.Number.Sum(infinity, infinity))
	ass.True(t, ele.Number.Sum(undefined, infinity).IsUndefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Number.Difference(minusi, infinity))
	ass.Equal(t, infinity, ele.Number.Difference(minusone, infinity))
	ass.Equal(t, infinity, ele.Number.Difference(zero, infinity))
	ass.Equal(t, infinity, ele.Number.Difference(one, infinity))
	ass.Equal(t, infinity, ele.Number.Difference(i, infinity))
	ass.True(t, ele.Number.Difference(infinity, infinity).IsUndefined())
	ass.True(t, ele.Number.Difference(undefined, infinity).IsUndefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Number.Difference(infinity, minusi))
	ass.Equal(t, infinity, ele.Number.Difference(infinity, minusone))
	ass.Equal(t, infinity, ele.Number.Difference(infinity, zero))
	ass.Equal(t, infinity, ele.Number.Difference(infinity, one))
	ass.Equal(t, infinity, ele.Number.Difference(infinity, i))
	ass.True(t, ele.Number.Difference(infinity, undefined).IsUndefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, ele.Number.Difference(minusi, minusi))
	ass.Equal(t, zero, ele.Number.Difference(minusone, minusone))
	ass.Equal(t, zero, ele.Number.Difference(zero, zero))
	ass.Equal(t, zero, ele.Number.Difference(one, one))
	ass.Equal(t, zero, ele.Number.Difference(i, i))
	ass.True(t, ele.Number.Difference(infinity, infinity).IsUndefined())
	ass.True(t, ele.Number.Difference(undefined, undefined).IsUndefined())

	//	z * r
	ass.Equal(t, minusi, ele.Number.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, ele.Number.Scaled(minusone, 0.5))
	ass.Equal(t, zero, ele.Number.Scaled(zero, 5.0))
	ass.Equal(t, half, ele.Number.Scaled(one, 0.5))
	ass.Equal(t, i, ele.Number.Scaled(i, 1.0))
	ass.Equal(t, infinity, ele.Number.Scaled(infinity, 5.0))
	ass.True(t, ele.Number.Scaled(undefined, 5.0).IsUndefined())

	//	/z
	ass.Equal(t, infinity, ele.Number.Reciprocal(zero))
	ass.Equal(t, two, ele.Number.Reciprocal(half))
	ass.Equal(t, one, ele.Number.Reciprocal(one))
	ass.Equal(t, minushalf, ele.Number.Reciprocal(minustwo))
	ass.Equal(t, minusi, ele.Number.Reciprocal(i))
	ass.Equal(t, zero, ele.Number.Reciprocal(infinity))
	ass.True(t, ele.Number.Reciprocal(undefined).IsUndefined())

	//	*z
	ass.Equal(t, zero, ele.Number.Conjugate(zero))
	ass.Equal(t, one, ele.Number.Conjugate(one))
	ass.Equal(t, minusi, ele.Number.Conjugate(i))
	ass.Equal(t, i, ele.Number.Conjugate(minusi))
	ass.True(t, ele.Number.Conjugate(undefined).IsUndefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, ele.Number.Product(zero, zero))
	ass.Equal(t, zero, ele.Number.Product(one, zero))
	ass.Equal(t, zero, ele.Number.Product(i, zero))
	ass.True(t, ele.Number.Product(infinity, zero).IsUndefined())
	ass.True(t, ele.Number.Product(undefined, zero).IsUndefined())

	//	z * one => z
	ass.Equal(t, zero, ele.Number.Product(zero, one))
	ass.Equal(t, one, ele.Number.Product(one, one))
	ass.Equal(t, i, ele.Number.Product(i, one))
	ass.Equal(t, infinity, ele.Number.Product(infinity, one))
	ass.True(t, ele.Number.Product(undefined, one).IsUndefined())

	//	z * infinity => infinity  {z != zero}
	ass.True(t, ele.Number.Product(zero, infinity).IsUndefined())
	ass.Equal(t, infinity, ele.Number.Product(one, infinity))
	ass.Equal(t, infinity, ele.Number.Product(i, infinity))
	ass.Equal(t, infinity, ele.Number.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.True(t, ele.Number.Quotient(zero, zero).IsUndefined())
	ass.Equal(t, zero, ele.Number.Quotient(zero, one))
	ass.Equal(t, zero, ele.Number.Quotient(zero, i))
	ass.Equal(t, zero, ele.Number.Quotient(zero, infinity))
	ass.True(t, ele.Number.Quotient(zero, undefined).IsUndefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, ele.Number.Quotient(one, zero))
	ass.Equal(t, infinity, ele.Number.Quotient(i, zero))
	ass.Equal(t, infinity, ele.Number.Quotient(infinity, zero))
	ass.True(t, ele.Number.Quotient(undefined, zero).IsUndefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, ele.Number.Quotient(one, infinity))
	ass.Equal(t, zero, ele.Number.Quotient(i, infinity))
	ass.True(t, ele.Number.Quotient(infinity, infinity).IsUndefined())
	ass.True(t, ele.Number.Quotient(undefined, infinity).IsUndefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, ele.Number.Quotient(infinity, zero))
	ass.Equal(t, infinity, ele.Number.Quotient(infinity, one))
	ass.Equal(t, infinity, ele.Number.Quotient(infinity, i))
	ass.True(t, ele.Number.Quotient(infinity, undefined).IsUndefined())

	//	y / z
	ass.Equal(t, one, ele.Number.Quotient(one, one))
	ass.Equal(t, one, ele.Number.Quotient(i, i))
	ass.Equal(t, i, ele.Number.Quotient(i, one))
	ass.Equal(t, two, ele.Number.Quotient(one, half))
	ass.Equal(t, one, ele.Number.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, ele.Number.Power(minusi, zero))
	ass.Equal(t, one, ele.Number.Power(minusone, zero))
	ass.Equal(t, one, ele.Number.Power(zero, zero))
	ass.Equal(t, one, ele.Number.Power(one, zero))
	ass.Equal(t, one, ele.Number.Power(i, zero))
	ass.Equal(t, one, ele.Number.Power(infinity, zero))
	ass.True(t, ele.Number.Power(undefined, zero).IsUndefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, ele.Number.Power(zero, one))
	ass.Equal(t, zero, ele.Number.Power(zero, i))
	ass.Equal(t, zero, ele.Number.Power(zero, infinity))
	ass.True(t, ele.Number.Power(zero, undefined).IsUndefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, ele.Number.Power(minustwo, infinity))
	ass.Equal(t, one, ele.Number.Power(minusi, infinity))
	ass.Equal(t, one, ele.Number.Power(minusone, infinity))
	ass.Equal(t, zero, ele.Number.Power(minushalf, infinity))
	ass.Equal(t, zero, ele.Number.Power(half, infinity))
	ass.Equal(t, one, ele.Number.Power(one, infinity))
	ass.Equal(t, one, ele.Number.Power(i, infinity))
	ass.Equal(t, infinity, ele.Number.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, ele.Number.Power(infinity, zero))
	ass.Equal(t, infinity, ele.Number.Power(infinity, one))
	ass.Equal(t, infinity, ele.Number.Power(infinity, i))
	ass.Equal(t, infinity, ele.Number.Power(infinity, infinity))
	ass.True(t, ele.Number.Power(infinity, undefined).IsUndefined())

	//	one ^ z => one
	ass.Equal(t, one, ele.Number.Power(one, one))
	ass.Equal(t, one, ele.Number.Power(one, i))
	ass.Equal(t, one, ele.Number.Power(one, minusone))
	ass.Equal(t, one, ele.Number.Power(one, minusi))

	//	log(zero, z) => zero
	ass.True(t, ele.Number.Logarithm(zero, zero).IsUndefined())
	ass.Equal(t, zero, ele.Number.Logarithm(zero, i))
	ass.Equal(t, zero, ele.Number.Logarithm(zero, one))
	ass.True(t, ele.Number.Logarithm(zero, infinity).IsUndefined())
	ass.True(t, ele.Number.Logarithm(zero, undefined).IsUndefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, ele.Number.Logarithm(one, zero))
	ass.True(t, ele.Number.Logarithm(one, one).IsUndefined())
	ass.Equal(t, infinity, ele.Number.Logarithm(one, infinity))
	ass.True(t, ele.Number.Logarithm(one, undefined).IsUndefined())

	//	log(infinity, z) => zero
	ass.True(t, ele.Number.Logarithm(infinity, zero).IsUndefined())
	ass.Equal(t, zero, ele.Number.Logarithm(infinity, one))
	ass.True(t, ele.Number.Logarithm(infinity, infinity).IsUndefined())
	ass.True(t, ele.Number.Logarithm(infinity, undefined).IsUndefined())
}
