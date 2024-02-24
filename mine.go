package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mine struct {
	GameObj *GameObj
	MinedBy *GameObj
	Resources []Resource
}

type Resource struct {
	Name string
	Amount int
	Remaining int
	Price int // per unit
}

func (*Mine) Id() string {
	return "mine"
}

type MineOption func(*Mine)

func (obj *GameObj) NewMine(
	opts ...MineOption) *Mine {

	mine := &Mine{
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(mine)
	}

	obj.AddComponents(mine)

	return mine
}

func (m *Mine) AddResource(name string, amount int, price int) *Mine {
	m.Resources = append(m.Resources, Resource{
		Name: name,
		Amount: amount,
		Remaining: amount,
		Price: price,
	})

	return m
}

func (m *Mine) Update() {

	// deplete if docked
	// if !m.GameObj.HasTag("docking") { return }

	if m.MinedBy == nil { return }
	if len(m.Resources) == 0 { return }

	minerBank := m.MinedBy.Components["bank"].(*Bank)

	r := m.Resources[0]
	
	if r.Remaining > 0 {
		r.Remaining--
		minerBank.Balance += r.Price
		m.Resources[0] = r
	}

	// remove if depleted
	if r.Remaining == 0 {
		m.Resources = m.Resources[1:]
	}
}

func (m *Mine) Draw() {

	if m.MinedBy == nil { return }
	if len(m.Resources) == 0 { return }
	
	// TEXT ABOVE PLANET
	for i, r := range m.Resources {

		// draw progress bar

		text := fmt.Sprintf("%s: %d of %d", r.Name, r.Remaining, r.Amount)

		itemSpacing := 42
		progressBarWidth := 180
		nudgeLeft := m.GameObj.Width() / 2 + 8
		distFromPlanet := m.GameObj.Height() / 2 + 42

		DrawProgressBar(
			int32(m.GameObj.PosGlobal().X - float32(nudgeLeft)),
			int32(m.GameObj.PosGlobal().Y) -
				int32(i * itemSpacing) -
				int32(distFromPlanet),
			int32(progressBarWidth),
			int32(r.Amount - r.Remaining),
			int32(r.Amount),
			text,
		)

		// 
		// fontSize := int32(16)

		// x := 	int32(m.GameObj.Position.X)
		// y := 	int32(m.GameObj.Position.Y) -
		// 		int32((2 + i) * int(fontSize + 6)) -
		// 		int32(m.GameObj.Height()/2)

		// DrawText(text, x, y, fontSize, 2, rl.White, Left)
		
	}

	// TEXT BELOW $SCORE
	text2 := fmt.Sprintf("%s @ $%d/unit", m.Resources[0].Name, m.Resources[0].Price)
	DrawText(text2, screenW/2, 62, 18, 8, rl.Green, Center)
	
}