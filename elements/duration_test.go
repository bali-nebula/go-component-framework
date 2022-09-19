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
	"testing"
)

func TestZeroDurations(t *testing.T) {
	var v = elements.Duration(0)
	assert.Equal(t, 0, int(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 0.0, v.AsMilliseconds())
	assert.Equal(t, 0.0, v.AsSeconds())
	assert.Equal(t, 0.0, v.AsMinutes())
	assert.Equal(t, 0.0, v.AsHours())
	assert.Equal(t, 0.0, v.AsDays())
	assert.Equal(t, 0.0, v.AsWeeks())
	assert.Equal(t, 0.0, v.AsMonths())
	assert.Equal(t, 0.0, v.AsYears())
	assert.Equal(t, "~P0Y0M0DT0H0M0.0S", v.AsString())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 0, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 0, v.GetDays())
	assert.Equal(t, 0, v.GetWeeks())
	assert.Equal(t, 0, v.GetMonths())
	assert.Equal(t, 0, v.GetYears())
}

func TestPositiveDurations(t *testing.T) {
	var v = elements.Duration(60000)
	assert.Equal(t, 60000, int(v))
	assert.False(t, v.IsNegative())
	assert.Equal(t, 60000.0, v.AsMilliseconds())
	assert.Equal(t, 60.0, v.AsSeconds())
	assert.Equal(t, 1.0, v.AsMinutes())
	assert.Equal(t, 0.016666666666666666, v.AsHours())
	assert.Equal(t, 0.0006944444444444445, v.AsDays())
	assert.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	assert.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	assert.Equal(t, 1.901324310408686e-06, v.AsYears())
	assert.Equal(t, "~P0Y0M0DT0H1M0.0S", v.AsString())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 1, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 0, v.GetDays())
	assert.Equal(t, 0, v.GetWeeks())
	assert.Equal(t, 0, v.GetMonths())
	assert.Equal(t, 0, v.GetYears())
}

func TestNegativeDurations(t *testing.T) {
	var v = elements.Duration(-60000)
	assert.Equal(t, -60000, int(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -60000.0, v.AsMilliseconds())
	assert.Equal(t, -60.0, v.AsSeconds())
	assert.Equal(t, -1.0, v.AsMinutes())
	assert.Equal(t, -0.016666666666666666, v.AsHours())
	assert.Equal(t, -0.0006944444444444445, v.AsDays())
	assert.Equal(t, -9.92063492063492e-05, v.AsWeeks())
	assert.Equal(t, -2.2815891724904232e-05, v.AsMonths())
	assert.Equal(t, -1.901324310408686e-06, v.AsYears())
	assert.Equal(t, "~-P0Y0M0DT0H1M0.0S", v.AsString())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 1, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 0, v.GetDays())
	assert.Equal(t, 0, v.GetWeeks())
	assert.Equal(t, 0, v.GetMonths())
	assert.Equal(t, 0, v.GetYears())
}

func TestStringDurations(t *testing.T) {
	var v, ok = elements.DurationFromString("~-P12Y3M4DT5H6M7.890S")
	assert.True(t, ok)
	assert.Equal(t, -386936629890, int(v))
	assert.True(t, v.IsNegative())
	assert.Equal(t, -386936629890.0, v.AsMilliseconds())
	assert.Equal(t, -386936629.89, v.AsSeconds())
	assert.Equal(t, -6448943.831499999, v.AsMinutes())
	assert.Equal(t, -107482.39719166666, v.AsHours())
	assert.Equal(t, -4478.433216319444, v.AsDays())
	assert.Equal(t, -639.7761737599205, v.AsWeeks())
	assert.Equal(t, -147.138404199493, v.AsMonths())
	assert.Equal(t, -12.261533683291084, v.AsYears())
	assert.Equal(t, "~-P12Y3M4DT5H6M7.890S", v.AsString())
	assert.Equal(t, 890, v.GetMilliseconds())
	assert.Equal(t, 7, v.GetSeconds())
	assert.Equal(t, 6, v.GetMinutes())
	assert.Equal(t, 5, v.GetHours())
	assert.Equal(t, 4, v.GetDays())
	assert.Equal(t, 13, v.GetWeeks())
	assert.Equal(t, 3, v.GetMonths())
	assert.Equal(t, 12, v.GetYears())

	v, ok = elements.DurationFromString("P12Y3M4DT5H6M7.890S")
	assert.False(t, ok)
	assert.Equal(t, 0, int(v))
}

func TestDurationsLibrary(t *testing.T) {
	var v0 = elements.Duration(0)
	var v1 = elements.Duration(5)
	var v2 = elements.Duration(-5)
	var v3 = elements.Duration(1)

	assert.Equal(t, v0, elements.Durations.Inverse(v0))
	assert.Equal(t, v2, elements.Durations.Inverse(v1))
	assert.Equal(t, v1, elements.Durations.Inverse(v2))

	assert.Equal(t, v0, elements.Durations.Sum(v0, v0))
	assert.Equal(t, v0, elements.Durations.Difference(v0, v0))
	assert.Equal(t, v1, elements.Durations.Sum(v0, v1))
	assert.Equal(t, v2, elements.Durations.Difference(v0, v1))
	assert.Equal(t, v0, elements.Durations.Sum(v1, v2))
	assert.Equal(t, v2, elements.Durations.Difference(v2, v0))

	assert.Equal(t, v0, elements.Durations.Scaled(v0, math.Pi))
	assert.Equal(t, v2, elements.Durations.Scaled(v1, -1.002))
	assert.Equal(t, v0, elements.Durations.Scaled(v2, 0.002))
	assert.Equal(t, v3, elements.Durations.Scaled(v1, 0.2))
}
