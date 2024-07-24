package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const STROKE_WEIGHT_ADJUST = 0.02

type VecText struct {
	GameObj *GameObj
	Text    string
	Weight  float32
	Gap     float32
	Color   rl.Color
}

func (o *GameObj) NewVecText(str string, weight float32, gap float32, c rl.Color, opts ...VecTextOption) *GameObj {

	vecText := &VecText{
		GameObj: o,
		Text:    str,
		Weight:  weight,
		Gap:     gap,
		Color:   c,
	}

	for _, opt := range opts {
		opt(vecText)
	}

	o.AddComponents(vecText)

	return o
}

type VecTextOption func(*VecText)

func (*VecText) Id() string {
	return "vecText"
}

func (vt *VecText) Update() {
	// no op
}

func (vt *VecText) ChangeColor(c rl.Color) {
	vt.Color = c
}

func (vt *VecText) Draw() {

	// CALCULATE LINE WEIGHT
	scaleAvg := (vt.GameObj.Scale.X + vt.GameObj.Scale.Y) / 2
	calculatedWeight := vt.Weight / scaleAvg

	// CALCULATE SIZE
	transparent := rl.Color{R: 0, G: 0, B: 0, A: 0}
	var totalW, totalH float32
	for _, char := range vt.Text {
		shapeW, shapeH := DrawSVGPath(letterForms[string(char)], calculatedWeight, transparent)
		totalW += shapeW + vt.Gap
		totalH = shapeH
	}

	// NOW DRAW
	rl.PushMatrix()

	// adjust to parent
	rl.Translatef(vt.GameObj.Position.X, vt.GameObj.Position.Y, 0)
	rl.Rotatef(vt.GameObj.Angle, 0, 0, 1) // rotate on the z axis only
	rl.Scalef(vt.GameObj.Scale.X, vt.GameObj.Scale.Y, 1)

	/////// END INNER LOOP

	// Shift the shape to the pivot point
	rl.Translatef((-totalW * vt.GameObj.Origin.X), (-totalH * vt.GameObj.Origin.Y), 0)

	// Draw all the letters again with the actual color
	for i, char := range vt.Text {
		letterOffsetX := float32(i) * (2 + vt.Gap) // 2 is the base width of the letter
		rl.PushMatrix()
		rl.Translatef(letterOffsetX, 0, 0)
		DrawSVGPath(letterForms[string(char)], calculatedWeight, vt.Color)
		rl.PopMatrix()
	}

	////////

	// Restore the drawing context (pop matrix)
	rl.PopMatrix()

}

var letterForms = map[string]string{
	"A": "M0 2 L0 1 L1 0 L2 1 L2 2 M0 1 L2 1",
	"B": "M0 2 L0 0 L1 0 L1 1 L2 1 L2 2 L0 2 M0 1 L2 1",
	"C": "M2 2 L0 2 L0 0 L2 0",
	"D": "M0 2 L0 0 L1 0 L2 1 L2 2 L0 2",
	"E": "M2 2 L0 2 L0 0 L2 0 M0 1 L1 1",
	"F": "M0 2 L0 0 L2 0 M0 1 L1 1",
	"G": "M2 0 L0 0 L0 2 L2 2 L2 1 L1 1",
	"H": "M0 0 L0 2 M2 0 L2 2 M0 1 L2 1",
	"I": "M1 0 L1 2 M0 0 L2 0 M0 2 L2 2",
	"J": "M2 0 L2 2 L0 2 L0 1",
	"K": "M0 0 L0 2 M2 0 L0 1 L2 2",
	"L": "M0 0 L0 2 L2 2",
	"M": "M0 2 L0 0 L1 1 L2 0 L2 2",
	"N": "M0 2 L0 0 L2 2 L2 0",
	"O": "M0 0 L0 2 L2 2 L2 0 L0 0",
	"P": "M0 2 L0 0 L2 0 L2 1 L0 1",
	"Q": "M0 0 L0 2 L2 2 L2 0 L0 0 M1 1 L2.5 2.5", // q is a little different
	"R": "M0 2 L0 0 L2 0 L2 1 L0 1 M1 1 L2 2",
	"S": "M2 0 L0 0 L0 1 L2 1 L2 2 L0 2",
	"T": "M0 0 L2 0 M1 0 L1 2",
	"U": "M0 0 L0 2 L2 2 L2 0",
	"V": "M0 0 L1 2 L2 0",
	"W": "M0 0 L0 2 L1 1 L2 2 L2 0",
	"X": "M0 0 L2 2 M2 0 L0 2",
	"Y": "M0 0 L1 1 L2 0 M1 1 L1 2",
	"Z": "M0 0 L2 0 L0 2 L2 2",
	"0": "M0 0 L0 2 L2 2 L2 0 L0 0 M0 2 L2 0",
	"1": "M0 0 L1 0 L1 2 M0 2 L2 2",
	"2": "M0 0 L2 0 L2 1 L0 1 L0 2 L2 2",
	"3": "M0 0 L2 0 L0 1 L2 1 L2 2 L0 2",
	"4": "M0 0 L0 1 L2 1 M2 0 L2 2 M0 1 L2,1",
	"5": "M2 0 L0 0 L0 1 L2 1 L2 2 L0 2",
	"6": "M2 0 L0 0 L0 2 L2 2 L2 1 L0 1",
	"7": "M0 0 L2 0 L0 2",
	"8": "M0 0 L0 2 L2 2 L2 0 L0 0 M0 1 L2 1",
	"9": "M0 2 L2 2 L2 0 L0 0 L0 1 L2 1",
	" ": "",
}
