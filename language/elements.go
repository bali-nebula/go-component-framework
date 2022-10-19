/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language

import (
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This method attempts to parse an angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (ele.Angle, *Token, bool) {
	var token *Token
	var angle ele.Angle
	token = v.nextToken()
	if token.Type != TokenAngle {
		v.backupOne()
		return angle, token, false
	}
	angle, _ = ele.AngleFromString(token.Value)
	return angle, token, true
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (ele.Boolean, *Token, bool) {
	var token *Token
	var boolean ele.Boolean
	token = v.nextToken()
	if token.Type != TokenBoolean {
		v.backupOne()
		return boolean, token, false
	}
	boolean, _ = ele.BooleanFromString(token.Value)
	return boolean, token, true
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (ele.Duration, *Token, bool) {
	var token *Token
	var duration ele.Duration
	token = v.nextToken()
	if token.Type != TokenDuration {
		v.backupOne()
		return duration, token, false
	}
	duration, _ = ele.DurationFromString(token.Value)
	return duration, token, true
}

// This method attempts to parse an element primitive. It returns the
// element primitive and whether or not the element primitive was
// successfully parsed.
func (v *parser) parseElement() (any, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var element any
	element, token, ok = v.parseAngle()
	if !ok {
		element, token, ok = v.parseBoolean()
	}
	if !ok {
		element, token, ok = v.parseDuration()
	}
	if !ok {
		element, token, ok = v.parseMoment()
	}
	if !ok {
		element, token, ok = v.parseNumber()
	}
	if !ok {
		element, token, ok = v.parsePattern()
	}
	if !ok {
		element, token, ok = v.parsePercentage()
	}
	if !ok {
		element, token, ok = v.parseProbability()
	}
	if !ok {
		element, token, ok = v.parseResource()
	}
	if !ok {
		element, token, ok = v.parseSymbol()
	}
	if !ok {
		element, token, ok = v.parseTag()
	}
	if !ok {
		// Override any zero values returned from failed parsing attempts.
		element = nil
	}
	return element, token, ok
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (ele.Moment, *Token, bool) {
	var token *Token
	var moment ele.Moment
	token = v.nextToken()
	if token.Type != TokenMoment {
		v.backupOne()
		return moment, token, false
	}
	moment, _ = ele.MomentFromString(token.Value)
	return moment, token, true
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (ele.Number, *Token, bool) {
	var token *Token
	var number ele.Number
	token = v.nextToken()
	if token.Type != TokenNumber {
		v.backupOne()
		return number, token, false
	}
	number, _ = ele.NumberFromString(token.Value)
	return number, token, true
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (ele.Pattern, *Token, bool) {
	var token *Token
	var pattern ele.Pattern
	token = v.nextToken()
	if token.Type != TokenPattern {
		v.backupOne()
		return pattern, token, false
	}
	pattern, _ = ele.PatternFromString(token.Value)
	return pattern, token, true
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (ele.Percentage, *Token, bool) {
	var token *Token
	var percentage ele.Percentage
	token = v.nextToken()
	if token.Type != TokenPercentage {
		v.backupOne()
		return percentage, token, false
	}
	percentage, _ = ele.PercentageFromString(token.Value)
	return percentage, token, true
}

func (v *parser) parseProbability() (ele.Probability, *Token, bool) {
	var token *Token
	var probability ele.Probability
	token = v.nextToken()
	if token.Type != TokenProbability {
		v.backupOne()
		return probability, token, false
	}
	probability, _ = ele.ProbabilityFromString(token.Value)
	return probability, token, true
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (ele.Resource, *Token, bool) {
	var token *Token
	var resource ele.Resource
	token = v.nextToken()
	if token.Type != TokenResource {
		v.backupOne()
		return resource, token, false
	}
	resource, _ = ele.ResourceFromString(token.Value)
	return resource, token, true
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (ele.Symbol, *Token, bool) {
	var token *Token
	var symbol ele.Symbol
	token = v.nextToken()
	if token.Type != TokenSymbol {
		v.backupOne()
		return symbol, token, false
	}
	symbol, _ = ele.SymbolFromString(token.Value)
	return symbol, token, true
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (ele.Tag, *Token, bool) {
	var token *Token
	var tag ele.Tag
	token = v.nextToken()
	if token.Type != TokenTag {
		v.backupOne()
		return tag, token, false
	}
	tag, _ = ele.TagFromString(token.Value)
	return tag, token, true
}
