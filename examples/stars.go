package main

import (
	gb "goboom"
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

	starfieldBottom := game.NewGameObject("starfield1Bottom", gb.WithPosition(0, 0)).
		NewMotion(gb.WithVelocity(0.2, 90), gb.WithWrap(false, true, game.Height))
	starfieldOffscreen := game.NewGameObject("starfield1Offscreen", gb.WithPosition(0, game.Height)).
		NewMotion(gb.WithVelocity(0.2, 90), gb.WithWrap(false, true, game.Height))

	gb.NewStarfield(starfieldBottom, int(game.Width), int(game.Height), STARFIELD_DENSITY)
	gb.NewStarfield(starfieldOffscreen, int(game.Width), int(game.Height), STARFIELD_DENSITY)

	scene := game.NewGameObject("intro")

	game.AddScene("myscene", scene)
	scene.AddChildren(starfieldBottom, starfieldOffscreen)

	game.SetScene("myscene")

}

func main() {
	game.Run()
}
