package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type VecText struct {
	GameObj *GameObj
	Text    string
	Size    float32
	Leading float32
	Color   rl.Color
}

func (o *GameObj) NewVecText(str string, s float32, lead float32, c rl.Color) *GameObj {

	vecText := &VecText{
		GameObj: o,
		Text:    str,
		Size:    s,
		Leading: lead,
		Color:   c,
	}

	o.AddComponents(vecText)

	return o
}

func (*VecText) Id() string {
	return "vecText"
}

func (p *VecText) Update() {
	// no op
}

func (p *VecText) Draw() {

	for i, char := range p.Text {

		offsetX := float32(i) * (2 + p.Leading) // 2 is the standard width of a character before scaling

		drawChar(
			string(char),
			int32(p.Size),
			rl.NewVector2(p.GameObj.Position.X+offsetX, p.GameObj.Position.Y),
			p.Color,
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
