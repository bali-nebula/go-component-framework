/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities

import (
	b64 "encoding/base64"
	b16 "encoding/hex"
	fmt "fmt"
	mat "math"
	sts "strings"
	uni "unicode"
)

// This function encodes the specified bytes as a base 16 encoded string using
// standard base 16 encoding with no padding.
func Base16Encode(bytes []byte) string {
	var encoded = b16.EncodeToString(bytes)
	return encoded
}

// This function decodes the specified base 16 encoded string into its
// corresponding decoded bytes.
func Base16Decode(encoded string) []byte {
	var bytes, err = b16.DecodeString(encoded)
	if err != nil {
		panic(fmt.Sprintf("The binary data was not encoded using base 16: %s", encoded))
	}
	return bytes
}

// This function encodes the specified bytes as a base 32 string using the
// ten digits and all UPPER CASE letters except 'E', 'I', 'O', and 'U'.
func Base32Encode(bytes []byte) string {
	// Encode each byte.
	var result sts.Builder
	var previousByte byte
	for index, currentByte := range bytes {
		// Encode the next one or two 5 bit chunks.
		base32EncodeBytes(previousByte, currentByte, index, &result)
		previousByte = currentByte
	}

	// Encode the last 5 bit chunk.
	var index = len(bytes) - 1
	base32EncodeLast(previousByte, index, &result)

	return result.String()
}

// This function decodes the specified base 32 string into its corresponding
// bytes.
func Base32Decode(encoded string) []byte {
	// Purify the base 32 encoded string.
	var buffer sts.Builder
	for _, c := range encoded {
		if !uni.IsSpace(c) {
			var r = uni.ToUpper(c)
			buffer.WriteRune(r)
		}
	}
	encoded = buffer.String()

	// Decode each base 32 character.
	var size = len(encoded)
	var bytes = make([]byte, int(mat.Trunc(float64(size)*5.0/8.0)))
	for index, r := range encoded {
		var chunk = byte(sts.Index(base32LookupTable, string(r)))
		if index < size-1 {
			base32DecodeBytes(chunk, index, bytes)
		} else {
			base32DecodeLast(chunk, index, bytes)
		}
	}
	return bytes
}

// This function encodes the specified bytes as a base 64 encoded string using
// standard base 64 encoding with no padding.
func Base64Encode(bytes []byte) string {
	var encoded = b64.RawStdEncoding.EncodeToString(bytes)
	return encoded
}

// This function decodes the specified base 64 encoded string into its
// corresponding decoded bytes.
func Base64Decode(encoded string) []byte {
	var bytes, err = b64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		panic(fmt.Sprintf("The binary data was not encoded using base 64: %s", encoded))
	}
	return bytes
}

// PRIVATE FUNCTIONS

// This lookup table maps the base 32 characters to the corresponding base 32
// digits. The letters 'E', 'I', 'O', and 'U' have been removed to avoid the
// possibility of randomly occurring  offensive words.
const base32LookupTable = "0123456789ABCDFGHJKLMNPQRSTVWXYZ"

/*
 * offset:    0        1        2        3        4        0
 * byte:  00000111|11222223|33334444|45555566|66677777|...
 * mask:   F8  07  C0 3E  01 F0  0F 80  7C 03  E0  1F   F8  07
 */
func base32EncodeBytes(previous byte, current byte, index int, base32 *sts.Builder) {
	var chunk byte
	switch index % 5 {
	case 0:
		chunk = (current & 0xF8) >> 3
		base32.WriteByte(base32LookupTable[chunk])
	case 1:
		chunk = ((previous & 0x07) << 2) | ((current & 0xC0) >> 6)
		base32.WriteByte(base32LookupTable[chunk])
		chunk = (current & 0x3E) >> 1
		base32.WriteByte(base32LookupTable[chunk])
	case 2:
		chunk = ((previous & 0x01) << 4) | ((current & 0xF0) >> 4)
		base32.WriteByte(base32LookupTable[chunk])
	case 3:
		chunk = ((previous & 0x0F) << 1) | ((current & 0x80) >> 7)
		base32.WriteByte(base32LookupTable[chunk])
		chunk = (current & 0x7C) >> 2
		base32.WriteByte(base32LookupTable[chunk])
	case 4:
		chunk = ((previous & 0x03) << 3) | ((current & 0xE0) >> 5)
		base32.WriteByte(base32LookupTable[chunk])
		chunk = current & 0x1F
		base32.WriteByte(base32LookupTable[chunk])
	}
}

/*
 * Same as normal, but pad with 0's in "next" byte
 * case:      0        1        2        3        4
 * byte:  xxxxx111|00xxxxx3|00004444|0xxxxx66|000xxxxx|...
 * mask:   F8  07  C0 3E  01 F0  0F 80  7C 03  E0  1F
 */
func base32EncodeLast(last byte, index int, base32 *sts.Builder) {
	var chunk byte
	switch index % 5 {
	case 0:
		chunk = (last & 0x07) << 2
		base32.WriteByte(base32LookupTable[chunk])
	case 1:
		chunk = (last & 0x01) << 4
		base32.WriteByte(base32LookupTable[chunk])
	case 2:
		chunk = (last & 0x0F) << 1
		base32.WriteByte(base32LookupTable[chunk])
	case 3:
		chunk = (last & 0x03) << 3
		base32.WriteByte(base32LookupTable[chunk])
	case 4:
		// nothing to do, was handled by previous call
	}
}

/*
 * offset:    0        1        2        3        4        0
 * byte:  00000111|11222223|33334444|45555566|66677777|...
 * mask:   F8  07  C0 3E  01 F0  0F 80  7C 03  E0  1F   F8  07
 */
func base32DecodeBytes(chunk byte, characterIndex int, bytes []byte) {
	var byteIndex = int(mat.Trunc(float64(characterIndex) * 5.0 / 8.0))
	switch characterIndex % 8 {
	case 0:
		bytes[byteIndex] |= chunk << 3
	case 1:
		bytes[byteIndex] |= chunk >> 2
		bytes[byteIndex+1] |= chunk << 6
	case 2:
		bytes[byteIndex] |= chunk << 1
	case 3:
		bytes[byteIndex] |= chunk >> 4
		bytes[byteIndex+1] |= chunk << 4
	case 4:
		bytes[byteIndex] |= chunk >> 1
		bytes[byteIndex+1] |= chunk << 7
	case 5:
		bytes[byteIndex] |= chunk << 2
	case 6:
		bytes[byteIndex] |= chunk >> 3
		bytes[byteIndex+1] |= chunk << 5
	case 7:
		bytes[byteIndex] |= chunk
	}
}

/*
 * Same as normal, but pad with 0's in "next" byte
 * case:      0        1        2        3        4
 * byte:  xxxxx111|00xxxxx3|00004444|0xxxxx66|00077777|...
 * mask:   F8  07  C0 3E  01 F0  0F 80  7C 03  E0  1F
 */
func base32DecodeLast(chunk byte, characterIndex int, bytes []byte) {
	var byteIndex = int(mat.Trunc(float64(characterIndex) * 5.0 / 8.0))
	switch characterIndex % 8 {
	case 1:
		bytes[byteIndex] |= chunk >> 2
	case 3:
		bytes[byteIndex] |= chunk >> 4
	case 4:
		bytes[byteIndex] |= chunk >> 1
	case 6:
		bytes[byteIndex] |= chunk >> 3
	case 7:
		bytes[byteIndex] |= chunk
	}
}
