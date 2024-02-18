package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 5
const THRUST_SPEED = 0.1

func createShip(x, y float32) *GameObj {

	// SHIP METHODS
	thrust := func(g *GameObj, speed float32) {
		g.Components["motion"].(*Motion).SetVelocity(speed, g.Angle)
	}

	rotateCW := func(g *GameObj) {
		g.Angle += ROTATE_SPEED
	}

	rotateCCW := func(g *GameObj) {
		g.Angle -= ROTATE_SPEED
	}

	// SHIP
	ship := NewGameObject("Spaceship")
	ship.NewSprite(textures["assets/ship.png"])

	ship.NewMotion(WithWrap(true, true, 0))
	ship.NewInput(
		KeyHandler{
			KeyPress{rl.KeyLeft, KEY_REPEAT},
			func() {
				rotateCCW(ship)
			},
		},
		KeyHandler{
			KeyPress{rl.KeyRight, KEY_REPEAT},
			func() {
				rotateCW(ship)
			},
		},
		KeyHandler{
			KeyPress{rl.KeyUp, KEY_REPEAT},
			func() {
				thrust(ship, THRUST_SPEED)
			},
		},
		
	)

	return ship

}