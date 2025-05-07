/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element_test

import (
	ele "github.com/bali-nebula/go-component-framework/v3/element"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestZeroAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(0)
	ass.Equal(t, 0.0, v.AsFloat())

	v = Angle.MakeFromString("~0")
	ass.Equal(t, "~0", v.AsString())
}

func TestPositiveAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = Angle.MakeFromString("~π")
	ass.Equal(t, "~π", v.AsString())
}

func TestNegativeAngles(t *tes.T) {
	var Angle = ele.Angle()
	var v = Angle.MakeFromFloat(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = Angle.MakeFromFloat(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var Angle = ele.Angle()
	var v0 = Angle.Zero()
	var v1 = Angle.MakeFromFloat(mat.Pi * 0.25)
	var v2 = Angle.MakeFromFloat(mat.Pi * 0.5)
	var v3 = Angle.MakeFromFloat(mat.Pi * 0.75)
	var v4 = Angle.Pi()
	var v5 = Angle.MakeFromFloat(mat.Pi * 1.25)
	var v6 = Angle.MakeFromFloat(mat.Pi * 1.5)
	var v7 = Angle.MakeFromFloat(mat.Pi * 1.75)
	var v8 = Angle.Tau()

	ass.Equal(t, v4, Angle.Inverse(v0))
	ass.Equal(t, v5, Angle.Inverse(v1))
	ass.Equal(t, v6, Angle.Inverse(v2))
	ass.Equal(t, v7, Angle.Inverse(v3))
	ass.Equal(t, v0, Angle.Inverse(v4))
	ass.Equal(t, v4, Angle.Inverse(v8))

	ass.Equal(t, v1, Angle.Sum(v0, v1))
	ass.Equal(t, v0, Angle.Difference(v1, v1))
	ass.Equal(t, v3, Angle.Sum(v1, v2))
	ass.Equal(t, v1, Angle.Difference(v3, v2))
	ass.Equal(t, v5, Angle.Sum(v2, v3))
	ass.Equal(t, v2, Angle.Difference(v5, v3))
	ass.Equal(t, v7, Angle.Sum(v3, v4))
	ass.Equal(t, v3, Angle.Difference(v7, v4))
	ass.Equal(t, v1, Angle.Sum(v8, v1))
	ass.Equal(t, v0, Angle.Difference(v8, v8))

	ass.Equal(t, v3, Angle.Scaled(v1, 3.0))
	ass.Equal(t, v0, Angle.Scaled(v4, 2.0))
	ass.Equal(t, v4, Angle.Scaled(v4, -1.0))
	ass.Equal(t, v0, Angle.Scaled(v8, 1.0))

	ass.Equal(t, v0, Angle.ArcCosine(Angle.Cosine(v0)))
	ass.Equal(t, v1, Angle.ArcCosine(Angle.Cosine(v1)))
	ass.Equal(t, v2, Angle.ArcCosine(Angle.Cosine(v2)))
	ass.Equal(t, v3, Angle.ArcCosine(Angle.Cosine(v3)))
	ass.Equal(t, v4, Angle.ArcCosine(Angle.Cosine(v4)))
	ass.Equal(t, v0, Angle.ArcCosine(Angle.Cosine(v8)))

	ass.Equal(t, v0, Angle.ArcSine(Angle.Sine(v0)))
	ass.Equal(t, v1, Angle.ArcSine(Angle.Sine(v1)))
	ass.Equal(t, v2, Angle.ArcSine(Angle.Sine(v2)))
	ass.Equal(t, v6, Angle.ArcSine(Angle.Sine(v6)))
	ass.Equal(t, v7, Angle.ArcSine(Angle.Sine(v7)))
	ass.Equal(t, v0, Angle.ArcSine(Angle.Sine(v8)))

	ass.Equal(t, v0, Angle.ArcTangent(Angle.Cosine(v0), Angle.Sine(v0)))
	ass.Equal(t, v1, Angle.ArcTangent(Angle.Cosine(v1), Angle.Sine(v1)))
	ass.Equal(t, v2, Angle.ArcTangent(Angle.Cosine(v2), Angle.Sine(v2)))
	ass.Equal(t, v3, Angle.ArcTangent(Angle.Cosine(v3), Angle.Sine(v3)))
	ass.Equal(t, v4, Angle.ArcTangent(Angle.Cosine(v4), Angle.Sine(v4)))
	ass.Equal(t, v5, Angle.ArcTangent(Angle.Cosine(v5), Angle.Sine(v5)))
	ass.Equal(t, v0, Angle.ArcTangent(Angle.Cosine(v8), Angle.Sine(v8)))
}

func TestFalseBooleans(t *tes.T) {
	var Boolean = ele.Boolean()
	var v = Boolean.MakeFromBoolean(false)
	ass.False(t, v.AsBoolean())
}

func TestTrueBooleans(t *tes.T) {
	var Boolean = ele.Boolean()
	var v = Boolean.MakeFromBoolean(true)
	ass.True(t, v.AsBoolean())
}

func TestBooleansLibrary(t *tes.T) {
	var Boolean = ele.Boolean()
	var T = Boolean.MakeFromBoolean(true)
	var F = Boolean.MakeFromBoolean(false)

	var andNot = Boolean.And(Boolean.Not(T), Boolean.Not(T))
	var notOr = Boolean.Not(Boolean.Or(T, T))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(T), Boolean.Not(F))
	notOr = Boolean.Not(Boolean.Or(T, F))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(F), Boolean.Not(T))
	notOr = Boolean.Not(Boolean.Or(F, T))
	ass.Equal(t, andNot, notOr)

	andNot = Boolean.And(Boolean.Not(F), Boolean.Not(F))
	notOr = Boolean.Not(Boolean.Or(F, F))
	ass.Equal(t, andNot, notOr)

	var sans = Boolean.And(T, Boolean.Not(T))
	ass.Equal(t, sans, Boolean.Sans(T, T))

	sans = Boolean.And(T, Boolean.Not(F))
	ass.Equal(t, sans, Boolean.Sans(T, F))

	sans = Boolean.And(F, Boolean.Not(T))
	ass.Equal(t, sans, Boolean.Sans(F, T))

	sans = Boolean.And(F, Boolean.Not(F))
	ass.Equal(t, sans, Boolean.Sans(F, F))

	var xor = Boolean.Or(Boolean.Sans(T, T), Boolean.Sans(T, T))
	ass.Equal(t, xor, Boolean.Xor(T, T))

	xor = Boolean.Or(Boolean.Sans(T, F), Boolean.Sans(F, T))
	ass.Equal(t, xor, Boolean.Xor(T, F))

	xor = Boolean.Or(Boolean.Sans(F, T), Boolean.Sans(T, F))
	ass.Equal(t, xor, Boolean.Xor(F, T))

	xor = Boolean.Or(Boolean.Sans(F, F), Boolean.Sans(F, F))
	ass.Equal(t, xor, Boolean.Xor(F, F))
}

func TestASCIICharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"a"`)
	ass.Equal(t, `"a"`, v.AsString())
	v = Character.ToUppercase(v)
	ass.Equal(t, `"A"`, v.AsString())
	v = Character.ToLowercase(v)
	ass.Equal(t, `"a"`, v.AsString())

	v = Character.MakeFromString(`"'"`)
	ass.Equal(t, `"'"`, v.AsString())
}

func TestUnicodeCharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"😊"`)
	ass.Equal(t, `"😊"`, v.AsString())

	v = Character.MakeFromString(`"界"`)
	ass.Equal(t, `"界"`, v.AsString())
}

func TestEscapedCharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"\""`)
	ass.Equal(t, `"\""`, v.AsString())

	v = Character.MakeFromString(`"\\"`)
	ass.Equal(t, `"\\"`, v.AsString())

	v = Character.MakeFromString(`"\n"`)
	ass.Equal(t, `"\n"`, v.AsString())

	v = Character.MakeFromString(`"\t"`)
	ass.Equal(t, `"\t"`, v.AsString())
}

func TestCitation(t *tes.T) {
	var Citation = ele.Citation()
	var v1 = Citation.MakeFromString("/bali/types/abstractions/String/v1.2.3")
	ass.Equal(t, "/bali/types/abstractions/String/v1.2.3", v1.AsString())
	ass.Equal(t, "/bali/types/abstractions/String", v1.GetName())
	ass.Equal(t, "v1.2.3", v1.GetVersion())
}

var zero int64
var one int64 = 1
var zilch float64

func TestZeroDurations(t *tes.T) {
	var Duration = ele.Duration()
	var v = Duration.MakeFromMilliseconds(0)
	ass.Equal(t, zero, v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, zilch, v.AsMilliseconds())
	ass.Equal(t, zilch, v.AsSeconds())
	ass.Equal(t, zilch, v.AsMinutes())
	ass.Equal(t, zilch, v.AsHours())
	ass.Equal(t, zilch, v.AsDays())
	ass.Equal(t, zilch, v.AsWeeks())
	ass.Equal(t, zilch, v.AsMonths())
	ass.Equal(t, zilch, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, zero, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

func TestPositiveDurations(t *tes.T) {
	var Duration = ele.Duration()
	var v = Duration.MakeFromMilliseconds(60000)
	ass.Equal(t, int64(60000), v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, one, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

func TestNegativeDurations(t *tes.T) {
	var Duration = ele.Duration()
	var v = Duration.MakeFromMilliseconds(-60000)
	ass.Equal(t, int64(-60000), v.AsInteger())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -60000.0, v.AsMilliseconds())
	ass.Equal(t, -60.0, v.AsSeconds())
	ass.Equal(t, -1.0, v.AsMinutes())
	ass.Equal(t, -0.016666666666666666, v.AsHours())
	ass.Equal(t, -0.0006944444444444445, v.AsDays())
	ass.Equal(t, -9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, -2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, -1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, one, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

func TestInfiniteFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`+infinity`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`+∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`∞`)
	ass.Equal(t, `∞`, v.AsString())

	v = Float.MakeFromString(`-infinity`)
	ass.Equal(t, `-∞`, v.AsString())

	v = Float.MakeFromString(`-∞`)
	ass.Equal(t, `-∞`, v.AsString())
}

func TestZeroFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`12.3`)
	ass.Equal(t, `12.3`, v.AsString())

	v = Float.MakeFromString(`+π`)
	ass.Equal(t, `π`, v.AsString())
}

func TestNegativeFloats(t *tes.T) {
	var Float = ele.Float()
	var v = Float.MakeFromString(`-e`)
	ass.Equal(t, `-e`, v.AsString())
}

func TestZeroIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`0`)
	ass.Equal(t, `0`, v.AsString())
}

func TestPositiveIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`42`)
	ass.Equal(t, `42`, v.AsString())

	v = Integer.MakeFromString(`+5`)
	ass.Equal(t, `5`, v.AsString())
}

func TestNegativeIntegers(t *tes.T) {
	var Integer = ele.Integer()
	var v = Integer.MakeFromString(`-1`)
	ass.Equal(t, `-1`, v.AsString())
}

func TestIntegerMoments(t *tes.T) {
	var Moment = ele.Moment()
	var v = Moment.MakeFromMilliseconds(1238589296789)
	ass.Equal(t, int64(1238589296789), v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391204, v.AsDays())
	ass.Equal(t, 2047.9320383416004, v.AsWeeks())
	ass.Equal(t, 470.9919881193849, v.AsMonths())
	ass.Equal(t, 39.24933234328208, v.AsYears())
	ass.Equal(t, int64(789), v.GetMilliseconds())
	ass.Equal(t, int64(56), v.GetSeconds())
	ass.Equal(t, int64(34), v.GetMinutes())
	ass.Equal(t, int64(12), v.GetHours())
	ass.Equal(t, int64(1), v.GetDays())
	ass.Equal(t, int64(14), v.GetWeeks())
	ass.Equal(t, int64(4), v.GetMonths())
	ass.Equal(t, int64(2009), v.GetYears())
}

func TestMomentsLibrary(t *tes.T) {
	var Duration = ele.Duration()
	var duration = Duration.MakeFromMilliseconds(12345)
	var Moment = ele.Moment()
	var before = Moment.Make()
	var after = Moment.MakeFromMilliseconds(before.AsInteger() + duration.AsInteger())

	ass.Equal(t, duration, Moment.Duration(before, after))
	ass.Equal(t, after, Moment.Later(before, duration))
	ass.Equal(t, before, Moment.Earlier(after, duration))
}
