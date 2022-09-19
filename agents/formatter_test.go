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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Define the zero value keys.
var k1 bool
var k2 uint
var k3 int
var k4 float64
var k5 complex128
var k6 rune
var k7 string
var k8 string = "array"

// Define the values.
var v1 bool = true
var v2 uint = 0xa
var v3 int = 42
var v4 float64 = 0.125
var v5 complex128 = 1 + 1i
var v6 rune = '☺'
var v7 string = "Hello World!"
var v8 = []int{1, 2, 3}

// Define the formatter.
var formatter = agents.Formatter()

func TestFormatterWithPrimitives(t *testing.T) {
	assert.Equal(t, 0, formatter.GetIndentation())

	s0 := formatter.FormatValue(nil)
	assert.Equal(t, "<nil>", s0)
	fmt.Println("\nNil: " + s0)

	s1 := formatter.FormatValue(v1)
	assert.Equal(t, "true", s1)
	fmt.Println("\nBoolean: " + s1)

	s2 := formatter.FormatValue(v2)
	assert.Equal(t, "0xa", s2)
	fmt.Println("\nUnsigned: " + s2)

	s3 := formatter.FormatValue(v3)
	assert.Equal(t, "42", s3)
	fmt.Println("\nInteger: " + s3)

	s4 := formatter.FormatValue(v4)
	assert.Equal(t, "0.125", s4)
	fmt.Println("\nFloat: " + s4)

	s5 := formatter.FormatValue(v5)
	assert.Equal(t, "(1+1i)", s5)
	fmt.Println("\nComplex: " + s5)

	s6 := formatter.FormatValue(v6)
	assert.Equal(t, "'☺'", s6)
	fmt.Println("\nRune: " + s6)

	s7 := formatter.FormatValue(v7)
	assert.Equal(t, "\"Hello World!\"", s7)
	fmt.Println("\nString: " + s7)
}

func TestFormatterWithEmptyArray(t *testing.T) {
	array := []any{}
	s := formatter.FormatValue(array)
	assert.Equal(t, "[ ](array)", s)
	fmt.Println("\nEmpty Array: " + s)
}

func TestFormatterWithArrayOfAny(t *testing.T) {
	array := []any{v1, v2, v3, v4, v5, v6, v7, v8}
	s := formatter.FormatValue(array)
	assert.Equal(t, "(array)", s[len(s)-7:])
	fmt.Println("\nArray of Any: " + s)
}

func TestFormatterWithArrayOfIntegers(t *testing.T) {
	array := []int{1, 2, 3, 4}
	s := formatter.FormatValue(array)
	assert.Equal(t, "(array)", s[len(s)-7:])
	fmt.Println("\nArray of Integers: " + s)
}

func TestFormatterWithEmptyMap(t *testing.T) {
	mapp := map[any]any{}
	s := formatter.FormatValue(mapp)
	assert.Equal(t, "[:](map)", s)
	fmt.Println("\nEmpty Map: " + s)
}

func TestFormatterWithMapOfAnyToAny(t *testing.T) {
	mapp := map[any]any{
		k1: v1,
		k2: v2,
		k3: v3,
		k4: v4,
		k5: v5,
		k6: v6,
		k7: v7,
		k8: v8}
	s := formatter.FormatValue(mapp)
	assert.Equal(t, "(map)", s[len(s)-5:])
	fmt.Println("\nMap of Any to Any: " + s)
}

func TestFormatterWithMapOfStringToInteger(t *testing.T) {
	mapp := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3}
	s := formatter.FormatValue(mapp)
	assert.Equal(t, "(map)", s[len(s)-5:])
	fmt.Println("\nMap of String to Integer: " + s)
}

func TestFormatterWithEmptyList(t *testing.T) {
	list := collections.List[any]()
	s := formatter.FormatValue(list)
	assert.Equal(t, "[ ](list)", s)
	fmt.Println("\nEmpty List: " + s)
}

func TestFormatterWithListOfAny(t *testing.T) {
	list := collections.List[any]()
	list.AddItems([]any{v1, v2, v3, v4, v5, v6, v7, v8})
	s := formatter.FormatValue(list)
	assert.Equal(t, "(list)", s[len(s)-6:])
	fmt.Println("\nList of Any: " + s)
}

func TestFormatterWithListOfBoolean(t *testing.T) {
	list := collections.List[bool]()
	list.AddItems([]bool{k1, v1})
	s := formatter.FormatValue(list)
	fmt.Println("\nList of Boolean: " + s)
}

