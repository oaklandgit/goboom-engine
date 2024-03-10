package main

import (
	"math/rand"

	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_LIFESPAN = 60 // 1 second

func createExplosion(x, y float32, texture string) *boom.GameObj {

	sound := 
		[]rl.Sound{
			game.Sounds["sounds/explosion1.wav"],
			game.Sounds["sounds/explosion2.wav"],
			game.Sounds["sounds/explosion3.wav"],
			game.Sounds["sounds/explosion4.wav"],
			game.Sounds["sounds/explosion5.wav"],
			game.Sounds["sounds/explosion6.wav"],
			game.Sounds["sounds/explosion7.wav"],
		}[rand.Intn(7)]

	rl.SetSoundVolume(sound, 0.2);
	
	rl.PlaySound(sound)

	// EXPLOSION
	e := boom.NewGameObject(
		"explosion",
		boom.WithOrigin(0.5, 0.5),
		boom.WithPosition(x, y),
		boom.WithTags("explosion"))

	e.NewLifespan(MAX_LIFESPAN)

	// SHRAPNEL
	for i := 0; i < 12; i++ {

		c := boom.NewGameObject(
			"shard",
			boom.WithScale(
				rand.Float32()*0.5,
				rand.Float32()*0.5),
			boom.WithTags("shard"))

		c.NewMotion(
			boom.WithVelocity(
				rand.Float32() * 1,
				float32(rand.Intn(361))))

		c.NewSprite(game.Textures[texture])
		c.NewRotate(float32(rand.Intn(7) - 3)) // -3 to 3
		c.NewTween(
			rand.Float32() * 0.05,
			func(g *boom.GameObj) *float32 {
				return &g.Components["sprite"].(*boom.Sprite).Opacity
			},
			func() {
				c.Delete()
			},
		)

		c.Components["tween"].(*boom.Tween).Play()

		e.AddChildren(c)
	}

	return e

}