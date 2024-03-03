Building an arcade-style game in Go, with the goal of generalizing the engine for other games.

This is a work in progress, learning as I… go 😉. Using [Raylib go bindings](https://github.com/gen2brain/raylib-go) for graphics and input.

Roughly inspired by the amazing [Kaboom!](https://kaboomjs.com/) engine for javascript.

So far, the engine contains (all rudimentary):

- a scene graph (parent/child relationships between game objects)
- a component system for adding behaviors to game objects (e.g. input handling, collision handling, movement, sprites, and custom components etc.)

https://github.com/oaklandgit/SpaceMiner2/assets/421615/18bf24ba-b01a-49b5-9117-466ab46435b9

# TO DO

- Fix failure to register collision when another object overlaps
- Enemy ships try to take the resources and shoot at you
- Expand Tween capability
- Use tweens for pulse behavior (e.g. when losing life icon or respawning ship)
- Somehow indicate what resources a planet has before docking
- Sometimes crashes (the program, not the ship) after crash landing
- Move any utility functions to a separate file
- Fix ship being pulled back to planet on undocking
- Collisions should not "cooldown" if collision is between two planets
  - in fact, collisions should be ignored between planets
- Delay before respawn
- embed textures in binary
- embed sounds in binary
- Test stand-alone binary
- Buggy when docked and moving past left or right edge of screen
- Leverage go time package for time-based events
- Moons and planets should be combined into a single struct

# JUICINESS

- Different planet textures (e.g. Rocky, gaseous, watery, earthlike)
- Game starts with instructions on objective and how to land
- Each scene starts with an intro to its solar system
- Mini map of the full solar system
- Animate touching down tutorial
- Artificial satellites
- worm holes (teleportation)
- more levels

# DONE

- For new components, return game object to allow chaining
- Sprite animations using go routine / generator
- Planets orbit their star (as opposed to just gliding across the screen)
- Docking sound
- Ship should attach to planet at correct landing point
- Don't switch target planets - only set when empty
- Play again
- Reset all variables on game start
- embed toml in binary
- More feedback on state of landing (e.g. too fast!)
- Move planet, moon and ore info into TOML file
- Collision with planet if too-fast or wrong-angle landing
- organize scenes into files
- Game Start state
- Game Over state (and scene switching generalized)
- Deal with memory leaks
  - Lifespan component // use for explosions
  - When toggling debug mode, print scene graph to console
- thrust and explosion sounds
- Explosion and death on wrong landing
- Mult lives / respawing
- Planets (and moons) should have radius param, not scale
- Depletion of resources reverse progress bars
- HUD showing docked planet and its resources
- attach ship to planet if landing zone collides
- collision detection
- sprite frames: flame on thrust
- rotation relative to parent // done for ship but not generalized
