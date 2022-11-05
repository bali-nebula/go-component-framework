/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ref "reflect"
)

// FORMATTER INTERFACE

// This function returns a canonical BDN string for the specified entity.
func FormatEntity(entity abs.EntityLike) string {
	var v = Formatter(0)
	return v.FormatEntity(entity)
}

// This function returns a canonical BDN string for the specified component.
func FormatComponent(component abs.ComponentLike) string {
	var v = Formatter(0)
	return v.FormatComponent(component)
}

// This function returns a canonical BDN string for the specified component
// using the specified indentation.
func FormatComponentWithIndentation(component abs.ComponentLike, indentation int) string {
	var v = Formatter(indentation)
	return v.FormatComponent(component)
}

// FORMATTER IMPLEMENTATION

// This constructor creates a new formatter using the specified indentation.
func Formatter(indentation int) *formatter {
	var v = &formatter{abs.FormatterStateWithIndentation(indentation)}
	return v
}

// This type defines the structure and methods for a canonical formatting agent.
type formatter struct {
	state abs.FormatterStateLike
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatter) GetIndentation() int {
	return v.state.GetIndentation()
}

// This method returns the canonical string for the specified entity.
func (v *formatter) FormatEntity(entity abs.EntityLike) string {
	v.formatEntity(entity)
	return v.state.GetResult()
}

// This method returns the canonical string for the specified component.
func (v *formatter) FormatComponent(component abs.ComponentLike) string {
	v.formatComponent(component)
	return v.state.GetResult()
}

// PRIVATE METHODS

// The maximum recursion level handles overly complex data models and cyclical
// references in a clean and efficient way.
const maximumFormattingDepth int = 8

// This method determines the actual type of the specified value and calls the
// corresponding format function for that type. Note that because the Go
// language doesn't really support polymorphism the selection of the actual
// function called must be done explicitly using reflection and a type switch.
func (v *formatter) formatAny(value any) {
	var r = ref.ValueOf(value)
	switch r.Kind() {

	// Handle all primitive types.
	case ref.Bool:
		v.formatBool(value)
	case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64:
		v.formatUnsigned(value)
	case ref.Int, ref.Int8, ref.Int16, ref.Int32, ref.Int64:
		v.formatInteger(value)
	case ref.Float32, ref.Float64:
		v.formatFloat(value)
	case ref.Complex64, ref.Complex128:
		v.formatComplex(value)
	case ref.String:
		v.formatString(value)

	// Handle all interfaces and pointers.
	case ref.Interface, ref.Pointer:
		v.formatInterface(value)
	case ref.Invalid:
		v.formatNil()

	// Handle all composites.
	case ref.Array, ref.Slice:
		v.formatArray(value)
	case ref.Map:
		v.formatMap(value)
	case ref.Struct:
		v.formatStructure(r)

	default:
		panic(fmt.Sprintf("Attempted to format:\n\tvalue: %v\n\ttype: %v\n\tkind: %v\n",
			r.Interface(),
			r.Type(),
			r.Kind()))
	}
}

// This method adds the canonical list format for the specified Go array to the
// state of the formatter.
func (v *formatter) formatArray(value any) {
	var r = ref.ValueOf(value)
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
				var value = r.Index(i)
				v.formatAny(value)
				v.state.DecrementDepth()
			}
			v.state.AppendNewline()
		}
	} else {
		v.state.AppendString(" ") // The array of values is empty: [ ]
	}
	v.state.AppendString("]")
}

// This method adds the canonical boolean format for the specified Go bool to
// the state of the formatter.
func (v *formatter) formatBool(value any) {
	boolean, ok := value.(ele.Boolean)
	if !ok {
		boolean = ele.Boolean(value.(bool))
	}
	v.formatBoolean(boolean)
}

// This method adds the canonical number format for the specified Go complex to
// the state of the formatter.
func (v *formatter) formatComplex(value any) {
	number, ok := value.(ele.Number)
	if ok {
		v.formatNumber(number)
		return
	}
	var c = ref.ValueOf(value).Complex()
	number = ele.Number(c)
	v.formatNumber(number)
}

