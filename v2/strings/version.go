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
	stc "strconv"
	sts "strings"
)

// CONSTANT DEFINITIONS

var V1 abs.VersionLike = Version("v1")

// VERSION STRING IMPLEMENTATION

// These constants are used to define a regular expression for matching
// version strings.
const (
	ordinal = `[1-9][0-9]*`
	version = `(` + ordinal + `(?:\.` + ordinal + `)*)`
)

// This scanner is used for matching version strings.
var versionScanner = reg.MustCompile(`^(?:` + version + `)$`)

// This constructor creates a new version string from the specified string.
func VersionFromString(string_ string) abs.VersionLike {
	if !versionScanner.MatchString(string_) {
		var message = fmt.Sprintf("Attempted to construct a version string from an invalid string: %v", string_)
		panic(message)
	}
	return Version(string_)
}

// This constructor attempts to create a new version string from the specified
// array of ordinal numbers. It returns the corresponding version string.
func VersionFromArray(array []abs.Ordinal) abs.VersionLike {
	var length = len(array)
	var version string
	for i, ordinal := range array {
		if ordinal < 1 {
			panic(fmt.Sprintf("All version numbers must be greater than zero: %v", array))
		}
		version += stc.FormatUint(uint64(ordinal), 10)
		if i < length-1 {
			version += "."
		}
	}
	return Version(version)
}

// This constructor creates a new version string from the specified sequence of
// ordinal numbers. It returns the corresponding version string.
func VersionFromSequence(sequence abs.Sequential[abs.Ordinal]) abs.VersionLike {
	var array = sequence.AsArray()
	return VersionFromArray(array)
}

// This type defines the methods associated with a version string that extends
// the native Go string type and represents the string of ordinals that make up
// the version string.
type Version string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
func (v Version) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Version) IsEmpty() bool {
	return false // A version must have at least one level.
}

// This method returns the number of ordinals contained in this string.
func (v Version) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the ordinals in this string. The ordinals retrieved
// are in the same order as they are in the string.
func (v Version) AsArray() []abs.Ordinal {
	var ordinals []abs.Ordinal
	var levels = sts.Split(string(v), ".")
	for _, level := range levels {
		var ordinal, _ = stc.ParseUint(level, 10, 0)
		ordinals = append(ordinals, abs.Ordinal(ordinal))
	}
	return ordinals
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the ordinal number that is associated
// with the specified index.
func (v Version) GetValue(index int) abs.Ordinal {
	var array = v.AsArray()
	var ordinals = col.Array[abs.Ordinal](array)
	return ordinals.GetValue(index)
}

// This method retrieves from this string all ordinals from the first index
// through the last index (inclusive).
func (v Version) GetValues(first int, last int) abs.Sequential[abs.Ordinal] {
	var array = v.AsArray()
	var ordinals = col.Array[abs.Ordinal](array)
	return ordinals.GetValues(first, last)
}

// VERSION LIBRARY

// This singleton creates a unique name space for the library functions for
// version strings.
var Version = &versions_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for version strings.
type versions_ struct{}

// This function returns the concatenation of the two specified version strings.
func (l *versions) Concatenate(first, second abs.VersionLike) abs.VersionLike {
	var version = first.AsString() + "." + second.AsString()
	return Version(version)
}

// This function returns a copy of the specified version string with the ordinal
// at the specified level incremented by one. For example:
//
//	           Current             Next          What Likely Changed
//	level 1:    v5.7              v6         (interface/symantic changes)
//	level 2:    v5.7              v5.8       (optimization/bug fixes)
//	level 3:    v5.7              v5.7.1     (changes being tested)
//
// A level of `-1` will result in the last ordinal in the copy of the current
// version string being incremented. A level that is greater than the size of
// current version will result in a new level with the value of `1` being
// appended to the copy of the current version string.
func (l *versions) GetNextVersion(current abs.VersionLike, level abs.Ordinal) abs.VersionLike {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = abs.Ordinal(len(ordinals))
	switch {
	case level == 0:
		level = size // Normalize the level to the current size.
	case level < size:
		// The next version will require fewer levels.
		ordinals = ordinals[:level]
	case level > size:
		// The next version will require another level.
		size++
		level = size // Normalize the level to the new size.
		ordinals = append(ordinals, 0)
	}

	// Increment the specified version level.
	var index = level - 1 // Convert to zero based indexing.
	ordinals[index]++

	var version = VersionFromArray(ordinals)
	return version
}

// This function determines whether or not the specified next version string
// is a valid next version string for the specified current version string. In
// order for the next version to be valid the last level in the next version
// string must be one more than the corresponding level in the current version
// string; or it must be '1' and the next version string must have one more
// level of version numbers than the current version string, for example:
//
//	           Current             Next          What Likely Changed
//	level 1:    v5.7              v6         (interface/symantic changes)
//	level 2:    v5.7              v5.8       (optimization/bug fixes)
//	level 3:    v5.7              v5.7.1     (changes being tested)
func (l *versions) IsValidNextVersion(current, next abs.VersionLike) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var currentIterator = col.Iterator[abs.Ordinal](current)
	var nextIterator = col.Iterator[abs.Ordinal](next)
	for currentIterator.HasNext() && nextIterator.HasNext() {
		var currentLevel = currentIterator.GetNext()
		var nextLevel = nextIterator.GetNext()
		if currentLevel == nextLevel {
			// So far the level values are the same.
			continue
		}
		// The last level for the next version must be one more.
		return !nextIterator.HasNext() && nextLevel == currentLevel+1
	}
	// The last level for the next version must be one.
	return nextIterator.HasNext() && nextIterator.GetNext() == 1
}
