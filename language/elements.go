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
func (v *parser) parseAngle() (elements.Angle, bool) {
	var ok bool
	var angle elements.Angle
	var token = v.nextToken()
	if token.typ != tokenAngle {
		v.backupOne()
		return angle, false
	}
	angle, ok = elements.AngleFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid angle token was found: %v", token))
	}
	return angle, true
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (elements.Boolean, bool) {
	var ok bool
	var boolean elements.Boolean
	var token = v.nextToken()
	if token.typ != tokenBoolean {
		v.backupOne()
		return boolean, false
	}
	boolean, ok = elements.BooleanFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid boolean token was found: %v", token))
	}
	return boolean, true
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (elements.Duration, bool) {
	var ok bool
	var duration elements.Duration
	var token = v.nextToken()
	if token.typ != tokenDuration {
		v.backupOne()
		return duration, false
	}
	duration, ok = elements.DurationFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid duration token was found: %v", token))
	}
	return duration, true
}

// This method attempts to parse an element primitive. It returns the
// element primitive and whether or not the element primitive was
// successfully parsed.
func (v *parser) parseElement() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var element any
	element, ok = v.parseAngle()
	if !ok {
		element, ok = v.parseBoolean()
	}
	if !ok {
		element, ok = v.parseDuration()
	}
	if !ok {
		element, ok = v.parseMoment()
	}
	if !ok {
		element, ok = v.parseNumber()
	}
	if !ok {
		element, ok = v.parsePattern()
	}
	if !ok {
		element, ok = v.parsePercentage()
	}
	if !ok {
		element, ok = v.parseProbability()
	}
	if !ok {
		element, ok = v.parseResource()
	}
	if !ok {
		element, ok = v.parseSymbol()
	}
	if !ok {
		element, ok = v.parseTag()
	}
	if !ok {
		element = nil
	}
	return element, ok
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (elements.Moment, bool) {
	var ok bool
	var moment elements.Moment
	var token = v.nextToken()
	if token.typ != tokenMoment {
		v.backupOne()
		return moment, false
	}
	moment, ok = elements.MomentFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid moment token was found: %v", token))
	}
	return moment, true
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (elements.Number, bool) {
	var ok bool
	var number elements.Number
	var token = v.nextToken()
	if token.typ != tokenNumber {
		v.backupOne()
		return number, false
	}
	number, ok = elements.NumberFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid number token was found: %v", token))
	}
	return number, true
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (elements.Pattern, bool) {
	var ok bool
	var pattern elements.Pattern
	var token = v.nextToken()
	if token.typ != tokenPattern {
		v.backupOne()
		return pattern, false
	}
	pattern, ok = elements.PatternFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid pattern token was found: %v", token))
	}
	return pattern, true
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (elements.Percentage, bool) {
	var ok bool
	var percentage elements.Percentage
	var token = v.nextToken()
	if token.typ != tokenPercentage {
		v.backupOne()
		return percentage, false
	}
	percentage, ok = elements.PercentageFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid percentage token was found: %v", token))
	}
	return percentage, true
}

func (v *parser) parseProbability() (elements.Probability, bool) {
	var ok bool
	var probability elements.Probability
	var token = v.nextToken()
	if token.typ != tokenProbability {
		v.backupOne()
		return probability, false
	}
	probability, ok = elements.ProbabilityFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid probability token was found: %v", token))
	}
	return probability, true
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (elements.Resource, bool) {
	var ok bool
	var resource elements.Resource
	var token = v.nextToken()
	if token.typ != tokenResource {
		v.backupOne()
		return resource, false
	}
	resource, ok = elements.ResourceFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid resource token was found: %v", token))
	}
	return resource, true
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (elements.Symbol, bool) {
	var ok bool
	var symbol elements.Symbol
	var token = v.nextToken()
	if token.typ != tokenSymbol {
		v.backupOne()
		return symbol, false
	}
	symbol, ok = elements.SymbolFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid symbol token was found: %v", token))
	}
	return symbol, true
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (elements.Tag, bool) {
	var ok bool
	var tag elements.Tag
	var token = v.nextToken()
	if token.typ != tokenTag {
		v.backupOne()
		return tag, false
	}
	tag, ok = elements.TagFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid tag token was found: %v", token))
	}
	return tag, true
}
