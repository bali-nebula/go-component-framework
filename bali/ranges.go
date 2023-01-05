/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ran "github.com/bali-nebula/go-component-framework/ranges"
	str "github.com/bali-nebula/go-component-framework/strings"
	mat "math"
	ref "reflect"
	stc "strconv"
	utf "unicode/utf8"
)

// This method attempts to parse an endpoint. It returns the endpoint and
// whether or not the endpoint was successfully parsed.
func (v *parser) parseEndpoint() (abs.Primitive, *Token, bool) {
	var ok bool
	var token *Token
	var endpoint abs.Primitive
	endpoint, token, ok = v.parseAngle()
	if !ok {
		endpoint, token, ok = v.parseBinary()
	}
	if !ok {
		endpoint, token, ok = v.parseBoolean()
	}
	if !ok {
		endpoint, token, ok = v.parseDuration()
	}
	if !ok {
		endpoint, token, ok = v.parseMoment()
	}
	if !ok {
		endpoint, token, ok = v.parseMoniker()
	}
	if !ok {
		endpoint, token, ok = v.parsePattern()
	}
	if !ok {
		endpoint, token, ok = v.parsePercentage()
	}
	if !ok {
		endpoint, token, ok = v.parseProbability()
	}
	if !ok {
		endpoint, token, ok = v.parseQuote()
	}
	if !ok {
		endpoint, token, ok = v.parseReal()
	}
	if !ok {
		endpoint, token, ok = v.parseResource()
	}
	if !ok {
		endpoint, token, ok = v.parseRune()
	}
	if !ok {
		endpoint, token, ok = v.parseSymbol()
	}
	if !ok {
		endpoint, token, ok = v.parseTag()
	}
	if !ok {
		endpoint, token, ok = v.parseVersion()
	}
	if !ok {
		// This must be explicitly set to nil since it is of type any.
		endpoint = nil
	}
	return endpoint, token, ok
}

