package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var textures map[string]rl.Texture2D
// var fonts map[string]rl.Font
var sounds map[string]rl.Sound

const (
	screenW = 800
	screenH = 450
	title = "Space Miner!"
	DEBUG = false
)

var game = NewGame(title, screenW, screenH, textures)

func main() {

	rl.InitWindow(screenW, screenH, title)
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	textures = LoadTextures(
		"assets/ship.png",
		"assets/rocky.png",
		"assets/shadow.png",
		"assets/rings.png",
		"assets/shard.png",
	)

	sounds = LoadSounds(
		"sounds/thrust.wav",
		"sounds/explosion1.wav",
		"sounds/explosion2.wav",
		"sounds/explosion3.wav",
		"sounds/gameover.wav",
	)

	// fonts = LoadFonts(
	// 	"assets/Monocraft.ttf",
	// )

	// FONT
	// fontPath := "assets/Monocraft.ttf"
	// fontSize := int32(16)
	// customFont := rl.LoadFontEx(fontPath, fontSize, nil)

	// SHIP
	ship := createShip(400, 120)

	// STARS
	starfield := NewGameObject("Starfield")
	starfield.NewStarfield(screenW, screenH, 40)

	// PLANETS
	var level1Map = `
	..........
	..........
	🪐.........
	.........🌎
	..🔴.......
	..........
	`

	earth := createPlanet("Earth", 0, 0, 0.2, -0.1, 0, 60, rl.Blue, 0.2, ship, 1)
	earth.AddChildren(
		createMoon("Moon", 0.4, 0.4, 20, 112, rl.RayWhite, 1),
	)
	earth.Components["mine"].(*Mine).
		AddResource("gold", 100, 1000).
		AddResource("silver", 200, 500).
		AddResource("copper", 300, 200).
		AddResource("iron", 400, 100)

	mars := createPlanet("Mars", 0, 0, 0.3, 0.6, 0, 40, rl.Red, 0.15, ship, 1)
	mars.AddChildren(
		createMoon("Phobos", -1.3, 3, 4, 50, rl.Pink, 1),
		createMoon("Deimos", -1, 0.1, 8, 100, rl.Pink, 1),
	)
	mars.Components["mine"].(*Mine).
		AddResource("lithium", 100, 1000).
		AddResource("uranium", 200, 500).
		AddResource("plutonium", 300, 200)

	saturn := createPlanet("Saturn", 0, 0, -0.1, 0.2, 0, 80, rl.Yellow, 0.1, ship, 1)
	
	saturn.Components["mine"].(*Mine).
		AddResource("diamond", 100, 1000).
		AddResource("ruby", 200, 500).
		AddResource("sapphire", 300, 200)
	
	
	rings := NewGameObject("Rings", WithScale(2.4, 2.4), WithAngle(30))
	rings.NewSprite(
		textures["assets/rings.png"],
		WithOpacity(0.2),
		WithColor(rl.Yellow),
	)
	saturn.AddChildren(rings)
		var level1MapTable = map[rune]func() *GameObj{
		'🪐': func() *GameObj {
			return saturn
		},
		'🌎': func() *GameObj {
			return earth
		},
		'🔴': func() *GameObj {
			return mars
		},

	}
	solarSystem := CreateLevel(
		"The Solar System",
		level1Map,
		60, 60,
		level1MapTable,
	)
	solarSystem.Size = rl.NewVector2(screenW, screenH)

	// SCENE
	scene1 := NewGameObject("Scene 1")
	scene1.Size = rl.NewVector2(screenW, screenH)
	scene1.AddChildren(starfield, solarSystem, ship)


	game.AddScene("level1", scene1)
	game.Run("level1")

}