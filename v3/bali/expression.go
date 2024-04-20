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

var expressionClass = &expressionClass_{
	// This class has no private constants to initialize.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *expressionClass_) MakeWithComponent(component ComponentLike) ExpressionLike {
	return &expression_{
		component_: component,
	}
}

func (c *expressionClass_) MakeWithIntrinsic(intrinsic IntrinsicLike) ExpressionLike {
	return &expression_{
		intrinsic_: intrinsic,
	}
}

func (c *expressionClass_) MakeWithVariable(variable VariableLike) ExpressionLike {
	return &expression_{
		variable_: variable,
	}
}

func (c *expressionClass_) MakeWithPrecedence(precedence PrecedenceLike) ExpressionLike {
	return &expression_{
		precedence_: precedence,
	}
}

func (c *expressionClass_) MakeWithDereference(dereference DereferenceLike) ExpressionLike {
	return &expression_{
		dereference_: dereference,
	}
}

func (c *expressionClass_) MakeWithInvocation(invocation InvocationLike) ExpressionLike {
	return &expression_{
		invocation_: invocation,
	}
}

func (c *expressionClass_) MakeWithSubcomponent(subcomponent SubcomponentLike) ExpressionLike {
	return &expression_{
		subcomponent_: subcomponent,
	}
}

func (c *expressionClass_) MakeWithChaining(chaining ChainingLike) ExpressionLike {
	return &expression_{
		chaining_: chaining,
	}
}

func (c *expressionClass_) MakeWithExponential(exponential ExponentialLike) ExpressionLike {
	return &expression_{
		exponential_: exponential,
	}
}

func (c *expressionClass_) MakeWithInversion(inversion InversionLike) ExpressionLike {
	return &expression_{
		inversion_: inversion,
	}
}

func (c *expressionClass_) MakeWithArithmetic(arithmetic ArithmeticLike) ExpressionLike {
	return &expression_{
		arithmetic_: arithmetic,
	}
}

func (c *expressionClass_) MakeWithMagnitude(magnitude MagnitudeLike) ExpressionLike {
	return &expression_{
		magnitude_: magnitude,
	}
}

func (c *expressionClass_) MakeWithComparison(comparison ComparisonLike) ExpressionLike {
	return &expression_{
		comparison_: comparison,
	}
}

func (c *expressionClass_) MakeWithComplement(complement ComplementLike) ExpressionLike {
	return &expression_{
		complement_: complement,
	}
}

func (c *expressionClass_) MakeWithLogical(logical LogicalLike) ExpressionLike {
	return &expression_{
		logical_: logical,
	}
}

// Functions

// INSTANCE METHODS

// Target

type expression_ struct {
	component_    ComponentLike
	intrinsic_    IntrinsicLike
	variable_     VariableLike
	precedence_   PrecedenceLike
	dereference_  DereferenceLike
	invocation_   InvocationLike
	subcomponent_ SubcomponentLike
	chaining_     ChainingLike
	exponential_  ExponentialLike
	inversion_    InversionLike
	arithmetic_   ArithmeticLike
	magnitude_    MagnitudeLike
	comparison_   ComparisonLike
	complement_   ComplementLike
	logical_      LogicalLike
}

// Attributes

func (v *expression_) GetComponent() ComponentLike {
	return v.component_
}

func (v *expression_) GetIntrinsic() IntrinsicLike {
	return v.intrinsic_
}

func (v *expression_) GetVariable() VariableLike {
	return v.variable_
}

func (v *expression_) GetPrecedence() PrecedenceLike {
	return v.precedence_
}

func (v *expression_) GetDereference() DereferenceLike {
	return v.dereference_
}

func (v *expression_) GetInvocation() InvocationLike {
	return v.invocation_
}

func (v *expression_) GetSubcomponent() SubcomponentLike {
	return v.subcomponent_
}

func (v *expression_) GetChaining() ChainingLike {
	return v.chaining_
}

func (v *expression_) GetExponential() ExponentialLike {
	return v.exponential_
}

func (v *expression_) GetInversion() InversionLike {
	return v.inversion_
}

func (v *expression_) GetArithmetic() ArithmeticLike {
	return v.arithmetic_
}

func (v *expression_) GetMagnitude() MagnitudeLike {
	return v.magnitude_
}

func (v *expression_) GetComparison() ComparisonLike {
	return v.comparison_
}

func (v *expression_) GetComplement() ComplementLike {
	return v.complement_
}

func (v *expression_) GetLogical() LogicalLike {
	return v.logical_
}

// Public

// Private
