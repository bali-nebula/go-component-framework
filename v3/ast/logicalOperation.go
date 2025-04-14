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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func LogicalOperationClass() LogicalOperationClassLike {
	return logicalOperationClass()
}

// Constructor Methods

func (c *logicalOperationClass_) LogicalOperation(
	logicOperator LogicOperatorLike,
	logical LogicalLike,
) LogicalOperationLike {
	if uti.IsUndefined(logicOperator) {
		panic("The \"logicOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(logical) {
		panic("The \"logical\" attribute is required by this class.")
	}
	var instance = &logicalOperation_{
		// Initialize the instance attributes.
		logicOperator_: logicOperator,
		logical_:       logical,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *logicalOperation_) GetClass() LogicalOperationClassLike {
	return logicalOperationClass()
}

// Attribute Methods

func (v *logicalOperation_) GetLogicOperator() LogicOperatorLike {
	return v.logicOperator_
}

func (v *logicalOperation_) GetLogical() LogicalLike {
	return v.logical_
}

// PROTECTED INTERFACE

// Instance Structure

type logicalOperation_ struct {
	// Declare the instance attributes.
	logicOperator_ LogicOperatorLike
	logical_       LogicalLike
}

// Class Structure

type logicalOperationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func logicalOperationClass() *logicalOperationClass_ {
	return logicalOperationClassReference_
}

var logicalOperationClassReference_ = &logicalOperationClass_{
	// Initialize the class constants.
}
