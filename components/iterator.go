/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework"
)

// ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an iterator that can be used to
// traverse the values in the specified sequence.
func Iterator[V abs.Value](sequence abs.Sequential[V]) abs.IteratorLike[V] {
	var raw = sequence.(col.Sequential[V])
	return col.Iterator[V](raw)
}
