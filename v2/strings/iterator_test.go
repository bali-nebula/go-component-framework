/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestByteIterators(t *tes.T) {
	var bytes = []byte{1, 2, 3}
	var b1 = bytes[0]
	var b2 = bytes[1]
	var b3 = bytes[2]
	var binary = str.BinaryFromArray(bytes)
	var iterator = str.ByteIterator(binary)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, b1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, b1, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, b3, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, b3, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, b1, iterator.GetNext())
	ass.Equal(t, b2, iterator.GetNext())
}

func TestInstructionIterators(t *tes.T) {
	var instructions = []abs.Instruction{1, 2, 3}
	var i1 = instructions[0]
	var i2 = instructions[1]
	var i3 = instructions[2]
	var bytecode = str.BytecodeFromArray(instructions)
	/*
	var bytecode = str.BytecodeFromString("000100020003")
	*/
	var iterator = str.InstructionIterator(bytecode)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, i1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, i1, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, i3, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, i3, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, i1, iterator.GetNext())
	ass.Equal(t, i2, iterator.GetNext())
}

func TestNameIterators(t *tes.T) {
	var bali = abs.Name("bali")
	var types = abs.Name("types")
	var abstractions = abs.Name("abstractions")
	var moniker = str.MonikerFromString("/bali/types/abstractions")
	var iterator = str.NameIterator(moniker)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, bali, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, bali, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, abstractions, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, abstractions, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, bali, iterator.GetNext())
	ass.Equal(t, types, iterator.GetNext())
}

func TestLineIterators(t *tes.T) {
	var first = abs.Line("This is the first line...")
	var second = abs.Line("This is the second line...")
	var third = abs.Line("This is the last line.")
	var array = []abs.Line{first, second, third}
	var narrative = str.NarrativeFromArray(array)
	var iterator = str.LineIterator(narrative)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.Equal(t, second, iterator.GetNext())
}

func TestRuneIterators(t *tes.T) {
	var first = rune('a')
	var second = rune('b')
	var third = rune('c')
	var array = []rune{first, second, third}
	var quote = str.QuoteFromArray(array)
	var iterator = str.RuneIterator(quote)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.Equal(t, second, iterator.GetNext())
}

func TestOrdinalIterators(t *tes.T) {
	var first = abs.Ordinal('1')
	var second = abs.Ordinal('2')
	var third = abs.Ordinal('3')
	var array = []abs.Ordinal{first, second, third}
	var version = str.VersionFromArray(array)
	var iterator = str.OrdinalIterator(version)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetPrevious())
	iterator.ToSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetNext())
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, third, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, first, iterator.GetNext())
	ass.Equal(t, second, iterator.GetNext())
}
