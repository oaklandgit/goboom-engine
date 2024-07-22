package main

import (
	"time"

	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ROTATE_SPEED       = 5
	THRUST_SPEED       = 0.04
	MAX_SPEED          = 2
	LIVES              = 3
	WARNING_DISTANCE   = 120
	SAFE_LANDING_SPEED = 1
	SHIP_RADIUS        = 8
	SHIP_DRAG          = 0.999
	SHIP_WRAP_PADDING  = 16

	FLYING        = 2
	FLYING_THRUST = 3

	LANDING        = 0
	LANDING_THRUST = 1
)

func createShip(x, y float32) *boom.GameObj {

	// SOUNDS
	thrustSound := game.Sounds["sounds/thrust.wav"]
	rl.SetSoundVolume(thrustSound, 0.3)

	// SHIP METHODS
	thrust := func(g *boom.GameObj, speed float32) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			g.Components["dock"].(*Dock).Undock()
		}
		g.Components["motion"].(*boom.Motion).SetVelocity(speed, g.Angle)
		g.Components["sprite"].(*boom.Sprite).CurrFrame = FLYING_THRUST

		if !rl.IsSoundPlaying(thrustSound) {
			rl.PlaySound(thrustSound)
		}

	}

	stopThrust := func(g *boom.GameObj) {
		g.Components["sprite"].(*boom.Sprite).CurrFrame = FLYING
		rl.StopSound(thrustSound)
	}

	dockWith := func(g *boom.GameObj, theObject *boom.GameObj, landingPosition rl.Vector2) {
		g.Components["dock"].(*Dock).DockWith(theObject, landingPosition)
	}

	// SHIP
	ship := game.NewGameObject("Spaceship",
		boom.WithPosition(x, y),
		boom.WithOrigin(0.7, 0.5),
		boom.WithTags("ship"),
		boom.WithScale(0.4, 0.4),
	)

	ship.NewSprite(
		game.Textures["assets/lander.png"],
		boom.WithFrames(2, 2, 4),
	)

	ship.Components["sprite"].(*boom.Sprite).CurrFrame = 2

	NewDock(ship)
	NewBank(ship)
	NewLives(ship, LIVES)

	rotateCW := func(g *boom.GameObj) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			return
		}
		g.Angle += ROTATE_SPEED
	}

	rotateCCW := func(g *boom.GameObj) {
		if g.Components["dock"].(*Dock).DockedWith != nil {
			return
		}
		g.Angle -= ROTATE_SPEED
	}

	ship.NewArea(boom.CircleCollider{Radius: SHIP_RADIUS},
		boom.WithCooldown(2*time.Second),
	)

	ship.NewMotion(
		boom.WithFriction(SHIP_DRAG),
		boom.WithMaxVelocity(MAX_SPEED),
		boom.WithWrap(true, true, SHIP_WRAP_PADDING),
	)

	NewApproach(
		ship,
		[]string{"planet"},
		WithSafeDistance(WARNING_DISTANCE),
		WithSafeSpeed(SAFE_LANDING_SPEED),
	)

	ship.NewInput(
		boom.KeyHandler{
			KeyPress: boom.KeyPress{Key: rl.KeyLeft, Mode: boom.KEY_REPEAT},
			Action: func() {
				rotateCCW(ship)
			},
		},
		boom.KeyHandler{
			KeyPress: boom.KeyPress{Key: rl.KeyRight, Mode: boom.KEY_REPEAT},
			Action: func() {
				rotateCW(ship)
			},
		},
		boom.KeyHandler{
			KeyPress: boom.KeyPress{Key: rl.KeyUp, Mode: boom.KEY_REPEAT},
			Action: func() {
				thrust(ship, THRUST_SPEED)
			},
		},
		boom.KeyHandler{
			KeyPress: boom.KeyPress{Key: rl.KeyUp, Mode: boom.KEY_UP},
			Action: func() {
				stopThrust(ship)
			},
		},
	)

	ship.Components["area"].(*boom.Area).AddCollisionHandler(
		"deadly",
		func(you, theObject *boom.GameObj) {

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
