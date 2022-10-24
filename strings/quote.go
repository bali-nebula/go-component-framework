/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// QUOTE STRING INTERFACE

// This constructor attempts to create a new quote string from the specified
// array of runes. It returns a quote string and whether or not the
// resulting string contained a valid quote.
func QuoteFromRunes(v []rune) Quote {
	var quote = Quote(string(v))
	return quote
}

// This type defines the methods associated with a quote string that extends
// the native Go string type and represents the string of runes that make up
// the quote string.
type Quote string

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Quote) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of runes contained in this string.
func (v Quote) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the runes in this string. The runes retrieved
// are in the same order as they are in the string.
func (v Quote) AsArray() []rune {
	return []rune(v)
}

// INDEXED INTERFACE

// This method retrieves from this string the rune that is associated
// with the specified index.
func (v Quote) GetItem(index int) rune {
	var runes = v.AsArray()
	var length = len(runes)
	index = abs.NormalizedIndex(index, length)
	return runes[index]
}

// This method retrieves from this string all runes from the first index
// through the last index (inclusive).
func (v Quote) GetItems(first int, last int) []rune {
	var runes = v.AsArray()
	var length = len(runes)
	first = abs.NormalizedIndex(first, length)
	last = abs.NormalizedIndex(last, length)
	return runes[first : last+1]
}

// This method returns the index of the FIRST occurence of the specified rune
// in this string, or zero if this string does not contain the rune.
func (v Quote) GetIndex(b rune) int {
	var runes = v.AsArray()
	for index, candidate := range runes {
		if candidate == b {
			// Found the rune.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The rune was not found.
	return 0
}

// QUOTES LIBRARY

// This singleton creates a unique name space for the library functions for
// quotes.
var Quotes = &quotes{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for quotes.
type quotes struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified quote
// strings.
func (l *quotes) Concatenate(first Quote, second Quote) Quote {
	return Quote(first + second)
}
