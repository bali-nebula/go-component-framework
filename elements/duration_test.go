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
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

func TestZeroDurations(t *tes.T) {
	var v = ele.Duration(0)
	ass.Equal(t, 0, int(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsMilliseconds())
	ass.Equal(t, 0.0, v.AsSeconds())
	ass.Equal(t, 0.0, v.AsMinutes())
	ass.Equal(t, 0.0, v.AsHours())
	ass.Equal(t, 0.0, v.AsDays())
	ass.Equal(t, 0.0, v.AsWeeks())
	ass.Equal(t, 0.0, v.AsMonths())
	ass.Equal(t, 0.0, v.AsYears())
	ass.Equal(t, "~P0Y0M0DT0H0M0.0S", v.AsString())
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
	var v = ele.Duration(60000)
	ass.Equal(t, 60000, int(v))
	ass.False(t, v.IsNegative())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.901324310408686e-06, v.AsYears())
	ass.Equal(t, "~P0Y0M0DT0H1M0.0S", v.AsString())
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
	var v = ele.Duration(-60000)
	ass.Equal(t, -60000, int(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -60000.0, v.AsMilliseconds())
	ass.Equal(t, -60.0, v.AsSeconds())
	ass.Equal(t, -1.0, v.AsMinutes())
	ass.Equal(t, -0.016666666666666666, v.AsHours())
	ass.Equal(t, -0.0006944444444444445, v.AsDays())
	ass.Equal(t, -9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, -2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, -1.901324310408686e-06, v.AsYears())
	ass.Equal(t, "~-P0Y0M0DT0H1M0.0S", v.AsString())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 1, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestStringDurations(t *tes.T) {
	var v, ok = ele.DurationFromString("~-P12Y3M4DT5H6M7.890S")
	ass.True(t, ok)
	ass.Equal(t, -386936629890, int(v))
	ass.True(t, v.IsNegative())
	ass.Equal(t, -386936629890.0, v.AsMilliseconds())
	ass.Equal(t, -386936629.89, v.AsSeconds())
	ass.Equal(t, -6448943.831499999, v.AsMinutes())
	ass.Equal(t, -107482.39719166666, v.AsHours())
	ass.Equal(t, -4478.433216319444, v.AsDays())
	ass.Equal(t, -639.7761737599205, v.AsWeeks())
	ass.Equal(t, -147.138404199493, v.AsMonths())
	ass.Equal(t, -12.261533683291084, v.AsYears())
	ass.Equal(t, "~-P12Y3M4DT5H6M7.890S", v.AsString())
	ass.Equal(t, 890, v.GetMilliseconds())
	ass.Equal(t, 7, v.GetSeconds())
	ass.Equal(t, 6, v.GetMinutes())
	ass.Equal(t, 5, v.GetHours())
	ass.Equal(t, 4, v.GetDays())
	ass.Equal(t, 13, v.GetWeeks())
	ass.Equal(t, 3, v.GetMonths())
	ass.Equal(t, 12, v.GetYears())

	v, ok = ele.DurationFromString("P12Y3M4DT5H6M7.890S")
	ass.False(t, ok)
	ass.Equal(t, 0, int(v))
}

func TestDurationsLibrary(t *tes.T) {
	var v0 = ele.Duration(0)
	var v1 = ele.Duration(5)
	var v2 = ele.Duration(-5)
	var v3 = ele.Duration(1)

	ass.Equal(t, v0, ele.Durations.Inverse(v0))
	ass.Equal(t, v2, ele.Durations.Inverse(v1))
	ass.Equal(t, v1, ele.Durations.Inverse(v2))

	ass.Equal(t, v0, ele.Durations.Sum(v0, v0))
	ass.Equal(t, v0, ele.Durations.Difference(v0, v0))
	ass.Equal(t, v1, ele.Durations.Sum(v0, v1))
	ass.Equal(t, v2, ele.Durations.Difference(v0, v1))
	ass.Equal(t, v0, ele.Durations.Sum(v1, v2))
	ass.Equal(t, v2, ele.Durations.Difference(v2, v0))

	ass.Equal(t, v0, ele.Durations.Scaled(v0, mat.Pi))
	ass.Equal(t, v2, ele.Durations.Scaled(v1, -1.002))
	ass.Equal(t, v0, ele.Durations.Scaled(v2, 0.002))
	ass.Equal(t, v3, ele.Durations.Scaled(v1, 0.2))
}
