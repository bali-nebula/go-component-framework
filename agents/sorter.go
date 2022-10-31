/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// SORTER INTERFACE

// This function sorts the specified array of items using the specified ranking
// function.
func SortArray[T abs.ItemLike](array []T, rank abs.RankingFunction) {
	var v = Sorter[T](rank)
	v.SortArray(array)
}

// This constructor creates a new instance of a sorter that can be used to
// sort an array using a specific ranking function.
func Sorter[T abs.ItemLike](rank abs.RankingFunction) abs.SorterLike[T] {
	return &sorter[T]{rank: rank}
}

// SORTER IMPLEMENTATION

// This type defines the structure and methods for a merge sorter agent.
type sorter[T abs.ItemLike] struct {
	rank abs.RankingFunction
}

// This method determines whether or not the specified values are equal.
func (v *sorter[T]) SortArray(array []T) {
	v.sortArray(array)
}

// The following methods are used to sort an array using the ranking function
// that is associated with the sorter. These methods implement a merge sort
// that has been optimized to be iterative rather than recursive. They also
// save on memory allocation by swapping between two arrays of the same size
// rather than allocating new arrays for each subarray.  This results in stable
// O[nlog(n)] time and O[n] space performance. The algorithm is documented here:
//   - https://en.wikipedia.org/wiki/Merge_sort#Bottom-up_implementation

func (v *sorter[T]) sortArray(array []T) {
	// Create a buffer array.
	var length = len(array)
	var buffer = make([]T, length)
	copy(buffer, array) // Make a copy of the original unsorted array.

	// Iterate through subarray widths of 2, 4, 8, ... length.
	for width := 1; width < length; width *= 2 {

		// Split the buffer array into two arrays.
		for left := 0; left < length; left += width * 2 {

			// Find the middle (it must be less than length).
			var middle = left + width
			if middle > length {
				middle = length
			}

			// Find the right side (it must be less than length).
			var right = middle + width
			if right > length {
				right = length
			}

			// Sort and merge the subarrays.
			v.mergeArrays(
				buffer[left:middle],
				buffer[middle:right],
				array[left:right])
		}

		// Swap the two arrays.
		buffer, array = array, buffer
	}

	// Synchronize the two arrays.
	copy(array, buffer) // Both arrays are now sorted.
}

func (v *sorter[T]) mergeArrays(left []T, right []T, merged []T) {
	var leftIndex = 0
	var leftLength = len(left)
	var rightIndex = 0
	var rightLength = len(right)
	var mergedIndex = 0
	var mergedLength = len(merged)

	// Work our way through filling the entire merged array.
	for mergedIndex < mergedLength {

		// Check to see if both left and right arrays still have items.
		if leftIndex < leftLength && rightIndex < rightLength {

			// Copy the next smallest item to the merged array.
			if v.rank(left[leftIndex], right[rightIndex]) < 0 {
				merged[mergedIndex] = left[leftIndex]
				leftIndex++
			} else {
				merged[mergedIndex] = right[rightIndex]
				rightIndex++
			}

		} else if leftIndex < leftLength {
			// Copy the rest of the left array to the merged array.
			copy(merged[mergedIndex:], left[leftIndex:])
			leftIndex++

		} else {
			// Copy the rest of the right array to the merged array.
			copy(merged[mergedIndex:], right[rightIndex:])
			rightIndex++
		}
		mergedIndex++
	}
}
