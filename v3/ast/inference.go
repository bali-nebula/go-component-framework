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

import (
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func InferenceClass() InferenceClassLike {
	return inferenceClass()
}

// Constructor Methods

func (c *inferenceClass_) Inference(
	logical LogicalLike,
	logicalOperations col.Sequential[LogicalOperationLike],
) InferenceLike {
	if uti.IsUndefined(logical) {
		panic("The \"logical\" attribute is required by this class.")
	}
	if uti.IsUndefined(logicalOperations) {
		panic("The \"logicalOperations\" attribute is required by this class.")
	}
	var instance = &inference_{
		// Initialize the instance attributes.
		logical_:           logical,
		logicalOperations_: logicalOperations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inference_) GetClass() InferenceClassLike {
	return inferenceClass()
}

// Attribute Methods

func (v *inference_) GetLogical() LogicalLike {
	return v.logical_
}

func (v *inference_) GetLogicalOperations() col.Sequential[LogicalOperationLike] {
	return v.logicalOperations_
}

// PROTECTED INTERFACE

// Instance Structure

type inference_ struct {
	// Declare the instance attributes.
	logical_           LogicalLike
	logicalOperations_ col.Sequential[LogicalOperationLike]
}

// Class Structure

type inferenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inferenceClass() *inferenceClass_ {
	return inferenceClassReference_
}

var inferenceClassReference_ = &inferenceClass_{
	// Initialize the class constants.
}
