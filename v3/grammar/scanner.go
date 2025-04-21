/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	fra "github.com/craterdog/go-collection-framework/v5"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS INTERFACE

// Access Function

func ScannerClass() ScannerClassLike {
	return scannerClass()
}

// Constructor Methods

func (c *scannerClass_) Scanner(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	if uti.IsUndefined(tokens) {
		panic("The \"tokens\" attribute is required by this class.")
	}
	var instance = &scanner_{
		// Initialize the instance attributes.
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go instance.scanTokens() // Start scanning tokens in the background.
	return instance
}

// Function Methods

func (c *scannerClass_) FormatToken(
	token TokenLike,
) string {
	var value = token.GetValue()
	value = fmt.Sprintf("%q", value)
	if len(value) > 40 {
		value = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_.GetValue(token.GetType()),
		token.GetLine(),
		token.GetPosition(),
		value,
	)
}

func (c *scannerClass_) FormatType(
	tokenType TokenType,
) string {
	return c.tokens_.GetValue(tokenType)
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var matcher = c.matchers_.GetValue(tokenType)
	var match = matcher.FindString(tokenValue)
	return uti.IsDefined(match)
}

// INSTANCE INTERFACE

// Principal Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerClass()
}

// PROTECTED INTERFACE

// Private Methods

func (v *scanner_) emitToken(
	tokenType TokenType,
) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = TokenClass().Token(v.line_, v.position_, tokenType, value)
	//fmt.Println(ScannerClass().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(
	tokenType TokenType,
) bool {
	// Attempt to match the specified token type.
	var class = scannerClass()
	var matcher = class.matchers_.GetValue(tokenType)
	var text = string(v.runes_[v.next_:])
	var match = matcher.FindString(text)
	if uti.IsUndefined(match) {
		return false
	}

	// Check for false delimiter matches.
	var token = []rune(match)
	var length = uint(len(token))
	var previous = token[length-1]
	if tokenType == DelimiterToken && uint(len(v.runes_)) > v.next_+length {
		var next = v.runes_[v.next_+length]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += length
	v.emitToken(tokenType)
	var count = uint(sts.Count(match, "\n"))
	if count > 0 {
		v.line_ += count
		v.position_ = v.indexOfLastEol(token)
	} else {
		v.position_ += v.next_ - v.first_
	}
	v.first_ = v.next_
	return true
}

func (v *scanner_) indexOfLastEol(
	runes []rune,
) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		// Find the next token type.
		case v.foundToken(DelimiterToken):
		case v.foundToken(NewlineToken):
		case v.foundToken(SpaceToken):
		case v.foundToken(AngleToken):
		case v.foundToken(BinaryToken):
		case v.foundToken(BooleanToken):
		case v.foundToken(BytecodeToken):
		case v.foundToken(CitationToken):
		case v.foundToken(CommentToken):
		case v.foundToken(DurationToken):
		case v.foundToken(MomentToken):
		case v.foundToken(NameToken):
		case v.foundToken(NarrativeToken):
		case v.foundToken(NoteToken):
		case v.foundToken(PatternToken):
		case v.foundToken(PercentageToken):
		case v.foundToken(ProbabilityToken):
		case v.foundToken(QuoteToken):
		case v.foundToken(NumberToken):
		case v.foundToken(ResourceToken):
		case v.foundToken(SymbolToken):
		case v.foundToken(TagToken):
		case v.foundToken(VersionToken):
		case v.foundToken(IdentifierToken):
		case v.foundToken(DotToken):
		case v.foundToken(ArrowToken):
		case v.foundToken(PlusToken):
		case v.foundToken(DashToken):
		case v.foundToken(StarToken):
		case v.foundToken(DoubleToken):
		case v.foundToken(SlashToken):
		case v.foundToken(CaretToken):
		case v.foundToken(AmpersandToken):
		case v.foundToken(LessToken):
		case v.foundToken(EqualToken):
		case v.foundToken(MoreToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseChannel()
}

// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   col.QueueLike[TokenLike]
}

// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   col.CatalogLike[TokenType, string]
	matchers_ col.CatalogLike[TokenType, *reg.Regexp]
}

// Class Reference

func scannerClass() *scannerClass_ {
	return scannerClassReference_
}

var scannerClassReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: fra.CatalogFromMap[TokenType, string](
		map[TokenType]string{
			// Define identifiers for each type of token.
			ErrorToken:       "error",
			AmpersandToken:   "ampersand",
			AngleToken:       "angle",
			ArrowToken:       "arrow",
			BinaryToken:      "binary",
			BooleanToken:     "boolean",
			BytecodeToken:    "bytecode",
			CaretToken:       "caret",
			CitationToken:    "citation",
			CommentToken:     "comment",
			DashToken:        "dash",
			DelimiterToken:   "delimiter",
			DotToken:         "dot",
			DoubleToken:      "double",
			DurationToken:    "duration",
			EqualToken:       "equal",
			IdentifierToken:  "identifier",
			LessToken:        "less",
			MomentToken:      "moment",
			MoreToken:        "more",
			NameToken:        "name",
			NarrativeToken:   "narrative",
			NewlineToken:     "newline",
			NoteToken:        "note",
			NumberToken:      "number",
			PatternToken:     "pattern",
			PercentageToken:  "percentage",
			PlusToken:        "plus",
			ProbabilityToken: "probability",
			QuoteToken:       "quote",
			ResourceToken:    "resource",
			SlashToken:       "slash",
			SpaceToken:       "space",
			StarToken:        "star",
			SymbolToken:      "symbol",
			TagToken:         "tag",
			VersionToken:     "version",
		},
	),
	matchers_: fra.CatalogFromMap[TokenType, *reg.Regexp](
		map[TokenType]*reg.Regexp{
			// Define pattern matchers for each type of token.
			AmpersandToken:   reg.MustCompile("^" + ampersand_),
			AngleToken:       reg.MustCompile("^" + angle_),
			ArrowToken:       reg.MustCompile("^" + arrow_),
			BinaryToken:      reg.MustCompile("^" + binary_),
			BooleanToken:     reg.MustCompile("^" + boolean_),
			BytecodeToken:    reg.MustCompile("^" + bytecode_),
			CaretToken:       reg.MustCompile("^" + caret_),
			CitationToken:    reg.MustCompile("^" + citation_),
			CommentToken:     reg.MustCompile("^" + comment_),
			DashToken:        reg.MustCompile("^" + dash_),
			DelimiterToken:   reg.MustCompile("^" + delimiter_),
			DotToken:         reg.MustCompile("^" + dot_),
			DoubleToken:      reg.MustCompile("^" + double_),
			DurationToken:    reg.MustCompile("^" + duration_),
			EqualToken:       reg.MustCompile("^" + equal_),
			IdentifierToken:  reg.MustCompile("^" + identifier_),
			LessToken:        reg.MustCompile("^" + less_),
			MomentToken:      reg.MustCompile("^" + moment_),
			MoreToken:        reg.MustCompile("^" + more_),
			NameToken:        reg.MustCompile("^" + name_),
			NarrativeToken:   reg.MustCompile("^" + narrative_),
			NewlineToken:     reg.MustCompile("^" + newline_),
			NoteToken:        reg.MustCompile("^" + note_),
			NumberToken:      reg.MustCompile("^" + number_),
			PatternToken:     reg.MustCompile("^" + pattern_),
			PercentageToken:  reg.MustCompile("^" + percentage_),
			PlusToken:        reg.MustCompile("^" + plus_),
			ProbabilityToken: reg.MustCompile("^" + probability_),
			QuoteToken:       reg.MustCompile("^" + quote_),
			ResourceToken:    reg.MustCompile("^" + resource_),
			SlashToken:       reg.MustCompile("^" + slash_),
			SpaceToken:       reg.MustCompile("^" + space_),
			StarToken:        reg.MustCompile("^" + star_),
			SymbolToken:      reg.MustCompile("^" + symbol_),
			TagToken:         reg.MustCompile("^" + tag_),
			VersionToken:     reg.MustCompile("^" + version_),
		},
	),
}

