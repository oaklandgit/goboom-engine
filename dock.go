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
	opts ...DockOption) *GameObj {

	dock := &Dock{
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(dock)
	}

	obj.AddComponents(dock)

	return obj
}

func angleAtPosition(rotation float32, center, position rl.Vector2) float32 {
    dx := position.X - center.X
    dy := position.Y - center.Y
    angle := float32(math.Atan2(float64(dy), float64(dx))) * rl.Rad2deg

	return adjustAngle(angle - rotation)
}

func adjustAngle(angle float32) float32 {
    // Use modulo to wrap around
    return float32(int(angle+360) % 360)
}

func (d *Dock) DockWith(other *GameObj, atPosition rl.Vector2) {

	sound := sounds["sounds/dock.wav"]
	rl.PlaySound(sound)

	d.DockedWith = other
	other.Components["mine"].(*Mine).MinedBy = d.GameObj
	d.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()

	d.AngleOffset = angleAtPosition(other.Angle, other.PosGlobal(), d.GameObj.PosGlobal())

}

func (d *Dock) Update() {	

	if d.DockedWith == nil { return }

	radius := d.DockedWith.Width() / 2 + DOCK_HEIGHT

	// angleDeg := adjustAngle(float32(math.Abs(float64(d.DockedWith.Angle - d.AngleOffset))) - 90)
	// angleDeg := adjustAngle(float32(math.Abs(float64(d.DockedWith.Angle - d.AngleOffset - 90))))
	angleRads := float64(adjustAngle(d.AngleOffset + d.DockedWith.Angle)) * rl.Deg2rad

	x := d.DockedWith.PosGlobal().X + (radius * float32(math.Cos(angleRads)))
	y := d.DockedWith.PosGlobal().Y + (radius * float32(math.Sin(angleRads)))

	d.GameObj.Position = rl.NewVector2(x, y)

	d.GameObj.Angle = adjustAngle(d.AngleOffset + d.DockedWith.Angle)

}

func (d *Dock) Undock() {

	if d.DockedWith == nil { return }

	sound := sounds["sounds/undock.wav"]
	rl.PlaySound(sound)

	// move ship a bit to avoid immediate re-docking
	displacement := displace(DOCK_HEIGHT, d.GameObj.Angle)
	d.GameObj.Position = rl.Vector2Add(d.GameObj.Position, displacement)

	// sever the connection
	d.DockedWith.Components["mine"].(*Mine).MinedBy = nil
	d.DockedWith = nil
}

func (d *Dock) Draw() {
	if d.DockedWith == nil { return }

	text := d.DockedWith.Name
	color := d.DockedWith.Components["sprite"].(*Sprite).Color
	DrawText(text, screenW/2, screenH - 62, 32, 3, color, Center)
}