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

	// letters := game.NewGameObject("letters", gb.WithPosition(10, 10)).
	// 	NewVecText("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 6, 16, rl.Yellow)

	numbers := game.NewGameObject("numbers", gb.WithPosition(10, 50)).
		NewVecText("ABCDEFG 0123456789", 6, 16, rl.Red)

	game.AddScene("myscene", numbers)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
