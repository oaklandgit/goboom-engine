package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func createTitleScene(game *Game) *GameObj {
	titleScene := NewGameObject(
		"Space Miner Start!",
		WithPosition(screenW/2, screenH/2),
		WithOrigin(0.5, 0.5),
		WithScale(2, 2))
	titleScene.Size = rl.NewVector2(screenW, screenH)
	titleScene.NewSprite(
		textures["assets/title.png"],
		)
	titleScene.NewInput(
		KeyHandler{
			KeyPress{rl.KeyX, KEY_ONCE},
			func() {
				game.SetScene("level1")
			},
		},
	)

	return titleScene
	
}