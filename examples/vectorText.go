package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"Vector Text",
	600,
	800,
	true,
)

func init() {

	game.Reset = func() {}

	message := game.NewGameObject("message", gb.WithPosition(10, 10)).
		NewVecText("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 6, 16, rl.Yellow)

	game.AddScene("myscene", message)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
