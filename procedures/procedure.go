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

// PROCEDURE IMPLEMENTATION

// This constructor creates a new procedure.
func Procedure(statements abs.Statements) abs.Procedural {
	var v = &procedure{}
	// Perform argument validation.
	v.SetStatements(statements)
	return v
}

// This type defines the structure and methods associated with a procedure.
type procedure struct {
	statements abs.Statements
}

// This method returns the catalog of statements for this procedure.
func (v *procedure) GetStatements() abs.Statements {
	return v.statements
}

// This method sets the catalog of statements for this procedure.
func (v *procedure) SetStatements(statements abs.Statements) {
	if statements == nil {
		panic("A procedure requires a (potentially empty) statement block.")
	}
	v.statements = statements
}
