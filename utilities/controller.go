/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v2"
)

// TYPE DEFINITIONS

type (
	Event string
	State string
	Events col.ListLike[Event]
	States col.ListLike[State]
	Table col.CatalogLike[State, States]
)

const (
	Invalid State = "invalid"
)

// CONTROLLER IMPLEMENTATION

// This constructor creates a new controller using the specified finite state
// machine and possible event types. It enforces the possible states of the
// state machine and allowed transitions between states given a finite set of
// possible event types. It implements a finite state machine with the following
// structure:
//     events:      [ "event1",  "event2", ...  "eventM"]
//     states: 
//         "state1: ["invalid",  "state2", ... "invalid"]
//         "state2: [ "state3",  "stateN", ...  "state1"]
   //                           ...
//         "stateN: [ "state1", "invalid", ...  "state3"]
//     ]
// The first state listed in the state table is the initial state of the finite
// state machine. Transitions to next states marked as "invalid" cannot occur.
func Controller(columns []Event, rows []State, table [][]State) *controller {
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
	var catalog = col.Catalog[State, col.ListLike[State]]()
	var row = 0
	var iterator = col.Iterator[State](states)
	for iterator.HasNext() {
		var state = iterator.GetNext()
		var transitions = col.ListFromArray(table[row])
		catalog.SetValue(state, transitions)
		row++
	}
	return &controller{events, catalog, states.GetValue(1)}
}

// This type defines the structure and methods associated with a controller agent.
type controller struct {
	events col.ListLike[Event]
	table col.CatalogLike[State, col.ListLike[State]]
	current State
}

// SEQUENTIAL INTERFACE

// This method retrieves the current state of the controller.
func (v *controller) GetState() State {
	return v.current
}

// This method determines whether or not the specified event would result in a
// valid state transition.
func (v *controller) IsValid(event Event) bool {
	var column = v.events.GetIndex(event)
	var row = v.table.GetValue(v.current)
	var next = row.GetValue(column)
	return next != Invalid
}

// This method determines whether or not this set is empty.
func (v *controller) TransitionState(event Event) State {
	var column = v.events.GetIndex(event)
	var row = v.table.GetValue(v.current)
	v.current = row.GetValue(column)
	return v.current
}

// This method resets the state of the controller to the initial state.
func (v *controller) ResetState() {
	var states = col.ListFromSequence(v.table.GetKeys())
	v.current = states.GetValue(1)
}
