/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// PUBLIC FUNCTIONS

// This function returns a string describing the type of the specified
// expression.
func GetType(expression abs.Expression) string {
	switch value := expression.(type) {
	case abs.ComponentLike:
		return "ComponentExpression"
	case *intrinsicExpression:
		return "IntrinsicExpression"
	case *variableExpression:
		return "VariableExpression"
	case *precedenceExpression:
		return "PrecedenceExpression"
	case *dereferenceExpression:
		return "DereferenceExpression"
	case *invocationExpression:
		return "InvocationExpression"
	case *subcomponentExpression:
		return "SubcomponentExpression"
	case *chainingExpression:
		return "ChainingExpression"
	case *exponentialExpression:
		return "ExponentialExpression"
	case *inversionExpression:
		return "InversionExpression"
	case *arithmeticExpression:
		return "ArithmeticExpression"
	case *magnitudeExpression:
		return "MagnitudeExpression"
	case *comparisonExpression:
		return "ComparisonExpression"
	case *complementExpression:
		return "ComplementExpression"
	case *logicalExpression:
		return "LogicalExpression"
	default:
		var message = fmt.Sprintf("An invalid expression type was found: %T", value)
		panic(message)
	}
}
