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
	uti "github.com/craterdog-bali/go-bali-document-notation/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRandomIntegers(t *tes.T) {
	var foundZero bool
	var foundFive bool
	for i := 0; i < 100; i++ {
		var random = uti.RandomInteger(5)
		if random == 0 {
			foundZero = true
		}
		if random == 5 {
			foundFive = true
		}
	}
	ass.True(t, foundZero)
	ass.True(t, foundFive)
}

func TestRandomProbabilities(t *tes.T) {
	var total float64
	for i := 0; i < 10000; i++ {
		total += uti.RandomProbability()
	}
	ass.True(t, total > 4800)
	ass.True(t, total < 5200)
}
