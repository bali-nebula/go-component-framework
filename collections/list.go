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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	col "github.com/craterdog/go-collection-framework"
)

// LIST IMPLEMENTATION

// This constructor creates a new empty list that uses the canonical compare
// function.
func List() abs.ListLike {
	return col.List[abs.ComponentLike]()
}

// This constructor creates a new empty list that uses the specified compare
// function.
func ListWithComparer(compare col.ComparisonFunction) abs.ListLike {
	return col.ListWithComparer[abs.ComponentLike](compare)
}

// This constructor creates a new list from the specified array that uses the
// canonical compare function.
func ListFromArray(array []abs.ComponentLike) abs.ListLike {
	return col.ListFromArray[abs.ComponentLike](array)
}

// This constructor creates a new list from the specified sequence. The list
// uses the natural compare function.
func ListFromSequence(sequence col.Sequential[abs.ComponentLike]) abs.ListLike {
	return col.ListFromSequence[abs.ComponentLike](sequence)
}
