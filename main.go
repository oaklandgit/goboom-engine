package main

import (
	_ "embed"
)

//go:embed systemSol.toml
var tomlData string

const (
	screenW = 600
	screenH = 800
	title = "Space Miner!"
	DEBUG = false
)

var game = NewGame(title, screenW, screenH)

func init() {

	game.LoadTextures(
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

	game.LoadSounds(
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
	game.AddScene("gameover", createGameOverScene(game))
	
	game.Reset = func() {
		game.AddScene("level1", createStarSystem(game, tomlData))
	}
	
	game.SetScene("titlescene")
}

func main() {

	game.Run()

}