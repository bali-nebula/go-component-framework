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

// PROCEDURE IMPLEMENTATION

// This constructor creates a new procedure.
func Procedure(statements abs.ListLike[abs.StatementLike]) abs.ProcedureLike {
	var v = &procedure{statements, statements, statements, statements}
	return v
}

// This type defines the structure and methods associated with a procedure.
type procedure struct {
	abs.Sequential[abs.StatementLike]
	abs.Indexed[abs.StatementLike]
	abs.Updatable[abs.StatementLike]
	abs.Malleable[abs.StatementLike]
}
