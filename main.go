package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 450, "SpaceMiner")
	rl.SetTargetFPS(60)

	planet := NewGameObject("Planet",
		WithPosition(400, 200),
		WithScale(0.4, 0.4),
	)
	planet.NewSprite(
		rl.LoadTexture("assets/planet.png"),
		WithColor(rl.Blue),
		WithOpacity(0.5))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, comp := range planet.Components {
			comp.Draw()
			comp.Update()
		}
		
		rl.EndDrawing()
	}
	

}