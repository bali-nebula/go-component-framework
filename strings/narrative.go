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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
	sts "strings"
)

// NARRATIVE STRING IMPLEMENTATION

// This constructor creates a new narrative string from the specified string.
func NarrativeFromString(string_ string) abs.NarrativeLike {
	// Since a narrative may be recursive we cannot (easily) use a regular
	// expression to scan it. We must trust it is correct here.
	return Narrative(string_)
}

// This constructor attempts to create a new narrative string from the specified
// array of runes. It returns the corresponding narrative string.
func NarrativeFromArray(array []abs.Line) abs.NarrativeLike {
	var builder sts.Builder
	var lines = col.Array[abs.Line](array)
	var iterator = col.Iterator[abs.Line](lines)
	if iterator.HasNext() {
		var line = string(iterator.GetNext())
		builder.WriteString(line)
	}
	for iterator.HasNext() {
		var line = string(iterator.GetNext())
		builder.WriteString("\n")
		builder.WriteString(line)
	}
	var narrative = Narrative(builder.String())
	return narrative
}

// This constructor creates a new narrative string from the specified sequence of
// runes. It returns the corresponding narrative string.
func NarrativeFromSequence(sequence abs.Sequential[abs.Line]) abs.NarrativeLike {
	var array = sequence.AsArray()
	var narrative = NarrativeFromArray(array)
	return narrative
}

// This type defines the methods associated with a narrative string that extends
// the native Go string type and represents the string of runes that make up
// the narrative string.
// For valid string formats for this type see `../abstractions/language.go`.
type Narrative string

// QUANTIZED INTERFACE

// This method returns a string value for this quantized string.
func (v Narrative) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Narrative) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of runes contained in this narrative.
func (v Narrative) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the runes in this string. The runes retrieved
// are in the same order as they are in the string.
func (v Narrative) AsArray() []abs.Line {
	var lines = sts.Split(string(v), "\n")
	var length = len(lines)
	var array = make([]abs.Line, length)
	for index, line := range lines {
		array[index] = abs.Line(line)
	}
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this narrative the line that is associated
// with the specified index.
func (v Narrative) GetValue(index int) abs.Line {
	var array = v.AsArray()
	var lines = col.Array[abs.Line](array)
	return lines.GetValue(index)
}

// This method retrieves from this narrative all lines from the first index
// through the last index (inclusive).
func (v Narrative) GetValues(first int, last int) abs.Sequential[abs.Line] {
	var array = v.AsArray()
	var lines = col.Array[abs.Line](array)
	return lines.GetValues(first, last)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all narrative string
// library functions.
const Narratives narratives = false

// This type defines the library functions associated with narrative strings.
type narratives bool

// This function returns the concatenation of the two specified narrative
// strings.
func (l narratives) Concatenate(first, second abs.NarrativeLike) abs.NarrativeLike {
	var a1 = first.AsArray()
	var a2 = second.AsArray()
	var array = make([]abs.Line, len(a1)+len(a2))
	copy(array, a1)
	copy(array[len(a1):], a2)
	var narrative = NarrativeFromArray(array)
	return narrative
}
