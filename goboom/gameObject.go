package goboom

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObj struct {
	Name string
	Tags map[string]struct{}
	//
	Game *Game
	Parent *GameObj
	Children []*GameObj
	//
	Position rl.Vector2
	Size rl.Vector2
	Layer int // TO DO
	Origin rl.Vector2
	Angle float32
	Scale	rl.Vector2
	//
	Components map[string]Component
	//
	Deleted bool
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
}

func (o *GameObj) RemoveComponent(id string) {
	delete(o.Components, id)
}

func (o *GameObj) HasTag(tag string) bool {
	_, exists := o.Tags[tag]
	return exists
}

func (o *GameObj) Profile(tags... string) {
	// fmt.Printf("====== %v ======\n", time.Now())
	fmt.Println("================")
	for _, t := range tags {
		fmt.Printf("%-12s %d\n", t, len(o.FindChildrenByTags(true, t)))
	}
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

	// printObjs(objs)
	return objs
}

func WithTags(tags ...string) GameObjOption {
	return func(o *GameObj) {
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

func (game *Game) NewGameObject(name string, opts ...GameObjOption) *GameObj {

	obj := &GameObj{
		Name: name,
		Game: game,
		Scale: rl.Vector2{X: 1, Y: 1},
		Origin: rl.Vector2{X: 0.5, Y: 0.5},
		Components: make(map[string]Component),
		Tags: make(map[string]struct{}),
		Deleted: false,
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}

func (o *GameObj) AddChildren(children ...*GameObj) {
	for _, c := range children {
		c.Parent = o
	}

	o.Children = append(o.Children, children...)
}

func (o *GameObj) Delete() {
	o.Deleted = true
}

func (o *GameObj) Update() {

	// update components
	for _, comp := range o.Components {
		comp.Update()
	}

	// update and remove children
	var newChildren []*GameObj
	for _, c := range o.Children {
		if !c.Deleted {
			newChildren = append(newChildren, c)
		}

		c.Update()
	}
	o.Children = newChildren
}

func (o *GameObj) Draw() {

	for _, c := range o.Components {
		c.Draw()
	}

	for _, c := range o.Children {
		c.Draw()
	}
}