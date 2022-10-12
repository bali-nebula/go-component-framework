/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase32EmptyRoundTrip(t *testing.T) {
	var bytes = make([]byte, 0)

	// Encode as base 32.
	var base32 = utilities.Base32Encode(bytes)

	// Decode base 32 to bytes.
	var decoded = utilities.Base32Decode(base32)
	assert.Equal(t, bytes, decoded)

	// Encode as base 32 again.
	var encoded = utilities.Base32Encode(decoded)
	assert.Equal(t, base32, encoded)
}

func TestBase32RoundTrip(t *testing.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	// Encode as base 32.
	var base32 = utilities.Base32Encode(bytes)

	// Decode base 32 to bytes.
	var decoded = utilities.Base32Decode(base32)
	assert.Equal(t, bytes, decoded)

	// Encode as base 32 again.
	var encoded = utilities.Base32Encode(decoded)
	assert.Equal(t, base32, encoded)
}

func TestBase64EmptyRoundTrip(t *testing.T) {
	var bytes = make([]byte, 0)

	// Encode as base 64.
	var base64 = utilities.Base64Encode(bytes)

	// Decode base 64 to bytes.
	var decoded = utilities.Base64Decode(base64)
	assert.Equal(t, bytes, decoded)

	// Encode as base 64 again.
	var encoded = utilities.Base64Encode(decoded)
	assert.Equal(t, base64, encoded)
}

func TestBase64RoundTrip(t *testing.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index, _ := range bytes {
		bytes[index] = byte(index)
	}

	// Encode as base 64.
	var base64 = utilities.Base64Encode(bytes)

	// Decode base 64 to bytes.
	var decoded = utilities.Base64Decode(base64)
	assert.Equal(t, bytes, decoded)

	// Encode as base 64 again.
	var encoded = utilities.Base64Encode(decoded)
	assert.Equal(t, base64, encoded)
}
