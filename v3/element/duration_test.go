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
