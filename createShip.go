package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 5
const THRUST_SPEED = 0.04
const MAX_SPEED = 2	

func createShip(x, y float32) *GameObj {

	// SHIP METHODS
	thrust := func(g *GameObj, speed float32) {
		g.Components["motion"].(*Motion).SetVelocity(speed, g.Angle)
		g.Components["sprite"].(*Sprite).CurrFrame = 1
	}

	stopThrust := func(g *GameObj) {
		g.Components["sprite"].(*Sprite).CurrFrame = 0
	}

	dockWith := func(g *GameObj, thePlanet *GameObj) {
		g.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
		g.Tags = append(g.Tags, "docked")
		thePlanet.Tags = append(thePlanet.Tags, "docking")
		fmt.Printf("Landed on %s!\n", thePlanet.Name)
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

	landingZone := NewGameObject(
		"Landing Zone",
		WithPosition(-12, 0),
	)
	landingZone.NewArea(CircleCollider{Radius: 3})
	ship.AddChildren(landingZone)

	rotateLandingWithShip := func() {
		landingZone.Position.X = float32(-12 * math.Cos(float64(landingZone.Angle * rl.Deg2rad)))
		landingZone.Position.Y = float32(-12 * math.Sin(float64(landingZone.Angle * rl.Deg2rad)))
	}


	rotateCW := func(g *GameObj) {
		g.Angle += ROTATE_SPEED
		landingZone.Angle += ROTATE_SPEED
		rotateLandingWithShip()
	}

	rotateCCW := func(g *GameObj) {
		g.Angle -= ROTATE_SPEED
		landingZone.Angle -= ROTATE_SPEED
		rotateLandingWithShip()
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
			dockWith(ship, thePlanet)			
		})

	return ship

}