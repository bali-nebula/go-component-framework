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

var mainClauseClass = &mainClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func MainClause() MainClauseClassLike {
	return mainClauseClass
}

// CLASS METHODS

// Target

type mainClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *mainClauseClass_) MakeWithFlow(flow FlowLike) MainClauseLike {
	return &mainClause_{
		flow_: flow,
	}
}

func (c *mainClauseClass_) MakeWithAssignment(assignment AssignmentLike) MainClauseLike {
	return &mainClause_{
		assignment_: assignment,
	}
}

func (c *mainClauseClass_) MakeWithMessaging(messaging MessagingLike) MainClauseLike {
	return &mainClause_{
		messaging_: messaging,
	}
}

func (c *mainClauseClass_) MakeWithRepository(repository RepositoryLike) MainClauseLike {
	return &mainClause_{
		repository_: repository,
	}
}

// Functions

// INSTANCE METHODS

// Target

type mainClause_ struct {
	flow_       FlowLike
	assignment_ AssignmentLike
	messaging_  MessagingLike
	repository_ RepositoryLike
}

// Attributes

func (v *mainClause_) GetFlow() FlowLike {
	return v.flow_
}

func (v *mainClause_) GetAssignment() AssignmentLike {
	return v.assignment_
}

func (v *mainClause_) GetMessaging() MessagingLike {
	return v.messaging_
}

func (v *mainClause_) GetRepository() RepositoryLike {
	return v.repository_
}

// Public

// Private