// Private Constants

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be used in
this way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expression patterns for each token type.
	delimiter_    = "(?:xor|with|while|to|throw|select|save|san|return|retrieve|reject|publish|post|on|notarize|not|matching|matches|loop|level|let|is|ior|in|if|from|each|do|discard|continue|checkout|break|at|as|and|accept|\\}|\\||\\{|\\]|\\[|\\?=|\\.\\.|\\+=|\\*=|\\)|\\(|@|;|:=|:|/=|-=|,)"
	newline_      = "(?:" + eol_ + ")"
	space_        = "(?:[ \\t]+)"
	alpha_        = "(?:[A-Za-z])"
	alphanumeric_ = "(?:(?:" + alpha_ + ")|(?:" + base10_ + "))"
	angle_        = "(?:~(0|(?:" + amplitude_ + ")|pi|π|tau|τ))"
	authority_    = "(?:[^/" + control_ + "]+)"
	base10_       = "(?:[0-9])"
	base16_       = "(?:(?:" + base10_ + ")|[a-f])"
	base32_       = "(?:(?:" + base10_ + ")|[A-DF-HJ-NP-TV-Z])"
	base64_       = "(?:(?:" + alphanumeric_ + ")|[\\+/])"
	binary_       = "(?:'>(" + eol_ + "(?:" + block_ + ")+)? *<')"
	block_        = "(?: *(?:" + base64_ + "){1,60}" + eol_ + ")"
	boolean_      = "(?:false|true)"
	bytecode_     = "(?:'((?:" + instruction_ + ")( (?:" + instruction_ + "))*)*')"
	character_    = "(?:(?:" + escape_ + ")|[^\"" + control_ + "])"
	citation_     = "(?:(?:" + name_ + ")/(?:" + version_ + "))"
	comment_      = "(?:!>" + eol_ + "(" + any_ + "|" + eol_ + ")*?" + eol_ + " *<!)"
	day_          = "(?:([0-2][1-9])|(3[0-1]))"
	days_         = "(?:(?:" + timespan_ + ")D)"
	duration_     = "(?:~(?:" + sign_ + ")?P((?:" + weeks_ + ")|((?:" + years_ + ")?(?:" + months_ + ")?(?:" + days_ + ")?(T(?:" + hours_ + ")?(?:" + minutes_ + ")?(?:" + seconds_ + ")?)?)))"
	escape_       = "(?:\\\\((?:" + unicode_ + ")|[abfnrtv\"\\\\]))"
	exponent_     = "(?:E(?:" + sign_ + ")?(?:" + ordinal_ + "))"
	float_        = "(?:(?:" + sign_ + ")?(?:" + amplitude_ + "))"
	fraction_     = "(?:\\.(?:" + base10_ + ")+)"
	fragment_     = "(?:[^>" + control_ + "]*)"
	hour_         = "(?:([0-1][0-9])|(2[0-3]))"
	hours_        = "(?:(?:" + timespan_ + ")H)"
	imaginary_    = "(?:(?:" + sign_ + ")?(?:" + amplitude_ + ")i)"
	infinity_     = "(?:(?:" + sign_ + ")?(infinity|∞))"
	instruction_  = "(?:(?:" + base16_ + "){4})"
	letter_       = "(?:" + lower_ + "|" + upper_ + ")"
	amplitude_    = "(?:(0(?:" + fraction_ + ")|(?:" + ordinal_ + ")(?:" + fraction_ + ")?)(?:" + exponent_ + ")?)"
	minute_       = "(?:[0-5][0-9])"
	minutes_      = "(?:(?:" + timespan_ + ")M)"
	moment_       = "(?:<(?:" + sign_ + ")?(?:" + year_ + ")(-(?:" + month_ + ")(-(?:" + day_ + ")(T(?:" + hour_ + ")(:(?:" + minute_ + ")(:(?:" + second_ + ")(?:" + fraction_ + ")?)?)?)?)?)?>)"
	month_        = "(?:(0[1-9])|(1[0-2]))"
	months_       = "(?:(?:" + timespan_ + ")M)"
	name_         = "(?:(/(?:" + identifier_ + "))+)"
	narrative_    = "(?:\">" + eol_ + "(" + any_ + "|" + eol_ + ")*?" + eol_ + " *<\")"
	note_         = "(?:! [^" + control_ + "]*)"
	ordinal_      = "(?:[1-9](?:" + base10_ + ")*)"
	path_         = "(?:[^\\?#>" + control_ + "]*)"
	pattern_      = "(?:none|(?:" + regex_ + ")|any)"
	percentage_   = "(?:(?:" + real_ + ")%)"
	probability_  = "(?:(?:" + fraction_ + ")|one)"
	polar_        = "(?:(?:" + amplitude_ + ")e\\^(?:" + angle_ + ")?i)"
	query_        = "(?:[^#>" + control_ + "]*)"
	quote_        = "(?:\"(?:" + character_ + ")*\")"
	real_         = "(?:0|(?:" + float_ + ")|(?:" + infinity_ + ")|(?:" + undefined_ + "))"
	rectangular_  = "(?:(?:" + float_ + ")(?:" + sign_ + ")(?:" + float_ + ")?i)"
	number_       = "(?:(?:" + polar_ + ")|(?:" + rectangular_ + ")|(?:" + imaginary_ + ")|(?:" + real_ + "))"
	regex_        = "(?:\"(?:" + character_ + ")+\"\\?)"
	resource_     = "(?:<(?:" + scheme_ + "):(//(?:" + authority_ + "))?/(?:" + path_ + ")(\\?(?:" + query_ + "))?(#(?:" + fragment_ + "))?>)"
	scheme_       = "(?:(?:" + alpha_ + ")((?:" + alphanumeric_ + ")|\\+|-|\\.)*)"
	second_       = "(?:([0-5][0-9])|(6[0-1]))"
	seconds_      = "(?:(?:" + timespan_ + ")S)"
	sign_         = "(?:\\+|-)"
	symbol_       = "(?:\\$(?:" + identifier_ + "))"
	tag_          = "(?:#(?:" + base32_ + ")+)"
	timespan_     = "(?:0|((?:" + ordinal_ + ")(?:" + fraction_ + ")?))"
	undefined_    = "(?:undefined)"
	unicode_      = "(?:(u(?:" + base16_ + "){4})|(U(?:" + base16_ + "){8}))"
	version_      = "(?:v(?:" + ordinal_ + ")(\\.(?:" + ordinal_ + "))*)"
	identifier_   = "(?:(?:" + letter_ + ")((?:" + letter_ + ")|" + digit_ + ")*)"
	weeks_        = "(?:(?:" + timespan_ + ")W)"
	year_         = "(?:0|(?:" + ordinal_ + "))"
	years_        = "(?:(?:" + timespan_ + ")Y)"
	dot_          = "(?:\\.)"
	arrow_        = "(?:<-)"
	plus_         = "(?:\\+)"
	dash_         = "(?:-)"
	star_         = "(?:\\*)"
	double_       = "(?://)"
	slash_        = "(?:/)"
	caret_        = "(?:\\^)"
	ampersand_    = "(?:&)"
	less_         = "(?:<)"
	equal_        = "(?:=)"
	more_         = "(?:>)"
)
