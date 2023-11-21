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

func TestZeroDurations(t *tes.T) {
	var v = ele.Duration.FromMilliseconds(0)
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
	var v = ele.Duration.FromMilliseconds(60000)
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
	var v = ele.Duration.FromMilliseconds(-60000)
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
