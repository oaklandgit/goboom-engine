package main

import (
	"fmt"
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

	// let the mine know I'm here
	other.Components["mine"].(*Mine).MinedBy = d.GameObj

	// stop motion
	d.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
	
	// stop attracting and colliding
	d.GameObj.Tags["docked"] = struct{}{}
	other.Tags["docking"] = struct{}{}

	// fanfare
	fmt.Printf("Landed on %s!\n", other.Name)
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

	// delete the tags
	delete(d.GameObj.Tags, "docked")
	delete(d.DockedWith.Tags, "docking")
	
	// move ship a bit to avoid immediate re-docking
	displacement := displace(DOCK_HEIGHT, d.GameObj.Angle)
	d.GameObj.Position = rl.Vector2Add(d.GameObj.Position, displacement)

	// sever the connection
	d.DockedWith.Components["mine"].(*Mine).MinedBy = nil
	d.DockedWith = nil
}

func (d *Dock) Update() {

	if d.DockedWith == nil {
		return
	}

	radius := d.DockedWith.Width() / 2 + DOCK_HEIGHT
	angle := d.DockedWith.Angle * rl.Deg2rad

	x := d.DockedWith.PosGlobal().X + (radius * float32(math.Cos(float64(angle))))
	y := d.DockedWith.PosGlobal().Y + (radius * float32(math.Sin(float64(angle))))

	d.GameObj.Position = rl.NewVector2(x, y)

	d.GameObj.Angle = d.DockedWith.Angle

}

func (d *Dock) Draw() {
	if d.DockedWith == nil { return }

	fontSize := int32(20)
	text := fmt.Sprintf("Docked with %s", d.DockedWith.Name)
	textWidth := rl.MeasureText(text, fontSize)

	color := d.DockedWith.Components["sprite"].(*Sprite).Color

	rl.DrawText(text, 400 - textWidth/2, 408, fontSize, color)
}