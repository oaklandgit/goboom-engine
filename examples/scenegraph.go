package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createThing() *gb.Node {
	return &gb.Node{
		Visible:  true,
		Origin:   rl.Vector2{X: 0, Y: 0},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Position: rl.Vector2{X: 0, Y: 0},
		Rotation: 0,
		Alpha:    1,
		DrawFunc: gb.Drawable{
			Draw: func() {
				// rl.DrawCircleV(rl.Vector2{X: 0, Y: 0}, 20, rl.Red)
				rl.DrawRectangleV(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 50, Y: 25}, rl.Green)

			},
			GetSize: func() rl.Vector2 {
				return rl.Vector2{X: 50, Y: 25}
			},
		},
	}
}

// func createBarrier(p rl.Vector2, r float32) *gb.Node {
// 	return &gb.Node{
// 		Visible:  true,
// 		Scale:    rl.Vector2{X: 1, Y: 1},
// 		Position: p,
// 		Alpha:    1,
// 		Rotation: r,
// 		DrawFunc: gb.Drawable{
// 			Draw: func() {
// 				rl.DrawRectangleV(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 12, Y: 50}, rl.White)
// 			},
// 			GetSize: func() rl.Vector2 {
// 				return rl.Vector2{X: 12, Y: 50}
// 			},
// 		},
// 	}
// }

// func createBlockade() *gb.Node {
// 	blockade := &gb.Node{
// 		Visible:  true,
// 		Scale:    rl.Vector2{X: 1, Y: 1},
// 		Position: rl.Vector2{X: 300, Y: 200},
// 		Alpha:    1,
// 		DrawFunc: gb.Drawable{
// 			Draw:    func() {},
// 			GetSize: func() rl.Vector2 { return rl.Vector2{X: 0, Y: 0} },
// 		},
// 	}

// 	fudgeAngle := float32(14) // this is a hack
// 	numChildren := 10
// 	radius := 100.0 // Adjust the radius as needed
// 	angleIncrement := 2 * math.Pi / float64(numChildren)

// 	for i := 0; i < numChildren; i++ {
// 		angle := float64(i) * angleIncrement
// 		x := radius * math.Cos(angle)
// 		y := radius * math.Sin(angle)
// 		blockade.AddChildren(createBarrier(rl.Vector2{X: float32(x), Y: float32(y)}, float32(angle*180/math.Pi)+fudgeAngle))
// 	}

// 	return blockade
// }

func main() {
	rl.InitWindow(600, 400, "SCENE GRAPH")
	rl.SetTargetFPS(60)

	root := gb.CreateRootNode()
	// blockade := createBlockade()
	// block := createBarrier(rl.Vector2{X: 0, Y: 0}, 0)
	thing := createThing()
	root.AddChildren(thing)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		root.Render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
