package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createExplosion(x, y float32, texture string) *GameObj {

	sound := 
		[]rl.Sound{
			sounds["sounds/explosion1.wav"],
			sounds["sounds/explosion2.wav"],
			sounds["sounds/explosion3.wav"],
		}[rand.Intn(3)]

	rl.SetSoundVolume(sound, 0.2);
	
	rl.PlaySound(sound)

	// EXPLOSION
	e := NewGameObject(
		"splosion",
		WithOrigin(0.5, 0.5),
		WithPosition(x, y))

	// SMOKE CLOUDS
	for i := 0; i < 12; i++ {

		c := NewGameObject(
			fmt.Sprintf("splosion-%d", i),
			WithScale(
				rand.Float32()*0.5,
				rand.Float32()*0.5))

		c.NewMotion(
			WithVelocity(
				rand.Float32() * 1,
				float32(rand.Intn(361))))

		c.NewSprite(textures[texture])
		c.NewRotate(float32(rand.Intn(7) - 3)) // -3 to 3
		c.NewTween(
			rand.Float32() * 0.05,
			func(g *GameObj) *float32 {
				return &g.Components["sprite"].(*Sprite).Opacity
			},
			func() {
				c.Delete()
			},
		)

		e.AddChildren(c)
	}

	return e

}