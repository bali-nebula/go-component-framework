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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	str "github.com/bali-nebula/go-component-framework/v2/strings"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	stc "strconv"
	sts "strings"
	uni "unicode"
)

// UNIVERSAL CONSTRUCTORS

// This constructor returns a new binary string initialized with the specified
// value.
func Binary(value abs.Value) abs.BinaryLike {
	var binary abs.BinaryLike
	switch actual := value.(type) {
	case int:
		binary = str.BinaryFromBytes(uti.RandomBytes(actual))
	case []byte:
		binary = str.BinaryFromBytes(actual)
	case abs.Sequential[byte]:
		binary = str.BinaryFromSequence(actual)
	case string:
		binary = str.BinaryFromString(actual)
	case abs.ComponentLike:
		binary = actual.GetEntity().(abs.BinaryLike)
	default:
		var message = fmt.Sprintf(
			"The value (of type %T) cannot be converted to a binary string: %v",
			actual, actual)
		panic(message)
	}
	return binary
}

// This constructor returns a new bytecode string initialized with the specified
// value.
func Bytecode(value abs.Value) abs.BytecodeLike {
	var bytecode abs.BytecodeLike
	switch actual := value.(type) {
	case []abs.Instruction:
		bytecode = str.BytecodeFromArray(actual)
	case abs.Sequential[abs.Instruction]:
		bytecode = str.BytecodeFromSequence(actual)
	case string:
		bytecode = str.BytecodeFromString(actual)
	case abs.ComponentLike:
		bytecode = actual.GetEntity().(abs.BytecodeLike)
	default:
		var message = fmt.Sprintf(
			"The value (of type %T) cannot be converted to a bytecode string: %v",
			actual, actual)
		panic(message)
	}
	return bytecode
}

// This constructor returns a new narrative string initialized with the specified
// value.
func Narrative(value abs.Value) abs.NarrativeLike {
	var narrative abs.NarrativeLike
	switch actual := value.(type) {
	case []abs.Line:
		narrative = str.NarrativeFromArray(actual)
	case abs.Sequential[abs.Line]:
		narrative = str.NarrativeFromSequence(actual)
	case string:
		narrative = str.NarrativeFromString(actual)
	case abs.ComponentLike:
		narrative = actual.GetEntity().(abs.NarrativeLike)
	default:
		var message = fmt.Sprintf(
			"The value (of type %T) cannot be converted to a narrative string: %v",
			actual, actual)
		panic(message)
	}
	return narrative
}

// This constructor returns a new quote string initialized with the specified
// value.
func Quote(value abs.Value) abs.QuoteLike {
	var quote abs.QuoteLike
	switch actual := value.(type) {
	case []rune:
		quote = str.QuoteFromString(string(actual))
	case abs.Sequential[rune]:
		quote = str.QuoteFromSequence(actual)
	case string:
		quote = str.QuoteFromString(actual)
	case abs.ComponentLike:
		quote = actual.GetEntity().(abs.QuoteLike)
	default:
		var message = fmt.Sprintf(
			"The value (of type %T) cannot be converted to a quote string: %v",
			actual, actual)
		panic(message)
	}
	return quote
}

