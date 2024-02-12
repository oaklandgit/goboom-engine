package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameObj struct {
	Name string
	Tags []string
	Position rl.Vector2
	Rotation float32
	Origin rl.Vector2
	Scale	rl.Vector2
	Components []Component
}

type Component interface {
	Update()
	Draw()
}

type GameObjOption func(*GameObj)

func (o *GameObj) AddComponents(comps ...Component) {
	o.Components = append(o.Components, comps...)
}

func WithPosition(x, y float32) GameObjOption {
	return func(e *GameObj) {
		e.Position = rl.NewVector2(x, y)
	}
}

func WithRotation(rotation float32) GameObjOption {
	return func(e *GameObj) {
		e.Rotation = rotation
	}
}

func WithScale(x, y float32) GameObjOption {
	return func(e *GameObj) {
		e.Scale = rl.NewVector2(x, y)
	}
}

func WithOrigin(x, y float32) GameObjOption {
	return func(e *GameObj) {
		e.Origin = rl.NewVector2(x, y)
	}
}

func NewGameObject(name string, opts ...GameObjOption) *GameObj {

	// default values that would otherwise be zero
	obj := &GameObj{
		Name: name,
		Scale: rl.Vector2{X: 1, Y: 1},
		Origin: rl.Vector2{X: 0.5, Y: 0.5},
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}