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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// SYMBOL INTERFACE

// This constructor attempts to create a new symbol from the specified
// formatted string. It returns a symbol value and whether or not the string
// contained a valid symbol.
// For valid string formats for this type see `../abstractions/language.go`.
func SymbolFromString(v string) (Symbol, bool) {
	var ok = true
	var symbol string
	var matches = abstractions.ScanSymbol([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		symbol = matches[0]
	}
	return Symbol(symbol), ok
}

// This constructor attempts to create a new symbol string from the specified
// array of runes. It returns the corresponding symbol string.
func SymbolFromRunes(v []rune) Symbol {
	var symbol, ok = SymbolFromString(`$` + string(v))
	if !ok {
		panic(fmt.Sprintf("The runes contain an illegal character: %v", string(v)))
	}
	return symbol
}

// This type defines the methods associated with a symbol string that
// extends the native Go string type and represents the string of runes that
// make up the symbol.
type Symbol string

// LEXICAL INTERFACE

// This method returns the canonical string for this string.
func (v Symbol) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Symbol) String() string {
	return v.AsString()
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Symbol) IsEmpty() bool {
	return false // A symbol must have at least one character.
}

// This method returns the number of runes contained in this string.
func (v Symbol) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the runes in this string. The runes retrieved are in
// the same order as they are in the string.
func (v Symbol) AsArray() []rune {
	return []rune(v[1:]) // Remove the leading dollar character.
}

// INDEXED INTERFACE

// This method retrieves from this string the rune that is associated with the
// specified index.
func (v Symbol) GetItem(index int) rune {
	var runes = v.AsArray()
	var length = len(runes)
	index = abstractions.NormalizedIndex(index, length)
	return runes[index]
}

// This method retrieves from this string all runes from the first index through
// the last index (inclusive).
func (v Symbol) GetItems(first int, last int) []rune {
	var runes = v.AsArray()
	var length = len(runes)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	return runes[first : last+1]
}

// This method returns the index of the FIRST occurence of the specified rune in
// this string, or zero if this string does not contain the rune.
func (v Symbol) GetIndex(r rune) int {
	var runes = v.AsArray()
	for index, candidate := range runes {
		if candidate == r {
			// Found the rune.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The rune was not found.
	return 0
}

// SYMBOLS LIBRARY

// This singleton creates a unique name space for the library functions for
// symbols.
var Symbols = &symbols{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for symbols.
type symbols struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified symbol
// strings.
func (f *symbols) Concatenate(first Symbol, second Symbol) Symbol {
	var symbol = first + second[1:]
	return Symbol(symbol)
}
