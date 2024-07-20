package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegPoly struct {
	GameObj     *GameObj
	Sides       int32
	Radius      float32
	StrokeColor rl.Color
	FillColor   rl.Color
	StrokeWidth float32
	Opacity     float32
	Rotation    float32
}

func (o *GameObj) NewRegPoly(s int32, r float32, c rl.Color) *GameObj {

	regPoly := &RegPoly{
		GameObj:     o,
		Sides:       s,
		Radius:      r,
		Opacity:     1.0,
		FillColor:   c,
		StrokeColor: rl.White,
		StrokeWidth: 1.0,
		Rotation:    0.0,
	}

	o.AddComponents(regPoly)

	return o
}

func (*RegPoly) Id() string {
	return "regularPolygon"
}

func (p *RegPoly) Update() {
	// no op
}

func (p *RegPoly) Draw() {
	rl.DrawPoly(p.GameObj.Position, p.Sides, p.Radius, p.Rotation, p.FillColor)
	rl.DrawPolyLinesEx(p.GameObj.Position, p.Sides, p.Radius, p.Rotation, p.StrokeWidth, p.StrokeColor)
}
