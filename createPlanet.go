package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createPlanet(
	name string,
	x, y float32,
	speed float32,
	rotationSpeed float32,
	heading float32,
	scale float32,
	color rl.Color,
	opacity float32) *GameObj {

	tex := textures["assets/planet.png"]

	p := NewGameObject(name,
		WithOrigin(0.5, 0.5),
		WithPosition(x, y),
		WithScale(scale, scale),
	)

	p.NewSprite(
		tex,
		WithColor(color),
		WithOpacity(opacity),
	)

	p.NewMotion(
		WithVelocity(speed, heading),
		WithWrap(true, false, float32(tex.Width) * scale / 2),
	)

	p.NewRotate(rotationSpeed)

	return p

}