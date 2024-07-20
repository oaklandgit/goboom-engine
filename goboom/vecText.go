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

			rl.DrawLineV(start, end, color)
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
		{P: []Point{{0, 2}, {0, 1}, {1, 0}, {2, 1}, {2, 2}}},
		{P: []Point{{0, 1}, {2, 1}}},
	},
}
