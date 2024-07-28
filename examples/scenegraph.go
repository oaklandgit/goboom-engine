package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPlanet(x, y, r float32) *gb.Node {

	// raylib circles are drawn from the center
	// so we'll keep the origin at 0, 0
	const (
		originX = 0
		originY = 0
	)

	return &gb.Node{
		Visible:  true,
		Origin:   rl.Vector2{X: originX, Y: originY},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Position: rl.Vector2{X: x, Y: y},
		Rotation: 0,
		Alpha:    1,
		OnDraw: func(n *gb.Node) {
			rl.DrawCircleLines(0, 0, r, rl.Red)
		},
		GetBounds: func(n *gb.Node) rl.Rectangle {
			return rl.Rectangle{X: n.Position.X, Y: n.Position.Y, Width: r * 2, Height: r * 2}
		},
	}
}

func createSatellite(x, y float32) *gb.Node {

	const (
		w       float32 = 22
		h       float32 = 12
		originX         = 0.5
		originY         = 8
	)

	return &gb.Node{
		Visible:  true,
		Origin:   rl.Vector2{X: originX, Y: originY},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Position: rl.Vector2{X: x, Y: y},
		Rotation: 0,
		Alpha:    1,
		OnDraw: func(n *gb.Node) {
			// rl.DrawRectangleV(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: w, Y: h}, rl.Green)
			rl.DrawRectangleLines(0, 0, int32(w), int32(h), rl.White)
		},
		GetBounds: func(n *gb.Node) rl.Rectangle {
			return rl.Rectangle{X: n.Position.X, Y: n.Position.Y, Width: w, Height: h}
		},
	}
}

func main() {
	rl.InitWindow(600, 400, "SCENE GRAPH")
	rl.SetTargetFPS(60)

	root := gb.CreateRootNode(600, 400)

	planet := createPlanet(300, 200, 20)
	satellites := []*gb.Node{}

	for i := 0; i < 12; i += 1 {
		s := createSatellite(0, 0)
		s.Rotation = float32(i) * 30
		satellites = append(satellites, s)
	}

	planet.AddChildren(satellites...)

	// root.AddChildren(satellites...)
	root.AddChildren(planet)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		root.Render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
