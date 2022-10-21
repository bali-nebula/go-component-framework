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
	mat "math"
	str "strconv"
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatAngle(angle ele.Angle) {
	var s string
	if float64(angle) == mat.Pi {
		s = "~π"
	} else {
		s = "~" + str.FormatFloat(float64(angle), 'G', -1, 64)
	}
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatBoolean(boolean ele.Boolean) {
	var s = str.FormatBool(bool(boolean))
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatDuration(duration ele.Duration) {
	var result str.Builder
	result.WriteString("~")
	if duration.IsNegative() {
		result.WriteString("-")
	}
	result.WriteString("P")
	result.WriteString(stc.FormatInt(int64(duration.GetYears()), 10))
	result.WriteString("Y")
	result.WriteString(stc.FormatInt(int64(duration.GetMonths()), 10))
	result.WriteString("M")
	result.WriteString(stc.FormatInt(int64(duration.GetDays()), 10))
	result.WriteString("DT")
	result.WriteString(stc.FormatInt(int64(duration.GetHours()), 10))
	result.WriteString("H")
	result.WriteString(stc.FormatInt(int64(duration.GetMinutes()), 10))
	result.WriteString("M")
	result.WriteString(stc.FormatInt(int64(duration.GetSeconds()), 10))
	result.WriteString(".")
	result.WriteString(stc.FormatInt(int64(duration.GetMilliseconds()), 10))
	result.WriteString("S")
	var s = result.String()
	v.state.AppendString(s)
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

// This method adds the canonical format for the specified element primitive to the
// state of the formatter.
func (v *formatter) formatElement(element any) {
	var value = ref.ValueOf(element)
	switch {
	case value.MethodByName("IsAngle").IsValid():
		v.formatAngle(element)
	case value.MethodByName("IsBoolean").IsValid():
		v.formatBoolean(element)
	case value.MethodByName("IsDuration").IsValid():
		v.formatDuration(element)
	case value.MethodByName("IsMoment").IsValid():
		v.formatMoment(element)
	case value.MethodByName("IsNumber").IsValid():
		v.formatNumber(element)
	case value.MethodByName("IsPattern").IsValid():
		v.formatPattern(element)
	case value.MethodByName("IsPercentage").IsValid():
		v.formatPercentage(element)
	case value.MethodByName("IsProbability").IsValid():
		v.formatProbability(element)
	case value.MethodByName("IsResource").IsValid():
		v.formatResource(element)
	case value.MethodByName("IsSymbol").IsValid():
		v.formatSymbol(element)
	default:
		v.formatTag(element)
	}
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatMoment(moment ele.Moment) {
	var result str.Builder
	result.WriteString("<")
	var year = moment.GetYears()
	result.WriteString(stc.FormatInt(int64(year), 10))
	var month = moment.GetMonths()
	if month > 1 {
		result.WriteString("-")
		result.WriteString(formatOrdinal(month))
		var day = moment.GetDays()
		if day > 1 {
			result.WriteString("-")
			result.WriteString(formatOrdinal(day))
			var hour = moment.GetHours()
			if hour > 0 {
				result.WriteString("T")
				result.WriteString(formatOrdinal(hour))
				var minute = moment.GetMinutes()
				if minute > 0 {
					result.WriteString(":")
					result.WriteString(formatOrdinal(minute))
					var second = moment.GetSeconds()
					if second > 0 {
						result.WriteString(":")
						result.WriteString(formatOrdinal(second))
						var millisecond = moment.GetMilliseconds()
						if millisecond > 0 {
							result.WriteString(".")
							result.WriteString(formatOrdinal(millisecond))
						}
					}
				}
			}
		}
	}
	result.WriteString(">")
	var s = result.String()
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatNumber(number ele.Number) {
	var s = number.AsRectangular()
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatPattern(pattern ele.Pattern) {
	var s = string(pattern)
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatPercentage(percentage ele.Percentage) {
	var s = str.FormatFloat(float64(percentage), 'G', -1, 64) + "%"
	v.state.AppendString(s)
}

// This method attempts to parse a probability element. It returns the probability
// element and whether or not the probability element was successfully parsed.
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatProbability(probability ele.Probability) {
	var s = str.FormatFloat(float64(probability), 'f', -1, 64)
	switch {
	// Zero is formatted as ".0".
	case probability == 0:
		s = "." + s
	// One is formatted as "1.".
	case probability == 1:
		s += "."
	// Otherwise, remove the leading "0".
	default:
		s = s[1:]
	}
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatResource(resource ele.Resource) {
	var s = string(resource)
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatSymbol(symbol ele.Symbol) {
	var s = string(symbol)
	v.state.AppendString(s)
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

// This method adds the canonical format for this element to the state of the
// formatter.
func (v *formatter) formatTag(tag ele.Tag) {
	var s = string(tag)
	v.state.AppendString(s)
}
