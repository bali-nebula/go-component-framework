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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	cox "github.com/craterdog/go-collection-framework"
)

// QUOTE STRING INTERFACE

// This constructor attempts to create a new quote string from the specified
// array of runes. It returns a quote string and whether or not the
// resulting string contained a valid quote.
func QuoteFromArray(array []rune) Quote {
	var quote = Quote(string(array))
	return quote
}

// This constructor creates a new quote string from the specified sequence of
// runes. It returns the corresponding quote string.
func QuoteFromSequence(sequence abs.Sequential[rune]) Quote {
	var array = sequence.AsArray()
	var quote = QuoteFromArray(array)
	return quote
}

// This type defines the methods associated with a quote string that extends
// the native Go string type and represents the string of runes that make up
// the quote string.
type Quote string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
func (v Quote) AsString() string {
	return string(v)
}

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
func (v Quote) GetValue(index int) rune {
	var array = v.AsArray()
	var runes = cox.Array[rune](array)
	return runes.GetValue(index)
}

// This method retrieves from this string all runes from the first index
// through the last index (inclusive).
func (v Quote) GetValues(first int, last int) abs.Sequential[rune] {
	var array = v.AsArray()
	var runes = cox.Array[rune](array)
	return runes.GetValues(first, last)
}

// This method returns the index of the FIRST occurence of the specified rune
// in this string, or zero if this string does not contain the rune.
func (v Quote) GetIndex(r rune) int {
	var array = v.AsArray()
	var runes = cox.Array[rune](array)
	return runes.GetIndex(r)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all quote string
// library functions.
const Quotes quotes = false

// This type defines the library functions associated with quote strings.
type quotes bool

// This function returns the concatenation of the two specified quote strings.
func (l quotes) Concatenate(first Quote, second Quote) Quote {
	return Quote(first + second)
}
