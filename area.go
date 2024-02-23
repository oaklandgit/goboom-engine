package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionShape interface {
	Width() float32
	Height() float32
	Draw(x, y float32, color rl.Color) // for debugging
}

type Area struct {
	GameObj *GameObj
	Shape CollisionShape
	Collided bool
	CollisionHandlers map[string]func(*GameObj, *GameObj)
	LastCollisionTime time.Time
	CollisionCooldown time.Duration
}

func (*Area) Id() string {
	return "area"
}


type AreaOptions func(*Area)

func (obj *GameObj) NewArea(
	shape CollisionShape,
	opts ...AreaOptions) *Area {

	area := &Area{
		GameObj: obj,
		Shape: shape,
		CollisionHandlers: make(map[string]func(*GameObj, *GameObj)),
	}

	for _, opt := range opts {
		opt(area)
	}

	obj.AddComponents(area)

	return area
}

func WithCooldown(cooldown time.Duration) AreaOptions {
	return func(a *Area) {
		a.CollisionCooldown = cooldown
	}
}

func (a *Area) AddCollisionHandler(
		tag string,
		handler func(*GameObj, *GameObj),
		) {
	a.CollisionHandlers[tag] = handler
}

func (a *Area) CollidedWith(other *GameObj) bool {
	b := other.Components["area"].(*Area)

	if time.Since(a.LastCollisionTime) < a.CollisionCooldown {
		return false
	}

	collided := rl.CheckCollisionCircles(
		a.GameObj.PosGlobal(),
		a.Shape.Width()/2,
		b.GameObj.PosGlobal(),
		b.Shape.Width()/2)

	if collided {
		a.LastCollisionTime = time.Now()
	}

	return collided
}

func (a *Area) Update() {
	// no op
}

func (a *Area) Draw() {
	if !DEBUG {
		return
	}

	color := rl.Green
	if a.Collided {
		color = rl.Red
	}
	a.Shape.Draw(a.GameObj.PosGlobal().X, a.GameObj.PosGlobal().Y, color)
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

func (c CircleCollider) Collides(s *CollisionShape) bool {
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

func (r RectangleCollider) Collides(s *CollisionShape) bool {
	return false
}

