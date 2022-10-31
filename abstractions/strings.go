/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	fmt "fmt"
)

// STRING INTERFACES

type StringLike any

// This interface defines the methods supported by all sequential strings of
// items.
type Sequential[T any] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []T
}

// This interface defines the methods supported by all strings of items whose
// items can be indexed. The indices of an indexed string are ORDINAL rather
// than ZERO based (which is "SO last century"). This allows for positive
// indices starting at the beginning of the string, and negative indices
// starting at the end of the string as follows:
//
//	    1          2          3            N
//	[item 1] . [item 2] . [item 3] ... [item N]
//	   -N       -(N-1)     -(N-2)         -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Indexed[T any] interface {
	SetComparer(compare ComparisonFunction)
	GetItem(index int) T
	GetItems(first int, last int) Sequential[T]
	GetIndex(item T) int
}

// This function normalizes an index to match the Go (zero based) indexing. The
// following transformation is performed:
//
//	[-length..-1] and [1..length] => [0..length)
//
// Notice that the specified index cannot be zero since zero is not an ORDINAL.
func NormalizedIndex(index int, length int) int {
	switch {
	case index < -length || index == 0 || index > length:
		// The index is outside the bounds of the specified range.
		panic(fmt.Sprintf(
			"The specified index is outside the allowed ranges [-%v..-1] and [1..%v]: %v",
			length,
			length,
			index))
	case index < 0:
		// Convert a negative index.
		return index + length
	case index > 0:
		// Convert a positive index.
		return index - 1
	default:
		// This should never happen so time to panic...
		panic(fmt.Sprintf("Compiler problem, unexpected index value: %v", index))
	}
}

// LIBRARY INTERFACES

// This library interface defines the functions supported by all libraries of
// chainable items.
type Chainable[T any] interface {
	Concatenate(first Sequential[T], second Sequential[T]) Sequential[T]
}
