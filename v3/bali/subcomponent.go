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

var subcomponentClass = &subcomponentClass_{
	// This class has no private constants to initialize.
}

// Function

func Subcomponent() SubcomponentClassLike {
	return subcomponentClass
}

// CLASS METHODS

// Target

type subcomponentClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *subcomponentClass_) MakeWithAttributes(
	composite CompositeLike,
	indices IndicesLike,
) SubcomponentLike {
	return &subcomponent_{
		composite_: composite,
		indices_:   indices,
	}
}

// Functions

// INSTANCE METHODS

// Target

type subcomponent_ struct {
	composite_ CompositeLike
	indices_   IndicesLike
}

// Attributes

func (v *subcomponent_) GetComposite() CompositeLike {
	return v.composite_
}

func (v *subcomponent_) GetIndices() IndicesLike {
	return v.indices_
}

// Public

// Private
