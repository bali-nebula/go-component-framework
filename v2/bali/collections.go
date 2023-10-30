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
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	cox "github.com/craterdog/go-collection-framework/v2"
)

// TYPE DEFINITIONS

type (
	Components   []any
	Associations [][2]any
)

// UNIVERSAL CONSTRUCTORS

// This function parses a catalog from a source string.
func Catalog(source string) abs.CatalogLike {
	return ParseEntity(source).(abs.CatalogLike)
}

// This constructor returns a new catalog component initialized with the
// specified associations and parameterized with the specified context.
func CatalogWithContext(associations Associations, context abs.ContextLike) abs.ComponentLike {
	var catalog = col.Catalog()
	for _, association := range associations {
		var key = Component(association[0]).GetEntity()
		var value = Component(association[1])
		catalog.SetValue(key, value)
	}
	return ComponentWithContext(catalog, context)
}

// This function parses a catalog from a source string.
func List(source string) abs.ListLike {
	return ParseEntity(source).(abs.ListLike)
}

// This constructor returns a new list component initialized with the
// specified associations and parameterized with the specified context.
func ListWithContext(components Components, context abs.ContextLike) abs.ComponentLike {
	var component abs.ComponentLike
	var list = col.List()
	for _, value := range components {
		component = Component(value)
		list.AddValue(component)
	}
	return ComponentWithContext(list, context)
}

// This function parses a catalog from a source string.
func Queue(source string) abs.QueueLike {
	return ParseEntity(source).(abs.QueueLike)
}

// This constructor returns a new queue component with the specified capacity
// and parameterized with the specified context.
func QueueWithContext(capacity int, context abs.ContextLike) abs.ComponentLike {
	var queue = col.QueueWithCapacity(capacity)
	return ComponentWithContext(queue, context)
}

// This function parses a catalog from a source string.
func Set(source string) abs.SetLike {
	return ParseEntity(source).(abs.SetLike)
}

// This constructor returns a new set component initialized with the
// specified associations and parameterized with the specified context.
func SetWithContext(components Components, context abs.ContextLike) abs.ComponentLike {
	var component abs.ComponentLike
	var set = col.Set()
	for _, value := range components {
		component = Component(value)
		set.AddValue(component)
	}
	return ComponentWithContext(set, context)
}

// This function parses a catalog from a source string.
func Stack(source string) abs.StackLike {
	return ParseEntity(source).(abs.StackLike)
}

// This constructor returns a new stack component with the specified capacity
// and parameterized with the specified context.
func StackWithContext(capacity int, context abs.ContextLike) abs.ComponentLike {
	var stack = col.StackWithCapacity(capacity)
	return ComponentWithContext(stack, context)
}

// PRIVATE METHODS

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
		message += generateGrammar("value",
			"$association",
			"$key",
			"$value")
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

// This method attempts to parse a collection of components. It returns the
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
func (v *parser) parseInlineMapping() (abs.MappingLike, *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var mapping = cox.List[abs.AssociationLike]()
	_, token, ok = v.parseDelimiter(":")
	if ok {
		// This is an empty mapping.
		return mapping, token, true
	}
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty series.
		v.backupOne() // Put back the ']' character.
		return mapping, token, false
	}
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty mapping must have at least one association.
		return mapping, token, false
	}
	for {
		mapping.AddValue(association)
		// Every subsequent association must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more associations.
			break
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("association",
				"$associations",
				"$association",
				"$key",
				"$value")
			panic(message)
		}
	}
	return mapping, token, true
}

// This method attempts to parse a series collection with inline components. It
// returns the series collection and whether or not the series collection was
// successfully parsed.
func (v *parser) parseInlineSeries() (abs.SeriesLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var series = col.List()
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty series.
		v.backupOne() // Put back the ']' token.
		return series, token, true
	}
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty series must have at least one component value.
		var message = v.formatError(token)
		message += generateGrammar("component",
			"$collection",
			"$components",
			"$component")
		panic(message)
	}
	for {
		series.AddValue(value)
		// Every subsequent value must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more components.
			break
		}
		value, token, ok = v.parseComponent()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("component",
				"$collection",
				"$components",
				"$component")
			panic(message)
		}
	}
	return series, token, true
}

// This method attempts to parse a mapping collection. It returns the
// mapping collection and whether or not the mapping collection was
// successfully parsed.
func (v *parser) parseMapping() (abs.MappingLike, *Token, bool) {
	var ok bool
	var token *Token
	var mapping abs.MappingLike
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return mapping, token, false
	}
	_, _, ok = v.parseEOL()
	if !ok {
		mapping, token, ok = v.parseInlineMapping()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return mapping, token, false
		}
	} else {
		mapping, token, ok = v.parseMultilineMapping()
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
			"$associations",
			"$association")
		panic(message)
	}
	return mapping, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatMapping(mapping abs.MappingLike) {
	v.AppendString("[")
	var iterator = col.AssociationIterator(mapping)
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
func (v *parser) parseMultilineMapping() (abs.MappingLike, *Token, bool) {
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
				"$collection",
				"$associations",
				"$association")
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

// This method attempts to parse a series collection with multiline components.
// It returns the series collection and whether or not the series collection was
// successfully parsed.
func (v *parser) parseMultilineSeries() (abs.SeriesLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var series = col.List()
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty series must have at least one value.
		var message = v.formatError(token)
		message += generateGrammar("component",
			"$collection",
			"$components",
			"$component")
		panic(message)
	}
	for {
		series.AddValue(value)
		// Every value must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$collection",
				"$components")
			panic(message)
		}
		value, token, ok = v.parseComponent()
		if !ok {
			// No more components.
			break
		}
	}
	return series, token, true
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
	// The order of these cases is very important since Go only compares the
	// set of methods supported by each interface. An interface that is a subset
	// of another interface must be checked AFTER that interface.
	case abs.BinaryLike:
		v.formatBinary(value)
	case abs.BytecodeLike:
		v.formatBytecode(value)
	case abs.MonikerLike:
		v.formatMoniker(value)
	case abs.NarrativeLike:
		v.formatNarrative(value)
	case abs.QuoteLike:
		v.formatQuote(value)
	case abs.VersionLike:
		v.formatVersion(value)
	case abs.DurationLike:
		v.formatDuration(value)
	case abs.MomentLike:
		v.formatMoment(value)
	case abs.NumberLike:
		v.formatNumber(value)
	case abs.PercentageLike:
		v.formatPercentage(value)
	case abs.ProbabilityLike:
		v.formatProbability(value)
	case abs.AngleLike:
		v.formatAngle(value)
	case abs.BooleanLike:
		v.formatBoolean(value)
	case abs.PatternLike:
		v.formatPattern(value)
	case abs.ResourceLike:
		v.formatResource(value)
	case abs.TagLike:
		v.formatTag(value)
	case abs.SymbolLike:
		v.formatSymbol(value)
	default:
		panic(fmt.Sprintf("An invalid primitive (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a series of components. It returns the
// series of components and whether or not the series of components was
// successfully parsed.
func (v *parser) parseSeries() (abs.SeriesLike, *Token, bool) {
	var ok bool
	var token *Token
	var series abs.SeriesLike
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return series, token, false
	}
	_, _, ok = v.parseEOL()
	if !ok {
		series, token, ok = v.parseInlineSeries()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return series, token, false
		}
	} else {
		series, token, ok = v.parseMultilineSeries()
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
			"$collection",
			"$components")
		panic(message)
	}
	return series, token, ok
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatSeries(series abs.SeriesLike) {
	v.AppendString("[")
	var iterator = col.ComponentIterator(series)
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
