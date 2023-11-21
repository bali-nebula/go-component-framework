/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
	reg "regexp"
)

// QUOTE STRING IMPLEMENTATION

// These constants are used to define a regular expression for matching
// quote strings. NOTE: Since the string does not contain the delimiting
// quotes, any embedded double quote need not be escaped. This is different
// than the regular expression the scanner uses for quoted strings.
const (
	base16    = `[0-9a-f]`
	character = `(?:` + escape + `|[^\f\r\n\t]` + `)`
	unicode   = `u` + base16 + `{4}`
	escape    = `\\(?:` + unicode + `|[frnt\\])` // No escaped double quote.
	quote     = `(` + character + `*)`
)

// This scanner is used for matching quote strings.
var quoteScanner = reg.MustCompile(`^(?:` + quote + `)$`)

// This constructor creates a new quote string from the specified string.
func QuoteFromString(string_ string) abs.QuoteLike {
	if !quoteScanner.MatchString(string_) {
		var message = fmt.Sprintf("Attempted to construct a quote string from an invalid string: %v", string_)
		panic(message)
	}
	return Quote(string_)
}

// This constructor attempts to create a new quote string from the specified
// array of runes. It returns a quote string and whether or not the
// resulting string contained a valid quote.
func QuoteFromArray(array []rune) abs.QuoteLike {
	var quote = Quote(string(array))
	return quote
}

// This constructor creates a new quote string from the specified sequence of
// runes. It returns the corresponding quote string.
func QuoteFromSequence(sequence abs.Sequential[rune]) abs.QuoteLike {
	var array = sequence.AsArray()
	var quote = QuoteFromArray(array)
	return quote
}

// This type defines the methods associated with a quote string that extends
// the native Go string type and represents the string of runes that make up
// the quote string.
type Quote string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
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
	var runes = []rune(v)
	var array = make([]rune, len(runes))
	for index, r := range runes {
		array[index] = rune(r)
	}
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the rune that is associated
// with the specified index.
func (v Quote) GetValue(index int) rune {
	var array = v.AsArray()
	var runes = col.Array[rune](array)
	return runes.GetValue(index)
}

// This method retrieves from this string all runes from the first index
// through the last index (inclusive).
func (v Quote) GetValues(first int, last int) abs.Sequential[rune] {
	var array = v.AsArray()
	var runes = col.Array[rune](array)
	return runes.GetValues(first, last)
}

// QUOTE LIBRARY

// This singleton creates a unique name space for the library functions for
// quote strings.
var Quote = &quotes_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for quote strings.
type quotes_ struct{}

// This function returns the concatenation of the two specified quote strings.
func (l *quotes) Concatenate(first, second abs.QuoteLike) abs.QuoteLike {
	return Quote(first.AsString() + second.AsString())
}
