package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ROTATE_SPEED = 5
	THRUST_SPEED = 0.04
	MAX_SPEED = 2
	LIVES = 3
	WARNING_DISTANCE = 120
	SAFE_LANDING_SPEED = 1
	SHIP_RADIUS = 8
	SHIP_DRAG = 0.999
	SHIP_WRAP_PADDING = 16

	FLYING = 2
	FLYING_THRUST = 3

	LANDING = 0
	LANDING_THRUST = 1
)

func createShip(x, y float32) *GameObj {

	// SOUNDS
	thrustSound := game.Sounds["sounds/thrust.wav"]
	rl.SetSoundVolume(thrustSound, 0.3);

	// SHIP METHODS
	thrust := func(g *GameObj, speed float32) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			g.Components["dock"].(*Dock).Undock()
		}
		g.Components["motion"].(*Motion).SetVelocity(speed, g.Angle)
		g.Components["sprite"].(*Sprite).CurrFrame = FLYING_THRUST

		if !rl.IsSoundPlaying(thrustSound) {
			rl.PlaySound(thrustSound)
		}
		
	}

	stopThrust := func(g *GameObj) {
		g.Components["sprite"].(*Sprite).CurrFrame = FLYING
		rl.StopSound(thrustSound)
	}

	dockWith := func(g *GameObj, theObject *GameObj, landingPosition rl.Vector2) {
		g.Components["dock"].(*Dock).DockWith(theObject, landingPosition)
	}

	// SHIP
	ship := NewGameObject("Spaceship",
		WithPosition(x, y),
		WithOrigin(0.7, 0.5),
		WithTags("ship"),
		WithScale(0.4, 0.4),
	)

	ship.NewSprite(
		game.Textures["assets/lander.png"],
		WithFrames(2, 2, 4),
	)

	ship.Components["sprite"].(*Sprite).CurrFrame = 2

	ship.NewDock()
	ship.NewBank()
	ship.NewLives(LIVES)

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

	ship.NewArea(CircleCollider{Radius: SHIP_RADIUS},
		WithCooldown(2 * time.Second),
	)

	ship.NewMotion(
		WithFriction(SHIP_DRAG),
		WithMaxVelocity(MAX_SPEED),
		WithWrap(true, true, SHIP_WRAP_PADDING),
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
			func(you, theObject *GameObj) {

				// // don't crash if docked with this planet
				if you.Components["dock"].(*Dock).DockedWith != nil &&
					you.Components["dock"].(*Dock).DockedWith == theObject { 
						return
				}

				// // land if good speed and angle
				if you.Components["approach"].(*Approach).IsSafeSpeed() &&
					!you.Components["approach"].(*Approach).IsPointingToward(theObject) {
						
					dockWith(you, theObject, rl.Vector2{})
					return
				}

				// otherwise, crash
				you.Parent.AddChildren(
					createExplosion(
						you.PosGlobal().X,
						you.PosGlobal().Y,
						"assets/shard.png",
						))
				you.Components["lives"].(*Lives).RemoveLife()
			})

	return ship

}