package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Approach struct {
	GameObj *GameObj
	OtherTags []string

	SafeDistance float32
	SafeSpeed float32

	Message string
}

func (*Approach) Id() string {
	return "approach"
}

type ApproachOption func(*Approach)

func (obj *GameObj) NewApproach(
	others []string,
	opts ...ApproachOption) *Approach {

	approach := &Approach{
		OtherTags: others,
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(approach)
	}

	obj.AddComponents(approach)

	return approach
}

func WithSafeDistance(distance float32) ApproachOption {
	return func(a *Approach) {
		a.SafeDistance = distance
	}
}

func WithSafeSpeed(speed float32) ApproachOption {
	return func(a *Approach) {
		a.SafeSpeed = speed
	}
}

func (a *Approach) IsMovingTowards(o *GameObj) bool {

	dir := rl.Vector2Subtract(o.Position, a.GameObj.Position)
	dir = rl.Vector2Normalize(dir)
	dot := rl.Vector2DotProduct(
		a.GameObj.Components["motion"].(*Motion).Velocity, dir)

	return dot > 0 // moving towards if > 0
}

func (a *Approach) IsPointingToward(o *GameObj) bool {

	rads := float64(a.GameObj.Angle) * rl.Deg2rad

	headingX := float32(math.Cos(rads))
	headingY := float32(math.Sin(rads))
	headingVector := rl.NewVector2(headingX, headingY)

	dir := rl.Vector2Subtract(o.Position, a.GameObj.Position)
	dir = rl.Vector2Normalize(dir)
	dot := rl.Vector2DotProduct(headingVector, dir)

	return dot > 0 // moving towards if > 0
}


func (a *Approach) Update() {

	a.Message = ""

	if a.GameObj.Components["dock"].(*Dock).DockedWith != nil {
		return
	}

	// check distance of all objects with tags
	for _, tag := range a.OtherTags {
		objs := a.GameObj.Parent.FindChildrenByTags(true, tag)
		for _, obj := range objs {
			if a.IsClose(obj) && a.IsMovingTowards(obj) {

				a.Message = fmt.Sprintf("Approaching %s", obj.Name)

				if a.IsPointingToward(obj) {
					a.Message = "Turn to dock!"
					return
				}

				if !a.IsSafeSpeed(obj) {
					a.Message = "Adjust your speed!"
					return
				}

			}
		}
	}

}

func (a *Approach) Draw() {
	DrawText(a.Message, 400, 400, 14, 2, rl.White, Center)
}

func (a *Approach) IsClose(target *GameObj) bool {

	// Vector from the ship to planet
	planetVec := rl.Vector2Subtract(target.Position, a.GameObj.Position)
	planetRadius := target.Width() / 2

	close := rl.Vector2Length(planetVec) <= a.SafeDistance + planetRadius
	
	return close
}

func (a *Approach) IsSafeSpeed(target *GameObj) bool {
	shipVel := a.GameObj.Components["motion"].(*Motion).Velocity
	safe := rl.Vector2Length(shipVel) <= a.SafeSpeed
	return safe
}