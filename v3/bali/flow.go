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

var flowClass = &flowClass_{
	// This class has no private constants to initialize.
}

// Function

func Flow() FlowClassLike {
	return flowClass
}

// CLASS METHODS

// Target

type flowClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *flowClass_) MakeWithIfClause(ifClause IfClauseLike) FlowLike {
	return &flow_{
		ifClause_: ifClause,
	}
}

func (c *flowClass_) MakeWithSelectClause(selectClause SelectClauseLike) FlowLike {
	return &flow_{
		selectClause_: selectClause,
	}
}

func (c *flowClass_) MakeWithWhileClause(whileClause WhileClauseLike) FlowLike {
	return &flow_{
		whileClause_: whileClause,
	}
}

func (c *flowClass_) MakeWithWithClause(withClause WithClauseLike) FlowLike {
	return &flow_{
		withClause_: withClause,
	}
}

func (c *flowClass_) MakeWithContinueClause(continueClause ContinueClauseLike) FlowLike {
	return &flow_{
		continueClause_: continueClause,
	}
}

func (c *flowClass_) MakeWithBreakClause(breakClause BreakClauseLike) FlowLike {
	return &flow_{
		breakClause_: breakClause,
	}
}

func (c *flowClass_) MakeWithReturnClause(returnClause ReturnClauseLike) FlowLike {
	return &flow_{
		returnClause_: returnClause,
	}
}

func (c *flowClass_) MakeWithThrowClause(throwClause ThrowClauseLike) FlowLike {
	return &flow_{
		throwClause_: throwClause,
	}
}

// Functions

// INSTANCE METHODS

// Target

type flow_ struct {
	ifClause_       IfClauseLike
	selectClause_   SelectClauseLike
	whileClause_    WhileClauseLike
	withClause_     WithClauseLike
	continueClause_ ContinueClauseLike
	breakClause_    BreakClauseLike
	returnClause_   ReturnClauseLike
	throwClause_    ThrowClauseLike
}

// Attributes

func (v *flow_) GetIfClause() IfClauseLike {
	return v.ifClause_
}

func (v *flow_) GetSelectClause() SelectClauseLike {
	return v.selectClause_
}

func (v *flow_) GetWhileClause() WhileClauseLike {
	return v.whileClause_
}

func (v *flow_) GetWithClause() WithClauseLike {
	return v.withClause_
}

func (v *flow_) GetContinueClause() ContinueClauseLike {
	return v.continueClause_
}

func (v *flow_) GetBreakClause() BreakClauseLike {
	return v.breakClause_
}

func (v *flow_) GetReturnClause() ReturnClauseLike {
	return v.returnClause_
}

func (v *flow_) GetThrowClause() ThrowClauseLike {
	return v.throwClause_
}

// Public

// Private
