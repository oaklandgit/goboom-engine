package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createPlanet(
	name string,
	x, y float32,
	scale float32,
	color rl.Color,
	opacity float32) *GameObj {

	tex := rl.LoadTexture("assets/planet.png")

	p := NewGameObject(name,
		WithPosition(x, y),
		WithScale(scale, scale),
	)

	p.NewSprite(
		tex,
		WithColor(color),
		WithOpacity(opacity),
	)

	return p

}