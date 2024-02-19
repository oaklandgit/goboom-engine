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

func (d *Dock) DockWith(other *GameObj) {
	d.DockedWith = other

	// stop motion
	d.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
	
	// stop attracting and colliding
	d.GameObj.Tags = append(d.GameObj.Tags, "docked")
	other.Tags = append(other.Tags, "docking")

	// fanfare
	fmt.Printf("Landed on %s!\n", other.Name)
}

func remove(s []string, r string) []string {
    for i, v := range s {
        if v == r {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}

func (d *Dock) Undock() {
	d.GameObj.Tags = remove(d.GameObj.Tags, "docked")
	d.DockedWith.Tags = remove(d.DockedWith.Tags, "docking")
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
	// d.GameObj.Position = d.DockedWith.PosGlobal()

}

func (d *Dock) Draw() {
	// no op
}