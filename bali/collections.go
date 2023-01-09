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
	col "github.com/bali-nebula/go-component-framework/collections"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
	cox "github.com/craterdog/go-collection-framework"
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
	var association = cox.Association[abs.Primitive, abs.ComponentLike](key, value)
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

// This method attempts to parse a collection of values. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (abs.Collection, *Token, bool) {
	var ok bool
	var token *Token
	var collection abs.Collection
	collection, token, ok = v.parseMapping()
	if !ok {
		// The series must be attempted last since it may start with a component
		// which cannot be put back as a single token.
		collection, token, ok = v.parseSeries()
	}
	return collection, token, ok
}

// It returns the mapping collection and whether or not the mapping collection
// was successfully parsed.
func (v *parser) parseInlineAssociations() (abs.Sequential[abs.AssociationLike], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var associations = cox.List[abs.AssociationLike]()
	_, token, ok = v.parseDelimiter(":")
	if ok {
		// This is an empty mapping.
		return associations, token, true
	}
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty series.
		v.backupOne() // Put back the ']' character.
		return associations, token, false
	}
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty mapping must have at least one association.
		return associations, token, false
	}
	for {
		associations.AddValue(association)
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
				"$mapping",
				"$association",
				"$primitive",
				"$component")
			panic(message)
		}
	}
	return associations, token, true
}

// This method attempts to parse a series collection with inline values. It
// returns the series collection and whether or not the series collection was
// successfully parsed.
func (v *parser) parseInlineValues() (abs.Sequential[abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var values = cox.List[abs.ComponentLike]()
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty series.
		v.backupOne() // Put back the ']' token.
		return values, token, true
	}
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty series must have at least one component value.
		var message = v.formatError(token)
		message += generateGrammar("$component",
			"$series",
			"$component")
		panic(message)
	}
	for {
		values.AddValue(value)
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
				"$series",
				"$component")
			panic(message)
		}
	}
	return values, token, true
}

// This method attempts to parse a mapping collection. It returns the
// mapping collection and whether or not the mapping collection was
// successfully parsed.
func (v *parser) parseMapping() (abs.Collection, *Token, bool) {
	var ok bool
	var token *Token
	var mapping abs.Collection
	var associations abs.Sequential[abs.AssociationLike]
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return mapping, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		associations, token, ok = v.parseInlineAssociations()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return mapping, token, false
		}
	} else {
		associations, token, ok = v.parseMultilineAssociations()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '[' character.
			return mapping, token, false
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
	mapping = col.MappingFromSequence(associations)
	return mapping, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatMapping(mapping abs.MappingLike) {
	v.AppendString("[")
	var iterator = cox.Iterator[abs.AssociationLike](mapping)
	switch mapping.GetSize() {
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

// This method attempts to parse a mapping collection with multiline associations.
// It returns the mapping collection and whether or not the mapping collection
// was successfully parsed.
func (v *parser) parseMultilineAssociations() (abs.Sequential[abs.AssociationLike], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var associations = cox.List[abs.AssociationLike]()
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty mapping must have at least one association.
		return associations, token, false
	}
	for {
		associations.AddValue(association)
		// Every association must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$mapping",
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
	return associations, token, true
}

// This method attempts to parse a series collection with multiline values. It
// returns the series collection and whether or not the series collection was
// successfully parsed.
func (v *parser) parseMultilineValues() (abs.Sequential[abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var values = cox.List[abs.ComponentLike]()
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty series must have at least one value.
		var message = v.formatError(token)
		message += generateGrammar("$component",
			"$series",
			"$component")
		panic(message)
	}
	for {
		values.AddValue(value)
		// Every value must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$series",
				"$component")
			panic(message)
		}
		value, token, ok = v.parseComponent()
		if !ok {
			// No more values.
			break
		}
	}
	return values, token, true
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

// This method attempts to parse a series of values. It returns the
// series of values and whether or not the series of values was
// successfully parsed.
func (v *parser) parseSeries() (abs.Collection, *Token, bool) {
	var ok bool
	var token *Token
	var series abs.Collection
	var values abs.Sequential[abs.ComponentLike]
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return series, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		values, token, ok = v.parseInlineValues()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return series, token, false
		}
	} else {
		values, token, ok = v.parseMultilineValues()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '[' character.
			return series, token, false
		}
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("]",
			"$series",
			"$component")
		panic(message)
	}
	series = col.SeriesFromSequence(values)
	return series, token, ok
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatSeries(series abs.SeriesLike) {
	v.AppendString("[")
	var iterator = cox.Iterator[abs.ComponentLike](series)
	switch series.GetSize() {
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
