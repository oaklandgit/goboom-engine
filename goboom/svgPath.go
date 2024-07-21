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
	rl.Translatef(p.GameObj.PosGlobal().X, p.GameObj.PosGlobal().Y, 0)
	DrawSVGPath(p.Path, p.StrokeWidth, p.Color)
	rl.PopMatrix()

}
