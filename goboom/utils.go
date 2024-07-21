package goboom

import (
	"fmt"
	"math"
	"reflect"
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

func Cycle(callback func(), delay time.Duration) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			callback()
		}
	}
}

func CycleByInterface(scene *GameObj, iface interface{}, array []rl.Color, delay time.Duration) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	index := 0
	ifaceType := reflect.TypeOf(iface).Elem()

	for {
		select {
		case <-ticker.C:
			for _, child := range scene.Children {
				for _, component := range child.Components {
					// Check if the component implements the given interface
					componentValue := reflect.ValueOf(component)
					if componentValue.Type().Implements(ifaceType) {
						// Call the ChangeColor method using reflection
						method := componentValue.MethodByName("ChangeColor")
						if method.IsValid() {
							method.Call([]reflect.Value{reflect.ValueOf(array[index])})
						}
					}
				}
			}
			index = (index + 1) % len(array)
		}
	}
}

func PrintObjs(objs []*GameObj) {
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
	DrawText(text, x, y+fontSize, fontSize, 4, rl.Green, Left)

}
