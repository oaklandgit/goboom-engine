package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 5
const THRUST_SPEED = 0.04
const MAX_SPEED = 2	


func createShip(x, y float32) *GameObj {

	// SHIP METHODS
	thrust := func(g *GameObj, speed float32) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			g.Components["dock"].(*Dock).Undock()
		}
		g.Components["motion"].(*Motion).SetVelocity(speed, g.Angle)
		g.Components["sprite"].(*Sprite).CurrFrame = 1
	}

	stopThrust := func(g *GameObj) {
		g.Components["sprite"].(*Sprite).CurrFrame = 0
	}

	dockWith := func(g *GameObj, thePlanet *GameObj, landingPosition rl.Vector2) {
		// landingPosition is the angle of where on the edge it landed
		g.Components["dock"].(*Dock).DockWith(thePlanet, landingPosition)
	}

	// SHIP
	ship := NewGameObject("Spaceship",
		WithPosition(x, y),
		WithOrigin(0.6, 0.5),
		WithTags("ship"),
	)

	ship.NewSprite(
		textures["assets/ship.png"],
		WithFrames(1, 2, 2),
	)

	ship.NewDock()
	ship.NewBank()

	landingZone := NewGameObject(
		"Landing Zone",
		WithPosition(-12, 0),
		WithTags("landingZone"),
	)
	landingZone.NewArea(CircleCollider{Radius: 3})
	ship.AddChildren(landingZone)

	rotateCW := func(g *GameObj) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			return
		}
		g.Angle += ROTATE_SPEED
	}

	rotateCCW := func(g *GameObj) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			return
		}
		g.Angle -= ROTATE_SPEED
	}

	ship.NewArea(CircleCollider{Radius: 8})

	ship.NewMotion(
		WithFriction(0.999),
		WithMaxVelocity(MAX_SPEED),
		WithWrap(true, true, 16),
	)
	ship.NewInput(
		KeyHandler{
			KeyPress{rl.KeyLeft, KEY_REPEAT},
			func() {
				rotateCCW(ship)
			},
		},
		KeyHandler{
			KeyPress{rl.KeyRight, KEY_REPEAT},
			func() {
				rotateCW(ship)
			},
		},
		KeyHandler{
			KeyPress{rl.KeyUp, KEY_REPEAT},
			func() {
				thrust(ship, THRUST_SPEED)
			},
		},
		KeyHandler{
			KeyPress{rl.KeyUp, KEY_UP},
			func() {
				stopThrust(ship)
			},
		},
		
	)

	ship.Components["area"].(*Area).AddCollisionHandler(
			"planet",
			func(you *GameObj, thePlanet *GameObj) {
				if !thePlanet.HasTag("docking") {
					fmt.Printf("BOOM! You crashed with %s\n", thePlanet.Name)
				}
			})

	landingZone.Components["area"].(*Area).AddCollisionHandler(
		"planet",
		func(you *GameObj, thePlanet *GameObj) {
			dockWith(ship, thePlanet, landingZone.PosGlobal())			
		})

	return ship

}