/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	"bytes"
	"regexp"
)

// LEXICAL INTERFACES

// This interface defines the methods supported by all lexical elements.
type Lexical interface {
	AsString() string
}

// ANGLE ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// an angle element. It is useful to include in parsing error messages.
const AngleSyntax = `
	E        = "e"
	PI       = "pi" | "π"
	PHI      = "phi" | "φ"
	TAU      = "tau" | "τ"
	CONSTANT = E | PI | PHI | TAU
	SIGN     = [+-]
	ZERO     = "0"
	ORDINAL  = [1-9][0-9]*
	FRACTION = .[0-9]+
	EXPONENT = "E" SIGN? ORDINAL
	SCALAR   = (ORDINAL FRACTION? | ZERO FRACTION) EXPONENT?
	ANGLE    = "~" (CONSTANT | SCALAR | ZERO)
`

// These constants are used to form a regular expression for valid angle
// entities.
const (
	e        = `e`
	pi       = `pi|π`
	phi      = `phi|φ`
	tau      = `tau|τ`
	constant = e + `|` + pi + `|` + phi + `|` + tau
	sign     = `[+-]`
	zero     = `0`
	ordinal  = `[1-9][0-9]*`
	fraction = `\.[0-9]+`
	exponent = `E` + sign + `?` + ordinal
	scalar   = `(?:(?:` + ordinal + `(?:` + fraction + `)?)|(?:` + zero + fraction + `))(?:` + exponent + `)?`
	angle    = `~(` + constant + `|` + scalar + `|` + zero + `)`
)

// This scanner is used for matching angle entities.
var angleScanner = regexp.MustCompile(`^` + angle)

// This function returns for the specified string an array of the matching
// subgroups for an angle element. The first item in the array is the entire
// matched string.
func ScanAngle(v []byte) []string {
	return bytesToStrings(angleScanner.FindSubmatch(v))
}

// BINARY STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a binary string. It is useful to include in parsing error messages.
const BinarySyntax = `
	WHITESPACE = \s
	BASE64     = [A-Za-z0-9+/]
	BINARY     = "'" (BASE64 | WHITESPACE)* "'"
`

// These constants are used to form a regular expression for valid binary
// strings.
const (
	whitespace = `\s`
	base64     = `[A-Za-z0-9+/]`
	binary     = `'((?:` + base64 + `|` + whitespace + `)*)'`
)

// This scanner is used for matching binary strings.
var binaryScanner = regexp.MustCompile(`^` + binary)

// This function returns for the specified string an array of the matching
// subgroups for a binary string. The first item in the array is the entire
// matched string.
func ScanBinary(v []byte) []string {
	return bytesToStrings(binaryScanner.FindSubmatch(v))
}

// BOOLEAN ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a boolean element. It is useful to include in parsing error messages.
const BooleanSyntax = `
	FALSE   = "false"
	TRUE    = "true"
	BOOLEAN = FALSE | TRUE
`

// These constants are used to form a regular expression for valid boolean
// entities.
const (
	boolean = `(false|true)`
)

// This scanner is used for matching boolean entities.
var booleanScanner = regexp.MustCompile(`^` + boolean)

// This function returns for the specified string an array of the matching
// subgroups for a boolean element. The first item in the array is the entire
// matched string.
func ScanBoolean(v []byte) []string {
	return bytesToStrings(booleanScanner.FindSubmatch(v))
}

// COMMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a comment. It is useful to include in parsing error messages.
const CommentSyntax = `
	EOL       = "\n"
	TAB       = "\t"
	CHARACTER = .*
	COMMENT = "!*" EOL (COMMENT | CHARACTER)*? EOL TAB* "*!"
`

// This function returns for the specified string an array of the matching
// subgroups for a commment. The first item in the array is the entire matched
// string. Since a comment can be recursive a regular expression is not used in
// this implementation.
func ScanComment(v []byte) []string {
	var result []string
	if !bytes.HasPrefix(v, bangStar) {
		return result
	}
	var current = 2 // Skip the leading "!*" characters.
	var level = 1
	for level > 0 {
		s := v[current:]
		switch {
		case len(s) == 0:
			return result
		case bytes.HasPrefix(s, bangStar):
			current++ // Skip the "*" character.
			level++   // Start a nested comment.
		case bytes.HasPrefix(s, starBang):
			current++ // Skip the "!" character.
			level--   // Terminate the current comment.
		}
		current++ // Accept the next character.
	}
	result = append(result, string(v[:current])) // Includes bang delimeters.
	return result
}

