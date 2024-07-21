package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 600

var game = gb.NewGame(
	"Vector Text",
	WIDTH,
	HEIGHT,
	true,
)

func init() {

	game.Reset = func() {}

	numbers := game.NewGameObject("numbers", gb.WithPosition(WIDTH/2, 60), gb.WithScale(8, 8)).
		NewVecText("ABCDEFG 0123456789", 1, rl.Red, gb.WithAlignment(gb.TextCenter))

	game.AddScene("myscene", numbers)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
