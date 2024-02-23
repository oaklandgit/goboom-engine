package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CHART_CELL_SIZE = 60
	RING_TO_PLANET_SIZE_RATIO = 3.8
)

func findPos(chart string, code rune, size int) (x, y int) {

	rows := strings.Split(chart, "\n")

	for i, row := range rows {
		col := strings.IndexRune(row, code)
		if col != -1 {
			return i * size, col * size
		}
	}

	return -1, -1

}

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
	Texture string
	Radius float32
	Speed float32
	Heading float32
	Rotation float32
	Color []int
	Gravity float32
	Products map[string]Product
	Satellites map[string]Satellite
	HasRings bool
}

func createLevel1Scene(g *Game) *GameObj {

	// load toml file
	var system System
    _, err := toml.DecodeFile("systemSol.toml", &system)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

	// SHIP ////////////////////
	shipY, shipX := findPos(system.Chart, rune('>'), CHART_CELL_SIZE)
	ship := createShip(float32(shipX), float32(shipY))
	
	// STARS ///////////////////
	starfield := NewGameObject("Starfield")
	starfield.NewStarfield(screenW, screenH, 40)

	// STAR SYSTEM //////////////
	starSystem := NewGameObject("Solar System")
	starSystem.Size = rl.NewVector2(screenW, screenH)

	for _, p := range system.Planets {

		planetColor := rl.NewColor(
			uint8(p.Color[0]),
			uint8(p.Color[1]),
			uint8(p.Color[2]),
			uint8(255),
		)

		// find its position in the chart

		symbol := rune(p.Symbol[0])
		posY, posX := findPos(system.Chart, symbol, CHART_CELL_SIZE)

		planet := createPlanet(
			p.Name,
			float32(posX),
			float32(posY),
			p.Speed,
			p.Rotation,
			p.Heading,
			p.Radius,
			planetColor,
			p.Gravity,
			ship,
		)

		if p.HasRings {

			fmt.Println("Creating rings")

			ringTex := textures["assets/rings.png"]
			ringW := ringTex.Width
			ringScale := p.Radius *
				RING_TO_PLANET_SIZE_RATIO / float32(ringW)
			ringAngle := rand.Float32() * 45

			rings := NewGameObject(
				"Rings",
				WithAngle(ringAngle),
				WithScale(ringScale, ringScale),
			)
			rings.NewSprite(
				textures["assets/rings.png"],
				WithOpacity(0.2),
				WithColor(planetColor),
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

			planet.AddChildren(
				createMoon(
					s.Name,
					s.Speed,
					s.Rotation,
					s.Radius,
					s.Distance,
					satColor,
					1,
				),
			)

		}


		mine := planet.NewMine()

		for _, product := range p.Products {
			mine.AddResource(product.Name, product.Amount, product.Value)
		}

		starSystem.AddChildren(planet)
	}

	scene1 := NewGameObject("Scene 1")
	scene1.Size = rl.NewVector2(screenW, screenH)
	scene1.AddChildren(starfield, starSystem, ship)

	return scene1
}