// DELIMITER SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a delimiter. It is useful to include in parsing error messages.
const DelimiterSyntax = `
	DELIMITER = "&" | "(" | ... | "}"
`

// This function returns for the specified string an array of the matching
// subgroups for a delimiter. The first item in the array is the entire
// matched string.
func ScanDelimiter(v []byte) []string {
	var result []string
	for _, delimiter := range delimiters {
		if bytes.HasPrefix(v, delimiter) {
			result = append(result, string(delimiter))
		}
	}
	return result
}

// DURATION ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a duration element. It is useful to include in parsing error messages.
const DurationSyntax = `
	TSPAN    = ZERO | ORDINAL FRACTION?
	WEEKS    = TSPAN "W"
	DATES    = (TSPAN "Y")? (TSPAN "M")? (TSPAN "D")?
	TIMES    = "t" (TSPAN "H")? (TSPAN "M")? (TSPAN "S")?
	DURATION = "~" SIGN? "P" (WEEKS | DATES TIMES?)
`

// These constants are used to form a regular expression for valid duration
// entities.
const (
	tspan    = `(?:` + zero + `|` + ordinal + `(?:` + fraction + `)?)`
	weeks    = `(` + tspan + `W)`
	dates    = `(` + tspan + `Y)?(` + tspan + `M)?(` + tspan + `D)?`
	times    = `(T)(` + tspan + `H)?(` + tspan + `M)?(` + tspan + `S)?`
	duration = `~(` + sign + `?)P(?:` + weeks + `|` + dates + `(?:` + times + `)?)`
)

// This scanner is used for matching duration entities.
var durationScanner = regexp.MustCompile(`^` + duration)

// This function returns for the specified string an array of the matching
// subgroups for a duration element. The first item in the array is the entire
// matched string.
func ScanDuration(v []byte) []string {
	return bytesToStrings(durationScanner.FindSubmatch(v))
}

// IDENTIFIER SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// an identifier. It is useful to include in parsing error messages.
const IdentifierSyntax = `
	LETTER    = {unicode letter}
	DIGIT     = {unicode digit}
	IDENTIFIER = LETTER (LETTER | DIGIT)*
`

// These constants are used to form a regular expression for valid identifiers.
const (
	letter     = `\pL` // All unicode letters and connectors like underscores.
	digit      = `\pN` // All unicode digits.
	identifier = letter + `(?:` + letter + `|` + digit + `)*`
)

// This scanner is used for matching identifiers.
var identifierScanner = regexp.MustCompile(`^` + identifier)

// This function returns for the specified string an array of the matching
// subgroups for an identifier. The first item in the array is the entire
// matched string.
func ScanIdentifier(v []byte) []string {
	return bytesToStrings(identifierScanner.FindSubmatch(v))
}

// KEYWORD SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a keyword. It is useful to include in parsing error messages.
const KeywordSyntax = `
	KEYWORD = "accept" | "any" | ... | "with"
`

// This function returns for the specified string an array of the matching
// subgroups for a keyword. The first item in the array is the entire
// matched string.
func ScanKeyword(v []byte) []string {
	var result []string
	for _, keyword := range keywords {
		if bytes.HasPrefix(v, keyword) {
			result = append(result, string(keyword))
		}
	}
	return result
}

// MOMENT ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a moment element. It is useful to include in parsing error messages.
const MomentSyntax = `
	SIGN     = [+-]
	ORDINAL  = [1-9][0-9]*
	YEAR     = SIGN? ORDINAL
	MONTH    = [0][1-9] | [1][012]
	DAY      = [012][1-9] | [3][01]
	HOUR     = [01][0-9] | [2][0-3]
	MINUTE   = [0-5][0-9]
	SECOND   = [0-5][0-9] | [6][01]
	FRACTION = .[0-9]+
	MOMENT   = "<" YEAR ("-" MONTH ("-" DAY ("T" HOUR (":" MINUTE (":" SECOND FRACTION?)?)?)?)?)? ">"
`

