package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	var level1Map = `
	..........
	..........
	...m......
	.........e
	..v.......
	..........
	`

	var level1MapTable = map[rune]func() *GameObj{
	'm': func() *GameObj {
		return createPlanet("Mars", 0, 0, 1, 0, 0.1, rl.Red, 1)
	},
	'e': func() *GameObj {
		return createPlanet("Earth", 0, 0, 2, 180, 0.3, rl.Blue, 1)
	},
	'v': func() *GameObj {
		return createPlanet("Venus", 0, 0, 3, 0, 0.2, rl.Yellow, 1)
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