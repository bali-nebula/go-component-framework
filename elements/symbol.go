/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	reg "regexp"
)

// SYMBOL IMPLEMENTATION

// These constants are used to define a regular expression for matching
// symbols.
const (
	letter     = `\pL` // All unicode letters and connectors like underscores.
	digit      = `\pN` // All unicode digits.
	identifier = letter + `(?:` + letter + `|` + digit + `)*`
	ordinal    = `[1-9][0-9]*`
	symbol     = `(` + identifier + `(?:-` + ordinal + `)?)`
)

// This scanner is used for matching symbol strings.
var symbolScanner = reg.MustCompile(`^(?:` + symbol + `)$`)

// This constructor creates a new symbol from the specified string.
func SymbolFromString(string_ string) abs.SymbolLike {
	if !symbolScanner.MatchString(string_) {
		var message = fmt.Sprintf("Attempted to construct a symbol from an invalid string: %v", string_)
		panic(message)
	}
	return Symbol(string_)
}

// This type defines the methods associated with a symbol element that
// extends the native Go string type and represents the runes that
// make up the symbol.
type Symbol string

// QUANTIZED INTERFACE

// This method returns a string value for this quantized element.
func (v Symbol) AsString() string {
	return string(v)
}
