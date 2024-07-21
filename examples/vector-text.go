package main

import (
	gb "goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 600
const STARFIELD_DENSITY = 100

var game = gb.NewGame(
	"Vector Text",
	WIDTH,
	HEIGHT,
	true,
)

func init() {

	game.Reset = func() {}

	letters := game.NewGameObject("letters", gb.WithPosition(4, 0), gb.WithScale(8, 8)).
		NewVecText("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, 0.5, rl.Blue)

	numbers := game.NewGameObject("numbers", gb.WithPosition(4, 40), gb.WithScale(8, 8)).
		NewVecText("0123456789", 1, 0.5, rl.Yellow)

	weights1 := game.NewGameObject("weights1", gb.WithPosition(4, 80), gb.WithScale(8, 8)).
		NewVecText("LINE WEIGHTS", 2, 0.5, rl.Yellow)
	weights2 := game.NewGameObject("weights2", gb.WithPosition(4, 120), gb.WithScale(8, 8)).
		NewVecText("LINE WEIGHTS", 3, 0.5, rl.Green)

	spacing1 := game.NewGameObject("spacing1", gb.WithPosition(4, 160), gb.WithScale(8, 8)).
		NewVecText("LETTER SPACING", 1, 2, rl.Blue)

	spacing2 := game.NewGameObject("spacing2", gb.WithPosition(4, 200), gb.WithScale(8, 8)).
		NewVecText("LETTER SPACING", 1, 4, rl.Blue)

	scaling := game.NewGameObject("scaling", gb.WithPosition(4, 260), gb.WithScale(24, 24)).
		NewVecText("SCALING", 1, 1, rl.Blue)

	squishing := game.NewGameObject("squishing", gb.WithPosition(4, 320), gb.WithScale(6, 12)).
		NewVecText("SQUISHING", 1, 1, rl.Blue)

	stretching := game.NewGameObject("stretching", gb.WithPosition(4, 370), gb.WithScale(24, 8)).
		NewVecText("STRETCHING", 1, 0.5, rl.Blue)

	scene := game.NewGameObject("text examples")

	game.AddScene("myscene", scene)
	scene.AddChildren(letters, numbers, weights1, weights2, spacing1, spacing2, scaling, squishing, stretching)

	game.SetScene("myscene")

	colors := []rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple}
	index := 0

	changeAllTextColor := func() {
		targets := scene.FindChildrenByComponent(true, "vecText")
		for _, target := range targets {
			target.Components["vecText"].(*gb.VecText).Color = colors[index]
		}
		index = (index + 1) % len(colors)
	}

	go gb.Cycle(changeAllTextColor, 1*time.Second)

}

func main() {
	game.Run()
}
