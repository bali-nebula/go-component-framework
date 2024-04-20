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

var messagingClass = &messagingClass_{
	// This class has no private constants to initialize.
}

// Function

func Messaging() MessagingClassLike {
	return messagingClass
}

// CLASS METHODS

// Target

type messagingClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *messagingClass_) MakeWithPostClause(postClause PostClauseLike) MessagingLike {
	return &messaging_{
		postClause_: postClause,
	}
}

func (c *messagingClass_) MakeWithRetrieveClause(retrieveClause RetrieveClauseLike) MessagingLike {
	return &messaging_{
		retrieveClause_: retrieveClause,
	}
}

func (c *messagingClass_) MakeWithAcceptClause(acceptClause AcceptClauseLike) MessagingLike {
	return &messaging_{
		acceptClause_: acceptClause,
	}
}

func (c *messagingClass_) MakeWithRejectClause(rejectClause RejectClauseLike) MessagingLike {
	return &messaging_{
		rejectClause_: rejectClause,
	}
}

func (c *messagingClass_) MakeWithPublishClause(publishClause PublishClauseLike) MessagingLike {
	return &messaging_{
		publishClause_: publishClause,
	}
}

// Functions

// INSTANCE METHODS

// Target

type messaging_ struct {
	postClause_     PostClauseLike
	retrieveClause_ RetrieveClauseLike
	acceptClause_   AcceptClauseLike
	rejectClause_   RejectClauseLike
	publishClause_  PublishClauseLike
}

// Attributes

func (v *messaging_) GetPostClause() PostClauseLike {
	return v.postClause_
}

func (v *messaging_) GetRetrieveClause() RetrieveClauseLike {
	return v.retrieveClause_
}

func (v *messaging_) GetAcceptClause() AcceptClauseLike {
	return v.acceptClause_
}

func (v *messaging_) GetRejectClause() RejectClauseLike {
	return v.rejectClause_
}

func (v *messaging_) GetPublishClause() PublishClauseLike {
	return v.publishClause_
}

// Public

// Private
