package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createGameOverScene(g *Game) *GameObj {

	gameOver := NewGameObject(
		"Game Over",
		WithPosition(screenW/2, screenH/2),
		WithOrigin(0.5, 0.5),
		WithScale(2, 2))
	gameOver.Size = rl.NewVector2(screenW, screenH)
	gameOver.NewSprite(
		textures["assets/gameover.png"],
		WithColor(rl.Red),
	)
	// gameOver.NewInput(
	// 	KeyHandler{
	// 		KeyPress{rl.KeyX, KEY_ONCE},
	// 		func() {
	// 			game.SetScene("level1")
	// 		},
	// 	},
	// )
	return gameOver
}