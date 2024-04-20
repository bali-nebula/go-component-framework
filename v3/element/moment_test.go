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
	tes "testing"
)

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
