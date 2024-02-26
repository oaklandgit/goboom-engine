# IN PROGRESS

- Show resources on HUD when close to target planet

# TO DO

- Generalized pulse behavior (e.g. when losing life icon or respawning ship)
- Sometimes crashes (the program, not the ship) after crash landing
- Ship should attach to planet at correct landing point
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

- Game starts with instructions on objective and how to land
- Each scene starts with an intro to its solar system
- Planets orbit their star (as opposed to just gliing across the screen)
- Mini map of the full solar system
- Docking sound
- Animate touching down tutorial
- Artificial satellites
- Different planet textures (e.g. Rocky, gaseous, watery, earthlike)
- worm holes (teleportation)
- more levels

# PROBABLY DON'T NEED

- Generalized Animate function - Use for sprite frames or tweening eg “pulse” behavior. How would Van Dur Spey do it? Sprite animations (separate component from sprite)

# DONE

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

# LIKELY IMPOSSIBLE

- Give GameObject direct access to components

```
// can't do this at runtime
// go is statically typed
func (o *GameObj) Sprite() *Sprite (
   return o.Components[“sprite”].(*Sprite)
}
```
