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
	rl.PushMatrix()

	// Translate to the parent's position
	rl.Translatef(p.GameObj.Position.X, p.GameObj.Position.Y, 0)

	// Scale (including negative values for flipping) <-- negative doesn't work for some reason
	rl.Scalef(p.GameObj.Scale.X, p.GameObj.Scale.Y, 1)

	// Rotate around the center of the shape
	rl.Rotatef(p.GameObj.Angle, 0, 0, 1) // rotate on the z axis only

	// DRAW A BLANK TO CALCULATE THE SIZE
	transparent := rl.Color{R: 0, G: 0, B: 0, A: 0}
	shapeW, shapeH := DrawSVGPath(p.Path, p.StrokeWidth, transparent)

	// Shift the shape to the center
	rl.Translatef(-shapeW/2, -shapeH/2, 0)

	// Draw the shape with the actual color
	DrawSVGPath(p.Path, p.StrokeWidth, p.Color)

	// Restore the drawing context (pop matrix)
	rl.PopMatrix()
}
