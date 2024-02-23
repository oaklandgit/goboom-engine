package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 5
const THRUST_SPEED = 0.04
const MAX_SPEED = 2
const LIVES = 3
const WARNING_DISTANCE = 120
const SAFE_LANDING_SPEED = 0.5

func createShip(x, y float32) *GameObj {

	// SOUNDS
	thrustSound := sounds["sounds/thrust.wav"]
	rl.SetSoundVolume(thrustSound, 0.1);

	// SHIP METHODS
	thrust := func(g *GameObj, speed float32) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			g.Components["dock"].(*Dock).Undock()
		}
		g.Components["motion"].(*Motion).SetVelocity(speed, g.Angle)
		g.Components["sprite"].(*Sprite).CurrFrame = 1

		

		if !rl.IsSoundPlaying(thrustSound) {
			rl.PlaySound(thrustSound)
		}
		
	}

	stopThrust := func(g *GameObj) {
		g.Components["sprite"].(*Sprite).CurrFrame = 0
		rl.StopSound(thrustSound)
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
	ship.NewLives(LIVES)

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

	ship.NewApproach(
		[]string{"planet"},
		WithSafeDistance(WARNING_DISTANCE),
		WithSafeSpeed(SAFE_LANDING_SPEED ),
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
			"deadly",
			func(you *GameObj, thePlanet *GameObj) {

				// not a collision if it's the planet you're docked with
				if you.Components["dock"].(*Dock).DockedWith != nil &&
					you.Components["dock"].(*Dock).DockedWith == thePlanet { 
						return
				}

				ship.Parent.AddChildren(
					createExplosion(
						you.PosGlobal().X,
						you.PosGlobal().Y,
						"assets/shard.png",
						))
				fmt.Printf("BOOM! You crashed with %s\n", thePlanet.Name)
				you.Components["lives"].(*Lives).RemoveLife()
			})

	landingZone.Components["area"].(*Area).AddCollisionHandler(
		"planet",
		func(you *GameObj, thePlanet *GameObj) {

			if ship.Components["approach"].(*Approach).IsSafeSpeed() {
				dockWith(ship, thePlanet, landingZone.PosGlobal())
			}

			// no need to deal with the landingzone collision
			// after this point, because the ship will hit the planet
			// and trigger its own collision.
					
		})

	return ship

}