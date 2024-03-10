package main

import (
	"time"

	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY_THRESHOLD = 300
const IGNORE_WHEN = "docked"

func createPlanet(
	name string,
	texture rl.Texture2D,
	x, y float32,
	speed float32,
	rotationSpeed float32,
	radius float32,
	color rl.Color,
	gravity float32,
	target *boom.GameObj) *boom.GameObj {

	scale := (radius * 2) / float32(texture.Width)

	p := game.NewGameObject(name,
		boom.WithTags("planet", "deadly"),
		boom.WithOrigin(0.5, 0.5),
		boom.WithPosition(x, y),
		boom.WithScale(scale, scale),
	)

	p.NewSprite(
		texture,
		boom.WithColor(color),
	)

	p.NewArea(boom.CircleCollider{Radius: float32(texture.Width) * scale / 2},
		boom.WithCooldown(2 * time.Second),
	)
	
	// p.NewAttract(
	// 	[]*GameObj{target},
	// 	gravity,
	// 	GRAVITY_THRESHOLD,
	// 	WithIgnored(IGNORE_WHEN),
	// )

	p.NewRotate(rotationSpeed)
	NewMine(p)

	return p

}