package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"SVG Path",
	600,
	800,
	true,
)

func init() {

	game.Reset = func() {}

	// shape := "L10 0 L20 10 L60 10 L80 20 L80 30 L50 30 L30 40 L10 40 L20 30 L10 30 L10 20 Z"
	chassis := "M0 2 L0 2 L1 0 L1 0 L2 2"
	// flame := "M1 2 L1 3"

	ship := game.NewGameObject("ship", gb.WithPosition(200, 200), gb.WithScale(10, 10), gb.WithAngle(0)).
		NewSvgPath(chassis, 0.2, rl.Yellow).
		// NewSvgPath(flame, 0.2, rl.Red).
		NewRotate(1)

	game.AddScene("myscene", ship)
	game.SetScene("myscene")

}

func main() {
	game.Run()
}
