package main

import (
	boom "goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const RESPAWN_BERTH = 22

type Lives struct {
	GameObj *boom.GameObj
	Total int
	Remaining int
}

func (*Lives) Id() string {
	return "lives"
}

type LivesOption func(*Lives)

func NewLives(
	obj *boom.GameObj, 
	total int, opts ...LivesOption) *boom.GameObj {

	lives := &Lives{
		GameObj: obj,
		Total: total,
		Remaining: total,
	}

	for _, opt := range opts {
		opt(lives)
	}

	obj.AddComponents(lives)

	return obj
}

func (l *Lives) AddLives(count int) *Lives {
	l.Remaining++
	return l
}

func (l *Lives) Respawn() *Lives {

	// stop the ship
	l.GameObj.Components["motion"].(*boom.Motion).Velocity = rl.Vector2Zero()
	l.GameObj.Components["sprite"].(*boom.Sprite).CurrFrame = 0
	l.GameObj.Components["dock"].(*Dock).Undock()

	randomSpot := func() rl.Vector2 {
		x := float32(rl.GetRandomValue(0, int32(game.Width)))
		y := float32(rl.GetRandomValue(0, int32(game.Width)))
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
		

		// reset lives for next game
		l.Remaining = l.Total
		game.Reset()

		//
		game.SetScene("gameover")
		game.Scenes["gameover"].Components["timer"].(*boom.Timer).Start()

		// should move this to the gameover scene:
		music := game.Sounds["sounds/gameover.wav"]
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

	for i := 0; i < l.Remaining; i++ {
		rl.DrawTextureEx(
			game.Textures["assets/icon-life.png"],
			rl.NewVector2(game.Width - 52.0 - (float32(i) * 28.0), 22),
			0,
			0.6,
			rl.Green,
		)
	}
	
}