/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities_test

import (
	uti "github.com/bali-nebula/go-component-framework/v1/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBase16EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 16.
	var base16 = uti.Base16Encode(bytes)

	// Decode base 16 to bytes.
	var decoded = uti.Base16Decode(base16)
	ass.Equal(t, bytes, decoded)

	// Encode as base 16 again.
	var encoded = uti.Base16Encode(decoded)
	ass.Equal(t, base16, encoded)

	// Decode base 16 again.
	var again = uti.Base16Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase16RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 16.
		var base16 = uti.Base16Encode(bytes[:index])

		// Decode base 16 to bytes.
		var decoded = uti.Base16Decode(base16)
		ass.Equal(t, bytes[:index], decoded)

		// Encode as base 16 again.
		var encoded = uti.Base16Encode(decoded)
		ass.Equal(t, base16, encoded)

		// Decode base 16 again.
		var again = uti.Base16Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}

func TestBase32EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 32.
	var base32 = uti.Base32Encode(bytes)

	// Decode base 32 to bytes.
	var decoded = uti.Base32Decode(base32)
	ass.Equal(t, bytes, decoded)

	// Encode as base 32 again.
	var encoded = uti.Base32Encode(decoded)
	ass.Equal(t, base32, encoded)

	// Decode base 32 again.
	var again = uti.Base32Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase32RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 32.
		var base32 = uti.Base32Encode(bytes[:index])

		// Decode base 32 to bytes.
		var decoded = uti.Base32Decode(base32)
		ass.Equal(t, bytes[:index], decoded)

		// Encode as base 32 again.
		var encoded = uti.Base32Encode(decoded)
		ass.Equal(t, base32, encoded)

		// Decode base 32 again.
		var again = uti.Base32Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}

func TestBase64EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 64.
	var base64 = uti.Base64Encode(bytes)

	// Decode base 64 to bytes.
	var decoded = uti.Base64Decode(base64)
	ass.Equal(t, bytes, decoded)

	// Encode as base 64 again.
	var encoded = uti.Base64Encode(decoded)
	ass.Equal(t, base64, encoded)

	// Decode base 64 again.
	var again = uti.Base64Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase64RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 64.
		var base64 = uti.Base64Encode(bytes)

		// Decode base 64 to bytes.
		var decoded = uti.Base64Decode(base64)
		ass.Equal(t, bytes, decoded)

		// Encode as base 64 again.
		var encoded = uti.Base64Encode(decoded)
		ass.Equal(t, base64, encoded)

		// Decode base 64 again.
		var again = uti.Base64Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}
