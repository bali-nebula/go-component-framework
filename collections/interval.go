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
	tem "github.com/bali-nebula/go-component-framework/temporary"
)

// INTERVAL IMPLEMENTATION

// This constructor creates a new discrete range of values covering the
// specified endpoints with the specified extent.
func Interval(first abs.Discrete, extent abs.Extent, last abs.Discrete) abs.IntervalLike {
	return tem.Interval[abs.Discrete](first, extent, last)
}

