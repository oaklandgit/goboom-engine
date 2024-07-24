package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SvgPath struct {
	GameObj     *GameObj
	Path        string
	StrokeWidth float32
	Color       rl.Color
}

func (*SvgPath) Id() string {
	return "svgPath"
}

func (o *GameObj) NewSvgPath(str string, w float32, c rl.Color) *GameObj {

	// default values that would otherwise be zero
	svg := &SvgPath{
		GameObj:     o,
		Path:        str,
		StrokeWidth: w,
		Color:       c,
	}

	o.AddComponents(svg)

	return o
}

func (p *SvgPath) Update() {}

func (p *SvgPath) Draw() {

	// const myPivot = "center"

	// type Pivot rl.Vector2

	// pivots := map[string]Pivot{
	// 	"top-left":     {0, 0},
	// 	"top-right":    {1, 0},
	// 	"bottom-left":  {0, 1},
	// 	"bottom-right": {1, 1},
	// 	"center":       {0.5, 0.5},
	// }

	rl.PushMatrix()

	// adjust to parent
	rl.Translatef(p.GameObj.Position.X, p.GameObj.Position.Y, 0)
	rl.Rotatef(p.GameObj.Angle, 0, 0, 1) // rotate on the z axis only
	rl.Scalef(p.GameObj.Scale.X, p.GameObj.Scale.Y, 1)

	// calculate size by drawing the shape with a transparent color
	transparent := rl.Color{R: 0, G: 0, B: 0, A: 0}
	shapeW, shapeH := DrawSVGPath(p.Path, p.StrokeWidth, transparent)

	// Shift the shape to the pivot point
	rl.Translatef((-shapeW * p.GameObj.Origin.X), (-shapeH * p.GameObj.Origin.Y), 0)

	// Draw the shape with the actual color
	DrawSVGPath(p.Path, p.StrokeWidth, p.Color)

	// Restore the drawing context (pop matrix)
	rl.PopMatrix()
}
