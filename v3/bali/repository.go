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

var repositoryClass = &repositoryClass_{
	// This class has no private constants to initialize.
}

// Function

func Repository() RepositoryClassLike {
	return repositoryClass
}

// CLASS METHODS

// Target

type repositoryClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *repositoryClass_) MakeWithCheckoutClause(checkoutClause CheckoutClauseLike) RepositoryLike {
	return &repository_{
		checkoutClause_: checkoutClause,
	}
}

func (c *repositoryClass_) MakeWithSaveClause(saveClause SaveClauseLike) RepositoryLike {
	return &repository_{
		saveClause_: saveClause,
	}
}

func (c *repositoryClass_) MakeWithDiscardClause(discardClause DiscardClauseLike) RepositoryLike {
	return &repository_{
		discardClause_: discardClause,
	}
}

func (c *repositoryClass_) MakeWithNotarizeClause(notarizeClause NotarizeClauseLike) RepositoryLike {
	return &repository_{
		notarizeClause_: notarizeClause,
	}
}

// Functions

// INSTANCE METHODS

// Target

type repository_ struct {
	checkoutClause_ CheckoutClauseLike
	saveClause_     SaveClauseLike
	discardClause_  DiscardClauseLike
	notarizeClause_ NotarizeClauseLike
}

// Attributes

func (v *repository_) GetCheckoutClause() CheckoutClauseLike {
	return v.checkoutClause_
}

func (v *repository_) GetSaveClause() SaveClauseLike {
	return v.saveClause_
}

func (v *repository_) GetDiscardClause() DiscardClauseLike {
	return v.discardClause_
}

func (v *repository_) GetNotarizeClause() NotarizeClauseLike {
	return v.notarizeClause_
}

// Public

// Private
