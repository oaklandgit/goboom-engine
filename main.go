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
	..........
	🪐.........
	.........🌎
	..🔴.......
	..........
	`

	ship := createShip(400, 120)

	earth := createPlanet("Earth", 0, 0, 0.2, 0.1, 0, 0.3, rl.Blue, 0.4, ship, 1)
	earth.AddChildren(
		createMoon("Moon", 0.4, 0.4, 0.08, 112, rl.White, 0.5),
	)

	mars := createPlanet("Mars", 0, 0, 0.3, 0.6, 0, 0.2, rl.Red, 0.3, ship, 1)
	mars.AddChildren(
		createMoon("Phobos", -1.3, 3, 0.04, 80, rl.Brown, 1),
		createMoon("Deimos", -1, 0.1, 0.02, 62, rl.Gray, 1),
	)

		var level1MapTable = map[rune]func() *GameObj{
		'🪐': func() *GameObj {
			return createPlanet("Saturn", 0, 0, -0.1, 0.2, 0, 0.1, rl.Yellow, 0.2, ship, 1)
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
		level1Map,
		60, 60,
		level1MapTable,
	)

	solarSystem.Size = rl.NewVector2(screenW, screenH)

	// solarSystem.Size = rl.NewVector2(screenW, screenH)

	starfield := NewGameObject("Starfield")
	starfield.NewStarfield(screenW, screenH, 40)

	scene1 := NewGameObject("Scene 1")
	scene1.Size = rl.NewVector2(screenW, screenH)
	scene1.AddChildren(starfield, solarSystem, ship)
	
	game.AddScene("level1", scene1)
	game.Run("level1")


}