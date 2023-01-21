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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// CONTROLLER IMPLEMENTATION

// This constructor creates a new controller using the specified finite state
// machine and possible event types. It enforces the possible states of the
// state machine and allowed transitions between states given a finite set of
// possible event types. It implements a finite state machine with the following
// structure:
//
//	events:      [ "event1",  "event2", ...  "eventM"]
//	states:
//	    "state1: ["invalid",  "state2", ... "invalid"]
//	    "state2: [ "state3",  "stateN", ...  "state1"]
//	                        ...
//	    "stateN: [ "state1", "invalid", ...  "state3"]
//	]
//
// The first state listed in the state table is the initial state of the finite
// state machine. Transitions to next states marked as "invalid" cannot occur.
func Controller(columns []abs.Event, rows []abs.State, table [][]abs.State, state abs.State) abs.ControllerLike {
	// Check for empty event types or states.
	var events = col.ListFromArray(columns)
	var states = col.ListFromArray(rows)
	if events.IsEmpty() || states.IsEmpty() {
		var message = "The controller must manage at least one event type and one state."
		panic(message)
	}
	// Validate the size of the state table.
	var width = events.GetSize()
	var height = states.GetSize()
	if len(table[0]) != width {
		var message = fmt.Sprintf(
			"The width of the state table must be the same as the number of event types: %v\n",
			width)
		panic(message)
	}
	if len(table) != height {
		var message = fmt.Sprintf(
			"The height of the state table must be the same as the number of state types: %v\n",
			height)
		panic(message)
	}
	// Construct the catalog of state transitions.
	var catalog = col.Catalog[abs.State, col.ListLike[abs.State]]()
	var row = 0
	var iterator = col.Iterator[abs.State](states)
	for iterator.HasNext() {
		var state = iterator.GetNext()
		var transitions = col.ListFromArray(table[row])
		catalog.SetValue(state, transitions)
		row++
	}
	return &controller{events, states, catalog, state}
}

// This type defines the structure and methods associated with a controller agent.
type controller struct {
	events  col.ListLike[abs.Event]
	states  col.ListLike[abs.State]
	table   col.CatalogLike[abs.State, col.ListLike[abs.State]]
	current abs.State
}

// PUBLIC INTERFACE

// This method retrieves the current state of the controller.
func (v *controller) GetState() abs.State {
	return v.current
}

// This method sets the state of the controller to the specified state.
func (v *controller) SetState(state abs.State) {
	if !v.states.ContainsValue(state) {
		var message = fmt.Sprintf("Attempted to set an invalid state: %v", state)
		panic(message)
	}
	v.current = state
}

// This method determines whether or not this set is empty.
func (v *controller) TransitionState(event abs.Event) abs.State {
	var column = v.events.GetIndex(event)
	var row = v.table.GetValue(v.current)
	if row == nil {
		var message = fmt.Sprintf("The controller is in an invalid state: %v", v.current)
		panic(message)
	}
	var next = row.GetValue(column)
	if !v.states.ContainsValue(next) {
		var message = fmt.Sprintf("Attempted to transition to an invalid state: %v", next)
		panic(message)
	}
	v.current = next
	return v.current
}
