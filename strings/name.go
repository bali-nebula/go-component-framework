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
	"strings"
)

// NAME SPACE INTERFACE

// This constructor attempts to create a new name space from the specified
// formatted string. It returns a name value and whether or not the string
// contained a valid name.
// For valid string formats for this type see `../abstractions/language.go`.
func NameFromString(v string) (Name, bool) {
	var ok = true
	var name string
	var matches = abstractions.ScanName([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		name = matches[0]
	}
	return Name(name), ok
}

// This constructor attempts to create a new name space from the specified
// array of labels. It returns the corresponding name value.
func NameFromLabels(v []string) Name {
	var name, ok = NameFromString("/" + strings.Join(v, "/"))
	if !ok {
		panic(fmt.Sprintf("The labels contain an illegal character: %v", v))
	}
	return name
}

// This type defines the methods associated with a name space string that
// extends the native Go string type and represents the string of labels that
// make up the name space.
type Name string

// LEXICAL INTERFACE

// This method returns the canonical string for this string.
func (v Name) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Name) String() string {
	return v.AsString()
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Name) IsEmpty() bool {
	return false // Names must have at least one label.
}

// This method returns the number of labels contained in this string.
func (v Name) GetSize() int {
	var labels = v.AsArray()
	return len(labels)
}

// This method returns all the labels in this string. The labels retrieved are in
// the same order as they are in the string.
func (v Name) AsArray() []string {
	return strings.Split(string(v[1:]), "/")
}

// INDEXED INTERFACE

// This method retrieves from this string the label that is associated with the
// specified index.
func (v Name) GetItem(index int) string {
	var labels = v.AsArray()
	var length = len(labels)
	index = abstractions.NormalizedIndex(index, length)
	return labels[index]
}

// This method retrieves from this string all labels from the first index through
// the last index (inclusive).
func (v Name) GetItems(first int, last int) []string {
	var labels = v.AsArray()
	var length = len(labels)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	var size = last - first + 1
	return labels[first:size]
}

// This method returns the index of the FIRST occurence of the specified label in
// this string, or zero if this string does not contain the label.
func (v Name) GetIndex(label string) int {
	var labels = v.AsArray()
	for index, candidate := range labels {
		if candidate == label {
			// Found the label.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The label was not found.
	return 0
}

// NAMES LIBRARY

// This singleton creates a unique name space for the library functions for
// names.
var Names = &names{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for names.
type names struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified name
// strings.
func (l *names) Concatenate(first Name, second Name) Name {
	var name = first + second
	return Name(name)
}
