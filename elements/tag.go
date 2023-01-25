/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

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
func TagFromString(v string) abs.TagLike {
	if !tagScanner.MatchString(v) {
		var message = fmt.Sprintf("Attempted to construct a tag from an invalid string: %v", v)
		panic(message)
	}
	return Tag(v)
}

// This constructor creates a new tag element from the specified string.
func TagFromArray(v []byte) abs.TagLike {
	var base32 = uti.Base32Encode(v)
	return Tag(base32)
}

// This constructor creates a new random tag element with the specified number
// of bytes.
func TagOfSize(size int) abs.TagLike {
	if size < 1 {
		panic("A tag must be at least one byte long!")
	}
	var bytes = uti.RandomBytes(size)
	var tag = uti.Base32Encode(bytes)
	return Tag(tag)
}

// This type defines the methods associated with a tag identifier that extends
// the native Go string type.
type Tag string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
func (v Tag) AsString() string {
	return string(v)
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
func (v Tag) AsArray() []byte {
	var base32 = string(v)
	var bytes = uti.Base32Decode(base32)
	return bytes
}
