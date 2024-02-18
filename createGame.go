package main

import rl "github.com/gen2brain/raylib-go/raylib"

type State int

const (
	Running State = iota
	Paused
	Stopped
)

type Game struct {
	Title string
	Width int32
	Height int32
	Scenes map[string]*GameObj
	// CurrScene *GameObj
	State State
	// Textures []string
	// Textures map[string]rl.Texture2D
}

func NewGame(title string, w, h int32, textures map[string]rl.Texture2D) *Game {
	game := &Game{
		Title: title,
		Width: w,
		Height: h,
		State: Stopped,
		Scenes: make(map[string]*GameObj),
		// Textures: textures,
	}

	return game
}

func (g *Game) AddScene(name string, scene *GameObj) {
	g.Scenes[name] = scene
}

func (g *Game) Run(scene string) {
	g.State = Running

	// rl.InitWindow(g.Width, g.Height, g.Title)
	// rl.SetTargetFPS(60)

	// textures = LoadTextures("assets/planet.png", "assets/ship.png")
	// textures = LoadTextures(g.Textures...)
	
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		game.Scenes[scene].Update()
		game.Scenes[scene].Draw()
		
		rl.EndDrawing()
	}
}

func (g *Game) Pause() {
	g.State = Paused
}

func (g *Game) Stop() {
	g.State = Stopped
}

