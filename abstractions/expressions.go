/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// EXPRESSION INTERFACES

// This interface defines the methods supported by all intrinsic-like types.
type IntrinsicLike interface {
	GetFunction() string
	GetArgument(index int) any // Ordinal based indexing.
	GetArguments() ArrayLike[any]
}

// This interface defines the methods supported by all precedence-like types.
type PrecedenceLike interface {
	GetExpression() any
	SetExpression(expression any)
}
