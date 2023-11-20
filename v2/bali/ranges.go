/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ran "github.com/bali-nebula/go-component-framework/v2/ranges"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	mat "math"
	ref "reflect"
	stc "strconv"
	utf "unicode/utf8"
)

// UNIVERSAL CONSTRUCTORS

// This constructor returns a new rune endpoint with the specified value.
func Character(value abs.Value) abs.CharacterLike {
	var character abs.CharacterLike
	switch actual := value.(type) {
	case uint:
		character = ele.Character(int(actual))
	case uint8:
		character = ele.Character(int(actual))
	case uint16:
		character = ele.Character(int(actual))
	case uint32:
		character = ele.Character(int(actual))
	case uint64:
		character = ele.Character(int(actual))
	case int:
		character = ele.Character(int(actual))
	case int8:
		character = ele.Character(int(actual))
	case int16:
		character = ele.Character(int(actual))
	case int32:
		character = ele.Character(int(actual))
	case int64:
		character = ele.Character(int(actual))
	case abs.CharacterLike:
		character = actual
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a rune: %v", actual, actual)
		panic(message)
	}
	return character
}

// This constructor returns a new integer endpoint with the specified value.
func Integer(value abs.Value) abs.IntegerLike {
	var integer abs.IntegerLike
	switch actual := value.(type) {
	case uint:
		integer = ele.Integer(int(actual))
	case uint8:
		integer = ele.Integer(int(actual))
	case uint16:
		integer = ele.Integer(int(actual))
	case uint32:
		integer = ele.Integer(int(actual))
	case uint64:
		integer = ele.Integer(int(actual))
	case int:
		integer = ele.Integer(int(actual))
	case int8:
		integer = ele.Integer(int(actual))
	case int16:
		integer = ele.Integer(int(actual))
	case int32:
		integer = ele.Integer(int(actual))
	case int64:
		integer = ele.Integer(int(actual))
	case abs.IntegerLike:
		integer = actual
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to an integer: %v", actual, actual)
		panic(message)
	}
	return integer
}

// This constructor returns a new real endpoint with the specified value.
func Float(value abs.Value) abs.FloatLike {
	var real_ abs.FloatLike
	switch actual := value.(type) {
	case uint:
		real_ = ele.Float(float64(actual))
	case uint8:
		real_ = ele.Float(float64(actual))
	case uint16:
		real_ = ele.Float(float64(actual))
	case uint32:
		real_ = ele.Float(float64(actual))
	case uint64:
		real_ = ele.Float(float64(actual))
	case int:
		real_ = ele.Float(float64(actual))
	case int8:
		real_ = ele.Float(float64(actual))
	case int16:
		real_ = ele.Float(float64(actual))
	case int32:
		real_ = ele.Float(float64(actual))
	case int64:
		real_ = ele.Float(float64(actual))
	case float32:
		real_ = ele.Float(float64(actual))
	case float64:
		real_ = ele.Float(float64(actual))
	case abs.FloatLike:
		real_ = actual
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a real: %v", actual, actual)
		panic(message)
	}
	return real_
}

