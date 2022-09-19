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
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortingEmpty(t *testing.T) {
	empty := []any{}
	agents.SortArray(empty, agents.RankValues)
}

func TestSortingIntegers(t *testing.T) {
	unsorted := []int{4, 3, 1, 5, 2}
	sorted := []int{1, 2, 3, 4, 5}
	agents.SortArray(unsorted, agents.RankValues)
	assert.Equal(t, sorted, unsorted)
}

func TestSortingStrings(t *testing.T) {
	unsorted := []string{"alpha", "beta", "gamma", "delta"}
	sorted := []string{"alpha", "beta", "delta", "gamma"}
	agents.SortArray(unsorted, agents.RankValues)
	assert.Equal(t, sorted, unsorted)
}
