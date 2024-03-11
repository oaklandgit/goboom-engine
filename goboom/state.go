package goboom

type StateMachine struct {
	GameObj *GameObj
	States map[string]State
	CurrState string
}

func (*StateMachine) Id() string {
	return "stateMachine"
}

type StateMachineOption func(*StateMachine)

func (obj *GameObj) NewStateMachine(
	states map[string]State,
	opts ...StateMachineOption) *GameObj {

	sm := &StateMachine{
		GameObj: obj,
		States: states,
	}

	for _, opt := range opts {
		opt(sm)
	}

	obj.AddComponents(sm)

	return obj
}

func (sm *StateMachine) SetState(state string) {
	sm.CurrState = state
}

func (sm *StateMachine) Update() {
	// sm.States[sm.CurrState].Update()
}

func (sm *StateMachine) AddState(name string, state State) {
	sm.States[name] = state
}

func (sm *StateMachine) GetState(name string) State {
	return sm.States[name]
}

func (sm *StateMachine) Draw() {

}