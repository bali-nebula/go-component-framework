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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	col "github.com/craterdog-bali/go-component-framework/collections"
	ele "github.com/craterdog-bali/go-component-framework/elements"
	str "github.com/craterdog-bali/go-component-framework/strings"
	co2 "github.com/craterdog/go-collection-framework"
	mat "math"
	ref "reflect"
	stc "strconv"
	utf "unicode/utf8"
)

// This method attempts to parse an association between a key and value. It
// returns the association and whether or not the association was successfully
// parsed.
func (v *parser) parseAssociation() (abs.AssociationLike, *Token, bool) {
	var ok bool
	var token *Token
	var key abs.Primitive
	var value abs.ComponentLike
	key, token, ok = v.parsePrimitive()
	if !ok {
		// This is not an association.
		return nil, token, false
	}
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		// This is not an association.
		v.backupOne() // Put back the primitive key token.
		return nil, token, false
	}
	// This must be an association.
	value, token, ok = v.parseComponent()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$component",
			"$association",
			"$primitive",
			"$component")
		panic(message)
	}
	var association = col.Association(key, value)
	return association, token, true
}

// This method adds the canonical format for the specified association to the
// state of the formatter.
func (v *formatter) formatAssociation(association abs.AssociationLike) {
	var key = association.GetKey()
	v.formatPrimitive(key)
	v.AppendString(": ")
	var value = association.GetValue()
	v.formatComponent(value)
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abs.CatalogLike, *Token, bool) {
	var ok bool
	var token *Token
	var catalog = col.Catalog()
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return catalog, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		catalog, token, ok = v.parseInlineAssociations()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return catalog, token, false
		}
	} else {
		catalog, token, ok = v.parseMultilineAssociations()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '[' character.
			return catalog, token, false
		}
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("]",
			"$collection",
			"$values")
		panic(message)
	}
	return catalog, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCatalog(catalog abs.CatalogLike) {
	v.AppendString("[")
	var iterator = co2.Iterator[abs.Binding](catalog)
	switch catalog.GetSize() {
	case 0:
		v.AppendString(":")
	case 1:
		var association = iterator.GetNext()
		v.formatAssociation(association)
	default:
		v.depth++
		for iterator.HasNext() {
			v.AppendNewline()
			var association = iterator.GetNext()
			v.formatAssociation(association)
		}
		v.depth--
		v.AppendNewline()
	}
	v.AppendString("]")
}

// This method attempts to parse a collection of values. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (abs.Collection, *Token, bool) {
	var ok bool
	var token *Token
	var collection abs.Collection
	collection, token, ok = v.parseCatalog()
	if !ok {
		collection, token, ok = v.parseRange()
	}
	if !ok {
		// The list must be attempted last since it may start with a component
		// which cannot be put back as a single token.
		collection, token, ok = v.parseList()
	}
	return collection, token, ok
}

// This method attempts to parse a catalog collection with inline associations.
// It returns the catalog collection and whether or not the catalog collection
// was successfully parsed.
func (v *parser) parseInlineAssociations() (abs.CatalogLike, *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var catalog = col.Catalog()
	_, token, ok = v.parseDelimiter(":")
	if ok {
		// This is an empty catalog.
		return catalog, token, true
	}
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty list.
		v.backupOne() // Put back the ']' character.
		return catalog, token, false
	}
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty catalog must have at least one association.
		return catalog, token, false
	}
	for {
		catalog.AddAssociation(association)
		// Every subsequent association must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more associations.
			break
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$association",
				"$catalog",
				"$association",
				"$primitive",
				"$component")
			panic(message)
		}
	}
	return catalog, token, true
}

// This method attempts to parse a list collection with inline values. It returns
// the list collection and whether or not the list collection was
// successfully parsed.
func (v *parser) parseInlineValues() (abs.ListLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var list = col.List()
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty list.
		v.backupOne() // Put back the ']' token.
		return list, token, true
	}
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one component value.
		var message = v.formatError(token)
		message += generateGrammar("$component",
			"$list",
			"$component")
		panic(message)
	}
	for {
		list.AddValue(value)
		// Every subsequent value must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more values.
			break
		}
		value, token, ok = v.parseComponent()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$component",
				"$list",
				"$component")
			panic(message)
		}
	}
	return list, token, true
}

