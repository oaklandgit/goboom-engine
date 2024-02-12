package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Motion struct {
	GameObj *GameObj
	Speed float32
	Heading float32
}

func (m *Motion) Draw() {
	// no op
}

type MotionOption func(*Motion)

func (o *GameObj) NewMotion(opts ...MotionOption) *Motion {

	motion := &Motion{
		GameObj: o,
	}

	for _, opt := range opts {
		opt(motion)
	}

	o.AddComponents(motion)

	return motion
}

func WithSpeed(speed float32) MotionOption {
	return func(m *Motion) {
		m.Speed = speed
	}
}

func WithHeading(heading float32) MotionOption {
	return func(m *Motion) {
		m.Heading = heading
	}
}

func (m *Motion) Move(speed float32, heading float32) {
	m.Speed = speed
	m.Heading = heading
}

func (m *Motion) Stop() {
	m.Speed = 0
}

func (m *Motion) Update() {

	headingRad := m.Heading * (rl.Pi / 180)
	headingVector := rl.NewVector2(
		float32(math.Cos(float64(headingRad))),
		float32(math.Sin(float64(headingRad))),
	)

	headingVector = rl.Vector2Scale(headingVector, m.Speed)
	
	m.GameObj.Position = rl.Vector2Add(m.GameObj.Position, headingVector)

}

