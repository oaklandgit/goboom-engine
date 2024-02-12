package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Graphics struct {
	GameObj *GameObj
	Width float32
	Height float32
	Texture rl.Texture2D
	Image rl.Image
}

func (o *GameObj) NewGraphics(drawing rl.Image) *Graphics {
	graphics := &Graphics{
		GameObj: o,
		Image: drawing,

	}

	graphics.Texture = rl.LoadTextureFromImage(&drawing)

	o.AddComponents(graphics)

	return graphics
}

func (g *Graphics) Update() {
	// no op
}

func (g *Graphics) Draw() {
	rl.DrawTextureEx(
		g.Texture,
		rl.NewVector2(g.GameObj.Position.X, g.GameObj.Position.Y),
		g.GameObj.Rotation, 1, rl.White)
}