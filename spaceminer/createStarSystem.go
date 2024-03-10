package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"

	boom "goboom"

	"github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	RING_TO_PLANET_SIZE_RATIO = 5.0
	SHADOW_TO_PLANET_SIZE_RATIO = 0.015
	STARFIELD_DENSITY = 60
	SUN_POS_Y_OFFSET = 260
)

type System struct {
	Name string
	Suns int
	Chart string
	Planets map[string]Planet
}

type Product struct {
	Name string
	Amount int
	Value int
}

type Satellite struct {
	Name string
	Texture string
	Radius float32
	Speed float32
	Rotation float32
	Distance float32
	Color []int
}

type Planet struct {
	Name string
	Symbol string
	Pos []float32
	Distance float32 // from its sun
	Texture string
	Radius float32
	Speed float32
	Rotation float32
	Color []int
	Gravity float32
	Products map[string]Product
	Satellites map[string]Satellite
	HasRings bool
	InitialAngle float32
}

func createStarSystem(g *boom.Game, tomlStr string) *boom.GameObj {

	// load toml file
	var system System
    _, err := toml.Decode(tomlStr, &system)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

	// SHIP ////////////////////
	ship := createShip(200, 200)
	
	// STARS ///////////////////
	starfield := boom.NewGameObject("Starfield")
	NewStarfield(starfield, screenW, screenH, STARFIELD_DENSITY)

	// STAR SYSTEM //////////////
	starSystem := boom.NewGameObject("Solar System", boom.WithScale(2, 2))
	starSystem.Size = rl.NewVector2(screenW, screenH)
	starSystem.NewSprite(game.Textures["assets/sun.png"], boom.WithOpacity(0.4))

	// star is located at bottom center of screen
	starSystem.Position = rl.NewVector2(screenW/2, screenH - SUN_POS_Y_OFFSET)

	for _, p := range system.Planets {

		planetColor := rl.NewColor(
			uint8(p.Color[0]),
			uint8(p.Color[1]),
			uint8(p.Color[2]),
			255,
		)

		planetTex := game.Textures[fmt.Sprintf("assets/%s.png", p.Texture)]

		planet := createPlanet(
			p.Name,
			planetTex,
			0, // temporary
			0,
			p.Speed,
			p.Rotation,
			p.Radius,
			planetColor,
			p.Gravity,
			ship,
		)

		if p.HasRings {

			ringTex := game.Textures["assets/rings.png"]
			ringW := ringTex.Width
			ringScale := p.Radius *
				RING_TO_PLANET_SIZE_RATIO / float32(ringW)
			ringAngle := rand.Float32() * 30

			rings := boom.NewGameObject(
				"Rings",
				boom.WithAngle(ringAngle),
				boom.WithScale(ringScale, ringScale),
			)
			rings.NewSprite(
				game.Textures["assets/rings.png"],
				boom.WithOpacity(0.3),
				boom.WithColor(planetColor),
			)
			planet.AddChildren(rings)
		}

		for _, s := range p.Satellites {

			satColor := rl.NewColor(
				uint8(s.Color[0]),
				uint8(s.Color[1]),
				uint8(s.Color[2]),
				255,
			)

			moonTex := game.Textures[fmt.Sprintf("assets/%s.png", s.Texture)]

			planet.AddChildren(
				createMoon(
					s.Name,
					moonTex,
					s.Speed,
					s.Rotation,
					s.Radius,
					s.Distance,
					satColor,
					1,
				),
			)

		}

		// mine := planet.NewMine()
		NewMine(planet)

		for _, product := range p.Products {
			planet.Components["mine"].(*Mine).
				AddResource(product.Name, product.Amount, product.Value)
		}

		// add shadow
		shadow := boom.NewGameObject("Shadow",
			boom.WithScale(
				p.Radius * SHADOW_TO_PLANET_SIZE_RATIO,
				p.Radius * SHADOW_TO_PLANET_SIZE_RATIO))

		shadow.NewSprite(game.Textures["assets/shadow.png"],
			boom.WithOpacity(0.8))

		shadow.NewPointAt(starSystem)
		planet.AddChildren(shadow)

		planet.NewOrbit(p.Speed, p.Distance, boom.WithOrbitAngle(p.InitialAngle))
		starSystem.AddChildren(planet)
	}

	// enemies
	enemy :=	boom.NewGameObject("Enemy1",
					boom.WithTags("enemy"),
					boom.WithScale(0.5, 0.5),
					boom.WithPosition(400, 400)).
				NewMotion(
					boom.WithVelocity(2, 0),
					boom.WithWrap(true, true, 20)).
				NewSprite(game.Textures["assets/ufo.png"],
					boom.WithFrames(1, 5, 4)).
				NewArea(boom.CircleCollider{Radius: 20}).
				NewPointAt(starSystem)

	enemy.Components["sprite"].(*boom.Sprite).NewAnimation("open", 0, 4, 8, true)
	enemy.Components["sprite"].(*boom.Sprite).Play("open")
		

	scene1 := boom.NewGameObject("Scene 1")
	scene1.Size = rl.NewVector2(screenW, screenH)
	scene1.AddChildren(starfield, starSystem, enemy, ship)

	return scene1
}