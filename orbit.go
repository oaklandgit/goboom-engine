package main

import (
	"fmt"
	"math"
)

type Orbit struct {
	GameObj *GameObj
	Target *GameObj
	Angle float32
	Speed float32
	Distance float32
}

type OrbitOptions func(*Orbit)

func (obj *GameObj) NewOrbit(
	target *GameObj,
	speed float32,
	distance float32,
	opts ...OrbitOptions) *Orbit {

	orbit := &Orbit{
		GameObj: obj,
		Target: target,
		Speed: speed,
		Distance: distance,
	}

	for _, opt := range opts {
		opt(orbit)
	}

	obj.AddComponents(orbit)

	return orbit
}

func (orb *Orbit) Update() {

	orb.Angle += orb.Speed

	if orb.Angle > 360 {
		orb.Angle = 0
	}

	// update position of orbiting object
	
	// o.Angle += o.Speed * (180 / math.Pi)
	radians := orb.Angle * (math.Pi / 180)
	orb.GameObj.Position.X = orb.Target.Position.X + orb.Distance * float32(math.Cos(float64(radians)))
	orb.GameObj.Position.Y = orb.Target.Position.Y + orb.Distance * float32(math.Sin(float64(radians)))
	// o.GameObj.Position.X = o.Target.Position.X + o.Distance
	// o.GameObj.Position.Y = o.Target.Position.Y + o.Distance

	// orb.GameObj.Position.X = orb.Target.Position.X + orb.Distance

	fmt.Println(orb.Distance)

	// fmt.Printf("Angle: %v, GameObj Position: (%v, %v), Target Position: (%v, %v), Speed: %v, Distance: %v\n",
	// o.Angle, o.GameObj.Position.X, o.GameObj.Position.Y, o.Target.Position.X, o.Target.Position.Y, o.Speed, o.Distance)
}

func (o *Orbit) Draw() {
	// no op
}