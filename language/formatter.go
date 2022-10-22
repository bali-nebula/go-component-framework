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
	st2 "github.com/craterdog-bali/go-bali-document-notation/strings"
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
	v.formatAny(value)
	return v.state.GetResult()
}

// PRIVATE METHODS

// The maximum recursion level handles overly complex data models and cyclical
// references in a clean and efficient way.
const maximumFormattingDepth int = 8

// This private method determines the actual type of the specified value and
// calls the corresponding format function for that type. Note that because the
// Go language doesn't really support polymorphism the selection of the actual
// function called must be done explicitly using reflection and a type switch.
func (v *formatter) formatAny(value any) {
	var r = ref.ValueOf(value)
	switch r.Kind() {

	// Handle all primitive types.
	case ref.Bool:
		var boolean, ok = value.(ele.Boolean)
		if !ok {
			boolean = ele.Boolean(r.Bool())
		}
		v.formatBoolean(boolean)
	case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64:
		var number = ele.Number(complex(r.Uint(), 0))
		v.formatNumber(number)
	case ref.Int, ref.Int8, ref.Int16, ref.Int32, ref.Int64:
		var integer, ok = value.(ele.Duration)
		if ok {
			v.formatDuration(integer)
		} else {
			integer, ok = value.(ele.Moment)
			if ok {
				v.formatMoment(integer)
			} else {
				integer = ele.Number(complex(r.Int(), 0))
				v.formatNumber(integer)
			}
		}
	case ref.Float32, ref.Float64:
		var float, ok = value.(ele.Angle)
		if ok {
			v.formatAngle(float)
		} else {
			float, ok = value.(ele.Percentage)
			if ok {
				v.formatPercentage(float)
			} else {
				float, ok = value.(ele.Probability)
				if ok {
					v.formatProbability(float)
				} else {
					float = ele.Number(complex(r.Float(), 0))
					v.formatNumber(float)
				}
			}
		}
	case ref.Complex64, ref.Complex128:
		var number, ok = value.(ele.Number)
		if ok {
			v.formatNumber(number)
		} else {
			number = ele.Number(r.Complex())
			v.formatNumber(number)
		}
	case ref.String:
		var s, ok = value.(ele.Pattern)
		if ok {
			v.formatPattern(s)
		} else {
			s, ok = value.(ele.Resource)
			if ok {
				v.formatResource(s)
			} else {
				s, ok = value.(ele.Symbol)
				if ok {
					v.formatSymbol(s)
				} else {
					s, ok = value.(ele.Tag)
					if ok {
						v.formatTag(s)
					} else {
						s, ok = value.(st2.Binary)
						if ok {
							v.formatBinary(s)
						} else {
							s, ok = value.(st2.Moniker)
							if ok {
								v.formatMoniker(s)
							} else {
								s, ok = value.(st2.Narrative)
								if ok {
									v.formatNarrative(s)
								} else {
									s, ok = value.(st2.Quote)
									if ok {
										v.formatQuote(s)
									} else {
										s, ok = value.(st2.Version)
										if ok {
											v.formatVersion(s)
										} else {
											s = st2.QuoteFromString(r.String())
											v.formatQuote(s)
										}
									}
								}
							}
						}
					}
				}
			}
		}

	// Handle all primitive collections.
	case ref.Array, ref.Slice:
		v.formatArray(r)
	case ref.Map:
		v.formatMap(r)

	// Handle all structures.
	case ref.Struct:
		v.formatStructure(r)

	// Handle all interfaces and pointers.
	case ref.Interface, ref.Pointer:
		if r.MethodByName("GetParameter").IsValid() {
			// The value is a component.
			var component = r.(abs.ComponentLike)
			v.formatComponent(component)
		} else if r.MethodByName("GetKey").IsValid() {
			// The value is an association.
			var association = r.(abs.AssociationLike)
			v.formatAssociation(association)
		} else if r.MethodByName("AsArray").IsValid() {
			// The value is a collection.
			var array = r.MethodByName("AsArray").Call([]ref.Value{})[0]
			v.formatArray(array)
		} else {
			// The value is a pointer to the value to be formatted (or type 'any').
			r = r.Elem()
			v.formatAny(r)
		}

	default:
		if !r.IsValid() {
			v.formatNone()
		} else {
			panic(fmt.Sprintf(
				"Attempted to format:\n\tvalue: %v\n\ttype: %v\n\tkind: %v\n",
				r.Interface(),
				r.Type(),
				r.Kind()))
		}
	}
}

// This private method appends the nil string for the specified value to the
// result.
func (v *formatter) formatNone() {
	v.state.AppendString("none")
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
				v.formatAny(item)
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

// This private method appends the string for the specified structure to
// the result.
func (v *formatter) formatStructure(r ref.Value) {
}
