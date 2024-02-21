package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// could also repel if force is negative
type Approach struct {
	GameObj *GameObj
	OtherTags []string

	SafeAngle float32
	CallbackAngle func()

	SafeDistance float32
	CallbackDistance func()

	SafeSpeed float32
	CallbackSpeed func()
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

func WithSafeDistance(distance float32, callback func()) ApproachOption {
	return func(a *Approach) {
		a.SafeDistance = distance
		a.CallbackDistance = callback
	}
}

func WithSafeAngle(angle float32, callback func()) ApproachOption {
	return func(a *Approach) {
		a.SafeAngle = angle
		a.CallbackAngle = callback
	}
}

func WithSafeSpeed(speed float32, callback func()) ApproachOption {
	return func(a *Approach) {
		a.SafeSpeed = speed
		a.CallbackSpeed = callback
	}
}


func (a *Approach) Update() {

	// check distance of all objects with tags
	for _, tag := range a.OtherTags {
		objs := a.GameObj.Parent.FindChildrenByTags(true, tag)
		for _, obj := range objs {
			if a.IsClose(obj) {
				if a.IsSafe(obj) {
					fmt.Println("Safe angle and speed")
				} else {
					fmt.Println("CAREFUL!")
				}
			}
		}
	}

}

func (a *Approach) Draw() {
	// no op
}

func (a *Approach) IsClose(target *GameObj) bool {
	// Vector from the ship to planet
	planetVec := rl.Vector2Subtract(target.Position, a.GameObj.Position)
	
	// Within thresholds?
	close := rl.Vector2Length(planetVec) <= a.SafeDistance
	
	if close && a.CallbackDistance != nil {
		a.CallbackDistance()
	}
	
	return close
}

func (a *Approach) IsSafe(target *GameObj) bool {
	
		// Vector from the ship to planet
		planetVec := rl.Vector2Subtract(target.Position, a.GameObj.Position)
		shipVel := a.GameObj.Components["motion"].(*Motion).Velocity
	
		// Normalize vectors
		planetVec = rl.Vector2Normalize(planetVec)
		shipVel = rl.Vector2Normalize(shipVel)
	
		// Angle of approach in degrees
		angleOfApproach := rl.Vector2Angle(shipVel, planetVec)
	
		// Within thresholds?
		safe := float32(
			math.Abs(float64(angleOfApproach))) <=
			a.SafeAngle && rl.Vector2Length(shipVel) <= 
			a.SafeSpeed
			
		return safe
	}