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

var statementClass = &statementClass_{
	// This class has no private constants to initialize.
}

// Function

func Statement() StatementClassLike {
	return statementClass
}

// CLASS METHODS

// Target

type statementClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *statementClass_) MakeWithAttributes(
	mainClause MainClauseLike,
	onClause OnClauseLike,
) StatementLike {
	return &statement_{
		mainClause_: mainClause,
		onClause_:   onClause,
	}
}

// Functions

// INSTANCE METHODS

// Target

type statement_ struct {
	mainClause_ MainClauseLike
	onClause_   OnClauseLike
}

// Attributes

func (v *statement_) GetMainClause() MainClauseLike {
	return v.mainClause_
}

func (v *statement_) GetOnClause() OnClauseLike {
	return v.onClause_
}

// Public

// Private
