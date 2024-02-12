package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	planets := []*GameObj{
		createPlanet(
			"Planet Claire",
			400, 200,
			1,
			0,
			0.4,
			rl.Blue,
			1,
		),
		createPlanet(
			"Prime Beta Z",
			100, 300,
			2,
			0,
			0.2,
			rl.Red,
			0.2,
		),
	
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, p := range planets {

			for _, comp := range p.Components {
				comp.Draw()
				comp.Update()
			}

		}
		
		rl.EndDrawing()
	}
	

}