package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createMoon(
	name string,
	// planet *GameObj,
	speed float32,
	rotationSpeed float32,
	scale float32,
	distance float32,
	color rl.Color,
	opacity float32) *GameObj {

	tex := textures["assets/planet.png"]

	m := NewGameObject(name,
		WithOrigin(0.5, 0.5),
		WithScale(scale, scale),
	)

	m.NewSprite(
		tex,
		WithColor(color),
		WithOpacity(opacity),
	)

	m.NewRotate(rotationSpeed)

	m.NewOrbit(speed, distance)

	return m

}