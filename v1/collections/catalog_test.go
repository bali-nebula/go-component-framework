/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	abs "github.com/bali-nebula/go-component-framework/v1/abstractions"
	col "github.com/bali-nebula/go-component-framework/v1/collections"
	com "github.com/bali-nebula/go-component-framework/v1/components"
	ele "github.com/bali-nebula/go-component-framework/v1/elements"
	str "github.com/bali-nebula/go-component-framework/v1/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestCatalogs(t *tes.T) {
	var foo = str.SymbolFromString("foo")
	var bar = str.SymbolFromString("bar")
	var baz = str.SymbolFromString("baz")
	var one = com.Component(ele.NumberFromComplex(1))
	var two = com.Component(ele.NumberFromComplex(2))
	var three = com.Component(ele.NumberFromComplex(3))
	var catalog = col.Catalog()
	ass.True(t, catalog.IsEmpty())
	ass.Equal(t, 0, catalog.GetSize())
	catalog.SetValue(foo, one)
	ass.False(t, catalog.IsEmpty())
	ass.Equal(t, 1, catalog.GetSize())
	ass.Equal(t, one, catalog.GetValue(foo))
	catalog.SetValue(bar, two)
	ass.Equal(t, 2, catalog.GetSize())
	ass.Equal(t, two, catalog.GetValue(bar))
	catalog.SetValue(baz, three)
	ass.Equal(t, 3, catalog.GetSize())
	ass.Equal(t, three, catalog.GetValue(baz))
	ass.Equal(t, []abs.Primitive{foo, bar, baz}, catalog.GetKeys().AsArray())
	catalog.SortValues()
	catalog.ShuffleValues()
	catalog.RemoveAll()
}
