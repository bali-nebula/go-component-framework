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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
)

// This method attempts to parse an association between a key and value. It
// returns the association and whether or not the association was successfully
// parsed.
func (v *parser) parseAssociation() (abs.AssociationLike[abs.KeyLike, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var key abs.KeyLike
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
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$component",
			"$association",
			"$primitive",
			"$component")
		panic(message)
	}
	var association = col.Association[abs.KeyLike, abs.ComponentLike](key, value)
	return association, token, true
}

// This method adds the canonical format for the specified association to the
// state of the formatter.
func (v *formatter) formatAssociation(association abs.AssociationLike[abs.KeyLike, abs.ComponentLike]) {
	var key = association.GetKey()
	v.formatAny(key)
	v.state.AppendString(": ")
	var value = association.GetValue()
	v.formatComponent(value)
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abs.CatalogLike[abs.KeyLike, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var catalog = col.Catalog[abs.KeyLike, abs.ComponentLike]()
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
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("]",
			"$collection",
			"$values")
		panic(message)
	}
	return catalog, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCatalog(catalog abs.CatalogLike[abs.KeyLike, abs.ComponentLike]) {
	v.state.AppendString("[")
	var iterator = age.Iterator[abs.AssociationLike[abs.KeyLike, abs.ComponentLike]](catalog)
	switch catalog.GetSize() {
	case 0:
		v.state.AppendString(":")
	case 1:
		var association = iterator.GetNext()
		v.formatAssociation(association)
	default:
		v.state.IncrementDepth()
		for iterator.HasNext() {
			v.state.AppendNewline()
			var association = iterator.GetNext()
			v.formatAssociation(association)
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	}
	v.state.AppendString("]")
}

// This method attempts to parse a collection of values. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (abs.SequenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var collection abs.SequenceLike
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
func (v *parser) parseInlineAssociations() (abs.CatalogLike[abs.KeyLike, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike[abs.KeyLike, abs.ComponentLike]
	var catalog = col.Catalog[abs.KeyLike, abs.ComponentLike]()
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
			var message = v.formatError("An unexpected token was received by the parser:", token)
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
func (v *parser) parseInlineValues() (abs.ListLike[abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var list = col.List[abs.ComponentLike]()
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty list.
		v.backupOne() // Put back the ']' token.
		return list, token, true
	}
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one component value.
		var message = v.formatError("An unexpected token was received by the parser:", token)
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
			var message = v.formatError("An unexpected token was received by the parser:", token)
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
func (v *parser) parseList() (abs.ListLike[abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var list = col.List[abs.ComponentLike]()
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
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("]",
			"$list",
			"$component")
		panic(message)
	}
	return list, token, ok
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatList(list abs.ListLike[abs.ComponentLike]) {
	v.state.AppendString("[")
	switch list.GetSize() {
	case 0:
		v.state.AppendString(" ")
	case 1:
		var value = list.GetValue(1)
		v.formatAny(value)
	default:
		var iterator = age.Iterator[abs.ComponentLike](list)
		v.state.IncrementDepth()
		for iterator.HasNext() {
			v.state.AppendNewline()
			var value = iterator.GetNext()
			v.formatAny(value)
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	}
	v.state.AppendString("]")
}

// This method attempts to parse a catalog collection with multiline associations.
// It returns the catalog collection and whether or not the catalog collection
// was successfully parsed.
func (v *parser) parseMultilineAssociations() (abs.CatalogLike[abs.KeyLike, abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike[abs.KeyLike, abs.ComponentLike]
	var catalog = col.Catalog[abs.KeyLike, abs.ComponentLike]()
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
			var message = v.formatError("An unexpected token was received by the parser:", token)
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
func (v *parser) parseMultilineValues() (abs.ListLike[abs.ComponentLike], *Token, bool) {
	var ok bool
	var token *Token
	var value abs.ComponentLike
	var list = col.List[abs.ComponentLike]()
	value, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one value.
		var message = v.formatError("An unexpected token was received by the parser:", token)
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
			var message = v.formatError("An unexpected token was received by the parser:", token)
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

// This method attempts to parse a primitive. It returns the primitive and
// whether or not the primitive was successfully parsed.
func (v *parser) parsePrimitive() (abs.ValueLike, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var primitive abs.ValueLike
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

// This method attempts to parse a range collection. It returns the range
// collection and whether or not the range collection was successfully parsed.
func (v *parser) parseRange() (abs.RangeLike[abs.ValueLike], *Token, bool) {
	var ok bool
	var token *Token
	var left, right string
	var first abs.ValueLike
	var extent abs.Extent
	var last abs.ValueLike
	var rng abs.RangeLike[abs.ValueLike]
	left, token, ok = v.parseDelimiter("[")
	if !ok {
		left, token, ok = v.parseDelimiter("(")
		if !ok {
			return rng, token, false
		}
	}
	first, token, _ = v.parsePrimitive() // The first value is optional.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a range collection.
		if first != nil {
			v.backupOne() // Put back the first value token.
		}
		v.backupOne() // Put back the left bracket character.
		return rng, token, false
	}
	last, token, _ = v.parsePrimitive() // The last value is optional.
	right, token, ok = v.parseDelimiter("]")
	if !ok {
		right, token, ok = v.parseDelimiter(")")
		if !ok {
			var message = v.formatError("An unexpected token was received by the parser:", token)
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
	rng = col.Range(first, extent, last)
	return rng, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatRange(rng abs.RangeLike[abs.ValueLike]) {
	var extent = rng.GetExtent()
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
	v.state.AppendString(left)
	var first = rng.GetFirst()
	if first != nil {
		v.formatAny(first)
	}
	v.state.AppendString("..")
	var last = rng.GetLast()
	if last != nil {
		v.formatAny(last)
	}
	v.state.AppendString(right)
}
