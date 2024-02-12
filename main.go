package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	var level1Map = `
	..........
	..........
	🪐.........
	.........🌎
	..🔴.......
	..........
	`

	var level1MapTable = map[rune]func() *GameObj{
	'🪐': func() *GameObj {
		return createPlanet("Saturn", 0, 0, 1, 0, 0.1, rl.Yellow, 1)
	},
	'🌎': func() *GameObj {
		return createPlanet("Earth", 0, 0, 2, 180, 0.3, rl.Blue, 1)
	},
	'🔴': func() *GameObj {
		return createPlanet("Mars", 0, 0, 3, 0, 0.2, rl.Red, 1)
	},
}

	solarSystem := CreateLevel(
		"The Solar System",
		level1Map,
		60, 60,
		level1MapTable,
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		solarSystem.Update()
		solarSystem.Draw()
		
		rl.EndDrawing()
	}
	

}