// This constructor returns a new lexical range with the specified endpoints.
func Spectrum(first abs.Value, extent abs.Extent, last abs.Value, context abs.ContextLike) abs.ComponentLike {
	var entity abs.Entity
	switch actual := first.(type) {
	case string:
		var actualFirst = str.QuoteFromString(first.(string))
		var actualLast = str.QuoteFromString(last.(string))
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.NameLike:
		var actualFirst = first.(abs.NameLike)
		var actualLast = last.(abs.NameLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.PatternLike:
		var actualFirst = first.(abs.PatternLike)
		var actualLast = last.(abs.PatternLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.QuoteLike:
		var actualFirst = first.(abs.QuoteLike)
		var actualLast = last.(abs.QuoteLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.ResourceLike:
		var actualFirst = first.(abs.ResourceLike)
		var actualLast = last.(abs.ResourceLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.SymbolLike:
		var actualFirst = first.(abs.SymbolLike)
		var actualLast = last.(abs.SymbolLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.TagLike:
		var actualFirst = first.(abs.TagLike)
		var actualLast = last.(abs.TagLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	case abs.VersionLike:
		var actualFirst = first.(abs.VersionLike)
		var actualLast = last.(abs.VersionLike)
		entity = ran.Spectrum(actualFirst, extent, actualLast)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be a spectrum endpoint: %v", actual, actual)
		panic(message)
	}
	return ComponentWithContext(entity, context)
}

// This constructor returns a new continuous range with the specified endpoints.
// This constructor returns a new discrete range with the specified endpoints.
func Interval(first abs.Value, extent abs.Extent, last abs.Value, context abs.ContextLike) abs.ComponentLike {
	var entity abs.Entity
	switch actual := first.(type) {
	case uint:
		var actualFirst = ele.Integer(int(first.(uint)))
		var actualLast = ele.Integer(int(last.(uint)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case uint8:
		var actualFirst = ele.Integer(int(first.(uint8)))
		var actualLast = ele.Integer(int(last.(uint8)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case uint16:
		var actualFirst = ele.Integer(int(first.(uint16)))
		var actualLast = ele.Integer(int(last.(uint16)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case uint32:
		var actualFirst = ele.Integer(int(first.(uint32)))
		var actualLast = ele.Integer(int(last.(uint32)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case uint64:
		var actualFirst = ele.Integer(int(first.(uint64)))
		var actualLast = ele.Integer(int(last.(uint64)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case int:
		var actualFirst = ele.Integer(int(first.(int)))
		var actualLast = ele.Integer(int(last.(int)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case int8:
		var actualFirst = ele.Integer(int(first.(int8)))
		var actualLast = ele.Integer(int(last.(int8)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case int16:
		var actualFirst = ele.Integer(int(first.(int16)))
		var actualLast = ele.Integer(int(last.(int16)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case int32:
		var actualFirst = ele.Integer(int(first.(int32)))
		var actualLast = ele.Integer(int(last.(int32)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case int64:
		var actualFirst = ele.Integer(int(first.(int64)))
		var actualLast = ele.Integer(int(last.(int64)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case float32:
		var actualFirst = ele.Integer(int(first.(float32)))
		var actualLast = ele.Integer(int(last.(float32)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case float64:
		var actualFirst = ele.Integer(int(first.(float64)))
		var actualLast = ele.Integer(int(last.(float64)))
		entity = ran.Interval(actualFirst, extent, actualLast)
	case abs.DurationLike:
		var actualFirst = first.(abs.DurationLike)
		var actualLast = last.(abs.DurationLike)
		entity = ran.Interval(actualFirst, extent, actualLast)
	case abs.MomentLike:
		var actualFirst = first.(abs.MomentLike)
		var actualLast = last.(abs.MomentLike)
		entity = ran.Interval(actualFirst, extent, actualLast)
	case abs.IntegerLike:
		var actualFirst = first.(abs.IntegerLike)
		var actualLast = last.(abs.IntegerLike)
		entity = ran.Interval(actualFirst, extent, actualLast)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be an interval endpoint: %v", actual, actual)
		panic(message)
	}
	return ComponentWithContext(entity, context)
}

// This constructor returns a new continuous range with the specified endpoints.
func Continuum(first abs.Value, extent abs.Extent, last abs.Value, context abs.ContextLike) abs.ComponentLike {
	var entity abs.Entity
	switch actual := first.(type) {
	case uint:
		var actualFirst = ele.Float(float64(first.(uint)))
		var actualLast = ele.Float(float64(last.(uint)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case uint8:
		var actualFirst = ele.Float(float64(first.(uint8)))
		var actualLast = ele.Float(float64(last.(uint8)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case uint16:
		var actualFirst = ele.Float(float64(first.(uint16)))
		var actualLast = ele.Float(float64(last.(uint16)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case uint32:
		var actualFirst = ele.Float(float64(first.(uint32)))
		var actualLast = ele.Float(float64(last.(uint32)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case uint64:
		var actualFirst = ele.Float(float64(first.(uint64)))
		var actualLast = ele.Float(float64(last.(uint64)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case int:
		var actualFirst = ele.Float(float64(first.(int)))
		var actualLast = ele.Float(float64(last.(int)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case int8:
		var actualFirst = ele.Float(float64(first.(int8)))
		var actualLast = ele.Float(float64(last.(int8)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case int16:
		var actualFirst = ele.Float(float64(first.(int16)))
		var actualLast = ele.Float(float64(last.(int16)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case int32:
		var actualFirst = ele.Float(float64(first.(int32)))
		var actualLast = ele.Float(float64(last.(int32)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case int64:
		var actualFirst = ele.Float(float64(first.(int64)))
		var actualLast = ele.Float(float64(last.(int64)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case float32:
		var actualFirst = ele.Float(float64(first.(float32)))
		var actualLast = ele.Float(float64(last.(float32)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case float64:
		var actualFirst = ele.Float(float64(first.(float64)))
		var actualLast = ele.Float(float64(last.(float64)))
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case abs.PercentageLike:
		var actualFirst = first.(abs.PercentageLike)
		var actualLast = last.(abs.PercentageLike)
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case abs.ProbabilityLike:
		var actualFirst = first.(abs.ProbabilityLike)
		var actualLast = last.(abs.ProbabilityLike)
		entity = ran.Continuum(actualFirst, extent, actualLast)
	case abs.AngleLike:
		var actualFirst = first.(abs.AngleLike)
		var actualLast = last.(abs.AngleLike)
		entity = ran.Continuum(actualFirst, extent, actualLast)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be a continuum endpoint: %v", actual, actual)
		panic(message)
	}
	return ComponentWithContext(entity, context)
}

// This method attempts to parse an endpoint. It returns the endpoint and
// whether or not the endpoint was successfully parsed.
func (v *parser) parseEndpoint() (abs.Primitive, *Token, bool) {
	var ok bool
	var token *Token
	var endpoint abs.Primitive
	endpoint, token, ok = v.parseAngle()
	if !ok {
		endpoint, token, ok = v.parseDuration()
	}
	if !ok {
		endpoint, token, ok = v.parseMoment()
	}
	if !ok {
		endpoint, token, ok = v.parseName()
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
		endpoint, token, ok = v.parseFloat()
	}
	if !ok {
		endpoint, token, ok = v.parseInteger()
	}
	if !ok {
		endpoint, token, ok = v.parseResource()
	}
	if !ok {
		endpoint, token, ok = v.parseCharacter()
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
	switch value := endpoint.(type) {
	// The order of these cases is very important since Go only compares the
	// set of methods supported by each interface. An interface that is a subset
	// of another interface must be checked AFTER that interface.
	case abs.NameLike:
		v.formatName(value)
	case abs.QuoteLike:
		v.formatQuote(value)
	case abs.VersionLike:
		v.formatVersion(value)
	case abs.DurationLike:
		v.formatDuration(value)
	case abs.MomentLike:
		v.formatMoment(value)
	case abs.PercentageLike:
		v.formatPercentage(value)
	case abs.FloatLike:
		v.formatFloat(value)
	case abs.IntegerLike:
		v.formatInteger(value)
	case abs.ProbabilityLike:
		v.formatProbability(value)
	case abs.AngleLike:
		v.formatAngle(value)
	case abs.CharacterLike:
		v.formatCharacter(value)
	case abs.PatternLike:
		v.formatPattern(value)
	case abs.ResourceLike:
		v.formatResource(value)
	case abs.TagLike:
		v.formatTag(value)
	case abs.SymbolLike:
		v.formatSymbol(value)
	default:
		panic(fmt.Sprintf("An invalid endpoint (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a range. It returns the range and whether or
// not the range was successfully parsed.
func (v *parser) parseRange() (abs.Range, *Token, bool) {
	var ok bool
	var token, bracketToken, endpointToken *Token
	var left, right string
	var first abs.Primitive
	var extent abs.Extent
	var last abs.Primitive
	var range_ abs.Range
	left, bracketToken, ok = v.parseDelimiter("[")
	if !ok {
		left, bracketToken, ok = v.parseDelimiter("(")
		if !ok {
			return range_, bracketToken, false
		}
	}
	first, endpointToken, ok = v.parseEndpoint()
	if !ok {
		// This is not a range.
		v.backupOne(bracketToken) // Put back the left bracket character.
		return range_, token, false
	}
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a range.
		v.backupOne(endpointToken) // Put back the first endpoint token.
		v.backupOne(bracketToken)  // Put back the left bracket character.
		return range_, token, false
	}
	last, token, ok = v.parseEndpoint()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("primitive",
			"$range",
			"$primitive")
		panic(message)
	}
	right, token, ok = v.parseDelimiter("]")
	if !ok {
		right, token, ok = v.parseDelimiter(")")
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("bracket",
				"$range")
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
	if ref.TypeOf(first) != ref.TypeOf(last) {
		var message = fmt.Sprintf("The range endpoints have two different types: %T and %T\n", first, last)
		panic(message)
	}
	switch first.(type) {
	case abs.Continuous:
		range_ = ran.Continuum(first.(abs.Continuous), extent, last.(abs.Continuous))
	case abs.Discrete:
		range_ = ran.Interval(first.(abs.Discrete), extent, last.(abs.Discrete))
	case abs.Lexical:
		range_ = ran.Spectrum(first.(abs.Lexical), extent, last.(abs.Lexical))
	default:
		var message = fmt.Sprintf("An invalid range endpoint (of type %T) was parsed: %v", first, first)
		panic(message)
	}
	return range_, token, true
}

// This method adds the canonical format for the specified interval to the
// state of the formatter.
func (v *formatter) formatInterval(interval abs.IntervalLike) {
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
func (v *formatter) formatSpectrum(spectrum abs.SpectrumLike) {
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
func (v *formatter) formatContinuum(continuum abs.ContinuumLike) {
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

// This method attempts to parse an integer. It returns the integer and whether
// or not the integer was successfully parsed.
func (v *parser) parseInteger() (ele.Integer, *Token, bool) {
	var number ele.Integer
	var token = v.nextToken()
	if token.Type != TokenNUMBER {
		// The token is not numerical.
		v.backupOne(token)
		return number, token, false
	}
	var matches = bytesToStrings(realScanner.FindSubmatch([]byte(token.Value)))
	if len(matches) > 0 {
		// The token is a real number.
		v.backupOne(token)
		return number, token, false
	}
	var err any
	var integer int64
	matches = bytesToStrings(integerScanner.FindSubmatch([]byte(token.Value)))
	if len(matches) == 0 {
		// The token is not an integer.
		v.backupOne(token)
		return number, token, false
	}
	integer, err = stc.ParseInt(matches[0], 10, 0)
	if err != nil {
		panic(fmt.Sprintf("The integer was not parsable: %v", err))
	}
	number = ele.Integer(integer)
	return number, token, true
}

// This method adds the canonical format for the specified integer to the state of
// the formatter.
func (v *formatter) formatInteger(number abs.IntegerLike) {
	var integer = number.AsInteger()
	v.AppendString(stc.FormatInt(int64(integer), 10))
}

// This method attempts to parse a real number. It returns the real number
// and whether or not the real number was successfully parsed.
func (v *parser) parseFloat() (ele.Float, *Token, bool) {
	var real_ ele.Float
	var token = v.nextToken()
	if token.Type != TokenREAL {
		v.backupOne(token)
		return real_, token, false
	}
	var real_ = ele.FloatFromString(token.Value)
	return real_, token, true
}

// This method adds the canonical format for the specified real number to the
// state of the formatter.
func (v *formatter) formatFloat(real_ abs.FloatLike) {
	v.AppendString(real_.AsString())
}

// This method attempts to parse a character. It returns the character and whether or not
// the character was successfully parsed.
func (v *parser) parseCharacter() (ele.Character, *Token, bool) {
	var character ele.Character
	var token = v.nextToken()
	if token.Type != TokenCHARACTER {
		v.backupOne(token)
		return character, token, false
	}
	var character = ele.CharacterFromString(token.Value)
	return character, token, true
}

// This method adds the canonical format for the specified character to the state of
// the formatter.
func (v *formatter) formatCharacter(character abs.CharacterLike) {
	v.AppendString(character.AsString())
}
