package main

import (
	_ "embed"
	"log"
	"net/http"
	_ "net/http/pprof"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed systemSol.toml
var tomlData string

var textures map[string]rl.Texture2D
// var fonts map[string]rl.Font
var sounds map[string]rl.Sound

const (
	screenW = 600
	screenH = 800
	title = "Space Miner!"
	DEBUG = true
)

var game = NewGame(title, screenW, screenH)

func main() {

	if DEBUG {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	rl.InitWindow(screenW, screenH, title)
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()

	textures = LoadTextures(
		"assets/ship.png",
		"assets/rocky.png",
		"assets/earthy.png",
		"assets/gassy.png",
		"assets/cratery.png",
		"assets/rocky3.png",
		"assets/shadow.png",
		"assets/rings.png",
		"assets/shard.png",
		"assets/gameover.png",
		"assets/title.png",
		"assets/sun.png",
		"assets/icon-life.png",
	)

	sounds = LoadSounds(
		"sounds/thrust.wav",
		"sounds/explosion1.wav",
		"sounds/explosion2.wav",
		"sounds/explosion3.wav",
		"sounds/gameover.wav",
	)

	game.AddScene("titlescene", createTitleScene(game))
	game.Reset()
	// game.AddScene("level1", createLevel(game, tomlData))
	game.AddScene("gameover", createGameOverScene(game))

	// RUN!	
	game.SetScene("titlescene")
	game.Run()

}