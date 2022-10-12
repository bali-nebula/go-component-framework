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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This method attempts to parse an angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (elements.Angle, *Token, bool) {
	var ok bool
	var token *Token
	var angle elements.Angle
	token = v.nextToken()
	if token.Type != TokenAngle {
		v.backupOne()
		return angle, token, false
	}
	angle, ok = elements.AngleFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid angle token was found: %v", token))
	}
	return angle, token, true
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (elements.Boolean, *Token, bool) {
	var ok bool
	var token *Token
	var boolean elements.Boolean
	token = v.nextToken()
	if token.Type != TokenBoolean {
		v.backupOne()
		return boolean, token, false
	}
	boolean, ok = elements.BooleanFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid boolean token was found: %v", token))
	}
	return boolean, token, true
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (elements.Duration, *Token, bool) {
	var ok bool
	var token *Token
	var duration elements.Duration
	token = v.nextToken()
	if token.Type != TokenDuration {
		v.backupOne()
		return duration, token, false
	}
	duration, ok = elements.DurationFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid duration token was found: %v", token))
	}
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
		element = nil
	}
	return element, token, ok
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (elements.Moment, *Token, bool) {
	var ok bool
	var token *Token
	var moment elements.Moment
	token = v.nextToken()
	if token.Type != TokenMoment {
		v.backupOne()
		return moment, token, false
	}
	moment, ok = elements.MomentFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid moment token was found: %v", token))
	}
	return moment, token, true
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (elements.Number, *Token, bool) {
	var ok bool
	var token *Token
	var number elements.Number
	token = v.nextToken()
	if token.Type != TokenNumber {
		v.backupOne()
		return number, token, false
	}
	number, ok = elements.NumberFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid number token was found: %v", token))
	}
	return number, token, true
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (elements.Pattern, *Token, bool) {
	var ok bool
	var token *Token
	var pattern elements.Pattern
	token = v.nextToken()
	if token.Type != TokenPattern {
		v.backupOne()
		return pattern, token, false
	}
	pattern, ok = elements.PatternFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid pattern token was found: %v", token))
	}
	return pattern, token, true
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (elements.Percentage, *Token, bool) {
	var ok bool
	var token *Token
	var percentage elements.Percentage
	token = v.nextToken()
	if token.Type != TokenPercentage {
		v.backupOne()
		return percentage, token, false
	}
	percentage, ok = elements.PercentageFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid percentage token was found: %v", token))
	}
	return percentage, token, true
}

func (v *parser) parseProbability() (elements.Probability, *Token, bool) {
	var ok bool
	var token *Token
	var probability elements.Probability
	token = v.nextToken()
	if token.Type != TokenProbability {
		v.backupOne()
		return probability, token, false
	}
	probability, ok = elements.ProbabilityFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid probability token was found: %v", token))
	}
	return probability, token, true
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (elements.Resource, *Token, bool) {
	var ok bool
	var token *Token
	var resource elements.Resource
	token = v.nextToken()
	if token.Type != TokenResource {
		v.backupOne()
		return resource, token, false
	}
	resource, ok = elements.ResourceFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid resource token was found: %v", token))
	}
	return resource, token, true
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (elements.Symbol, *Token, bool) {
	var ok bool
	var token *Token
	var symbol elements.Symbol
	token = v.nextToken()
	if token.Type != TokenSymbol {
		v.backupOne()
		return symbol, token, false
	}
	symbol, ok = elements.SymbolFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid symbol token was found: %v", token))
	}
	return symbol, token, true
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (elements.Tag, *Token, bool) {
	var ok bool
	var token *Token
	var tag elements.Tag
	token = v.nextToken()
	if token.Type != TokenTag {
		v.backupOne()
		return tag, token, false
	}
	tag, ok = elements.TagFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid tag token was found: %v", token))
	}
	return tag, token, true
}
