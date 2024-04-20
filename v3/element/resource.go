/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import ()

// CLASS ACCESS

// Reference

var resourceClass = &resourceClass_{
	// This class has no private constants to initialize.
}

// Function

func Resource() ResourceClassLike {
	return resourceClass
}

// CLASS METHODS

// Target

type resourceClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *resourceClass_) MakeFromString(string_ string) ResourceLike {
	return &resource_{}
}

// Functions

// INSTANCE METHODS

// Target

type resource_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Lexical

func (v *resource_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Segmented

func (v *resource_) GetScheme() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *resource_) GetAuthority() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *resource_) GetPath() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *resource_) GetQuery() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *resource_) GetFragment() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Public

// Private
