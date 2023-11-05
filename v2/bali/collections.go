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
	com "github.com/bali-nebula/go-component-framework/v2/components"
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
	_, _, ok = v.parseDelimiter(":")
	if !ok {
		// This is not an association.
		v.backupOne(token) // Put back the primitive key token.
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
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not a collection.
		return collection, token, false
	}
	collection, _, ok = v.parseAssociations()
	if !ok {
		// The values must be attempted second since it may start with a component
		// which cannot be put back as a single token.
		collection, _, ok = v.parseValues()
		if !ok {
			// Not a collection (must be a range).
			v.backupOne(token) // Put back the "[" character.
			return collection, token, false
		}
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("]",
			"$collection",
			"$associations",
			"$association",
		)
		panic(message)
	}
	return collection, token, true
}

// It returns the associations collection and whether or not the associations collection
// was successfully parsed.
func (v *parser) parseInlineAssociations() (abs.AssociationsLike, *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var associations = col.Catalog()
	association, token, ok = v.parseAssociation()
	if !ok {
		// This is not an inline associations.
		return associations, token, false
	}
	for {
		var key = association.GetKey()
		var value = association.GetValue()
		associations.SetValue(key, value)
		// Every subsequent association must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more associations.
			return associations, token, true
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("association",
				"$collection",
				"$associations",
				"$association",
			)
			panic(message)
		}
	}
}

// This method attempts to parse a values collection with inline values. It
// returns the values collection and whether or not the values collection was
// successfully parsed.
func (v *parser) parseInlineValues() (abs.ValuesLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var values = col.List()
	value, token, ok = v.parseComponent()
	if !ok {
		// This is not an inline values.
		return values, token, false
	}
	for {
		values.AddValue(value)
		// Every subsequent value must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more values.
			return values, token, true
		}
		value, token, ok = v.parseComponent()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("value",
				"$collection",
				"$values",
				"$value",
			)
			panic(message)
		}
	}
}

// This method attempts to parse an associations collection. It returns the
// associations collection and whether or not the associations collection was
// successfully parsed.
func (v *parser) parseAssociations() (abs.AssociationsLike, *Token, bool) {
	var ok bool
	var token *Token
	var associations abs.AssociationsLike
	_, token, ok = v.parseDelimiter(":")
	if ok {
		// The associations is empty.
		associations = col.Catalog()
		return associations, token, true
	}
	_, token, ok = v.parseEOL()
	if ok {
		associations, _, ok = v.parseMultilineAssociations()
		if !ok {
			v.backupOne(token) // Put back the EOL character.
			return associations, token, false
		}
	} else {
		associations, token, ok = v.parseInlineAssociations()
		if !ok {
			return associations, token, false
		}
	}
	return associations, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatAssociations(associations abs.AssociationsLike) {
	v.AppendString("[")
	var iterator = col.AssociationIterator(associations)
	switch associations.GetSize() {
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

// This method attempts to parse an associations collection with multiline associations.
// It returns the associations collection and whether or not the associations collection
// was successfully parsed.
func (v *parser) parseMultilineAssociations() (abs.AssociationsLike, *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike
	var associations = col.Catalog()

	association, token, ok = v.parseAssociation()
	if !ok {
		// This is not a multiline associations.
		return associations, token, false
	}
	// Every association must be followed by an EOL.
	_, token, ok = v.parseEOL()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("EOL",
			"$collection",
			"$associations",
			"$association",
		)
		panic(message)
	}
	for {
		var key = association.GetKey()
		var value = association.GetValue()
		associations.SetValue(key, value)
		association, token, ok = v.parseAssociation()
		if !ok {
			// No more associations.
			return associations, token, true
		}
		// Every association must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$collection",
				"$associations",
				"$association",
			)
			panic(message)
		}
	}
}

// This method attempts to parse a values collection with multiline values.
// It returns the values collection and whether or not the values collection was
// successfully parsed.
func (v *parser) parseMultilineValues() (abs.ValuesLike, *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var values = col.List()

	value, token, ok = v.parseComponent()
	if !ok {
		// This is not a multiline values.
		return values, token, false
	}
	// Every value must be followed by an EOL.
	_, token, ok = v.parseEOL()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("EOL",
			"$collection",
			"$values",
			"$value",
		)
		panic(message)
	}
	for {
		values.AddValue(value)
		value, token, ok = v.parseComponent()
		if !ok {
			// No more values.
			return values, token, true
		}
		// Every value must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$collection",
				"$values",
				"$value",
			)
			panic(message)
		}
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

// This method attempts to parse a values of values. It returns the
// values of values and whether or not the values of values was
// successfully parsed.
func (v *parser) parseValues() (abs.ValuesLike, *Token, bool) {
	var ok bool
	var token *Token
	var values abs.ValuesLike
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// The values is empty.
		v.backupOne(token) // Put back the "]" character.
		values = col.List()
		return values, token, true
	}
	_, token, ok = v.parseEOL()
	if ok {
		values, _, ok = v.parseMultilineValues()
		if !ok {
			v.backupOne(token) // Put back the EOL character.
			return values, token, false
		}
	} else {
		values, token, ok = v.parseInlineValues()
		if !ok {
			return values, token, false
		}
	}
	return values, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatValues(values abs.ValuesLike) {
	v.AppendString("[")
	var iterator = com.ComponentIterator(values)
	switch values.GetSize() {
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
