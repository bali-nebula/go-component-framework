/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
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

// This type defines the methods associated with a symbol element that
// extends the native Go string type and represents the runes that
// make up the symbol.
type Symbol string

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Symbol) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Symbol) String() string {
	return v.AsString()
}
