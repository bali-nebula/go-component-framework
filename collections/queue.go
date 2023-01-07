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

// QUEUE IMPLEMENTATION

// This constructor creates a new empty queue with the default capacity.
// The default capacity is 16 values.
func Queue() abs.QueueLike {
	return col.Queue[abs.ComponentLike]().(abs.QueueLike)
}

// This constructor creates a new empty queue with the specified capacity.
func QueueWithCapacity(capacity int) abs.QueueLike {
	return col.QueueWithCapacity[abs.ComponentLike](capacity).(abs.QueueLike)
}
