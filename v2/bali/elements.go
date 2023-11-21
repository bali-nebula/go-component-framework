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
	mat "math"
	cmp "math/cmplx"
	uri "net/url"
	stc "strconv"
	sts "strings"
	tim "time"
)

// UNIVERSAL CONSTRUCTORS

// This constructor returns a new angle element initialized with the specified
// value.
func Angle(value abs.Value) abs.AngleLike {
	var angle abs.AngleLike
	switch actual := value.(type) {
	case uint:
		angle = ele.AngleFromFloat(float64(actual))
	case uint8:
		angle = ele.AngleFromFloat(float64(actual))
	case uint16:
		angle = ele.AngleFromFloat(float64(actual))
	case uint32:
		angle = ele.AngleFromFloat(float64(actual))
	case uint64:
		angle = ele.AngleFromFloat(float64(actual))
	case int:
		angle = ele.AngleFromFloat(float64(actual))
	case int8:
		angle = ele.AngleFromFloat(float64(actual))
	case int16:
		angle = ele.AngleFromFloat(float64(actual))
	case int32:
		angle = ele.AngleFromFloat(float64(actual))
	case int64:
		angle = ele.AngleFromFloat(float64(actual))
	case float32:
		angle = ele.AngleFromFloat(float64(actual))
	case float64:
		angle = ele.AngleFromFloat(float64(actual))
	case string:
		angle = ele.AngleFromString(actual)
	case abs.AngleLike:
		angle = actual
	case abs.ComponentLike:
		angle = actual.GetEntity().(abs.AngleLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to an angle: %v", actual, actual)
		panic(message)
	}
	return angle
}

// This constructor returns a new boolean element initialized with the specified
// value.
func Boolean(value abs.Value) abs.BooleanLike {
	var boolean abs.BooleanLike
	switch actual := value.(type) {
	case bool:
		boolean = ele.BooleanFromBoolean(actual)
	case string:
		boolean = ele.BooleanFromString(actual)
	case abs.BooleanLike:
		boolean = actual
	case abs.ComponentLike:
		boolean = actual.GetEntity().(abs.BooleanLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a boolean: %v", actual, actual)
		panic(message)
	}
	return boolean
}

// This constructor returns a new duration element initialized with the specified
// value.
func Duration(value abs.Value) abs.DurationLike {
	var duration abs.DurationLike
	switch actual := value.(type) {
	case uint:
		duration = ele.DurationFromMilliseconds(int(actual))
	case uint8:
		duration = ele.DurationFromMilliseconds(int(actual))
	case uint16:
		duration = ele.DurationFromMilliseconds(int(actual))
	case uint32:
		duration = ele.DurationFromMilliseconds(int(actual))
	case uint64:
		duration = ele.DurationFromMilliseconds(int(actual))
	case int:
		duration = ele.DurationFromMilliseconds(int(actual))
	case int8:
		duration = ele.DurationFromMilliseconds(int(actual))
	case int16:
		duration = ele.DurationFromMilliseconds(int(actual))
	case int32:
		duration = ele.DurationFromMilliseconds(int(actual))
	case int64:
		duration = ele.DurationFromMilliseconds(int(actual))
	case string:
		duration = ele.DurationFromString(actual)
	case abs.DurationLike:
		duration = actual
	case abs.ComponentLike:
		duration = actual.GetEntity().(abs.DurationLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a duration: %v", actual, actual)
		panic(message)
	}
	return duration
}

// This constructor returns a new moment element initialized with the specified
// value.
func Moment(value abs.Value) abs.MomentLike {
	var moment abs.MomentLike
	switch actual := value.(type) {
	case uint:
		moment = ele.MomentFromMilliseconds(int(actual))
	case uint8:
		moment = ele.MomentFromMilliseconds(int(actual))
	case uint16:
		moment = ele.MomentFromMilliseconds(int(actual))
	case uint32:
		moment = ele.MomentFromMilliseconds(int(actual))
	case uint64:
		moment = ele.MomentFromMilliseconds(int(actual))
	case int:
		moment = ele.MomentFromMilliseconds(int(actual))
	case int8:
		moment = ele.MomentFromMilliseconds(int(actual))
	case int16:
		moment = ele.MomentFromMilliseconds(int(actual))
	case int32:
		moment = ele.MomentFromMilliseconds(int(actual))
	case int64:
		moment = ele.MomentFromMilliseconds(int(actual))
	case string:
		moment = ele.MomentFromString(actual)
	case abs.MomentLike:
		moment = actual
	case abs.ComponentLike:
		moment = actual.GetEntity().(abs.MomentLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a moment: %v", actual, actual)
		panic(message)
	}
	return moment
}

// This constructor returns a new number element initialized with the specified
// value.
func Number(value abs.Value) abs.NumberLike {
	var number abs.NumberLike
	switch actual := value.(type) {
	case uint:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint8:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint16:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int8:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int16:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case float32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case float64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case complex64:
		number = ele.NumberFromComplex(complex128(actual))
	case complex128:
		number = ele.NumberFromComplex(complex128(actual))
	case string:
		number = ele.NumberFromString(actual)
	case abs.NumberLike:
		number = actual
	case abs.ComponentLike:
		number = actual.GetEntity().(abs.NumberLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a number: %v", actual, actual)
		panic(message)
	}
	return number
}

// This constructor returns a new pattern element initialized with the specified
// value.
func Pattern(value abs.Value) abs.PatternLike {
	var pattern abs.PatternLike
	switch actual := value.(type) {
	case string:
		pattern = ele.PatternFromString(actual)
	case abs.PatternLike:
		pattern = actual
	case abs.ComponentLike:
		pattern = actual.GetEntity().(abs.PatternLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a pattern: %v", actual, actual)
		panic(message)
	}
	return pattern
}

// This constructor returns a new percentage element initialized with the specified
// value.
func Percentage(value abs.Value) abs.PercentageLike {
	var percentage abs.PercentageLike
	switch actual := value.(type) {
	case uint:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint8:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint16:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int8:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int16:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case float32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case float64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case string:
		percentage = ele.PercentageFromString(actual)
	case abs.PercentageLike:
		percentage = actual
	case abs.ComponentLike:
		percentage = actual.GetEntity().(abs.PercentageLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a percentage: %v", actual, actual)
		panic(message)
	}
	return percentage
}

// This constructor returns a new probability element initialized with the specified
// value.
func Probability(value abs.Value) abs.ProbabilityLike {
	var probability abs.ProbabilityLike
	switch actual := value.(type) {
	case bool:
		probability = ele.ProbabilityFromBool(actual)
	case uint:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint8:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint16:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int8:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int16:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case float32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case float64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case string:
		probability = ele.ProbabilityFromString(actual)
	case abs.ProbabilityLike:
		probability = actual
	case abs.ComponentLike:
		probability = actual.GetEntity().(abs.ProbabilityLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a probability: %v", actual, actual)
		panic(message)
	}
	return probability
}

// This constructor returns a new resource element initialized with the specified
// value.
func Resource(value abs.Value) abs.ResourceLike {
	var resource abs.ResourceLike
	switch actual := value.(type) {
	case *uri.URL:
		resource = ele.ResourceFromURI(actual)
	case string:
		resource = ele.ResourceFromString(actual)
	case abs.ResourceLike:
		resource = actual
	case abs.ComponentLike:
		resource = actual.GetEntity().(abs.ResourceLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a resource: %v", actual, actual)
		panic(message)
	}
	return resource
}

// PRIVATE METHODS

// This method attempts to parse an angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (abs.AngleLike, *Token, bool) {
	var token *Token
	var angle abs.AngleLike
	token = v.nextToken()
	if token.Type != TokenANGLE {
		v.backupOne(token)
		return angle, token, false
	}
	angle = ele.AngleFromString(token.Value)
	return angle, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatAngle(angle abs.AngleLike) {
	var string_ = angle.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (abs.BooleanLike, *Token, bool) {
	var token *Token
	var boolean abs.BooleanLike
	token = v.nextToken()
	if token.Type != TokenBOOLEAN {
		v.backupOne(token)
		return boolean, token, false
	}
	boolean = ele.BooleanFromString(token.Value)
	return boolean, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatBoolean(boolean abs.BooleanLike) {
	var string_ = boolean.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (abs.DurationLike, *Token, bool) {
	var token *Token
	var duration abs.DurationLike
	token = v.nextToken()
	if token.Type != TokenDURATION {
		v.backupOne(token)
		return duration, token, false
	}
	duration = ele.DurationFromString(token.Value)
	return duration, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatDuration(duration abs.DurationLike) {
	var string_ = duration.AsString()
	v.AppendString(string_)
}

// This method attempts to parse an element primitive. It returns the
// element primitive and whether or not the element primitive was
// successfully parsed.
func (v *parser) parseElement() (abs.Element, *Token, bool) {
	var ok bool
	var token *Token
	var element abs.Element
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
		// Override any zero values returned from failed parsing attempts.
		element = nil
	}
	return element, token, ok
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (abs.MomentLike, *Token, bool) {
	var token *Token
	var moment abs.MomentLike
	token = v.nextToken()
	if token.Type != TokenMOMENT {
		v.backupOne(token)
		return moment, token, false
	}
	moment = ele.MomentFromString(token.Value)
	return moment, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatMoment(moment abs.MomentLike) {
	var string_ = moment.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (abs.NumberLike, *Token, bool) {
	var number abs.NumberLike
	var token = v.nextToken()
	if token.Type != TokenNUMBER {
		v.backupOne(token)
		return number, token, false
	}
	number = ele.NumberFromString(token.Value)
	return number, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatNumber(number abs.NumberLike) {
	var string_ = number.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (abs.PatternLike, *Token, bool) {
	var token *Token
	var pattern abs.PatternLike
	token = v.nextToken()
	if token.Type != TokenPATTERN {
		v.backupOne(token)
		return pattern, token, false
	}
	pattern = ele.PatternFromString(token.Value)
	return pattern, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatPattern(pattern abs.PatternLike) {
	var string_ = pattern.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (abs.PercentageLike, *Token, bool) {
	var token *Token
	var percentage abs.PercentageLike
	token = v.nextToken()
	if token.Type != TokenPERCENTAGE {
		v.backupOne(token)
		return percentage, token, false
	}
	percentage = ele.PercentageFromString(token.Value)
	return percentage, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatPercentage(percentage abs.PercentageLike) {
	var string_ = percentage.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a probability element. It returns the probability
// element and whether or not the probability element was successfully parsed.
func (v *parser) parseProbability() (abs.ProbabilityLike, *Token, bool) {
	var token *Token
	var probability abs.ProbabilityLike
	token = v.nextToken()
	if token.Type != TokenPROBABILITY {
		v.backupOne(token)
		return probability, token, false
	}
	probability = ele.ProbabilityFromString(token.Value)
	return probability, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatProbability(probability abs.ProbabilityLike) {
	var string_ = probability.AsString()
	v.AppendString(string_)
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (abs.ResourceLike, *Token, bool) {
	var token *Token
	var resource abs.ResourceLike
	token = v.nextToken()
	if token.Type != TokenRESOURCE {
		v.backupOne(token)
		return resource, token, false
	}
	resource = ele.ResourceFromString(token.Value)
	return resource, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatResource(resource abs.ResourceLike) {
	var string_ = resource.AsString()
	v.AppendString(string_)
}
