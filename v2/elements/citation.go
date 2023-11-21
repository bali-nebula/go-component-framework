/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	sts "strings"
)

// CITATION CLASS DEFINITION

// This public singleton creates a unique name space for the citation class.
var Citation = &citations_{}

// This private type defines the structure associated with the class constants
// and class methods for the citation elements.
type citations_ struct {
	// This class has no class constants.
}

// CITATION CLASS METHODS

// This constructor creates a new citation from the specified string value.
func (t *citations_) FromString(string_ string) abs.CitationLike {
	var matches = uti.CitationMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a citation from an invalid string: %v", string_)
		panic(message)
	}
	var citation = citation_(matches[0])
	return citation
}

// CITATION INSTANCE METHODS

// This private type implements the CitationLike interface.  It extends the
// native Go `string` type.
type citation_ string

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v citation_) AsString() string {
	return string(v)
}

// NAMED INTERFACE

// This method returns the name part of this citation.
func (v citation_) GetName() string {
	var index = sts.LastIndex(string(v), "/")
	return string(v[:index])
}

// VERSIONED INTERFACE

// This method returns the version part of this citation.
func (v citation_) GetVersion() string {
	var index = sts.LastIndex(string(v), "/") + 1 // Skip the '/' character.
	return string(v[index:])
}
