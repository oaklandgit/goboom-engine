package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	GameObj *GameObj
	Texture rl.Texture2D
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

func (s *Sprite) AnchorPoint() rl.Vector2 {
	return rl.NewVector2(
		s.GameObj.Origin.X * float32(s.Texture.Width) * s.GameObj.Scale.X,
		s.GameObj.Origin.Y * float32(s.Texture.Height) * s.GameObj.Scale.Y,
	)
}

func (s *Sprite) Update() {
	// no op
}

func (s *Sprite) Draw() {

	texW := float32(s.Texture.Width)
	texH := float32(s.Texture.Height)

	objW := s.GameObj.Width()
	objH := s.GameObj.Height()
	objR := s.GameObj.Angle

	if s.FlipX {
		texW *= -1
	}

	if s.FlipY {
		texH *= -1
	}

	source := rl.NewRectangle(0, 0, texW, texH)

	dest := rl.NewRectangle(
		s.GameObj.PosGlobal().X,
		s.GameObj.PosGlobal().Y,
        objW,
        objH)

	origin := s.AnchorPoint()
	color := s.Color
	if s.Opacity < 1.0 {
	color.A = uint8(float32(color.A) * s.Opacity)
	}
	
	rl.DrawRectanglePro(dest, origin, objR, rl.Color{R: 255, G: 0, B: 0, A: 60})
	rl.DrawTexturePro(s.Texture, source, dest, origin, objR, color)
	posText := fmt.Sprintf("%v", s.GameObj.PosGlobal())
	rl.DrawText(posText, int32(dest.X), int32(dest.Y), 12, rl.White)
}