// These constants are used to form a regular expression for valid moment
// entities.
const (
	year   = sign + `?` + ordinal
	month  = `(?:[0][1-9])|(?:[1][012])`
	day    = `(?:[012][1-9])|(?:[3][01])`
	hour   = `(?:[01][0-9])|(?:[2][0-3])`
	minute = `[0-5][0-9]`
	second = `(?:[0-5][0-9])|(?:[6][01])`
	moment = `<(` + year + `)(?:-(` + month + `)(?:-(` + day + `)` +
		`(?:T(` + hour + `)(?::(` + minute + `)(?::((?:` + second + `)(?:` + fraction + `)?))?)?)?)?)?>`
)

// This scanner is used for matching moment entities.
var momentScanner = regexp.MustCompile(`^` + moment)

// This function returns for the specified string an array of the matching
// subgroups for a moment element. The first item in the array is the entire
// matched string.
func ScanMoment(v []byte) []string {
	return bytesToStrings(momentScanner.FindSubmatch(v))
}

// MONIKER STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a moniker string. It is useful to include in parsing error messages.
const MonikerSyntax = `
	LETTER    = {unicode letter}
	DIGIT     = {unicode digit}
	SEPARATOR = [-+.]
	NAME      = LETTER (SEPARATOR? (LETTER | DIGIT))*
	MONIKER   = ("/" NAME)+
`

// These constants are used to form a regular expression for valid moniker strings.
const (
	separator = `[-+.]`
	name      = letter + `(?:` + separator + `?(?:` + letter + `|` + digit + `))*`
	moniker   = `((?:/` + name + `)+)` // Cannot capture each name...
)

// This scanner is used for matching moniker strings.
var monikerScanner = regexp.MustCompile(`^` + moniker)

// This function returns for the specified string an array of the matching
// subgroups for a moniker string. The first item in the array is the entire
// matched string.
func ScanMoniker(v []byte) []string {
	return bytesToStrings(monikerScanner.FindSubmatch(v))
}

// NARRATIVE STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a narrative string. It is useful to include in parsing error messages.
const NarrativeSyntax = `
	EOL       = "\n"
	TAB       = "\t"
	CHARACTER = .*
	NARRATIVE = '"' EOL (NARRATIVE | CHARACTER)* EOL TAB* '"'
`

// This function returns for the specified string an array of the matching
// subgroups for a narrative string. The first item in the array is the entire
// matched string. Since a narrative can be recursive a regular expression is
// not used in this implementation.
func ScanNarrative(v []byte) []string {
	var result []string
	if !bytes.HasPrefix(v, doubleQuoteEOL) {
		return result
	}
	var endPossible = false
	var current = 2 // Skip the double quote and EOL.
	var level = 1
	for level > 0 {
		s := v[current:]
		switch {
		case len(s) == 0:
			return result
		case !endPossible && bytes.HasPrefix(s, doubleQuoteEOL):
			// Start nested narrative.
			current++ // Skip the double quote.
			level++
		case bytes.HasPrefix(s, eol):
			// A closing double quote may be pending.
			endPossible = true
		case endPossible && bytes.HasPrefix(s, tab):
			// Ignore any tabs before the closing double quote.
		case endPossible && bytes.HasPrefix(s, doubleQuote):
			// Terminate this narrative.
			level--
			endPossible = false
		default:
			// A closing double quote did not appear.
			endPossible = false
		}
		current++ // Accept the next character.
	}
	result = append(result, string(v[:current])) // Includes double quote delimeters.
	return result
}

// NOTE SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a note. It is useful to include in parsing error messages.
const NoteSyntax = `
	EOL       = "\n"
	NOTE: "! " !EOL*
`

// These constants are used to form a regular expression for valid notes.
const (
	note = `(! [^\n]*)`
)

// This scanner is used for matching notes.
var noteScanner = regexp.MustCompile(`^` + note)

// This function returns for the specified string an array of the matching
// subgroups for a note. The first item in the array is the entire matched
// string.
func ScanNote(v []byte) []string {
	return bytesToStrings(noteScanner.FindSubmatch(v))
}

