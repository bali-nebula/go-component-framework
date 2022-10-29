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
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
)

// This method attempts to parse an association between a key and value. It
// returns the association and whether or not the association was successfully
// parsed.
func (v *parser) parseAssociation() (abs.AssociationLike[any, any], *Token, bool) {
	var ok bool
	var token *Token
	var key any
	var value any
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
	var association = col.Association[any, any](key, value)
	return association, token, true
}

// This method adds the canonical format for the specified association to the
// state of the formatter.
func (v *formatter) formatAssociation(association abs.AssociationLike[any, any]) {
	var key = association.GetKey()
	v.formatAny(key)
	v.state.AppendString(": ")
	var value = association.GetValue()
	v.formatAny(value)
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abs.CatalogLike[any, any], *Token, bool) {
	var ok bool
	var token *Token
	var catalog = col.Catalog[any, any]()
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
			"$items")
		panic(message)
	}
	return catalog, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCatalog(catalog abs.CatalogLike[any, any]) {
	v.state.AppendString("[")
	switch catalog.GetSize() {
	case 0:
		v.state.AppendString(":")
	case 1:
		var association = catalog.GetItem(1)
		v.formatAssociation(association)
	default:
		var iterator = age.Iterator[abs.AssociationLike[any, any]](catalog)
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

// This method attempts to parse a collection of items. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (any, *Token, bool) {
	var ok bool
	var token *Token
	var collection any
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

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCollection(collection any) {
	rang, ok := collection.(abs.RangeLike[any])
	if ok {
		v.formatRange(rang)
		return
	}
	catalog, ok := collection.(abs.CatalogLike[any, any])
	if ok {
		v.formatCatalog(catalog)
		return
	}
	var list = collection.(abs.ListLike[any])
	v.formatList(list)
}

// This method attempts to parse a catalog collection with inline associations.
// It returns the catalog collection and whether or not the catalog collection
// was successfully parsed.
func (v *parser) parseInlineAssociations() (abs.CatalogLike[any, any], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike[any, any]
	var catalog = col.Catalog[any, any]()
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

// This method attempts to parse a list collection with inline items. It returns
// the list collection and whether or not the list collection was
// successfully parsed.
func (v *parser) parseInlineItems() (abs.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var item any
	var list = col.List[any]()
	_, token, ok = v.parseDelimiter("]")
	if ok {
		// This is an empty list.
		v.backupOne() // Put back the ']' token.
		return list, token, true
	}
	item, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one component item.
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$component",
			"$list",
			"$component")
		panic(message)
	}
	for {
		list.AddItem(item)
		// Every subsequent item must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more items.
			break
		}
		item, token, ok = v.parseComponent()
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

// This method attempts to parse a list of items. It returns the
// list of items and whether or not the list of items was
// successfully parsed.
func (v *parser) parseList() (abs.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var list = col.List[any]()
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return list, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		list, token, ok = v.parseInlineItems()
		if !ok {
			v.backupOne() // Put back the '[' character.
			return list, token, false
		}
	} else {
		list, token, ok = v.parseMultilineItems()
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
func (v *formatter) formatList(list abs.ListLike[any]) {
	v.state.AppendString("[")
	switch list.GetSize() {
	case 0:
		v.state.AppendString(" ")
	case 1:
		var item = list.GetItem(1)
		v.formatAny(item)
	default:
		var iterator = age.Iterator[any](list)
		v.state.IncrementDepth()
		for iterator.HasNext() {
			v.state.AppendNewline()
			var item = iterator.GetNext()
			v.formatAny(item)
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	}
	v.state.AppendString("]")
}

// This method attempts to parse a catalog collection with multiline associations.
// It returns the catalog collection and whether or not the catalog collection
// was successfully parsed.
func (v *parser) parseMultilineAssociations() (abs.CatalogLike[any, any], *Token, bool) {
	var ok bool
	var token *Token
	var association abs.AssociationLike[any, any]
	var catalog = col.Catalog[any, any]()
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

// This method attempts to parse a list collection with multiline items. It
// returns the list collection and whether or not the list collection was
// successfully parsed.
func (v *parser) parseMultilineItems() (abs.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var item any
	var list = col.List[any]()
	item, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one item.
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$component",
			"$list",
			"$component")
		panic(message)
	}
	for {
		list.AddItem(item)
		// Every item must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError("An unexpected token was received by the parser:", token)
			message += generateGrammar("EOL",
				"$list",
				"$component")
			panic(message)
		}
		item, token, ok = v.parseComponent()
		if !ok {
			// No more items.
			break
		}
	}
	return list, token, true
}

// This method attempts to parse a primitive. It returns the primitive and
// whether or not the primitive was successfully parsed.
func (v *parser) parsePrimitive() (any, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var primitive any
	primitive, token, ok = v.parseElement()
	if !ok {
		primitive, token, ok = v.parseString()
	}
	if !ok {
		// This must be explicitly set to nil since it is or type any.
		primitive = nil
	}
	return primitive, token, ok
}

// This method attempts to parse a range collection. It returns the range
// collection and whether or not the range collection was successfully parsed.
func (v *parser) parseRange() (abs.RangeLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var left, right string
	var first any
	var extent abs.Extent
	var last any
	var rng abs.RangeLike[any]
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
func (v *formatter) formatRange(rng abs.RangeLike[any]) {
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
