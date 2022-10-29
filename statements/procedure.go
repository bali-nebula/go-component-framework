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

// PROCEDURE IMPLEMENTATION

// This constructor creates a new procedure.
func Procedure(statements abs.ListLike[any]) abs.ProcedureLike {
	var v = &procedure{statements, statements, statements, statements}
	return v
}

// This type defines the structure and methods associated with a procedure.
type procedure struct {
	abs.Sequential[any]
	abs.Indexed[any]
	abs.Updatable[any]
	abs.Malleable[any]
}
