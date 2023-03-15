/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// CONTROLLER IMPLEMENTATION

// This constructor creates a new controller using the specified finite state
// machine and possible event types. It enforces the possible states of the
// state machine and allowed transitions between states given a finite set of
// possible event types. It implements a finite state machine with the following
// table structure:
//
//	states: [events, event1,  event2,  ... eventM ]
//	        [state1, invalid, state2,  ... invalid]
//	        [state2, state3,  stateN,  ... invalid]
//	                         ...
//	        [stateN, state1,  invalid, ... state3 ]
//
// The first row of the state table defines the possible events that can occur.
// Each subsequent row defines a state and the possible transitions from that
// state to the next state for each possible event. Transitions marked as
// "invalid" cannot occur. The second row of the state table defines the initial
// state of the finite state machine.
func Controller(states [][]int) abs.ControllerLike {
	// Validate the size of the state table.
	var height = len(states)
	if height < 3 {
		var message = fmt.Sprintf("The state table must have at least two possible states: %v\n", height)
		panic(message)
	}
	var width = len(states[0])
	if width < 2 {
		var message = fmt.Sprintf("The state table must have at least one possible event: %v\n", width)
		panic(message)
	}
	for _, row := range states {
		if len(row) != width {
			var message = fmt.Sprintf("Each row in the state table must be the same width: %v\n", width)
			panic(message)
		}
	}
	return &controller{states, 1}
}

// This type defines the structure and methods associated with a controller agent.
type controller struct {
	states  [][]int
	current int
}

// PUBLIC INTERFACE

// This method retrieves the current state of the controller.
func (v *controller) GetState() int {
	return v.current
}

// This method sets the state of the controller to the specified state.
func (v *controller) SetState(state int) {
	if state < 1 || state >= len(v.states) {
		var message = fmt.Sprintf("Attempted to set an invalid state: %v", state)
		panic(message)
	}
	v.current = state
}

// This method determines whether or not this set is empty.
func (v *controller) TransitionState(event int) int {
	if event < 1 || event >= len(v.states[0]) {
		var message = fmt.Sprintf("Received an invalid event: %v", event)
		panic(message)
	}
	var next = v.states[v.current][event]
	if next < 1 || next >= len(v.states) {
		var message = fmt.Sprintf("Attempted to transition to an invalid state: %v", next)
		panic(message)
	}
	v.current = next
	return v.current
}
