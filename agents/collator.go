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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"math/cmplx"
	"reflect"
	"strings"
)

// COLLATOR INTERFACE

// This function determines whether or not the specified values are equal using
// their natural comparison criteria.
func CompareValues(first any, second any) bool {
	var v = Collator()
	return v.CompareValues(first, second)
}

// This function ranks the specified values based on their natural ordering.
func RankValues(first any, second any) int {
	var v = Collator()
	return v.RankValues(first, second)
}

// This constructor creates a new instance of a formatter that can be used to
// generate the natural string format for any value.
func Collator() abstractions.CollatorLike {
	return &collator{abstractions.AgentState()}
}

// COLLATOR IMPLEMENTATION

// This type defines the structure and methods for a natural formatting agent.
type collator struct {
	state abstractions.AgentStateLike
}

// This method determines whether or not the specified values are equal.
func (v *collator) CompareValues(first any, second any) bool {
	return v.compareValues(reflect.ValueOf(first), reflect.ValueOf(second))
}

// This method ranks the specified values using their natural ordering.
func (v *collator) RankValues(first any, second any) int {
	return v.rankValues(reflect.ValueOf(first), reflect.ValueOf(second))
}

// This function determines whether or not the specified reflective values are
// equal using reflection and a recursive descent algorithm.
func (v *collator) compareValues(first reflect.Value, second reflect.Value) bool {
	// Handle nil values.
	if !first.IsValid() {
		return !second.IsValid()
	}
	if !second.IsValid() {
		return false
	}

	// At this point, neither of the types are nil.
	var firstType = baseTypeName(first.Type())
	var secondType = baseTypeName(second.Type())
	if firstType != secondType && firstType != "any" && secondType != "any" {
		// The values have different types.
		return false
	}

	// At this point, the types of the values are the same, and neither of
	// the values is nil.
	switch first.Kind() {

	// Handle all primitive elements.
	case reflect.Bool,
		reflect.Uint8, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.String:
		return v.compareElements(first, second)

	// Handle all primitive collections.
	case reflect.Array, reflect.Slice:
		return v.compareArrays(first, second)
	case reflect.Map:
		return v.compareMaps(first, second)

	// Handle all sequential collections.
	case reflect.Interface, reflect.Pointer:
		if first.MethodByName("AsArray").IsValid() {
			// The value is a collection.
			return v.compareCollections(first, second)
		} else if first.MethodByName("GetKey").IsValid() {
			// The value is an association.
			return v.compareAssociations(first, second)
		} else {
			// The values are pointers to the values to be compared.
			first = first.Elem()
			second = second.Elem()
			return v.compareValues(first, second)
		}

	default:
		panic(fmt.Sprintf(
			"Attempted to compare:\n\tfirst: %v\n\ttype: %v\n\tkind: %v\nand\n\tsecond: %v\n\ttype: %v\n\tkind: %v\n",
			first.Interface(),
			first.Type(),
			first.Kind(),
			second.Interface(),
			second.Type(),
			second.Kind()))
	}
}

// This private method determines whether or not the specified elements have
// the same value.
func (v *collator) compareElements(first reflect.Value, second reflect.Value) bool {
	return first.Interface() == second.Interface()
}

// This private method determines whether or not the specified arrays have
// the same value.
func (v *collator) compareArrays(first reflect.Value, second reflect.Value) bool {
	var size = first.Len()
	if second.Len() != size {
		// The arrays are of different lengths.
		return false
	}
	for i := 0; i < size; i++ {
		v.state.IncrementDepth()
		if !v.compareValues(first.Index(i), second.Index(i)) {
			// Two of the items in the arrays are different.
			v.state.DecrementDepth()
			return false
		}
		v.state.DecrementDepth()
	}
	// All items in the arrays are equal.
	return true
}

