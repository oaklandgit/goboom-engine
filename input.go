package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPACEBAR = 32
	UP = rl.KeyUp
	DOWN = rl.KeyDown
	LEFT = rl.KeyLeft
	RIGHT = rl.KeyRight
)

type Press int

const (
	KEY_ONCE Press = iota
	KEY_REPEAT
	KEY_UP
)

type KeyPress struct {
	Key int32
	Mode Press
}

type Action func()

type KeyHandler struct {
	KeyPress
	Action
}

type Input struct {
	GameObj *GameObj
	Handlers map[KeyPress]Action
}

func (o *GameObj) NewInput(actions ...KeyHandler) *Input {

	input := &Input{
		GameObj: o,
		Handlers: make(map[KeyPress]Action),
	}

	for _, action := range actions {
		input.Handlers[action.KeyPress] = action.Action
	}

	o.AddComponents(input)
	return input
}

func AddAction(key int32, mode Press, action Action) (KeyPress, Action) {
	return KeyPress{key, mode}, action
}

func (i *Input) Update() {
	for key, action := range i.Handlers {
		if key.Mode == KEY_ONCE{
			if rl.IsKeyPressed(key.Key) { action() }
		} else if key.Mode == KEY_REPEAT {
			if rl.IsKeyDown(key.Key) { action() }
		} else {
			if rl.IsKeyReleased(key.Key) { action() }
		}
	}
}

func (i *Input) Draw() {
	// no op
}


