package main

import (
	gb "goboom"
)

var game = gb.NewGame(
	"Hello World",
	600,
	800,
	true,
)

func init() {

	game.Reset = func() {}
	game.LoadTextures("assets/ship.png")

	ship := game.	NewGameObject("ship", gb.WithPosition(300, 400)).
					NewSprite(game.Textures["assets/ship.png"])

	game.AddScene("myscene", ship)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
