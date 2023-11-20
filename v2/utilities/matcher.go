/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities

import (
	reg "regexp"
)

// MATCHER IMPLEMENTATION

// These compiled regular expressions are able to match all of the token types.
var (
	AngleMatcher       = reg.MustCompile(`^(?:` + angle + `)`)
	BinaryMatcher      = reg.MustCompile(`^(?:` + binary + `)`)
	BooleanMatcher     = reg.MustCompile(`^(?:` + boolean + `)`)
	BytecodeMatcher    = reg.MustCompile(`^(?:` + bytecode + `)`)
	CharacterMatcher   = reg.MustCompile(`^(?:` + character + `)`)
	CommentMatcher     = reg.MustCompile(`^(?:` + comment + `)`)
	CitationMatcher    = reg.MustCompile(`^(?:` + citation + `)`)
	DelimiterMatcher   = reg.MustCompile(`^(?:` + delimiter + `)`)
	DurationMatcher    = reg.MustCompile(`^(?:` + duration + `)`)
	EolMatcher         = reg.MustCompile(`^(?:` + eol + `)`)
	FloatMatcher       = reg.MustCompile(`^(?:` + float + `)`)
	IdentifierMatcher  = reg.MustCompile(`^(?:` + identifier + `)`)
	IntegerMatcher     = reg.MustCompile(`^(?:` + integer + `)`)
	IntrinsicMatcher   = reg.MustCompile(`^(?:` + intrinsic + `)`)
	KeywordMatcher     = reg.MustCompile(`^(?:` + keyword + `)`)
	MomentMatcher      = reg.MustCompile(`^(?:` + moment + `)`)
	NameMatcher        = reg.MustCompile(`^(?:` + name + `)`)
	NarrativeMatcher   = reg.MustCompile(`^(?:` + narrative + `)`)
	NoteMatcher        = reg.MustCompile(`^(?:` + note + `)`)
	NumberMatcher      = reg.MustCompile(`^(?:` + number + `)`)
	PatternMatcher     = reg.MustCompile(`^(?:` + pattern + `)`)
	PercentageMatcher  = reg.MustCompile(`^(?:` + percentage + `)`)
	ProbabilityMatcher = reg.MustCompile(`^(?:` + probability + `)`)
	QuoteMatcher       = reg.MustCompile(`^(?:` + quote + `)`)
	RealMatcher        = reg.MustCompile(`^(?:` + real_ + `)`)
	ResourceMatcher    = reg.MustCompile(`^(?:` + resource + `)`)
	SymbolMatcher      = reg.MustCompile(`^(?:` + symbol + `)`)
	TagMatcher         = reg.MustCompile(`^(?:` + tag + `)`)
	VersionMatcher     = reg.MustCompile(`^(?:` + version + `)`)
	WhitespaceMatcher  = reg.MustCompile(`^(?:` + whitespace + `)`)
)

// PRIVATE CONSTANTS

