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
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	stc "strconv"
	str "strings"
)

// VERSION STRING INTERFACE

// This constructor attempts to create a new version string from the specified
// formatted string. It returns a version value and whether or not the string
// contained a valid version.
// For valid string formats for this type see `../abstractions/language.go`.
func VersionFromString(v string) (Version, bool) {
	var ok = true
	var version string
	var matches = abs.ScanVersion([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		version = matches[0]
	}
	return Version(version), ok
}

// This constructor attempts to create a new version string from the specified
// array of ordinals. It returns the corresponding version string.
func VersionFromOrdinals(v []int) Version {
	var length = len(v)
	var version = "v"
	for i, ordinal := range v {
		if ordinal < 1 {
			panic(fmt.Sprintf("All version numbers must be greater than zero: %v", v))
		}
		version += stc.FormatInt(int64(ordinal), 10)
		if i < length-1 {
			version += "."
		}
	}
	return Version(version)
}

// This type defines the methods associated with a version string that extends
// the native Go string type and represents the string of ordinals that make up
// the version string.
type Version string

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
func (v Version) AsArray() []int {
	var ordinals []int
	var levels = str.Split(string(v[1:]), ".")
	for _, level := range levels {
		var ordinal, _ = stc.ParseInt(level, 10, 0)
		ordinals = append(ordinals, int(ordinal))
	}
	return ordinals
}

// INDEXED INTERFACE

// This method retrieves from this string the ordinal number that is associated
// with the specified index.
func (v Version) GetItem(index int) int {
	var ordinals = v.AsArray()
	var length = len(ordinals)
	index = abs.NormalizedIndex(index, length)
	return ordinals[index]
}

// This method retrieves from this string all ordinals from the first index
// through the last index (inclusive).
func (v Version) GetItems(first int, last int) []int {
	var ordinals = v.AsArray()
	var length = len(ordinals)
	first = abs.NormalizedIndex(first, length)
	last = abs.NormalizedIndex(last, length)
	var size = last - first + 1
	return ordinals[first:size]
}

// This method returns the index of the FIRST occurence of the specified ordinal
// number in this string, or zero if this string does not contain the ordinal
// number.
func (v Version) GetIndex(ordinal int) int {
	var ordinals = v.AsArray()
	for index, candidate := range ordinals {
		if candidate == ordinal {
			// Found the ordinal number.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The ordinal number was not found.
	return 0
}

// VERSIONS LIBRARY

// This singleton creates a unique name space for the library functions for
// versions.
var Versions = &versions{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for versions.
type versions struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified version
// strings.
func (l *versions) Concatenate(first Version, second Version) Version {
	var version = first + "." + second[1:]
	return Version(version)
}

// This library function returns a copy of the specified version string with the
// ordinal at the specified level incremented by one. For example:
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
func (l *versions) GetNextVersion(current Version, level int) Version {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = len(ordinals)
	switch {
	case level < 1:
		panic("The level of the next version must be positive!")
	case level < size:
		// The next version will require fewer levels.
		size = level
		ordinals = ordinals[:level]
	case level > size:
		// The next version will require another level.
		size++
		level = size // Normalize the level.
		ordinals = append(ordinals, 0)
	default:
		// The level is equal to the current size.
	}

	// Increment the specified version level.
	var index = level - 1 // Convert to zero based indexing.
	ordinals[index]++

	var version = VersionFromOrdinals(ordinals)
	return version
}

// This library function determines whether or not the specified next version
// string is a valid next version string for the specified current version
// string. In order for the next version to be valid the last level in the next
// version string must be one more than the corresponding level in the current
// version string; or it must be '1' and the next version string must have one
// more level of version numbers than the current version string, for example:
//
//	           Current             Next          What Likely Changed
//	level 1:    v5.7              v6         (interface/symantic changes)
//	level 2:    v5.7              v5.8       (optimization/bug fixes)
//	level 3:    v5.7              v5.7.1     (changes being tested)
func (l *versions) IsValidNextVersion(current Version, next Version) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var currentIterator = age.Iterator[int](current)
	var nextIterator = age.Iterator[int](next)
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
