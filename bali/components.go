/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
	cof "github.com/craterdog/go-collection-framework/v2"
	sts "strings"
)

// UNIVERSAL CONSTRUCTORS

// This constructor returns a new component initialized with the specified
// value.
func Component(value abs.Value) abs.ComponentLike {
	var component abs.ComponentLike
	switch actual := value.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		component = com.Component(Number(actual))
	case string:
		component = com.Component(str.QuoteFromString(actual))
	case abs.ComponentLike:
		component = actual
	default:
		component = com.Component(actual)
	}
	return component
}

// PRIVATE METHODS

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
	var iterator = cof.Iterator[string](comment)
	for iterator.HasNext() {
		var line = iterator.GetNext()
		if len(line) > 0 {
			v.AppendNewline()
			v.AppendString("    " + line)
		} else {
			v.AppendString("\n")
		}
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
	if !ok {
		return component, token, false
	}
	context, token, _ = v.parseContext()   // The context is optional.
	entity = adjustEntity(entity, context) // Set the real collection type.
	note, token, _ = v.parseNote()         // The note is optional.
	component = com.ComponentWithContext(entity, context)
	component.SetNote(note)
	return component, token, true
}

// This method adds the canonical format for the specified component to the
// state of the formatter.
func (v *formatter) formatComponent(component abs.ComponentLike) {
	var entity = component.GetEntity()
	var context = adjustContext(component)
	v.formatEntity(entity)
	if component.IsParameterized() {
		v.formatContext(context)
	}
	if component.IsAnnotated() {
		v.AppendString("  ")
		var note = component.GetNote()
		v.formatNote(note)
	}
}

// This method attempts to parse a component context. It returns the
// component context and whether or not the component context was
// successfully parsed.
func (v *parser) parseContext() (abs.ContextLike, *Token, bool) {
	var ok bool
	var token *Token
	var context abs.ContextLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		return context, token, false
	}
	_, token, ok = v.parseEOL()
	if !ok {
		context, token, ok = v.parseInlineContext()
		if !ok {
			v.backupOne() // Put back the '(' character.
			return context, token, false
		}
	} else {
		context, token, ok = v.parseMultilineContext()
		if !ok {
			v.backupOne() // Put back the EOL character.
			v.backupOne() // Put back the '(' character.
			return context, token, false
		}
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar(")",
			"$context",
			"$parameters",
			"$parameter",
			"$name",
			"$value")
		panic(message)
	}
	return context, token, true
}

// This method adds the canonical format for the specified context to the
// state of the formatter.
func (v *formatter) formatContext(context abs.ContextLike) {
	v.AppendString("(")
	var names = context.GetNames()
	var iterator = cof.Iterator[abs.SymbolLike](names)
	switch names.GetSize() {
	case 1:
		var name = iterator.GetNext()
		var value = context.GetValue(name)
		var parameter = com.Parameter(name, value)
		v.formatParameter(parameter)
	default:
		v.depth++
		for iterator.HasNext() {
			v.AppendNewline()
			var name = iterator.GetNext()
			var value = context.GetValue(name)
			var parameter = com.Parameter(name, value)
			v.formatParameter(parameter)
		}
		v.depth--
		v.AppendNewline()
	}
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
	// The order of these cases is very important since Go only compares the
	// set of methods supported by each interface. An interface that is a subset
	// of another interface must be checked AFTER that interface.
	case abs.BinaryLike:
		v.formatBinary(value)
	case abs.MonikerLike:
		v.formatMoniker(value)
	case abs.NarrativeLike:
		v.formatNarrative(value)
	case abs.QuoteLike:
		v.formatQuote(value)
	case abs.VersionLike:
		v.formatVersion(value)
	case abs.DurationLike:
		v.formatDuration(value)
	case abs.MomentLike:
		v.formatMoment(value)
	case abs.NumberLike:
		v.formatNumber(value)
	case abs.PercentageLike:
		v.formatPercentage(value)
	case abs.ProbabilityLike:
		v.formatProbability(value)
	case abs.AngleLike:
		v.formatAngle(value)
	case abs.BooleanLike:
		v.formatBoolean(value)
	case abs.PatternLike:
		v.formatPattern(value)
	case abs.ResourceLike:
		v.formatResource(value)
	case abs.TagLike:
		v.formatTag(value)
	case abs.SymbolLike:
		v.formatSymbol(value)
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
	case abs.ProcedureLike:
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

// It returns the component context and whether or not the component context
// was successfully parsed.
func (v *parser) parseInlineContext() (abs.ContextLike, *Token, bool) {
	var ok bool
	var token *Token
	var context abs.ContextLike
	var parameter abs.ParameterLike
	_, token, ok = v.parseDelimiter(":")
	if ok {
		// A context must have at least one parameter.
		var message = v.formatError(token)
		message += generateGrammar("$parameter",
			"$context",
			"$parameters",
			"$parameter",
			"$name",
			"$value")
		panic(message)
	}
	parameter, token, ok = v.parseParameter()
	if !ok {
		// This is not an inline context.
		return context, token, false
	}
	context = com.Context()
	for {
		var name = parameter.GetKey()
		var value = parameter.GetValue()
		context.SetValue(name, value)
		// Every subsequent parameter must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more parameters.
			break
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$parameter",
				"$context",
				"$parameters",
				"$parameter",
				"$name",
				"$value")
			panic(message)
		}
	}
	return context, token, true
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

