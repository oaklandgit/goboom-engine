package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// could also repel if force is negative
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

		text := fmt.Sprintf("%s: %d of %d", r.Name, r.Remaining, r.Amount)
		fontSize := int32(16)
		textWidth := rl.MeasureText(text, fontSize)

		rl.DrawText(
			text,
			int32(m.GameObj.Position.X) - int32(textWidth)/2,
			int32(m.GameObj.Position.Y) -
				int32((2 + i) * int(fontSize + 6)) -
				int32(m.GameObj.Height()/2),
			fontSize, rl.White)
		
	}

	// TEXT BELOW $SCORE
	text2 := fmt.Sprintf("%s @ $%d/unit", m.Resources[0].Name, m.Resources[0].Price)
	fontSize := int32(18)
	textWidth := rl.MeasureText(text2, fontSize)
	rl.DrawText(text2, 400 - textWidth/2, 46, fontSize, rl.White)
	
}