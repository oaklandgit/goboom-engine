package main

import (
	gb "goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"Vector Text",
	800,
	600,
	true,
)

func init() {

	game.Reset = func() {}

	center := float32(rl.GetScreenWidth()) / 2

	letters := game.NewGameObject("letters",
		gb.WithPosition(center, 40),
		gb.WithOrigin(0.5, 0.5),
		gb.WithScale(8, 8)).
		NewVecText("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0.5, 1, rl.Blue)

	numbers := game.NewGameObject("numbers",
		gb.WithPosition(center, 80),
		gb.WithOrigin(0.5, 0.5),
		gb.WithScale(8, 8)).
		NewVecText("0123456789", 1, 0.5, rl.Yellow)

	weights1 := game.NewGameObject("weights1",
		gb.WithPosition(center, 120),
		gb.WithOrigin(0.5, 0.5),
		gb.WithScale(8, 8)).
		NewVecText("LINE", 2, 0.5, rl.Yellow)

	weights2 := game.NewGameObject("weights2",
		gb.WithPosition(center, 160),
		gb.WithOrigin(0.5, 0.5),
		gb.WithScale(8, 8)).
		NewVecText("WEIGHTS", 3, 0.5, rl.Green)

	spacing1 := game.NewGameObject("spacing1",
		gb.WithPosition(center, 200),
		gb.WithScale(8, 8)).
		NewVecText("LETTER SPACING", 1, 2, rl.Blue)

	spacing2 := game.NewGameObject("spacing2",
		gb.WithPosition(center, 240),
		gb.WithScale(8, 8)).
		NewVecText("LETTER SPACING", 1, 4, rl.Blue)

	scaling := game.NewGameObject("scaling",
		gb.WithPosition(center, 300),
		gb.WithScale(24, 24)).
		NewVecText("SCALING", 1, 1, rl.Blue)

	squishing := game.NewGameObject("squishing",
		gb.WithPosition(center, 360),
		gb.WithScale(6, 12)).
		NewVecText("SQUISHING", 1, 1, rl.Blue)

	stretching := game.NewGameObject("stretching",
		gb.WithPosition(center, 420),
		gb.WithScale(24, 8)).
		NewVecText("STRETCHING", 1, 0.5, rl.Blue)

	scene := game.NewGameObject("text examples")

	game.AddScene("myscene", scene)
	scene.AddChildren(
		letters,
		numbers,
		weights1,
		weights2,
		spacing1,
		spacing2,
		scaling,
		squishing,
		stretching)

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
