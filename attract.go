package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// could also repel if force is negative
type Attract struct {
	GameObj *GameObj
	Targets []*GameObj
	Ignored []string
	Force float32
	Threshold float32
}

func (*Attract) Id() string {
	return "attract"
}

type AttractOption func(*Attract)

func (obj *GameObj) NewAttract(
	targets []*GameObj,
	force float32,
	threshold float32,
	opts ...AttractOption) *Attract {

	attract := &Attract{
		Targets: targets,
		GameObj: obj,
		Force: force,
		Threshold: threshold,
	}

	for _, opt := range opts {
		opt(attract)
	}

	obj.AddComponents(attract)

	return attract
}

func WithIgnored(ignored... string) AttractOption {
	return func(a *Attract) {
		a.Ignored = ignored
	}
}


func (a *Attract) Update() {

	if len(a.Targets) == 0 {
		return
	}

	attract := func(t *GameObj) {
		
		p1 := a.GameObj.PosGlobal()
		p2 := t.PosGlobal()

		dist := rl.Vector2Distance(p2, p1)

		// if dist > a.Threshold {
		// 	return
		// }

		// strength := a.Force / (dist * dist) + 0.001 // avoid divide by zero
		strength := a.Force / dist

		dir := rl.Vector2Subtract(p1, p2)
		pull := rl.Vector2Scale(rl.Vector2Normalize(dir), strength)

		obj2Motion := t.Components["motion"].(*Motion)

		v2 := obj2Motion.Velocity
		obj2Motion.Velocity = rl.Vector2Add(v2, pull)
	}

	for _, target := range a.Targets {
		for _, ignore := range a.Ignored {
			if target.HasTag(ignore) {
				return
			}
		}
		attract(target)
	}

}

func (a *Attract) Draw() {
	// no op
}