/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

func TestZeroAngles(t *tes.T) {
	var v = bal.Angle(0)
	ass.Equal(t, 0.0, v.AsReal())
}

func TestPositiveAngles(t *tes.T) {
	var v = bal.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsReal())
}

func TestNegativeAngles(t *tes.T) {
	var v1 = bal.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v1.AsReal())

	var v2 = bal.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v2.AsReal())
}

func TestStringBooleans(t *tes.T) {
	var v1 = bal.Boolean("false")
	ass.False(t, v1.AsBoolean())

	var v2 = bal.Boolean("true")
	ass.True(t, v2.AsBoolean())
}

func TestFalseBooleans(t *tes.T) {
	var v = bal.Boolean(false)
	ass.False(t, v.AsBoolean())
}

func TestTrueBooleans(t *tes.T) {
	var v = bal.Boolean(true)
	ass.True(t, v.AsBoolean())
}

func TestZeroDurations(t *tes.T) {
	var v = bal.Duration("~P0W")
	ass.Equal(t, 0, v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsMilliseconds())
	ass.Equal(t, 0.0, v.AsSeconds())
	ass.Equal(t, 0.0, v.AsMinutes())
	ass.Equal(t, 0.0, v.AsHours())
	ass.Equal(t, 0.0, v.AsDays())
	ass.Equal(t, 0.0, v.AsWeeks())
	ass.Equal(t, 0.0, v.AsMonths())
	ass.Equal(t, 0.0, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestPositiveDurations(t *tes.T) {
	var v = bal.Duration(60000)
	ass.Equal(t, 60000, v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.901324310408686e-06, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 1, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestNegativeDurations(t *tes.T) {
	var v = bal.Duration(-60000)
	ass.Equal(t, -60000, v.AsInteger())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -60000.0, v.AsMilliseconds())
	ass.Equal(t, -60.0, v.AsSeconds())
	ass.Equal(t, -1.0, v.AsMinutes())
	ass.Equal(t, -0.016666666666666666, v.AsHours())
	ass.Equal(t, -0.0006944444444444445, v.AsDays())
	ass.Equal(t, -9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, -2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, -1.901324310408686e-06, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 1, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestStringMoments(t *tes.T) {
	var v = bal.Moment("<2009-04-01T15:30:27.123>")
	ass.Equal(t, 1238599827123, v.AsInteger())
	ass.Equal(t, 1238599827123.0, v.AsMilliseconds())
	ass.Equal(t, 1238599827.123, v.AsSeconds())
	ass.Equal(t, 20643330.452049997, v.AsMinutes())
	ass.Equal(t, 344055.5075341666, v.AsHours())
	ass.Equal(t, 14335.646147256943, v.AsDays())
	ass.Equal(t, 2047.9494496081347, v.AsWeeks())
	ass.Equal(t, 470.99599243539103, v.AsMonths())
	ass.Equal(t, 39.24966603628258, v.AsYears())
	ass.Equal(t, 123, v.GetMilliseconds())
	ass.Equal(t, 27, v.GetSeconds())
	ass.Equal(t, 30, v.GetMinutes())
	ass.Equal(t, 15, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestIntegerMoments(t *tes.T) {
	var v = bal.Moment(1238589296789)
	ass.Equal(t, 1238589296789, v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391202, v.AsDays())
	ass.Equal(t, 2047.9320383416002, v.AsWeeks())
	ass.Equal(t, 470.99198811938487, v.AsMonths())
	ass.Equal(t, 39.24933234328207, v.AsYears())
	ass.Equal(t, 789, v.GetMilliseconds())
	ass.Equal(t, 56, v.GetSeconds())
	ass.Equal(t, 34, v.GetMinutes())
	ass.Equal(t, 12, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestZero(t *tes.T) {
	var v = bal.Number(0 + 0i)
	ass.Equal(t, 0+0i, v.AsComplex())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestInfinity(t *tes.T) {
	var v = bal.Number(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.AsComplex())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, mat.Inf(1), v.AsReal())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
}

func TestUndefined(t *tes.T) {
	var v = bal.Number(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.AsComplex()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsReal()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositiveReals(t *tes.T) {
	var v = bal.Number(0.25)
	ass.Equal(t, 0.25+0i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsReal())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestPositiveImaginaries(t *tes.T) {
	var v = bal.Number(0.25i)
	ass.Equal(t, 0+0.25i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativeReals(t *tes.T) {
	var v = bal.Number(-0.75)
	ass.Equal(t, -0.75+0i, v.AsComplex())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsReal())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativeImaginaries(t *tes.T) {
	var v = bal.Number(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsComplex())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsReal())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNonePattern(t *tes.T) {
	var v = bal.Pattern("none")
	ass.Equal(t, `^none$`, v.AsString())

	var text = ""
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *tes.T) {
	var v = bal.Pattern("any")
	ass.Equal(t, `.*`, v.AsString())

	var text = ""
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *tes.T) {
	var v = bal.Pattern(`"c(.+t)"?`)
	ass.Equal(t, `c(.+t)`, v.AsString())

	var text = "ct"
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}

func TestStringPercentages(t *tes.T) {
	var v = bal.Percentage("100%")
	ass.Equal(t, 1.0, v.AsReal())
}

func TestZeroPercentages(t *tes.T) {
	var v = bal.Percentage(0.0)
	ass.Equal(t, 0.0, v.AsReal())
}

func TestPositivePercentages(t *tes.T) {
	var v = bal.Percentage(25)
	ass.Equal(t, 25, v.AsInteger())
	ass.Equal(t, 0.25, v.AsReal())
}

func TestNegativePercentages(t *tes.T) {
	var v = bal.Percentage(-75)
	ass.Equal(t, -75, v.AsInteger())
	ass.Equal(t, -0.75, v.AsReal())
}

func TestStringProbabilities(t *tes.T) {
	var v1 = bal.Probability(".0")
	ass.False(t, v1.AsBoolean())

	var v2 = bal.Probability("1.")
	ass.True(t, v2.AsBoolean())
}

func TestBooleanProbabilities(t *tes.T) {
	var v1 = bal.Probability(false)
	ass.Equal(t, 0.0, v1.AsReal())

	var v2 = bal.Probability(true)
	ass.Equal(t, 1.0, v2.AsReal())
}

func TestZeroProbabilities(t *tes.T) {
	var v = bal.Probability(0.0)
	ass.Equal(t, 0.0, v.AsReal())
}

func TestOneProbabilities(t *tes.T) {
	var v = bal.Probability(1.0)
	ass.Equal(t, 1.0, v.AsReal())
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = bal.Probability(0.25)
	ass.Equal(t, 0.25, v1.AsReal())

	var v2 = bal.Probability(0.5)
	ass.Equal(t, 0.5, v2.AsReal())

	var v3 = bal.Probability(0.75)
	ass.Equal(t, 0.75, v3.AsReal())
}

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = bal.Resource("<https://craterdog.com/About.html>")
	ass.Equal(t, "https://craterdog.com/About.html", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = bal.Resource("<mailto:craterdog@google.com>")
	ass.Equal(t, "mailto:craterdog@google.com", v.AsString())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = bal.Resource("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.Equal(t, "https://craterdog.com/?foo=bar;bar=baz", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = bal.Resource("<https://craterdog.com/#Home>")
	ass.Equal(t, "https://craterdog.com/#Home", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = bal.Resource("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.Equal(t, "https://craterdog.com/?foo=bar;bar=baz#Home", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestSymbol(t *tes.T) {
	var v = bal.Symbol("$foobar")
	ass.Equal(t, "foobar", v.AsString())
}

func TestTags(t *tes.T) {
	for i := 1; i < 33; i++ {
		var t1 = bal.Tag(i)
		ass.Equal(t, len(t1.AsString()), int(mat.Ceil(float64(i)*8.0/5.0)))
		var s1 = "#" + t1.AsString()
		var t2 = bal.Tag(s1)
		ass.Equal(t, t1, t2)
		var s2 = "#" + t2.AsString()
		ass.Equal(t, s1, s2)
	}
}
