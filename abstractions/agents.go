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
	sts "strings"
)

// ITERATION INTERFACES

// This interface defines the methods supported by all iterator-like agents that
// are capable of moving forward and backward over the values in a sequence. It
// is used to implement the GoF Iterator Pattern:
//   - https://en.wikipedia.org/wiki/Iterator_pattern
//
// An iterator locks into the slots that reside between each value in the
// sequence:
//
//	    [value 1] . [value 2] . [value 3] ... [value N]
//	  ^           ^           ^                         ^
//	slot 0      slot 1      slot 2                    slot N
//
// It moves from slot to slot and has access to the values (if they exist) on
// each side of the slot.
type IteratorLike[T ValueLike] interface {
	GetSlot() int
	ToSlot(slot int)
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() T
	HasNext() bool
	GetNext() T
}

// COLLATION INTERFACES

// This interface defines the methods supported by all collator-like agent
// types that can compare and rank two values.
type CollatorLike interface {
	CompareValues(first ValueLike, second ValueLike) bool
	RankValues(first ValueLike, second ValueLike) int
}

// This type defines the function signature for any function that can determine
// whether or not two specified values are equal to each other.
type ComparisonFunction func(first ValueLike, second ValueLike) bool

// This type defines the function signature for any function that can determine
// the relative ordering of two specified values. The result must be one of
// the following:
//   - -1: The first value is less than the second value.
//   - 0: The first value is equal to the second value.
//   - 1: The first value is more than the second value.
type RankingFunction func(first ValueLike, second ValueLike) int

// SORTING INTERFACES

// This interface defines the methods supported by all sorter-like agents that
// can sort an array of values using a ranking function.
type SorterLike[T ValueLike] interface {
	SortArray(array []T)
}

// This type defines the function signature for any function that can sort an
// array of values using a ranking function.
type Sort[T ValueLike] func(array []T, rank RankingFunction)

// AGENT STATE

// This interface defines the methods that all formatter states must implement.
type AgentStateLike interface {
	IncrementDepth() int
	CurrentDepth() int
	DecrementDepth() int
}

// This constructor creates a new agent state.
func AgentState() AgentStateLike {
	return &agentState{depth: 0}
}

// The maximum recursion level handles overly complex data models and cyclical
// references in a clean and efficient way.
const maximumRecursion int = 100

// This type defines the structure and methods for the state of any agent that
// performs a recursive descent on a compound value.
type agentState struct {
	depth int
}

// This method increases the depth level of recursion by one and returns the new
// depth level.
func (v *agentState) IncrementDepth() int {
	v.depth++
	if v.depth > maximumRecursion {
		panic(fmt.Sprintf("The maximum recursion depth was exceeded: %v", maximumRecursion))
	}
	return v.depth
}

// This method returns the current depth level of recursion.
func (v *agentState) CurrentDepth() int {
	return v.depth
}

// This method decreases the depth level of recursion by one and returns the new
// depth level.
func (v *agentState) DecrementDepth() int {
	v.depth--
	return v.depth
}

// FORMATTER STATE

// This interface defines the methods that all formatter states must implement.
type FormatterStateLike interface {
	AgentStateLike
	GetIndentation() int
	AppendString(s string)
	AppendNewline()
	GetResult() string
}

// This constructor creates a new formatter state with no indentation.
func FormatterState() FormatterStateLike {
	return FormatterStateWithIndentation(0)
}

// This constructor creates a new formatter state with the specified
// indentation.
func FormatterStateWithIndentation(indentation int) FormatterStateLike {
	return &formatterState{AgentStateLike: AgentState(), indentation: indentation}
}

// This type defines the structure and methods for the state of any formatting
// agent. The formatted state uses a string builder to maintain the growing
// formatted string.
type formatterState struct {
	AgentStateLike
	indentation int
	result      sts.Builder
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatterState) GetIndentation() int {
	return v.indentation
}

// This method appends the specified string to the result.
func (v *formatterState) AppendString(s string) {
	v.result.WriteString(s)
}

// This method appends a properly indented newline to the result.
func (v *formatterState) AppendNewline() {
	var separator = "\n"
	var levels = v.CurrentDepth() + v.GetIndentation()
	for level := 0; level < levels; level++ {
		separator += "    "
	}
	v.result.WriteString(separator)
}

// This method returns the canonically formatted string result.
func (v *formatterState) GetResult() string {
	var result = v.result.String()
	v.result.Reset()
	return result
}
