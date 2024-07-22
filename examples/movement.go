// TO DO:
// Runtime crash when using WithWrap, unless you first create a scene (other than the game itself)
// because the game doesn't have a scene to wrap around. This is a bug.

package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"Movement",
	600,
	800,
	false,
)

func init() {

	game.Reset = func() {}

	shape := "L10 0 L20 10 L60 10 L80 20 L80 30 L50 30 L30 40 L10 40 L20 30 L10 30 L10 20 Z"

	ship := game.NewGameObject("ship", gb.WithScale(1, 0.5), gb.WithPosition(300, 400)).
		NewMotion(gb.WithVelocity(0.1, 0)).
		NewSvgPath(shape, 2, rl.Yellow)

	ship.NewInput(
		gb.KeyHandler{
			KeyPress: gb.KeyPress{Key: rl.KeyLeft, Mode: gb.KEY_REPEAT},
			Action: func() {
				ship.Components["motion"].(*gb.Motion).SetVelocity(0.1, 180)
			},
		},
		gb.KeyHandler{
			KeyPress: gb.KeyPress{Key: rl.KeyRight, Mode: gb.KEY_REPEAT},
			Action: func() {
				ship.Components["motion"].(*gb.Motion).SetVelocity(0.1, 0)
			},
		},
	)

	game.AddScene("myscene", ship)
	game.SetScene("myscene")

}

func main() {
	game.Run()
}
