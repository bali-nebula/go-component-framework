/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language

import (
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ref "reflect"
	stc "strconv"
	str "strings"
)

// FORMATTER INTERFACE

// This function returns a canonical BDN string for the specified value.
func FormatValue(value any) string {
	var v = Formatter()
	return v.FormatValue(value)
}

// This constructor creates a new instance of a formatter that can be used to
// generate the canonical string format for any value.
func Formatter() abs.FormatterLike {
	return FormatterWithIndentation(0)
}

// This constructor creates a new instance of a formatter that can be used to
// generate the canonical string format for any value using the specified
// initial indentation level.
func FormatterWithIndentation(indentation int) abs.FormatterLike {
	return &formatter{abs.FormatterStateWithIndentation(indentation)}
}

// FORMATTER IMPLEMENTATION

// This type defines the structure and methods for a canonical formatting agent.
type formatter struct {
	state abs.FormatterStateLike
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatter) GetIndentation() int {
	return v.state.GetIndentation()
}

// This method returns the canonical string for the specified value.
func (v *formatter) FormatValue(value any) string {
	v.formatValue(ref.ValueOf(value))
	return v.state.GetResult()
}

// The maximum recursion level handles overly complex data models and cyclical
// references in a clean and efficient way.
const maximumFormattingDepth int = 8

// This private method determines the actual type of the specified value and
// calls the corresponding format function for that type. Note that because the
// Go language doesn't really support polymorphism the selection of the actual
// function called must be done explicitly using reflection and a type switch.
func (v *formatter) formatValue(value ref.Value) {
	switch value.Kind() {

	// Handle all primitive types.
	case ref.Bool:
		v.formatBoolean(value)
	case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64, ref.Uintptr:
		v.formatUnsigned(value)
	case ref.Int, ref.Int8, ref.Int16, ref.Int64:
		v.formatInteger(value)
	case ref.Float32, ref.Float64:
		v.formatFloat(value)
	case ref.Complex64, ref.Complex128:
		v.formatComplex(value)
	case ref.Int32:
		v.formatRune(value)
	case ref.String:
		v.formatString(value)

	// Handle all primitive collections.
	case ref.Array, ref.Slice:
		v.formatArray(value, "array")
	case ref.Map:
		v.formatMap(value)

	// Handle all structures.
	case ref.Struct:
		v.formatStructure(value)

	// Handle all interfaces and pointers.
	case ref.Interface, ref.Pointer:
		if value.MethodByName("GetParameter").IsValid() {
			// The value is a component.
			v.formatComponent(value)
		} else if value.MethodByName("GetKey").IsValid() {
			// The value is an association.
			v.formatAssociation(value)
		} else if value.MethodByName("AsArray").IsValid() {
			// The value is a collection.
			v.formatCollection(value)
		} else {
			// The value is a pointer to the value to be formatted (or type 'any').
			value = value.Elem()
			v.formatValue(value)
		}

	default:
		if !value.IsValid() {
			v.formatNil(value)
		} else {
			panic(fmt.Sprintf(
				"Attempted to format:\n\tvalue: %v\n\ttype: %v\n\tkind: %v\n",
				value.Interface(),
				value.Type(),
				value.Kind()))
		}
	}
}

// This private method appends the string for the specified component to
// the result.
func (v *formatter) formatComponent(r ref.Value) {
}

// This private method appends the nil string for the specified value to the
// result.
func (v *formatter) formatNil(r ref.Value) {
	v.state.AppendString("<nil>")
}

// This private method appends the name of the specified boolean value to the
// result.
func (v *formatter) formatBoolean(r ref.Value) {
	var b = r.Bool()
	v.state.AppendString(stc.FormatBool(b))
}

// This private method appends the base 10 string for the specified integer
// value to the result.
func (v *formatter) formatInteger(r ref.Value) {
	var i = r.Int()
	v.state.AppendString(stc.FormatInt(int64(i), 10))
}

// This private method appends the base 10 string for the specified unsigned
// integer value to the result.
func (v *formatter) formatUnsigned(r ref.Value) {
	var u = r.Uint()
	v.state.AppendString("0x" + stc.FormatUint(uint64(u), 16))
}

// This private method appends the base 10 string for the specified floating
// point value to the result using scientific notation if necessary.
func (v *formatter) formatFloat(r ref.Value) {
	var flt = r.Float()
	var str = stc.FormatFloat(flt, 'G', -1, 64)
	if !str.Contains(str, ".") && !str.Contains(str, "E") {
		str += ".0"
	}
	v.state.AppendString(str)
}

