package goboom

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rotate struct {
	GameObj *GameObj
	Speed float32 // can be negative
}

func (*Rotate) Id() string {
	return "rotate"
}

type RotateOption func(*Rotate)

func (obj *GameObj) NewRotate(speed float32, opts ...RotateOption) *GameObj {

	rotate := &Rotate{
		GameObj: obj,
		Speed: speed,
	}

	for _, opt := range opts {
		opt(rotate)
	}

	obj.AddComponents(rotate)

	return obj
}

func (r *Rotate) Update() {

	r.GameObj.Angle += r.Speed
	if r.GameObj.Angle > 360 {
		r.GameObj.Angle = 0
	}
	if r.GameObj.Angle < 0 {
		r.GameObj.Angle = 360
	}
	
}

func (r *Rotate) Draw() {

	if !r.GameObj.Game.Debug {
		return
	}

	p1 := r.GameObj.PosGlobal()
	radius := r.GameObj.Width() / 2
	rads := float64(r.GameObj.Angle * rl.Deg2rad)

	p2 := rl.Vector2{
		X: p1.X + radius * float32(math.Cos(rads)),
		Y: p1.Y + radius * float32(math.Sin(rads)),
	} 

	rl.DrawLine(
		int32(p1.X),
		int32(p1.Y),
		int32(p2.X), int32(p2.Y),
		rl.White)

	rl.DrawText(
		fmt.Sprintf("%.0f", r.GameObj.Angle),
		int32(p1.X + 10),
		int32(p1.Y + 10),
		20, rl.White)
}

