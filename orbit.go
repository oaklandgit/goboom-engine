package main

import (
	"math"
)

type Orbit struct {
	GameObj *GameObj
	// Target *GameObj
	Angle float32
	Speed float32
	Distance float32
}

func (*Orbit) Id() string {
	return "orbit"
}


type OrbitOptions func(*Orbit)

func (obj *GameObj) NewOrbit(
	// target *GameObj,
	speed float32,
	distance float32,
	opts ...OrbitOptions) *Orbit {

	orbit := &Orbit{
		GameObj: obj,
		// Target: target,
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

	radians := orb.Angle * (math.Pi / 180)
	orb.GameObj.Position.X = orb.Distance * float32(math.Cos(float64(radians)))
	orb.GameObj.Position.Y = orb.Distance * float32(math.Sin(float64(radians)))

}

func (o *Orbit) Draw() {
	// no op
}