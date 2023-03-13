// Package fsm is implementation of finite state machine
package fsm

import "fmt"

type AlertMachine struct {
	BaseMachine // inheritance
}

func NewAlertMachine(transitions map[State]Transition, currentState State) *AlertMachine {
	return &AlertMachine{
		BaseMachine: *NewBaseMachine(transitions, currentState),
	}
}

func (m *AlertMachine) Fire(event string) error {
	if err := m.BaseMachine.Fire(event); err != nil {
		fmt.Printf(">> Sending alert to user, err=%v\n", err) // Simulate Alert
	}
	return nil
}
