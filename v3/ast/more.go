/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import ()

// CLASS INTERFACE

// Access Function

func MoreClass() MoreClassLike {
	return moreClass()
}

// Constructor Methods

func (c *moreClass_) More() MoreLike {
	var instance = &more_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *more_) GetClass() MoreClassLike {
	return moreClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type more_ struct {
	// Declare the instance attributes.
}

// Class Structure

type moreClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moreClass() *moreClass_ {
	return moreClassReference_
}

var moreClassReference_ = &moreClass_{
	// Initialize the class constants.
}
