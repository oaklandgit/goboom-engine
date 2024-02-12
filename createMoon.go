package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createMoon(
	name string,
	planet *GameObj,
	speed float32,
	angle float32,
	scale float32,
	distance float32,
	color rl.Color,
	opacity float32) *GameObj {

	tex := rl.LoadTexture("assets/planet.png")

	m := NewGameObject(name,
		WithScale(scale, scale),
	)

	m.NewSprite(
		tex,
		WithColor(color),
		WithOpacity(opacity),
	)

	m.NewOrbit(planet, speed, distance)

	return m

}