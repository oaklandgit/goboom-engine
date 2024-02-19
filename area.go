package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape interface {
	Width() float32
	Height() float32
	Collides(*Shape) bool
	Draw(x, y float32, color rl.Color) // for debugging
}

type Area struct {
	GameObj *GameObj
	Shape Shape
	CollisionHandlers map[string]func(*GameObj)
}

func (*Area) Id() string {
	return "area"
}


type AreaOptions func(*Area)

func (obj *GameObj) NewArea(
	shape Shape,
	opts ...AreaOptions) *Area {

	area := &Area{
		GameObj: obj,
		Shape: shape,
	}

	for _, opt := range opts {
		opt(area)
	}

	obj.AddComponents(area)

	return area
}

func (a *Area) AddCollisionHandler(tag string, handler func(*GameObj)) {
	a.CollisionHandlers[tag] = handler
}

func (a *Area) Update() {
}

func (a *Area) Draw() {
	// DEBUG
	// a.Shape.Draw(a.GameObj.PosGlobal().X, a.GameObj.PosGlobal().Y, rl.Green)
}

type CircleCollider struct {
	Radius float32
}

func (c CircleCollider) Width() float32 {
	return c.Radius * 2
}

func (c CircleCollider) Height() float32 {
	return c.Radius * 2
}

func (c CircleCollider) Collides(s *Shape) bool {
	return false
}

func (c CircleCollider) Draw(x, y float32, color rl.Color) {
	rl.DrawCircleLines(int32(x), int32(y), c.Radius, color)
}

type RectangleCollider struct {
	W float32
	H float32
}

func (r RectangleCollider) Width() float32 {
	return r.W
}

func (r RectangleCollider) Height() float32 {
	return r.H
}

func (r RectangleCollider) Draw(x, y float32, color rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(r.W), int32(r.H), color)
}

func (r RectangleCollider) Collides(s *Shape) bool {
	return false
}

