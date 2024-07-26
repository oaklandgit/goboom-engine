package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

func draw() {
	rl.DrawCircleV(rl.Vector2{X: 100, Y: 100}, 20, rl.Red)
}

func getSize() rl.Vector2 {
	return rl.Vector2{X: 500, Y: 500}
}

func createThingNode() *gb.Node {
	return &gb.Node{
		Visible: true,
		Scale:   rl.Vector2{X: 1, Y: 1},
		Alpha:   1,
		DrawFunc: gb.Drawable{
			Draw:    draw,
			GetSize: getSize,
		},
	}
}

func main() {
	rl.InitWindow(600, 400, "SCENE GRAPH")
	rl.SetTargetFPS(60)

	root := gb.CreateRootNode()
	root.AddChildren(createThingNode())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		root.Render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
