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

var documentClass = &documentClass_{
	// This class has no private constants to initialize.
}

// Function

func Document() DocumentClassLike {
	return documentClass
}

// CLASS METHODS

// Target

type documentClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *documentClass_) MakeWithAttributes(
	header HeaderLike,
	component ComponentLike,
) DocumentLike {
	return &document_{
		header_:    header,
		component_: component,
	}
}

// Functions

// INSTANCE METHODS

// Target

type document_ struct {
	header_    HeaderLike
	component_ ComponentLike
}

// Attributes

func (v *document_) GetHeader() HeaderLike {
	return v.header_
}

func (v *document_) GetComponent() ComponentLike {
	return v.component_
}

// Public

// Private
