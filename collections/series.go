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
	col "github.com/craterdog/go-collection-framework"
)

// SERIES IMPLEMENTATION

// This constructor creates a new empty series of component values.
func Series() abs.SeriesLike {
	return col.Array[abs.ComponentLike]([]abs.ComponentLike{})
}

// This constructor creates a new series of component values from the specified
// array.
func SeriesFromArray(array []abs.ComponentLike) abs.SeriesLike {
	return col.Array[abs.ComponentLike](array)
}