// This method attempts to parse a component context with multiline parameters.
// It returns the component context and whether or not the component context
// was successfully parsed.
func (v *parser) parseMultilineContext() (abs.ContextLike, *Token, bool) {
	var ok bool
	var token *Token
	var parameter abs.ParameterLike
	var context abs.ContextLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		// A context must have at least one parameter.
		var message = v.formatError(token)
		message += generateGrammar("parameter",
			"$context",
			"$parameters",
			"$parameter",
			"$name",
			"$value")
		panic(message)
	}
	context = com.Context()
	for {
		var name = parameter.GetKey()
		var value = parameter.GetValue()
		context.SetValue(name, value)
		// Every parameter must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("EOL",
				"$context",
				"$parameters",
				"$parameter",
				"$name",
				"$value")
			panic(message)
		}
		parameter, token, ok = v.parseParameter()
		if !ok {
			// No more parameters.
			break
		}
	}
	return context, token, true
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
	var name abs.SymbolLike
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
	v.formatIdentifier(key.AsString())
	v.AppendString(": ")
	var value = parameter.GetValue()
	v.formatComponent(value)
}

// PRIVATE FUNCTIONS

// This function removes the first and last line delimiters (shown as "xx")
// and the indentation from each line of the specified multi-line string.
//
// The following is an indented string with dashes showing the indentation:
//
//	----xx
//	--------This is the first line
//	--------of a multi-line
//	--------indented string.
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
	var last = lines[size]
	var indentation = len(last) + 2 // The number of spaces in the last line plus four more.
	lines = lines[1:size]           // Strip off the first and last delimitier lines.
	for _, line := range lines {
		if len(line) < indentation {
			trimmed += "\n" // This is an empty line.
		} else {
			trimmed += line[indentation:] + "\n" // Strip off the indentation.
		}
	}
	return trimmed[:len(trimmed)-1] // Strip off the extra end-of-line character.
}

// This function checks to see if the entity is a collection and if so adjusts
// it to be the right collection type.
func adjustEntity(entity abs.Entity, context abs.ContextLike) abs.Entity {
	// Check for an explicit component type.
	var type_ string
	if context != nil {
		var name = ele.SymbolFromString("type")
		var component = context.GetValue(name)
		if component != nil {
			var moniker = component.ExtractMoniker()
			type_ = moniker.AsString()
		}
	}
	// Check for a collection entity.
	switch entity.(type) {
	case abs.SeriesLike:
		// The type is a series of components.
		var sequence = entity.(abs.Sequential[abs.ComponentLike])
		switch type_ {
		case "/bali/collections/Set/v1":
			// The series type is a set.
			entity = col.SetFromSequence(sequence)
		case "/bali/collections/Queue/v1":
			// The series type is a queue.
			entity = col.QueueFromSequence(sequence)
		case "/bali/collections/Stack/v1":
			// The series type is a stack.
			entity = col.StackFromSequence(sequence)
		default:
			// The default series type is a list.
			entity = col.ListFromSequence(sequence)
		}
	case abs.MappingLike:
		// The mapping type is a catalog.
		var sequence = entity.(abs.Sequential[abs.AssociationLike])
		entity = col.CatalogFromSequence(sequence)
	default:
		// The entity is not a collection.
	}
	return entity
}

// This function checks to make sure the context for any collection component has
// the right type parameter.
func adjustContext(component abs.ComponentLike) abs.ContextLike {
	var type_ string
	var entity = component.GetEntity()
	var context = component.GetContext()
	switch entity.(type) {
	case abs.QueueLike:
		type_ = "/bali/collections/Queue/v1"
	case abs.SetLike:
		type_ = "/bali/collections/Set/v1"
	case abs.StackLike:
		type_ = "/bali/collections/Stack/v1"
	}
	if type_ != "" {
		if context == nil {
			context = com.Context()
			component.SetContext(context)
		}
		var name = ele.SymbolFromString("type")
		var value = com.Component(str.MonikerFromString(type_))
		context.SetValue(name, value)
	}
	return context
}
