package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"Polygon",
	600,
	800,
	true,
)

func init() {

	game.Reset = func() {}

	shape := game.NewGameObject("polygon", gb.WithPosition(300, 400)).
		NewRegPoly(12, 50, rl.Red)

	game.AddScene("myscene", shape)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
