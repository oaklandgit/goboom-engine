package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// could also repel if force is negative
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
	l.GameObj.Components["motion"].(*Motion).SetVelocity(0, 0)
	l.GameObj.Components["sprite"].(*Sprite).CurrFrame = 0
	l.GameObj.Position = rl.NewVector2(400, 120)
	
	return l
}

func (l *Lives) RemoveLife() {
	l.Remaining--
	if l.Remaining <= 0 {
		fmt.Println("Game Over!")
		return
	}

	l.Respawn()
}

func (l *Lives) Update() {
	
}

func (l *Lives) Draw() {

	text := fmt.Sprintf("Lives: %d of %d", l.Remaining, l.Total)
	DrawText(text, 12, 12, 20, 2, rl.White, Left)
	
}