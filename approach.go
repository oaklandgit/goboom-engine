package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ARC_SEGMENTS = 12
)

type Approach struct {
	GameObj *GameObj
	Target *GameObj

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

	dir := rl.Vector2Subtract(o.PosGlobal(), a.GameObj.PosGlobal())
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

	dir := rl.Vector2Subtract(o.PosGlobal(), a.GameObj.PosGlobal())
	dir = rl.Vector2Normalize(dir)
	dot := rl.Vector2DotProduct(headingVector, dir)

	return dot > 0 // moving towards if > 0
}

func (a *Approach) IsClose(target *GameObj) bool {

	// Vector from the ship to planet
	planetVec := rl.Vector2Subtract(target.PosGlobal(), a.GameObj.PosGlobal())
	planetRadius := target.Width() / 2

	close := rl.Vector2Length(planetVec) <= a.SafeDistance + planetRadius
	
	return close
}

func (a *Approach) IsSafeSpeed() bool {
	shipVel := a.GameObj.Components["motion"].(*Motion).Velocity
	safe := rl.Vector2Length(shipVel) <= a.SafeSpeed
	return safe
}

func (a *Approach) Update() {

	a.Message = ""
	a.Target = nil

	if a.GameObj.Components["dock"].(*Dock).DockedWith != nil {
		return
	}

	// check distance of all objects with tags
	for _, tag := range a.OtherTags {
		objs := a.GameObj.Parent.FindChildrenByTags(true, tag)
		for _, obj := range objs {
			if a.IsClose(obj) && a.IsMovingTowards(obj) {

				a.Message = fmt.Sprintf("Approaching %s", obj.Name)
				a.Target = obj

				if a.IsPointingToward(obj) {
					a.Message = "Turn to dock!"
					return
				}

				if !a.IsSafeSpeed() {
					a.Message = "Adjust your speed!"
					return
				}

			}
		}
	}

}

func (a *Approach) Draw() {
	// DrawText(a.Message, screenW/2, screenH - 14, 14, 2, rl.White, Center)

	if a.Target == nil {
		return
	}

	DrawText(
		a.Message,
		int32(a.GameObj.PosGlobal().X),
		int32(a.GameObj.PosGlobal().Y) + 16,
		14, 2, rl.Green, Center)

	arcRadius := calculateRadius(
		a.Target.PosGlobal(),
		a.Target.Parent.PosGlobal())

	targetAngle := calculateAngle(
		a.Target.PosGlobal(),
		a.Target.Parent.PosGlobal())
	startAngle := targetAngle - 45
	endAngle := targetAngle + 45

	rl.DrawCircleLines(
		int32(a.Target.PosGlobal().X),
		int32(a.Target.PosGlobal().Y),
		a.Target.Width()/2 + (a.Target.Width()/2 * 1.1),
		rl.Green)

	rl.DrawRingLines(
		a.Target.Parent.PosGlobal(),
		arcRadius,
		arcRadius,
		startAngle,
		endAngle,
		ARC_SEGMENTS,
		rl.Green)
}

func calculateRadius(targetPos, parentPos rl.Vector2) float32 {
    dx := targetPos.X - parentPos.X
    dy := targetPos.Y - parentPos.Y
    return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// func calculateAngle(targetPos, parentPos rl.Vector2) float32 {
//     dx := targetPos.X - parentPos.X
//     dy := targetPos.Y - parentPos.Y
//     return float32(math.Atan2(float64(dy), float64(dx))) * 180 / math.Pi
// }