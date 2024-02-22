package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	CurrScene string
	State State
}

func NewGame(title string, w, h int32) *Game {
	game := &Game{
		Title: title,
		Width: w,
		Height: h,
		State: Stopped,
		Scenes: make(map[string]*GameObj),
	}

	return game
}

func (g *Game) AddScene(name string, scene *GameObj) {
	g.Scenes[name] = scene
}

func (g *Game) SetScene(name string) {
	g.CurrScene = name
}

func (g *Game) Run() {
	g.State = Running

	scene := g.Scenes[g.CurrScene]
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{10, 10, 20, 255})

		scene.Update()
		CheckForCollisions(scene)
		scene.Draw()

		if DEBUG {
			game.Scenes[g.CurrScene].Profile(
				"ship",
				"planet",
				"moon",
				"deadly",
				"explosion",
				"shard")
		}
		
		rl.EndDrawing()

		// Check for scene change here
		// So that we're sure to finish drawing
		if scene != g.Scenes[g.CurrScene] {
			scene = g.Scenes[g.CurrScene]
		}
	}
}

func (g *Game) Pause() {
	g.State = Paused
}

func (g *Game) Stop() {
	g.State = Stopped
}

func CheckForCollisions(scope *GameObj) {
	objs := scope.FindChildrenByComponent(true, "area")

	// RESET
	for _, a := range objs {
		area := a.Components["area"].(*Area)
		area.Collided = false
	}

	// CHECK
	for i := 0; i < len(objs); i++ {
		this := objs[i]
		thisArea := this.Components["area"].(*Area)
		for j := i + 1; j < len(objs); j++ {
			that := objs[j]
		  	thatArea := that.Components["area"].(*Area)
		  	collision := thisArea.CollidedWith(that)
		  	if collision {
				thisArea.Collided = true
				thatArea.Collided = true

				// check for collision handlers on both objects!

				for tag, handler := range thisArea.CollisionHandlers {
					if that.HasTag(tag) {
						handler(this, that)
					}
				}

				for tag, handler := range thatArea.CollisionHandlers {
					if this.HasTag(tag) {
						handler(that, this)
					}
				}
		  	}
		}
	}
}

