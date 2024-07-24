package main

import (
	gb "goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"SVG Path",
	600,
	800,
	true,
)

const (
	ROTATE_SPEED      = 5
	THRUST_SPEED      = 0.04
	MAX_SPEED         = 3
	SHIP_DRAG         = 0.999
	SHIP_WRAP_PADDING = 16
)

func init() {

	var speed float32

	game.Reset = func() {}

	rotateCW := func(g *gb.GameObj) {
		g.Angle += ROTATE_SPEED
	}

	rotateCCW := func(g *gb.GameObj) {
		g.Angle -= ROTATE_SPEED
	}

	thrust := func(g *gb.GameObj) {
		speed += THRUST_SPEED
		g.Components["motion"].(*gb.Motion).SetVelocity(speed, g.Angle-90)

	}

	score := game.NewGameObject("score",
		gb.WithPosition(float32(game.Width/2), 24),
		gb.WithScale(8, 8),
		gb.WithOrigin(0.5, 0.5)).
		NewVecText("SCORE: 0", 1, 1, rl.Green)

	ship := game.NewGameObject("ship",
		gb.WithPosition(200, 200),
		gb.WithScale(10, 10),
		gb.WithAngle(0)).
		NewSvgPath("M0 2 L0 2 L1 0 L1 0 L2 2", 0.2, rl.White).
		NewMotion(
			gb.WithFriction(SHIP_DRAG),
			gb.WithMaxVelocity(MAX_SPEED),
			gb.WithWrap(true, true, SHIP_WRAP_PADDING),
		)

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
		gb.KeyHandler{
			KeyPress: gb.KeyPress{Key: rl.KeyUp, Mode: gb.KEY_REPEAT},
			Action: func() {
				thrust(ship)
			},
		},
	)

	scene1 := game.NewGameObject("scene1")
	scene1.AddChildren(score, ship)
	scene1.Size = rl.NewVector2(game.Width, game.Height) // need to automate this.

	game.AddScene("myscene", scene1)
	game.SetScene("myscene")

	// ADD COLOR CYCLING JUICINESS
	colors := []rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple}
	index := 0

	changeAllTextColor := func() {
		targets := scene1.FindChildrenByComponent(true, "vecText")
		for _, target := range targets {
			target.Components["vecText"].(*gb.VecText).Color = colors[index]
		}
		index = (index + 1) % len(colors)
	}

	go gb.Cycle(changeAllTextColor, 500*time.Millisecond)

}

func main() {
	game.Run()
}