// This method attempts to parse a list of values. It returns the
// list of values and whether or not the list of values was
// successfully parsed.
func (v *parser) parseList() (abs.ListLike, *Token, bool) {
	var ok bool
	var token *Token
	var list = col.List()
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return list, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		list, token, ok = v.parseInlineValues()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return list, token, false
		}
	} else {
		list, token, ok = v.parseMultilineValues()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '[' character.
			return list, token, false
		}
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("]",
			"$list",
			"$component")
		panic(message)
	}
	return list, token, ok
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatList(list abs.ListLike) {
	v.AppendString("[")
	var iterator = co2.Iterator[abs.ComponentLike](list)
	switch list.GetSize() {
	case 0:
		v.AppendString(" ")
	case 1:
		var value = iterator.GetNext()
		v.formatComponent(value)
	default:
		v.depth++
		for iterator.HasNext() {
			v.AppendNewline()
			var value = iterator.GetNext()
			v.formatComponent(value)
		}
		v.depth--
		v.AppendNewline()
	}
	v.AppendString("]")
}

// This method attempts to parse a catalog collection with multiline associations.
// It returns the catalog collection and whether or not the catalog collection
// was successfully parsed.
func (v *parser) parseMultilineAssociations() (abs.CatalogLike, *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var catalog = col.Catalog()
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty catalog must have at least one association.
		return catalog, token, false
	}
	for {
		catalog.AddAssociation(association)
		// Every association must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$catalog",
				"$association",
				"$primitive",
				"$component")
			panic(message)
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			// No more associations.
			break
		}
	}
	return catalog, token, true
}

// This method attempts to parse a list collection with multiline values. It
// returns the list collection and whether or not the list collection was
// successfully parsed.
func (v *parser) parseMultilineValues() (abs.ListLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var list = col.List()
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one value.
		var message = v.formatError(token)
		message += generateGrammar("$component",
			"$list",
			"$component")
		panic(message)
	}
	for {
		list.AddValue(value)
		// Every value must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$list",
				"$component")
			panic(message)
		}
		value, token, ok = v.parseComponent()
		if !ok {
			// No more values.
			break
		}
	}
	return list, token, true
}

// This method attempts to parse a numeric element. It returns the numeric
// element and whether or not the numeric element was successfully parsed.
func (v *parser) parseNumeric() (abs.Numeric, *Token, bool) {
	var ok bool
	var token *Token
	var numeric abs.Numeric
	numeric, token, ok = v.parseAngle()
	if !ok {
		numeric, token, ok = v.parseDuration()
	}
	if !ok {
		numeric, token, ok = v.parseMoment()
	}
	if !ok {
		numeric, token, ok = v.parsePercentage()
	}
	if !ok {
		numeric, token, ok = v.parseProbability()
	}
	if !ok {
		numeric, token, ok = v.parseReal()
	}
	if !ok {
		numeric, token, ok = v.parseRune()
	}
	if !ok {
		// Override any zero values returned from failed parsing attempts.
		numeric = nil
	}
	return numeric, token, ok
}

