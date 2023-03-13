package fsm

import (
	"testing"
)

func TestBaseBuilder(t *testing.T) {
	m := new(BaseBuilder).
		AddTransition(map[State]Transition{
			"s1": {
				"e1": "s2",
				"e2": "s3",
			},
			"s2": {
				"e3": "s1",
			},
		}).
		SetInitialState("s1").
		Build()
	m.Fire("e1")
	m.Fire("e3")
	m.Fire("e2")
	if err := m.Fire("XXX"); err == nil {
		t.Errorf("Error check failed. Should return error")
	}
}

func TestCustomBuilder(t *testing.T) {
	m := new(BaseBuilder).
		AddTransition(map[State]Transition{
			"s1": {
				"e1": "s2",
				"e2": "s3",
			},
			"s2": {
				"e3": "s1",
			},
		}).
		SetType("alert").
		SetInitialState("s1").
		Build()
	m.Fire("e1")
	m.Fire("e3")
	m.Fire("e2")
	m.Fire("XXX")
}