package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	var level1Map = `
	..........
	....🚀.....
	🪐.........
	.........🌎
	..🔴.......
	..........
	`

	earth := createPlanet("Earth", 0, 0, 2, 180, 0.3, rl.Blue, 1)
	moon := createMoon("Moon", earth, 1.3, 0, 0.08, 112, rl.White, 0.5)
	earth.AddChildren(moon)

	mars := createPlanet("Mars", 0, 0, 3, 0, 0.2, rl.Red, 1)
	phobos := createMoon("Phobos", mars, -1.3, 45, 0.04, 80, rl.Brown, 0.5)
	mars.AddChildren(phobos)

	var level1MapTable = map[rune]func() *GameObj{
	'🚀': func() *GameObj {
		return createShip(0, 0)
	},
	'🪐': func() *GameObj {
		return createPlanet("Saturn", 0, 0, 1, 0, 0.1, rl.Yellow, 1)
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

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		solarSystem.Update()
		solarSystem.Draw()
		
		rl.EndDrawing()
	}
	

}