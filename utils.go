package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Alignment int

const (
	Left Alignment = iota
	Center
)

func WaitAndTrigger(callback func(), delay time.Duration) {
    time.AfterFunc(delay, callback)
}

func printObjs (objs []*GameObj) {
	for _, o := range objs {
		fmt.Println(o.Name)
	}
}

func calculateAngle(obj1, obj2 rl.Vector2) float32 {
    dx := obj1.X - obj2.X
    dy := obj1.Y - obj2.Y
	
    return float32(math.Atan2(float64(dy), float64(dx))) * 180 / math.Pi
}

// func calculateAngle(targetPos, parentPos rl.Vector2) float32 {
//     dx := targetPos.X - parentPos.X
//     dy := targetPos.Y - parentPos.Y
//     return float32(math.Atan2(float64(dy), float64(dx))) * 180 / math.Pi
// }

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