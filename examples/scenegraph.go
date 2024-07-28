package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPlanet(x, y, r float32, c rl.Color) *gb.Node {

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
			// rl.DrawCircleLines(0, 0, r, rl.Red)
			rl.DrawCircle(0, 0, r, c)
		},
		GetBounds: func(n *gb.Node) rl.Rectangle {
			return rl.Rectangle{X: n.Position.X, Y: n.Position.Y, Width: r * 2, Height: r * 2}
		},
	}
}

func createSatellite(w, h float32, dist float32) *gb.Node {

	return &gb.Node{
		Visible:  true,
		Origin:   rl.Vector2{X: 0.5, Y: h * dist / 2},
		Scale:    rl.Vector2{X: 1, Y: 1},
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

func createSatellites(num int, dist float32, speed float32) *gb.Node {
	satellites := []*gb.Node{}

	for i := 0; i < num; i += 1 {
		s := createSatellite(20, 5, dist)
		s.Rotation = float32(i) * 360 / float32(num)
		satellites = append(satellites, s)
	}

	group := &gb.Node{
		Visible: true,
		Scale:   rl.Vector2{X: 1, Y: 1},
		OnUpdate: func(n *gb.Node) {
			n.Rotation += speed * rl.GetFrameTime()
		},
	}

	group.AddChildren(satellites...)

	return group

}

func Update(n *gb.Node) {

	for _, c := range n.Children {
		if c.OnUpdate != nil {
			c.OnUpdate(c)
		}
		Update(c)
	}

}

func main() {
	rl.InitWindow(600, 400, "SCENE GRAPH")
	rl.SetTargetFPS(60)

	root := gb.CreateRootNode(600, 400, rl.DarkBlue)

	planet1 := createPlanet(300, 200, 50, rl.Red).AddChildren(createSatellites(8, 10, 100))
	planet2 := createPlanet(75, 75, 16, rl.Green).AddChildren(createSatellites(12, 3, 30))

	root.AddChildren(planet1, planet2)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		Update(root)
		root.Render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
