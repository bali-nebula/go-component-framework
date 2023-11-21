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
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	col "github.com/craterdog/go-collection-framework/v2"
	sts "strings"
)

// NAME STRING INTERFACE

// This constructor creates a new name string from the specified string.
func NameFromString(string_ string) abs.NameLike {
	var matches = uti.NameMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a name string from an invalid string: %v", string_)
		panic(message)
	}
	var name = name_(matches[0])
	return name
}

// This constructor attempts to create a new name string from the specified
// array of identifiers. It returns the corresponding name string.
func NameFromArray(array []abs.Identifier) abs.NameLike {
	var builder sts.Builder
	for _, identifier := range array {
		builder.WriteString("/" + string(identifier))
	}
	var name = name_(builder.String())
	return name
}

// This constructor creates a new name string from the specified name space
// sequence. It returns the corresponding name string.
func NameFromSequence(sequence abs.Sequential[abs.Identifier]) abs.NameLike {
	var array = sequence.AsArray()
	var name = NameFromArray(array)
	return name
}

// NAME STRING IMPLEMENTATION

// This type defines the methods associated with a name string that
// extends the native Go string type and represents the string of identifiers
// that make up the name.
type name_ string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
func (v name_) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v name_) IsEmpty() bool {
	return false // Names must have at least one identifier.
}

// This method returns the number of identifiers contained in this string.
func (v name_) GetSize() int {
	var identifiers = v.AsArray()
	return len(identifiers)
}

// This method returns all the identifiers in this string. The identifiers
// retrieved are in the same order as they are in the string.
func (v name_) AsArray() []abs.Identifier {
	var identifiers = sts.Split(string(v[1:]), "/")
	var array = make([]abs.Identifier, len(identifiers))
	for index, identifier := range identifiers {
		array[index] = abs.Identifier(identifier)
	}
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the identifier that is associated with
// the specified index.
func (v name_) GetValue(index int) abs.Identifier {
	var array = v.AsArray()
	var identifiers = col.Array[abs.Identifier](array)
	return identifiers.GetValue(index)
}

// This method retrieves from this string all identifiers from the first index
// through the last index (inclusive).
func (v name_) GetValues(first int, last int) abs.Sequential[abs.Identifier] {
	var array = v.AsArray()
	var identifiers = col.Array[abs.Identifier](array)
	return identifiers.GetValues(first, last)
}

// NAME LIBRARY

// This singleton creates a unique name space for the library functions for
// name strings.
var Name = &names_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for name strings.
type names_ struct{}

// This function returns the concatenation of the two specified name strings.
func (l *names_) Concatenate(first, second abs.NameLike) abs.NameLike {
	var name = first.AsString() + second.AsString()
	return name_(name)
}
