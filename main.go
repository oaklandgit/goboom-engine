package main

import rl "github.com/gen2brain/raylib-go/raylib"

var textures map[string]rl.Texture2D

const (
	screenW = 800
	screenH = 450
	title = "Space Miner!"
)

var game = NewGame(title, screenW, screenH, textures)

func main() {

	rl.InitWindow(screenW, screenH, title)
	rl.SetTargetFPS(60)
	textures = LoadTextures(
		"assets/planet.png",
		"assets/ship.png",
	)

	var level1Map = `
	..........
	....🚀.....
	🪐.........
	.........🌎
	..🔴.......
	..........
	`

	earth := createPlanet("Earth", 0, 0, 1, 0.1, 180, 0.3, rl.Blue, 1)
	earth.AddChildren(
		createMoon("Moon", 0.4, 0.4, 0.08, 112, rl.White, 0.5),
	)

	mars := createPlanet("Mars", 0, 0, 2, 0.6, 0, 0.2, rl.Red, 1)
	mars.AddChildren(
		createMoon("Phobos", -1.3, 3, 0.04, 80, rl.Brown, 0.5),
	)

		var level1MapTable = map[rune]func() *GameObj{
		'🚀': func() *GameObj {
			return createShip(0, 0)
		},
		'🪐': func() *GameObj {
			return createPlanet("Saturn", 0, 0, -1.5, 0.2, 0, 0.1, rl.Yellow, 1)
		},
		'🌎': func() *GameObj {
			return earth
		},
		'🔴': func() *GameObj {
			return mars
		},

	}

	solarSystem := CreateLevel(
		"The Solar System",
		rl.Vector2{X: screenW, Y: screenH},
		level1Map,
		60, 60,
		level1MapTable,
	)

	solarSystem.Size = rl.NewVector2(screenW, screenH)
	
	game.AddScene("solarSystem", solarSystem)
	game.Run("solarSystem")


}