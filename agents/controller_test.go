/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	age "github.com/bali-nebula/go-component-framework/agents"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const (
	invalid int = iota
	state1
	state2
	state3
)

const (
	events int = iota
	initialized
	processed
	finalized
)

func TestController(t *tes.T) {
	var states = [][]int{
		{events, initialized, processed, finalized},
		{state1, state2,      invalid,   invalid},
		{state2, invalid,     state2,    state3},
		{state3, invalid,     invalid,   invalid},
	}

	var controller = age.Controller(states)
	ass.Equal(t, state1, controller.GetState())
	ass.Equal(t, state2, controller.TransitionState(initialized))
	ass.Equal(t, state2, controller.TransitionState(processed))
	ass.Equal(t, state3, controller.TransitionState(finalized))
	controller.SetState(state1)
	ass.Equal(t, state1, controller.GetState())
}
