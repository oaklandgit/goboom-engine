package main

import (
	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createMoon(
	name string,
	texture rl.Texture2D,
	speed float32,
	rotationSpeed float32,
	radius float32,
	distance float32,
	color rl.Color,
	opacity float32) *boom.GameObj {

	scale := (radius * 2) / float32(texture.Width)

	m := game.NewGameObject(name,
		boom.WithTags("moon", "deadly"),
		boom.WithOrigin(0.5, 0.5),
		boom.WithScale(scale, scale),
	)

	m.NewSprite(
		texture,
		boom.WithColor(color),
		boom.WithOpacity(opacity),
	)

	m.NewRotate(rotationSpeed)
	m.NewOrbit(speed, distance)
	m.NewArea(boom.CircleCollider{Radius: radius})

	shadow := game.NewGameObject("Shadow", boom.WithScale(scale * 0.9, scale * 0.9))
	shadow.NewSprite(
		game.Textures["assets/shadow.png"],
		boom.WithOpacity(0.9),
	)
	m.AddChildren(shadow)
	
	return m

}