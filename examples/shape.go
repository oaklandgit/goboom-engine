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

	shape := "L10 0 L20 10 L60 10 L80 20 L80 30 L50 30 L30 40 L10 40 L20 30 L10 30 L10 20 Z"

	ship := game.NewGameObject("ship",
		gb.WithPosition(200, 200),
		gb.WithOrigin(0.5, 0.5), // pivot around center
		gb.WithScale(1, 1),
		gb.WithAngle(0)).
		NewSvgPath(shape, 1, rl.Pink).
		NewRotate(1)

	game.AddScene("myscene", ship)
	game.SetScene("myscene")

}

func main() {
	game.Run()
}
