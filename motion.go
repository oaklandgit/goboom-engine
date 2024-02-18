package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Motion struct {
	GameObj *GameObj
	Speed float32
	Velocity rl.Vector2
	Heading float32
	WrapX bool
	WrapY bool
	WrapPadding float32
}

func (*Motion) Id() string {
	return "motion"
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

func WithVelocity(velocity rl.Vector2) MotionOption {
	return func(m *Motion) {
		m.Velocity = velocity
	}
}

func WithHeading(heading float32) MotionOption {
	return func(m *Motion) {
		m.Heading = heading
	}
}

func WithWrap(x, y bool, padding float32) MotionOption {
	return func(m *Motion) {
		m.WrapX = x
		m.WrapY = y
		m.WrapPadding = padding
	}
}

func (m *Motion) Move(speed float32, heading float32) {
	m.Speed = speed
	m.Heading = heading
}

func (m *Motion) Accelerate(force float32, heading float32) {
	headingVector := rl.Vector2{
        X: float32(math.Cos(float64(heading))),
        Y: float32(math.Sin(float64(heading))),
    }

	acceleration := rl.Vector2Scale(headingVector, force)

	m.Velocity = rl.Vector2Add(m.Velocity, acceleration)
}

func (m *Motion) Stop() {
	m.Velocity = rl.Vector2Zero()
}

func (m *Motion) Update() {

	headingRad := m.Heading * (rl.Pi / 180)
	headingVector := rl.NewVector2(
		float32(math.Cos(float64(headingRad))),
		float32(math.Sin(float64(headingRad))),
	)

	headingVector = rl.Vector2Scale(headingVector, m.Speed)

	if m.WrapX {
		if m.GameObj.Position.X < -m.WrapPadding {
			m.GameObj.Position.X = m.GameObj.Parent.Size.X + m.WrapPadding
		}

		if m.GameObj.Position.X > m.GameObj.Parent.Size.X + m.WrapPadding {
			m.GameObj.Position.X = -m.WrapPadding
		}

		if m.GameObj.Position.Y < -m.WrapPadding {
			m.GameObj.Position.Y = m.GameObj.Parent.Size.X + m.WrapPadding
		}

		if m.GameObj.Position.Y > m.GameObj.Parent.Size.X + m.WrapPadding {
			m.GameObj.Position.Y = -m.WrapPadding
		}
	}
	
	m.GameObj.Position = rl.Vector2Add(m.GameObj.Position, headingVector)

}

