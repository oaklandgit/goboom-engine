package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createMoon(
	name string,
	texture rl.Texture2D,
	speed float32,
	rotationSpeed float32,
	radius float32,
	distance float32,
	color rl.Color,
	opacity float32) *GameObj {

	// tex := textures["assets/rocky.png"]
	scale := (radius * 2) / float32(texture.Width)


	m := NewGameObject(name,
		WithOrigin(0.5, 0.5),
		WithScale(scale, scale),
		WithTags("moon", "deadly"),
	)

	m.NewSprite(
		texture,
		WithColor(color),
		WithOpacity(opacity),
	)

	m.NewRotate(rotationSpeed)
	m.NewOrbit(speed, distance)
	m.NewArea(CircleCollider{Radius: float32(texture.Width) * scale / 2})

	shadow := NewGameObject("Shadow", WithScale(scale * 0.9, scale * 0.9))
	shadow.NewSprite(
		textures["assets/shadow.png"],
		WithOpacity(0.9),
	)
	m.AddChildren(shadow)
	
	return m

}