// This private method determines whether or not the specified maps have
// the same value.
func (v *collator) compareMaps(first reflect.Value, second reflect.Value) bool {
	// Compare the sizes of the two maps.
	if first.Len() != second.Len() {
		// The maps are different sizes.
		return false
	}

	// Compare the keys and values for the two maps.
	var iterator = first.MapRange()
	for iterator.Next() {
		v.state.IncrementDepth()
		var key = iterator.Key()
		var firstValue = iterator.Value()
		var secondValue = second.MapIndex(key)
		if !v.compareValues(firstValue, secondValue) {
			// The values don't match.
			v.state.DecrementDepth()
			return false
		}
		v.state.DecrementDepth()
	}

	// All keys and values match.
	return true
}

// This private method determines whether or not the specified associations have
// the same key value pair.
func (v *collator) compareAssociations(first reflect.Value, second reflect.Value) bool {
	// Compare the keys of the two associations.
	var firstKey = first.MethodByName("GetKey").Call([]reflect.Value{})[0]
	var secondKey = second.MethodByName("GetKey").Call([]reflect.Value{})[0]
	if !v.compareValues(firstKey, secondKey) {
		// The keys don't match.
		return false
	}

	// The keys match so compare the values of the two associations.
	var firstValue = first.MethodByName("GetValue").Call([]reflect.Value{})[0]
	var secondValue = second.MethodByName("GetValue").Call([]reflect.Value{})[0]
	return v.compareValues(firstValue, secondValue)
}

// This private method determines whether or not the specified collections
// have the same value.
func (v *collator) compareCollections(first reflect.Value, second reflect.Value) bool {
	// Compare the arrays for the two collections.
	var firstArray = first.MethodByName("AsArray").Call([]reflect.Value{})[0]
	var secondArray = second.MethodByName("AsArray").Call([]reflect.Value{})[0]
	return v.compareArrays(firstArray, secondArray)
}

// PRIVATE FUNCTIONS

func (v *collator) rankReflective(first any, second any) int {
	var firstValue = first.(reflect.Value)
	var secondValue = second.(reflect.Value)
	return v.rankValues(firstValue, secondValue)
}

// This private method returns the ranking order of the specified values using
// reflection and a recursive descent algorithm.
func (v *collator) rankValues(first reflect.Value, second reflect.Value) int {
	if !first.IsValid() {
		if !second.IsValid() {
			// Both values are nil.
			return 0
		}
		// Only the first value is nil.
		return -1
	} else if !second.IsValid() {
		// Only the second value is nil.
		return 1
	}

	// At this point, neither of the values are nil.
	var firstType = baseTypeName(first.Type())
	var secondType = baseTypeName(second.Type())
	if firstType != secondType && firstType != "any" && secondType != "any" {
		// The values have different types.
		return RankValues(firstType, secondType)
	}

	// At this point, the types of the values are the same,
	// and neither of the values is nil.
	switch first.Kind() {

	// Handle all primitive element types.
	case reflect.Bool:
		return rankBooleans[bool](v, first, second)

	case reflect.Uint8:
		return rankNumbers[uint8](v, first, second)

	case reflect.Uint16:
		return rankNumbers[uint16](v, first, second)

	case reflect.Uint32:
		return rankNumbers[uint32](v, first, second)

	case reflect.Uint64:
		return rankNumbers[uint64](v, first, second)

	case reflect.Uint:
		return rankNumbers[uint](v, first, second)

	case reflect.Uintptr:
		return rankNumbers[uintptr](v, first, second)

	case reflect.Int8:
		return rankNumbers[int8](v, first, second)

	case reflect.Int16:
		return rankNumbers[int16](v, first, second)

	case reflect.Int32:
		return rankNumbers[int32](v, first, second)

	case reflect.Int64:
		return rankNumbers[int64](v, first, second)

	case reflect.Int:
		return rankNumbers[int](v, first, second)

	case reflect.Float32:
		return rankNumbers[float32](v, first, second)

	case reflect.Float64:
		return rankNumbers[float64](v, first, second)

	case reflect.Complex64:
		return rankVectors[complex64](v, first, second)

	case reflect.Complex128:
		return rankVectors[complex128](v, first, second)

	case reflect.String:
		return rankStrings[string](v, first, second)

	// Handle all primitive collection types.
	case reflect.Array, reflect.Slice:
		return v.rankArrays(first, second)
	case reflect.Map:
		return v.rankMaps(first, second)

	// Handle all sequential collections.
	case reflect.Interface, reflect.Pointer:
		if first.MethodByName("AsArray").IsValid() {
			// The value is a collection.
			return v.rankCollections(first, second)
		} else if first.MethodByName("GetKey").IsValid() {
			// The value is an association.
			return v.rankAssociations(first, second)
		} else {
			// The values are pointers to the values to be ranked.
			first = first.Elem()
			second = second.Elem()
			return v.rankValues(first, second)
		}

	default:
		panic(fmt.Sprintf(
			"Attempted to rank:\n\tfirst: %v\n\ttype: %v\n\tkind: %v\nand\n\tsecond: %v\n\ttype: %v\n\tkind: %v\n",
			first.Interface(),
			first.Type(),
			first.Kind(),
			second.Interface(),
			second.Type(),
			second.Kind()))
	}
}

