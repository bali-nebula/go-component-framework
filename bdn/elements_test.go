/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bdn_test

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	bdn "github.com/craterdog-bali/go-bali-document-notation/bdn"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithAngles(t *tes.T) {
	var angleStrings = []string{
		`~0`,
		`~π`,
		`~1.23456789E-10`,
	}

	for index, s := range angleStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, angleStrings[index], s)
	}
}

func TestRoundtripWithBooleans(t *tes.T) {
	var booleanStrings = []string{
		`false`,
		`true`,
	}

	for index, s := range booleanStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, booleanStrings[index], s)
	}
}

func TestRoundtripWithDurations(t *tes.T) {
	var durationStrings = []string{
		`~P0W`,
		`~-P13W`,
		`~P12Y3M4DT5H6M7.890S`,
		`~-P12Y3M4D`,
		`~P3M4DT5H6M`,
		`~-PT6M7.890S`,
		`~P12YT6M7.890S`,
		`~-P12Y4DT5H6M7S`,
	}

	for index, s := range durationStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, durationStrings[index], s)
	}
}

func TestRoundtripWithMoments(t *tes.T) {
	var momentStrings = []string{
		`<0>`,
		`<1>`,
		`<-1>`,
		`<1970>`,
		`<1776>`,
		`<10000>`,
		`<-10000>`,
		`<2009-04>`,
		`<1962-04-25>`,
		`<1-02-03T04:05:06.789>`,
		`<-1-02-03T04:05:06.789>`,
	}

	for index, s := range momentStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, momentStrings[index], s)
	}
}

func TestRoundtripWithNumbers(t *tes.T) {
	var numberStrings = []string{
		`0`,
		`e`,
		`-e`,
		`π`,
		`-π`,
		`i`,
		`-i`,
		`ei`,
		`-ei`,
		`πi`,
		`-πi`,
		`∞`,
		`(1, -i)`,
		`(-3, 4i)`,
	}

	for index, s := range numberStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, numberStrings[index], s)
	}
}

func TestRoundtripWithPatterns(t *tes.T) {
	var patternStrings = []string{
		`none`,
		`"c[aou]+t"?`,
		`any`,
	}

	for index, s := range patternStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, patternStrings[index], s)
	}
}

func TestRoundtripWithPercentages(t *tes.T) {
	var percentageStrings = []string{
		`0%`,
		`50%`,
		`100%`,
		`-1.7%`,
	}

	for index, s := range percentageStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, percentageStrings[index], s)
	}
}

func TestRoundtripWithProbabilities(t *tes.T) {
	var probabilityStrings = []string{
		`.0`,
		`.5`,
		`1.`,
	}

	for index, s := range probabilityStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, probabilityStrings[index], s)
	}
}

func TestRoundtripWithResources(t *tes.T) {
	var resourceStrings = []string{
		`<https://google.com>`,
		`<https://google.com/path>`,
		`<https://google.com/path?foo=bar>`,
		`<https://google.com/path#fragment>`,
		`<https://google.com/path?foo=bar#fragment>`,
	}

	for index, s := range resourceStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, resourceStrings[index], s)
	}
}

func TestRoundtripWithSymbols(t *tes.T) {
	var symbolStrings = []string{
		`$A`,
		`$foobar`,
	}

	for index, s := range symbolStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, symbolStrings[index], s)
	}
}

func TestRoundtripWithTags(t *tes.T) {
	var tagStrings = []string{
		`#A`,
		`#A3GHK57Z`,
	}

	for index, s := range tagStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatValue(component)
		ass.Equal(t, tagStrings[index], s)
	}
}
