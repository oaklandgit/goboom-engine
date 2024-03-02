package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Motion struct {
	GameObj *GameObj
	Velocity rl.Vector2
	MaxVelocity float32
	Friction float32
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

func (o *GameObj) NewMotion(opts ...MotionOption) *GameObj {

	motion := &Motion{
		GameObj: o,
	}

	for _, opt := range opts {
		opt(motion)
	}

	o.AddComponents(motion)

	return o
}

func (m *Motion) SetVelocity(speed float32, heading float32) {
	// rads := float64(heading) * (math.Pi / 180)
	rads := float64(heading * rl.Deg2rad)
	m.Velocity =
		rl.Vector2Add(m.Velocity, rl.Vector2{
			X: speed * float32(math.Cos(rads)),
			Y: speed * float32(math.Sin(rads)),
		})
}


func WithVelocity(speed float32, heading float32) MotionOption {
	return func(m *Motion) {
		m.SetVelocity(speed, heading)
	}
}

func WithFriction(friction float32) MotionOption {
	return func(m *Motion) {
		m.Friction = friction
	}
}

func WithMaxVelocity(max float32) MotionOption {
	return func(m *Motion) {
		m.MaxVelocity = max
	}
}

func WithWrap(x, y bool, padding float32) MotionOption {
	return func(m *Motion) {
		m.WrapX = x
		m.WrapY = y
		m.WrapPadding = padding
	}
}

func (m *Motion) Wrap() {
	
	if m.GameObj.Position.X < -m.WrapPadding {
		m.GameObj.Position.X = m.GameObj.Parent.Size.X + m.WrapPadding
	}

	if m.GameObj.Position.X > m.GameObj.Parent.Size.X + m.WrapPadding {
		m.GameObj.Position.X = -m.WrapPadding
	}

	if m.GameObj.Position.Y < -m.WrapPadding {
		m.GameObj.Position.Y = m.GameObj.Parent.Size.Y + m.WrapPadding
	}

	if m.GameObj.Position.Y > m.GameObj.Parent.Size.Y + m.WrapPadding {
		m.GameObj.Position.Y = -m.WrapPadding
	}
}

func (m *Motion) Update() {

	// WRAP
	if m.WrapX || m.WrapY {
		m.Wrap()
	}

	// FRICTION
	if m.Friction > 0 {
		m.Velocity = rl.Vector2Scale(m.Velocity, m.Friction)
	}

	// LIMIT
	if m.MaxVelocity > 0 {
		if rl.Vector2Length(m.Velocity) > m.MaxVelocity {
			m.Velocity = 
				rl.Vector2Scale(
					rl.Vector2Normalize(m.Velocity), m.MaxVelocity)
		}
	}

	// NEW POSITION
	m.GameObj.Position = rl.Vector2Add(m.GameObj.Position, m.Velocity)
	
}
