package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY_THRESHOLD = 300

func createPlanet(
	name string,
	x, y float32,
	speed float32,
	rotationSpeed float32,
	heading float32,
	scale float32,
	color rl.Color,
	gravity float32,
	target *GameObj,
	opacity float32) *GameObj {

	tex := textures["assets/rocky.png"]

	p := NewGameObject(name,
		WithTags("planet"),
		WithOrigin(0.5, 0.5),
		WithPosition(x, y),
		WithScale(scale, scale),
	)

	p.NewSprite(
		tex,
		WithColor(color),
		WithOpacity(opacity),
	)

	p.NewArea(CircleCollider{Radius: float32(tex.Width) * scale / 2})

	p.NewMotion(
		WithVelocity(speed, heading),
		WithWrap(true, false, float32(tex.Width) * scale / 2),
	)

	p.NewAttract([]*GameObj{target}, gravity, GRAVITY_THRESHOLD)

	p.NewRotate(rotationSpeed)

	shadow := NewGameObject("Shadow", WithScale(scale * 0.95, scale * 0.95))
	shadow.NewSprite(
		textures["assets/shadow.png"],
		WithOpacity(0.8),
	)
	p.AddChildren(shadow)

	return p

}