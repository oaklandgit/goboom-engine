package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type VecText struct {
	GameObj *GameObj
	Text    string
	Size    float32
	Gap     float32
	Color   rl.Color
	Align   TextAlignment
}

type TextAlignment int

const (
	TextLeft TextAlignment = iota
	TextRight
	TextCenter
)

func (o *GameObj) NewVecText(str string, gap float32, c rl.Color, opts ...VecTextOption) *GameObj {

	vecText := &VecText{
		GameObj: o,
		Text:    str,
		Gap:     gap,
		Color:   c,
		Align:   TextLeft,
	}

	for _, opt := range opts {
		opt(vecText)
	}

	o.AddComponents(vecText)

	return o
}

type VecTextOption func(*VecText)

func WithAlignment(a TextAlignment) VecTextOption {
	return func(vt *VecText) {
		vt.Align = a
	}
}

func (*VecText) Id() string {
	return "vecText"
}

func (vt *VecText) Update() {
	// no op
}

func (vt *VecText) Draw() {

	startX := float32(vt.GameObj.Position.X)

	numChars := len(vt.Text)
	totalWidth := float32(numChars) * (2 + vt.Gap) * vt.GameObj.Scale.X

	switch vt.Align {
	case TextLeft:
		// no op
	case TextRight:
		startX -= totalWidth
	case TextCenter:
		startX -= (totalWidth / 2)
	}

	for i, char := range vt.Text {

		offsetX := startX + float32(i)*(2+vt.Gap)*vt.GameObj.Scale.X // 2 is the standard width of a character before scaling

		// each char
		drawChar(
			string(char),
			int32(vt.GameObj.Scale.X),
			rl.NewVector2(offsetX, vt.GameObj.Position.Y),
			vt.Color,
		)
	}

}

func drawChar(char string, size int32, pos rl.Vector2, color rl.Color) {
	lines, ok := letterShapes[char]
	if !ok {
		return
	}

	for _, line := range lines {
		for i := 0; i < len(line.P)-1; i++ {
			start := rl.Vector2{X: line.P[i].X, Y: line.P[i].Y}
			end := rl.Vector2{X: line.P[i+1].X, Y: line.P[i+1].Y}

			start = rl.Vector2Multiply(start, rl.Vector2{X: float32(size), Y: float32(size)})
			end = rl.Vector2Multiply(end, rl.Vector2{X: float32(size), Y: float32(size)})

			// add position to each point
			start = rl.Vector2Add(start, pos)
			end = rl.Vector2Add(end, pos)

			// rl.DrawLineV(start, end, color)
			rl.DrawLineEx(start, end, 1.5, color)
		}
	}
}

type Point struct {
	X, Y float32
}

type Line struct {
	P []Point
}

var letterShapes = map[string][]Line{
	"A": {
		{P: []Point{{0, 2}, {0, 1}, {1, 0}, {2, 1}, {2, 2}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
	"B": {
		{P: []Point{{0, 2}, {0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}, {0, 2}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
	"C": {
		{P: []Point{{2, 2}, {0, 2}, {0, 0}, {2, 0}}},
	},
	"D": {
		{P: []Point{{0, 2}, {0, 0}, {1, 0}, {2, 1}, {2, 2}, {0, 2}}},
	},
	"E": {
		{P: []Point{{2, 2}, {0, 2}, {0, 0}, {2, 0}}},
		{P: []Point{{0, 1}, {1, 1}}},
	},
	"F": {
		{P: []Point{{0, 2}, {0, 0}, {2, 0}}},
		{P: []Point{{0, 1}, {1, 1}}},
	},
	"G": {
		{P: []Point{{2, 0}, {0, 0}, {0, 2}, {2, 2}, {2, 1}, {1, 1}}},
	},
	"H": {
		{P: []Point{{0, 0}, {0, 2}}},
		{P: []Point{{2, 0}, {2, 2}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
	"I": {
		{P: []Point{{1, 0}, {1, 2}}},
		{P: []Point{{0, 0}, {2, 0}}},
		{P: []Point{{0, 2}, {2, 2}}},
	},
	"J": {
		{P: []Point{{2, 0}, {2, 2}, {0, 2}, {0, 1}}},
	},
	"K": {
		{P: []Point{{0, 0}, {0, 2}}},
		{P: []Point{{2, 0}, {0, 1}, {2, 2}}},
	},
	"L": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}}},
	},
	"M": {
		{P: []Point{{0, 2}, {0, 0}, {1, 1}, {2, 0}, {2, 2}}},
	},
	"N": {
		{P: []Point{{0, 2}, {0, 0}, {2, 2}, {2, 0}}},
	},
	"O": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}, {2, 0}, {0, 0}}},
	},
	"P": {
		{P: []Point{{0, 2}, {0, 0}, {2, 0}, {2, 1}, {0, 1}}},
	},
	"Q": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}, {2, 0}, {0, 0}}},
		{P: []Point{{1, 1}, {2.5, 2.5}}},
	},
	"R": {
		{P: []Point{{0, 2}, {0, 0}, {2, 0}, {2, 1}, {0, 1}}},
		{P: []Point{{1, 1}, {2, 2}}},
	},
	"S": {
		{P: []Point{{2, 0}, {0, 0}, {0, 1}, {2, 1}, {2, 2}, {0, 2}}},
	},
	"T": {
		{P: []Point{{0, 0}, {2, 0}}},
		{P: []Point{{1, 0}, {1, 2}}},
	},
	"U": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}, {2, 0}}},
	},
	"V": {
		{P: []Point{{0, 0}, {1, 2}, {2, 0}}},
	},
	"W": {
		{P: []Point{{0, 0}, {0, 2}, {1, 1}, {2, 2}, {2, 0}}},
	},
	"X": {
		{P: []Point{{0, 0}, {2, 2}}},
		{P: []Point{{2, 0}, {0, 2}}},
	},
	"Y": {
		{P: []Point{{0, 0}, {1, 1}, {2, 0}}},
		{P: []Point{{1, 1}, {1, 2}}},
	},
	"Z": {
		{P: []Point{{0, 0}, {2, 0}, {0, 2}, {2, 2}}},
	},
	"0": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}, {2, 0}, {0, 0}}},
		{P: []Point{{0, 2}, {2, 0}}},
	},
	"1": {
		{P: []Point{{0, 0}, {1, 0}, {1, 2}}},
		{P: []Point{{0, 2}, {2, 2}}},
	},
	"2": {
		{P: []Point{{0, 0}, {2, 0}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}},
	},
	"3": {
		{P: []Point{{0, 0}, {2, 0}, {0, 1}, {2, 1}, {2, 2}, {0, 2}}},
	},
	"4": {
		{P: []Point{{0, 0}, {0, 1}, {2, 1}}},
		{P: []Point{{2, 0}, {2, 2}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
	"5": {
		{P: []Point{{2, 0}, {0, 0}, {0, 1}, {2, 1}, {2, 2}, {0, 2}}},
	},
	"6": {
		{P: []Point{{2, 0}, {0, 0}, {0, 2}, {2, 2}, {2, 1}, {0, 1}}},
	},
	"7": {
		{P: []Point{{0, 0}, {2, 0}, {0, 2}}},
	},
	"8": {
		{P: []Point{{0, 0}, {0, 2}, {2, 2}, {2, 0}, {0, 0}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
	"9": {
		{P: []Point{{0, 2}, {2, 2}, {2, 0}, {0, 0}, {0, 1}, {2, 1}}},
	},
	" ": {},
}
