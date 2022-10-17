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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
)

// This method attempts to parse an association between a key and value. It
// returns the association and whether or not the association was successfully
// parsed.
func (v *parser) parseAssociation() (abstractions.AssociationLike[any, any], *Token, bool) {
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
		panic("Expected a value after the ':' character.")
	}
	var association = collections.Association[any, any](key, value)
	return association, token, true
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abstractions.CatalogLike[any, any], *Token, bool) {
	var ok bool
	var token *Token
	var association abstractions.AssociationLike[any, any]
	var catalog = collections.Catalog[any, any]()
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
				panic("Expected an association after the ',' character.")
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
			panic("Expected an EOL character following the association.")
		}
		association, token, ok = v.parseAssociation()
		if !ok {
			// No more associations.
			break
		}
	}
	return catalog, token, true
}

// This method attempts to parse a collection of items. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (any, *Token, bool) {
	var ok bool
	var token *Token
	var sequence any
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		return nil, token, false
	}
	sequence, token, ok = v.parseSequence()
	if !ok {
		panic("Expected a sequence following the '[' character.")
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the sequence and received: %v", *token))
	}
	return sequence, token, true
}

// This method attempts to parse a list of items. It returns the
// list of items and whether or not the list of items was
// successfully parsed.
func (v *parser) parseList() (abstractions.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var item any
	var list = collections.List[any]()
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
			// A non-empty list must have at least one item.
			panic("Expected an item after the '[' character.")
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
				panic("Expected an item after the ',' character.")
			}
		}
		return list, token, true
	}
	// The items are on separate lines.
	item, token, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one item.
		panic("Expected an item after the '[' character.")
	}
	for {
		list.AddItem(item)
		// Every item must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the item.")
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

// This method attempts to parse a sequence of items. It returns the
// sequence and whether or not the sequence was successfully parsed.
func (v *parser) parseSequence() (any, *Token, bool) {
	var ok bool
	var token *Token
	var sequence any
	sequence, token, ok = v.parseCatalog()
	if !ok {
		sequence, token, ok = v.parseRange()
	}
	if !ok {
		// The list must be attempted last since it may start with a component
		// which cannot be put back as a single token.
		sequence, token, ok = v.parseList() // The list must be last.
	}
	return sequence, token, ok
}

// This method attempts to parse a range collection. It returns the range
// collection and whether or not the range collection was successfully parsed.
func (v *parser) parseRange() (abstractions.RangeLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var first any
	var connector string
	var last any
	first, token, _ = v.parsePrimitive() // The first value is optional.
	connector, token, ok = v.parseDelimiter("..")
	if !ok {
		connector, token, ok = v.parseDelimiter("<..")
	}
	if !ok {
		connector, token, ok = v.parseDelimiter("<..<")
	}
	if !ok {
		connector, token, ok = v.parseDelimiter("..<")
	}
	if !ok {
		// This is not a range collection.
		if first != nil {
			v.backupOne() // Put back the Value token.
		}
		return nil, token, false
	}
	last, token, _ = v.parsePrimitive() // The last value is optional.
	var rng = collections.Range(first, connector, last)
	return rng, token, true
}
