package fsm

type State string
type Event string
type Transition map[Event]State

type Machine interface {
	Fire(event string) error
	CurrentState() State
}
