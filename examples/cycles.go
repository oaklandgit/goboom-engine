package main

import (
	gb "goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game = gb.NewGame(
	"Vector Text",
	600,
	400,
	true,
)

func init() {

	game.Reset = func() {}

	// create text objects
	greeting := game.NewGameObject("greeting", gb.WithPosition(game.Width/2, 60), gb.WithScale(8, 8), gb.WithTags("colorme")).
		NewVecText("HELLO WORLD", 1, 2, rl.Yellow)

	shape := game.NewGameObject("shape", gb.WithPosition(game.Width/2, 160), gb.WithTags("colorme")).
		NewRegPoly(11, 40, rl.Red)

	// add text objects to scene
	scene := game.NewGameObject("intro")
	game.AddScene("myscene", scene)
	scene.AddChildren(greeting, shape)
	game.SetScene("myscene")

	// set up cycling effect
	colors := []rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple}
	index := 0

	// set up a callback function to change the color of the text
	// a la the 80s classic "Stargate" by Williams Electronics,
	// colors of various elements cycle through a set of colors
	// in sync with each other
	//

	type ColorChanger interface {
		ChangeColor(color rl.Color)
	}

	changeAllColors := func() {
		targets := scene.FindChildrenByTags(true, "colorme")
		for _, target := range targets {
			for _, component := range target.Components {
				if colorChanger, ok := component.(ColorChanger); ok {
					colorChanger.ChangeColor(colors[index])
				}
			}
		}
		index = (index + 1) % len(colors)
	}

	// start the cycling effect
	go gb.Cycle(changeAllColors, 500*time.Millisecond)

}

func main() {
	game.Run()
}
