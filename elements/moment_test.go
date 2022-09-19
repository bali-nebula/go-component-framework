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
	"testing"
)

func TestIntegerMoments(t *testing.T) {
	var v = elements.Moment(1238589296789)
	assert.Equal(t, 1238589296789, int(v))
	assert.Equal(t, 1238589296789.0, v.AsMilliseconds())
	assert.Equal(t, 1238589296.789, v.AsSeconds())
	assert.Equal(t, 20643154.946483333, v.AsMinutes())
	assert.Equal(t, 344052.58244138886, v.AsHours())
	assert.Equal(t, 14335.524268391202, v.AsDays())
	assert.Equal(t, 2047.9320383416002, v.AsWeeks())
	assert.Equal(t, 470.99198811938487, v.AsMonths())
	assert.Equal(t, 39.24933234328207, v.AsYears())
	assert.Equal(t, 789, v.GetMilliseconds())
	assert.Equal(t, 56, v.GetSeconds())
	assert.Equal(t, 34, v.GetMinutes())
	assert.Equal(t, 12, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestNegativeMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<-1009-08-07T06:05:04.321>")
	assert.True(t, ok)
	assert.Equal(t, "<-1009-08-07T06:05:04.321>", v.AsString())
	assert.Equal(t, -93989267695679, int(v))
	assert.Equal(t, -93989267695679.0, v.AsMilliseconds())
	assert.Equal(t, -93989267695.679, v.AsSeconds())
	assert.Equal(t, -1566487794.9279833, v.AsMinutes())
	assert.Equal(t, -26108129.915466387, v.AsHours())
	assert.Equal(t, -1087838.7464777662, v.AsDays())
	assert.Equal(t, -155405.53521110947, v.AsWeeks())
	assert.Equal(t, -35740.81591746085, v.AsMonths())
	assert.Equal(t, -2978.401326455071, v.AsYears())
	assert.Equal(t, 321, v.GetMilliseconds())
	assert.Equal(t, 4, v.GetSeconds())
	assert.Equal(t, 5, v.GetMinutes())
	assert.Equal(t, 6, v.GetHours())
	assert.Equal(t, 7, v.GetDays())
	assert.Equal(t, 31, v.GetWeeks())
	assert.Equal(t, 8, v.GetMonths())
	assert.Equal(t, -1009, v.GetYears())
}

func TestMillisecondsMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04-01T12:34:56.789>")
	assert.True(t, ok)
	assert.Equal(t, 1238589296789, int(v))
	assert.Equal(t, 1238589296789.0, v.AsMilliseconds())
	assert.Equal(t, 1238589296.789, v.AsSeconds())
	assert.Equal(t, 20643154.946483333, v.AsMinutes())
	assert.Equal(t, 344052.58244138886, v.AsHours())
	assert.Equal(t, 14335.524268391202, v.AsDays())
	assert.Equal(t, 2047.9320383416002, v.AsWeeks())
	assert.Equal(t, 470.99198811938487, v.AsMonths())
	assert.Equal(t, 39.24933234328207, v.AsYears())
	assert.Equal(t, 789, v.GetMilliseconds())
	assert.Equal(t, 56, v.GetSeconds())
	assert.Equal(t, 34, v.GetMinutes())
	assert.Equal(t, 12, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestSecondsMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04-01T12:34:56>")
	assert.True(t, ok)
	assert.Equal(t, 1238589296000, int(v))
	assert.Equal(t, 1238589296000.0, v.AsMilliseconds())
	assert.Equal(t, 1238589296.000, v.AsSeconds())
	assert.Equal(t, 20643154.933333334, v.AsMinutes())
	assert.Equal(t, 344052.58222222223, v.AsHours())
	assert.Equal(t, 14335.52425925926, v.AsDays())
	assert.Equal(t, 2047.9320370370372, v.AsWeeks())
	assert.Equal(t, 470.991987819356, v.AsMonths())
	assert.Equal(t, 39.24933231827966, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 56, v.GetSeconds())
	assert.Equal(t, 34, v.GetMinutes())
	assert.Equal(t, 12, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestMinutesMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04-01T12:34>")
	assert.True(t, ok)
	assert.Equal(t, 1238589240000, int(v))
	assert.Equal(t, 1238589240000.0, v.AsMilliseconds())
	assert.Equal(t, 1238589240.0, v.AsSeconds())
	assert.Equal(t, 20643154.0, v.AsMinutes())
	assert.Equal(t, 344052.56666666665, v.AsHours())
	assert.Equal(t, 14335.52361111111, v.AsDays())
	assert.Equal(t, 2047.9319444444443, v.AsWeeks())
	assert.Equal(t, 470.9919665245236, v.AsMonths())
	assert.Equal(t, 39.2493305437103, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 34, v.GetMinutes())
	assert.Equal(t, 12, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestHoursMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04-01T12>")
	assert.True(t, ok)
	assert.Equal(t, 1238587200000, int(v))
	assert.Equal(t, 1238587200000.0, v.AsMilliseconds())
	assert.Equal(t, 1238587200.0, v.AsSeconds())
	assert.Equal(t, 20643120.0, v.AsMinutes())
	assert.Equal(t, 344052.0, v.AsHours())
	assert.Equal(t, 14335.5, v.AsDays())
	assert.Equal(t, 2047.9285714285713, v.AsWeeks())
	assert.Equal(t, 470.991190784205, v.AsMonths())
	assert.Equal(t, 39.24926589868375, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 0, v.GetMinutes())
	assert.Equal(t, 12, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestDaysMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04-01>")
	assert.True(t, ok)
	assert.Equal(t, 1238544000000, int(v))
	assert.Equal(t, 1238544000000.0, v.AsMilliseconds())
	assert.Equal(t, 1238544000.0, v.AsSeconds())
	assert.Equal(t, 20642400.0, v.AsMinutes())
	assert.Equal(t, 344040.0, v.AsHours())
	assert.Equal(t, 14335.0, v.AsDays())
	assert.Equal(t, 2047.857142857143, v.AsWeeks())
	assert.Equal(t, 470.9747633421631, v.AsMonths())
	assert.Equal(t, 39.24789694518026, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 0, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestMonthsMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009-04>") // Note: same as "<2009-04-01>".
	assert.True(t, ok)
	assert.Equal(t, 1238544000000, int(v))
	assert.Equal(t, 1238544000000.0, v.AsMilliseconds())
	assert.Equal(t, 1238544000.0, v.AsSeconds())
	assert.Equal(t, 20642400.0, v.AsMinutes())
	assert.Equal(t, 344040.0, v.AsHours())
	assert.Equal(t, 14335.0, v.AsDays())
	assert.Equal(t, 2047.857142857143, v.AsWeeks())
	assert.Equal(t, 470.9747633421631, v.AsMonths())
	assert.Equal(t, 39.24789694518026, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 0, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 14, v.GetWeeks())
	assert.Equal(t, 4, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestYearsMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<2009>")
	assert.True(t, ok)
	assert.Equal(t, 1230768000000, int(v))
	assert.Equal(t, 1230768000000.0, v.AsMilliseconds())
	assert.Equal(t, 1230768000.0, v.AsSeconds())
	assert.Equal(t, 20512800.0, v.AsMinutes())
	assert.Equal(t, 341880.0, v.AsHours())
	assert.Equal(t, 14245.0, v.AsDays())
	assert.Equal(t, 2035.0, v.AsWeeks())
	assert.Equal(t, 468.01782377461547, v.AsMonths())
	assert.Equal(t, 39.00148531455129, v.AsYears())
	assert.Equal(t, 0, v.GetMilliseconds())
	assert.Equal(t, 0, v.GetSeconds())
	assert.Equal(t, 0, v.GetMinutes())
	assert.Equal(t, 0, v.GetHours())
	assert.Equal(t, 1, v.GetDays())
	assert.Equal(t, 1, v.GetWeeks())
	assert.Equal(t, 1, v.GetMonths())
	assert.Equal(t, 2009, v.GetYears())
}

func TestBadMoments(t *testing.T) {
	var v, ok = elements.MomentFromString("<>")
	assert.False(t, ok)
	assert.Equal(t, 0, int(v))
}

func TestMomentsLibrary(t *testing.T) {
	var before = elements.Now()
	var duration = elements.Duration(12345)
	var after = elements.Moment(int(before) + int(duration))

	assert.Equal(t, duration, elements.Moments.Duration(before, after))
	assert.Equal(t, after, elements.Moments.Later(before, duration))
	assert.Equal(t, before, elements.Moments.Earlier(after, duration))
}
