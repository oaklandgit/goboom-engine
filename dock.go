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

	// stop motion
	d.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
	
	// stop attracting and colliding
	d.GameObj.Tags["docked"] = struct{}{}
	other.Tags["docking"] = struct{}{}

	// fanfare
	fmt.Printf("Landed on %s!\n", other.Name)
}

// func remove(s []string, r string) []string {
//     for i, v := range s {
//         if v == r {
//             return append(s[:i], s[i+1:]...)
//         }
//     }
//     return s
// }

func displace(distance float32, angle float32) rl.Vector2 {
	rads := float64(angle * rl.Deg2rad)
	displacement := rl.Vector2{
    	X: distance * float32(math.Cos(rads)),
    	Y: distance * float32(math.Sin(rads)),
	}
	return displacement
}

func (d *Dock) Undock() {
	fmt.Println(d.GameObj.Tags)
	// d.GameObj.Tags = remove(d.GameObj.Tags, "docked")
	delete(d.GameObj.Tags, "docked")
	fmt.Println(d.GameObj.Tags)
	
	// d.DockedWith.Tags = remove(d.DockedWith.Tags, "docking")
	delete(d.DockedWith.Tags, "docking")
	d.DockedWith = nil
	// move it a bit to avoid immediate re-docking
	displacement := displace(DOCK_HEIGHT, d.GameObj.Angle)
	d.GameObj.Position = rl.Vector2Add(d.GameObj.Position, displacement)
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

	text := fmt.Sprintf("Docked with %s", d.DockedWith.Name)
	rl.DrawText(text, 10, 10, 20, rl.White)
}