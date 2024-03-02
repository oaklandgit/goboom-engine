Building an arcade-style game in Go, with the hopes of generalizing the engine for other games. A work in progress, learning as I… go. Using Raylib go bindings for graphics and input.

<video width="300" height="600" controls>
  <source src="demo/gameplay.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

# IN PROGRESS

- Tweening for pulse behavior (e.g. when losing life icon or respawning ship)

# TO DO

- Sprite animations (separate component from sprite?) Use go routine / generator for sprite animations
- For new components, return game object to allow chaining
- Enemy ships try to take the resources and shoot at you
- OK maybe some shooting would be good
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
