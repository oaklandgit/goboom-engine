package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Starfield struct {
	GameObj *GameObj
	Texture rl.Texture2D
}

func (*Starfield) Id() string {
	return "starfield"
}

func createStarfieldTexture(w, h, density int) rl.Texture2D {

	stars := rl.GenImageColor(w, h, rl.Color{0, 0, 0, 0})

	for i := 0; i < density; i++ {
		o := rl.GetRandomValue(40, 255) // opacity
		r := rl.GetRandomValue(160, 255)
		g := rl.GetRandomValue(160, 255)
		b := rl.GetRandomValue(160, 255)
		x := float32(rl.GetRandomValue(0, int32(w)))
		y := float32(rl.GetRandomValue(0, int32(h)))
		rl.ImageDrawPixel(
			stars,
			int32(x), int32(y),
			rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(o)))
	}

	tex := rl.LoadTextureFromImage(stars)

	return tex
}

func NewStarfield(
	obj *GameObj,
	w, h int, density int) {

	starfield := &Starfield{
		GameObj: obj,
		Texture: createStarfieldTexture(
			w,
			h,
			density),
	}

	obj.AddComponents(starfield)
}

func (s *Starfield) Update() {
}

func (s *Starfield) Draw() {
	rl.DrawTexture(s.Texture, int32(s.GameObj.PosGlobal().X), int32(s.GameObj.PosGlobal().Y), rl.White)
}
