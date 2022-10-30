/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package statements

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// STATEMENT IMPLEMENTATION

// This constructor creates a new statement.
func Statement(mainClause any) abs.StatementLike {
	var v = &statement{}
	// Perform argument validation.
	v.SetMainClause(mainClause)
	return v
}

// This constructor creates a new statement.
func StatementWithHandler(mainClause any, onClause abs.OnClauseLike) abs.StatementLike {
	var v = &statement{}
	// Perform argument validation.
	v.SetMainClause(mainClause)
	v.SetOnClause(onClause)
	return v
}

// This type defines the structure and methods associated with a statement.
type statement struct {
	annotation string
	mainClause any
	onClause   abs.OnClauseLike
	note string
}

// This method returns the annotation for this statement.
func (v *statement) GetAnnotation() string {
	return v.annotation
}

// This method sets the annotation for this statement.
func (v *statement) SetAnnotation(annotation string) {
	v.annotation = annotation
}

// This method returns the main clause for this statement.
func (v *statement) GetMainClause() any {
	return v.mainClause
}

// This method sets the main clause for this statement.
func (v *statement) SetMainClause(mainClause any) {
	v.mainClause = mainClause
}

// This method returns the note for this statement.
func (v *statement) GetNote() string {
	return v.note
}

// This method sets the note for this statement.
func (v *statement) SetNote(note string) {
	v.note = note
}

// This method returns the on clause for this statement.
func (v *statement) GetOnClause() abs.OnClauseLike {
	return v.onClause
}

// This method sets the on clause for this statement.
func (v *statement) SetOnClause(onClause abs.OnClauseLike) {
	v.onClause = onClause
}
