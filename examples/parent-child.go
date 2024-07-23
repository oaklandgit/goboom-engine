package main

import (
	gb "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth   = 256
	screenHeight  = 256
	rectangleSize = 128
	scaleFactor   = 0.2
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib example")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// START CONTEXT
		rl.PushMatrix()
		// make centerpoint of the screen the origin
		rl.Translatef(screenWidth/2, screenHeight/2, 0)

		// DO THE TRANSFORMATIONS
		// e.g. tranlate, rotate, or scale
		rl.Rotatef(30, 0, 0, 1) // rotate on the z axis if

		// DRAW THE THING
		// shift the thing half its size to the left and up
		shape := "L10 0 L20 10 L60 10 L80 20 L80 30 L50 30 L30 40 L10 40 L20 30 L10 30 L10 20 Z"

		// DRAW A BLANK TO CALCULATE THE SIZE
		shapeW, shapeH := gb.DrawSVGPath(shape, 1, rl.Color{R: 0, G: 0, B: 0, A: 0})

		// NOW SHIFT THE SHAPE TO THE CENTER
		rl.Translatef(-shapeW/2, -shapeH/2, 0)
		// AND DRAW FOR REALS
		gb.DrawSVGPath(shape, 1, rl.Black)
		// rl.DrawRectangleLines(-rectangleSize/2, -rectangleSize/2, rectangleSize, rectangleSize, rl.Black)
		// rl.DrawRectangle(-rectangleSize/2, -rectangleSize/2, rectangleSize, rectangleSize, rl.Color{R: 128, G: 128, B: 128, A: 128})

		// Restore the drawing context (pop matrix)
		rl.PopMatrix()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