// This method adds the canonical element format for the specified Go float to
// the state of the formatter.
func (v *formatter) formatFloat(value any) {
	angle, ok := value.(ele.Angle)
	if ok {
		v.formatAngle(angle)
		return
	}
	percentage, ok := value.(ele.Percentage)
	if ok {
		v.formatPercentage(percentage)
		return
	}
	probability, ok := value.(ele.Probability)
	if ok {
		v.formatProbability(probability)
		return
	}
	var float = ref.ValueOf(value).Float()
	var number = ele.Number(complex(float64(float), 0))
	v.formatNumber(number)
}

// This method adds the canonical element format for the specified Go int to
// the state of the formatter.
func (v *formatter) formatInteger(value any) {
	duration, ok := value.(ele.Duration)
	if ok {
		v.formatDuration(duration)
		return
	}
	moment, ok := value.(ele.Moment)
	if ok {
		v.formatMoment(moment)
		return
	}
	var integer = ref.ValueOf(value).Int()
	var number = ele.Number(complex(float64(integer), 0))
	v.formatNumber(number)
}

// This method adds the canonical entity format for the specified Go interface
// to the state of the formatter.
func (v *formatter) formatInterface(value any) {
	component, ok := value.(abs.ComponentLike)
	if ok {
		v.formatComponent(component)
		return
	}
	association, ok := value.(abs.AssociationLike[abs.KeyLike, abs.ComponentLike])
	if ok {
		v.formatAssociation(association)
		return
	}
	catalog, ok := value.(abs.CatalogLike[abs.KeyLike, abs.ComponentLike])
	if ok {
		v.formatCatalog(catalog)
		return
	}
	list, ok := value.(abs.ListLike[abs.ComponentLike])
	if ok {
		v.formatList(list)
		return
	}
	rng, ok := value.(abs.RangeLike[abs.PrimitiveLike])
	if ok {
		v.formatRange(rng)
		return
	}
	procedure, ok := value.(abs.ProcedureLike)
	if ok {
		v.formatProcedure(procedure)
		return
	}
	// The value is a pointer to the value to be formatted (or type 'any').
	var ptr = value.(*any)
	v.formatAny(*ptr)
}

// This method adds the canonical catalog format for the specified Go map to the
// state of the formatter.
func (v *formatter) formatMap(value any) {
	var r = ref.ValueOf(value)
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
				v.formatAny(key)
				v.state.AppendString(": ")
				v.formatAny(value)
				v.state.DecrementDepth()
			}
			v.state.AppendNewline()
		}
	} else {
		v.state.AppendString(":") // The map is empty: [:]
	}
	v.state.AppendString("](map)")
}

// This method adds the canonical pattern format for the specified Go nil to the
// state of the formatter.
func (v *formatter) formatNil() {
	v.state.AppendString("none")
}

// This method adds the canonical element format for the specified Go string to
// the state of the formatter.
func (v *formatter) formatString(s any) {
	switch value := s.(type) {
	case ele.Pattern:
		v.formatPattern(value)
	case ele.Symbol:
		v.formatSymbol(value)
	case ele.Resource:
		v.formatResource(value)
	case ele.Tag:
		v.formatTag(value)
	case str.Binary:
		v.formatBinary(value)
	case str.Moniker:
		v.formatMoniker(value)
	case str.Narrative:
		v.formatNarrative(value)
	case str.Quote:
		v.formatQuote(value)
	case str.Version:
		v.formatVersion(value)
	default:
		panic(fmt.Sprintf("An invalid string was passed to the formatter: %v", value))
	}
}

// This method adds the canonical entity format for the specified Go struct to
// the state of the formatter.
func (v *formatter) formatStructure(value any) {
	//TODO:
	//var r = ref.ValueOf(value)
	//...
}

// This method adds the canonical number format for the specified Go uint to the
// state of the formatter.
func (v *formatter) formatUnsigned(value any) {
	var unsigned = ref.ValueOf(value).Uint()
	var number = ele.Number(complex(float64(unsigned), 0))
	v.formatNumber(number)
}
