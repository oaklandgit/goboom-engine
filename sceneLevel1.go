package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	cellW = 60
	cellH = 60
)

func findPos(chart string, code rune, lineW int) (row, col int) {
	for i, char := range chart {
		if char == code {
			row = i / lineW
			col = i % lineW
			return
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
}

func createLevel1Scene(g *Game) *GameObj {

	// load toml file
	var sol System
    _, err := toml.DecodeFile("systemSol.toml", &sol)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Printf("Title: %s\n", sol.Name)
	fmt.Printf("Suns: %d\n", sol.Suns)
	fmt.Printf("Planets: %v\n", sol.Planets["earth"])
	fmt.Printf("Chart: %v\n", sol.Chart)

	// SHIP
	ship := createShip(400, 120)

	// STARS
	starfield := NewGameObject("Starfield")
	starfield.NewStarfield(screenW, screenH, 40)

	// SOLAR SYSTEM
	solarSystem := NewGameObject("Solar System")
	solarSystem.Size = rl.NewVector2(screenW, screenH)

	for _, p := range sol.Planets {

		color := rl.NewColor(
			uint8(p.Color[0]),
			uint8(p.Color[1]),
			uint8(p.Color[2]),
			255,
		)

		// find its position in the chart

		symbol := rune(p.Symbol[0])
		
		posX, posY := findPos(sol.Chart, symbol, 8)

		planet := createPlanet(
			p.Name,
			float32(posX * cellW),
			float32(posY * cellH),
			p.Speed,
			p.Rotation,
			p.Heading,
			p.Radius,
			color,
			p.Gravity,
			ship,
			1,
		)

		mine := planet.NewMine()

		for _, product := range p.Products {
			mine.AddResource(product.Name, product.Amount, product.Value)
		}

		solarSystem.AddChildren(planet)
	}

	// earth := createPlanet("Earth", 0, 0, 0.2, -0.1, 0, 60, rl.Blue, 0.2, ship, 1)
	// earth.AddChildren(
	// 	createMoon("Moon", 0.4, 0.4, 14, 112, rl.Fade(rl.RayWhite, 0.5), 1),
	// )
	// earth.Components["mine"].(*Mine).
	// 	AddResource("gold", 100, 1000).
	// 	AddResource("silver", 200, 500).
	// 	AddResource("copper", 300, 200).
	// 	AddResource("iron", 400, 100)

	// mars := createPlanet("Mars", 0, 0, 0.3, 0.6, 0, 40, rl.Fade(rl.Red, 0.7), 0.15, ship, 1)
	// mars.AddChildren(
	// 	createMoon("Phobos", -1.3, 3, 4, 50, rl.Fade(rl.Pink, 0.6), 1),
	// 	createMoon("Deimos", -1, 0.1, 8, 100, rl.Fade(rl.Pink, 0.6), 1),
	// )
	// mars.Components["mine"].(*Mine).
	// 	AddResource("lithium", 100, 1000).
	// 	AddResource("uranium", 200, 500).
	// 	AddResource("plutonium", 300, 200)

	// saturn := createPlanet("Saturn", 0, 0, -0.1, 0.2, 0, 80, rl.Fade(rl.Yellow, 0.6), 0.1, ship, 1)
	
	// saturn.Components["mine"].(*Mine).
	// 	AddResource("diamond", 100, 1000).
	// 	AddResource("ruby", 200, 500).
	// 	AddResource("sapphire", 300, 200)
	
	
	// rings := NewGameObject("Rings", WithScale(2.4, 2.4), WithAngle(30))
	// rings.NewSprite(
	// 	textures["assets/rings.png"],
	// 	WithOpacity(0.2),
	// 	WithColor(rl.Yellow),
	// )
	// saturn.AddChildren(rings)
	// 	var level1MapTable = map[rune]func() *GameObj{
	// 	'🪐': func() *GameObj {
	// 		return saturn
	// 	},
	// 	'🌎': func() *GameObj {
	// 		return earth
	// 	},
	// 	'🔴': func() *GameObj {
	// 		return mars
	// 	},

	// }
	// solarSystem := CreateLevel(
	// 	"The Solar System",
	// 	level1Map,
	// 	60, 60,
	// 	level1MapTable,
	// )
	// solarSystem.Size = rl.NewVector2(screenW, screenH)

	scene1 := NewGameObject("Scene 1")
	scene1.Size = rl.NewVector2(screenW, screenH)
	scene1.AddChildren(starfield, solarSystem, ship)

	return scene1
}