// This private method returns the ranking order of the specified arrays using
// a recursive descent algorithm.
func (v *collator) rankArrays(first reflect.Value, second reflect.Value) int {
	// Determine the smallest array.
	var firstSize = first.Len()
	var secondSize = second.Len()
	if firstSize > secondSize {
		// Swap the order of the arrays and reverse the sign of the result.
		return -1 * v.rankArrays(second, first)
	}

	// Iterate through the smallest array.
	for i := 0; i < firstSize; i++ {
		v.state.IncrementDepth()
		var rank = v.rankValues(first.Index(i), second.Index(i))
		if rank < 0 {
			// The item in the first array comes before its matching item.
			v.state.DecrementDepth()
			return -1
		}
		if rank > 0 {
			// The item in the first array comes after its matching item.
			v.state.DecrementDepth()
			return 1
		}
		// The two items match.
		v.state.DecrementDepth()
	}

	// The arrays contain the same initial items.
	if secondSize > firstSize {
		// The shorter array is ranked before the longer array.
		return -1
	}

	// The arrays are the same length and contain the same items.
	return 0
}

// This private method returns the ranking order of the specified maps using a
// recursive descent algorithm. Note: currently the implementation of Go maps is
// hashtable based. The order of the keys is random, even for two maps with the
// same keys if the associations were entered in different sequences. Therefore
// at this time it is necessary to sort the key arrays for each map. This
// introduces a circular dependency between the implementation of the collator
// and the sorter (i.e. rankMaps() -> SortArray() -> RankingFunction type).
func (v *collator) rankMaps(first reflect.Value, second reflect.Value) int {
	// Extract and sort the keys for the two maps.
	var firstKeys = first.MapKeys() // The returned keys are in random order.
	SortArray(firstKeys, v.rankReflective)
	var secondKeys = second.MapKeys() // The returned keys are in random order.
	SortArray(secondKeys, v.rankReflective)

	// Determine the smallest map.
	var firstSize = len(firstKeys)
	var secondSize = len(secondKeys)
	if firstSize > secondSize {
		// Swap the order of the maps and reverse the sign of the result.
		return -1 * v.rankMaps(second, first)
	}

	// Iterate through the smallest map.
	for i := 0; i < firstSize; i++ {
		// Rank the two keys.
		var firstKey = firstKeys[i]
		var secondKey = secondKeys[i]
		var keyRank = v.rankValues(firstKey, secondKey)
		if keyRank < 0 {
			// The key in the first map comes before its matching key.
			return -1
		}
		if keyRank > 0 {
			// The key in the first map comes after its matching key.
			return 1
		}

		// The two keys match so rank the corresponding values.
		var firstValue = first.MapIndex(firstKey)
		var secondValue = second.MapIndex(secondKey)
		var valueRank = v.rankValues(firstValue, secondValue)
		if valueRank < 0 {
			// The value in the first map comes before its matching value.
			return -1
		}
		if valueRank > 0 {
			// The value in the first map comes after its matching value.
			return 1
		}
	}

	// The maps contain the same initial associations.
	if secondSize > firstSize {
		// The shorter map is ranked before the longer map.
		return -1
	}

	// All keys and values match.
	return 0
}

