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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	uti "github.com/bali-nebula/go-component-framework/utilities"
	reg "regexp"
)

// TAG IMPLEMENTATION

// These constants are used to define a regular expression for matching
// tags.
const (
	base32 = `[0-9A-DF-HJ-NP-TV-Z]` // "E", "I", "O", and "U" have been removed.
	tag    = `(` + base32 + `+)`
)

// This scanner is used for matching tag strings.
var tagScanner = reg.MustCompile(`^(?:` + tag + `)$`)

// This constructor creates a new tag element from the specified string.
func TagFromString(string_ string) abs.TagLike {
	if !tagScanner.MatchString(string_) {
		var message = fmt.Sprintf("Attempted to construct a tag from an invalid string: %v", string_)
		panic(message)
	}
	return Tag(string_)
}

// This constructor creates a new binary string from the specified byte array.
// It returns the corresponding binary string.
func TagFromBytes(bytes []byte) abs.TagLike {
	var encoded = uti.Base32Encode(bytes)
	return Tag(encoded)
}

// This constructor creates a new tag element from the specified string.
func TagFromArray(array []abs.Byte) abs.TagLike {
	var bytes = make([]byte, len(array))
	for index, b := range array {
		bytes[index] = byte(b)
	}
	var base32 = uti.Base32Encode(bytes)
	return Tag(base32)
}

// This constructor creates a new random tag element with the specified number
// of bytes.
func TagOfSize(size int) abs.TagLike {
	if size < 1 {
		size = 20 // Default to 20 bytes.
	}
	var bytes = uti.RandomBytes(size)
	var tag = uti.Base32Encode(bytes)
	return Tag(tag)
}

// This type defines the methods associated with a tag identifier that extends
// the native Go string type.
type Tag string

// QUANTIZED INTERFACE

// This method returns a string value for this lexical element.
func (v Tag) AsString() string {
	return string(v)
}

// FUNDAMENTAL INTERFACE

// This method returns the byte array for this fundamental string.
func (v Tag) AsBytes() []byte {
	var encoded = string(v)
	var bytes = uti.Base32Decode(encoded)
	return bytes
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this tag is empty.
func (v Tag) IsEmpty() bool {
	return false
}

// This method returns the number of bytes contained in this tag.
func (v Tag) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the bytes in this tag. The bytes retrieved
// are in the same order as they are in the tag.
func (v Tag) AsArray() []abs.Byte {
	var bytes = v.AsBytes()
	var array = make([]abs.Byte, len(bytes))
	for index, b := range bytes {
		array[index] = abs.Byte(b)
	}
	return array
}
