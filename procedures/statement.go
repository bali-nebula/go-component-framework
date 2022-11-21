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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
)

// STATEMENT IMPLEMENTATION

// This constructor creates a new statement.
func Statement(mainClause abs.Clause) abs.StatementLike {
	var v = &statement{}
	// Perform argument validation.
	v.SetMainClause(mainClause)
	return v
}

// This constructor creates a new statement.
func StatementWithHandler(mainClause abs.Clause, onClause abs.OnClauseLike) abs.StatementLike {
	var v = &statement{}
	// Perform argument validation.
	v.SetMainClause(mainClause)
	v.SetOnClause(onClause)
	return v
}

// This type defines the structure and methods associated with a statement.
type statement struct {
	annotation abs.Annotation
	mainClause abs.Clause
	onClause   abs.OnClauseLike
	note       abs.NoteLike
}

// This method returns the annotation for this statement.
func (v *statement) GetAnnotation() abs.Annotation {
	return v.annotation
}

// This method sets the annotation for this statement.
func (v *statement) SetAnnotation(annotation abs.Annotation) {
	v.annotation = annotation
}

// This method returns the main clause for this statement.
func (v *statement) GetMainClause() abs.Clause {
	return v.mainClause
}

// This method sets the main clause for this statement.
func (v *statement) SetMainClause(mainClause abs.Clause) {
	v.mainClause = mainClause
}

// This method returns the note for this statement.
func (v *statement) GetNote() abs.NoteLike {
	return v.note
}

// This method sets the note for this statement.
func (v *statement) SetNote(note abs.NoteLike) {
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
