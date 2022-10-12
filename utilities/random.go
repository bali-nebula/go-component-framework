/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const maximum = 1 << 53

// This function generates a cryptographically secure random integer in the
// range [0..max].
func RandomInteger(max int) int {
	var random, err = rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	if err != nil {
		panic(err)
	}
	return int(random.Int64())
}

// This function generates a cryptographically secure random probability in the
// range [0..1].
func RandomProbability() float64 {
	return float64(RandomInteger(maximum)) / float64(maximum)
}

// This function generates a cryptographically secure array of random bytes.
func RandomBytes(size int) []byte {
	var bytes = make([]byte, size)
	var _, err = rand.Read(bytes)
	if err != nil {
		panic(fmt.Sprintf("The random number generator gave the following error: %v", err))
	}
	return bytes
}
