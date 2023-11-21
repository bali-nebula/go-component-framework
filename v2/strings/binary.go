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

// BINARY STRING INTERFACE

// This constructor creates a new binary string from the specified string.
func BinaryFromString(string_ string) abs.BinaryLike {
	var matches = uti.BinaryMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a binary string from an invalid string: %v", string_)
		panic(message)
	}
	var binary = matches[1]                   // Strip off the "'>\n" and "<'" delimiters.
	binary = sts.ReplaceAll(binary, " ", "")  // Remove all spaces.
	binary = sts.ReplaceAll(binary, "\n", "") // Remove all EOLs.
	return binary_(binary)
}

// This constructor creates a new binary string from the specified byte array.
// It returns the corresponding binary string.
func BinaryFromArray(array []byte) abs.BinaryLike {
	var encoded = uti.Base64Encode(array)
	return binary_(encoded)
}

// This constructor creates a new binary string from the specified byte
// sequence. It returns the corresponding binary string.
func BinaryFromSequence(sequence abs.Sequential[byte]) abs.BinaryLike {
	var array = sequence.AsArray()
	return BinaryFromArray(array)
}

// BINARY STRING IMPLEMENTATION

// This type defines the methods associated with a binary string that extends
// the native Go string type and represents the string of bytes that make up
// the binary string.
type binary_ string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
func (v binary_) AsString() string {
	var indentation = "    "
	var encoded = string(v)
	var width = 60
	var length = len(encoded)
	var string_ = "'>\n"
	if length > 0 {
		var index int
		for index = 0; index+width < length; index += width {
			string_ += indentation + encoded[index:index+width]
		}
		string_ += indentation + encoded[index:] + "\n"
	}
	string_ += "<'"
	return string_
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v binary_) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of bytes contained in this string.
func (v binary_) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the bytes in this string. The bytes retrieved
// are in the same order as they are in the string.
func (v binary_) AsArray() []byte {
	var encoded = string(v)
	var array = uti.Base64Decode(encoded)
	return array
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the byte that is associated
// with the specified index.
func (v binary_) GetValue(index int) byte {
	var array = v.AsArray()
	var bytes = col.Array[byte](array)
	return bytes.GetValue(index)
}

// This method retrieves from this string all bytes from the first index
// through the last index (inclusive).
func (v binary_) GetValues(first int, last int) abs.Sequential[byte] {
	var array = v.AsArray()
	var bytes = col.Array[byte](array)
	return bytes.GetValues(first, last)
}

// BINARY LIBRARY

// This singleton creates a unique name space for the library functions for
// binary strings.
var Binary = &binaries_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for binary strings.
type binaries_ struct{}

// This function returns the concatenation of the two specified binary
// strings.
func (l *binaries_) Concatenate(first, second abs.BinaryLike) abs.BinaryLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make([]byte, len(firstBytes)+len(secondBytes))
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return BinaryFromArray(allBytes)
}

// This function returns the logical (bitwise) inverse of the specified
// binary string.
func (l *binaries_) Not(binary abs.BinaryLike) abs.BinaryLike {
	var bytes = binary.AsArray()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return BinaryFromArray(bytes)
}

// This function returns the logical (bitwise) conjunction of the
// specified binary strings.
func (l *binaries_) And(first, second abs.BinaryLike) abs.BinaryLike {
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
	return BinaryFromArray(result)
}

// This function returns the logical (bitwise) material non-implication
// of the specified binary strings.
func (l *binaries_) Sans(first, second abs.BinaryLike) abs.BinaryLike {
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
	return BinaryFromArray(result)
}

// This function returns the logical (bitwise) disjunction of the
// specified binary strings.
func (l *binaries_) Or(first, second abs.BinaryLike) abs.BinaryLike {
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
	return BinaryFromArray(result)
}

// This function returns the logical (bitwise) exclusive disjunction of
// the specified binary strings.
func (l *binaries_) Xor(first, second abs.BinaryLike) abs.BinaryLike {
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
	return BinaryFromArray(result)
}
