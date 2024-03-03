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
	DEBUG = false
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
		"assets/enemy.png",
		"assets/ufo.png",
		"assets/rocky.png",
		"assets/earthy.png",
		"assets/gassy.png",
		"assets/cratery.png",
		"assets/scarry.png",
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
		"sounds/music.wav",
		"sounds/thrust.wav",
		"sounds/dock.wav",
		"sounds/collected.wav",
		"sounds/undock.wav",
		"sounds/explosion1.wav",
		"sounds/explosion2.wav",
		"sounds/explosion3.wav",
		"sounds/explosion4.wav",
		"sounds/explosion5.wav",
		"sounds/explosion6.wav",
		"sounds/explosion7.wav",
		"sounds/gameover.wav",
	)

	game.AddScene("titlescene", createTitleScene(game))
	game.Reset()
	game.AddScene("gameover", createGameOverScene(game))

	// RUN!	
	game.SetScene("titlescene")
	game.Run()

}