/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStacksWithStrings(t *testing.T) {
	var stack = collections.Stack[string]()
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.GetSize())
	stack.RemoveAll()
	stack.AddItem("foo")
	stack.AddItem("bar")
	stack.AddItem("baz")
	assert.Equal(t, 3, stack.GetSize())
	assert.Equal(t, "baz", string(stack.GetTop()))
	assert.Equal(t, "baz", string(stack.RemoveTop()))
	assert.Equal(t, 2, stack.GetSize())
	assert.Equal(t, "bar", string(stack.GetTop()))
	assert.Equal(t, "bar", string(stack.RemoveTop()))
	assert.Equal(t, 1, stack.GetSize())
	assert.Equal(t, "foo", string(stack.GetTop()))
	stack.RemoveAll()
}
