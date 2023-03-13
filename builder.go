package fsm

var types = map[string]struct{}{
	"alert": {},		// Go has no built-in set; thus need to implement manually
}

type BaseBuilder struct {
	transitions  map[State]Transition
	initialState State
	_type string
}

func (b *BaseBuilder) AddTransition(transitions map[State]Transition) *BaseBuilder {
	b.transitions = transitions
	return b
}

func (b *BaseBuilder) SetInitialState(state string) *BaseBuilder {
	b.initialState = State(state)
	return b
}

func (b *BaseBuilder) SetType(_type string) *BaseBuilder {
	if _, exist := types[_type]; exist {
		b._type = _type
	}
	return b
}

func(b *BaseBuilder) Build() Machine {
	switch b._type {
	case "alert":
		return NewAlertMachine(b.transitions, b.initialState)
	default:
		return NewBaseMachine(b.transitions, b.initialState)
	}
}
