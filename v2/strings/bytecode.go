/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	col "github.com/craterdog/go-collection-framework/v2"
	sts "strings"
)

// BYTECODE STRING INTERFACE

// This constructor creates a new bytecode string from the specified string.
func BytecodeFromString(string_ string) abs.BytecodeLike {
	var matches = uti.BytecodeMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a bytecode string from an invalid string: %v", string_)
		panic(message)
	}
	var encoded = matches[1]                    // Strip off the "'" delimiters.
	encoded = sts.ReplaceAll(encoded, " ", "")  // Remove all spaces.
	return bytecode_(encoded)
}

// This constructor creates a new bytecode string from the specified array of
// instructions.
func BytecodeFromArray(array []abs.Instruction) abs.BytecodeLike {
	var list = col.List[byte]()
	var instructions = col.ListFromArray(array)
	var iterator = col.Iterator[abs.Instruction](instructions)
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var leftByte = instruction.GetLeftByte()
		var rightByte = instruction.GetRightByte()
		list.AddValue(leftByte)
		list.AddValue(rightByte)
	}
	var bytes = list.AsArray()
	var encoded = uti.Base16Encode(bytes)
	return bytecode_(encoded)
}

// This constructor creates a new bytecode string from the specified sequence of
// instructions.
func BytecodeFromSequence(sequence abs.Sequential[abs.Instruction]) abs.BytecodeLike {
	var array = sequence.AsArray()
	return BytecodeFromArray(array)
}

// BYTECODE STRING IMPLEMENTATION

// This type defines the methods associated with a bytecode string that extends
// the native Go string type and represents the string of instructions that make up
// the bytecode string.
type bytecode_ string

// LEXICAL INTERFACE

// This method returns a string value for this lexical string.
func (v bytecode_) AsString() string {
	return "'" + string(v) + "'"
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this string is empty.
func (v bytecode_) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of instructions contained in this string.
func (v bytecode_) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the instructions in this string. The instructions retrieved
// are in the same order as they are in the string.
func (v bytecode_) AsArray() []abs.Instruction {
	var encoded = string(v)
	var decoded = uti.Base16Decode(encoded)
	var bytes = col.ListFromArray(decoded)
	var instructions = col.List[abs.Instruction]()
	var iterator = col.Iterator[byte](bytes)
	for iterator.HasNext() {
		var leftByte = iterator.GetNext()
		var rightByte = iterator.GetNext()
		var instruction = abs.InstructionFromBytes(leftByte, rightByte)
		instructions.AddValue(instruction)
	}
	return instructions.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this string the instruction that is associated
// with the specified index.
func (v bytecode_) GetValue(index int) abs.Instruction {
	var array = v.AsArray()
	var instructions = col.Array[abs.Instruction](array)
	return instructions.GetValue(index)
}

// This method retrieves from this string all instructions from the first index
// through the last index (inclusive).
func (v bytecode_) GetValues(first int, last int) abs.Sequential[abs.Instruction] {
	var array = v.AsArray()
	var instructions = col.Array[abs.Instruction](array)
	return instructions.GetValues(first, last)
}
