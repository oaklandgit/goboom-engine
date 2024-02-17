package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createShip(x, y float32) *GameObj {

	// SHIP
	ship := NewGameObject("Spaceship")
	shipShape := rl.GenImageColor(
		int(16),
		int(16),
		rl.Color{0, 0, 0, 0},) 
	rl.ImageDrawLine(shipShape, 2, 0, 14, 7, rl.White)
	rl.ImageDrawLine(shipShape, 2, 14, 14, 7, rl.White)
	ship.NewSprite(rl.LoadTextureFromImage(shipShape))

	// FLAME
	// flame := NewGameObject("Flame")
	// flameShape := rl.GenImageColor(
	// 	int(3),
	// 	int(3),
	// 	rl.Color{0, 0, 0, 0},)
	// rl.ImageDrawCircle(flameShape, 0, 0, 3, rl.Red)
	// flame.Position.Y = 5.5
	// flame.Position.X = -2
	// flame.NewSprite(rl.LoadTextureFromImage(flameShape))

	// ship.AddChildren(flame)

	ship.NewMotion()
	ship.NewInput(
		KeyHandler{
			KeyPress{rl.KeyLeft, KEY_REPEAT},
			func() {
				ship.Rotation -= 5
			},
		},
		KeyHandler{
			KeyPress{rl.KeyRight, KEY_REPEAT},
			func() {
				ship.Rotation += 5
			},
		},
		KeyHandler{
			KeyPress{rl.KeyUp, KEY_REPEAT},
			func() {
				ship.Components["motion"].(*Motion).Speed += 0.01
				ship.Components["motion"].(*Motion).Heading = ship.Rotation
			},
		},
	)



	



	return ship

}