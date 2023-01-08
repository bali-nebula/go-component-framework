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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
	col "github.com/craterdog/go-collection-framework"
	sts "strings"
)

// PARSING METHODS
//
// Each parsing method returns the following three results:
//  * The parsed component corresponding to the source string, or nil if the
//    parsing failed.
//  * The parsed token that caused the failure if the parsing did fail.
//  * Whether or not the parsing succeeded.

// This method attempts to parse annotation. It returns the annotation and
// whether or not the annotation was successfully parsed.
func (v *parser) parseAnnotation() (abs.Annotation, *Token, bool) {
	var ok bool
	var token *Token
	var annotation abs.Annotation
	annotation, token, ok = v.parseNote()
	if !ok {
		annotation, token, ok = v.parseComment()
	}
	return annotation, token, ok
}

// This method adds the canonical format for the specified annotation to the
// state of the formatter.
func (v *formatter) formatAnnotation(annotation abs.Annotation) {
	switch value := annotation.(type) {
	case abs.NoteLike:
		v.formatNote(value)
	case abs.CommentLike:
		v.formatComment(value)
	default:
		panic(fmt.Sprintf("An invalid annotation (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a comment. It returns a string containing the
// comment and whether or not the comment was successfully parsed.
func (v *parser) parseComment() (abs.CommentLike, *Token, bool) {
	var comment abs.CommentLike
	var token = v.nextToken()
	if token.Type != TokenComment {
		v.backupOne()
		return comment, token, false
	}
	comment = com.Comment(trimIndentation(token.Value))
	return comment, token, true
}

// This method adds the canonical format for the specified annotation to the
// state of the formatter.
func (v *formatter) formatComment(comment abs.CommentLike) {
	v.AppendString(`!>`)
	var iterator = col.Iterator[string](comment)
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.AppendNewline()
		v.AppendString(line)
	}
	v.AppendNewline()
	v.AppendString(`<!`)
}

// This method attempts to parse a component. It returns the component and
// whether or not the component was successfully parsed.
func (v *parser) parseComponent() (abs.ComponentLike, *Token, bool) {
	var component abs.ComponentLike
	var context abs.ContextLike
	var note abs.NoteLike
	var entity, token, ok = v.parseEntity()
	if ok {
		context, token, _ = v.parseContext() // The context is optional.
		note, token, _ = v.parseNote()       // The note is optional.
		component = com.ComponentWithContext(entity, context)
		component.SetNote(note)
	}
	return component, token, ok
}

// This method adds the canonical format for the specified component to the
// state of the formatter.
func (v *formatter) formatComponent(component abs.ComponentLike) {
	var entity = component.GetEntity()
	v.formatEntity(entity)
	if component.IsParameterized() {
		var context = component.GetContext()
		v.formatContext(context)
	}
	if component.IsAnnotated() {
		v.AppendString("  ")
		var note = component.GetNote()
		v.formatNote(note)
	}
}

// This method attempts to parse the context for a parameterized component. It
// returns an array of parameters and whether or not the context was
// successfully parsed.
func (v *parser) parseContext() (abs.ContextLike, *Token, bool) {
	var ok bool
	var token *Token
	var parameters abs.Sequential[abs.ParameterLike]
	var context abs.ContextLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		return context, token, false
	}
	parameters, token, ok = v.parseParameters()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$parameter",
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar(")",
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	context = com.Context()
	var iterator = col.Iterator[abs.ParameterLike](parameters)
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var name = parameter.GetKey()
		var value = parameter.GetValue()
		context.SetValue(name, value)
	}
	return context, token, true
}

// This method adds the canonical format for the specified context to the
// state of the formatter.
func (v *formatter) formatContext(context abs.ContextLike) {
	v.AppendString("(")
	var parameters = col.List[abs.ParameterLike]()
	var names = context.GetNames()
	var iterator = col.Iterator[abs.Symbolic](names)
	for iterator.HasNext() {
		var name = iterator.GetNext()
		var value = context.GetValue(name)
		var parameter = com.Parameter(name, value)
		parameters.AddValue(parameter)
	}
	v.formatParameters(parameters)
	v.AppendString(")")
}

// This method attempts to parse the specified delimiter. It returns
// the token and whether or not the delimiter was found.
func (v *parser) parseDelimiter(delimiter string) (string, *Token, bool) {
	var token = v.nextToken()
	if token.Type == TokenEOF || token.Value != delimiter {
		v.backupOne()
		return delimiter, token, false
	}
	return delimiter, token, true
}

// This method attempts to parse a component entity. It returns the component
// entity and whether or not the component entity was successfully parsed.
func (v *parser) parseEntity() (abs.Entity, *Token, bool) {
	var ok bool
	var token *Token
	var entity abs.Entity
	entity, token, ok = v.parseElement()
	if !ok {
		entity, token, ok = v.parseString()
	}
	if !ok {
		entity, token, ok = v.parseRange()
	}
	if !ok {
		entity, token, ok = v.parseCollection()
	}
	if !ok {
		entity, token, ok = v.parseProcedure()
	}
	return entity, token, ok
}

// This method adds the canonical format for the specified entity to the
// state of the formatter.
func (v *formatter) formatEntity(entity abs.Entity) {
	switch value := entity.(type) {
	case ele.Angle:
		v.formatAngle(value)
	case ele.Boolean:
		v.formatBoolean(value)
	case ele.Duration:
		v.formatDuration(value)
	case ele.Moment:
		v.formatMoment(value)
	case ele.Number:
		v.formatNumber(value)
	case ele.Pattern:
		v.formatPattern(value)
	case ele.Percentage:
		v.formatPercentage(value)
	case ele.Probability:
		v.formatProbability(value)
	case ele.Resource:
		v.formatResource(value)
	case ele.Symbol:
		v.formatSymbol(value)
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
	case abs.SeriesLike:
		v.formatSeries(value)
	case abs.MappingLike:
		v.formatMapping(value)
	case abs.IntervalLike[abs.Discrete]:
		v.formatInterval(value)
	case abs.SpectrumLike[abs.Spectral]:
		v.formatSpectrum(value)
	case abs.ContinuumLike[abs.Continuous]:
		v.formatContinuum(value)
	case abs.Procedural:
		v.formatProcedure(value)
	default:
		panic(fmt.Sprintf("An invalid entity (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse the end-of-file (EOF) marker. It returns
// the token and whether or not an EOL token was found.
func (v *parser) parseEOF() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOF {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse the end-of-line (EOL) marker. It returns
// the token and whether or not an EOF token was found.
func (v *parser) parseEOL() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOL {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseIdentifier() (string, *Token, bool) {
	var identifier string = "<UNKNOWN>"
	var token = v.nextToken()
	if token.Type != TokenIdentifier {
		v.backupOne()
		return identifier, token, false
	}
	identifier = token.Value
	return identifier, token, true
}

// This method adds the canonical format for the specified identifier to the
// state of the formatter.
func (v *formatter) formatIdentifier(identifier string) {
	v.AppendString(identifier)
}

// This method attempts to parse the specified keyword. It returns
// the token and whether or not the keyword was found.
func (v *parser) parseKeyword(keyword string) (string, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenKeyword || token.Value != keyword {
		v.backupOne()
		return keyword, token, false
	}
	return keyword, token, true
}

// This method attempts to parse a note. It returns a string containing the
// note and whether or not the note was successfully parsed.
func (v *parser) parseNote() (abs.NoteLike, *Token, bool) {
	var note abs.NoteLike
	var token = v.nextToken()
	if token.Type != TokenNote {
		v.backupOne()
		return note, token, false
	}
	note = com.Note(token.Value[2:]) // Remove the leading "! ".
	return note, token, true
}

// This method adds the canonical format for the specified annotation to the
// state of the formatter.
func (v *formatter) formatNote(note abs.NoteLike) {
	v.AppendString("! ")
	v.AppendString(string(note.AsArray()))
}

// This method attempts to parse a parameter containing a name and value. It
// returns the parameter and whether or not the parameter was successfully
// parsed.
func (v *parser) parseParameter() (abs.ParameterLike, *Token, bool) {
	var ok bool
	var token *Token
	var name abs.Symbolic
	var value abs.ComponentLike
	name, token, ok = v.parseSymbol()
	if !ok {
		// This is not a parameter.
		return nil, token, false
	}
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		// This is not a parameter.
		v.backupOne() // Put back the symbol token.
		return nil, token, false
	}
	// This must be a parameter.
	value, token, ok = v.parseComponent()
	if !ok {
		panic("Expected a value after the ':' character.")
	}
	var parameter = com.Parameter(name, value)
	return parameter, token, true
}

// This method adds the canonical format for the specified parameter to the
// state of the formatter.
func (v *formatter) formatParameter(parameter abs.ParameterLike) {
	var key = parameter.GetKey()
	v.AppendString("$")
	v.formatIdentifier(key.GetIdentifier())
	v.AppendString(": ")
	var value = parameter.GetValue()
	v.formatComponent(value)
}

// This method attempts to parse context parameters. It returns the
// context parameters and whether or not the context parameters were
// successfully parsed.
func (v *parser) parseParameters() (abs.Sequential[abs.ParameterLike], *Token, bool) {
	var ok bool
	var token *Token
	var parameter abs.ParameterLike
	var parameters = col.List[abs.ParameterLike]()
	_, token, ok = v.parseEOL()
	if !ok {
		// The parameters are on a single line.
		_, token, ok = v.parseDelimiter(":")
		if ok {
			panic("There must be at least one parameter in a context.")
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			panic("There must be at least one parameter in a context.")
		}
		for {
			parameters.AddValue(parameter)
			// Every subsequent parameter must be preceded by a ','.
			_, token, ok = v.parseDelimiter(",")
			if !ok {
				// No more parameters.
				break
			}
			parameter, token, ok = v.parseParameter()
			if !ok {
				panic("Expected a parameter after the ',' character.")
			}
		}
		return parameters, token, true
	}
	// The parameters are on separate lines.
	parameter, token, ok = v.parseParameter()
	if !ok {
		panic("There must be at least one parameter in a context.")
	}
	for {
		parameters.AddValue(parameter)
		// Every parameter must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the parameter.")
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			// No more parameters.
			break
		}
	}
	return parameters, token, true
}

// This method adds the canonical format for the specified parameters to the
// state of the formatter.
func (v *formatter) formatParameters(parameters abs.Sequential[abs.ParameterLike]) {
	var iterator = col.Iterator[abs.ParameterLike](parameters)
	switch parameters.GetSize() {
	case 0:
		panic("A context must have at least one parameter.")
	case 1:
		var parameter = iterator.GetNext()
		v.formatParameter(parameter)
	default:
		v.depth++
		for iterator.HasNext() {
			v.AppendNewline()
			var parameter = iterator.GetNext()
			v.formatParameter(parameter)
		}
		v.depth--
		v.AppendNewline()
	}
}

// PRIVATE FUNCTIONS

// This function removes the first and last line delimiters (shown as "xx")
// and the indentation from each line of the specified multi-line string.
//
// The following is an indented string with dashes showing the indentation:
//
//	----xx
//	----This is the first line
//	----of a multi-line
//	----indented string.
//	----xx
//
// It will be trimmed to:
//
//	This is the first line
//	of a multi-line
//	indented string.
//
// Only the indented lines remain and are no longer indented.
func trimIndentation(v string) string {
	var trimmed string
	var lines = sts.Split(v, "\n")
	var size = len(lines) - 1
	var last = lines[size]          // The last line provides the level of indentation.
	var indentation = len(last) - 2 // The number of spaces in the last line.
	lines = lines[1:size]           // Strip off the first and last delimitier lines.
	for _, line := range lines {
		trimmed += line[indentation:] + "\n" // Strip off the indentation.
	}
	return trimmed[:len(trimmed)-1] // Strip off the extra end-of-line character.
}
