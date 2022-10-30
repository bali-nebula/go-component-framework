/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// DISCARD CLAUSE IMPLEMENTATION

// This constructor creates a new discard clause.
func DiscardClause(citation any) abs.DiscardClauseLike {
	var v = &discardClause{}
	// Perform argument validation.
	v.SetCitation(citation)
	return v
}

// This type defines the structure and methods associated with a discard
// clause.
type discardClause struct {
	citation any
}

// This method returns the citation expression for this discard clause.
func (v *discardClause) GetCitation() any {
	return v.citation
}

// This method sets the citation expression for this discard clause.
func (v *discardClause) SetCitation(citation any) {
	if citation == nil {
		panic("A discard clause requires a citation.")
	}
	v.citation = citation
}
