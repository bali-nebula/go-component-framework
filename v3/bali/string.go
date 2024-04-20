/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import ()

// CLASS ACCESS

// Reference

var stringClass = &stringClass_{
	// This class has no private constants to initialize.
}

// Function

func String() StringClassLike {
	return stringClass
}

// CLASS METHODS

// Target

type stringClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *stringClass_) MakeWithBinary(binary string) StringLike {
	return &string_{
		binary_: binary,
	}
}

func (c *stringClass_) MakeWithBytecode(bytecode string) StringLike {
	return &string_{
		bytecode_: bytecode,
	}
}

func (c *stringClass_) MakeWithName(name string) StringLike {
	return &string_{
		name_: name,
	}
}

func (c *stringClass_) MakeWithNarrative(narrative string) StringLike {
	return &string_{
		narrative_: narrative,
	}
}

func (c *stringClass_) MakeWithQuote(quote string) StringLike {
	return &string_{
		quote_: quote,
	}
}

func (c *stringClass_) MakeWithSymbol(symbol string) StringLike {
	return &string_{
		symbol_: symbol,
	}
}

func (c *stringClass_) MakeWithTag(tag string) StringLike {
	return &string_{
		tag_: tag,
	}
}

func (c *stringClass_) MakeWithVersion(version string) StringLike {
	return &string_{
		version_: version,
	}
}

// Functions

// INSTANCE METHODS

// Target

type string_ struct {
	binary_    string
	bytecode_  string
	name_      string
	narrative_ string
	quote_     string
	symbol_    string
	tag_       string
	version_   string
}

// Attributes

func (v *string_) GetBinary() string {
	return v.binary_
}

func (v *string_) GetBytecode() string {
	return v.bytecode_
}

func (v *string_) GetName() string {
	return v.name_
}

func (v *string_) GetNarrative() string {
	return v.narrative_
}

func (v *string_) GetQuote() string {
	return v.quote_
}

func (v *string_) GetSymbol() string {
	return v.symbol_
}

func (v *string_) GetTag() string {
	return v.tag_
}

func (v *string_) GetVersion() string {
	return v.version_
}

// Public

// Private
