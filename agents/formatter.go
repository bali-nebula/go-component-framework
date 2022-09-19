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
	"reflect"
	"strconv"
	"strings"
)

// FORMATTER INTERFACE

// This function returns a canonical BDN string for the specified value.
func FormatValue(value any) string {
	v := Formatter()
	return v.FormatValue(value)
}

// This constructor creates a new instance of a formatter that can be used to
// generate the canonical string format for any value.
func Formatter() abstractions.FormatterLike {
	return FormatterWithIndentation(0)
}

// This constructor creates a new instance of a formatter that can be used to
// generate the canonical string format for any value using the specified
// initial indentation level.
func FormatterWithIndentation(indentation int) abstractions.FormatterLike {
	return &formatter{abstractions.FormatterStateWithIndentation(indentation)}
}

// FORMATTER IMPLEMENTATION

// This type defines the structure and methods for a canonical formatting agent.
type formatter struct {
	state abstractions.FormatterStateLike
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatter) GetIndentation() int {
	return v.state.GetIndentation()
}

// This method returns the canonical string for the specified value.
func (v *formatter) FormatValue(value any) string {
	v.formatValue(reflect.ValueOf(value))
	return v.state.GetResult()
}

// The maximum recursion level handles overly complex data models and cyclical
// references in a clean and efficient way.
const maximumFormattingDepth int = 8

// This private method determines the actual type of the specified value and
// calls the corresponding format function for that type. Note that because the
// Go language doesn't really support polymorphism the selection of the actual
// function called must be done explicitly using reflection and a type switch.
func (v *formatter) formatValue(value reflect.Value) {
	switch value.Kind() {

	// Handle all primitive types.
	case reflect.Bool:
		v.formatBoolean(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.formatUnsigned(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64:
		v.formatInteger(value)
	case reflect.Float32, reflect.Float64:
		v.formatFloat(value)
	case reflect.Complex64, reflect.Complex128:
		v.formatComplex(value)
	case reflect.Int32:
		v.formatRune(value)
	case reflect.String:
		v.formatString(value)

	// Handle all primitive collections.
	case reflect.Array, reflect.Slice:
		v.formatArray(value, "array")
	case reflect.Map:
		v.formatMap(value)

	// Handle all sequential collections.
	case reflect.Interface, reflect.Pointer:
		if value.MethodByName("AsArray").IsValid() {
			// The value is a collection.
			v.formatCollection(value)
		} else if value.MethodByName("GetKey").IsValid() {
			// The value is an association.
			v.formatAssociation(value)
		} else {
			// The value is a pointer to the value to be formatted.
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

// This private method appends the nil string for the specified value to the
// result.
func (v *formatter) formatNil(r reflect.Value) {
	v.state.AppendString("<nil>")
}

// This private method appends the name of the specified boolean value to the
// result.
func (v *formatter) formatBoolean(r reflect.Value) {
	b := r.Bool()
	v.state.AppendString(strconv.FormatBool(b))
}

// This private method appends the base 10 string for the specified integer
// value to the result.
func (v *formatter) formatInteger(r reflect.Value) {
	i := r.Int()
	v.state.AppendString(strconv.FormatInt(int64(i), 10))
}

// This private method appends the base 10 string for the specified unsigned
// integer value to the result.
func (v *formatter) formatUnsigned(r reflect.Value) {
	u := r.Uint()
	v.state.AppendString("0x" + strconv.FormatUint(uint64(u), 16))
}

// This private method appends the base 10 string for the specified floating
// point value to the result using scientific notation if necessary.
func (v *formatter) formatFloat(r reflect.Value) {
	flt := r.Float()
	str := strconv.FormatFloat(flt, 'G', -1, 64)
	if !strings.Contains(str, ".") && !strings.Contains(str, "E") {
		str += ".0"
	}
	v.state.AppendString(str)
}

// This private method appends the base 10 string for the specified complex
// number value to the result using scientific notation if necessary.
func (v *formatter) formatComplex(r reflect.Value) {
	c := r.Complex()
	v.state.AppendString(strconv.FormatComplex(c, 'G', -1, 128))
}

// This private method appends the string for the specified rune value to the
// result.
func (v *formatter) formatRune(r reflect.Value) {
	rn := r.Int()
	v.state.AppendString(strconv.QuoteRune(int32(rn)))
}

// This private method appends the string for the specified string value to the
// result.
func (v *formatter) formatString(r reflect.Value) {
	str := r.String()
	v.state.AppendString("\"" + str + "\"")
}

// This private method appends the string for the specified array of values to
// the result.
func (v *formatter) formatArray(r reflect.Value, typ string) {
	size := r.Len()
	v.state.AppendString("[")
	if size > 0 {
		if v.state.CurrentDepth()+1 > maximumFormattingDepth {
			// Truncate the recursion.
			v.state.AppendString("...")
		} else {
			for i := 0; i < size; i++ {
				v.state.IncrementDepth()
				v.state.AppendNewline()
				item := r.Index(i)
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
func (v *formatter) formatMap(r reflect.Value) {
	keys := r.MapKeys()
	size := len(keys)
	v.state.AppendString("[")
	if size > 0 {
		if v.state.CurrentDepth()+1 > maximumFormattingDepth {
			// Truncate the recursion.
			v.state.AppendString("...")
		} else {
			for i := 0; i < size; i++ {
				v.state.IncrementDepth()
				v.state.AppendNewline()
				key := keys[i]
				value := r.MapIndex(key)
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

// This private method appends the string for the specified catalog of
// key-value pairs to the result. It uses recursion to format each pair.
func (v *formatter) formatAssociation(r reflect.Value) {
	key := r.MethodByName("GetKey").Call([]reflect.Value{})[0]
	v.formatValue(key)
	v.state.AppendString(": ")
	value := r.MethodByName("GetValue").Call([]reflect.Value{})[0]
	v.formatValue(value)
}

// This private method appends the string for the specified collection of
// items to the result. It uses recursion to format each item.
func (v *formatter) formatCollection(r reflect.Value) {
	array := r.MethodByName("AsArray").Call([]reflect.Value{})[0]
	typ := extractType(r)
	v.formatArray(array, typ)
}

// PRIVATE FUNCTIONS

// This private function extracts the type name string from the full reflected
// type.
func extractType(r reflect.Value) string {
	t := r.Type().String()
	switch {
	case strings.HasPrefix(t, "[]"):
		return "array"
	case strings.HasPrefix(t, "map["):
		return "map"
	case strings.HasPrefix(t, "*collections.set"):
		return "set"
	case strings.HasPrefix(t, "abstractions.SetLike"):
		return "set"
	case strings.HasPrefix(t, "*collections.queue"):
		return "queue"
	case strings.HasPrefix(t, "abstractions.QueueLike"):
		return "queue"
	case strings.HasPrefix(t, "*collections.stack"):
		return "stack"
	case strings.HasPrefix(t, "abstractions.StackLike"):
		return "stack"
	case strings.HasPrefix(t, "*collections.list"):
		return "list"
	case strings.HasPrefix(t, "abstractions.ListLike"):
		return "list"
	case strings.HasPrefix(t, "*collections.catalog"):
		return "catalog"
	case strings.HasPrefix(t, "abstractions.CatalogLike"):
		return "catalog"
	default:
		fmt.Printf("UNKNOWN: %v\n", t)
		return "unknown"
	}
}
