package main

import (
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

	o.AddComponents(sprite)

	return sprite
}

func (s *Sprite) Center() rl.Vector2 {
	return rl.NewVector2(float32(s.Texture.Width/2), float32(s.Texture.Height/2))
}

func (s *Sprite) AnchorPoint() rl.Vector2 {
	return rl.NewVector2(
		s.GameObj.Origin.X * float32(s.Texture.Width),
		s.GameObj.Origin.Y * float32(s.Texture.Height),
	)
}

func (s *Sprite) Update() {
	// no op
}

func (s *Sprite) Draw() {

	width := s.Texture.Width
	height := s.Texture.Height

	if s.FlipX {
		width = -s.Texture.Width
	}

	if s.FlipY {
		height = -s.Texture.Height
	}

	 source := rl.NewRectangle(
		0, 0,
		float32(width),
		float32(height))

    scaledWidth := float32(s.Texture.Width) * s.GameObj.Scale.X
    scaledHeight := float32(s.Texture.Height) * s.GameObj.Scale.Y

	dest := rl.NewRectangle(
		s.GameObj.Position.X - s.GameObj.Origin.X * scaledWidth,
		s.GameObj.Position.Y - s.GameObj.Origin.Y * scaledHeight,
        scaledWidth,
        scaledHeight)

	 origin := s.AnchorPoint()
	 color := s.Color
	 if s.Opacity < 1.0 {
		color.A = uint8(float32(color.A) * s.Opacity)
	 }
	 rl.DrawTexturePro(s.Texture, source, dest, origin, s.GameObj.Rotation, color)

}