// NUMBER ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a number element. It is useful to include in parsing error messages.
const NumberSyntax = `
	E        = "e"
	PI       = "pi" | "π"
	PHI      = "phi" | "φ"
	TAU      = "tau" | "τ"
	CONSTANT = E | PI | PHI | TAU
	SIGN      = [+-]
	ZERO      = "0"
	ORDINAL   = [1-9][0-9]*
	FRACTION  = .[0-9]+
	EXPONENT  = "E" SIGN? ORDINAL
	SCALAR    = (ORDINAL FRACTION? | ZERO FRACTION) EXPONENT?
	SIGNED    = SIGN? (CONSTANT | SCALAR)
	INFINITY  = "infinity" | "∞"
	UNDEFINED = "undefined"
	NUMBER    = SIGNED "i"? | ZERO | INFINITY | UNDEFINED | "(" SIGNED ", " (SIGNED | ANGLE) "i" ")"
`

// These constants are used to form a regular expression for valid number
// entities.
const (
	signed    = sign + `?(?:` + constant + `|` + scalar + `)`
	infinity  = `infinity|∞`
	undefined = `undefined`
	number    = `(` +
		signed + `i?` + `|` +
		zero + `|` +
		infinity + `|` +
		undefined + `|` + `\(` +
		`(` + signed + `)` +
		`(?:, )` +
		`(` + signed + `|` + angle + `)i` +
		`\)` +
		`)`
)

// This scanner is used for matching number entities.
var numberScanner = regexp.MustCompile(`^` + number)

// This function returns for the specified string an array of the matching
// subgroups for a number element. The first item in the array is the entire
// matched string.
func ScanNumber(v []byte) []string {
	return bytesToStrings(numberScanner.FindSubmatch(v))
}

// PATTERN ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a pattern element. It is useful to include in parsing error messages.
const PatternSyntax = `
	BASE16  = [0-9a-f]
	UNICODE = "u" BASE16{4}
	ESCAPE  = "\" (UNICODE | [bfrnt\])
	RUNE    = ESCAPE | '\"' | !["\r\n]
	REGEX   = '"' RUNE+ '"?'
	NONE    = "none"
	ANY     = "any"
	PATTERN = NONE | REGEX | ANY
`

// These constants are used to form a regular expression for valid pattern
// entities.
const (
	base16  = `[0-9a-f]`
	unicode = `u` + base16 + `{4}`
	escape  = `\\(?:` + unicode + `|[bfrnt\\])`
	roon    = `(?:` + escape + `|\\"|[^"\r\n]` + `)`
	regex   = `"(` + roon + `+)"\?`
	pattern = `none` + `|` + regex + `|` + `any`
)

// This scanner is used for matching pattern entities.
var patternScanner = regexp.MustCompile(`^` + pattern)

// This function returns for the specified string an array of the matching
// subgroups for a pattern element. The first item in the array is the entire
// matched string.
func ScanPattern(v []byte) []string {
	return bytesToStrings(patternScanner.FindSubmatch(v))
}

// PERCENTAGE ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a percentage element. It is useful to include in parsing error messages.
const PercentageSyntax = `
	E        = "e"
	PI       = "pi" | "π"
	PHI      = "phi" | "φ"
	TAU      = "tau" | "τ"
	CONSTANT = E | PI | PHI | TAU
	SIGN       = [+-]
	ZERO       = "0"
	ORDINAL    = [1-9][0-9]*
	FRACTION   = .[0-9]+
	EXPONENT   = "E" SIGN? ORDINAL
	SCALAR     = (ORDINAL FRACTION? | ZERO FRACTION) EXPONENT?
	SIGNED     = SIGN? (CONSTANT | SCALAR)
	PERCENTAGE = (SIGNED | ZERO) "%"
`

// These constants are used to form a regular expression for valid percentage
// entities.
const (
	percentage = `(` + signed + `|` + zero + `)%`
)

// This scanner is used for matching percentage entities.
var percentageScanner = regexp.MustCompile(`^` + percentage)

// This function returns for the specified string an array of the matching
// subgroups for a percentage element. The first item in the array is the entire
// matched string.
func ScanPercentage(v []byte) []string {
	return bytesToStrings(percentageScanner.FindSubmatch(v))
}

