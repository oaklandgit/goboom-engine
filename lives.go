package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const RESPAWN_BERTH = 22

type Lives struct {
	GameObj *GameObj
	Total int
	Remaining int
}

func (*Lives) Id() string {
	return "lives"
}

type LivesOption func(*Lives)

func (obj *GameObj) NewLives(total int, opts ...LivesOption) *Lives {

	lives := &Lives{
		GameObj: obj,
		Total: total,
		Remaining: total,
	}

	for _, opt := range opts {
		opt(lives)
	}

	obj.AddComponents(lives)

	return lives
}

func (l *Lives) AddLives(count int) *Lives {
	l.Remaining++
	return l
}

func (l *Lives) Respawn() *Lives {

	// stop the ship
	l.GameObj.Components["motion"].(*Motion).Velocity = rl.Vector2Zero()
	l.GameObj.Components["sprite"].(*Sprite).CurrFrame = 0
	l.GameObj.Components["dock"].(*Dock).Undock()

	randomSpot := func() rl.Vector2 {
		x := float32(rl.GetRandomValue(20, screenW - 20))
		y := float32(rl.GetRandomValue(20, screenH - 20))
		return rl.NewVector2(x, y)
	}

	safe := false
	maxAttempts := 1_000 // prevent a potential infinite loop
	attempts := 0

	safeLoop:
	for !safe && attempts < maxAttempts {

		// try a random spot
		l.GameObj.Position = randomSpot()

		for _, sibling := range l.GameObj.Parent.FindChildrenByTags(true, "deadly") {

			if rl.CheckCollisionCircles(
				l.GameObj.PosGlobal(),
				l.GameObj.Width() + RESPAWN_BERTH,
				sibling.PosGlobal(),
				sibling.Width() + RESPAWN_BERTH,
			) {
				attempts++
				continue safeLoop
			}
			
		}
		safe = true
	}

	return l
}

func (l *Lives) RemoveLife() {
	l.Remaining--
	if l.Remaining <= 0 {
		game.SetScene("gameover")
		
		// should move this to the gameover scene:
		music := sounds["sounds/gameover.wav"]
		rl.SetSoundVolume(music, 0.2);
		rl.PlaySound(music);
	} else {
		// l.ReadyToRespawn = true
		l.Respawn()
	}
	
}

func (l *Lives) Update() {
	// no op
}

func (l *Lives) Draw() {

	text := fmt.Sprintf("Lives: %d of %d", l.Remaining, l.Total)
	DrawText(text, 12, 12, 20, 2, rl.White, Left)
	
}