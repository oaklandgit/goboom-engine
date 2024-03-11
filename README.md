A simple composition-based game engine in Go and [Raylib](https://github.com/gen2brain/raylib-go). Roughly inspired by the amazing [Kaboom!](https://kaboomjs.com/) engine for javascript.

## Example

```go
package main

import (
	gb "goboom"
)

var game = gb.NewGame(
	"Hello World",
	600,
	800,
	true, // debug mode
)

func init() {

	game.Reset = func() {}
	game.LoadTextures("assets/ship.png")

	ship := game.	NewGameObject("ship", gb.WithPosition(300, 400)).
					NewSprite(game.Textures["assets/ship.png"])

	game.AddScene("myscene", ship)
	game.SetScene("myscene")
}

func main() {
	game.Run()
}
```
