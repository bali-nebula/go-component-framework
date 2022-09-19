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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/utilities"
	"strings"
	"unicode"
)

// BINARY STRING INTERFACE

// This constructor attempts to create a new binary string from the specified
// base 64 encoded string. It returns a binary value and whether or not the
// string contained valid base64 encoding.
// For valid string formats for this type see `../abstractions/language.go`.
func BinaryFromString(v string) (Binary, bool) {
	var ok = true
	var binary string
	var matches = abstractions.ScanBinary([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		binary = removeWhiteSpace(matches[0])
	}
	return Binary(binary), ok
}

// This constructor attempts to create a new binary string from the specified
// byte array. It returns the corresponding binary string.
func BinaryFromBytes(v []byte) Binary {
	var encoded = "'" + utilities.Base64Encode(v) + "'"
	return Binary(encoded)
}

// This type defines the methods associated with a binary string that extends
// the native Go string type and represents the string of bytes that make up
// the binary string.
type Binary string

// LEXICAL INTERFACE

// This method returns the canonical string for this string.
func (v Binary) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Binary) String() string {
	return v.AsString()
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Binary) IsEmpty() bool {
	return len(v) == 2 // The empty binary string is: "''".
}

// This method returns the number of bytes contained in this string.
func (v Binary) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the bytes in this string. The bytes retrieved
// are in the same order as they are in the string.
func (v Binary) AsArray() []byte {
	var encoded = string(v[1 : len(v)-1]) // Remove the single quotes.
	var bytes = utilities.Base64Decode(encoded)
	return bytes
}

// INDEXED INTERFACE

// This method retrieves from this string the byte that is associated
// with the specified index.
func (v Binary) GetItem(index int) byte {
	var bytes = v.AsArray()
	var length = len(bytes)
	index = abstractions.NormalizedIndex(index, length)
	return bytes[index]
}

// This method retrieves from this string all bytes from the first index
// through the last index (inclusive).
func (v Binary) GetItems(first int, last int) []byte {
	var bytes = v.AsArray()
	var length = len(bytes)
	first = abstractions.NormalizedIndex(first, length)
	last = abstractions.NormalizedIndex(last, length)
	return bytes[first : last+1]
}

// This method returns the index of the FIRST occurence of the specified byte
// in this string, or zero if this string does not contain the byte.
func (v Binary) GetIndex(b byte) int {
	var bytes = v.AsArray()
	for index, candidate := range bytes {
		if candidate == b {
			// Found the byte.
			return index + 1 // Convert to an ORDINAL based index.
		}
	}
	// The byte was not found.
	return 0
}

func removeWhiteSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// BINARIES LIBRARY

// This singleton creates a unique name space for the library functions for
// binaries.
var Binaries = &binaries{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for binaries.
type binaries struct{}

// CHAINABLE INTERFACE

// This library function returns the concatenation of the two specified binary
// strings.
func (l *binaries) Concatenate(first Binary, second Binary) Binary {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make([]byte, len(firstBytes)+len(secondBytes))
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return BinaryFromBytes(allBytes)
}

// LOGICAL INTERFACE

// This library function returns the logical (bitwise) inverse of the specified
// binary string.
func (l *binaries) Not(binary Binary) Binary {
	var bytes = binary.AsArray()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return BinaryFromBytes(bytes)
}

// This library function returns the logical (bitwise) conjunction of the
// specified binary strings.
func (l *binaries) And(first Binary, second Binary) Binary {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] & secondBytes[i]
	}
	return BinaryFromBytes(result)
}

// This library function returns the logical (bitwise) material non-implication
// of the specified binary strings.
func (l *binaries) Sans(first Binary, second Binary) Binary {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] &^ secondBytes[i]
	}
	return BinaryFromBytes(result)
}

// This library function returns the logical (bitwise) disjunction of the
// specified binary strings.
func (l *binaries) Or(first Binary, second Binary) Binary {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] | secondBytes[i]
	}
	return BinaryFromBytes(result)
}

// This library function returns the logical (bitwise) exclusive disjunction of
// the specified binary strings.
func (l *binaries) Xor(first Binary, second Binary) Binary {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] ^ secondBytes[i]
	}
	return BinaryFromBytes(result)
}