// This private method appends the base 10 string for the specified complex
// number value to the result using scientific notation if necessary.
func (v *formatter) formatComplex(r ref.Value) {
	var c = r.Complex()
	v.state.AppendString(stc.FormatComplex(c, 'G', -1, 128))
}

// This private method appends the string for the specified rune value to the
// result.
func (v *formatter) formatRune(r ref.Value) {
	var rn = r.Int()
	v.state.AppendString(stc.QuoteRune(int32(rn)))
}

// This private method appends the string for the specified string value to the
// result.
func (v *formatter) formatString(r ref.Value) {
	var str = r.String()
	v.state.AppendString("\"" + str + "\"")
}

// This private method appends the string for the specified array of values to
// the result.
func (v *formatter) formatArray(r ref.Value, typ string) {
	var size = r.Len()
	v.state.AppendString("[")
	if size > 0 {
		if v.state.CurrentDepth()+1 > maximumFormattingDepth {
			// Truncate the recursion.
			v.state.AppendString("...")
		} else {
			for i := 0; i < size; i++ {
				v.state.IncrementDepth()
				v.state.AppendNewline()
				var item = r.Index(i)
				v.formatValue(item)
				v.state.DecrementDepth()
			}
			v.state.AppendNewline()
		}
	} else {
		if typ == "catalog" {
			v.state.AppendString(":") // The array of associations is empty: [:]
		} else {
			v.state.AppendString(" ") // The array of items is empty: [ ]
		}
	}
	v.state.AppendString("](" + typ + ")")
}

// This private method appends the string for the specified map of key-value
// pairs to the result.
func (v *formatter) formatMap(r ref.Value) {
	var keys = r.MapKeys()
	var size = len(keys)
	v.state.AppendString("[")
	if size > 0 {
		if v.state.CurrentDepth()+1 > maximumFormattingDepth {
			// Truncate the recursion.
			v.state.AppendString("...")
		} else {
			for i := 0; i < size; i++ {
				v.state.IncrementDepth()
				v.state.AppendNewline()
				var key = keys[i]
				var value = r.MapIndex(key)
				v.formatValue(key)
				v.state.AppendString(": ")
				v.formatValue(value)
				v.state.DecrementDepth()
			}
			v.state.AppendNewline()
		}
	} else {
		v.state.AppendString(":") // The map is empty: [:]
	}
	v.state.AppendString("](map)")
}

// This private method appends the string for the specified structure to
// the result.
func (v *formatter) formatStructure(r ref.Value) {
}

// This private method appends the string for the specified catalog of
// key-value pairs to the result. It uses recursion to format each pair.
func (v *formatter) formatAssociation(r ref.Value) {
	var key = r.MethodByName("GetKey").Call([]ref.Value{})[0]
	v.formatValue(key)
	v.state.AppendString(": ")
	var value = r.MethodByName("GetValue").Call([]ref.Value{})[0]
	v.formatValue(value)
}

// This private method appends the string for the specified collection of
// items to the result. It uses recursion to format each item.
func (v *formatter) formatCollection(r ref.Value) {
	var array = r.MethodByName("AsArray").Call([]ref.Value{})[0]
	var typ = extractType(r)
	v.formatArray(array, typ)
}

// PRIVATE FUNCTIONS

// This private function extracts the type name string from the full reflected
// type.
func extractType(r ref.Value) string {
	var t = r.Type().String()
	switch {
	case str.HasPrefix(t, "[]"):
		return "array"
	case str.HasPrefix(t, "map["):
		return "map"
	case str.HasPrefix(t, "*collections.set"):
		return "set"
	case str.HasPrefix(t, "abs.SetLike"):
		return "set"
	case str.HasPrefix(t, "*collections.queue"):
		return "queue"
	case str.HasPrefix(t, "abs.QueueLike"):
		return "queue"
	case str.HasPrefix(t, "*collections.stack"):
		return "stack"
	case str.HasPrefix(t, "abs.StackLike"):
		return "stack"
	case str.HasPrefix(t, "*collections.list"):
		return "list"
	case str.HasPrefix(t, "abs.ListLike"):
		return "list"
	case str.HasPrefix(t, "*collections.catalog"):
		return "catalog"
	case str.HasPrefix(t, "abs.CatalogLike"):
		return "catalog"
	default:
		fmt.Printf("UNKNOWN: %v\n", t)
		return "unknown"
	}
}
