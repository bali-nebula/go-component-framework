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
	"github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This method attempts to parse an angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (elements.Angle, *Token, bool) {
	var token *Token
	var angle elements.Angle
	token = v.nextToken()
	if token.Type != TokenAngle {
		v.backupOne()
		return angle, token, false
	}
	angle, _ = elements.AngleFromString(token.Value)
	return angle, token, true
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (elements.Boolean, *Token, bool) {
	var token *Token
	var boolean elements.Boolean
	token = v.nextToken()
	if token.Type != TokenBoolean {
		v.backupOne()
		return boolean, token, false
	}
	boolean, _ = elements.BooleanFromString(token.Value)
	return boolean, token, true
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (elements.Duration, *Token, bool) {
	var token *Token
	var duration elements.Duration
	token = v.nextToken()
	if token.Type != TokenDuration {
		v.backupOne()
		return duration, token, false
	}
	duration, _ = elements.DurationFromString(token.Value)
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
func (v *parser) parseMoment() (elements.Moment, *Token, bool) {
	var token *Token
	var moment elements.Moment
	token = v.nextToken()
	if token.Type != TokenMoment {
		v.backupOne()
		return moment, token, false
	}
	moment, _ = elements.MomentFromString(token.Value)
	return moment, token, true
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (elements.Number, *Token, bool) {
	var token *Token
	var number elements.Number
	token = v.nextToken()
	if token.Type != TokenNumber {
		v.backupOne()
		return number, token, false
	}
	number, _ = elements.NumberFromString(token.Value)
	return number, token, true
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (elements.Pattern, *Token, bool) {
	var token *Token
	var pattern elements.Pattern
	token = v.nextToken()
	if token.Type != TokenPattern {
		v.backupOne()
		return pattern, token, false
	}
	pattern, _ = elements.PatternFromString(token.Value)
	return pattern, token, true
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (elements.Percentage, *Token, bool) {
	var token *Token
	var percentage elements.Percentage
	token = v.nextToken()
	if token.Type != TokenPercentage {
		v.backupOne()
		return percentage, token, false
	}
	percentage, _ = elements.PercentageFromString(token.Value)
	return percentage, token, true
}

func (v *parser) parseProbability() (elements.Probability, *Token, bool) {
	var token *Token
	var probability elements.Probability
	token = v.nextToken()
	if token.Type != TokenProbability {
		v.backupOne()
		return probability, token, false
	}
	probability, _ = elements.ProbabilityFromString(token.Value)
	return probability, token, true
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (elements.Resource, *Token, bool) {
	var token *Token
	var resource elements.Resource
	token = v.nextToken()
	if token.Type != TokenResource {
		v.backupOne()
		return resource, token, false
	}
	resource, _ = elements.ResourceFromString(token.Value)
	return resource, token, true
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (elements.Symbol, *Token, bool) {
	var token *Token
	var symbol elements.Symbol
	token = v.nextToken()
	if token.Type != TokenSymbol {
		v.backupOne()
		return symbol, token, false
	}
	symbol, _ = elements.SymbolFromString(token.Value)
	return symbol, token, true
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (elements.Tag, *Token, bool) {
	var token *Token
	var tag elements.Tag
	token = v.nextToken()
	if token.Type != TokenTag {
		v.backupOne()
		return tag, token, false
	}
	tag, _ = elements.TagFromString(token.Value)
	return tag, token, true
}
