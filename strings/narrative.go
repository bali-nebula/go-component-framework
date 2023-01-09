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
	sts "strings"
)

// NARRATIVE STRING INTERFACE

// This constructor attempts to create a new narrative string from the specified
// array of runes. It returns the corresponding narrative string.
func NarrativeFromArray(array []string) Narrative {
	var s = sts.Join(array, "\n")
	var narrative = Narrative(s)
	return narrative
}

// This constructor creates a new narrative string from the specified sequence of
// runes. It returns the corresponding narrative string.
func NarrativeFromSequence(sequence abs.Sequential[string]) Narrative {
	var array = sequence.AsArray()
	var narrative = NarrativeFromArray(array)
	return narrative
}

// This type defines the methods associated with a narrative string that extends
// the native Go string type and represents the string of runes that make up
// the narrative string.
// For valid string formats for this type see `../abstractions/language.go`.
type Narrative string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
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
func (v Narrative) AsArray() []string {
	var array = sts.Split(string(v), "\n")
	return array
}

// INDEXED INTERFACE

// This method retrieves from this narrative the line that is associated
// with the specified index.
func (v Narrative) GetValue(index int) string {
	var array = v.AsArray()
	var lines = cox.Array[string](array)
	return lines.GetValue(index)
}

// This method retrieves from this narrative all lines from the first index
// through the last index (inclusive).
func (v Narrative) GetValues(first int, last int) abs.Sequential[string] {
	var array = v.AsArray()
	var lines = cox.Array[string](array)
	return lines.GetValues(first, last)
}

// This method returns the index of the FIRST occurence of the specified rune
// in this string, or zero if this string does not contain the rune.
func (v Narrative) GetIndex(line string) int {
	var array = v.AsArray()
	var lines = cox.Array[string](array)
	return lines.GetIndex(line)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all narrative string
// library functions.
const Narratives narratives = false

// This type defines the library functions associated with narrative strings.
type narratives bool

// This function returns the concatenation of the two specified narrative
// strings.
func (l narratives) Concatenate(first Narrative, second Narrative) Narrative {
	var a1 = first.AsArray()
	var a2 = second.AsArray()
	var array = make([]string, len(a1)+len(a2))
	copy(array, a1)
	copy(array[len(a1):], a2)
	var narrative = NarrativeFromArray(array)
	return narrative
}
