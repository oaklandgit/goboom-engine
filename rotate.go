package main

import (
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

func (o *GameObj) NewRotate(speed float32, opts ...RotateOption) *Rotate {

	rotate := &Rotate{
		GameObj: o,
		Speed: speed,
	}

	for _, opt := range opts {
		opt(rotate)
	}

	o.AddComponents(rotate)

	return rotate
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
}

