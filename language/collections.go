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
func (v *parser) parseAssociation() (abstractions.AssociationLike[any, any], bool) {
	var ok bool
	var key any
	var value any
	key, ok = v.parsePrimitive()
	if !ok {
		// This is not an association.
		return nil, false
	}
	_, ok = v.parseDelimiter(":")
	if !ok {
		// This is not an association.
		v.backupOne() // Put back the primitive token.
		return nil, false
	}
	// This must be an association.
	value, ok = v.parseComponent()
	if !ok {
		panic("Expected a value after the ':' character.")
	}
	var association = collections.Association[any, any](key, value)
	return association, true
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abstractions.CatalogLike[any, any], bool) {
	var ok bool
	var association abstractions.AssociationLike[any, any]
	var catalog = collections.Catalog[any, any]()
	_, ok = v.parseEOL()
	if !ok {
		// The associations are on a single line.
		_, ok = v.parseDelimiter(":")
		if ok {
			// This is an empty catalog.
			return catalog, true
		}
		association, ok = v.parseAssociation()
		if !ok {
			// A non-empty catalog must have at least one association.
			return nil, false
		}
		for {
			catalog.AddAssociation(association)
			// Every subsequent association must be preceded by a ','.
			_, ok = v.parseDelimiter(",")
			if !ok {
				// No more associations.
				break
			}
			association, ok = v.parseAssociation()
			if !ok {
				panic("Expected an association after the ',' character.")
			}
		}
		return catalog, true
	}
	// The associations are on separate lines.
	association, ok = v.parseAssociation()
	if !ok {
		// A non-empty catalog must have at least one association.
		v.backupOne() // Put back the EOL token.
		return nil, false
	}
	for {
		catalog.AddAssociation(association)
		// Every association must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the association.")
		}
		association, ok = v.parseAssociation()
		if !ok {
			// No more associations.
			break
		}
	}
	return catalog, true
}

// This method attempts to parse a collection of items. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (any, bool) {
	var ok bool
	var bad *token
	var sequence any
	_, ok = v.parseDelimiter("[")
	if !ok {
		return nil, false
	}
	sequence, ok = v.parseSequence()
	if !ok {
		panic("Expected a sequence following the '[' character.")
	}
	bad, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the sequence and received: %v", *bad))
	}
	return sequence, true
}

// This method attempts to parse a list of items. It returns the
// list of items and whether or not the list of items was
// successfully parsed.
func (v *parser) parseList() (abstractions.ListLike[any], bool) {
	var ok bool
	var item any
	var list = collections.List[any]()
	_, ok = v.parseEOL()
	if !ok {
		// The items are on a single line.
		_, ok = v.parseDelimiter("]")
		if ok {
			// This is an empty list.
			v.backupOne() // Put back the ']' token.
			return list, true
		}
		item, ok = v.parseComponent()
		if !ok {
			// A non-empty list must have at least one item.
			panic("Expected an item after the '[' character.")
		}
		for {
			list.AddItem(item)
			// Every subsequent item must be preceded by a ','.
			_, ok = v.parseDelimiter(",")
			if !ok {
				// No more items.
				break
			}
			item, ok = v.parseComponent()
			if !ok {
				panic("Expected an item after the ',' character.")
			}
		}
		return list, true
	}
	// The items are on separate lines.
	item, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one item.
		panic("Expected an item after the '[' character.")
	}
	for {
		list.AddItem(item)
		// Every item must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the item.")
		}
		item, ok = v.parseComponent()
		if !ok {
			// No more items.
			break
		}
	}
	return list, true
}

// This method attempts to parse a primitive. It returns the primitive and
// whether or not the primitive was successfully parsed.
func (v *parser) parsePrimitive() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var primitive any
	primitive, ok = v.parseElement()
	if !ok {
		primitive, ok = v.parseString()
	}
	if !ok {
		// This must be explicitly set to nil since it is or type any.
		primitive = nil
	}
	return primitive, ok
}

// This method attempts to parse a sequence of items. It returns the
// sequence and whether or not the sequence was successfully parsed.
func (v *parser) parseSequence() (any, bool) {
	var ok bool
	var sequence any
	sequence, ok = v.parseCatalog()
	if !ok {
		sequence, ok = v.parseSlice()
	}
	if !ok {
		// The list must be attempted last since it may start with a component
		// which cannot be put back as a single token.
		sequence, ok = v.parseList() // The list must be last.
	}
	return sequence, ok
}

// This method attempts to parse a slice collection. It returns the slice
// collection and whether or not the slice collection was successfully parsed.
func (v *parser) parseSlice() (abstractions.SliceLike[any], bool) {
	var ok bool
	var t *token
	var first any
	var last any
	first, _ = v.parseValue() // The first value is optional.
	t, ok = v.parseDelimiter("..")
	if !ok {
		t, ok = v.parseDelimiter("<..")
	}
	if !ok {
		t, ok = v.parseDelimiter("<..<")
	}
	if !ok {
		t, ok = v.parseDelimiter("..<")
	}
	if !ok {
		// This is not a slice collection.
		if first != nil {
			v.backupOne() // Put back the Value token.
		}
		return nil, false
	}
	var connector = t.val
	last, _ = v.parseValue() // The last value is optional.
	var slice = collections.Slice(first, connector, last)
	return slice, true
}

// This method attempts to parse a primitive value. It returns the primitive
// value and whether or not the primitive value was successfully parsed.
func (v *parser) parseValue() (any, bool) {
	var ok bool
	var value any
	value, ok = v.parseElement()
	if !ok {
		value, ok = v.parseString()
	}
	if !ok {
		value, ok = v.parseIdentifier()
	}
	if !ok {
		// This must be explicitly set to nil since it is or type any.
		value = nil
	}
	return value, ok
}
