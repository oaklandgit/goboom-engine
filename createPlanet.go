package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY_THRESHOLD = 300
const IGNORE_WHEN = "docked"

func createPlanet(
	name string,
	x, y float32,
	speed float32,
	rotationSpeed float32,
	heading float32,
	radius float32,
	color rl.Color,
	gravity float32,
	target *GameObj,
	opacity float32) *GameObj {

	tex := textures["assets/rocky3.png"]

	scale := (radius * 2) / float32(tex.Width)

	p := NewGameObject(name,
		WithTags("planet", "deadly"),
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

	p.NewAttract(
		[]*GameObj{target},
		gravity,
		GRAVITY_THRESHOLD,
		WithIgnored(IGNORE_WHEN),
	)

	p.NewRotate(rotationSpeed)
	
	p.NewMine()

	shadow := NewGameObject("Shadow", WithScale(scale * 0.88, scale * 0.88))
	shadow.NewSprite(
		textures["assets/shadow.png"],
		WithOpacity(0.8),
	)
	p.AddChildren(shadow)

	return p

}