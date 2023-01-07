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

// SET IMPLEMENTATION

// This constructor creates a new empty set that uses the canonical ranking
// function.
func Set() abs.SetLike {
	return col.Set[abs.ComponentLike]().(abs.SetLike)
}

// This constructor creates a new set from the specified array that uses the
// natural ranking function.
func SetFromSeries(series abs.SeriesLike) abs.SetLike {
	var set = Set()
	var iterator = ent.Iterator[abs.ComponentLike](series)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		set.AddValue(value)
	}
	return set
}
