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

var intrinsicClass = &intrinsicClass_{
	// This class has no private constants to initialize.
}

// Function

func Intrinsic() IntrinsicClassLike {
	return intrinsicClass
}

// CLASS METHODS

// Target

type intrinsicClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *intrinsicClass_) MakeWithAttributes(
	function FunctionLike,
	arguments ArgumentsLike,
) IntrinsicLike {
	return &intrinsic_{
		function_:  function,
		arguments_: arguments,
	}
}

// Functions

// INSTANCE METHODS

// Target

type intrinsic_ struct {
	function_  FunctionLike
	arguments_ ArgumentsLike
}

// Attributes

func (v *intrinsic_) GetFunction() FunctionLike {
	return v.function_
}

func (v *intrinsic_) GetArguments() ArgumentsLike {
	return v.arguments_
}

// Public

// Private
