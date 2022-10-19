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
	tes "testing"
)

func TestIntegerMoments(t *tes.T) {
	var v = ele.Moment(1238589296789)
	ass.Equal(t, 1238589296789, int(v))
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

func TestNegativeMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<-1009-08-07T06:05:04.321>")
	ass.True(t, ok)
	ass.Equal(t, "<-1009-08-07T06:05:04.321>", v.AsString())
	ass.Equal(t, -93989267695679, int(v))
	ass.Equal(t, -93989267695679.0, v.AsMilliseconds())
	ass.Equal(t, -93989267695.679, v.AsSeconds())
	ass.Equal(t, -1566487794.9279833, v.AsMinutes())
	ass.Equal(t, -26108129.915466387, v.AsHours())
	ass.Equal(t, -1087838.7464777662, v.AsDays())
	ass.Equal(t, -155405.53521110947, v.AsWeeks())
	ass.Equal(t, -35740.81591746085, v.AsMonths())
	ass.Equal(t, -2978.401326455071, v.AsYears())
	ass.Equal(t, 321, v.GetMilliseconds())
	ass.Equal(t, 4, v.GetSeconds())
	ass.Equal(t, 5, v.GetMinutes())
	ass.Equal(t, 6, v.GetHours())
	ass.Equal(t, 7, v.GetDays())
	ass.Equal(t, 31, v.GetWeeks())
	ass.Equal(t, 8, v.GetMonths())
	ass.Equal(t, -1009, v.GetYears())
}

func TestMillisecondsMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04-01T12:34:56.789>")
	ass.True(t, ok)
	ass.Equal(t, 1238589296789, int(v))
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

func TestSecondsMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04-01T12:34:56>")
	ass.True(t, ok)
	ass.Equal(t, 1238589296000, int(v))
	ass.Equal(t, 1238589296000.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.000, v.AsSeconds())
	ass.Equal(t, 20643154.933333334, v.AsMinutes())
	ass.Equal(t, 344052.58222222223, v.AsHours())
	ass.Equal(t, 14335.52425925926, v.AsDays())
	ass.Equal(t, 2047.9320370370372, v.AsWeeks())
	ass.Equal(t, 470.991987819356, v.AsMonths())
	ass.Equal(t, 39.24933231827966, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 56, v.GetSeconds())
	ass.Equal(t, 34, v.GetMinutes())
	ass.Equal(t, 12, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestMinutesMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04-01T12:34>")
	ass.True(t, ok)
	ass.Equal(t, 1238589240000, int(v))
	ass.Equal(t, 1238589240000.0, v.AsMilliseconds())
	ass.Equal(t, 1238589240.0, v.AsSeconds())
	ass.Equal(t, 20643154.0, v.AsMinutes())
	ass.Equal(t, 344052.56666666665, v.AsHours())
	ass.Equal(t, 14335.52361111111, v.AsDays())
	ass.Equal(t, 2047.9319444444443, v.AsWeeks())
	ass.Equal(t, 470.9919665245236, v.AsMonths())
	ass.Equal(t, 39.2493305437103, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 34, v.GetMinutes())
	ass.Equal(t, 12, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestHoursMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04-01T12>")
	ass.True(t, ok)
	ass.Equal(t, 1238587200000, int(v))
	ass.Equal(t, 1238587200000.0, v.AsMilliseconds())
	ass.Equal(t, 1238587200.0, v.AsSeconds())
	ass.Equal(t, 20643120.0, v.AsMinutes())
	ass.Equal(t, 344052.0, v.AsHours())
	ass.Equal(t, 14335.5, v.AsDays())
	ass.Equal(t, 2047.9285714285713, v.AsWeeks())
	ass.Equal(t, 470.991190784205, v.AsMonths())
	ass.Equal(t, 39.24926589868375, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 12, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestDaysMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04-01>")
	ass.True(t, ok)
	ass.Equal(t, 1238544000000, int(v))
	ass.Equal(t, 1238544000000.0, v.AsMilliseconds())
	ass.Equal(t, 1238544000.0, v.AsSeconds())
	ass.Equal(t, 20642400.0, v.AsMinutes())
	ass.Equal(t, 344040.0, v.AsHours())
	ass.Equal(t, 14335.0, v.AsDays())
	ass.Equal(t, 2047.857142857143, v.AsWeeks())
	ass.Equal(t, 470.9747633421631, v.AsMonths())
	ass.Equal(t, 39.24789694518026, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestMonthsMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009-04>") // Note: same as "<2009-04-01>".
	ass.True(t, ok)
	ass.Equal(t, 1238544000000, int(v))
	ass.Equal(t, 1238544000000.0, v.AsMilliseconds())
	ass.Equal(t, 1238544000.0, v.AsSeconds())
	ass.Equal(t, 20642400.0, v.AsMinutes())
	ass.Equal(t, 344040.0, v.AsHours())
	ass.Equal(t, 14335.0, v.AsDays())
	ass.Equal(t, 2047.857142857143, v.AsWeeks())
	ass.Equal(t, 470.9747633421631, v.AsMonths())
	ass.Equal(t, 39.24789694518026, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestYearsMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<2009>")
	ass.True(t, ok)
	ass.Equal(t, 1230768000000, int(v))
	ass.Equal(t, 1230768000000.0, v.AsMilliseconds())
	ass.Equal(t, 1230768000.0, v.AsSeconds())
	ass.Equal(t, 20512800.0, v.AsMinutes())
	ass.Equal(t, 341880.0, v.AsHours())
	ass.Equal(t, 14245.0, v.AsDays())
	ass.Equal(t, 2035.0, v.AsWeeks())
	ass.Equal(t, 468.01782377461547, v.AsMonths())
	ass.Equal(t, 39.00148531455129, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 1, v.GetWeeks())
	ass.Equal(t, 1, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestBadMoments(t *tes.T) {
	var v, ok = ele.MomentFromString("<>")
	ass.False(t, ok)
	ass.Equal(t, 0, int(v))
}

func TestMomentsLibrary(t *tes.T) {
	var before = ele.Now()
	var duration = ele.Duration(12345)
	var after = ele.Moment(int(before) + int(duration))

	ass.Equal(t, duration, ele.Moments.Duration(before, after))
	ass.Equal(t, after, ele.Moments.Later(before, duration))
	ass.Equal(t, before, ele.Moments.Earlier(after, duration))
}
