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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework"
	stc "strconv"
	sts "strings"
)

// VERSION STRING INTERFACE

// This constructor attempts to create a new version string from the specified
// array of ordinal numbers. It returns the corresponding version string.
func VersionFromArray(array []int) Version {
	var length = len(array)
	var version string
	for i, ordinal := range array {
		if ordinal < 1 {
			panic(fmt.Sprintf("All version numbers must be greater than zero: %v", array))
		}
		version += stc.FormatInt(int64(ordinal), 10)
		if i < length-1 {
			version += "."
		}
	}
	return Version(version)
}

// This constructor creates a new version string from the specified sequence of
// ordinal numbers. It returns the corresponding version string.
func VersionFromSequence(sequence abs.Ordinals) Version {
	var array = sequence.AsArray()
	return VersionFromArray(array)
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
	var levels = sts.Split(string(v), ".")
	for _, level := range levels {
		var ordinal, _ = stc.ParseInt(level, 10, 0)
		ordinals = append(ordinals, int(ordinal))
	}
	return ordinals
}

// INDEXED INTERFACE

// This method retrieves from this string the ordinal number that is associated
// with the specified index.
func (v Version) GetValue(index int) int {
	var array = v.AsArray()
	var ordinals = col.Array[int](array)
	return ordinals.GetValue(index)
}

// This method retrieves from this string all ordinals from the first index
// through the last index (inclusive).
func (v Version) GetValues(first int, last int) abs.Ordinals {
	var array = v.AsArray()
	var ordinals = col.Array[int](array)
	return ordinals.GetValues(first, last)
}

// This method returns the index of the FIRST occurence of the specified ordinal
// number in this string, or zero if this string does not contain the ordinal
// number.
func (v Version) GetIndex(ordinal int) int {
	var array = v.AsArray()
	var ordinals = col.Array[int](array)
	return ordinals.GetIndex(ordinal)
}

// This method normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func (v Version) GoIndex(index int) int {
	var array = v.AsArray()
	var ordinals = col.Array[int](array)
	return ordinals.GoIndex(index)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all version string
// library functions.
const Versions versions = false

// This type defines the library functions associated with version strings.
type versions bool

// This function returns the concatenation of the two specified version strings.
func (l versions) Concatenate(first Version, second Version) Version {
	var version = first + "." + second
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
func (l versions) GetNextVersion(current Version, level int) Version {
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
func (l versions) IsValidNextVersion(current Version, next Version) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var currentIterator = col.Iterator[int](current)
	var nextIterator = col.Iterator[int](next)
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
