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
	tes "testing"
)

var Moment = ele.Moment()

func TestIntegerMoments(t *tes.T) {
	var v = Moment.FromMilliseconds(1238589296789)
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

func TestMomentsLibrary(t *tes.T) {
	var before = Moment.Now()
	var duration = Duration.FromMilliseconds(12345)
	var after = Moment.FromMilliseconds(before.AsInteger() + duration.AsInteger())

	ass.Equal(t, duration, Moment.Duration(before, after))
	ass.Equal(t, after, Moment.Later(before, duration))
	ass.Equal(t, before, Moment.Earlier(after, duration))
}
