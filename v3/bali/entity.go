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

var entityClass = &entityClass_{
	// This class has no private constants to initialize.
}

// Function

func Entity() EntityClassLike {
	return entityClass
}

// CLASS METHODS

// Target

type entityClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *entityClass_) MakeWithElement(element ElementLike) EntityLike {
	return &entity_{
		element_: element,
	}
}

func (c *entityClass_) MakeWithString(string_ StringLike) EntityLike {
	return &entity_{
		string_: string_,
	}
}

func (c *entityClass_) MakeWithRange(range_ RangeLike) EntityLike {
	return &entity_{
		range_: range_,
	}
}

func (c *entityClass_) MakeWithCollection(collection CollectionLike) EntityLike {
	return &entity_{
		collection_: collection,
	}
}

func (c *entityClass_) MakeWithProcedure(procedure ProcedureLike) EntityLike {
	return &entity_{
		procedure_: procedure,
	}
}

// Functions

// INSTANCE METHODS

// Target

type entity_ struct {
	element_    ElementLike
	string_     StringLike
	range_      RangeLike
	collection_ CollectionLike
	procedure_  ProcedureLike
}

// Attributes

func (v *entity_) GetElement() ElementLike {
	return v.element_
}

func (v *entity_) GetString() StringLike {
	return v.string_
}

func (v *entity_) GetRange() RangeLike {
	return v.range_
}

func (v *entity_) GetCollection() CollectionLike {
	return v.collection_
}

func (v *entity_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// Public

// Private