// This private method returns the ranking order of the specified associations.
func (v *collator) rankAssociations(first reflect.Value, second reflect.Value) int {
	// Rank the keys of the two associations.
	var firstKey = first.MethodByName("GetKey").Call([]reflect.Value{})[0]
	var secondKey = second.MethodByName("GetKey").Call([]reflect.Value{})[0]
	var keyRank = v.rankValues(firstKey, secondKey)
	if keyRank < 0 {
		// The key in the first association comes before the second.
		return -1
	}
	if keyRank > 0 {
		// The key in the first association comes after the second.
		return 1
	}

	// The keys match so rank the values of the two associations.
	var firstValue = first.MethodByName("GetValue").Call([]reflect.Value{})[0]
	var secondValue = second.MethodByName("GetValue").Call([]reflect.Value{})[0]
	return v.rankValues(firstValue, secondValue)
}

// This private method returns the ranking order of the specified collections
// using a recursive descent algorithm.
func (v *collator) rankCollections(first reflect.Value, second reflect.Value) int {
	// Rank the arrays for the two collections.
	var firstArray = first.MethodByName("AsArray").Call([]reflect.Value{})[0]
	var secondArray = second.MethodByName("AsArray").Call([]reflect.Value{})[0]
	return v.rankArrays(firstArray, secondArray)
}

// This private function returns the ranking order of the specified boolean
// values.
func rankBooleans[T ~bool](v *collator, first reflect.Value, second reflect.Value) int {
	var firstBoolean = first.Interface().(T)
	var secondBoolean = second.Interface().(T)
	if !firstBoolean && secondBoolean {
		return -1
	}
	if firstBoolean && !secondBoolean {
		return 1
	}
	return 0
}

// This private function returns the ranking order of the specified numberic
// values.
func rankNumbers[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr |
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
	~float32 | ~float64](v *collator, first reflect.Value, second reflect.Value) int {
	var firstNumber = first.Interface().(T)
	var secondNumber = second.Interface().(T)
	if firstNumber < secondNumber {
		return -1
	}
	if firstNumber > secondNumber {
		return 1
	}
	return 0
}

// This private function returns the ranking order of the specified complex
// number values.
func rankVectors[T ~complex64 | ~complex128](v *collator, first reflect.Value, second reflect.Value) int {
	// The cmplx library requires type complex128.
	var firstVector = (complex128)(first.Interface().(T))
	var secondVector = (complex128)(second.Interface().(T))
	if firstVector == secondVector {
		return 0
	}
	switch {
	case cmplx.Abs(firstVector) < cmplx.Abs(secondVector):
		// The magnitude of the first vector is less than the second.
		return -1
	case cmplx.Abs(firstVector) > cmplx.Abs(secondVector):
		// The magnitude of the first vector is greater than the second.
		return 1
	default:
		// The magnitudes of the vectors are equal.
		switch {
		case cmplx.Phase(firstVector) < cmplx.Phase(secondVector):
			// The phase of the first vector is less than the second.
			return -1
		case cmplx.Phase(firstVector) > cmplx.Phase(secondVector):
			// The phase of the first vector is greater than the second.
			return 1
		default:
			// The phases of the vectors are also equal.
			return 0
		}
	}
}

// This private function returns the ranking order of the specified string
// values.
func rankStrings[T ~string](v *collator, first reflect.Value, second reflect.Value) int {
	var firstString = first.Interface().(T)
	var secondString = second.Interface().(T)
	if firstString < secondString {
		// The first string comes before the second string lexigraphically.
		return -1
	}
	if firstString > secondString {
		// The first string comes after the second string lexigraphically.
		return 1
	}
	// The two strings are the same.
	return 0
}

// This function removes the generics from the type string for the specified
// type and converts an empty interface into type "any".
func baseTypeName(t reflect.Type) string {
	var result = t.String()
	var index = strings.Index(result, "[")
	if index > -1 {
		result = result[:index]
	}
	if result == "interface {}" {
		result = "any"
	}
	return result
}
