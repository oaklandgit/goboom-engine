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

	rl.PushMatrix()

	// translate to center of screen
	rl.Translatef(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2, 0)
	// rl.Translatef(vt.GameObj.PosGlobal().X, vt.GameObj.PosGlobal().Y, 0)

	rl.Scalef(vt.GameObj.Scale.X, vt.GameObj.Scale.Y, 1)

	var shapeW, shapeH float32

	for _, char := range vt.Text {

		// letterOffsetX := float32(i) * (2 + vt.Gap)
		// fullWidth += int(letterOffsetX)

		scaleAvg := (vt.GameObj.Scale.X + vt.GameObj.Scale.Y) / 2
		weight := vt.Weight / scaleAvg

		rl.PushMatrix()

		// rl.Translatef(letterOffsetX, 0, 0)

		// FIRST, CALCULATE THE SIZE OF THE LETTER
		// AND ADD IT TO THE TOTAL WIDTH
		transparent := rl.Color{R: 0, G: 0, B: 0, A: 0}
		w, h := DrawSVGPath(letterForms[string(char)], weight, transparent)
		shapeW += w + vt.Gap
		shapeH = h // tho all letters are the same height

		// Shift the shape to the center
		rl.Translatef(-shapeW/2, -shapeH/2, 0)

		// Draw the shape with the actual color
		DrawSVGPath(letterForms[string(char)], weight, vt.Color)
		rl.PopMatrix()

	}

	// rl.Translatef(-float32(fullWidth), 0, 0)
	rl.PopMatrix()

	// if vt.GameObj.Game.Debug {

	// 	width := float32(len(vt.Text)) * (2 + vt.Gap) * vt.GameObj.Scale.X
	// 	height := 2 * vt.GameObj.Scale.Y
	// 	centerX := vt.GameObj.PosGlobal().X + width/2
	// 	centerY := vt.GameObj.PosGlobal().Y + height/2

	// 	// bounding box
	// 	rl.DrawRectangleLines(int32(vt.GameObj.PosGlobal().X), int32(vt.GameObj.PosGlobal().Y), int32(width), int32(height), rl.Red)
	// 	// center point
	// 	rl.DrawCircleLines(int32(centerX), int32(centerY), 3, rl.Yellow)
	// }
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
