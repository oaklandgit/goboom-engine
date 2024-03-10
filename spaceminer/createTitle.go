package main

import (
	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createTitleScene(game *boom.Game) *boom.GameObj {
	titleScene := game.NewGameObject(
		"Space Miner Start!",
		boom.WithPosition(game.Width/2, game.Height/2),
		boom.WithOrigin(0.5, 0.5),
		boom.WithScale(2, 2))
	titleScene.Size = rl.NewVector2(game.Width/2, game.Height/2)
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