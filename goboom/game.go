package goboom

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
	Width float32
	Height float32
	Scenes map[string]*GameObj
	CurrScene string
	State State
	Debug bool
	Textures map[string]rl.Texture2D
	Sounds map[string]rl.Sound
	Soundtrack string
	Fonts map[string]rl.Font
	Reset func()
}

func NewGame(title string, w, h float32, debug bool) *Game {
	game := &Game{
		Title: title,
		Width: w,
		Height: h,
		State: Stopped,
		Scenes: make(map[string]*GameObj),
		Debug: debug,
	}

	rl.InitWindow(int32(game.Width), int32(game.Height), title)
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	return game
}

func (g *Game) LoadTextures(paths ...string) {
	textureMap := make(map[string]rl.Texture2D)

	for _, path := range paths {
		textureMap[path] = rl.LoadTexture(path)
	}

	g.Textures = textureMap
}

func (g *Game) LoadSounds(paths ...string) {
	soundMap := make(map[string]rl.Sound)

	for _, path := range paths {
		soundMap[path] = rl.LoadSound(path)
	}

	g.Sounds = soundMap
}

func (g *Game) LoadFonts(paths ...string) {
	fontMap := make(map[string]rl.Font)

	for _, path := range paths {
		fontMap[path] = rl.LoadFont(path)
	}

	g.Fonts = fontMap
}

func (g *Game) AddScene(name string, scene *GameObj) {
	g.Scenes[name] = scene
}

func (g *Game) SetScene(name string) {
	g.CurrScene = name
}

func (game *Game) NewGameObject(name string, opts ...GameObjOption) *GameObj {

	obj := &GameObj{
		Name: name,
		Game: game,
		Scale: rl.Vector2{X: 1, Y: 1},
		Origin: rl.Vector2{X: 0.5, Y: 0.5},
		Components: make(map[string]Component),
		Tags: make(map[string]struct{}),
		Deleted: false,
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}

func (g *Game) Run() {

	g.Reset()
	g.State = Running

	scene := g.Scenes[g.CurrScene]
	soundtrack := rl.LoadMusicStream(g.Soundtrack)

	rl.PlayMusicStream(soundtrack)
	rl.SetMusicVolume(soundtrack, 0.4)

	for !rl.WindowShouldClose() {

		rl.UpdateMusicStream(soundtrack)
		if !rl.IsMusicStreamPlaying(soundtrack) {
			rl.PlayMusicStream(soundtrack)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{10, 10, 20, 255})

		scene.Update()
		checkForCollisions(scene)
		scene.Draw()
		
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

func checkForCollisions(scope *GameObj) {
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
						// cooldown here?
					}
				}

				for tag, handler := range thatArea.CollisionHandlers {
					if this.HasTag(tag) {
						handler(that, this)
						// cooldown here?
					}
				}
		  	}
		}
	}
}