// PROBABILITY ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a probability element. It is useful to include in parsing error messages.
const ProbabilitySyntax = `
	FRACTION    = .[0-9]+
	ONE         = "1."
	PROBABILITY = FRACTION | ONE
`

// These constants are used to form a regular expression for valid probability
// entities.
const (
	probability = `(` + fraction + `|1\.)`
)

// This scanner is used for matching probability entities.
var probabilityScanner = regexp.MustCompile(`^` + probability)

// This function returns for the specified string an array of the matching
// subgroups for a probability element. The first item in the array is the
// entire matched string.
func ScanProbability(v []byte) []string {
	return bytesToStrings(probabilityScanner.FindSubmatch(v))
}

// QUOTE STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a quoted string. It is useful to include in parsing error messages.
const QuoteSyntax = `
	BASE16  = [0-9a-f]
	UNICODE = "u" BASE16{4}
	ESCAPE  = "\" (UNICODE | [bfrnt\])
	RUNE    = ESCAPE | '\"' | !["\r\n]
	QUOTE   = '"' RUNE* '"'
`

// These constants are used to form a regular expression for valid quoted
// strings.
const (
	quote = `"(` + roon + `*)"`
)

// This scanner is used for matching quoted strings.
var quoteScanner = regexp.MustCompile(`^` + quote)

// This function returns for the specified string an array of the matching
// subgroups for a quoted string. The first item in the array is the entire
// matched string.
func ScanQuote(v []byte) []string {
	return bytesToStrings(quoteScanner.FindSubmatch(v))
}

// RESOURCE ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a resource element. It is useful to include in parsing error messages.
const ResourceSyntax = `
	SCHEME    = [a-zA-Z][0-9a-zA-Z+-.]*
	AUTHORITY = ![/]+
	PATH      = ![?#]*
	QUERY     = ![#>]*
	FRAGMENT  = ![>]+
	RESOURCE  = "<" SCHEME ":" ("//" AUTHORITY)? "/" PATH ("?" QUERY)? ("#" FRAGMENT)? ">"
`

// These constants are used to form a regular expression for valid resource
// entities.
const (
	scheme    = `[a-zA-Z][0-9a-zA-Z+-.]*`
	authority = `[^/]+`
	path      = `[^?#]*`
	query     = `[^#>]*`
	fragment  = `[^>]+`
	resource  = `<(` +
		`(` + scheme + `):` +
		`(?:` + `//(` + authority + `)` + `)?` +
		`(` + path + `)` +
		`(?:` + `\?(` + query + `)` + `)?` +
		`(?:` + `#(` + fragment + `)` + `)?` +
		`)>`
)

// This scanner is used for matching resource entities.
var resourceScanner = regexp.MustCompile(`^` + resource)

// This function returns for the specified string an array of the matching
// subgroups for a resource element. The first item in the array is the entire
// matched string.
func ScanResource(v []byte) []string {
	return bytesToStrings(resourceScanner.FindSubmatch(v))
}

// SYMBOL STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a symbol string. It is useful to include in parsing error messages.
const SymbolSyntax = `
	LETTER    = {unicode letter}
	DIGIT     = {unicode digit}
	IDENTIFIER = LETTER (LETTER | DIGIT)*
	ORDINAL = [1-9][0-9]*
	SYMBOL  = "$" IDENTIFIER+ ("-" ORDINAL)?
`

// These constants are used to form a regular expression for valid symbol
// strings.
const (
	symbol = `\$(` + identifier + `(?:-` + ordinal + `)?)`
)

// This scanner is used for matching symbol strings.
var symbolScanner = regexp.MustCompile(`^` + symbol)

// This function returns for the specified string an array of the matching
// subgroups for a symbol string. The first item in the array is the entire
// matched string.
func ScanSymbol(v []byte) []string {
	return bytesToStrings(symbolScanner.FindSubmatch(v))
}

// TAG ELEMENT SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a tag element. It is useful to include in parsing error messages.
const TagSyntax = `
	BASE32 = [0-9A-DF-HJ-NP-TV-Z]
	TAG    = "#" BASE32+
`

