package main

import (
	"time"

	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SWITCH_TO_START = 6 * time.Second

func createGameOverScene(g *boom.Game) *boom.GameObj {

	gameOver := boom.NewGameObject(
		"Game Over",
		boom.WithPosition(screenW/2, screenH/2),
		boom.WithOrigin(0.5, 0.5),
		boom.WithScale(2, 2))
	gameOver.Size = rl.NewVector2(screenW, screenH)
	gameOver.NewSprite(
		game.Textures["assets/gameover.png"],
		boom.WithColor(rl.Red),
	)

	gameOver.NewTimer(
		SWITCH_TO_START,
		func() {
			game.SetScene("titlescene")
		})

	gameOver.NewInput(
		boom.KeyHandler{
			boom.KeyPress{rl.KeyX, boom.KEY_ONCE},
			func() {
				game.SetScene("level1")
			},
		},
	)

	return gameOver
}