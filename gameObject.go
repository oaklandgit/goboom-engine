package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObj struct {
	Name string
	Tags map[string]struct{}
	Position rl.Vector2 // local
	Layer int
	Offset rl.Vector2
	Angle float32
	LocalAngle float32
	Origin rl.Vector2
	Scale	rl.Vector2
	Components map[string]Component
	Parent *GameObj
	Children []*GameObj
	Size rl.Vector2
}

func (o *GameObj) PosGlobal() rl.Vector2 {
	if o.Parent != nil {
	
		rads := float64(o.Parent.Angle * rl.Deg2rad)
        rotatedPosition := rl.Vector2{
            X: o.Position.X*float32(math.Cos(rads)) - o.Position.Y*float32(math.Sin(rads)),
            Y: o.Position.X*float32(math.Sin(rads)) + o.Position.Y*float32(math.Cos(rads)),
        }
        return rl.Vector2Add(rotatedPosition, o.Parent.PosGlobal())
		
	}

	if o.Position != rl.Vector2Zero() {
		return o.Position
	}
	return rl.Vector2Zero()
}

type Component interface {
	Id() string
	Update()
	Draw()
}

type GameObjOption func(*GameObj)

func (o *GameObj) Width() float32 {
	return o.Size.X
}

func (o *GameObj) Height() float32 {
	return o.Size.Y
}

func (o *GameObj) AddComponents(comps ...Component) {
	for _, c := range comps {
		o.Components[c.Id()] = c
	}
	// o.Components = append(o.Components, comps...)
}

func (o *GameObj) HasTag(tag string) bool {

	_, exists := o.Tags[tag]

	// for _, t := range o.Tags {
	// 	if t == tag {
	// 		return true
	// 	}
	// }
	return exists
}

func (o *GameObj) FindChildrenByTags(recurse bool, tags... string) []*GameObj {
	var objs []*GameObj
	for _, c := range o.Children {
		for _, t := range tags {
			if c.HasTag(t) {
				objs = append(objs, c)
			}
			if recurse && len(c.Children) > 0 {
				objs = append(objs, c.FindChildrenByTags(recurse, tags...)...)
			}
		}
	}
	return objs
}

func (o *GameObj) FindChildrenByComponent(recurse bool, comp string) []*GameObj {
	var objs []*GameObj
	for _, c := range o.Children {
		if c.Components[comp] != nil {
			objs = append(objs, c)
		}
		if recurse && len(c.Children) > 0 {
			objs = append(objs, c.FindChildrenByComponent(recurse, comp)...)
		}
	}
	return objs
}

func WithTags(tags ...string) GameObjOption {
	return func(o *GameObj) {
		// o.Tags = tags
		for _, t := range tags {
			o.Tags[t] = struct{}{}
		}
	}
}

func WithPosition(x, y float32) GameObjOption {
	return func(o *GameObj) {
		o.Position = rl.NewVector2(x, y)
	}
}

func WithScale(x, y float32) GameObjOption {
	return func(o *GameObj) {
		o.Scale = rl.NewVector2(x, y)
	}
}

func WithOrigin(x, y float32) GameObjOption {
	return func(o *GameObj) {
		o.Origin = rl.NewVector2(x, y)
	}
}

func WithAngle(angle float32) GameObjOption {
	return func(o *GameObj) {
		o.Angle = angle
	}
}

func NewGameObject(name string, opts ...GameObjOption) *GameObj {

	// default values that would otherwise be zero
	obj := &GameObj{
		Name: name,
		Scale: rl.Vector2{X: 1, Y: 1},
		Origin: rl.Vector2{X: 0.5, Y: 0.5},
		Components: make(map[string]Component),
		Tags: make(map[string]struct{}),
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}

func (o *GameObj) AddChildren(children ...*GameObj) {
	for _, c := range children {
		c.Parent = o
		// c.Offset = rl.Vector2Subtract(c.Position, o.Position)
		// c.LocalRotation = c.Rotation - o.Rotation
	}

	o.Children = append(o.Children, children...)
}

func (o *GameObj) Update() {

	// if o.Parent != nil {

		// Calculate the rotated offset
		// rotatedOffset := rl.Vector2{
		// 	X: o.Offset.X * float32(math.Cos(float64(o.Parent.Rotation))) - o.Offset.Y * float32(math.Sin(float64(o.Parent.Rotation))),
		// 	Y: o.Offset.X * float32(math.Sin(float64(o.Parent.Rotation))) + o.Offset.Y * float32(math.Cos(float64(o.Parent.Rotation))),
		// }

		// Update position and rotation
		// o.Position = rl.Vector2Add(o.Parent.Position, rotatedOffset)
		// o.Rotation = o.Parent.Rotation + o.LocalRotation

		// o.Position = rl.Vector2Add(o.Parent.Position, o.Offset)
		// o.Rotation = o.Parent.Rotation + o.LocalRotation
	// }

	for _, c := range o.Components {
		c.Update()
	}

	for _, c := range o.Children {
		c.Update()
	}
}

func (o *GameObj) Draw() {
	for _, c := range o.Components {
		c.Draw()
	}

	for _, c := range o.Children {
		c.Draw()
	}
}