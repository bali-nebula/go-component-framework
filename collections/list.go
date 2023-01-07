/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ent "github.com/bali-nebula/go-component-framework/entities"
	col "github.com/craterdog/go-collection-framework"
)

// LIST IMPLEMENTATION

// This constructor creates a new empty list that uses the canonical compare
// function.
func List() abs.ListLike {
	return col.List[abs.ComponentLike]().(abs.ListLike)
}

// This constructor creates a new list from the specified series of values.
func ListFromSeries(series abs.SeriesLike) abs.ListLike {
	var list = List()
	var iterator = ent.Iterator[abs.ComponentLike](series)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		list.AddValue(value)
	}
	return list
}
