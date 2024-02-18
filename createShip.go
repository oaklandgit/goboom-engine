package main

import rl "github.com/gen2brain/raylib-go/raylib"

const ROTATE_SPEED = 5

func createShip(x, y float32) *GameObj {

	// SHIP CUSTOM METHODS
	thrust := func(g *GameObj) {
		g.Components["motion"].(*Motion).Speed += 0.01
		g.Components["motion"].(*Motion).Heading = g.Angle
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

	ship.NewMotion()
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
				thrust(ship)
			},
		},
		
	)

	return ship

}