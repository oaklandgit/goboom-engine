package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	solarSystem := CreateScene(
		"The Solar System",
		createPlanet("Mercury", 100, 100, 1, 0, 0.1, rl.Gray, 1),
		createPlanet("Venus", 200, 200, 1.2, 180, 0.2, rl.Yellow, 1),
		createPlanet("Earth", 300, 300, 1.3, 0, 0.3, rl.Blue, 1),
		createPlanet("Mars", 400, 400, 1.4, 180, 0.4, rl.Red, 1),
		createPlanet("Jupiter", 500, 500, 1.5, 0, 0.5, rl.Brown, 1),
		createPlanet("Saturn", 600, 600, 1.6, 180, 0.6, rl.Green, 1),
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		solarSystem.Update()
		solarSystem.Draw()
		
		rl.EndDrawing()
	}
	

}