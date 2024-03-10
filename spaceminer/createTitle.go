package main

import (
	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createTitleScene(game *boom.Game) *boom.GameObj {
	titleScene := boom.NewGameObject(
		"Space Miner Start!",
		boom.WithPosition(screenW/2, screenH/2),
		boom.WithOrigin(0.5, 0.5),
		boom.WithScale(2, 2))
	titleScene.Size = rl.NewVector2(screenW, screenH)
	titleScene.NewSprite(
		game.Textures["assets/title.png"],
		)
	titleScene.NewInput(
		boom.KeyHandler{
			boom.KeyPress{rl.KeyX, boom.KEY_ONCE},
			func() {
				game.SetScene("level1")
			},
		},
	)

	return titleScene
	
}