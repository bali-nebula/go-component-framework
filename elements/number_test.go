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

func TestZeroNumbers(t *testing.T) {
	var v = elements.NumberFromComplex(0 + 0i)
	assert.Equal(t, 0+0i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, "0", v.AsString())
}

func TestPositiveNumbers(t *testing.T) {
	var v = elements.NumberFromComplex(0.25)
	assert.Equal(t, 0.25+0i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.25, v.AsReal())
	assert.Equal(t, "0.25", v.AsString())
}

func TestNegativeNumbers(t *testing.T) {
	var v = elements.NumberFromComplex(-0.75)
	assert.Equal(t, -0.75+0i, complex128(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -1, v.AsInteger())
	assert.Equal(t, -0.75, v.AsReal())
	assert.Equal(t, "-0.75", v.AsString())
}

func TestStringNumbers(t *testing.T) {
	var v, ok = elements.NumberFromString("0")
	assert.True(t, ok)
	assert.Equal(t, 0+0i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, "0", v.AsString())

	v, ok = elements.NumberFromString("0.4")
	assert.True(t, ok)
	assert.Equal(t, 0.4+0i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.4, v.AsReal())
	assert.Equal(t, "0.4", v.AsString())

	v, ok = elements.NumberFromString("-0.4i")
	assert.True(t, ok)
	assert.Equal(t, 0-0.4i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, "-0.4i", v.AsString())

	v, ok = elements.NumberFromString("-1.4E17i")
	assert.True(t, ok)
	assert.Equal(t, 0-1.4e17i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, 0.0, v.AsReal())
	assert.Equal(t, "-1.4E+17i", v.AsString())

	v, ok = elements.NumberFromString("(2.3, -1.4E17i)")
	assert.True(t, ok)
	assert.Equal(t, 2.3-1.4e17i, complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 2, v.AsInteger())
	assert.Equal(t, 2.3, v.AsReal())
	assert.Equal(t, "(2.3, -1.4E+17i)", v.AsString())

	v, ok = elements.NumberFromString("(-2.3, ~1.5i)")
	assert.True(t, ok)
	assert.Equal(t, -0.16269556383571668-2.294238469189325i, complex128(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, 0, v.AsInteger())
	assert.Equal(t, -0.16269556383571668, v.AsReal())
	assert.Equal(t, "(-0.16269556383571668, -2.294238469189325i)", v.AsString())

	v, ok = elements.NumberFromString("∞")
	assert.True(t, ok)
	assert.Equal(t, cmplx.Inf(), complex128(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, math.MaxInt, v.AsInteger())
	assert.Equal(t, math.Inf(1), v.AsReal())
	assert.Equal(t, "∞", v.AsString())

	v, ok = elements.NumberFromString("undefined")
	assert.True(t, ok)
	assert.True(t, cmplx.IsNaN(complex128(v))) // NaN != NaN
	assert.False(t, v.IsNegative())
	assert.Equal(t, math.MinInt, v.AsInteger())
	assert.True(t, math.IsNaN(v.AsReal())) // NaN != NaN
	assert.Equal(t, "undefined", v.AsString())

	v, ok = elements.NumberFromString("bad")
	assert.False(t, ok)
	assert.Equal(t, 0+0i, complex128(v))
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
