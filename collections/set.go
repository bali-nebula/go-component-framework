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

// SET IMPLEMENTATION

// This constructor creates a new empty set that uses the canonical ranking
// function.
func Set() abs.SetLike {
	return col.Set[abs.ComponentLike]()
}

// This constructor creates a new empty set that uses the specified ranking
// function.
func SetWithRanker(rank abs.RankingFunction) abs.SetLike {
	return col.SetWithRanker[abs.ComponentLike](rank)
}

// This constructor creates a new set from the specified array that uses the
// natural ranking function.
func SetFromArray(array []abs.ComponentLike) abs.SetLike {
	return col.SetFromArray[abs.ComponentLike](array)
}

// This constructor creates a new set from the specified sequence. The set
// uses the natural ranking function.
func SetFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.SetLike {
	var raw = sequence.(col.Sequential[abs.ComponentLike])
	return col.SetFromSequence[abs.ComponentLike](raw)
}
