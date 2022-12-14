/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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
	uti "github.com/bali-nebula/go-component-framework/utilities"
	col "github.com/craterdog/go-collection-framework"
)

// BINARY STRING INTERFACE

// This constructor creates a new binary string from the specified byte array.
// It returns the corresponding binary string.
func BinaryFromArray(array []byte) Binary {
	var encoded = uti.Base64Encode(array)
	return Binary(encoded)
}

// This constructor creates a new binary string from the specified byte
// sequence. It returns the corresponding binary string.
func BinaryFromSequence(sequence abs.Sequential[byte]) Binary {
	var array = sequence.AsArray()
	return BinaryFromArray(array)
}

// This type defines the methods associated with a binary string that extends
// the native Go string type and represents the string of bytes that make up
// the binary string.
type Binary string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
func (v Binary) AsString() string {
	return string(v)
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v Binary) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of bytes contained in this string.
func (v Binary) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the bytes in this string. The bytes retrieved
// are in the same order as they are in the string.
func (v Binary) AsArray() []byte {
	var encoded = string(v)
	var bytes = uti.Base64Decode(encoded)
	return bytes
}

// INDEXED INTERFACE

// This method retrieves from this string the byte that is associated
// with the specified index.
func (v Binary) GetValue(index int) byte {
	var array = v.AsArray()
	var bytes = col.Array[byte](array)
	return bytes.GetValue(index)
}

// This method retrieves from this string all bytes from the first index
// through the last index (inclusive).
func (v Binary) GetValues(first int, last int) abs.Sequential[byte] {
	var array = v.AsArray()
	var bytes = col.Array[byte](array)
	return bytes.GetValues(first, last)
}

// This method returns the index of the FIRST occurence of the specified byte
// in this string, or zero if this string does not contain the byte.
func (v Binary) GetIndex(b byte) int {
	var array = v.AsArray()
	var bytes = col.Array[byte](array)
	return bytes.GetIndex(b)
}

// LIBRARY FUNCTIONS

// This constant defines a namespace within this package for all binary string
// library functions.
const Binaries binaries = false

// This type defines the library functions associated with binary strings.
type binaries bool

// This function returns the concatenation of the two specified binary
// strings.
func (l binaries) Concatenate(first Binary, second Binary) Binary {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make([]byte, len(firstBytes)+len(secondBytes))
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return BinaryFromArray(allBytes)
}

// This function returns the logical (bitwise) inverse of the specified
// binary string.
func (l binaries) Not(binary Binary) Binary {
	var bytes = binary.AsArray()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return BinaryFromArray(bytes)
}

// This function returns the logical (bitwise) conjunction of the
// specified binary strings.
func (l binaries) And(first Binary, second Binary) Binary {
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
func (l binaries) Sans(first Binary, second Binary) Binary {
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
func (l binaries) Or(first Binary, second Binary) Binary {
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
func (l binaries) Xor(first Binary, second Binary) Binary {
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
