package main

import (
	"math"

	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const DOCK_HEIGHT = 12

// could also repel if force is negative
type Dock struct {
	GameObj *boom.GameObj
	DockedWith *boom.GameObj
	AngleOffset float32
}

func (*Dock) Id() string {
	return "dock"
}

type DockOption func(*Dock)

func NewDock(
	obj *boom.GameObj,
	opts ...DockOption) *boom.GameObj {

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

func (d *Dock) DockWith(other *boom.GameObj, atPosition rl.Vector2) {

	sound := game.Sounds["sounds/dock.wav"]
	rl.PlaySound(sound)

	d.DockedWith = other
	other.Components["mine"].(*Mine).MinedBy = d.GameObj
	d.GameObj.Components["motion"].(*boom.Motion).Velocity = rl.Vector2Zero()

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

	sound := game.Sounds["sounds/undock.wav"]
	rl.PlaySound(sound)

	// move ship a bit to avoid immediate re-docking
	displacement := boom.Displace(DOCK_HEIGHT, d.GameObj.Angle)
	d.GameObj.Position = rl.Vector2Add(d.GameObj.Position, displacement)

	// sever the connection
	d.DockedWith.Components["mine"].(*Mine).MinedBy = nil
	d.DockedWith = nil
}

func (d *Dock) Draw() {
	if d.DockedWith == nil { return }

	text := d.DockedWith.Name
	color := d.DockedWith.Components["sprite"].(*boom.Sprite).Color
	boom.DrawText(text, int32(game.Width/2), int32(game.Height/2) - 62, 32, 3, color, boom.Center)
}