package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createShip(
	x, y float32,
	speed float32,
	heading float32,
	) *GameObj {

	ship := NewGameObject("Spaceship",
		WithPosition(x, y),
	)
	
	img := rl.GenImageColor(
		int(16),
		int(16),
		rl.Color{0, 0, 0, 0},) 

	rl.ImageDrawLine(img, 7, 0, 0, 14, rl.White)
	rl.ImageDrawLine(img, 7, 0, 14, 14, rl.White)

	ship.NewGraphics(*img)

	ship.NewMotion(
		WithSpeed(speed),
		WithHeading(heading),
	)

	return ship

}