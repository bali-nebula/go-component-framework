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
		case v.foundToken(AmpersandToken):
		case v.foundToken(AndToken):
		case v.foundToken(AngleToken):
		case v.foundToken(BarToken):
		case v.foundToken(BinaryToken):
		case v.foundToken(BooleanToken):
		case v.foundToken(BytecodeToken):
		case v.foundToken(CaretToken):
		case v.foundToken(ColonEqualToken):
		case v.foundToken(ColonToken):
		case v.foundToken(CommentToken):
		case v.foundToken(DashEqualToken):
		case v.foundToken(DefaultEqualToken):
		case v.foundToken(DurationToken):
		case v.foundToken(EqualToken):
		case v.foundToken(IorToken):
		case v.foundToken(IsToken):
		case v.foundToken(MatchesToken):
		case v.foundToken(MomentToken):
		case v.foundToken(MoreToken):
		case v.foundToken(NameToken):
		case v.foundToken(NarrativeToken):
		case v.foundToken(NotToken):
		case v.foundToken(NoteToken):
		case v.foundToken(PatternToken):
		case v.foundToken(PlusEqualToken):
		case v.foundToken(PlusToken):
		case v.foundToken(QuoteToken):
		case v.foundToken(ResourceToken):
		case v.foundToken(SanToken):
		case v.foundToken(SlashEqualToken):
		case v.foundToken(SlashSlashToken):
		case v.foundToken(SlashToken):
		case v.foundToken(SnailToken):
		case v.foundToken(StarEqualToken):
		case v.foundToken(StarToken):
		case v.foundToken(SymbolToken):
		case v.foundToken(TagToken):
		case v.foundToken(VersionToken):
		case v.foundToken(XorToken):
		case v.foundToken(PercentageToken):
		case v.foundToken(ProbabilityToken):
		case v.foundToken(NumberToken):
		case v.foundToken(IdentifierToken):
		case v.foundToken(DotToken):
		case v.foundToken(DashToken):
		case v.foundToken(ArrowToken):
		case v.foundToken(LessToken):
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
			ErrorToken:        "error",
			AmpersandToken:    "ampersand",
			AndToken:          "and",
			AngleToken:        "angle",
			ArrowToken:        "arrow",
			BarToken:          "bar",
			BinaryToken:       "binary",
			BooleanToken:      "boolean",
			BytecodeToken:     "bytecode",
			CaretToken:        "caret",
			ColonToken:        "colon",
			ColonEqualToken:   "colonEqual",
			CommentToken:      "comment",
			DashToken:         "dash",
			DashEqualToken:    "dashEqual",
			DefaultEqualToken: "defaultEqual",
			DelimiterToken:    "delimiter",
			DotToken:          "dot",
			DurationToken:     "duration",
			EqualToken:        "equal",
			IdentifierToken:   "identifier",
			IorToken:          "ior",
			IsToken:           "is",
			LessToken:         "less",
			MatchesToken:      "matches",
			MomentToken:       "moment",
			MoreToken:         "more",
			NameToken:         "name",
			NarrativeToken:    "narrative",
			NewlineToken:      "newline",
			NotToken:          "not",
			NoteToken:         "note",
			NumberToken:       "number",
			PatternToken:      "pattern",
			PercentageToken:   "percentage",
			PlusToken:         "plus",
			PlusEqualToken:    "plusEqual",
			ProbabilityToken:  "probability",
			QuoteToken:        "quote",
			ResourceToken:     "resource",
			SanToken:          "san",
			SlashToken:        "slash",
			SlashEqualToken:   "slashEqual",
			SlashSlashToken:   "slashSlash",
			SnailToken:        "snail",
			SpaceToken:        "space",
			StarToken:         "star",
			StarEqualToken:    "starEqual",
			SymbolToken:       "symbol",
			TagToken:          "tag",
			VersionToken:      "version",
			XorToken:          "xor",
		},
	),
	matchers_: fra.CatalogFromMap[TokenType, *reg.Regexp](
		map[TokenType]*reg.Regexp{
			// Define pattern matchers for each type of token.
			AmpersandToken:    reg.MustCompile("^" + ampersand_),
			AndToken:          reg.MustCompile("^" + and_),
			AngleToken:        reg.MustCompile("^" + angle_),
			ArrowToken:        reg.MustCompile("^" + arrow_),
			BarToken:          reg.MustCompile("^" + bar_),
			BinaryToken:       reg.MustCompile("^" + binary_),
			BooleanToken:      reg.MustCompile("^" + boolean_),
			BytecodeToken:     reg.MustCompile("^" + bytecode_),
			CaretToken:        reg.MustCompile("^" + caret_),
			ColonToken:        reg.MustCompile("^" + colon_),
			ColonEqualToken:   reg.MustCompile("^" + colonEqual_),
			CommentToken:      reg.MustCompile("^" + comment_),
			DashToken:         reg.MustCompile("^" + dash_),
			DashEqualToken:    reg.MustCompile("^" + dashEqual_),
			DefaultEqualToken: reg.MustCompile("^" + defaultEqual_),
			DelimiterToken:    reg.MustCompile("^" + delimiter_),
			DotToken:          reg.MustCompile("^" + dot_),
			DurationToken:     reg.MustCompile("^" + duration_),
			EqualToken:        reg.MustCompile("^" + equal_),
			IdentifierToken:   reg.MustCompile("^" + identifier_),
			IorToken:          reg.MustCompile("^" + ior_),
			IsToken:           reg.MustCompile("^" + is_),
			LessToken:         reg.MustCompile("^" + less_),
			MatchesToken:      reg.MustCompile("^" + matches_),
			MomentToken:       reg.MustCompile("^" + moment_),
			MoreToken:         reg.MustCompile("^" + more_),
			NameToken:         reg.MustCompile("^" + name_),
			NarrativeToken:    reg.MustCompile("^" + narrative_),
			NewlineToken:      reg.MustCompile("^" + newline_),
			NotToken:          reg.MustCompile("^" + not_),
			NoteToken:         reg.MustCompile("^" + note_),
			NumberToken:       reg.MustCompile("^" + number_),
			PatternToken:      reg.MustCompile("^" + pattern_),
			PercentageToken:   reg.MustCompile("^" + percentage_),
			PlusToken:         reg.MustCompile("^" + plus_),
			PlusEqualToken:    reg.MustCompile("^" + plusEqual_),
			ProbabilityToken:  reg.MustCompile("^" + probability_),
			QuoteToken:        reg.MustCompile("^" + quote_),
			ResourceToken:     reg.MustCompile("^" + resource_),
			SanToken:          reg.MustCompile("^" + san_),
			SlashToken:        reg.MustCompile("^" + slash_),
			SlashEqualToken:   reg.MustCompile("^" + slashEqual_),
			SlashSlashToken:   reg.MustCompile("^" + slashSlash_),
			SnailToken:        reg.MustCompile("^" + snail_),
			SpaceToken:        reg.MustCompile("^" + space_),
			StarToken:         reg.MustCompile("^" + star_),
			StarEqualToken:    reg.MustCompile("^" + starEqual_),
			SymbolToken:       reg.MustCompile("^" + symbol_),
			TagToken:          reg.MustCompile("^" + tag_),
			VersionToken:      reg.MustCompile("^" + version_),
			XorToken:          reg.MustCompile("^" + xor_),
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
	delimiter_    = "(?:with|while|to|throw|select|save|return|retrieve|reject|publish|post|on|notarize|matching|loop|level|let|in|if|from|each|do|discard|continue|checkout|break|at|as|accept|\\}|\\{|\\]|\\[|\\.\\.|\\)|\\(|;|,)"
	newline_      = "(?:" + eol_ + ")"
	space_        = "(?: )"
	alpha_        = "(?:[A-Za-z])"
	alphanumeric_ = "(?:(?:" + alpha_ + ")|(?:" + base10_ + "))"
	ampersand_    = "(?:&)"
	and_          = "(?:AND)"
	angle_        = "(?:~((?:" + zero_ + ")|(?:" + magnitude_ + ")))"
	authority_    = "(?:[^/" + control_ + "]+)"
	bar_          = "(?:\\|)"
	base10_       = "(?:[0-9])"
	base16_       = "(?:(?:" + base10_ + ")|[a-f])"
	base32_       = "(?:(?:" + base10_ + ")|[A-DF-HJ-NP-TV-Z])"
	base64_       = "(?:(?:" + alphanumeric_ + ")|[\\+/])"
	binary_       = "(?:'>" + eol_ + "(?:" + space_ + ")*(?:" + base64_ + ")+" + eol_ + "*(?:" + space_ + ")*<')"
	boolean_      = "(?:false|true)"
	bytecode_     = "(?:'((?:" + instruction_ + ")((?:" + space_ + ")(?:" + instruction_ + "))*)*')"
	caret_        = "(?:\\^)"
	character_    = "(?:(?:" + escape_ + ")|[^\"" + control_ + "])"
	colonEqual_   = "(?::=)"
	colon_        = "(?::)"
	comment_      = "(?:!>" + any_ + "*?<!)"
	dashEqual_    = "(?:-=)"
	day_          = "(?:([0-2][1-9])|(3[0-1]))"
	days_         = "(?:(?:" + timespan_ + ")D)"
	defaultEqual_ = "(?:\\?=)"
	duration_     = "(?:~(?:" + sign_ + ")?P((?:" + weeks_ + ")|((?:" + years_ + ")?(?:" + months_ + ")?(?:" + days_ + ")?(T(?:" + hours_ + ")?(?:" + minutes_ + ")?(?:" + seconds_ + ")?)?)))"
	equal_        = "(?:=)"
	escape_       = "(?:\\\\((?:" + unicode_ + ")|[abfnrtv\"\\\\]))"
	exponent_     = "(?:E(?:" + sign_ + ")?(?:" + ordinal_ + "))"
	float_        = "(?:(?:" + sign_ + ")?(?:" + magnitude_ + "))"
	fraction_     = "(?:(?:" + dot_ + ")(?:" + base10_ + ")+)"
	fragment_     = "(?:[^>" + control_ + "]*)"
	hour_         = "(?:([0-1][0-9])|(2[0-3]))"
	hours_        = "(?:(?:" + timespan_ + ")H)"
	imaginary_    = "(?:(?:" + sign_ + ")?(?:" + magnitude_ + ")?i)"
	infinity_     = "(?:(?:" + sign_ + ")?(infinity|∞))"
	instruction_  = "(?:(?:" + base16_ + "){4})"
	ior_          = "(?:IOR)"
	is_           = "(?:IS)"
	letter_       = "(?:" + lower_ + "|" + upper_ + ")"
	magnitude_    = "(?:e|pi|π|phi|φ|tau|τ|(?:" + scalar_ + "))"
	matches_      = "(?:MATCHES)"
	minute_       = "(?:[0-5][0-9])"
	minutes_      = "(?:(?:" + timespan_ + ")M)"
	moment_       = "(?:<(?:" + sign_ + ")?(?:" + year_ + ")(-(?:" + month_ + ")(-(?:" + day_ + ")(T(?:" + hour_ + ")(:(?:" + minute_ + ")(:(?:" + second_ + ")(?:" + fraction_ + ")?)?)?)?)?)?>)"
	month_        = "(?:(0[1-9])|(1[0-2]))"
	months_       = "(?:(?:" + timespan_ + ")M)"
	more_         = "(?:>)"
	name_         = "(?:((?:" + slash_ + ")(?:" + identifier_ + "))+)"
	narrative_    = "(?:\">" + eol_ + "" + any_ + "*?" + eol_ + "(?:" + space_ + ")*<\")"
	not_          = "(?:NOT)"
	note_         = "(?:! [^" + control_ + "]*)"
	one_          = "(?:1\\.)"
	ordinal_      = "(?:[1-9](?:" + base10_ + ")*)"
	path_         = "(?:[^\\?#>" + control_ + "]*)"
	pattern_      = "(?:none|(?:" + regex_ + ")|any)"
	plusEqual_    = "(?:\\+=)"
	plus_         = "(?:\\+)"
	polar_        = "(?:(?:" + magnitude_ + ")e\\^(?:" + angle_ + ")?i)"
	query_        = "(?:[^#>" + control_ + "]*)"
	quote_        = "(?:\"(?:" + character_ + ")*\")"
	real_         = "(?:(?:" + zero_ + ")|(?:" + float_ + ")|(?:" + infinity_ + ")|(?:" + undefined_ + "))"
	rectangular_  = "(?:(?:" + float_ + ")(?:" + sign_ + ")(?:" + float_ + ")?i)"
	regex_        = "(?:\"(?:" + character_ + ")+\"\\?)"
	resource_     = "(?:<(?:" + scheme_ + "):(//(?:" + authority_ + "))?/(?:" + path_ + ")(\\?(?:" + query_ + "))?(#(?:" + fragment_ + "))?>)"
	san_          = "(?:SAN)"
	scalar_       = "(?:((?:" + zero_ + ")(?:" + fraction_ + ")|(?:" + ordinal_ + ")(?:" + fraction_ + ")?)(?:" + exponent_ + ")?)"
	scheme_       = "(?:(?:" + alpha_ + ")((?:" + alphanumeric_ + ")|(?:" + plus_ + ")|(?:" + dash_ + ")|(?:" + dot_ + "))*)"
	second_       = "(?:([0-5][0-9])|(6[0-1]))"
	seconds_      = "(?:(?:" + timespan_ + ")S)"
	sign_         = "(?:(?:" + plus_ + ")|(?:" + dash_ + "))"
	slashEqual_   = "(?:/=)"
	slashSlash_   = "(?://)"
	slash_        = "(?:/)"
	snail_        = "(?:@)"
	starEqual_    = "(?:\\*=)"
	star_         = "(?:\\*)"
	symbol_       = "(?:\\$(?:" + identifier_ + "))"
	tag_          = "(?:#(?:" + base32_ + ")+)"
	timespan_     = "(?:(?:" + zero_ + ")|((?:" + ordinal_ + ")(?:" + fraction_ + ")?))"
	undefined_    = "(?:undefined)"
	unicode_      = "(?:(u(?:" + base16_ + "){4})|(U(?:" + base16_ + "){8}))"
	version_      = "(?:v(?:" + ordinal_ + ")((?:" + dot_ + ")(?:" + ordinal_ + "))*)"
	weeks_        = "(?:(?:" + timespan_ + ")W)"
	xor_          = "(?:XOR)"
	year_         = "(?:(?:" + zero_ + ")|(?:" + ordinal_ + "))"
	years_        = "(?:(?:" + timespan_ + ")Y)"
	zero_         = "(?:0)"
	percentage_   = "(?:(?:" + real_ + ")%)"
	probability_  = "(?:(?:" + fraction_ + ")|(?:" + one_ + "))"
	number_       = "(?:(?:" + polar_ + ")|(?:" + rectangular_ + ")|(?:" + imaginary_ + ")|(?:" + real_ + "))"
	identifier_   = "(?:(?:" + letter_ + ")((?:" + letter_ + ")|" + digit_ + ")*)"
	dot_          = "(?:\\.)"
	dash_         = "(?:-)"
	arrow_        = "(?:<-)"
	less_         = "(?:<)"
)
