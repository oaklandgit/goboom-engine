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

const ROTATE_SPEED = 5

func init() {

	game.Reset = func() {}

	rotateCW := func(g *gb.GameObj) {
		g.Angle += ROTATE_SPEED
	}

	rotateCCW := func(g *gb.GameObj) {
		g.Angle -= ROTATE_SPEED
	}

	chassis := "M0 2 L0 2 L1 0 L1 0 L2 2"

	ship := game.NewGameObject("ship", gb.WithPosition(200, 200), gb.WithScale(10, 10), gb.WithAngle(0)).
		NewSvgPath(chassis, 0.2, rl.White).
		NewMotion(gb.WithVelocity(0, -90))

	ship.NewInput(
		gb.KeyHandler{
			KeyPress: gb.KeyPress{Key: rl.KeyLeft, Mode: gb.KEY_REPEAT},
			Action: func() {
				rotateCCW(ship)
			},
		},
		gb.KeyHandler{
			KeyPress: gb.KeyPress{Key: rl.KeyRight, Mode: gb.KEY_REPEAT},
			Action: func() {
				rotateCW(ship)
			},
		},
	)

	game.AddScene("myscene", ship)
	game.SetScene("myscene")

}

func main() {
	game.Run()
}
