package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

func draw() {
	rl.DrawCircleV(rl.Vector2{X: 0, Y: 0}, 20, rl.Red)
}

func getSize() rl.Vector2 {
	return rl.Vector2{X: 20, Y: 20}
}

func createThing() *gb.Node {
	return &gb.Node{
		Visible:  true,
		Scale:    rl.Vector2{X: 1, Y: 1},
		Position: rl.Vector2{X: 100, Y: 100},
		Alpha:    1,
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
	thing1 := createThing()
	thing2 := createThing()
	thing1.AddChildren(thing2)
	root.AddChildren(thing1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		root.RenderRoot()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
