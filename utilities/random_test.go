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

func TestRandomIntegers(t *testing.T) {
	var foundZero bool
	var foundFive bool
	for i := 0; i < 100; i++ {
		var random = utilities.RandomInteger(5)
		if random == 0 {
			foundZero = true
		}
		if random == 5 {
			foundFive = true
		}
	}
	assert.True(t, foundZero)
	assert.True(t, foundFive)
}

func TestRandomProbabilities(t *testing.T) {
	var total float64
	for i := 0; i < 10000; i++ {
		total += utilities.RandomProbability()
	}
	assert.True(t, total > 4800)
	assert.True(t, total < 5200)
}
