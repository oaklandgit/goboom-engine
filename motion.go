package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Motion struct {
	GameObj *GameObj
	Velocity rl.Vector2
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

func (m *Motion) SetVelocity(speed float32, heading float32) {
	rads := float64(heading) * (math.Pi / 180)
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
		m.GameObj.Position.Y = m.GameObj.Parent.Size.X + m.WrapPadding
	}

	if m.GameObj.Position.Y > m.GameObj.Parent.Size.X + m.WrapPadding {
		m.GameObj.Position.Y = -m.WrapPadding
	}
}

func (m *Motion) Update() {

	if m.WrapX || m.WrapY {
		m.Wrap()
	}

	m.GameObj.Position = rl.Vector2Add(m.GameObj.Position, m.Velocity)
	
}
