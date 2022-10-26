/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language_test

import (
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

func TestParserWithAngles(t *tes.T) {
	var angleStrings = []string{
		`~0`,
		`~π`,
		`~1.23456789E-10`,
	}

	var angleValues = []ele.Angle{
		0,
		mat.Pi,
		1.23456789e-10,
	}

	for index, s := range angleStrings {
		var angle = lan.ParseSource(s).GetEntity().(ele.Angle)
		ass.Equal(t, angleValues[index], angle)
	}

	for index, angle := range angleValues {
		var s = lan.FormatValue(angle)
		ass.Equal(t, angleStrings[index], s)
	}
}

func TestParserWithBooleans(t *tes.T) {
	var booleanStrings = []string{
		`false`,
		`true`,
	}

	var booleanValues = []ele.Boolean{
		false,
		true,
	}

	for index, s := range booleanStrings {
		var boolean = lan.ParseSource(s).GetEntity().(ele.Boolean)
		ass.Equal(t, booleanValues[index], boolean)
	}

	for index, boolean := range booleanValues {
		var s = lan.FormatValue(boolean)
		ass.Equal(t, booleanStrings[index], s)
	}
}

func TestParserWithDurations(t *tes.T) {
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

	var durationValues = []ele.Duration{
		0,
		-7862400000,
		386936629890,
		-386918262000,
		8253198000,
		-367890,
		378683791890,
		-379047391000,
	}

	for index, s := range durationStrings {
		var duration = lan.ParseSource(s).GetEntity().(ele.Duration)
		ass.Equal(t, durationValues[index], duration)
	}

	for index, duration := range durationValues {
		var s = lan.FormatValue(duration)
		ass.Equal(t, durationStrings[index], s)
	}
}

func TestParserWithMoments(t *tes.T) {
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

	var momentValues = []ele.Moment{
		-62167219200000,
		-62135596800000,
		-62198755200000,
		0,
		-6122044800000,
		253402300800000,
		-377736739200000,
		1238544000000,
		-242611200000,
		-62132730893211,
		-62195889293211,
	}

	for index, s := range momentStrings {
		var moment = lan.ParseSource(s).GetEntity().(ele.Moment)
		ass.Equal(t, momentValues[index], moment)
	}

	for index, moment := range momentValues {
		var s = lan.FormatValue(moment)
		ass.Equal(t, momentStrings[index], s)
	}
}

func TestParserWithNumbers(t *tes.T) {
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

	var numberValues = []ele.Number{
		0,
		mat.E,
		-mat.E,
		mat.Pi,
		-mat.Pi,
		complex(0, 1),
		complex(0, -1),
		complex(0, mat.E),
		complex(0, -mat.E),
		complex(0, mat.Pi),
		complex(0, -mat.Pi),
		ele.Number(cmp.Inf()),
		complex(1, -1),
		complex(-3, 4),
	}

	for index, s := range numberStrings {
		var number = lan.ParseSource(s).GetEntity().(ele.Number)
		ass.Equal(t, numberValues[index], number)
	}

	for index, number := range numberValues {
		var s = lan.FormatValue(number)
		ass.Equal(t, numberStrings[index], s)
	}
}

func TestParserWithPatterns(t *tes.T) {
	var patternStrings = []string{
		`none`,
		`"c[aou]+t"?`,
		`any`,
	}

	var patternValues = []ele.Pattern{
		"^none$",
		`c[aou]+t`,
		".*",
	}

	for index, s := range patternStrings {
		var pattern = lan.ParseSource(s).GetEntity().(ele.Pattern)
		ass.Equal(t, patternValues[index], pattern)
	}

	for index, pattern := range patternValues {
		var s = lan.FormatValue(pattern)
		ass.Equal(t, patternStrings[index], s)
	}
}

func TestParserWithPercentages(t *tes.T) {
	var percentageStrings = []string{
		`0%`,
		`50%`,
		`100%`,
		`-1.7%`,
	}

	var percentageValues = []ele.Percentage{
		0,
		50,
		100,
		-1.7,
	}

	for index, s := range percentageStrings {
		var percentage = lan.ParseSource(s).GetEntity().(ele.Percentage)
		ass.Equal(t, percentageValues[index], percentage)
	}

	for index, percentage := range percentageValues {
		var s = lan.FormatValue(percentage)
		ass.Equal(t, percentageStrings[index], s)
	}
}

func TestParserWithProbabilities(t *tes.T) {
	var probabilityStrings = []string{
		`.0`,
		`.5`,
		`1.`,
	}

	var probabilityValues = []ele.Probability{
		0,
		0.5,
		1,
	}

	for index, s := range probabilityStrings {
		var probability = lan.ParseSource(s).GetEntity().(ele.Probability)
		ass.Equal(t, probabilityValues[index], probability)
	}

	for index, probability := range probabilityValues {
		var s = lan.FormatValue(probability)
		ass.Equal(t, probabilityStrings[index], s)
	}
}

func TestParserWithResources(t *tes.T) {
	var resourceStrings = []string{
		`<https://google.com>`,
		`<https://google.com/path>`,
		`<https://google.com/path?foo=bar>`,
		`<https://google.com/path#fragment>`,
		`<https://google.com/path?foo=bar#fragment>`,
	}

	var resourceValues = []ele.Resource{
		"https://google.com",
		"https://google.com/path",
		"https://google.com/path?foo=bar",
		"https://google.com/path#fragment",
		"https://google.com/path?foo=bar#fragment",
	}

	for index, s := range resourceStrings {
		var resource = lan.ParseSource(s).GetEntity().(ele.Resource)
		ass.Equal(t, resourceValues[index], resource)
	}

	for index, resource := range resourceValues {
		var s = lan.FormatValue(resource)
		ass.Equal(t, resourceStrings[index], s)
	}
}

func TestParserWithSymbols(t *tes.T) {
	var symbolStrings = []string{
		`$A`,
		`$foobar`,
	}

	var symbolValues = []ele.Symbol{
		"A",
		"foobar",
	}

	for index, s := range symbolStrings {
		var symbol = lan.ParseSource(s).GetEntity().(ele.Symbol)
		ass.Equal(t, symbolValues[index], symbol)
	}

	for index, symbol := range symbolValues {
		var s = lan.FormatValue(symbol)
		ass.Equal(t, symbolStrings[index], s)
	}
}

func TestParserWithTags(t *tes.T) {
	var tagStrings = []string{
		`#A`,
		`#A3GHK57Z`,
	}

	var tagValues = []ele.Tag{
		`A`,
		`A3GHK57Z`,
	}

	for index, s := range tagStrings {
		var tag = lan.ParseSource(s).GetEntity().(ele.Tag)
		ass.Equal(t, tagValues[index], tag)
	}

	for index, tag := range tagValues {
		var s = lan.FormatValue(tag)
		ass.Equal(t, tagStrings[index], s)
	}
}
