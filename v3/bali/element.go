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

var elementClass = &elementClass_{
	// This class has no private constants to initialize.
}

// Function

func Element() ElementClassLike {
	return elementClass
}

// CLASS METHODS

// Target

type elementClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *elementClass_) MakeWithAngle(angle string) ElementLike {
	return &element_{
		angle_: angle,
	}
}

func (c *elementClass_) MakeWithBoolean(boolean string) ElementLike {
	return &element_{
		boolean_: boolean,
	}
}

func (c *elementClass_) MakeWithDuration(duration string) ElementLike {
	return &element_{
		duration_: duration,
	}
}

func (c *elementClass_) MakeWithMoment(moment string) ElementLike {
	return &element_{
		moment_: moment,
	}
}

func (c *elementClass_) MakeWithNumber(number string) ElementLike {
	return &element_{
		number_: number,
	}
}

func (c *elementClass_) MakeWithPattern(pattern string) ElementLike {
	return &element_{
		pattern_: pattern,
	}
}

func (c *elementClass_) MakeWithPercentage(percentage string) ElementLike {
	return &element_{
		percentage_: percentage,
	}
}

func (c *elementClass_) MakeWithProbability(probability string) ElementLike {
	return &element_{
		probability_: probability,
	}
}

func (c *elementClass_) MakeWithResource(resource string) ElementLike {
	return &element_{
		resource_: resource,
	}
}

// Functions

// INSTANCE METHODS

// Target

type element_ struct {
	angle_       string
	boolean_     string
	duration_    string
	moment_      string
	number_      string
	pattern_     string
	percentage_  string
	probability_ string
	resource_    string
}

// Attributes

func (v *element_) GetAngle() string {
	return v.angle_
}

func (v *element_) GetBoolean() string {
	return v.boolean_
}

func (v *element_) GetDuration() string {
	return v.duration_
}

func (v *element_) GetMoment() string {
	return v.moment_
}

func (v *element_) GetNumber() string {
	return v.number_
}

func (v *element_) GetPattern() string {
	return v.pattern_
}

func (v *element_) GetPercentage() string {
	return v.percentage_
}

func (v *element_) GetProbability() string {
	return v.probability_
}

func (v *element_) GetResource() string {
	return v.resource_
}

// Public

// Private
