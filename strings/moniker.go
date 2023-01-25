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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
	reg "regexp"
	sts "strings"
)

// MONIKER STRING IMPLEMENTATION

// These constants are used to define a regular expression for matching
// moniker strings.
const (
	separator = `[-+.]`
	letter    = `\pL` // All unicode letters and connectors like underscores.
	digit     = `\pN` // All unicode digits.
	name      = letter + `(?:` + separator + `?(?:` + letter + `|` + digit + `))*`
	moniker   = `(?:/` + name + `)+`
)

// This scanner is used for matching moniker strings.
var monikerScanner = reg.MustCompile(`^(?:` + moniker + `)$`)

// This constructor creates a new moniker string from the specified string.
func MonikerFromString(v string) abs.MonikerLike {
	if !monikerScanner.MatchString(v) {
		var message = fmt.Sprintf("Attempted to construct a moniker string from an invalid string: %v", v)
		panic(message)
	}
	return Moniker(v)
}

// This constructor attempts to create a new moniker string from the specified
// array of names. It returns the corresponding moniker string.
func MonikerFromArray(array []abs.Name) abs.MonikerLike {
	var builder sts.Builder
	for _, name := range array {
		builder.WriteString("/" + string(name))
	}
	var moniker = Moniker(builder.String())
	return moniker
}

// This constructor creates a new moniker string from the specified name space
// sequence. It returns the corresponding moniker string.
func MonikerFromSequence(sequence abs.Sequential[abs.Name]) abs.MonikerLike {
	var array = sequence.AsArray()
	return MonikerFromArray(array)
}

// This type defines the methods associated with a moniker string that
// extends the native Go string type and represents the string of names that
// make up the moniker.
type Moniker string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
func (v Moniker) AsString() string {
	return string(v)
}

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
func (v Moniker) AsArray() []abs.Name {
	var names = sts.Split(string(v[1:]), "/")
	var array = make([]abs.Name, len(names))
	for index, name := range names {
		array[index] = abs.Name(name)
	}
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the name that is associated with the
// specified index.
func (v Moniker) GetValue(index int) abs.Name {
	var array = v.AsArray()
	var names = col.Array[abs.Name](array)
	return names.GetValue(index)
}

// This method retrieves from this string all names from the first index through
// the last index (inclusive).
func (v Moniker) GetValues(first int, last int) abs.Sequential[abs.Name] {
	var array = v.AsArray()
	var names = col.Array[abs.Name](array)
	return names.GetValues(first, last)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all moniker string
// library functions.
const Monikers monikers = false

// This type defines the library functions associated with moniker strings.
type monikers bool

// This function returns the concatenation of the two specified moniker
// strings.
func (l monikers) Concatenate(first, second abs.MonikerLike) abs.MonikerLike {
	var moniker = first.AsString() + second.AsString()
	return Moniker(moniker)
}
