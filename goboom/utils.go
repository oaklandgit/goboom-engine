package goboom

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

func PrintObjs (objs []*GameObj) {
	for _, o := range objs {
		fmt.Println(o.Name)
	}
}

func Displace(distance float32, angle float32) rl.Vector2 {
	rads := float64(angle * rl.Deg2rad)
	displacement := rl.Vector2{
    	X: distance * float32(math.Cos(rads)),
    	Y: distance * float32(math.Sin(rads)),
	}
	return displacement
}

func CalculateAngle(obj1, obj2 rl.Vector2) float32 {
    dx := obj1.X - obj2.X
    dy := obj1.Y - obj2.Y
	
    return float32(math.Atan2(float64(dy), float64(dx))) * 180 / math.Pi
}

// func adjustAngle(angle float32) float32 {
//     // Use modulo to wrap around
//     return float32(int(angle+360) % 360)
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
	height := int32(2)
	fontSize := int32(12)

	rl.DrawRectangle(x, y, w, height, rl.Fade(rl.Green, 0.25))
	rl.DrawRectangle(x, y, lineW, height, rl.Green)
	DrawText(text, x, y + fontSize, fontSize, 4, rl.Green, Left)
	
	
}