// These constant definitions capture the regular expression subpatterns for all
// of the token types.
const (
	angle       = `~(` + magnitude + `|` + zero + `)`
	authority   = `[^/` + control + `]+`
	base16      = `[0-9a-f]`
	base32      = `[0-9A-DF-HJ-NP-TV-Z]` // "E", "I", "O", and "U" have been removed.
	base64      = `[A-Za-z0-9+/]`
	binary      = `'>` + eol + `((?:` + space + `*` + base64 + `+` + eol + `)*)` + space + `*<'`
	boolean     = `false|true`
	bytecode    = `'((?:` + instruction + `(?:` + space + instruction + `)*)*)'`
	character   = `"(` + glyph + `)"`
	citation    = `(` + name + `)/(` + version + `)`
	comment     = `!>` + eol + `((?:.|` + eol + `)*?)` + eol + space + `*<!`
	complex_    = `\((?:` + rectangular + `|` + polar + `)\)`
	control     = `\a\f\n\r\t\v`
	dates       = years + `?` + months + `?` + days + `?`
	day         = `(?:[012][1-9])|(?:[3][01])`
	days        = `(` + span + `D)`
	delimiter   = `≠|~|\}|\||\{|\^|\]|\[|@|\?=|>|=|<-|<|;|:=|:|/=|//|/|\.\.|\.|-=|-|,|\+=|\+|\*=|\*|\)|\(|&`
	digit       = `\pN` // All unicode digits.
	duration    = `~(` + sign + `?)P(?:` + weeks + `|` + dates + `(?:` + times + `)?)`
	e           = `e`
	eol         = `\n`
	escape      = unicode + `|\\[afnrtv"\\]`
	exponent    = `E` + sign + `?` + ordinal
	float       = real_ + `|` + zero + `|` + sign + `?(?:` + infinity + `)`
	fraction    = `\.[0-9]+`
	fragment    = `[^>` + control + `]+`
	glyph       = `(?:` + escape + `|[^"` + control + `])`
	hour        = `(?:[01][0-9])|(?:[2][0-3])`
	hours       = `(` + span + `H)`
	identifier  = letter + `(?:` + letter + `|` + digit + `)*`
	imaginary   = `(?:` + sign + `|` + real_ + `)?i`
	infinity    = `infinity|∞`
	instruction = base16 + base16 + base16 + base16
	integer     = zero + `|` + sign + `?` + ordinal
	intrinsic   = `ANY|LOWER_CASE|UPPER_CASE|DIGIT|ESCAPE|CONTROL|EOF`
	keyword     = `AND|IS|MATCHES|NOT|OR|SANS|XOR|accept|as|at|break|checkout|continue|discard|do|each|from|if|in|let|level|loop|matching|notarize|on|post|publish|reject|retrieve|return|save|select|throw|to|while|with`
	letter      = `\pL` // All unicode letters and connectors like underscores.
	magnitude   = `(?:` + e + `|` + pi + `|` + phi + `|` + tau + `|` + scalar + `)`
	minute      = `[0-5][0-9]`
	minutes     = `(` + span + `M)`
	moment      = `<(` + sign + `)?(` + year + `)(?:-(` + month + `)(?:-(` + day + `)` +
		`(?:T(` + hour + `)(?::(` + minute + `)(?::((?:` + second + `)(?:` + fraction + `)?))?)?)?)?)?>`
	month       = `(?:[0][1-9])|(?:[1][012])`
	months      = `(` + span + `M)`
	name        = `(?:/` + identifier + `)+` // Cannot capture each identifier...
	narrative   = `">` + eol + `((?:.|` + eol + `)*?)` + eol + space + `*<"`
	note        = `! [^` + control + `]*`
	number      = imaginary + `|` + real_ + `|` + complex_ + `|` + zero + `|` + infinity + `|` + undefined
	ordinal     = `[1-9][0-9]*`
	path        = `[^?#>` + control + `]*`
	pattern     = `none` + `|` + regex + `|` + `any`
	percentage  = `(` + real_ + `|` + zero + `)%`
	pi          = `pi|π`
	phi         = `phi|φ`
	polar       = `(` + magnitude + `)(e\^)(` + angle + `i)`
	probability = fraction + `|1\.`
	query       = `[^#>` + control + `]*`
	quote       = `"(` + glyph + `*)"`
	real_       = sign + `?` + magnitude
	rectangular = `(` + real_ + `)(, )(` + imaginary + `)`
	regex       = `"(` + glyph + `+)"\?`
	resource    = `<(` + `(` + scheme + `):` + `(?:` + `//(` + authority + `)` + `)?` + `(` + path + `)` + `(?:` + `\?(` + query + `)` + `)?` + `(?:` + `#(` + fragment + `)` + `)?` + `)>`
	scalar      = `(?:` + ordinal + `(?:` + fraction + `)?|` + zero + fraction + `)(?:` + exponent + `)?`
	scheme      = `[a-zA-Z][0-9a-zA-Z+-.]*`
	second      = `(?:[0-5][0-9])|(?:[6][01])`
	seconds     = `(` + span + `S)`
	separator   = `[-+.]`
	sign        = `[+-]`
	space       = ` `
	symbol      = `\$(` + identifier + `(?:-` + ordinal + `)?)`
	tag         = `#(` + base32 + `+)`
	tau         = `tau|τ`
	times       = `(T)` + hours + `?` + minutes + `?` + seconds + `?`
	span        = `(?:` + zero + `|` + ordinal + `(?:` + fraction + `)?)`
	undefined   = `undefined`
	unicode     = `\\x\{` + base16 + `\}`
	version     = `v(` + ordinal + `(?:\.` + ordinal + `)*)` // Cannot capture each ordinal...
	weeks       = `(` + span + `W)`
	whitespace  = `[` + space + `]+`
	year        = `(?:` + ordinal + `|` + zero + `)`
	years       = `(` + span + `Y)`
	zero        = `0`
)