// This method adds the canonical format for the specified numeric element to
// the state of the formatter.
func (v *formatter) formatNumeric(numeric abs.Numeric) {
	// NOTE: A type switch will only work if each case specifies a unique
	// interface. If two different interfaces define the same method signatures
	// they are indistinguishable as types. To get around this we have added as
	// necessary a unique "dummy" method to each interface to guarantee that it
	// is unique.
	switch value := numeric.(type) {
	case ele.Angle:
		v.formatAngle(value)
	case ele.Duration:
		v.formatDuration(value)
	case ele.Moment:
		v.formatMoment(value)
	case ele.Percentage:
		v.formatPercentage(value)
	case ele.Probability:
		v.formatProbability(value)
	case col.Real:
		v.formatReal(value)
	case col.Rune:
		v.formatRune(value)
	default:
		panic(fmt.Sprintf("An invalid numeric (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a primitive. It returns the primitive and
// whether or not the primitive was successfully parsed.
func (v *parser) parsePrimitive() (abs.Primitive, *Token, bool) {
	var ok bool
	var token *Token
	var primitive abs.Primitive
	primitive, token, ok = v.parseElement()
	if !ok {
		primitive, token, ok = v.parseString()
	}
	if !ok {
		// This must be explicitly set to nil since it is of type any.
		primitive = nil
	}
	return primitive, token, ok
}

// This method adds the canonical format for the specified primitive to the
// state of the formatter.
func (v *formatter) formatPrimitive(primitive abs.Primitive) {
	// NOTE: A type switch will only work if each case specifies a unique
	// interface. If two different interfaces define the same method signatures
	// they are indistinguishable as types. To get around this we have added as
	// necessary a unique "dummy" method to each interface to guarantee that it
	// is unique.
	switch value := primitive.(type) {
	case ele.Angle:
		v.formatAngle(value)
	case ele.Boolean:
		v.formatBoolean(value)
	case ele.Duration:
		v.formatDuration(value)
	case ele.Moment:
		v.formatMoment(value)
	case ele.Number:
		v.formatNumber(value)
	case ele.Pattern:
		v.formatPattern(value)
	case ele.Percentage:
		v.formatPercentage(value)
	case ele.Probability:
		v.formatProbability(value)
	case ele.Resource:
		v.formatResource(value)
	case ele.Symbol:
		v.formatSymbol(value)
	case ele.Tag:
		v.formatTag(value)
	case str.Binary:
		v.formatBinary(value)
	case str.Moniker:
		v.formatMoniker(value)
	case str.Narrative:
		v.formatNarrative(value)
	case str.Quote:
		v.formatQuote(value)
	case str.Version:
		v.formatVersion(value)
	default:
		panic(fmt.Sprintf("An invalid primitive (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a range collection. It returns the range
// collection and whether or not the range collection was successfully parsed.
func (v *parser) parseRange() (abs.RangeLike[abs.Numeric], *Token, bool) {
	var ok bool
	var token *Token
	var left, right string
	var first abs.Numeric
	var extent abs.Extent
	var last abs.Numeric
	var range_ abs.RangeLike[abs.Numeric]
	left, token, ok = v.parseDelimiter("[")
	if !ok {
		left, token, ok = v.parseDelimiter("(")
		if !ok {
			return range_, token, false
		}
	}
	first, token, _ = v.parseNumeric() // The first value is optional.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a range collection.
		if first != nil {
			v.backupOne() // Put back the first value token.
		}
		v.backupOne() // Put back the left bracket character.
		return range_, token, false
	}
	last, token, _ = v.parseNumeric() // The last value is optional.
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
		var message = fmt.Sprintf("Expected valid range brackets but received:%q and %q\n\n", left, right)
		panic(message)
	}
	if !ref.ValueOf(first).IsValid() && !ref.ValueOf(last).IsValid() {
		// Default to real endpoints if neither is specified.
		first = col.MinimumReal()
		last = col.MaximumReal()
	}
	range_ = col.Range(first, extent, last)
	return range_, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatRange(range_ abs.RangeLike[abs.Numeric]) {
	var extent = range_.GetExtent()
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
		panic(fmt.Sprintf("The range contains an unknown extent: %v", extent))
	}
	v.AppendString(left)
	var first = range_.GetFirst()
	if first != range_.GetMinimum() {
		v.formatNumeric(first)
	}
	v.AppendString("..")
	var last = range_.GetLast()
	if last != range_.GetMaximum() {
		v.formatNumeric(last)
	}
	v.AppendString(right)
}

// This method attempts to parse a real number. It returns the real number
// and whether or not the real number was successfully parsed.
func (v *parser) parseReal() (col.Real, *Token, bool) {
	var number col.Real
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
	case matches[0] == "infinity" || matches[0] == "∞":
		r = mat.Inf(1)
	case matches[0] == "pi", matches[0] == "-pi", matches[0] == "phi", matches[0] == "-phi":
		r = stringToFloat(matches[0])
	default:
		r, err = stc.ParseFloat(matches[0], 64)
		if err != nil {
			r = stringToFloat(matches[0])
		}
	}
	number = col.Real(r)
	return number, token, true
}

// This method adds the canonical format for the specified real number to the
// state of the formatter.
func (v *formatter) formatReal(number col.Real) {
	switch {
	case number.IsZero():
		v.AppendString("0")
	case number.IsInfinite():
		if number.IsNegative() {
			v.AppendString("-")
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
func (v *parser) parseRune() (col.Rune, *Token, bool) {
	var number = col.Rune(-1)
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
	number = col.Rune(r)
	return number, token, true
}

// This method adds the canonical format for the specified rune to the state of
// the formatter.
func (v *formatter) formatRune(number col.Rune) {
	var runes = []rune{rune(number)}
	var quote = str.Quote(string(runes))
	v.formatQuote(quote)
}
