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
	uti "github.com/bali-nebula/go-component-framework/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

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
}

func TestBase32RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	// Encode as base 32.
	var base32 = uti.Base32Encode(bytes)

	// Decode base 32 to bytes.
	var decoded = uti.Base32Decode(base32)
	ass.Equal(t, bytes, decoded)

	// Encode as base 32 again.
	var encoded = uti.Base32Encode(decoded)
	ass.Equal(t, base32, encoded)
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
}

func TestBase64RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	// Encode as base 64.
	var base64 = uti.Base64Encode(bytes)

	// Decode base 64 to bytes.
	var decoded = uti.Base64Decode(base64)
	ass.Equal(t, bytes, decoded)

	// Encode as base 64 again.
	var encoded = uti.Base64Encode(decoded)
	ass.Equal(t, base64, encoded)
}
