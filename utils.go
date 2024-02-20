package main

import rl "github.com/gen2brain/raylib-go/raylib"

// const FONT = "assets/VT323-Regular.ttf"

type Alignment int

const (
	Left Alignment = iota
	Center
	// Right
)

func DrawText(
	text string,
	x int32,
	y int32,
	fontSize int32,
	spacing float32,
	color rl.Color,
	alignment Alignment) {

	if alignment == Center {
		x = x - int32(rl.MeasureText(text, fontSize)/2)
	}
	
	rl.DrawTextEx(
		rl.GetFontDefault(),
		text,
		rl.Vector2{X: float32(x), Y: float32(y)},
		float32(fontSize),
		spacing,
		color)
	
}

func DrawProgressBar(
	x int32,
	y int32,
	w int32,
	completed int32,
	total int32,
	text string) {

	lineW := int32(float32(w) * (float32(completed) / float32(total)))
	height := int32(6)
	fontSize := int32(13)

	rl.DrawRectangle(x, y, w, height, rl.White)
	rl.DrawRectangle(x, y, lineW, height, rl.Green)
	DrawText(text, x, y + fontSize, fontSize, 4, rl.White, Left)
	
	
}