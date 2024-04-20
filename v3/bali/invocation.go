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

var invocationClass = &invocationClass_{
	// This class has no private constants to initialize.
}

// Function

func Invocation() InvocationClassLike {
	return invocationClass
}

// CLASS METHODS

// Target

type invocationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *invocationClass_) MakeWithAttributes(
	target TargetLike,
	method MethodLike,
	arguments ArgumentsLike,
) InvocationLike {
	return &invocation_{
		target_:    target,
		method_:    method,
		arguments_: arguments,
	}
}

// Functions

// INSTANCE METHODS

// Target

type invocation_ struct {
	target_    TargetLike
	method_    MethodLike
	arguments_ ArgumentsLike
}

// Attributes

func (v *invocation_) GetTarget() TargetLike {
	return v.target_
}

func (v *invocation_) GetMethod() MethodLike {
	return v.method_
}

func (v *invocation_) GetArguments() ArgumentsLike {
	return v.arguments_
}

// Public

// Private