// These constants are used to form a regular expression for valid tag
// entities.
const (
	base32 = `[0-9A-DF-HJ-NP-TV-Z]` // "E", "I", "O", and "U" have been removed.
	tag    = `#(` + base32 + `+)`
)

// This scanner is used for matching tag entities.
var tagScanner = regexp.MustCompile(`^` + tag)

// This function returns for the specified string an array of the matching
// subgroups for a tag element. The first item in the array is the entire
// matched string.
func ScanTag(v []byte) []string {
	return bytesToStrings(tagScanner.FindSubmatch(v))
}

// VERSION STRING SYNTAX

// This raw string captures the pseudo-grammar used to define the syntax of
// a version string. It is useful to include in parsing error messages.
const VersionSyntax = `
	ORDINAL = [1-9][0-9]*
	VERSION = "v" ORDINAL ("." ORDINAL)*
`

// These constants are used to form a regular expression for valid version
// strings.
const (
	version = `v(` + ordinal + `(?:\.` + ordinal + `)*)` // Cannot capture each ordinal...
)

// This scanner is used for matching version strings.
var versionScanner = regexp.MustCompile(`^` + version)

// This function returns for the specified string an array of the matching
// subgroups for a version string. The first item in the array is the entire
// matched string.
func ScanVersion(v []byte) []string {
	return bytesToStrings(versionScanner.FindSubmatch(v))
}

// PRIVATE CONSTANTS

// The following constants define some important byte sequences.
var (
	eol            = []byte("\n")
	tab            = []byte("\t")
	doubleQuote    = []byte(`"`)
	doubleQuoteEOL = []byte(`"` + "\n")
	bangStar       = []byte("!*")
	starBang       = []byte("*!")
)

// This array contains the full set of keywords used by the Bali Document
// Notation™ (BDN). They are in reverse order for proper matching.
var keywords = [][]byte{
	[]byte("with"),
	[]byte("while"),
	[]byte("undefined"),
	[]byte("true"),
	[]byte("to"),
	[]byte("throw"),
	[]byte("select"),
	[]byte("save"),
	[]byte("return"),
	[]byte("retrieve"),
	[]byte("reject"),
	[]byte("publish"),
	[]byte("post"),
	[]byte("notarize"),
	[]byte("none"),
	[]byte("matching"),
	[]byte("loop"),
	[]byte("level"),
	[]byte("infinity"),
	[]byte("in"),
	[]byte("if"),
	[]byte("handle"),
	[]byte("from"),
	[]byte("false"),
	[]byte("each"),
	[]byte("do"),
	[]byte("discard"),
	[]byte("continue"),
	[]byte("checkout"),
	[]byte("break"),
	[]byte("at"),
	[]byte("as"),
	[]byte("any"),
	[]byte("accept"),
	[]byte("XOR"),
	[]byte("SANS"),
	[]byte("OR"),
	[]byte("NOT"),
	[]byte("MATCHES"),
	[]byte("IS"),
	[]byte("AND")}

// This array contains the full set of delimiters used by the Bali Document
// Notation™ (BDN). They are in reverse order for proper matching.
var delimiters = [][]byte{
	[]byte("}"),
	[]byte("|"),
	[]byte("{"),
	[]byte("^"),
	[]byte("]"),
	[]byte("["),
	[]byte("@"),
	[]byte("?"),
	[]byte(">"),
	[]byte("="),
	[]byte("<..<"),
	[]byte("<.."),
	[]byte("<-"),
	[]byte("<"),
	[]byte(";"),
	[]byte(":="),
	[]byte(":"),
	[]byte("/="),
	[]byte("//"),
	[]byte("/"),
	[]byte("..<"),
	[]byte(".."),
	[]byte("."),
	[]byte("-="),
	[]byte("-"),
	[]byte(","),
	[]byte("+="),
	[]byte("+"),
	[]byte("*="),
	[]byte("*"),
	[]byte(")"),
	[]byte("("),
	[]byte("&")}

// PRIVATE FUNCTIONS

func bytesToStrings(bytes [][]byte) []string {
	strings := make([]string, len(bytes))
	for index, array := range bytes {
		strings[index] = string(array)
	}
	return strings
}
