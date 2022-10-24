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
	sts "strings"
)

// MONIKER STRING INTERFACE

// This constructor attempts to create a new moniker string from the specified
// array of names. It returns the corresponding moniker value.
func MonikerFromNames(v []string) Moniker {
	var moniker = Moniker("/" + sts.Join(v, "/"))
	return moniker
}

// This type defines the methods associated with a moniker string that
// extends the native Go string type and represents the string of names that
// make up the moniker.
type Moniker string

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Moniker) IsEmpty() bool {
	return false // Monikers must have at least one name.
}

// This method returns the number of names contained in this string.
func (v Moniker) GetSize() int {
	var names = v.AsArray()
	return len(names)
}

// This method returns all the names in this string. The names retrieved are in
// the same order as they are in the string.
func (v Moniker) AsArray() []string {
	return sts.Split(string(v[1:]), "/")
}

// INDEXED INTERFACE

// This method retrieves from this string the name that is associated with the
// specified index.
func (v Moniker) GetItem(index int) string {
	var names = v.AsArray()
	var length = len(names)
	index = abs.NormalizedIndex(index, length)
	return names[index]
}

// This method retrieves from this string all names from the first index through
// the last index (inclusive).
func (v Moniker) GetItems(first int, last int) []string {
	var names = v.AsArray()
	var length = len(names)
	first = abs.NormalizedIndex(first, length)
	last = abs.NormalizedIndex(last, length)
	var size = last - first + 1
	return names[first:size]
}

// This method returns the index of the FIRST occurence of the specified name in
// this string, or zero if this string does not contain the name.
func (v Moniker) GetIndex(name string) int {
	var names = v.AsArray()
	for index, candidate := range names {
		if candidate == name {
			// Found the name.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The name was not found.
	return 0
}

// MONIKERS LIBRARY

// This singleton creates a unique name space for the library functions for
// monikers.
var Monikers = &monikers{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for monikers.
type monikers struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified moniker
// strings.
func (l *monikers) Concatenate(first Moniker, second Moniker) Moniker {
	var moniker = first + second
	return Moniker(moniker)
}
