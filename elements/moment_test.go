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
	ele "github.com/bali-nebula/go-component-framework/elements"
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

func TestMomentsLibrary(t *tes.T) {
	var before = ele.Now()
	var duration = ele.Duration(12345)
	var after = ele.Moment(int(before) + int(duration))

	ass.Equal(t, duration, ele.Moments.Duration(before, after))
	ass.Equal(t, after, ele.Moments.Later(before, duration))
	ass.Equal(t, before, ele.Moments.Earlier(after, duration))
}
