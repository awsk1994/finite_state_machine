package fsm

import (
	"errors"
	"fmt"
)

type BaseMachine struct {
	Transitions  map[State]Transition
	currentState State
}

func NewBaseMachine(transitions map[State]Transition, currentState State) *BaseMachine {
	return &BaseMachine{
		Transitions:  transitions,
		currentState: currentState,
	}
}

func (m *BaseMachine) Fire(event string) error {
	if dst, exist := m.currentEvents()[Event(event)]; exist {
		fmt.Printf("src(%s)--event(%s)-->dst(%s)\n", m.CurrentState(), event, dst)
		m.currentState = dst
		return nil
	}
	return errors.New(fmt.Sprintf(CurrentStateNoEventErrorTemplate, m.CurrentState(), event))
}

func (m BaseMachine) CurrentState() State {
	return m.currentState
}

func (m BaseMachine) currentEvents() Transition {
	return m.Transitions[m.currentState]
}
