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

// NARRATIVE STRING INTERFACE

// This constructor creates a new narrative string from the specified string.
func NarrativeFromString(string_ string) abs.NarrativeLike {
	var matches = uti.NarrativeMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a narrative string from an invalid string: %v", string_)
		panic(message)
	}
	var narrative = matches[1] // Strip off the ">\n and <" delimiters.
	// TODO: Remove indentation from all lines.
	return narrative_(narrative)
}

// This constructor attempts to create a new narrative string from the specified
// array of runes. It returns the corresponding narrative string.
func NarrativeFromArray(array []abs.Line) abs.NarrativeLike {
	var builder sts.Builder
	var indentation = "    "
	var lines = col.Array[abs.Line](array)
	var iterator = col.Iterator[abs.Line](lines)
	for iterator.HasNext() {
		var line = string(iterator.GetNext())
		builder.WriteString(indentation)
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	var narrative = narrative_(builder.String())
	return narrative
}

// This constructor creates a new narrative string from the specified sequence of
// runes. It returns the corresponding narrative string.
func NarrativeFromSequence(sequence abs.Sequential[abs.Line]) abs.NarrativeLike {
	var array = sequence.AsArray()
	var narrative = NarrativeFromArray(array)
	return narrative
}

// NARRATIVE STRING IMPLEMENTATION

// This type defines the methods associated with a narrative string that extends
// the native Go string type and represents the string of runes that make up
// the narrative string.
// For valid string formats for this type see `../abstractions/language.go`.
type narrative_ string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
func (v narrative_) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v narrative_) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of runes contained in this narrative.
func (v narrative_) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the runes in this string. The runes retrieved
// are in the same order as they are in the string.
func (v narrative_) AsArray() []abs.Line {
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
func (v narrative_) GetValue(index int) abs.Line {
	var array = v.AsArray()
	var lines = col.Array[abs.Line](array)
	return lines.GetValue(index)
}

// This method retrieves from this narrative all lines from the first index
// through the last index (inclusive).
func (v narrative_) GetValues(first int, last int) abs.Sequential[abs.Line] {
	var array = v.AsArray()
	var lines = col.Array[abs.Line](array)
	return lines.GetValues(first, last)
}

// NARRATIVE LIBRARY

// This singleton creates a unique name space for the library functions for
// narrative strings.
var Narrative = &narratives_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for narrative strings.
type narratives_ struct{}

// This function returns the concatenation of the two specified narrative
// strings.
func (l *narratives_) Concatenate(first, second abs.NarrativeLike) abs.NarrativeLike {
	var a1 = first.AsArray()
	var a2 = second.AsArray()
	var array = make([]abs.Line, len(a1)+len(a2))
	copy(array, a1)
	copy(array[len(a1):], a2)
	var narrative = NarrativeFromArray(array)
	return narrative
}