func TestFormatterWithSetOfAny(t *testing.T) {
	set := collections.Set[any]()
	set.AddItems([]any{v1, v2, v3, v4, v5, v6, v7, v8})
	s := formatter.FormatValue(set)
	assert.Equal(t, "(set)", s[len(s)-5:])
	fmt.Println("\nSet of Any: " + s)
}

func TestFormatterWithSetOfSet(t *testing.T) {
	set1 := collections.Set[any]()
	set1.AddItems([]any{v1, v2, v3, v4, v5, v6, v7, v8})
	set2 := collections.Set[any]()
	set2.AddItems([]any{k1, k2, k3, k4, k5, k6, k7, k8})
	set := collections.Set[abstractions.SetLike[any]]()
	set.AddItems([]abstractions.SetLike[any]{set1, set2})
	s := formatter.FormatValue(set)
	fmt.Println("\nSet of Set: " + s)
}

func TestFormatterWithStackOfAny(t *testing.T) {
	stack := collections.Stack[any]()
	stack.AddItem(v1)
	stack.AddItem(v2)
	stack.AddItem(v3)
	stack.AddItem(v4)
	stack.AddItem(v5)
	stack.AddItem(v6)
	stack.AddItem(v7)
	s := formatter.FormatValue(stack)
	assert.Equal(t, "(stack)", s[len(s)-7:])
	fmt.Println("\nStack: " + s)
}

func TestFormatterWithQueueOfAny(t *testing.T) {
	var queue abstractions.FIFO[any] = collections.Queue[any]()
	queue.AddItem(v1)
	queue.AddItem(v2)
	queue.AddItem(v3)
	queue.AddItem(v4)
	queue.AddItem(v5)
	queue.AddItem(v6)
	queue.AddItem(v7)
	s := formatter.FormatValue(queue)
	assert.Equal(t, "(queue)", s[len(s)-7:])
	fmt.Println("\nQueue: " + s)
}

func TestFormatterWithAssociationOfAnyToAny(t *testing.T) {
	association := collections.Association[any, any]("foo", 5)
	s := formatter.FormatValue(association)
	fmt.Println("\nAssociation of Any to Any: " + s)
}

func TestFormatterWithAssociationOfStringToInteger(t *testing.T) {
	association := collections.Association[string, int]("bar", 42)
	s := formatter.FormatValue(association)
	fmt.Println("\nAssociation of String to Integer: " + s)
}

func TestFormatterWithEmptyCatalog(t *testing.T) {
	catalog := collections.Catalog[any, any]()
	s := formatter.FormatValue(catalog)
	assert.Equal(t, "[:](catalog)", s)
	fmt.Println("\nEmpty Catalog: " + s)
}

func TestFormatterWithCatalogOfAnyToAny(t *testing.T) {
	catalog := collections.Catalog[any, any]()
	catalog.SetValue(k1, v1)
	catalog.SetValue(k2, v2)
	catalog.SetValue(k3, v3)
	catalog.SetValue(k4, v4)
	catalog.SetValue(k5, v5)
	catalog.SetValue(k6, v6)
	catalog.SetValue(k7, v7)
	s := formatter.FormatValue(catalog)
	assert.Equal(t, "(catalog)", s[len(s)-9:])
	fmt.Println("\nCatalog: " + s)
}

func TestFormatterWithCatalogOfStringToAny(t *testing.T) {
	catalog := collections.Catalog[string, any]()
	catalog.SetValue("key1", v1)
	catalog.SetValue("key2", v2)
	catalog.SetValue("key3", v3)
	catalog.SetValue("key4", v4)
	catalog.SetValue("key5", v5)
	catalog.SetValue("key6", v6)
	catalog.SetValue("key7", v7)
	s := formatter.FormatValue(catalog)
	fmt.Println("\nCatalog of String to Any: " + s)
}

func TestFormatterWithCatalogOfStringToInteger(t *testing.T) {
	catalog := collections.Catalog[string, int]()
	catalog.SetValue("key1", 1)
	catalog.SetValue("key2", 2)
	catalog.SetValue("key3", 3)
	catalog.SetValue("key4", 4)
	catalog.SetValue("key5", 5)
	catalog.SetValue("key6", 6)
	catalog.SetValue("key7", 7)
	s := formatter.FormatValue(catalog)
	fmt.Println("\nCatalog of String to Integer: " + s)
}
