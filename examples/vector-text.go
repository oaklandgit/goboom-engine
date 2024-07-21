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

	greeting := game.NewGameObject("greeting", gb.WithPosition(WIDTH/2, HEIGHT/2), gb.WithScale(8, 8)).
		NewVecText("HELLO WORLD", 1, rl.Yellow, gb.WithAlignment(gb.TextCenter))

	score := game.NewGameObject("score", gb.WithPosition(WIDTH/2, 20), gb.WithScale(6, 6)).
		NewVecText("123456", 1, rl.Yellow, gb.WithAlignment(gb.TextCenter))

	scene := game.NewGameObject("intro")

	game.AddScene("myscene", scene)
	scene.AddChildren(greeting, score)

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

	// go gb.Cycle(changeTextColor, 1*time.Second)
	go gb.Cycle(changeAllTextColor, 1*time.Second)

}

func main() {
	game.Run()
}