// This constructor returns a new symbol element initialized with the specified
// value.
func Symbol(value abs.Value) abs.SymbolLike {
	var symbol abs.SymbolLike
	switch actual := value.(type) {
	case string:
		symbol = str.SymbolFromString(actual)
	case abs.SymbolLike:
		symbol = actual
	case abs.ComponentLike:
		symbol = actual.GetEntity().(abs.SymbolLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a symbol: %v", actual, actual)
		panic(message)
	}
	return symbol
}

// This constructor returns a new random tag element of the default size.
func NewTag() abs.TagLike {
	return str.TagOfSize(0)
}

// This constructor returns a new tag element initialized with the specified
// value.
func Tag(value abs.Value) abs.TagLike {
	var tag abs.TagLike
	switch actual := value.(type) {
	case uint:
		tag = str.TagOfSize(int(actual))
	case uint8:
		tag = str.TagOfSize(int(actual))
	case uint16:
		tag = str.TagOfSize(int(actual))
	case uint32:
		tag = str.TagOfSize(int(actual))
	case uint64:
		tag = str.TagOfSize(int(actual))
	case int:
		tag = str.TagOfSize(int(actual))
	case int8:
		tag = str.TagOfSize(int(actual))
	case int16:
		tag = str.TagOfSize(int(actual))
	case int32:
		tag = str.TagOfSize(int(actual))
	case int64:
		tag = str.TagOfSize(int(actual))
	case []byte:
		tag = str.TagFromBytes(actual)
	case string:
		tag = str.TagFromString(actual)
	case abs.TagLike:
		tag = actual
	case abs.ComponentLike:
		tag = actual.GetEntity().(abs.TagLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a tag: %v", actual, actual)
		panic(message)
	}
	return tag
}

// This constructor returns a new version string initialized with the specified
// value.
func Version(value abs.Value) abs.VersionLike {
	var version abs.VersionLike
	switch actual := value.(type) {
	case []abs.Ordinal:
		version = str.VersionFromArray(actual)
	case abs.Sequential[abs.Ordinal]:
		version = str.VersionFromSequence(actual)
	case string:
		version = str.VersionFromString(actual)
	case abs.ComponentLike:
		version = actual.GetEntity().(abs.VersionLike)
	default:
		var message = fmt.Sprintf(
			"The value (of type %T) cannot be converted to a version string: %v",
			actual, actual)
		panic(message)
	}
	return version
}

// This constructor returns the next version string for the specified version.
func NextVersion(current abs.VersionLike) abs.VersionLike {
	return str.Versions.GetNextVersion(current, 0)
}

// PRIVATE METHODS

// This method attempts to parse a binary string. It returns the binary
// string and whether or not the binary string was successfully parsed.
func (v *parser) parseBinary() (abs.BinaryLike, *Token, bool) {
	var token *Token
	var binary abs.BinaryLike
	token = v.nextToken()
	if token.Type != TokenBINARY {
		v.backupOne(token)
		return binary, token, false
	}
	var matches = bytesToStrings(binaryScanner.FindSubmatch([]byte(token.Value)))
	// Remove all whitespace and the "'>" and "<'" delimiters.
	binary = str.BinaryFromString(sts.Map(func(r rune) rune {
		if uni.IsSpace(r) {
			return -1
		}
		return r
	}, matches[1]))
	return binary, token, true
}

const lineLength = 60 // 60 base 64 characters encode 45 bytes per line.

func (v *formatter) formatName(name abs.NameLike) {
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatBinary(binary abs.BinaryLike) {
	v.AppendString("'>")
	v.depth++
	var s = binary.AsString()
	var length = len(s)
	if length > 0 {
		for index := 0; index < length; {
			v.AppendNewline()
			var next = index + lineLength
			if next > length {
				next = length
			}
			v.AppendString(s[index:next])
			index = next
		}
	}
	v.depth--
	v.AppendNewline()
	v.AppendString("<'")
}

// This method attempts to parse a bytecode string. It returns the bytecode
// string and whether or not the bytecode string was successfully parsed.
func (v *parser) parseBytecode() (abs.BytecodeLike, *Token, bool) {
	var token *Token
	var bytecode abs.BytecodeLike
	token = v.nextToken()
	if token.Type != TokenBYTECODE {
		v.backupOne(token)
		return bytecode, token, false
	}
	var matches = bytesToStrings(bytecodeScanner.FindSubmatch([]byte(token.Value)))
	// Remove all whitespace and the "'" delimiters.
	bytecode = str.BytecodeFromString(sts.Map(func(r rune) rune {
		if uni.IsSpace(r) {
			return -1
		}
		return r
	}, matches[1]))
	return bytecode, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatBytecode(bytecode abs.BytecodeLike) {
	v.AppendString("'")
	var s = bytecode.AsString()
	var length = len(s)
	if length > 0 {
		v.AppendString(s[0:4])
	}
	for index := 4; index < length; index += 4 {
		v.AppendString(" ")
		v.AppendString(s[index : index+4])
	}
	v.AppendString("'")
}

// This method attempts to parse a narrative string. It returns the narrative
// string and whether or not the narrative string was successfully parsed.
func (v *parser) parseNarrative() (abs.NarrativeLike, *Token, bool) {
	var token *Token
	var narrative abs.NarrativeLike
	token = v.nextToken()
	if token.Type != TokenNARRATIVE {
		v.backupOne(token)
		return narrative, token, false
	}
	narrative = str.NarrativeFromString(trimIndentation(token.Value))
	return narrative, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatNarrative(narrative abs.NarrativeLike) {
	var s = narrative.AsString()
	var lines = sts.Split(s, EOL)
	v.AppendString(`">`)
	v.depth++
	for _, line := range lines {
		if len(line) > 0 {
			v.AppendNewline()
			v.AppendString(line)
		} else {
			v.AppendString(EOL)
		}
	}
	v.depth--
	v.AppendNewline()
	v.AppendString(`<"`)
}

// This method attempts to parse a quote string. It returns the quote string
// and whether or not the quote string was successfully parsed.
func (v *parser) parseQuote() (abs.QuoteLike, *Token, bool) {
	var token *Token
	var quote abs.QuoteLike
	token = v.nextToken()
	if token.Type != TokenQUOTE {
		v.backupOne(token)
		return quote, token, false
	}
	var matches = bytesToStrings(quoteScanner.FindSubmatch([]byte(token.Value)))
	// We must unquote the full token string properly.
	var unquoted, _ = stc.Unquote(matches[0])
	quote = str.QuoteFromString(unquoted)
	return quote, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatQuote(quote abs.QuoteLike) {
	// We must requote the string string properly.
	var s = stc.Quote(quote.AsString())
	v.AppendString(s)
}

// This method attempts to parse a string sequence. It returns the
// string sequence and whether or not the string sequence was
// successfully parsed.
func (v *parser) parseString() (abs.String, *Token, bool) {
	var ok bool
	var token *Token
	var s abs.String
	s, token, ok = v.parseBinary()
	if !ok {
		s, token, ok = v.parseBytecode()
	}
	if !ok {
		s, token, ok = v.parseName()
	}
	if !ok {
		s, token, ok = v.parseNarrative()
	}
	if !ok {
		s, token, ok = v.parseQuote()
	}
	if !ok {
		s, token, ok = v.parseSymbol()
	}
	if !ok {
		s, token, ok = v.parseTag()
	}
	if !ok {
		s, token, ok = v.parseVersion()
	}
	if !ok {
		// Override any empty strings returned from failed parsing attempts.
		s = nil
	}
	return s, token, ok
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (abs.SymbolLike, *Token, bool) {
	var token *Token
	var symbol abs.SymbolLike
	token = v.nextToken()
	if token.Type != TokenSYMBOL {
		v.backupOne(token)
		return symbol, token, false
	}
	var matches = bytesToStrings(symbolScanner.FindSubmatch([]byte(token.Value)))
	symbol = str.SymbolFromString(matches[1]) // Remove the leading '$'.
	return symbol, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatSymbol(symbol abs.SymbolLike) {
	var identifier = symbol.AsString()
	var s = "$" + identifier
	v.AppendString(s)
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (abs.TagLike, *Token, bool) {
	var token *Token
	var tag abs.TagLike
	token = v.nextToken()
	if token.Type != TokenTAG {
		v.backupOne(token)
		return tag, token, false
	}
	var matches = bytesToStrings(tagScanner.FindSubmatch([]byte(token.Value)))
	tag = str.TagFromString(matches[1]) // Remove the leading "#".
	return tag, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatTag(tag abs.TagLike) {
	var encoded = tag.AsString()
	var s = "#" + encoded
	v.AppendString(s)
}

// This method attempts to parse a version string. It returns the version
// string and whether or not the version string was successfully parsed.
func (v *parser) parseVersion() (abs.VersionLike, *Token, bool) {
	var token *Token
	var version abs.VersionLike
	token = v.nextToken()
	if token.Type != TokenVERSION {
		v.backupOne(token)
		return version, token, false
	}
	var matches = bytesToStrings(versionScanner.FindSubmatch([]byte(token.Value)))
	version = str.VersionFromString(matches[1]) // Remove the leading "v".
	return version, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatVersion(version abs.VersionLike) {
	var s = "v" + version.AsString()
	v.AppendString(s)
}
