Building an arcade-style game in Go, with the goal of generalizing the engine for other games.

This is a work in progress, learning as I… go 😉. Using [Raylib go bindings](https://github.com/gen2brain/raylib-go) for graphics and input.

Roughly inspired by the amazing [Kaboom!](https://kaboomjs.com/) engine for javascript.

So far, the engine contains (all rudimentary):

- a scene graph (parent/child relationships between game objects)
- a component system for adding behaviors to game objects (e.g. input handling, collision handling, movement, sprites, and custom components etc.)

https://github.com/oaklandgit/SpaceMiner2/assets/421615/18bf24ba-b01a-49b5-9117-466ab46435b9
