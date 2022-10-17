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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// This type defines the node structure associated with a Bali Document
// Notation (BDN) statement containing a main clause and an optional
// exception handling clause.
type Statement struct {
	MainClause      any
	ExceptionClause *ExceptionClause
}