// This method adds the canonical format for the specified endpoint to the
// state of the formatter.
func (v *formatter) formatEndpoint(endpoint abs.Primitive) {
	// NOTE: A type switch will only work if each case specifies a unique
	// interface. If two different interfaces define the same method signatures
	// they are indistinguishable as types. To get around this we have added as
	// necessary a unique "dummy" method to each interface to guarantee that it
	// is unique.
	switch value := endpoint.(type) {
	case ele.Angle:
		v.formatAngle(value)
	case str.Binary:
		v.formatBinary(value)
	case ele.Boolean:
		v.formatBoolean(value)
	case ele.Duration:
		v.formatDuration(value)
	case ele.Moment:
		v.formatMoment(value)
	case str.Moniker:
		v.formatMoniker(value)
	case ele.Pattern:
		v.formatPattern(value)
	case ele.Percentage:
		v.formatPercentage(value)
	case ele.Probability:
		v.formatProbability(value)
	case str.Quote:
		v.formatQuote(value)
	case ran.Real:
		v.formatReal(value)
	case ele.Resource:
		v.formatResource(value)
	case ran.Rune:
		v.formatRune(value)
	case ele.Symbol:
		v.formatSymbol(value)
	case ele.Tag:
		v.formatTag(value)
	case str.Version:
		v.formatVersion(value)
	default:
		panic(fmt.Sprintf("An invalid endpoint (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a range. It returns the range and whether or
// not the range was successfully parsed.
func (v *parser) parseRange() (abs.Range, *Token, bool) {
	var ok bool
	var token *Token
	var left, right string
	var first abs.Primitive
	var extent abs.Extent
	var last abs.Primitive
	var range_ abs.Range
	left, token, ok = v.parseDelimiter("[")
	if !ok {
		left, token, ok = v.parseDelimiter("(")
		if !ok {
			return range_, token, false
		}
	}
	first, token, ok = v.parseEndpoint()
	if !ok {
		// This is not a range.
		v.backupOne() // Put back the left bracket character.
		return range_, token, false
	}
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a range.
		v.backupOne() // Put back the first endpoint token.
		v.backupOne() // Put back the left bracket character.
		return range_, token, false
	}
	last, token, ok = v.parseEndpoint()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("right endpoint",
			"$range",
			"$primitive")
		panic(message)
	}
	right, token, ok = v.parseDelimiter("]")
	if !ok {
		right, token, ok = v.parseDelimiter(")")
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("right bracket",
				"$range",
				"$primitive")
			panic(message)
		}
	}
	switch {
	case left == "[" && right == "]":
		extent = abs.INCLUSIVE
	case left == "(" && right == "]":
		extent = abs.RIGHT
	case left == "[" && right == ")":
		extent = abs.LEFT
	case left == "(" && right == ")":
		extent = abs.EXCLUSIVE
	default:
		// This should never happen unless there is a bug in the parser.
		var message = fmt.Sprintf("Expected valid range brackets but received:%q and %q\n", left, right)
		panic(message)
	}
	if ref.TypeOf(first) != ref.TypeOf(first) {
		var message = fmt.Sprintf("The range endpoints have two different types: %T and %T\n", first, last)
		panic(message)
	}
	switch first.(type) {
	case abs.Discrete:
		range_ = ran.Interval(first.(abs.Discrete), extent, last.(abs.Discrete))
	case abs.Spectral:
		range_ = ran.Spectrum(first.(abs.Spectral), extent, last.(abs.Spectral))
	case abs.Continuous:
		range_ = ran.Continuum(first.(abs.Continuous), extent, last.(abs.Continuous))
	default:
		var message = fmt.Sprintf("An invalid range endpoint (of type %T) was parsed: %v", first, first)
		panic(message)
	}
	return range_, token, true
}

// This method adds the canonical format for the specified interval to the
// state of the formatter.
func (v *formatter) formatInterval(interval abs.IntervalLike[abs.Discrete]) {
	var extent = interval.GetExtent()
	var left, right string
	switch extent {
	case abs.INCLUSIVE:
		left, right = "[", "]"
	case abs.RIGHT:
		left, right = "(", "]"
	case abs.LEFT:
		left, right = "[", ")"
	case abs.EXCLUSIVE:
		left, right = "(", ")"
	default:
		panic(fmt.Sprintf("The range contains an unknown extent type: %v", extent))
	}
	v.AppendString(left)
	var first = interval.GetFirst()
	v.formatEndpoint(first)
	v.AppendString("..")
	var last = interval.GetLast()
	v.formatEndpoint(last)
	v.AppendString(right)
}

// This method adds the canonical format for the specified spectrum to the
// state of the formatter.
func (v *formatter) formatSpectrum(spectrum abs.SpectrumLike[abs.Spectral]) {
	var extent = spectrum.GetExtent()
	var left, right string
	switch extent {
	case abs.INCLUSIVE:
		left, right = "[", "]"
	case abs.RIGHT:
		left, right = "(", "]"
	case abs.LEFT:
		left, right = "[", ")"
	case abs.EXCLUSIVE:
		left, right = "(", ")"
	default:
		panic(fmt.Sprintf("The range contains an unknown extent type: %v", extent))
	}
	v.AppendString(left)
	var first = spectrum.GetFirst()
	v.formatEndpoint(first)
	v.AppendString("..")
	var last = spectrum.GetLast()
	v.formatEndpoint(last)
	v.AppendString(right)
}

// This method adds the canonical format for the specified continuum to the
// state of the formatter.
func (v *formatter) formatContinuum(continuum abs.ContinuumLike[abs.Continuous]) {
	var extent = continuum.GetExtent()
	var left, right string
	switch extent {
	case abs.INCLUSIVE:
		left, right = "[", "]"
	case abs.RIGHT:
		left, right = "(", "]"
	case abs.LEFT:
		left, right = "[", ")"
	case abs.EXCLUSIVE:
		left, right = "(", ")"
	default:
		panic(fmt.Sprintf("The range contains an unknown extent type: %v", extent))
	}
	v.AppendString(left)
	var first = continuum.GetFirst()
	v.formatEndpoint(first)
	v.AppendString("..")
	var last = continuum.GetLast()
	v.formatEndpoint(last)
	v.AppendString(right)
}

// This method attempts to parse a real number. It returns the real number
// and whether or not the real number was successfully parsed.
func (v *parser) parseReal() (ran.Real, *Token, bool) {
	var number ran.Real
	var token = v.nextToken()
	if token.Type != TokenNumber {
		v.backupOne()
		return number, token, false
	}
	var err any
	var r float64
	var matches = scanReal([]byte(token.Value))
	switch {
	case matches[0] == "undefined":
		r = mat.NaN()
	case matches[0] == "+infinity" || matches[0] == "+∞":
		r = mat.Inf(1)
	case matches[0] == "-infinity" || matches[0] == "-∞":
		r = mat.Inf(-1)
	case matches[0] == "pi", matches[0] == "-pi", matches[0] == "phi", matches[0] == "-phi":
		r = stringToFloat(matches[0])
	default:
		r, err = stc.ParseFloat(matches[0], 64)
		if err != nil {
			r = stringToFloat(matches[0])
		}
	}
	number = ran.Real(r)
	return number, token, true
}

// This method adds the canonical format for the specified real number to the
// state of the formatter.
func (v *formatter) formatReal(number ran.Real) {
	switch {
	case number.IsZero():
		v.AppendString("0")
	case number.IsInfinite():
		if number.IsNegative() {
			v.AppendString("-")
		} else {
			v.AppendString("+")
		}
		v.AppendString("∞")
	case number.IsUndefined():
		v.AppendString("undefined")
	default:
		v.formatFloat(float64(number))
	}
}

// This method attempts to parse a rune. It returns the rune and whether or not
// the rune was successfully parsed.
func (v *parser) parseRune() (ran.Rune, *Token, bool) {
	var number = ran.Rune(-1)
	var quote, token, ok = v.parseQuote()
	if !ok {
		return number, token, false
	}
	var s = string(quote)
	var r, size = utf.DecodeRuneInString(s)
	if len(s) != size {
		// This is not a rune.
		v.backupOne() // Put back the quote token.
		return number, token, false
	}
	number = ran.Rune(r)
	return number, token, true
}

// This method adds the canonical format for the specified rune to the state of
// the formatter.
func (v *formatter) formatRune(number ran.Rune) {
	var runes = []rune{rune(number)}
	var quote = str.Quote(string(runes))
	v.formatQuote(quote)
}
