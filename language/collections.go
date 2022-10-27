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
		v.backupOne() // Put back the primitive token.
		return nil, token, false
	}
	// This must be an association.
	value, token, ok = v.parseComponent()
	if !ok {
		var message = fmt.Sprintf("Expected a component after the ':' character but received:\n%v\n\n", token)
		message += generateGrammar(
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
	var association abs.AssociationLike[any, any]
	var catalog = col.Catalog[any, any]()
	_, token, ok = v.parseEOL()
	if !ok {
		// The associations are on a single line.
		_, token, ok = v.parseDelimiter(":")
		if ok {
			// This is an empty catalog.
			return catalog, token, true
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			// A non-empty catalog must have at least one association.
			return nil, token, false
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
				var message = fmt.Sprintf(
					"Expected an association after the ',' character but received:\n%v\n\n", token)
				message += generateGrammar(
					"$catalog",
					"$association",
					"$primitive",
					"$component")
				panic(message)
			}
		}
		return catalog, token, true
	}
	// The associations are on separate lines.
	association, token, ok = v.parseAssociation()
	if !ok {
		// A non-empty catalog must have at least one association.
		v.backupOne() // Put back the EOL token.
		return nil, token, false
	}
	for {
		catalog.AddAssociation(association)
		// Every association must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = fmt.Sprintf(
				"Expected an EOL character following the association but received:\n%v\n\n", token)
			message += generateGrammar(
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

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCatalog(catalog abs.CatalogLike[any, any]) {
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
}

// This method attempts to parse a collection of items. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (any, *Token, bool) {
	var ok bool
	var token *Token
	var items any
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return nil, token, false
	}
	items, token, ok = v.parseItems()
	if !ok {
		var message = fmt.Sprintf(
			"Expected a sequence of items following the '[' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$collection",
			"$items")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = fmt.Sprintf(
			"Expected a ']' character following the sequence of items but received:\n%v\n\n", token)
		message += generateGrammar(
			"$collection",
			"$items")
		panic(message)
	}
	return items, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatCollection(collection any) {
	v.state.AppendString("[")
	v.formatItems(collection)
	v.state.AppendString("]")
}

// This method attempts to parse a list of items. It returns the
// list of items and whether or not the list of items was
// successfully parsed.
func (v *parser) parseList() (abs.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var item any
	var list = col.List[any]()
	_, token, ok = v.parseEOL()
	if !ok {
		// The items are on a single line.
		_, token, ok = v.parseDelimiter("]")
		if ok {
			// This is an empty list.
			v.backupOne() // Put back the ']' token.
			return list, token, true
		}
		item, token, ok = v.parseComponent()
		if !ok {
			// A non-empty list must have at least one component item.
			var message = fmt.Sprintf(
				"Expected a component item after the '[' character but received:\n%v\n\n", token)
			message += generateGrammar(
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
				var message = fmt.Sprintf(
					"Expected a component item after the ',' character but received:\n%v\n\n", token)
				message += generateGrammar(
					"$list",
					"$component")
				panic(message)
			}
		}
		return list, token, true
	}
	// The items are on separate lines.
	item, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one item.
		var message = fmt.Sprintf(
			"Expected a component item after the '[' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$list",
			"$component")
		panic(message)
	}
	for {
		list.AddItem(item)
		// Every item must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = fmt.Sprintf(
				"Expected an EOL character following the component item but received:\n%v\n\n", token)
			message += generateGrammar(
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

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatList(list abs.ListLike[any]) {
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

// This method attempts to parse a sequence of items. It returns the
// sequence and whether or not the sequence was successfully parsed.
func (v *parser) parseItems() (any, *Token, bool) {
	var ok bool
	var token *Token
	var items any
	items, token, ok = v.parseCatalog()
	if !ok {
		items, token, ok = v.parseRange()
	}
	if !ok {
		// The list must be attempted last since it may start with a component
		// which cannot be put back as a single token.
		items, token, ok = v.parseList() // The list must be last.
	}
	return items, token, ok
}

// This method adds the canonical format for the specified items to the
// state of the formatter.
func (v *formatter) formatItems(items any) {
	rang, ok := items.(abs.RangeLike[any])
	if ok {
		v.formatRange(rang)
		return
	}
	catalog, ok := items.(abs.CatalogLike[any, any])
	if ok {
		v.formatCatalog(catalog)
		return
	}
	var list = items.(abs.ListLike[any])
	v.formatList(list)
}

// This method attempts to parse a range collection. It returns the range
// collection and whether or not the range collection was successfully parsed.
func (v *parser) parseRange() (abs.RangeLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var first any
	var extent abs.Extent
	var last any
	first, token, _ = v.parsePrimitive() // The first value is optional.
	delimiter, token, ok = v.parseDelimiter("..")
	if !ok {
		delimiter, token, ok = v.parseDelimiter("<..")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("<..<")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("..<")
	}
	if !ok {
		// This is not a range collection.
		if first != nil {
			v.backupOne() // Put back the Value token.
		}
		return nil, token, false
	}
	switch delimiter {
	case "..":
		extent = abs.INCLUSIVE
	case "<..":
		extent = abs.RIGHT
	case "..<":
		extent = abs.LEFT
	case "<..<":
		extent = abs.EXCLUSIVE
	default:
		panic(fmt.Sprintf("The range contains an unknown extent: %v", delimiter))
	}
	last, token, _ = v.parsePrimitive() // The last value is optional.
	var rng = col.Range(first, extent, last)
	return rng, token, true
}

// This method adds the canonical format for the specified collection to the
// state of the formatter.
func (v *formatter) formatRange(r abs.RangeLike[any]) {
	var first = r.GetFirst()
	if first != nil {
		v.formatAny(first)
	}
	var extent = r.GetExtent()
	var delimiter string
	switch extent {
	case abs.INCLUSIVE:
		delimiter = ".."
	case abs.RIGHT:
		delimiter = "<.."
	case abs.LEFT:
		delimiter = "..<"
	case abs.EXCLUSIVE:
		delimiter = "<..<"
	default:
		panic(fmt.Sprintf("The range contains an unknown extent: %v", extent))
	}
	v.state.AppendString(delimiter)
	var last = r.GetLast()
	if last != nil {
		v.formatAny(last)
	}
}
