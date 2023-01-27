/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	bal "github.com/bali-nebula/go-component-framework/bali"
	tes "testing"
)

var firstAngle, lastAngle = bal.Angle("~0"), bal.Angle("~π")
var firstBinary, lastBinary = bal.Binary("'abcd'"), bal.Binary("'abce'")
var firstDuration, lastDuration = bal.Duration(12), bal.Duration(123)
var firstFloat, lastFloat = -1.2, 3.4
var firstInt, lastInt = 3, 7
var firstInteger, lastInteger = bal.Integer(-1), bal.Integer(1)
var firstMoment, lastMoment = bal.Moment(-12), bal.Moment(123)
var firstMoniker, lastMoniker = bal.Moniker("/bali/foo"), bal.Moniker("/bali/fuz")
var firstPattern, lastPattern = bal.Pattern(`"ca+d"?`), bal.Pattern(`"ca+t"?`)
var firstPercentage, lastPercentage = bal.Percentage(-10), bal.Percentage(150)
var firstProbability, lastProbability = bal.Probability(false), bal.Probability(true)
var firstQuote, lastQuote = bal.Quote(`"High Five"`), bal.Quote(`"High Seven"`)
var firstReal, lastReal = bal.Real(0), bal.Real(1.23456)
var firstResource, lastResource = bal.Resource("<https://craterdog.com>"), bal.Resource("<https://google.com>")
var firstString, lastString = "aa", "fe"
var firstSymbol, lastSymbol = bal.Symbol("$bar"), bal.Symbol("$foo")
var firstTag, lastTag = bal.Tag("#TAB"), bal.Tag("#TAG")
var firstVersion, lastVersion = bal.Version("v1.2"), bal.Version("v1.2.3")

func TestSpectra(t *tes.T) {
	bal.Spectrum(firstBinary, abs.INCLUSIVE, lastBinary, nil)
	bal.Spectrum(firstMoniker, abs.INCLUSIVE, lastMoniker, nil)
	bal.Spectrum(firstPattern, abs.INCLUSIVE, lastPattern, nil)
	bal.Spectrum(firstQuote, abs.INCLUSIVE, lastQuote, nil)
	bal.Spectrum(firstResource, abs.INCLUSIVE, lastResource, nil)
	bal.Spectrum(firstString, abs.INCLUSIVE, lastString, nil)
	bal.Spectrum(firstSymbol, abs.INCLUSIVE, lastSymbol, nil)
	bal.Spectrum(firstTag, abs.INCLUSIVE, lastTag, nil)
	bal.Spectrum(firstVersion, abs.INCLUSIVE, lastVersion, nil)
}

func TestIntervals(t *tes.T) {
	bal.Interval(firstDuration, abs.INCLUSIVE, lastDuration, nil)
	bal.Interval(firstFloat, abs.INCLUSIVE, lastFloat, nil)
	bal.Interval(firstInt, abs.INCLUSIVE, lastInt, nil)
	bal.Interval(firstInteger, abs.INCLUSIVE, lastInteger, nil)
	bal.Interval(firstMoment, abs.INCLUSIVE, lastMoment, nil)
}

func TestContinua(t *tes.T) {
	bal.Continuum(firstAngle, abs.INCLUSIVE, lastAngle, nil)
	bal.Continuum(firstFloat, abs.INCLUSIVE, lastFloat, nil)
	bal.Continuum(firstInt, abs.INCLUSIVE, lastInt, nil)
	bal.Continuum(firstPercentage, abs.INCLUSIVE, lastPercentage, nil)
	bal.Continuum(firstProbability, abs.INCLUSIVE, lastProbability, nil)
	bal.Continuum(firstReal, abs.INCLUSIVE, lastReal, nil)
}
