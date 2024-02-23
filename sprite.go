package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	GameObj *GameObj
	Texture rl.Texture2D
	Frames []rl.Rectangle
	CurrFrame int
	Opacity float32
	Color rl.Color
	FlipX bool
	FlipY bool
}

func (*Sprite) Id() string {
	return "sprite"
}

type SpriteOption func(*Sprite)

func WithOpacity(opacity float32) SpriteOption {
	return func(s *Sprite) {
		s.Opacity = opacity
	}
}

func WithColor(color rl.Color) SpriteOption {
	return func(s *Sprite) {
		s.Color = color
	}	
}

func WithFrames(rows, cols, frames int) SpriteOption {
	return func(s *Sprite) {
		for i := 0; i < frames; i++ {
			frame := rl.NewRectangle(
				float32(i % cols) * float32(s.Texture.Width) / float32(cols),
				float32(i / cols) * float32(s.Texture.Height) / float32(rows),
				float32(s.Texture.Width) / float32(cols),
				float32(s.Texture.Height) / float32(rows),
			)
			s.Frames = append(s.Frames, frame)
		}
	}
}

func WithFlip(X bool, Y bool) SpriteOption {
	return func(s *Sprite) {
		s.FlipX = X
		s.FlipY = Y
	}
}

func (o *GameObj) NewSprite(tex rl.Texture2D, opts ...SpriteOption) *Sprite {

	// default values that would otherwise be zero
	sprite := &Sprite{
		GameObj: o,
		Texture: tex,
		Opacity: 1.0,
		Color: rl.White,
	}

	for _, opt := range opts {
		opt(sprite)
	}

	// apply the texture size to the game object size
	w := float32(tex.Width) * o.Scale.X
	h := float32(tex.Height) * o.Scale.Y
	o.Size = rl.NewVector2(w, h)

	o.AddComponents(sprite)

	return sprite
}

func (s *Sprite) Center() rl.Vector2 {
	return rl.NewVector2(float32(s.Texture.Width/2), float32(s.Texture.Height/2))
}

func (s *Sprite) AnchorPoint(w, h float32) rl.Vector2 {
	return rl.NewVector2(
		s.GameObj.Origin.X * w,
		s.GameObj.Origin.Y * h,
	)
}

func (s *Sprite) Update() {
	// no op
}

func (s *Sprite) GetSpriteRect() (rl.Rectangle, rl.Rectangle) {

	// SOURCE
	texW := float32(s.Texture.Width)
	texH := float32(s.Texture.Height)

	var source rl.Rectangle
	if len(s.Frames) == 0 {
		source = rl.NewRectangle(0, 0, texW, texH)
	} else {
		source = s.Frames[s.CurrFrame]
	}

	// DEST
	dest := rl.NewRectangle(
		s.GameObj.PosGlobal().X,
		s.GameObj.PosGlobal().Y,
        source.Width * s.GameObj.Scale.X,
        source.Height * s.GameObj.Scale.Y,)

	return source, dest

}

func (s *Sprite) Draw() {

	source, dest := s.GetSpriteRect()
	objR := s.GameObj.Angle

	texW := float32(s.Texture.Width)
	texH := float32(s.Texture.Height)

	if s.FlipX {
		texW *= -1
	}

	if s.FlipY {
		texH *= -1
	}

	origin := s.AnchorPoint(dest.Width, dest.Height)
	color := s.Color
	if s.Opacity < 1.0 {
		color.A = uint8(255 * s.Opacity)
	}
	
	rl.DrawTexturePro(s.Texture, source, dest, origin, objR, color)
}