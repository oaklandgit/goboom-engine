# IN PROGRESS

- Collision with planet if too-fast or wrong-angle landing
- More feedback on state of landing (e.g. too fast!)

# TO DO

- Docking sound
- Different planet textures (e.g. Rocky, gaseous, watery, earthlike)
- Delay before respawn
- Game Start state
- Animate touching down tutorial

- Generalized Animate function - Use for sprite frames or tweening eg “pulse” behavior. How would Van Dur Spey do it? Sprite animations (separate component from sprite)
- Buggy when docked and moving past left or right edge of screen
- Game Over state
- Ship should attach to planet at correct landing point
- worm holes (teleportation)
- more levels

# DONE

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
