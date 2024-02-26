package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const DOCK_HEIGHT = 12

// could also repel if force is negative
type Dock struct {
	GameObj *GameObj
	DockedWith *GameObj
	AngleOffset float32
}

func (*Dock) Id() string {
	return "dock"
}

type DockOption func(*Dock)

func (obj *GameObj) NewDock(
	opts ...DockOption) *Dock {

	dock := &Dock{
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(dock)
	}

	obj.AddComponents(dock)

	return dock
}

func (d *Dock) DockWith(other *GameObj, atPosition rl.Vector2) {
	d.DockedWith = other
	other.Components["mine"].(*Mine).MinedBy = d.GameObj
	d.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
}

func displace(distance float32, angle float32) rl.Vector2 {
	rads := float64(angle * rl.Deg2rad)
	displacement := rl.Vector2{
    	X: distance * float32(math.Cos(rads)),
    	Y: distance * float32(math.Sin(rads)),
	}
	return displacement
}

func (d *Dock) Undock() {

	if d.DockedWith == nil { return }

	// move ship a bit to avoid immediate re-docking
	displacement := displace(DOCK_HEIGHT, d.GameObj.Angle)
	d.GameObj.Position = rl.Vector2Add(d.GameObj.Position, displacement)

	// sever the connection
	d.DockedWith.Components["mine"].(*Mine).MinedBy = nil
	d.DockedWith = nil
}

func (d *Dock) Update() {

	// testAngleOffset := 90

	if d.DockedWith == nil { return }

	radius := d.DockedWith.Width() / 2 + DOCK_HEIGHT
	angle := float64(d.DockedWith.Angle * rl.Deg2rad)

	x := d.DockedWith.PosGlobal().X + (radius * float32(math.Cos(angle)))
	y := d.DockedWith.PosGlobal().Y + (radius * float32(math.Sin(angle)))

	d.GameObj.Position = rl.NewVector2(x, y)

	d.GameObj.Angle = d.DockedWith.Angle

}

func (d *Dock) Draw() {
	if d.DockedWith == nil { return }

	text := d.DockedWith.Name
	color := d.DockedWith.Components["sprite"].(*Sprite).Color
	DrawText(text, screenW/2, screenH - 62, 32, 3, color, Center)
}