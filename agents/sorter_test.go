/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSortingEmpty(t *tes.T) {
	var empty = []any{}
	age.SortArray(empty, age.RankValues)
}

func TestSortingIntegers(t *tes.T) {
	var unsorted = []int{4, 3, 1, 5, 2}
	var sorted = []int{1, 2, 3, 4, 5}
	age.SortArray(unsorted, age.RankValues)
	ass.Equal(t, sorted, unsorted)
}

func TestSortingStrings(t *tes.T) {
	var unsorted = []string{"alpha", "beta", "gamma", "delta"}
	var sorted = []string{"alpha", "beta", "delta", "gamma"}
	age.SortArray(unsorted, age.RankValues)
	ass.Equal(t, sorted, unsorted)
}
