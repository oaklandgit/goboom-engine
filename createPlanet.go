package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY_THRESHOLD = 300
const IGNORE_WHEN = "docked"

func createPlanet(
	name string,
	texture rl.Texture2D,
	x, y float32,
	speed float32,
	orbitPosition float32,
	rotationSpeed float32,
	radius float32,
	color rl.Color,
	gravity float32,
	target *GameObj) *GameObj {

	scale := (radius * 2) / float32(texture.Width)

	p := NewGameObject(name,
		WithTags("planet", "deadly"),
		WithOrigin(0.5, 0.5),
		WithPosition(x, y),
		WithScale(scale, scale),
	)

	p.NewSprite(
		texture,
		WithColor(color),
	)

	p.NewArea(CircleCollider{Radius: float32(texture.Width) * scale / 2},
		WithCooldown(2 * time.Second),
	)

	// p.NewMotion(
	// 	WithVelocity(speed, heading),
	// 	WithWrap(true, false, float32(texture.Width) * scale / 2),
	// )

	// COMMENTING OUT TO SEE IF IT'S THE PROBLEM WITH HANGS
	
	p.NewAttract(
		[]*GameObj{target},
		gravity,
		GRAVITY_THRESHOLD,
		WithIgnored(IGNORE_WHEN),
	)

	p.NewRotate(rotationSpeed)
	
	p.NewMine()

	shadow := NewGameObject("Shadow", WithScale(scale * 0.94, scale * 0.94))
	shadow.NewSprite(
		textures["assets/shadow.png"],
		WithOpacity(0.8),
	)
	p.AddChildren(shadow)

	return p

}